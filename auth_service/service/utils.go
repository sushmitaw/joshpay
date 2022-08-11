package service

import (
	"encoding/base64"
	"github.com/golang-jwt/jwt"
	"time"
)

func decodeToken(token string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(token)
}

func encodeToken(token string) string {
	return base64.StdEncoding.EncodeToString([]byte(token))
}

func signToken(user LoginRequest) (tokenString string, err error) {

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(1 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: expirationTime.Unix(),
		Id:        user.username,
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err = token.SignedString(hmacSampleSecret)
	if err != nil {
		return
	}
	return
}
