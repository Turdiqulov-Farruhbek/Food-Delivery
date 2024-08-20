package main

import (
	// "fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader to upgrade HTTP connections to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Struct to hold WebSocket clients and a broadcast channel
type Server struct {
	clients   map[*websocket.Conn]bool
	broadcast chan Message
}

// Message structure
type Message struct {
	Content string `json:"content"`
}

func main() {
	server := &Server{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan Message),
	}

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ws", server.handleConnections)

	go server.handleMessages()

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func (server *Server) handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	server.clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(server.clients, ws)
			break
		}

		server.broadcast <- msg
	}
}

func (server *Server) handleMessages() {
	for {
		msg := <-server.broadcast
		for client := range server.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(server.clients, client)
			}
		}
	}
}
