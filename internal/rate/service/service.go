package service

import (
	"btcApp/internal/utils"
	"net/http"
)

func GetResponseFromBtcRateService() *http.Response {
	response, getRateError := http.Get(utils.BtcUahRateServiceUrl)
	if getRateError != nil {
		panic(getRateError)
	}
	return response
}
