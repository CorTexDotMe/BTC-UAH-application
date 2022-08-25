package routes

import (
	"btcApp/internal/handlers"
	"btcApp/internal/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var router *mux.Router

func StartService() {
	initializeRouter()
	setHandlers()
	runService()
}

func initializeRouter() {
	defaultRouter := mux.NewRouter()
	router = defaultRouter.Host("gses2.app").PathPrefix("/api").Subrouter()
}

func setHandlers() {
	router.HandleFunc("/rate", handlers.GetRateBTC).Methods("GET")
	router.HandleFunc("/subscribe", handlers.SubscribeEmail).Methods("POST")
	router.HandleFunc("/sendEmails", handlers.SendRateToEmails).Methods("POST")
}

func runService() {
	log.Fatal(http.ListenAndServe(utils.Port, router))
}
