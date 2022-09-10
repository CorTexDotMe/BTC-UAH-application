package repository

import (
	"btcApp/internal/common/utils"
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func (d *Database) Contains(email string) bool {
	file, openingError := os.Open(d.FullPath)
	utils.PanicIfUnexpectedErrorOccurs(openingError)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if fileScanner.Text() == email {
			return true
		}
	}
	utils.PanicIfUnexpectedErrorOccurs(fileScanner.Err())

	return false
}

func (d *Database) Add(email string) error {
	file, openingError := os.OpenFile(d.FullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	utils.PanicIfUnexpectedErrorOccurs(openingError)
	defer file.Close()

	var stringToWriteInFile string
	if d.firstElementAdded {
		stringToWriteInFile = "\n" + email
	} else {
		stringToWriteInFile = email
		d.firstElementAdded = true
	}
	_, writingError := file.WriteString(stringToWriteInFile)
	utils.PanicIfUnexpectedErrorOccurs(writingError)

	return nil
}

func (d *Database) GetAllEmails() []string {
	fileContent, readingError := ioutil.ReadFile(d.FullPath)
	utils.PanicIfUnexpectedErrorOccurs(readingError)

	stringWithAllEmails := string(fileContent)
	return strings.Split(stringWithAllEmails, "\n")
}
