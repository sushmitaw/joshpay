package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strings"
)

func LoginService(req LoginRequest) (string, error) {
	return signToken(req)
}

var hmacSampleSecret = []byte("c2VjcmV0OktleQ==")

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
	log.Print("Authenticated using  jwt token")

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
	log.Print("Authenticated using base64 encoded token")

	return nil
}

func AuthJwtToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"])
	} else {
		return err
	}

	return nil

}
