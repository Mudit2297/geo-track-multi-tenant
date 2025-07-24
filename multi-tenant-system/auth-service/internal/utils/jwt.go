package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type CognitoJWK struct {
	Keys []map[string]interface{} `json:"keys"`
}

func ValidateJWT(tokenString string, region, userPoolID string) (jwt.MapClaims, error) {
	jwksURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolID)
	resp, err := http.Get(jwksURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var jwks CognitoJWK
	json.Unmarshal(body, &jwks)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		kid := token.Header["kid"]
		for _, key := range jwks.Keys {
			if key["kid"] == kid {
				return jwt.ParseRSAPublicKeyFromPEM([]byte(key["x5c"].(string)))
			}
		}
		return nil, errors.New("key not found")
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot extract claims")
	}

	return claims, nil
}
