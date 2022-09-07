package controller

import (
	"btcApp/internal/repository"
	"btcApp/internal/services/email"
	"btcApp/internal/services/rate"
	"log"
)

func SubscribeEmail(email string) bool {
	if repository.DB.Contains(email) {
		return false
	}

	databaseError := repository.DB.Add(email)
	if databaseError != nil {
		panic(databaseError)
	}
	return true
}

func SendRateToEmails() {
	btcRate := rate.GetBtcRateInUah()
	dialer := email.InitializeDialer()
	initialMsg := email.InitializeMessage(btcRate)

	emails := repository.DB.GetAllEmails()
	for _, emailToSendRate := range emails {
		log.Print(emailToSendRate, btcRate)
		email.SendEmail(emailToSendRate, initialMsg, dialer)
	}
}
