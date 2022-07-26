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
	decodedToken, err := decodeToken(token)
	if err != nil {
		return err
	}
	tokenParams := strings.Split(string(decodedToken), ":")

	if tokenParams[0] != tokenPrefix {
		return errors.New("invalid token prefix")
	}

	if tokenParams[1] != validUser {
		return errors.New("invalid token user")
	}

	return nil
}
