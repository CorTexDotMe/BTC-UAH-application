package handlers

import (
	"btcApp/internal/controller"
	"btcApp/internal/services/rate"
	"btcApp/internal/utils"
	"encoding/json"
	"net/http"
	"net/mail"
)

func GetBtcRateInUah(writer http.ResponseWriter, request *http.Request) {
	defer utils.RecoverInternalError(writer)
	btcUahRate := rate.GetBtcRateInUah()
	writeBtcRateInResponse(writer, btcUahRate)
}

func SubscribeEmail(writer http.ResponseWriter, request *http.Request) {
	defer utils.RecoverInternalError(writer)
	parsingError := request.ParseForm()
	utils.HandleUnexpectedError(parsingError)

	email := request.Form.Get("email")
	if emailIsInvalid(email) {
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
	defer utils.RecoverInternalError(writer)
	controller.SendRateToEmails()
}

func writeBtcRateInResponse(writer http.ResponseWriter, btcUahRate int) {
	writeResponseError := json.NewEncoder(writer).Encode(btcUahRate)
	utils.HandleUnexpectedError(writeResponseError)
}

func emailIsInvalid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}
