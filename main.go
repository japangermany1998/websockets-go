package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"log"
	"os"
)


var ws = neffos.New(websocket.DefaultGorillaUpgrader, neffos.Namespaces{
	"v1": neffos.Events{
	"chat": serverReceived,
	},
})

func serverReceived(c *neffos.NSConn, msg neffos.Message) error {
	log.Println(string(msg.Body))
	c.Conn.Server().Broadcast(nil, neffos.Message{
		Namespace: msg.Namespace,
		Room: msg.Room,
		Event: msg.Event,
		Body: msg.Body,
	})
	return nil
}

func clientReceive(c *neffos.NSConn, msg neffos.Message) error {
	fmt.Print("\r")
	log.Println(string(msg.Body))
	fmt.Print("Enter message: ")
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("expected program to start with 'server' or 'client' argument")
	}
	side := args[0]

	switch side {
	case "server":
		runServer()
	case "client1":
		runClient("room1", "client1", color.Yellow)
	case "client2":
		runClient("room1", "client2", color.Blue)
	default:
		log.Fatalf("unexpected argument, expected 'server' or 'client' but got '%s'", side)
	}
}

func runServer() {
	// [...]
	app := iris.New()

	app.Get("/websocket_endpoint", websocket.Handler(ws))
	log.Println("Serving websockets on http://localhost:8080/websocket_endpoint")
	log.Fatal(app.Listen(":8080"))
}

