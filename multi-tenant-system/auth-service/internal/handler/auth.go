package handlers

import (
	"auth-service/internal/utils"
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login payload"})
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		log.Printf("error loading config: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AWS config error"})
		return
	}

	secretHash := utils.GenerateSecretHash(req.Username, os.Getenv("COGNITO_CLIENT_ID"), os.Getenv("COGNITO_CLIENT_SECRET"))

	cognito := cognitoidentityprovider.NewFromConfig(cfg)
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		ClientId: &[]string{os.Getenv("COGNITO_CLIENT_ID")}[0],
		AuthParameters: map[string]string{
			"USERNAME":    req.Username,
			"PASSWORD":    req.Password,
			"SECRET_HASH": secretHash,
		},
	}

	authResp, err := cognito.InitiateAuth(context.TODO(), authInput)
	if err != nil {
		log.Printf("login failed: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": authResp.AuthenticationResult})
}

func ValidateTokenHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Decode the token header to get the kid
	parts := strings.Split(tokenString, ".")
	if len(parts) < 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	headerBytes, _ := base64.RawURLEncoding.DecodeString(parts[0])
	var tokenHeader struct {
		Kid string `json:"kid"`
	}
	_ = json.Unmarshal(headerBytes, &tokenHeader)

	jwksURL := os.Getenv("COGNITO_JWKS_URL")
	jwks, err := utils.GetCognitoJWKs(jwksURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch JWKS"})
		return
	}

	// Find the key with matching kid
	var publicKey interface{}
	for _, key := range jwks.Keys {
		if key.Kid == tokenHeader.Kid {
			pubKey, err := utils.KeyToPublicKey(key)
			if err == nil {
				publicKey = pubKey
			}
			break
		}
	}

	if publicKey == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Public key not found"})
		return
	}

	// Parse and verify the token
	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil || !parsedToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"claims": parsedToken.Claims})
}
