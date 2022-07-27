package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getRateBTC(writer http.ResponseWriter, request *http.Request) {
	defer recoverInternalError(writer)

	btcUahRate := getBtcUahRate()
	returnBtcUahRate(writer, btcUahRate)
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
