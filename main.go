package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var wsupgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool{
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)

    if err!= nil {
        return
    }

    defer conn.Close()
    clients[conn] = true
    
	for {
		_, msg, err := conn.ReadMessage()

		if err!= nil {
            delete(clients, conn)
            break
        }

		broadcast <- string(msg)
	}
}


