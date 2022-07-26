package service

import (
	"encoding/json"
	auth "github.com/sushmitaw/joshpay/auth_service/service"
	"log"
	"net/http"
)

func InitiatePaymentHandler(rw http.ResponseWriter, req *http.Request) {

	auth.Authenticate(req)
	paymentRequest := PaymentRequest{}
	err := json.NewDecoder(req.Body).Decode(&paymentRequest)
	if err != nil {
		log.Fatal("error reading request body", "error", err.Error())
		rw.Write([]byte("invalid request body"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = InitiatePaymentService(paymentRequest)
	if err != nil {
		log.Fatal("error payment service failed", "error", err.Error())
		rw.Write([]byte("payment failed"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write([]byte("payment sucess"))
	rw.WriteHeader(http.StatusOK)
	return
}
