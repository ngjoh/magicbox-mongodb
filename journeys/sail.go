package journeys

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func Sail() error {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// // Responding to a request message
	// nc.Subscribe("request", func(m *nats.Msg) {
	// 	m.Respond([]byte("answer is 42"))
	// })

	// // Simple Sync Subscriber
	// sub, err := nc.SubscribeSync("foo")
	// if err != nil {
	// 	return err
	// }
	// timeout := 5 * time.Second
	// m, err := sub.NextMsg(timeout)
	// if err != nil {
	// 	return err
	// }
	// m.Respond([]byte("answer is 41"))
	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe("foo", ch)
	if err != nil {
		return err
	}
	msg := <-ch
	msg.Respond([]byte("answer is 42"))
	// Unsubscribe
	sub.Unsubscribe()

	// Drain
	sub.Drain()

	// Requests
	msg, err = nc.Request("help", []byte("help me"), 10*time.Millisecond)
	if err != nil {
		return err
	}
	msg.Respond([]byte("answer is 43"))
	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()
	return nil
}
