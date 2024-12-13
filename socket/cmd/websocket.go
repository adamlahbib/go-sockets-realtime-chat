package main

import (
	"log"

	"github.com/adamlahbib/go-socket-chat/internal/config/db"
	"github.com/adamlahbib/go-socket-chat/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Initialize Redis client
	redisClient := db.InitializeRedis()

	// Start the HTTP server
	err := server.StartServer(redisClient)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
