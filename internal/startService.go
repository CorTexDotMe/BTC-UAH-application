package internal

import (
	"btcApp/internal/router"
	"btcApp/internal/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartService() {
	runService(router.CreateInitialRouter())
}

func runService(router *mux.Router) {
	log.Fatal(http.ListenAndServe(utils.Port, router))
}
