package main

import (
	"context"
	"fmt"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

func GetStockNameAndPrice(symbol string, finnhubClient *finnhub.DefaultApiService) (StockPrice, error) {
	profile, _, err := finnhubClient.CompanyProfile2(context.Background()).Symbol(symbol).Execute()

	if err != nil {
		return StockPrice{}, err
	}

	quote, _, err := finnhubClient.Quote(context.Background()).Symbol(symbol).Execute()

	if err != nil {
		return StockPrice{}, err
	}

	currentPrice := fmt.Sprintf("$%.2f", *quote.C)
	name := *profile.Name

	stockPrice := StockPrice{
		Symbol: symbol,
		Name: name,
		Price: currentPrice,
	}

	return stockPrice, err
}