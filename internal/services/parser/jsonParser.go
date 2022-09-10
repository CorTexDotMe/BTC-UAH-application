package parser

import (
	"btcApp/internal/common/utils"
	"encoding/json"
)

const IdOfUahInJsonFromRateService = 0

type coinsFromBtcRateService struct {
	Coins []coin
}

type coin struct {
	Price float64
}

func (r *coinsFromBtcRateService) getCoinPrice() float64 {
	return r.Coins[IdOfUahInJsonFromRateService].Price
}

func ParseJsonResponse(jsonWithRate []byte) int {
	var coinsFromService coinsFromBtcRateService
	parseJsonError := json.Unmarshal(jsonWithRate, &coinsFromService)
	utils.PanicIfUnexpectedErrorOccurs(parseJsonError)

	return int(coinsFromService.getCoinPrice())
}
