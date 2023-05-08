package controllers

import (
	"backend/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	MealPointsOffer uint16 `json:"mealPointsOffer"`
}

type DeleteRequest struct {
	OfferID uint64 `json:"offerID"`
}

type UpdateRequest struct {
	OfferID         uint64 `json:"offerID"`
	MealPointsOffer uint16 `json:"mealPointsOffer"`
}

// gets offers endpoint
func GetOffers(context *gin.Context) {
	offers, err := database.GetOffers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"offers": offers})
}

// creates new offer endpoint
func CreateOffer(context *gin.Context) {
	var offerAmount CreateRequest

	//bind request params
	if err := context.ShouldBindJSON(&offerAmount); err != nil {
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

	err := database.CreateNewOffer(userID, offerAmount.MealPointsOffer)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"response": "Created new offer!"})
}

// deletes offer endpoint
func DeleteOffer(context *gin.Context) {
	var offerID DeleteRequest

	//bind request params
	if err := context.ShouldBindJSON(&offerID); err != nil {
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

	err := database.DeleteOffer(userID, offerID.OfferID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"body": "You sucessfully deleted the offer!"})
}

// updates offer endpoint
func UpdateOffer(context *gin.Context) {
	var updatedOffer UpdateRequest

	//bind request params
	if err := context.ShouldBindJSON(&updatedOffer); err != nil {
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

	err := database.UpdateOffer(userID, updatedOffer.OfferID, updatedOffer.MealPointsOffer)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"body": "You sucessfully updated the offer!"})
}
