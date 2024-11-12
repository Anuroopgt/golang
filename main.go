package main

import (
	"log"
	"os"
)

func main() {
	// Initialize the database
	InitializeDB()

	// Start the cron job to fetch prices every 5 minutes
	apiKey := "368d2b03-8cb2-4927-8c3a-85a237e83352" // Replace with your actual API key
	go StartScheduledTask(apiKey)

	// Set up the Gin server and run it
	router := SetupRouter()
	err := router.Run(":8080") // Port 8080
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}