package service

import (
	"bufio"
	"net/http"
	"net/mail"
	"os"
)

func subscribeEmail(writer http.ResponseWriter, request *http.Request) {
	defer recoverInternalError(writer)

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

	file, openingError := os.OpenFile(EmailsFileName, os.O_RDWR|os.O_CREATE, 0600)
	if openingError != nil {
		panic(openingError)
	}
	defer file.Close()

	if alreadySubscribed(file, email) {
		writer.WriteHeader(http.StatusConflict)
	} else {
		addEmail(file, email)
		writer.WriteHeader(http.StatusOK)
	}
}

func alreadySubscribed(file *os.File, email string) bool {
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if fileScanner.Text() == email {
			return true
		}
	}

	if fileScanner.Err() != nil {
		panic(fileScanner.Err())
	}

	return false
}

func addEmail(file *os.File, email string) {
	_, writingError := file.WriteString(email + "\n")

	if writingError != nil {
		panic(writingError)
	}
}
