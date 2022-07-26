package server

import (
	"github.com/gorilla/mux"
	"github.com/sushmitaw/joshpay/callback_service/service"
	"net/http"
)

func initRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/callback", service.CallbackHandler).Methods(http.MethodPost)

	return
}
