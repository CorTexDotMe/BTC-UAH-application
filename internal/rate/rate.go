package rate

import (
	"btcApp/internal/rate/parser"
	"btcApp/internal/rate/service"
	"io/ioutil"
	"net/http"
)

func GetBtcRateInUah() int {
	response := service.GetResponseFromBtcRateService()
	defer closeResponseBody(response)

	responseAsJson := readJson(response)
	BtcRateInUah := parser.ParseJsonResponse(responseAsJson)
	return BtcRateInUah
}

func closeResponseBody(response *http.Response) {
	closingError := response.Body.Close()
	if closingError != nil {
		panic(closingError)
	}
}

func readJson(response *http.Response) []byte {
	body, readJsonError := ioutil.ReadAll(response.Body)
	if readJsonError != nil {
		panic(readJsonError)
	}
	return body
}
