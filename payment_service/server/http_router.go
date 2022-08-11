package server

import (
	"github.com/gorilla/mux"
	"github.com/sushmitaw/joshpay/payment_service/service"
	"net/http"
)

func initRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/pay", service.InitiatePaymentHandler).Methods(http.MethodPost)

	return
}
