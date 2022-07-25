package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type rate struct {
	BtcRate int `json:"rate"`
}

func getRateBTC(w http.ResponseWriter, r *http.Request) {
	tmpRate := 800000

	err := json.NewEncoder(w).Encode(rate{tmpRate})
	if err != nil {
		return
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/rate", getRateBTC).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
