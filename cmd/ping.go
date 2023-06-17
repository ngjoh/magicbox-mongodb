/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func pingNats() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Responding to a request message
	nc.Subscribe("request", func(m *nats.Msg) {
		m.Respond([]byte("answer is 42"))
	})

	// Simple Sync Subscriber
	sub, err := nc.SubscribeSync("foo")
	if err != nil {
		log.Fatal(err)
	}
	m, err := sub.NextMsg(10 * time.Second)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read message: %s\n", string(m.Data))
	}

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err = nc.ChanSubscribe("foo", ch)
	if err != nil {
		log.Fatal(err)
	}

	msg := <-ch
	fmt.Printf("Msg ch: %s\n", msg)
	// Unsubscribe
	sub.Unsubscribe()

	// Drain
	sub.Drain()

	// Requests
	msg, err = nc.Request("help", []byte("help me"), 10*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()

}

// accessCmd represents the access command
var pingCmd = &cobra.Command{
	Use:   "ping [service]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Verify connectivity to a service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]
		switch service {
		case "nats":
			pingNats()
		default:

			log.Fatalln("Unknown service", subject)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accessCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accessCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
