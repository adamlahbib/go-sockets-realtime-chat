package subscriber

import (
	"log"

	"github.com/adamlahbib/go-storage-chat/internal/config/db"
	"github.com/go-redis/redis/v8"
)

func SubscribeAndSave(redisClient *redis.Client) {
	pubsub := redisClient.Subscribe(db.Ctx, "chat_channel")
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		err := redisClient.LPush(db.Ctx, "chat_history", msg.Payload).Err()
		if err != nil {
			log.Println("Failed to save message:", err)
		} else {
			log.Println("Message saved:", msg.Payload)
		}
	}
}
