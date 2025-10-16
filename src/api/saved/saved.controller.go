package saved

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"net/http"
	"strconv"
)

func GetSavedByID(c *gin.Context) {
	var json models.GetSavedByIDRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	if err := verifyGetSavedById(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	user, err := getSavedByIDService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"user": user})
	return
}

func GetSavedForUser(c *gin.Context) {
	var json models.GetSavedForUserRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.TargetUserID = targetID
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyGetSavedForUserRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	saved, user, count, err := getSavedForUserService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"count": count, "user": user, "saved": saved})
	return
}

func ListOrder(c *gin.Context) {
	var json models.ListSavedRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListSavedRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	saved, count, err := listSavedService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"count": count, "saved": saved})
	return
}

func CreateSaved(c *gin.Context) {
	var json models.CreateSavedRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyCreateSaved(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	saved, err := createSavedService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"saved": saved})
}

func UpdateSaved(c *gin.Context) {
	var json models.UpdateSavedRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.TargetID = targetID
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyUpdateSavedRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	saved, err := updateSavedService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"saved": saved})
	return
}

func DeleteSaved(c *gin.Context) {
	var json models.DeleteSavedRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.TargetID = targetID
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyDeleteSavedRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	if err := deleteSavedService(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Saved successfully deleted"})
}
