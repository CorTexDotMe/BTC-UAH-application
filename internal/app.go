package internal

import (
	"btcApp/internal/repository"
	"btcApp/internal/router"
	"btcApp/internal/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	router *mux.Router
}

func (a *App) Initialize() {
	repository.InitializeDatabase()
	a.router = router.CreateInitialRouter()
}

func (a *App) Run() {
	runService(a.router)
}

func runService(router *mux.Router) {
	log.Fatal(http.ListenAndServe(utils.Port, router))
}
