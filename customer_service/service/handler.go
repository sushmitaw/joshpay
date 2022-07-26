package service

import (
	"encoding/json"
	auth "github.com/sushmitaw/joshpay/auth_service/service"
	"log"
	"net/http"
)

func SaveCustomerDetailsHandler(rw http.ResponseWriter, req *http.Request) {

	err := auth.Authenticate(*req)
	if err != nil {
		log.Println("error in authentication", "error", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("authentication failed"))
		return
	}

	customerDetailsRequest := CustomerDetailsRequest{}
	err = json.NewDecoder(req.Body).Decode(&customerDetailsRequest)
	if err != nil {
		log.Println("error reading request body", "error", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("invalid request body"))
		return
	}

	err = SaveCustomerDetailsService(customerDetailsRequest)
	if err != nil {
		log.Println("error customer service failed", "error", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("save customer failed"))
		return
	}

	rw.Write([]byte("customer saved successfully"))
	return
}
