package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strconv"
)

func GetUserByID(c *gin.Context) {
	var json models.GetUserByIDRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OrderID is wrong"})
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	if err := verifyGetUserById(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	user, err := GetUserByIDService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
	return
}

func GetUserByEmail(c *gin.Context) {
	var json models.GetUserByEmailRequest

	json.UserID = c.GetFloat64("userID")
	json.UserRole = c.GetString("userRole")

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := verifyGetUserByEmail(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	user, err := GetUserByEmailService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func ListAllUsers(c *gin.Context) {
	var json models.ListUserRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListUserData(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	users, count, err := ListUsersService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count, "users": users})
	return
}

func CreateUser(c *gin.Context) {
	var json models.CreateUserRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := verifyCreateUserData(json); err != nil {
		c.JSON(err.StatusCode, err.Message)
		return
	}

	user, err := CreateUserService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	var json models.UpdateUserRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TargetID is wrong"})
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	if err := verifyUpdateUserRequest(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	if err := UpdateUserService(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	return
}

func DeleteUser(c *gin.Context) {
	var json models.DeleteUserRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TargetID is wrong"})
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	if err := verifyDeleteUserRequest(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	if err := DeleteUserService(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
	return
}
