package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type rate struct {
	UahRate int `json:"rate"`
}

func getRateBTC(w http.ResponseWriter, r *http.Request) {
	//Get json with bitcoin rate from third-party service
	response, getRateError := http.Get("https://api.coinstats.app/public/v1/coins?skip=0&limit=1&currency=UAH")
	if getRateError != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	//Read json from response
	body, readJsonError := ioutil.ReadAll(response.Body)
	if readJsonError != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Parse json from third-party service
	var coins map[string][]map[string]float64
	json.Unmarshal(body, &coins)
	rateBtcUah := int(coins["coins"][0]["price"])

	//Write response with btc-uah rate
	writeResponseError := json.NewEncoder(w).Encode(rate{UahRate: rateBtcUah})
	if writeResponseError != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/rate", getRateBTC).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
