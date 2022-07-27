package service

import (
	"bufio"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"net/http"
	"os"
)

func sendRateToEmails(writer http.ResponseWriter, request *http.Request) {
	defer recoverInternalError(writer)

	host := "smtp.gmail.com"
	port := 587

	subject := "Subject: BTC-UAH rate\n"
	body := fmt.Sprintf("BTC rate in UAH: %d", getBtcUahRate())

	msg := gomail.NewMessage()
	msg.SetHeader("From", AppEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)

	dialer := gomail.NewDialer(
		host,
		port,
		AppEmail,
		AppEmailPassword,
	)

	file, openFileError := os.OpenFile(EmailsFileName, os.O_RDONLY|os.O_CREATE, 0)
	if openFileError != nil {
		panic(openFileError)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		to := fileScanner.Text()
		msg.SetHeader("To", to)
		if err := dialer.DialAndSend(msg); err != nil {
			log.Println(err)
		}
	}
}
