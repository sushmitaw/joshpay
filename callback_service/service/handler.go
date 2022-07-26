package service

import (
	"encoding/json"
	auth "github.com/sushmitaw/joshpay/auth_service/service"
	"log"
	"net/http"
)

func CallbackHandler(rw http.ResponseWriter, req *http.Request) {

	err := auth.Authenticate(*req)
	if err != nil {
		log.Println("error in authentication", "error", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("authentication failed"))
		return
	}

	callbackRequest := CallBackRequest{}
	err = json.NewDecoder(req.Body).Decode(&callbackRequest)
	if err != nil {
		log.Println("error reading request body", "error", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("invalid request body"))
		return
	}

	err = ProcessCallbackService(callbackRequest)
	if err != nil {
		log.Println("error callback service failed", "error", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("callback status failed"))
		return
	}

	rw.Write([]byte("callback success"))
	return
}
