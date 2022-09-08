package utils

import "net/http"

func RecoverInternalError(writer http.ResponseWriter) {
	if r := recover(); r != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func HandleUnexpectedError(err error) {
	if err != nil {
		panic(err)
	}
}
