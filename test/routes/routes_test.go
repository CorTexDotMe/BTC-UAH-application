package routes

import (
	"btcApp/internal/routes"
	"btcApp/test/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

type request struct {
	route        string
	method       string
	expectedCode int
}

func TestRouter(t *testing.T) {
	router := routes.CreateInitialRouter()
	serverUrl := "http://gses2.app/api"

	data := []request{
		{
			"/rate",
			http.MethodGet,
			http.StatusOK,
		},
		{
			"/subscribe",
			http.MethodPost,
			http.StatusOK,
		},
		{
			"/sendEmails",
			http.MethodPost,
			http.StatusOK,
		},
		{
			"/wrong",
			http.MethodGet,
			http.StatusNotFound,
		},
	}

	for _, testRequestData := range data {
		request := httptest.NewRequest(testRequestData.method, serverUrl+testRequestData.route, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		if recorder.Code == testRequestData.expectedCode {
			utils.Success(t, "%s returned code %d", testRequestData.route, testRequestData.expectedCode)
		} else {
			utils.Failure(
				t,
				"%s. Expected code %d , result: code %d",
				testRequestData.route,
				testRequestData.expectedCode,
				recorder.Code,
			)
		}
	}
}
