package main

import (
	"go-rest-api/config"
	"go-rest-api/models"
	"go-rest-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	r := gin.Default()

	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Migration failed: ", err)
	}
	
	api := r.Group("/api")
	// Auth Routes
	routes.AuthRoutes(api)
	// User Routes
	routes.UserRoutes(api)
	// Todo Routes
	routes.BlogRoutes(api)

	r.Run()
}