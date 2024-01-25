package connectors

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/nats-io/nats.go"
)

type Connector struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Url         string `json:"url"`
	JSON        any    `json:"json"`
	IsCurrent   bool   `json:"isCurrent"`

	// Name of the connector
}

type Options struct {
	Channel string
	Dir     string
	Env     []string
}

type MessageData struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	IsError   bool      `json:"isError"`
}
type Message struct {
	Channel string      `json:"channel"`
	Data    MessageData `json:"data"`
}

func CentrifugePost(channel string, data string, isError bool) error {
	// http post to localhost:8000/api/publish/centrifuge including X-API-Key: secret

	messageData := MessageData{
		Timestamp: time.Now(),
		Message:   data,
		IsError:   isError,
	}

	body := Message{
		Channel: channel,
		Data:    messageData,
	}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:8000/api/publish", bytes.NewReader(bodyJson))
	if err != nil {
		return err
	}

	req.Header.Add("X-API-Key", "913f84d9-797c-49e7-b2ac-8bacb40f7637")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return nil
}
func Execute(program string, options Options, args ...string) (output []byte, err error,
) {
	cmd := exec.Command(program, args...)
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()
	if options.Dir != "" {
		cmd.Dir = options.Dir
	}
	if options.Env != nil {
		cmd.Env = options.Env
	}
	pipe, _ := cmd.StdoutPipe()
	combinedOutput := []byte{}

	err = cmd.Start()
	go func(p io.ReadCloser) {
		reader := bufio.NewReader(pipe)
		line, err := reader.ReadString('\n')
		//
		// nc.Publish(program, []byte(line))
		// // if options.Channel != "" {
		// 	CentrifugePost(options.Channel, line, err == nil)
		// }
		for err == nil {
			//log.Print(line)
			combinedOutput = append(combinedOutput, []byte(line)...)
			line, err = reader.ReadString('\n')
		}
		//
	}(pipe)
	err = cmd.Wait()

	if err != nil {
		log.Println(fmt.Sprint(err) + ": " + string(combinedOutput))
		return nil, errors.New(fmt.Sprint(err))
	}

	return combinedOutput, nil
}
