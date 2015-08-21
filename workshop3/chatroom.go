package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var websockets = []*websocket.Conn{}

// ChatServer implements a websocket-based chatroom
func ChatServer(ws *websocket.Conn) {
	var message string
	if err := websocket.Message.Receive(ws, &message); err != nil {
		log.Printf("receive: %v", err)
	}

	if err := websocket.Message.Send(ws, message); err != nil {
		log.Printf("send: %v", err)
	}
}

// This example demonstrates a trivial echo server.
func main() {
	// serve the client.html
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	wsSever := websocket.Server{
		Handler: ChatServer,
	}
	http.Handle("/ws", wsSever)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
