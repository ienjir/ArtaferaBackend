package user

import (
	"fmt"
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

	users, count, err := ListUsersService(json.Offset)
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

	if err := VerifyDeleteUserRequest(requestUserID, requestUserRole, targetUserID); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	if err := DeleteUserService(targetUserID); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
	return
}

func GetUserByID(c *gin.Context) {
	var user *models.User
	var err *models.ServiceError
	requestUserID := c.GetInt64("userID")
	requestUserRole := c.GetString("userRole")
	targetUserID := c.Param("id")

	if err := VerifyGetUserById(requestUserID, requestUserRole, targetUserID); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	if user, err = GetUserByIDService(targetUserID); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
	return
}

func GetUserByEmail(c *gin.Context) {
	var user *models.User
	var err *models.ServiceError
	var Data models.GetUserByEmail

	Data.RequestID = c.GetFloat64("userID")
	Data.RequestRole = c.GetString("userRole")

	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if err := VerifyGetUserByEmail(Data); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	if user, err = GetUserByEmailService(Data); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	requestUserID := c.GetInt64("userID")
	requestUserRole := c.GetString("userRole")
	targetUserID := c.Param("id")

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateUpdateUserRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("RequestID: %d, RequestRole: %s, TargetId: %s \n", requestUserID, requestUserRole, targetUserID)

	// Attempt to update user
	if err := UpdateUserService(requestUserID, requestUserRole, targetUserID, req); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
