package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

type server struct {
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	connection, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Fatal("Cannot accept web socket connection")
	}

	defer connection.CloseNow()

	err = echo(connection)
	if err != nil {
		log.Fatalf("Failed to echo message. %v", err)
	}
}

func echo(connection *websocket.Conn) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	messageType, reader, err := connection.Reader(ctx)
	if err != nil {
		return err
	}

	writer, err := connection.Writer(ctx, messageType)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}

	return writer.Close()
}
