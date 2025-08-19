package art

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strconv"
)

func GetArtByID(c *gin.Context) {
	var json models.GetArtByIDRequest

	artID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "ArtID is wrong"})
		return
	}

	userID := c.GetInt64("userID")
	userRole := c.GetString("userRole")

	json.TargetID = artID
	json.UserID = userID
	json.UserRole = userRole

	if err := verifyGetArtByID(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	art, err := getArtByIDService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"art": art})
	return
}

func ListArts(c *gin.Context) {
	var json models.ListArtRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListArt(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	arts, count, err := listArtService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count, "arts": arts})
	return
}

func CreateArt(c *gin.Context) {
	var json models.CreateArtRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyCreateArt(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	art, err := createArtService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"art": art})
	return
}

func UpdateArt(c *gin.Context) {
	var json models.UpdateArtRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	artID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ArtID is wrong"})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = artID

	if err := verifyUpdateArt(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	art, err := updateArtService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"art": art})
	return
}

func DeleteArt(c *gin.Context) {
	var json models.DeleteArtRequest

	artID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ArtID is wrong"})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = artID

	if err := verifyDeleteArt(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	if err := deleteArtService(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Art successfully deleted"})
	return
}
