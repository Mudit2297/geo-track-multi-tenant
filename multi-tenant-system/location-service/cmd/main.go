package main

import (
	"location-service/internal/client"
	"location-service/internal/handler"
	"location-service/internal/middleware"
	"location-service/internal/streamer"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to streaming service via WebSocket
	streamer.Connect()

	client.InitDB()
	defer client.DB.Close()

	r := gin.Default()

	authServiceURL := "http://auth-service:8084" // Docker service name

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(authServiceURL))
	{
		protected.POST("/location", middleware.RoleMiddleware("tenant", "admin"), handler.SubmitLocationHandler)
	}

	r.Run(":8082")
}
