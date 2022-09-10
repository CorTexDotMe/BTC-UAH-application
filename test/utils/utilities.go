package utils

import (
	"btcApp/internal/common/constants"
	"fmt"
	"testing"
)

func Success(t *testing.T, message string, args ...any) {
	t.Logf("\t%s\t%s", constants.Success, fmt.Sprintf(message, args...))
}

func Failure(t *testing.T, message string, args ...any) {
	t.Fatalf("\t%s\t%s", constants.Failed, fmt.Sprintf(message, args...))
}
