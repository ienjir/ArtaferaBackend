package ArtTranslation

import (
	"github.com/gin-gonic/gin"
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
