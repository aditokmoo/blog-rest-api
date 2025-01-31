package controller

import (
	"go-rest-api/config"
	"go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func UpdateUser(c *gin.Context) {}

func DeleteUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var blog models.Blog
	var user models.User

	if err := config.DB.Unscoped().Where("user_id = ?", userID).Delete(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete user blogs"})
		return
	}

	if err := config.DB.Unscoped().Where("id = ?", userID).Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Account has been deleted"})
}