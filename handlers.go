package main

import (
	"strings"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/gin-gonic/gin"
)


func GetQuoteHandler(finnhubClient *finnhub.DefaultApiService) gin.HandlerFunc {
    return func(c *gin.Context) {
        symbol := c.Query("symbol")
		stockPrice, err := GetStockNameAndPrice(symbol, finnhubClient)

        if err != nil {
            c.JSON(500, gin.H{
                "error": err.Error(),
            })
            return
        }

		result := []StockPrice{stockPrice}
		
		c.Header("Content-Type", "text/html")
		
		err = quoteTemplate.Execute(c.Writer, result)
		if err != nil {
            c.JSON(500, gin.H{
                "error": err.Error(),
            })
            return
        }
	}
}

func GetQuotesHandler(finnhubClient *finnhub.DefaultApiService) gin.HandlerFunc {
	return func(c *gin.Context) {
		symbolsQuery := c.Query("symbols")
		
		
		symbols := strings.Split(symbolsQuery, ",")
		var stockPrices []StockPrice

		for _, symbol := range symbols {
			
			stockPrice, err := GetStockNameAndPrice(strings.TrimSpace(symbol), finnhubClient)

			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			stockPrices = append(stockPrices, stockPrice)
		}

		err := quoteTemplate.Execute(c.Writer, stockPrices)
		if err != nil {
            c.JSON(500, gin.H{
                "error": err.Error(),
            })
            return
        }

	}
}

func GetFormHandler(quoteType TemplateData) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		err := formTemplate.Execute(c.Writer, quoteType)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	
}