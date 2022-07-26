package service

import (
	"errors"
	"net/http"
	"strings"
)

func LoginService(req LoginRequest) (string, error) {
	return encodeToken(validUser), nil
}

func Authenticate(req http.Request) error {
	token := req.Header.Get("Authorization")
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

	return nil
}
