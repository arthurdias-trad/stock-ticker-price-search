package main

import (
	"fmt"
	"html/template"
	"os"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var quoteTemplate = template.Must(template.ParseFiles("./templates/quote_template.html"))
var formTemplate = template.Must(template.ParseFiles("./templates/form_template.html"))

func main () {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found. Proceeding with environment variables.")
	}
	
	PORT := os.Getenv("PORT")
	FINNHUB_API_KEY := os.Getenv("FINNHUB_API_KEY")

	cfg := finnhub.NewConfiguration()
    cfg.AddDefaultHeader("X-Finnhub-Token", FINNHUB_API_KEY)
    finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

	router := gin.Default()

	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
        c.File("./static/index.html")
    })

	router.GET("/single-quote", GetFormHandler(TemplateData{IsMultiQuote: false}))
	router.GET("/multi-quote", GetFormHandler(TemplateData{IsMultiQuote: true}))

	router.GET("/quote", GetQuoteHandler(finnhubClient))
	router.GET("/quotes", GetQuotesHandler(finnhubClient))

	fmt.Printf("Starting server at port %s\n", PORT)
	router.Run(fmt.Sprintf(":%s", PORT))
}