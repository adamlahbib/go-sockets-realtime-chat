package server

import (
	"log"
	"net/http"

	socket "github.com/adamlahbib/go-socket-chat/internal/websocket"
	"github.com/go-redis/redis/v8"
)

// StartServer initializes and starts the HTTP server
func StartServer(redisClient *redis.Client) error {
	// HTTP route for WebSocket connection
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		socket.HandleConnections(w, r, redisClient)
	})

	log.Println("WebSocket service started on :8080")
	return http.ListenAndServe(":8080", nil)
}
