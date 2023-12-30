package sockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Message is a object used to pass data on sockets.
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

// FindHandler is a type that defines handler finding functions.
type FindHandler func(Event) (Handler, bool)

// Client is a type that reads and writes on sockets.
type Client struct {
	send        Message
	socket      *websocket.Conn
	findHandler FindHandler
}

// NewClient accepts a socket and returns an initialized Client.
func NewClient(socket *websocket.Conn, findHandler FindHandler) *Client {
	return &Client{
		socket:      socket,
		findHandler: findHandler,
	}
}

// Write receives messages from the channel and writes to the socket.
func (c *Client) Write() {
	msg := c.send
	err := c.socket.WriteJSON(msg)
	if err != nil {
		log.Printf("socket write error: %v\n", err)
	}
}

// Read intercepts messages on the socket and assigns them to a handler function.
func (c *Client) Read() {
	var msg Message
	for {
		// read incoming message from socket
		if err := c.socket.ReadJSON(&msg); err != nil {
			log.Printf("socket read error: %v\n", err)
			break
		}
		// assign message to a function handler
		if handler, found := c.findHandler(Event(msg.Name)); found {
			handler(c, msg.Data)
		}
	}
	log.Println("exiting read loop")

	// close interrupted socket connection
	c.socket.Close()
}

// Handler is a type representing functions which resolve requests.
type Handler func(*Client, interface{})

// Event is a type representing request names.
type Event string

// Router is a message routing object mapping events to function handlers.
type Router struct {
	rules map[Event]Handler // rules maps events to functions.
}

// NewRouter returns an initialized Router.
func NewRouter() *Router {
	return &Router{
		rules: make(map[Event]Handler),
	}
}

// ServeHTTP creates the socket connection and begins the read routine.
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// configure upgrader
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// accept all?
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	// upgrade connection to socket
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("socket server configuration error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := NewClient(socket, rt.FindHandler)

	// running method for reading from sockets, in main routine
	client.Read()
}

// FindHandler implements a handler finding function for router.
func (rt *Router) FindHandler(event Event) (Handler, bool) {
	handler, found := rt.rules["helloFromClient"]
	return handler, found
}

// Handle is a function to add handlers to the router.
func (rt *Router) Handle(event Event, handler Handler) {
	// store in to router rules
	rt.rules[event] = handler

}

// helloFromClient is a method that handles messages from the app client.
func helloFromClient(c *Client, data interface{}) {
	log.Printf("hello from client! message: %v\n", data)

	// set and write response message
	c.send = Message{Name: "helloFromServer", Data: "hello client!"}
	c.Write()
}

func Serve() {
	// create router instance
	router := NewRouter()

	// handle events with messages named `helloFromClient` with handler
	// helloFromClient (from above).
	router.Handle("helloFromClient", helloFromClient)

	// handle all requests to /, upgrade to WebSocket via our router handler.
	http.Handle("/", router)

	// start server.
	http.ListenAndServe(":4000", nil)
}
