package artPicture

import (
	"fmt"
	"github.com/gin-gonic/gin"
	picture "github.com/ienjir/ArtaferaBackend/src/api/picutre"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

const (
	UploadDir  = "./uploads"
	BucketName = "artpictures"
)

func CreateArtPicture(c *gin.Context) {
	var json models.CreateArtPictureRequest
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	fmt.Printf("ArtID: %d \n", json.ArtID)
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	pictureRequest := models.CreatePictureRequest{
		UserID:    json.UserID,
		UserRole:  json.UserRole,
		ImageName: json.ImageName,
	}

	if err := verifyCreateArtPicture(json, c); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	createdPicture, err := picture.CreatePictureFromRequest(c, pictureRequest)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	artPicture := models.ArtPicture{
		ArtID:     json.ArtID,
		PictureID: createdPicture.ID,
	}

	if err := database.DB.Create(&artPicture).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create art picture association"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"picture":     createdPicture,
		"art_picture": artPicture,
	})
}
