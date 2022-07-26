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

	//invite user
	//router.Handle("/dashboard/users/invite",
	//	middleware.Adapt(
	//		dashboardregister.HandleUserInvite(dependencies.RegisterMerchantService),
	//		middleware.AuthenticateUser(
	//			config.JWTKey(), []constants.Role{
	//				constants.Admin,
	//				constants.Owner,
	//			}),
	//	),
	//).Methods(
	//	http.MethodPost,
	//)

	return
}
