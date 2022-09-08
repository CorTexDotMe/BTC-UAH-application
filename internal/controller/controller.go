package controller

import (
	"btcApp/internal/repository"
	"btcApp/internal/services/email"
	"btcApp/internal/services/rate"
	"btcApp/internal/utils"
	"log"
)

func SubscribeEmail(email string) bool {
	if repository.DB.Contains(email) {
		return false
	}

	databaseError := repository.DB.Add(email)
	utils.HandleUnexpectedError(databaseError)

	return true
}

func SendRateToEmails() {
	btcRate := rate.GetBtcRateInUah()
	initialMsg := email.InitializeMessage(btcRate)
	dialer := email.InitializeDialer()

	emails := repository.DB.GetAllEmails()
	for _, emailToSendRate := range emails {
		log.Println(emailToSendRate)
		email.SendEmail(emailToSendRate, initialMsg, dialer)
	}
}
