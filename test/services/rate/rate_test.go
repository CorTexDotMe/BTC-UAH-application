package rate

import (
	"btcApp/internal/services/rate"
	"btcApp/test/utils"
	"testing"
)

func TestRateServiceAvailability(t *testing.T) {
	if rate.GetResponseFromBtcRateService() != nil {
		utils.Success(t, "Response from rate service")
	} else {
		utils.Failure(t, "Response from rate service is nil")
	}
}
