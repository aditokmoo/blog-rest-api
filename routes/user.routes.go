package routes

import (
	"go-rest-api/controller"
	"go-rest-api/utils"

	"github.com/gin-gonic/gin"
)

func UserRoutes(c *gin.RouterGroup) {
	userRoute := c.Group("/users")

	userRoute.GET("/:id", controller.GetUser)
	userRoute.PUT("/me", utils.Protect(), controller.UpdateUser)
	userRoute.DELETE("/me", utils.Protect(), controller.DeleteUser)
}