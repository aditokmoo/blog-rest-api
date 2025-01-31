package controller

import (
	"go-rest-api/config"
	"go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateBlog(c *gin.Context) {
	var body struct {
		Title string `json:"title" binding:"required"`
		Content  string `json:"content" binding:"required"`
	}

	var errorMessages []string
	if err := c.ShouldBindJSON(&body); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, vErr := range validationErrors {
				errorMessages = append(errorMessages, vErr.Field()+" is "+vErr.Tag())
			}
		} else {
			errorMessages = append(errorMessages, "Invalid request body")
		}

		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": errorMessages})
		return
	}

	userID := c.MustGet("userID").(uint)
	blog := models.Blog{Title: body.Title, Content: body.Content, UserID: userID}
	if err := config.DB.Create(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Blog created", "data": blog})
}

func GetBlogs(c *gin.Context) {}

func GetBlog(c *gin.Context) {}

func UpdateBlog(c *gin.Context) {}

func DeleteBlog(c *gin.Context) {}
