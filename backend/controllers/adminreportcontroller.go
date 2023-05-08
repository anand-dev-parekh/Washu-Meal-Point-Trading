package controllers

import (
	"backend/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BanRequest struct {
	BanUserID uint64 `json:"banUserID"`
}

// gets reports endpoints
func GetReports(context *gin.Context) {
	reports, err := database.GetReports()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"reports": reports})
}

// bans users endpoint
func BanUser(context *gin.Context) {
	var userBanned BanRequest

	//bind request params
	if err := context.ShouldBindJSON(&userBanned); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	err := database.BanUser(userBanned.BanUserID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"body": "Sucessfully banned user"})
}
