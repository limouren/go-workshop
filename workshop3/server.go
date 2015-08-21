package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var websockets = []*websocket.Conn{}

// ChatServer echoes the data received on the WebSocket.
func ChatServer(conn *websocket.Conn) {
	websockets = append(websockets, conn)

	var message string
	for {
		if err := websocket.Message.Receive(conn, &message); err != nil {
			log.Printf("receive: %v", err)
			break
		}

		for _, ws := range websockets {
			err := websocket.Message.Send(ws, message)
			if err != nil {
				log.Printf("send: %v", err)
			}
		}
	}

	// remove the websockets
	for i, ws := range websockets {
		if ws == conn {
			websockets = append(websockets[:i], websockets[i+1:]...)
		}
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
