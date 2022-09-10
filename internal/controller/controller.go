package controller

import (
	"btcApp/internal/common/utils"
	"btcApp/internal/repository"
	"btcApp/internal/services/email"
	"btcApp/internal/services/parser"
	"btcApp/internal/services/rate"
	"io/ioutil"
	"log"
	"net/http"
)

func GetBtcRateInUah() int {
	response := rate.GetResponseFromBtcRateService()
	defer closeResponseBody(response)

	responseAsJson := readJson(response)
	BtcRateInUah := parser.ParseJsonResponse(responseAsJson)
	return BtcRateInUah
}

func closeResponseBody(response *http.Response) {
	closingError := response.Body.Close()
	utils.PanicIfUnexpectedErrorOccurs(closingError)
}

func readJson(response *http.Response) []byte {
	body, readJsonError := ioutil.ReadAll(response.Body)
	utils.PanicIfUnexpectedErrorOccurs(readJsonError)
	return body
}

func SubscribeEmail(email string) bool {
	if repository.DB.Contains(email) {
		return false
	}

	databaseError := repository.DB.Add(email)
	utils.PanicIfUnexpectedErrorOccurs(databaseError)

	return true
}

func SendRateToEmails() {
	btcRate := GetBtcRateInUah()
	initialMsg := email.InitializeMessage(btcRate)
	dialer := email.InitializeDialer()

	emails := repository.DB.GetAllEmails()
	for _, emailToSendRate := range emails {
		log.Println(emailToSendRate)
		email.SendEmail(emailToSendRate, initialMsg, dialer)
	}
}
