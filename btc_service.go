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

const EmailsFileName = "emails.txt"
const Port = ":80"
const BtcUahRateUrl = "https://api.coinstats.app/public/v1/coins?skip=0&limit=1&currency=UAH"

func main() {
	router := mux.NewRouter()
	apiRouter := router.Host("gses2.app").PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/rate", getRateBTC).Methods("GET")
	apiRouter.HandleFunc("/subscribe", subscribeEmail).Methods("POST")
	apiRouter.HandleFunc("/sendEmails", sendRateToEmails).Methods("POST")

	log.Fatal(http.ListenAndServe(Port, apiRouter))
}

func getRateBTC(writer http.ResponseWriter, request *http.Request) {
	defer recoverGetRateError(writer)

	btcUahRate := getBtcUahRate()
	returnBtcUahRate(writer, btcUahRate)
}

func recoverGetRateError(writer http.ResponseWriter) {
	if r := recover(); r != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func getBtcUahRate() int {
	response := rateServiceHttpGet()
	defer response.Body.Close()

	return parseJson(readJson(response))
}

func rateServiceHttpGet() *http.Response {
	response, getRateError := http.Get(BtcUahRateUrl)
	if getRateError != nil {
		panic(getRateError)
	}
	return response
}

func readJson(response *http.Response) []byte {
	body, readJsonError := ioutil.ReadAll(response.Body)
	if readJsonError != nil {
		panic(readJsonError)
	}
	return body
}

func parseJson(jsonWithRate []byte) int {
	var coins map[string][]map[string]float64
	json.Unmarshal(jsonWithRate, &coins)
	return int(coins["coins"][0]["price"])
}

func returnBtcUahRate(writer http.ResponseWriter, btcUahRate int) {
	writeResponseError := json.NewEncoder(writer).Encode(btcUahRate)
	if writeResponseError != nil {
		panic(writeResponseError)
	}
}

func subscribeEmail(writer http.ResponseWriter, request *http.Request) {
	defer recoverSubscriptionError(writer)

	parsingError := request.ParseForm()
	if parsingError != nil {
		panic(parsingError)
	}

	email := request.Form.Get("email")

	file, openingError := os.OpenFile(EmailsFileName, os.O_RDWR|os.O_CREATE, 0600)
	if openingError != nil {
		panic(openingError)
	}
	defer file.Close()

	if alreadySubscribed(file, email) {
		writer.WriteHeader(http.StatusConflict)
	} else {
		addEmail(file, email)
		writer.WriteHeader(http.StatusOK)
	}
}

func recoverSubscriptionError(writer http.ResponseWriter) {
	if r := recover(); r != nil {
		writer.WriteHeader(http.StatusConflict)
	}
}

func alreadySubscribed(file *os.File, email string) bool {
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if fileScanner.Text() == email {
			return true
		}
	}

	if fileScanner.Err() != nil {
		panic(fileScanner.Err())
	}

	return false
}

func addEmail(file *os.File, email string) {
	_, writingError := file.WriteString(email + "\n")

	if writingError != nil {
		panic(writingError)
	}
}

func sendRateToEmails(writer http.ResponseWriter, request *http.Request) {

}
