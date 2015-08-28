package main

import (
	"io"
	"log"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

var chatroom = Chatroom{}

// ChatServer implements a websocket-based chatroom
func ChatServer(ws *websocket.Conn) {
	log.Printf("%p is joining", ws)
	chatroom.Add(ws)

	var message string
	for {
		if err := websocket.Message.Receive(ws, &message); err == io.EOF {
			log.Printf("%p is leaving", ws)
			break
		} else if err != nil {
			log.Printf("receive: %v", err)
			break
		}

		// since errors are being printed within Broadcast, we are ignoring
		// the error it is returning
		chatroom.Broadcast(message)
	}

	chatroom.Remove(ws)
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

type Chatroom struct {
	websockets []*websocket.Conn
	sync.Mutex
}

func (rm *Chatroom) Add(ws *websocket.Conn) {
	rm.Lock()
	defer rm.Unlock()

	rm.websockets = append(rm.websockets, ws)
}

func (rm *Chatroom) Remove(ws *websocket.Conn) {
	rm.Lock()
	defer rm.Unlock()

	for i, conn := range rm.websockets {
		if conn == ws {
			rm.websockets = append(rm.websockets[:i], rm.websockets[i+1:]...)
			break
		}
	}
}

func (rm *Chatroom) Broadcast(message string) error {
	rm.Lock()
	defer rm.Unlock()

	var lasterr error
	for _, ws := range rm.websockets {
		if err := websocket.Message.Send(ws, message); err != nil {
			log.Printf("send: %v", err)
			lasterr = err
		}
	}

	// returning only the last error encountered
	// in practice we should return a custom error that has all errors collected
	return lasterr
}
