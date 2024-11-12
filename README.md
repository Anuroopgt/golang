# CoinMarketCap Currency Price Updater

This is a Go-based microservice that fetches cryptocurrency prices from the CoinMarketCap API and updates a local database with the latest prices. Additionally, it provides an API endpoint to retrieve all stored prices.

## Overview

The CoinMarketCap Currency Price Updater is a Go-based microservice that:
- Fetches cryptocurrency prices from the [CoinMarketCap API](https://coinmarketcap.com/api/) at regular intervals (every 5 minutes).
- Updates or inserts cryptocurrency prices into a database (e.g., MySQL).
- Provides an API endpoint to retrieve all stored cryptocurrency prices.

This service is built using **Go**, **Gin** (for the HTTP server), and a **MySQL** database to store the prices.

## Prerequisites

Before you start, ensure that you have the following installed:

- **Go** (v1.17+): The Go programming language.
  - [Go Installation Guide](https://golang.org/doc/install)
  
- **MySQL** (or a compatible database):
  - [MySQL Installation Guide](https://dev.mysql.com/doc/refman/8.0/en/installing.html)

- **CoinMarketCap API Key**:
  - Sign up for a free or paid account at [CoinMarketCap API](https://coinmarketcap.com/api/).
  - Get your API key after signing in.

## Setup and Installation

 Step-by-step instructions for setting up and running the project, including cloning the repository, installing dependencies, setting up the database, configuring the API key, and running the service.

 **License**: MIT license

git clone https://github.com/anuroopgt/golang.git
cd CoinMarketCapUpdater
