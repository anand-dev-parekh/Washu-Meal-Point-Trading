package controllers

import (
	"backend/auth"
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// generate token endpoint/login endpoint
func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User

	//bind request params
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	//scan user obj w/ user from database
	record := database.GetUser(&user, request.Email)
	if record != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error()})
		context.Abort()
		return
	}

	//checks request password matches database password
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	//generate jwt token w/ user email
	tokenString, err := auth.GenerateJWT(user.Email, user.AuthLevel, user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": tokenString, "userID": user.Id})
}
