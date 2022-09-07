package controller

import (
	"btcApp/internal/rate"
	"btcApp/internal/repository"
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
	emails := repository.DB.GetAllEmails()

	for _, email := range emails {
		log.Print(email, btcRate)
	}
}
