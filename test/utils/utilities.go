package utils

import (
	"btcApp/internal/utils"
	"fmt"
	"testing"
)

func Success(t *testing.T, message string, args ...any) {
	t.Logf("\t%s\t%s", utils.Success, fmt.Sprintf(message, args...))
}

func Failure(t *testing.T, message string, args ...any) {
	t.Fatalf("\t%s\t%s", utils.Failed, fmt.Sprintf(message, args...))
}
