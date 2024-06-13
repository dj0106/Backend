package handler

import (
	"InvitKaro/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	// This is a mockup; you should validate user credentials here
	userID := c.PostForm("user_id")
	password := c.PostForm("password")

	// Validate the user credentials (this is a placeholder logic)
	if userID == "valid_user" && password == "valid_password" {
		token, err := Middleware.GenerateJWT(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

func RegisterHandler(c *gin.Context) {
	// Handle registration logic here
	c.JSON(http.StatusOK, gin.H{"message": "Registration endpoint"})
}

func GetLocationHandler(c *gin.Context) {
	// Handle getting location logic here
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get location", "id": id})
}
