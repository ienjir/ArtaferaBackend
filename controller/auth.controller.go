package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/ienjir/ArtaferaBackend/models"
	"github.com/ienjir/ArtaferaBackend/services"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	var u models.User

	// Bind JSON body to the user model
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Printf("The user request value: %v\n", u)

	if u.Username == "Chek" && u.Password == "123456" {
		tokenString, err := services.CreateToken(u.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

// ProtectedHandler handles requests to a protected endpoint
func ProtectedHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return
	}

	// Remove "Bearer " prefix
	tokenString = tokenString[len("Bearer "):]

	if err := services.VerifyToken(tokenString); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected area"})
}
