package controllers

import (
	"backend/auth"
	"backend/database"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

type EmailConfirm struct {
	Email string `json:"email"`
	ID    uint64 `json:"id"`
}

// sends confirmation email endpoint
func SendConfirmationEmail(context *gin.Context) {
	//Get jwt token
	tokenString := context.GetHeader("Authorization")
	if tokenString == "" {
		context.JSON(401, gin.H{"error": "request does not contain an access token"})
		context.Abort()
		return
	}

	//confirm token
	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		context.JSON(401, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	from := "anand.d.parekh@gmail.com"
	password := os.Getenv("EMAILPASS")

	to := []string{
		claims.Email,
	}

	//create type to be encoded
	info := EmailConfirm{
		Email: claims.Email,
		ID:    claims.UserID,
	}
	//turn to string
	jsonBytes, err := json.Marshal(info)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	//encode to base64
	encodedData := base64.StdEncoding.EncodeToString(jsonBytes)

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	//template email
	subject := "Subject: Washu Meal Points Trader\r\n"
	body := "Thanks for using Washu Meal Points! Please verify your account using this link: http://localhost:8080/verify-email/" + encodedData + "\r\n"
	message := []byte(subject + "\r\n" + body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)
	// Sending email.
	sendErr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if sendErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": sendErr.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"response": "Sent Confirmation Email!"})
}

// verifies email endpoint
func VerifyEmail(context *gin.Context) {
	encodedInfo := context.Param("encodedInfo")

	decodedBytes, err := base64.StdEncoding.DecodeString(encodedInfo)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var info EmailConfirm
	err = json.Unmarshal(decodedBytes, &info)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	err = database.VerifyUser(info.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"response": "You verified your email!"})
}
