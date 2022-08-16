package main

import (
	"github.com/TwiN/go-color"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"log"
	"os"
)

//Khi server tiếp nhận message tại event "chat" sẽ truyền phát đến các client thuộc cùng 1 namespace, 1 room
func serverReceived(c *neffos.NSConn, msg neffos.Message) error {
	c.Conn.Server().Broadcast(nil, neffos.Message{
		Namespace: msg.Namespace,
		Room:      msg.Room,
		Event:     msg.Event,
		Body:      msg.Body,
	})
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("expected program to start with 'server' or 'client' argument")
	}
	side := args[0]

	//Khởi tạo server và client
	switch side {
	case "server":
		runServer()
	//client1 và client2 tại room1
	case "client1":
		runClient("room1", "client1", color.Yellow)
	case "client2":
		runClient("room1", "client2", color.Blue)
	//client3 và client4 tại room2
	case "client3":
		runClient("room2", "client3", color.Green)
	case "client4":
		runClient("room2", "client4", color.Cyan)
	default:
		log.Fatalf("unexpected argument, expected 'server' or 'client' but got '%s'", side)
	}
}

func runServer() {
	// [...]
	app := iris.New()

	var ws = neffos.New(websocket.DefaultGorillaUpgrader, neffos.Namespaces{
		"v1": neffos.Events{
			"chat": serverReceived,
		},
	})

	app.Get("/websocket_endpoint", websocket.Handler(ws))
	log.Println("Serving websockets on http://localhost:8080/websocket_endpoint")
	log.Fatal(app.Listen(":8080"))
}
