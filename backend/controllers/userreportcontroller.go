package controllers

import (
	"backend/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportRequest struct {
	ReportID uint64 `json:"reportID"`
	Message  string `json:"message"`
}

// reports new user endpoint
func ReportUser(context *gin.Context) {
	var report ReportRequest

	//bind request params
	if err := context.ShouldBindJSON(&report); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	//check userID passed thru jwt token
	userID, doesExist := context.Get("userID")
	if doesExist == false {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	err := database.ReportUser(report.ReportID, userID, report.Message)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"response": "Sucessfully Reported!"})
}
