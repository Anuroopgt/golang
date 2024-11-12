package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetAllPrices handles the API route for getting all stored prices
func GetAllPrices(c *gin.Context) {
	currencies, err := GetAllCurrencyPrices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from database"})
		return
	}

	c.JSON(http.StatusOK, currencies)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/get-all-prices", GetAllPrices)
	return r
}