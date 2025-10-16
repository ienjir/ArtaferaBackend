package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"net/http"
	"strconv"
	"strings"
)

func GetUserByID(c *gin.Context) {
	var json models.GetUserByIDRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	if err := verifyGetUserById(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	user, err := getUserByIDService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"user": user})
	return
}

func GetUserByEmail(c *gin.Context) {
	var json models.GetUserByEmailRequest

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.Email = strings.ToLower(json.Email)

	if err := verifyGetUserByEmail(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	user, err := getUserByEmailService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"user": user})
}

func ListAllUsers(c *gin.Context) {
	var json models.ListUserRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListUserData(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	users, count, err := listUsersService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"count": count, "users": users})
	return
}

func CreateUser(c *gin.Context) {
	var json models.CreateUserRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.Email = strings.ToLower(json.Email)

	if err := verifyCreateUserData(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	user, err := createUserService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	var json models.UpdateUserRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	if err := verifyUpdateUserRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	user, err := updateUserService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"user": user})
	return
}

func DeleteUser(c *gin.Context) {
	var json models.DeleteUserRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	if err := verifyDeleteUserRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	if err := deleteUserService(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, nil, "User successfully deleted")
	return
}
