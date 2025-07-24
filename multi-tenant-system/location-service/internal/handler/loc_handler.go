package handler

import (
	"fmt"
	"location-service/internal/model"
	"location-service/internal/operator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SubmitLocationHandler(c *gin.Context) {
	var req model.LocationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid payload - %v", err)})
		return
	}

	claimsTenantID, exists := c.Get("custom_tenant_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	req.TenantID = claimsTenantID.(string)
	err := operator.InsertLocation(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("could not store location - %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "location submitted"})
}
