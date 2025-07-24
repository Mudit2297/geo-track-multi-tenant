package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// GenerateSecretHash creates the SECRET_HASH required by AWS Cognito
func GenerateSecretHash(username, clientID, clientSecret string) string {
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(username + clientID))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
