package handlers

import (
	"btcApp/internal/handlers"
	"btcApp/internal/repository"
	"btcApp/test/utils"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestHandleRate(t *testing.T) {
	serverUrl := "http://gses2.app/api/rate"
	request := httptest.NewRequest(http.MethodGet, serverUrl, nil)
	recorder := httptest.NewRecorder()

	handlers.GetBtcRateInUah(recorder, request)
	result := recorder.Result()
	defer result.Body.Close()

	data, readingError := ioutil.ReadAll(result.Body)
	if readingError != nil {
		utils.Failure(t, "Error while reading result body")
	}

	parsedBtcRate, parsingError := strconv.Atoi(string(data))
	if parsingError != nil {
		utils.Failure(t, "Error while reading result body")
	} else {
		utils.Success(t, "Btc rate acquired: %d", parsedBtcRate)
	}
}

func TestHandleSubscribe(t *testing.T) {
	setUp()
	defer tearDown()

	serverUrl := "http://gses2.app/api/rate"
	testMail := "testMail@mail.com"

	data := url.Values{}
	data.Add("email", testMail)
	request := httptest.NewRequest(http.MethodPost, serverUrl, strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	recorder := httptest.NewRecorder()

	handlers.SubscribeEmail(recorder, request)

	if repository.DB.Contains(testMail) {
		utils.Success(t, "Database contains email after request")
	} else {
		utils.Failure(t, "No email added after /subscribe request")
	}
}

func TestHandleSendEmails(t *testing.T) {
	setUp()
	defer tearDown()

	serverUrl := "http://gses2.app/api/rate"
	request := httptest.NewRequest(http.MethodPost, serverUrl, nil)
	recorder := httptest.NewRecorder()

	testData := []string{
		"firstTestEmail",
		"secondTestEmail",
		"thirdTestEmail",
	}
	addTestDataToDatabase(testData)

	buf := prepareBufferForLogging()
	defer recoverSystemLogging()

	handlers.SendRateToEmails(recorder, request)

	if handlerLogContainsAllEmails(buf, testData) {
		utils.Success(t, "Handler send messages to all emails")
	} else {
		utils.Failure(t, "Handler missed some emails while sending messages")
	}
}

func setUp() {
	repository.DB = &repository.Database{FullPath: "testDatabase.txt"}
	repository.DB.Initialize()
}

func tearDown() {
	os.Remove(repository.DB.FullPath)
}

func addTestDataToDatabase(testData []string) {
	for _, email := range testData {
		repository.DB.Add(email)
	}
}

func prepareBufferForLogging() *bytes.Buffer {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	return &buf
}

func recoverSystemLogging() {
	log.SetOutput(os.Stderr)
}

func handlerLogContainsAllEmails(buf *bytes.Buffer, testData []string) bool {
	buffer := buf.String()
	for _, email := range testData {
		if !strings.Contains(buffer, email) {
			return false
		}
	}
	return true
}
