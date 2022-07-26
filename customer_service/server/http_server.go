package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"strconv"
)

func StartHTTPServer() (err error) {
	port := 8083
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	muxRouter := initRouter()

	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "User-Agent"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "PATCH", "DELETE"})

	corsHandler := handlers.CORS(headersOk, originsOk, methodsOk)(muxRouter)

	server := &http.Server{
		Addr:    addr,
		Handler: corsHandler,
	}

	log.Printf("starting API server on %s", addr)
	server.ListenAndServe()
	return
}
