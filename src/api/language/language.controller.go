package language

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"net/http"
	"strconv"
)

func GetLanguageByID(c *gin.Context) {
	var json models.GetLanguageByIDRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	if err := verifyGetLanguageByIDRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	language, err := getLanguageByIDService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"language": language})
	return
}

func ListLanguages(c *gin.Context) {
	var json models.ListLanguageRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListLanguagesRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	languages, count, err := listLanguageService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"count": count, "languages": languages})
	return
}

func CreateLanguage(c *gin.Context) {
	var json models.CreateLanguageRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyCreateLanguageRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	createdRole, err := createLanguageService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"language": createdRole})
	return
}

func UpdateLanguage(c *gin.Context) {
	var json models.UpdateLanguageRequest

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

	if err := verifyUpdateLanguageRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	updatedLanguage, err := updateLanguageService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"language": updatedLanguage})
	return
}

func DeleteLanguage(c *gin.Context) {
	var json models.DeleteLanguageRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	if err := verifyDeleteLanguageRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	if err := deleteLanguageService(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Language successfully deleted"})
	return
}
