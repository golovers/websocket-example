package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("error: failed to upgrades the HTTP server connection to the WebSocket protocol: %v\n", err)
			return
		}
		for {
			typ, msg, err := conn.ReadMessage()
			if err != nil {
				log.Printf("error: failed to read message from client: %v\n", err)
				return
			}
			if err := conn.WriteMessage(typ, msg); err != nil {
				log.Printf("error: failed to write message to client: %v\n", err)
				return
			}
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":8080", nil)
}
