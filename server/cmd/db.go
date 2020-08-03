package cmd

import (
	"log"
	"time"

	"github.com/solidiquis/djangolang/config/defaults"
	"github.com/solidiquis/djangolang/models"

	"github.com/jinzhu/gorm"
	// Init
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func initDB() {
	var err error

	// Establish connection to postgres
	for {
		log.Println("Establishing connection to postgres.")
		db, err = gorm.Open("postgres", defaults.PgDev)
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
}

func migrate() {
	db.AutoMigrate(
		models.User{},
	)
}
