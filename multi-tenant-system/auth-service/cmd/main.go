package main

import (
	"github.com/gin-gonic/gin"

	handlers "auth-service/internal/handler"
)

func main() {
	r := gin.Default()

	r.POST("/login", handlers.LoginHandler)
	r.GET("/validate", handlers.ValidateTokenHandler)

	r.Run(":8084")
}
