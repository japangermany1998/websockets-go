package main

import (
	"atomicgo.dev/cursor"
	"bufio"
	"context"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
	"log"
	"os"
)

//Client thực hiện lệnh khi nhận được tín hiệu từ server tại event "chat"
func clientReceive(c *neffos.NSConn, msg neffos.Message) error {
	fmt.Print("\r")
	log.Println(string(msg.Body))
	fmt.Print("Enter message: ")
	return nil
}

//khởi chạy client, xác định room để join, tự đặt tên và có thể chọn màu để highlight trên console
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

	//Khởi tạo instance client kết nối đến namespace cụ thể, ở đây là v1
	c, err := client.Connect(ctx, "v1")
	if err != nil {
		panic(err)
	}
	//Khởi tạo instance client join room
	r, err := c.JoinRoom(ctx, room)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello, I'm " + color.Colorize(colors, full_name) + "!\n\n")

out:
	for {
		select {
		//Trường hợp mất kết nối websocket
		case <-client.NotifyClose:
			log.Println(color.Colorize(color.Red, "Websocket connection closed"))
			break out
		//Trường hợp kết nối socket thành công
		default:
			//Nhập message
			fmt.Print("Enter message: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan() // use `for scanner.Scan()` to keep reading
			msg := scanner.Text()
			cursor.ClearLinesUp(1)

			//Client bắn message và gửi đến cho server tại room và namespace xác định và tại event "chat"
			r.Emit("chat", []byte(color.Colorize(colors, full_name+": ")+msg))
			break
		}
	}

}
