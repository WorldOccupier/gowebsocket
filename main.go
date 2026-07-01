package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Could not listen on 8080")
	}
	echoServer := &http.Server{
		Handler: server{},
	}
	log.Println("Server started")
	go SendMessage()

	err = echoServer.Serve(listener)
	if err != nil {
		log.Fatal("Could not serve web socket request")
	}
}
