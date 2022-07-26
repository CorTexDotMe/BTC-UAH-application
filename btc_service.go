package main

import (
	"bufio"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const emailsFileName = "emails.txt"

func main() {
	router := mux.NewRouter()
	apiRouter := router.Host("gses2.app").PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/rate", getRateBTC).Methods("GET")
	apiRouter.HandleFunc("/subscribe", subscribeEmail).Methods("POST")

	log.Fatal(http.ListenAndServe(":80", apiRouter))
}

func getRateBTC(writer http.ResponseWriter, request *http.Request) {
	//Get json with bitcoin rate from third-party service
	response, getRateError := http.Get("https://api.coinstats.app/public/v1/coins?skip=0&limit=1&currency=UAH")
	if getRateError != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	//Read json from response
	body, readJsonError := ioutil.ReadAll(response.Body)
	if readJsonError != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	//Parse json from third-party service
	var coins map[string][]map[string]float64
	json.Unmarshal(body, &coins)
	rateBtcUah := int(coins["coins"][0]["price"])

	//Write response with btc-uah rate
	writeResponseError := json.NewEncoder(writer).Encode(rateBtcUah)
	if writeResponseError != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}

func subscribeEmail(writer http.ResponseWriter, request *http.Request) {
	parsingError := request.ParseForm()
	if parsingError != nil {
		writer.WriteHeader(http.StatusConflict)
		return
	}

	success := addEmail(request.Form.Get("email"))
	if !success {
		writer.WriteHeader(http.StatusConflict)
		return
	}
}

func addEmail(email string) bool {
	file, openingError := os.OpenFile(emailsFileName, os.O_RDWR|os.O_CREATE, 0600)
	if openingError != nil {
		return false
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if fileScanner.Text() == email {
			return false
		}
	}
	if fileScanner.Err() != nil {
		return false
	}

	if _, writingError := file.WriteString(email + "\n"); writingError != nil {
		return false
	}

	return true
}
