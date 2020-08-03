package cmd

import (
	"log"

	// Init
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// Execute ...starts app
func Execute() {
	// load .env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	initDB()
	defer db.Close()

	migrate()

	startServer()
}
