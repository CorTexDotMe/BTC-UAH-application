package service

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const Port = ":80"
const BtcUahRateUrl = "https://api.coinstats.app/public/v1/coins?skip=0&limit=1&currency=UAH"
const EmailsFileName = "emails.txt"
const AppEmail = "gses2.app.nechyporchuk@gmail.com"
const AppEmailPassword = "hriendevvanbzuvg"

func StartService() {
	router := mux.NewRouter()
	apiRouter := router.Host("gses2.app").PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/rate", getRateBTC).Methods("GET")
	apiRouter.HandleFunc("/subscribe", subscribeEmail).Methods("POST")
	apiRouter.HandleFunc("/sendEmails", sendRateToEmails).Methods("POST")

	log.Fatal(http.ListenAndServe(Port, apiRouter))
}

func recoverInternalError(writer http.ResponseWriter) {
	if r := recover(); r != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
