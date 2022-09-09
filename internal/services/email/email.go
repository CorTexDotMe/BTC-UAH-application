package email

import (
	"btcApp/internal/utils"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
)

func InitializeDialer() *gomail.Dialer {
	host := "smtp.gmail.com"
	port := 587

	dialer := gomail.NewDialer(
		host,
		port,
		utils.AppEmail,
		utils.AppEmailPassword,
	)

	return dialer
}

func InitializeMessage(btcRate int) *gomail.Message {
	subject := "Subject: BTC-UAH rate\n"
	body := fmt.Sprintf("BTC rate in UAH: %d", btcRate)

	msg := gomail.NewMessage()
	msg.SetHeader("From", utils.AppEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)

	return msg
}

func SendEmail(email string, message *gomail.Message, dialer *gomail.Dialer) {
	message.SetHeader("To", email)
	if err := dialer.DialAndSend(message); err != nil {
		log.Printf("Unable to send message to \"%s\"\n", email)
	}
}
