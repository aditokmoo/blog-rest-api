package routes

import (
	"go-rest-api/controller"
	"go-rest-api/utils"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(r *gin.RouterGroup) {
	blogRoutes := r.Group("/blogs")

	blogRoutes.POST("/", utils.Protect(), controller.CreateBlog)
	blogRoutes.GET("/", controller.GetBlogs)
	blogRoutes.GET("/:id", controller.GetBlog)
	blogRoutes.PATCH("/:id", controller.UpdateBlog)
	blogRoutes.DELETE("/:id", controller.DeleteBlog)
}