package main

import (
	"go-rest-api/config"
	"go-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	r := gin.Default()
	
	api := r.Group("/api")
	// Auth Routes
	routes.AuthRoutes(api)
	// User Routes
	routes.UserRoutes(api)
	// Todo Routes
	routes.BlogRoutes(api)

	r.Run()
}