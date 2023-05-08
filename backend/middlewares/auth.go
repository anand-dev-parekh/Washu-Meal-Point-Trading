package middlewares

import (
	"backend/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// confirms user auth
func UserAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		//check whether auth token passed in header
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		claims, err := auth.ValidateToken(tokenString) //validate auth token
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		//1 corresponds to unverified email, if so do not allow
		if claims.AuthLevel == 1 {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "unverified email"})
			context.Abort()
			return
		}

		//4 corresponds to banned user, if so do not allow
		if claims.AuthLevel == 3 {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "You have been banned"})
			context.Abort()
			return
		}
		context.Set("userID", claims.UserID)
		context.Next()
	}
}

// confirms admin auth
func AdminAuth() gin.HandlerFunc {
	return func(context *gin.Context) {

		//checks auth token passed in header
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		claims, err := auth.ValidateToken(tokenString) //validate auth token
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		//1 corresponds to unverified email, if so do not allow
		if claims.AuthLevel == 1 {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "unverified email"})
			context.Abort()
			return
		}
		//2 corresponds to regular user, if so do not allow
		if claims.AuthLevel == 2 {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "You are not a admin"})
			context.Abort()
			return
		}

		//4 corresponds to banned user, if so do not allow
		if claims.AuthLevel == 3 {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "You have been banned"})
			context.Abort()
			return
		}

		context.Set("userID", claims.UserID)
		context.Next()
	}
}
