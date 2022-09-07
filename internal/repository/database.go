package repository

import (
	"btcApp/internal/utils"
	"log"
	"os"
)

var DB *Database

type Database struct {
	FullPath string
}

func (d *Database) Initialize() {
	database := createFile(d.FullPath)
	database.Close()
}

func createFile(path string) *os.File {
	file, creatingError := os.Create(path)
	if creatingError != nil {
		log.Print("File not created")
	}

	return file
}

func InitializeDatabase() {
	DB = &Database{FullPath: utils.EmailsFilePath + utils.EmailsFileName}
	DB.Initialize()
}
