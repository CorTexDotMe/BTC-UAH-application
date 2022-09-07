package parser

import (
	"btcApp/internal/utils"
	"encoding/json"
)

type coinsFromBtcRateService struct {
	Coins []coin
}

type coin struct {
	Price float64
}

func (r *coinsFromBtcRateService) getCoinPrice() float64 {
	return r.Coins[utils.IdOfUahInJsonFromRateService].Price
}

func ParseJsonResponse(jsonWithRate []byte) int {
	var coinsFromService coinsFromBtcRateService
	parseJsonError := json.Unmarshal(jsonWithRate, &coinsFromService)
	if parseJsonError != nil {
		panic(parseJsonError)
	}

	return int(coinsFromService.getCoinPrice())
}
