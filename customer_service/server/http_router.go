package server

import (
	"github.com/gorilla/mux"
	"github.com/sushmitaw/joshpay/customer_service/service"
	"net/http"
)

func initRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/customer", service.SaveCustomerDetailsHandler).Methods(http.MethodPost)

	return
}
