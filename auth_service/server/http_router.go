package server

import (
	"github.com/gorilla/mux"
	"github.com/sushmitaw/joshpay/auth_service/service"
	"net/http"
)

func initRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/login", service.LoginHandler).Methods(http.MethodPost)

	return
}
