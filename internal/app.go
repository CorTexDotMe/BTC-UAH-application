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
	router   *mux.Router
	database *repository.Database
}

func (a *App) Initialize() {
	a.database = &repository.Database{FullPath: utils.EmailsFilePath + utils.EmailsFileName}
	a.database.Initialize()

	a.router = router.CreateInitialRouter()
}

func (a *App) Run() {
	runService(a.router)
}

func runService(router *mux.Router) {
	log.Fatal(http.ListenAndServe(utils.Port, router))
}
