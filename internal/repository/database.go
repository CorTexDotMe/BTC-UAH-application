package repository

import (
	"log"
	"os"
)

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
