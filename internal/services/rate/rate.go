package rate

import (
	"btcApp/internal/common/constants"
	"btcApp/internal/common/utils"
	"net/http"
)

func GetResponseFromBtcRateService() *http.Response {
	response, getRateError := http.Get(constants.BtcUahRateServiceUrl)
	utils.PanicIfUnexpectedErrorOccurs(getRateError)
	return response
}
