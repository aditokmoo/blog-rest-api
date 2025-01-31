package main

import (
	"go-rest-api/config"
	"go-rest-api/models"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	config.DB.AutoMigrate(&models.User{})
}