package routes

import (
	"go-rest-api/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	authRoutes := r.Group("/auth")

	authRoutes.POST("/register", controller.Register)
	authRoutes.POST("/login", controller.Login)
}