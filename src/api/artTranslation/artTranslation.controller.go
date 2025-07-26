package artTranslation

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/api/language"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strconv"
)

func GetArtTranslationByID(c *gin.Context) {
	var json models.GetArtTranslationByIDRequest

	artTranslationID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "ArtTranslationID is wrong"})
	}

	userID := c.GetInt64("userID")
	userRole := c.GetString("userRole")

	json.TargetID = artTranslationID
	json.UserID = userID
	json.UserRole = userRole

	if err := verifyGetArtTranslationByIDRequest(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	artTranslation, err := getArtTranslationByIDService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"artTranslation": artTranslation})
	return
}

func ListArtTranslations(c *gin.Context) {
	var json models.ListArtTranslationRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListArtTranslation(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	artTranslations, count, err := listArtTranslationService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count, "artTranslations": artTranslations})
	return
}

func CreateArtTranslation(c *gin.Context) {
	var json models.CreateArtTranslationRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyCreateArtTranslation(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	id, langErr := language.LanguageCodeToID(json.LanguageCode)
	if langErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": langErr.Error()})
		return
	}

	createdRole, err := createArtTranslationService(json, id.ID)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"language": createdRole})
	return
}

func UpdateArtTranslation(c *gin.Context) {
	var json models.UpdateArtTranslationRequest
	var artTranslation models.ArtTranslation

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TranslationID is wrong"})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	if err := verifyUpdateArtTranslation(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	if json.LanguageCode != nil {
		languageID, langErr := language.LanguageCodeToID(*json.LanguageCode)
		if langErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": langErr.Error()})
			return
		}
		artTranslation.LanguageID = languageID.ID
	}

	updatedLanguage, err := updateArtTranslation(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"language": updatedLanguage})
	return
}
