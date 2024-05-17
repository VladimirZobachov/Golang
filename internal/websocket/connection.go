package websocket

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"time"
)

func SetupSocketIO() *socketio.Server {
	options := &engineio.Options{
		PingInterval: 25 * time.Second, // How often a ping is sent
		PingTimeout:  60 * time.Second, // How long to wait for a ping response before considering the connection closed
	}

	server := socketio.NewServer(options)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "join_department", func(s socketio.Conn, departmentID string) {
		s.Join(departmentID)
		fmt.Println("joined department: ", departmentID)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	return server
}
