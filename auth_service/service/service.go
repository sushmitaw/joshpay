package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func LoginService(req LoginRequest) (string, error) {
	return encodeToken(validUser), nil
}

var hmacSampleSecret []byte

func Authenticate(req http.Request) error {
	tokenString := req.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["id"], claims["name"])
	} else {
		return err
	}

	return nil
}

func AuthEncodedToken(token string) error {
	tokenParams := strings.Split(token, " ")

	if tokenParams[0] != tokenPrefix {
		return errors.New("invalid token prefix")
	}

	decodedToken, err := decodeToken(tokenParams[1])
	if err != nil {
		return err
	}

	if string(decodedToken) != validUser {
		return errors.New("invalid token user")
	}
	return
}
