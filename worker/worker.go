package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/koksmat-com/koksmat/powershell"
	"github.com/nats-io/nats.go"
)

type PowerShellRequest struct {
	AppId      string `json:"appid"`
	ScriptName string `json:"scriptname"`
	Input      string `json:"input"`
}

type PowerShellResponse struct {
	AppId      string `json:"appid"`
	ScriptName string `json:"scriptname"`

	Input    string `json:"input"`
	Output   string `json:"result"`
	HasError bool   `json:"haserror"`
	Console  string `json:"console"`
}

func PowerShellExecute(appId string, scriptName string, input string) (PowerShellResponse, error) {
	request := &PowerShellRequest{}

	request.AppId = appId
	request.ScriptName = scriptName
	request.Input = input

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	response := PowerShellResponse{}
	r, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	m, err := nc.Request("powershell.exchange", r, 30*time.Second)
	if err != nil {
		return response, err
	}

	jsonErr := json.Unmarshal(m.Data, &response)
	if jsonErr != nil {
		return response, jsonErr
	}

	return response, nil
}

func Work() {

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.QueueSubscribe("powershell.exchange", "worker", func(m *nats.Msg) {
		request := PowerShellRequest{}
		response := PowerShellResponse{}
		jsonErr := json.Unmarshal(m.Data, &request)
		if jsonErr != nil {
			log.Fatal(jsonErr)
			return
		}

		outout, err, console := powershell.Execute("koksmat", request.ScriptName, request.Input, powershell.SetupExchange, "", powershell.CallbackMockup)

		response.Output = string(outout)
		response.AppId = request.AppId
		response.ScriptName = request.ScriptName
		response.Input = request.Input

		response.Console = console
		if err != nil {
			response.HasError = true
			response.Output = err.Error()

		} else {
			response.HasError = false
		}

		r, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
			return
		}

		m.Respond(r)
	})

	for {
		log.Println("Waiting for a message")
		time.Sleep(5 * time.Second)

	}
}

func Send(message string) error {

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()
	log.Println(fmt.Sprintf("Sending %s", message))
	m, err := nc.Request("foo", nil, 10*time.Second)

	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("Reply %s", m.Reply))
	return nil
}
