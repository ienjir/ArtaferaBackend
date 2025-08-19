package artTranslation

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/language"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"net/http"
	"strconv"
)

func GetArtTranslationByID(c *gin.Context) {
	var json models.GetArtTranslationByIDRequest

	artTranslationID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	userID := c.GetInt64("userID")
	userRole := c.GetString("userRole")

	json.TargetID = artTranslationID
	json.UserID = userID
	json.UserRole = userRole

	if err := verifyGetArtTranslationByIDRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	artTranslation, err := getArtTranslationByIDService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"artTranslation": artTranslation})
	return
}

func ListArtTranslations(c *gin.Context) {
	var json models.ListArtTranslationRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListArtTranslation(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	artTranslations, count, err := listArtTranslationService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"count": count, "artTranslations": artTranslations})
	return
}

func CreateArtTranslation(c *gin.Context) {
	var json models.CreateArtTranslationRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyCreateArtTranslation(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	id, langErr := language.LanguageCodeToID(json.LanguageCode)
	if langErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrLanguageNotFound)
		return
	}

	createdRole, err := createArtTranslationService(json, id.ID)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"language": createdRole})
	return
}

func UpdateArtTranslation(c *gin.Context) {
	var json models.UpdateArtTranslationRequest
	var artTranslation models.ArtTranslation

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

	if err := verifyUpdateArtTranslation(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	if json.LanguageCode != nil {
		languageID, langErr := language.LanguageCodeToID(*json.LanguageCode)
		if langErr != nil {
			utils.RespondWithError(c, http.StatusBadRequest, utils.ErrLanguageNotFound)
			return
		}
		artTranslation.LanguageID = languageID.ID
	}

	updatedLanguage, err := updateArtTranslation(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"language": updatedLanguage})
	return
}

func DeleteArtTranslation(c *gin.Context) {
	var json models.DeleteArtTranslationRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.TargetID = targetID
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyDeleteArtTranslation(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	if err := deleteArtTranslationService(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"success": "Art translation successfully deleted"})
	return
}
