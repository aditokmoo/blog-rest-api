package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB * gorm.DB

func ConnectDB() {
    connStr := os.Getenv("DB_URL")

    db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: " + err.Error())
    }

	log.Println("Database connected")

	DB = db
}
