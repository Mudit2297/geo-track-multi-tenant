package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tenant-service/internal/helper"

	"github.com/gin-gonic/gin"
)

type ValidateRequest struct {
	Token string `json:"token"`
}

type ValidateResponse struct {
	Claims struct {
		Aud             string `json:"aud"`
		AuthTime        int    `json:"auth_time"`
		CognitoUsername string `json:"cognito:username"`
		CustomRole      string `json:"custom:role"`
		CustomTenantID  string `json:"custom:tenant_id"`
		Email           string `json:"email"`
		EmailVerified   bool   `json:"email_verified"`
		EventID         string `json:"event_id"`
		Exp             int    `json:"exp"`
		Iat             int    `json:"iat"`
		Iss             string `json:"iss"`
		Jti             string `json:"jti"`
		Name            string `json:"name"`
		OriginJti       string `json:"origin_jti"`
		Sub             string `json:"sub"`
		TokenUse        string `json:"token_use"`
	} `json:"claims"`
}

func AuthMiddleware(authServiceURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := authServiceURL + "/validate"
		method := "GET"

		body, err := helper.ExecHttpRequest(method, url, c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Error in executing request - %v", err)})
			return
		}

		var claims ValidateResponse
		if err := json.Unmarshal(body, &claims); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unable to parse token claims"})
			return
		}

		// Save claims in context
		c.Set("claims", claims)
		c.Set("custom_role", claims.Claims.CustomRole)
		c.Set("custom_tenant_id", claims.Claims.CustomTenantID)
		c.Next()
	}
}

func RoleMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsRole, exists := c.Get("custom_role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		role := claimsRole.(string)
		roleMatch := false

		for _, r := range requiredRoles {
			if r == role {
				roleMatch = true
				c.Next()
				return
			}
		}

		if !roleMatch {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("Forbidden: insufficient privileges, no role match. Current role is: %v", role)})
			return
		}
	}
}
