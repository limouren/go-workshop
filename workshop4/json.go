package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type Message struct {
	Event string
	Data  MessageData
}

type MessageData struct {
	Username string
	Message  string
}

// ChatServer implements a websocket-based chatroom
func ChatServer(ws *websocket.Conn) {
	message := Message{
		Event: "some event",
		Data: MessageData{
			Username: "some username",
			Message:  "Here is a message :D",
		},
	}

	if err := websocket.JSON.Send(ws, message); err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}

// This example demonstrates a trivial echo server.
func main() {
	// serve the client.html
	http.Handle("/", http.FileServer(http.Dir("static")))

	wsSever := websocket.Server{
		Handler: ChatServer,
	}
	http.Handle("/ws", wsSever)

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
