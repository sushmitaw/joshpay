package service

import (
	"encoding/base64"
	"fmt"
)

func decodeToken(token string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(token)
}

func encodeToken(token string) string {
	token = fmt.Sprintf("%s:%s", tokenPrefix, token)
	return base64.StdEncoding.EncodeToString([]byte(token))
}
