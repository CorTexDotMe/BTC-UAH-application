package repository

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func (d *Database) Contains(email string) bool {
	file, openingError := os.Open(d.FullPath)
	if openingError != nil {
		panic(openingError)
	}
	defer file.Close()

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

func (d *Database) Add(email string) error {
	file, openingError := os.OpenFile(d.FullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if openingError != nil {
		panic(openingError)
	}
	defer file.Close()

	_, writingError := file.WriteString(email + "\n")
	if writingError != nil {
		return writingError
	}

	return nil
}

func (d *Database) GetAllEmails() []string {
	fileContent, err := ioutil.ReadFile(d.FullPath)
	if err != nil {
		log.Print("error while reading")
	}

	return strings.Split(string(fileContent), "\n")
}
