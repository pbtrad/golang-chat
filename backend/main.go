package main

import (
	"fmt"
	"net/http"

	"github.com/pbtrad/golang-chat/pkg/websocket"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request)

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	fmt.Println("Paul's chat!")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
