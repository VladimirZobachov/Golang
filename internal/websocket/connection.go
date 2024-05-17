package websocket

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
)

func SetupSocketIO() *socketio.Server {
	server := socketio.NewServer(nil)

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
		// server.Remove(s.ID())
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		// Add the Remove session id. Fixed the connection & mem leak
		//server.Remove(s.ID())
		fmt.Println("closed", reason)
	})

	return server
}
