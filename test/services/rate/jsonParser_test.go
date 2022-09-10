package rate

import (
	"btcApp/internal/services/parser"
	"btcApp/test/utils"
	"testing"
)

func TestJsonParser(t *testing.T) {
	data := []struct {
		outputAsCost int
		inputAsJson  string
	}{
		{-123445, "{\"coins\":[{\"id\":\"bitcoin\",\"icon\":\"https://static.coinstats.app/coins/1650455588819.png\",\"name\":\"Bitcoin\",\"symbol\":\"BTC\",\"rank\":1,\"price\":-123445.2424,\"priceBtc\":1,\"volume\":1587101186398.7573,\"marketCap\":13246639971606.812,\"availableSupply\":19143912,\"totalSupply\":21000000,\"priceChange1h\":-0.1,\"priceChange1d\":-5.88,\"priceChange1w\":-7.78,\"websiteUrl\":\"http://www.bitcoin.org\",\"twitterUrl\":\"https://twitter.com/bitcoin\",\"exp\":[\"https://blockchair.com/bitcoin/\",\"https://btc.com/\",\"https://btc.tokenview.com/\"]}]}"},
		{0, "{\"coins\":[{\"id\":\"bitcoin\",\"icon\":\"https://static.coinstats.app/coins/1650455588819.png\",\"name\":\"Bitcoin\",\"symbol\":\"BTC\",\"rank\":1,\"price\":0,\"priceBtc\":1,\"volume\":1587101186398.7573,\"marketCap\":13246639971606.812,\"availableSupply\":19143912,\"totalSupply\":21000000,\"priceChange1h\":-0.1,\"priceChange1d\":-5.88,\"priceChange1w\":-7.78,\"websiteUrl\":\"http://www.bitcoin.org\",\"twitterUrl\":\"https://twitter.com/bitcoin\",\"exp\":[\"https://blockchair.com/bitcoin/\",\"https://btc.com/\",\"https://btc.tokenview.com/\"]}]}"},
		{13131313, "{\"coins\":[{\"id\":\"bitcoin\",\"icon\":\"https://static.coinstats.app/coins/1650455588819.png\",\"name\":\"Bitcoin\",\"symbol\":\"BTC\",\"rank\":1,\"price\":13131313.13131313,\"priceBtc\":1,\"volume\":1587101186398.7573,\"marketCap\":13246639971606.812,\"availableSupply\":19143912,\"totalSupply\":21000000,\"priceChange1h\":-0.1,\"priceChange1d\":-5.88,\"priceChange1w\":-7.78,\"websiteUrl\":\"http://www.bitcoin.org\",\"twitterUrl\":\"https://twitter.com/bitcoin\",\"exp\":[\"https://blockchair.com/bitcoin/\",\"https://btc.com/\",\"https://btc.tokenview.com/\"]}]}"},
	}

	for _, testData := range data {
		result := parser.ParseJsonResponse([]byte(testData.inputAsJson))
		if result == testData.outputAsCost {
			utils.Success(t, "Cost %d parsed correctly", result)
		} else {
			utils.Failure(t, "Values parsed incorrectly.%d, expected: %d", result, testData.outputAsCost)
		}
	}
}
