package controller

import (
	"btcApp/internal/rate"
	"btcApp/internal/repository"
	"log"
)

func GetRateBTC() {
	rate.GetBtcRateInUah()
}

func SubscribeEmail() {
	error := repository.Add("")
	if error != nil {

	}
}

func SendRateToEmails() {
	btcRate := rate.GetBtcRateInUah()
	emails := repository.GetAllEmails()

	for _, email := range emails {
		log.Print(email, btcRate)
	}
}
