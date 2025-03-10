package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func CreateJWT(userID uint) (string, error) {
	payload := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateConfirmToken() string {
	b := make([]byte, 20)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func SendConfirmationMail(email, token string) error {
	from := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASS")
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	subject := "Please Confirm Your Registration"
	body := fmt.Sprintf(
		"%s", `<h1>Confirm Your Registration</h1>
		<p>Click the link below to confirm your registration:</p>
		<a href=`+os.Getenv("HOST")+`/api/auth/verify/`+token+`>Confirm your registration</a>`,
	)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func Protect() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized"})
			c.Abort();
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		
		token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))
		c.Set("userID", userID)
		c.Next()
	}
}