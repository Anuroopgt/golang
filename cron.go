package main

import (
	"log"
	"github.com/robfig/cron/v3"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type CoinMarketCapResponse struct {
	Data map[string]struct {
		Symbol string `json:"symbol"`
		Quote  struct {
			USD struct {
				Price float64 `json:"price"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"data"`
}

func StartScheduledTask(apiKey string) {
	// Create a new cron job scheduler
	c := cron.New()

	// Define the task to fetch prices every 5 minutes
	_, err := c.AddFunc("@every 5m", func() {
		log.Println("Fetching currency prices...")
		coins, err := FetchCurrencyPrices(apiKey)
		if err != nil {
			log.Printf("Error fetching currency prices: %v", err)
			return
		}

		// Save or update each currency in the database
		for _, coin := range coins {
			err := SaveOrUpdateCurrencyPrice(coin.Symbol, coin.Quote.USD.Price)
			if err != nil {
				log.Printf("Error saving/updating currency %s: %v", coin.Symbol, err)
			}
		}
	})

	if err != nil {
		log.Fatalf("Error scheduling task: %v", err)
	}

	// Start the cron job
	c.Start()

	// Block the main thread so the cron job keeps running
	select {}
}

// FetchCurrencyPrices fetches current prices from the CoinMarketCap API
func FetchCurrencyPrices(apiKey string) ([]CoinMarketCapResponse, error) {
	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set the API key in the request header
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CoinMarketCapResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}