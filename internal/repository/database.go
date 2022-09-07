package repository

import (
	"btcApp/internal/utils"
	"log"
	"os"
)

var DB *Database

type Database struct {
	FullPath          string
	firstElementAdded bool
}

func (d *Database) Initialize() {
	database := createFile(d.FullPath)
	database.Close()

	d.firstElementAdded = false
}

func createFile(path string) *os.File {
	file, creatingError := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if creatingError != nil {
		log.Print("File not created")
	}

	return file
}

func InitializeDatabase() {
	DB = &Database{FullPath: utils.EmailsFilePath + utils.EmailsFileName}
	DB.Initialize()
}
