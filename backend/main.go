package main

import (
	"backend/controllers"
	"backend/database"
	"backend/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	// Add the CORS middleware to the router
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	//add public endpoints
	router.POST("/create-user", controllers.RegisterUser)
	router.POST("/generate-token", controllers.GenerateToken)

	router.PUT("/send-confirmation-email", controllers.SendConfirmationEmail)
	router.GET("/verify-email/:encodedInfo", controllers.VerifyEmail)

	//endpoints for users logged in
	secured_user := router.Group("/secure").Use(middlewares.UserAuth())
	{
		secured_user.GET("/get-offers", controllers.GetOffers)

		secured_user.POST("/create-offer", controllers.CreateOffer)
		secured_user.DELETE("/delete-offer", controllers.DeleteOffer)
		secured_user.PUT("/update-offer", controllers.UpdateOffer)

		secured_user.POST("/report-user", controllers.ReportUser)
	}

	//endpoints for admins
	admin_user := router.Group("admin").Use(middlewares.AdminAuth())
	{
		admin_user.GET("/get-reports", controllers.GetReports)
		admin_user.PUT("/ban-user", controllers.BanUser)
	}

	router.Run("localhost:8080")
}
