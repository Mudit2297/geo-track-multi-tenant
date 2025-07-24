package main

import (
	"log"
	"streaming-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		handler.HandleConnections(c.Writer, c.Request)
	})

	// Start broadcasting in the background
	go handler.HandleBroadcast()

	log.Println("Streaming service running on :8082")
	if err := r.Run(":8083"); err != nil {
		log.Fatalf("Gin server failed: %v", err)
	}
}
