package service

import (
	"encoding/base64"
)

func decodeToken(token string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(token)
}

func encodeToken(token string) string {
	return base64.StdEncoding.EncodeToString([]byte(token))
}
