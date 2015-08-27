package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var websockets = []*websocket.Conn{}

// ChatServer implements a websocket-based chatroom
func ChatServer(ws *websocket.Conn) {
	websockets = append(websockets, ws)

	var message string
	for {
		if err := websocket.Message.Receive(ws, &message); err != nil {
			log.Printf("receive: %v", err)
			break
		}

		for _, conn := range websockets {
			if err := websocket.Message.Send(conn, message); err != nil {
				log.Printf("send: %v", err)
				break
			}
		}
	}

	for i, conn := range websockets {
		if conn == ws {
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
