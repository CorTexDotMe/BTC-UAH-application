package routes

import (
	"btcApp/internal/handlers"
	"btcApp/internal/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartService() {
	runService(CreateInitialRouter())
}

func runService(router *mux.Router) {
	log.Fatal(http.ListenAndServe(utils.Port, router))
}

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
	router.HandleFunc("/rate", handlers.GetRateBTC).Methods("GET")
	router.HandleFunc("/subscribe", handlers.SubscribeEmail).Methods("POST")
	router.HandleFunc("/sendEmails", handlers.SendRateToEmails).Methods("POST")
}
