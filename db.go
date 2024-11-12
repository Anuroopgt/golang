package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitializeDB initializes the MySQL database connection
func InitializeDB() {
	var err error
	// Change this with your actual database credentials
	dsn := "root:password@tcp(localhost:3306)/coinmarketcap" // Replace with your credentials
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Ensure the database connection is alive
	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// Create table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS currency_prices (
		id INT AUTO_INCREMENT PRIMARY KEY,
		currency VARCHAR(10) NOT NULL,
		price DECIMAL(15, 8) NOT NULL,
		last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		UNIQUE (currency)
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	log.Println("Database initialized and table created (if not exists)")
}

// SaveOrUpdateCurrencyPrice saves or updates the currency price in the database
func SaveOrUpdateCurrencyPrice(currency string, price float64) error {
	// Query to insert or update the currency price
	query := `
		INSERT INTO currency_prices (currency, price)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE price = ?, last_updated = NOW();`

	_, err := db.Exec(query, currency, price, price)
	return err
}

// GetAllCurrencyPrices fetches all stored currency prices
func GetAllCurrencyPrices() ([]Currency, error) {
	rows, err := db.Query("SELECT id, currency, price, last_updated FROM currency_prices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currencies []Currency
	for rows.Next() {
		var c Currency
		if err := rows.Scan(&c.ID, &c.Currency, &c.Price, &c.LastUpdated); err != nil {
			return nil, err
		}
		currencies = append(currencies, c)
	}

	return currencies, nil
}

// Currency struct to hold the currency data
type Currency struct {
	ID         int
	Currency   string
	Price      float64
	LastUpdated string
}