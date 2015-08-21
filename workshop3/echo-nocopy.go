package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

// EchoServer echoes the data received on the WebSocket.
func EchoServer(conn *websocket.Conn) {
	// re-implement echo server without io.Copy
}

// This example demonstrates a trivial echo server.
func main() {
	// serve the client.html
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	wsSever := websocket.Server{
		Handler: EchoServer,
	}
	http.Handle("/ws", wsSever)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
