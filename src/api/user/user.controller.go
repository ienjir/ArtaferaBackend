package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var json models.CreateUserRequest

	// Validate the input
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err2 := VerifyCreateUserData(json)
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
		return
	}

	if err := VerifyListUserData(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	users, count, err := ListUsers(json.Offset)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count, "users": users})
	return
}

func DeleteUser(c *gin.Context) {
	requestUserID := c.GetFloat64("userID")
	requestUserRole := c.GetString("userRole")
	targetUserID := c.Param("id")
	var targetUserIDFloat float64
	var err error

	targetUserIDFloat, err = strconv.ParseFloat(targetUserID, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while parsing ID's"})
	}

	if requestUserRole != "admin" && requestUserID != targetUserIDFloat {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own account", "requestId": requestUserID, "targetID": targetUserID})
		return
	}

	if err := database.DB.Where("id = ?", targetUserID).Delete(&models.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting user"})
		return
	}

	fmt.Println(requestUserID)
	fmt.Println(targetUserID)

	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
	return
}
