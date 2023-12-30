package sockets

import (
	"log"

	"github.com/labstack/echo"

	socketio "github.com/googollee/go-socket.io"
)

func SocketIO() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")

		log.Println("connected:", s.ID(), s.Namespace())
		log.Println(server.Count(), "connected clients")
		//server.BroadcastToNamespace(s.Namespace(), "notice", s.ID()+" joined the room")
		//log.Println("broadcasted to namespace")
		return nil
	})

	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
		server.BroadcastToNamespace(s.Namespace(), "notice", s.ID()+" said "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "echo", func(s socketio.Conn, msg interface{}) {
		s.Emit("echo", msg)
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	e := echo.New()
	e.HideBanner = true

	e.Static("/", "../asset")
	e.Any("/socket.io/", func(context echo.Context) error {
		server.ServeHTTP(context.Response(), context.Request())
		return nil
	})
	e.Logger.Fatal(e.Start(":4000"))
}
