package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// registers a new user
func RegisterUser(context *gin.Context) {
	var user models.User

	//if req body is not user obj format return
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	//if hashpassword err return error
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	//create new user obj in sql database
	record := database.CreateUser(&user)
	if record != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"body": "Sucessfuly created new user!"})
}
