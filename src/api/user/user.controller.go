package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"reflect"
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
	}

	if err := DeleteUserService(targetUserID); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
	return
}

func GetUserByID(c *gin.Context) {
	requestUserID := c.GetUint("userID")

	c.JSON(http.StatusOK, gin.H{"number": requestUserID, "test": reflect.TypeOf(requestUserID)})
}
