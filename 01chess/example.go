package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// for {
	// 	messageType, p, err := conn.ReadMessage()
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// 	if err := conn.WriteMessage(messageType, p); err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// }
	for {
		messageType, r, err := conn.NextReader()
		if err != nil {
			return
		}
		w, err := conn.NextWriter(messageType)
		if err != nil {
			return
		}
		if _, err := io.Copy(w, r); err != nil {
			return
		}
		if err := w.Close(); err != nil {
			return
		}
	}
}
