package handlers

import (
	"btcApp/internal/controller"
	"btcApp/internal/rate"
	"btcApp/internal/utils"
	"encoding/json"
	"net/http"
	"net/mail"
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
	defer utils.RecoverInternalError(writer)
	parsingError := request.ParseForm()
	if parsingError != nil {
		panic(parsingError)
	}

	email := request.Form.Get("email")
	_, err := mail.ParseAddress(email)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if controller.SubscribeEmail(email) {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusConflict)
	}
}

func SendRateToEmails(writer http.ResponseWriter, request *http.Request) {
	//controller.SendRateToEmails()
}
