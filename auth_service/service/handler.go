package service

import (
	"encoding/json"
	"log"
	"net/http"
)

func LoginHandler(rw http.ResponseWriter, req *http.Request) {
	loginRequest := LoginRequest{}
	err := json.NewDecoder(req.Body).Decode(&loginRequest)
	if err != nil {
		log.Fatal("error reading request body", "error", err.Error())
		rw.Write([]byte("invalid request body"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = LoginService(loginRequest)
	if err != nil {
		log.Fatal("error auth service failed", "error", err.Error())
		rw.Write([]byte("login failed"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write([]byte("login sucess"))
	rw.WriteHeader(http.StatusOK)
	return
}
