package internal

import (
	"btcApp/internal/common/utils"
	"github.com/joho/godotenv"
)

func StartService() {
	loadEnv()

	app := App{}
	app.Initialize()
	app.Run()
}

func loadEnv() {
	loadingError := godotenv.Load()
	utils.PanicIfUnexpectedErrorOccurs(loadingError)
}
