package router

import (
	"btcApp/internal/handlers"
	"github.com/gorilla/mux"
)

func CreateInitialRouter() *mux.Router {
	router := basicRouterInit()
	setHandlers(router)
	return router
}

func basicRouterInit() *mux.Router {
	defaultRouter := mux.NewRouter()
	router := defaultRouter.Host("gses2.app").PathPrefix("/api").Subrouter()
	return router
}

func setHandlers(router *mux.Router) {
	router.HandleFunc("/rate", handlers.GetBtcRateInUah).Methods("GET")
	router.HandleFunc("/subscribe", handlers.SubscribeEmail).Methods("POST")
	router.HandleFunc("/sendEmails", handlers.SendRateToEmails).Methods("POST")
}
