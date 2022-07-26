package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func LoginHandler(rw http.ResponseWriter, req *http.Request) {
	loginRequest := LoginRequest{}
	err := json.NewDecoder(req.Body).Decode(&loginRequest)
	if err != nil {
		log.Println("error reading request body", "error", err.Error())
		rw.Write([]byte("invalid request body"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := LoginService(loginRequest)
	if err != nil {
		log.Println("error auth service failed", "error", err.Error())
		rw.Write([]byte("login failed"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write([]byte(fmt.Sprintf("login sucess, your token is %s", token)))
	rw.WriteHeader(http.StatusOK)
	return
}
