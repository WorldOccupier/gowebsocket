package main

import (
	"context"
	"log"
	"time"

	"github.com/coder/websocket"
)

func SendMessage() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	connection, _, err := websocket.Dial(ctx, "ws://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer connection.CloseNow()
	err = connection.Write(ctx, websocket.MessageText, []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	_, resp, err := connection.Read(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("received %s\n", resp)
}
