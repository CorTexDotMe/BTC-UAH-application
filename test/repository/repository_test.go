package repository

import (
	"btcApp/internal/repository"
	"btcApp/test/utils"
	"fmt"
	"os"
	"reflect"
	"testing"
)

var DB *repository.Database

func setUp() {
	DB = &repository.Database{FullPath: "test.txt"}
	DB.Initialize()
}

func tearDown() {
	os.Remove(DB.FullPath)
}

func TestContains(t *testing.T) {
	setUp()
	defer tearDown()

	testEmail := "testEmail@test.com"

	firstContains := DB.Contains(testEmail)
	if firstContains {
		utils.Failure(t, "Empty database has testEmail")
	}

	DB.Add(testEmail)

	containsAfterAdding := DB.Contains(testEmail)
	if containsAfterAdding {
		utils.Success(t, "Contains works fine")
	} else {
		utils.Failure(t, "Database after adding email doesn't have testEmail")
	}
}

func TestAdd(t *testing.T) {
	setUp()
	defer tearDown()

	data := []string{
		"testEmail1",
		"testEmail2",
		"testEmail3",
		"testEmail4",
		"testEmail5",
	}

	for _, email := range data {
		DB.Add(email)
	}

	for _, email := range data {
		if !DB.Contains(email) {
			utils.Failure(t, "Database doesn't contains email:%s after adding", email)
		}
	}

	utils.Success(t, "Adding works fine")
}

func TestGetAll(t *testing.T) {
	setUp()
	defer tearDown()

	data := []string{
		"testEmail1",
		"testEmail2",
		"testEmail3",
		"testEmail4",
		"testEmail5",
	}

	for _, email := range data {
		DB.Add(email)
	}

	dataInDatabase := DB.GetAllEmails()
	dataIsInDatabase := reflect.DeepEqual(data, dataInDatabase)

	if dataIsInDatabase {
		utils.Success(t, "GetAllEmails works fine")
	} else {
		actual := fmt.Sprint(dataInDatabase)
		expected := fmt.Sprint(data)
		utils.Failure(t, "\nGetAllEmails result is different from data entered.\n"+
			"\tActual: %s\n\tExpected: %s", actual, expected)
	}
}
