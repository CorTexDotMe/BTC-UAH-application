package service

import (
	"btcApp/internal/common/constants"
	"net/http"
)

func GetResponseFromBtcRateService() *http.Response {
	response, getRateError := http.Get(constants.BtcUahRateServiceUrl)
	if getRateError != nil {
		panic(getRateError)
	}
	return response
}
