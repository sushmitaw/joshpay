package service

import (
	"encoding/json"
	auth "github.com/sushmitaw/joshpay/auth_service/service"
	"log"
	"net/http"
)

func CreateOrderHandler(rw http.ResponseWriter, req *http.Request) {

	err := auth.Authenticate(*req)
	if err != nil {
		log.Println("error in authentication", "error", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("authentication failed"))
		return
	}

	createOrderRequest := CreateOrderRequest{}
	err = json.NewDecoder(req.Body).Decode(&createOrderRequest)
	if err != nil {
		log.Println("error reading request body", "error", err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("invalid request body"))
		return
	}

	err = CreateOrderService(createOrderRequest)
	if err != nil {
		log.Println("error order service failed", "error", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("create order failed"))
		return
	}

	rw.Write([]byte("create order success"))
	return
}
