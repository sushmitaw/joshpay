package service

import (
	"encoding/json"
	auth "github.com/sushmitaw/joshpay/auth_service/service"
	"log"
	"net/http"
)

func InitiatePaymentHandler(rw http.ResponseWriter, req *http.Request) {

	err := auth.Authenticate(*req)
	if err != nil {
		log.Println("error in authentication", "error", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("authentication failed"))
		return
	}

	paymentRequest := PaymentRequest{}
	err = json.NewDecoder(req.Body).Decode(&paymentRequest)
	if err != nil {
		log.Println("error reading request body", "error", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("invalid request body"))
		return
	}

	err = InitiatePaymentService(paymentRequest)
	if err != nil {
		log.Println("error payment service failed", "error", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("payment failed"))
		return
	}

	rw.Write([]byte("payment success"))
	return
}
