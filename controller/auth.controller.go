package controller

import (
	"go-rest-api/config"
	"go-rest-api/models"
	"go-rest-api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Register(c *gin.Context) {
	var input struct {
		Name      string `json:"name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
	}

	// Validate data
	if err := c.ShouldBindJSON(&input); err != nil {
		var errorMessages []string

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

	// Check if user  exist
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "User already exists"})
		return	
	}

	// hash password and generate confirm token
	hashedPassword, _ := utils.HashPassword(input.Password)
	confirmToken := utils.GenerateConfirmToken()

	user := models.User{
		Name: input.Name,
		Email: input.Email,
		Password: hashedPassword,
		Confirmed: false,
		ConfirmToken: confirmToken,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Belaj neki"})
		return
	}

	if err := utils.SendConfirmationMail(user.Email, confirmToken); err != nil {
		log.Println("Failed to send confirmation email: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to send confirmation email"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{ "status": "success", "User": user })
}

func Login(c *gin.Context) {
	log.Println("Login Controller")
}

func VerifyAccount(c *gin.Context) {
	confirmToken := c.Param("confirmToken")

	var user models.User
	if err := config.DB.Where("confirm_token = ?", confirmToken).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid token"})
		return
	}

	user.Confirmed = true
	user.ConfirmToken = ""
	
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Account verified"})
}