package utils

import (
	"btcApp/internal/utils"
	"fmt"
	"testing"
)

var ServiceForTestingMailingApiKey = "7f69f952d63214a853e4e43483cef8b3fd481185876092222dba82ca1316e092"

func Success(t *testing.T, message string, args ...any) {
	t.Logf("\t%s\t%s", utils.Success, fmt.Sprintf(message, args...))
}

func Failure(t *testing.T, message string, args ...any) {
	t.Fatalf("\t%s\t%s", utils.Failed, fmt.Sprintf(message, args...))
}
