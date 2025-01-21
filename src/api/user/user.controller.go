package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var json models.CreateUserRequest

	// Validate the input
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err2 := verifyCreateUserData(json)
	if err2 != nil {
		c.JSON(err2.StatusCode, err2.Message)
		return
	}

	// Call the service to handle user creation
	user, err3 := CreateUserService(json)
	if err3 != nil {
		c.JSON(err3.StatusCode, gin.H{"error": err3.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func ListAllUsers(c *gin.Context) {
	var json models.ListUserRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := verifyListUserData(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
	}

	if err := ListUsers(json.Offset); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
	}
}
