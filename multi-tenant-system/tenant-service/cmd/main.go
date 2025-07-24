package main

import (
	"tenant-service/internal/client"
	"tenant-service/internal/handler"
	"tenant-service/internal/middleware"
	"tenant-service/internal/operator"

	"github.com/gin-gonic/gin"
)

func main() {
	db := client.InitDB()
	defer client.DB.Close()

	op := operator.NewTenantOperator(db)
	h := handler.NewTenantHandler(op)

	r := gin.Default()

	authServiceURL := "http://auth-service:8084" // Docker service name

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(authServiceURL))
	{
		protected.POST("/tenants", middleware.RoleMiddleware("admin"), h.CreateTenant)
		protected.GET("/tenants", middleware.RoleMiddleware("admin"), h.GetAllTenants)
		protected.GET("/tenant/:id", middleware.RoleMiddleware("admin"), h.GetTenantByID)
	}

	r.Run(":8081")
}
