package language

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func GetLanguageByID(c *gin.Context) {
	requestUserID := c.GetInt64("userID")
	requestUserRole := c.GetString("userRole")
	targetLanguageID := c.Param("id")

	if err := verifyGetLanguageByID(requestUserID, requestUserRole, targetLanguageID); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	language, err := getLanguageByIDService(targetLanguageID)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"language": language})
	return
}

func ListLanguages(c *gin.Context) {
	var Data models.ListLanguageRequest

	if err := c.ShouldBindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	languages, count, err := listLanguageService(Data.Offset)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count, "languages": languages})
	return
}

func CreateLanguage(c *gin.Context) {
	var language models.CreateLanguageRequest

	if err := c.ShouldBindJSON(&language); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdRole, err := createLanguageService(language)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"language": createdRole})
	return
}

func UpdateLanguage(c *gin.Context) {
	var language models.UpdateLanguageRequest

	language.LanguageID = c.Param("id")

	if err := c.ShouldBindJSON(&language); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := verifyUpdateLanguage(language); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.StatusCode})
		return
	}

	updatedLanguage, err := updateRoleService(language)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"language": updatedLanguage})
	return
}
