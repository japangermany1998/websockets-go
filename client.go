package main

import (
	"atomicgo.dev/cursor"
	"context"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
	"log"
)

func runClient(room, full_name, colors string) {
	ctx := context.Background()
	client, err := neffos.Dial(ctx, gorilla.DefaultDialer, "ws://localhost:8080/websocket_endpoint", neffos.Namespaces{
		"v1": neffos.Events{
			"chat": clientReceive,
		},
	})
	if err != nil {
		panic(err)
	}

	c, err := client.Connect(ctx, "v1")
	if err != nil {
		panic(err)
	}
	r, err := c.JoinRoom(ctx, room)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello, I'm " + color.Colorize(colors, full_name))
out:
	for {
		select {
		case <-client.NotifyClose:
			log.Println(color.Colorize(color.Red, "Websocket connection closed"))
			break out
		default:
			var msg string
			fmt.Print("Enter message: ")
			_, _ = fmt.Scan(&msg)
			cursor.ClearLinesUp(1)
			r.Emit("chat", []byte(color.Colorize(colors, full_name+": ")+msg))
			break
		}
	}

}
