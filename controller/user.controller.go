package controller

import (
	"go-rest-api/config"
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"
	"strings"

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

func UpdateUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var body struct {
		Name *string `json:"name" bindng:"omitempty",min=2,max=100`
		Email * string `json:"email" bidng:"omitempty",email`
		Password *string `json:"password" binding:"omitempty,min=6"`
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Body is required"})
		return
	}

	updates := make(map[string]interface{})
	if body.Name != nil {
		updates["Name"] = strings.TrimSpace(*body.Name)
	}
	if body.Email != nil {
		// Email cannot be changed at this time
		c.JSON(http.StatusNotAcceptable, gin.H{"status": "error", "messsage": "Email can't be updated"})
		return
	}
	if body.Password != nil {
		hashedPassword, err := utils.HashPassword(*body.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Failed to hash password"})
			return
		}
		updates["Password"] = hashedPassword
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No fields updated"})
		return
	}

	result := config.DB.Model(&user).Updates(updates)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Failed to update user"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No changes detected"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User updated successfully", "user": user})
}

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