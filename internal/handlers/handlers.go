package handlers

import (
	"btcApp/internal/rate"
	"btcApp/internal/utils"
	"encoding/json"
	"net/http"
)

func GetBtcRateInUah(writer http.ResponseWriter, request *http.Request) {
	defer utils.RecoverInternalError(writer)
	btcUahRate := rate.GetBtcRateInUah()

	returnBtcRateInUah(writer, btcUahRate)
}

func returnBtcRateInUah(writer http.ResponseWriter, btcUahRate int) {
	writeResponseError := json.NewEncoder(writer).Encode(btcUahRate)
	if writeResponseError != nil {
		panic(writeResponseError)
	}
}

func SubscribeEmail(writer http.ResponseWriter, request *http.Request) {
	//controller.SubscribeEmail()
}

func SendRateToEmails(writer http.ResponseWriter, request *http.Request) {
	//controller.SendRateToEmails()
}
