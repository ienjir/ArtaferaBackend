package language

import (
	"github.com/gin-gonic/gin"
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
