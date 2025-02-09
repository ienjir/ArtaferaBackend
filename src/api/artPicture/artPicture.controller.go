package artPicture

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strconv"
)

const (
	UploadDir  = "./uploads"
	BucketName = "artpictures"
)

func CreateArtPicture(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	artID := c.PostForm("artID")

	parsedArtID, err := strconv.ParseInt(artID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid artID format"})
		return
	}

	json := models.CreateArtPictureRequest{
		ArtID:     parsedArtID,
		ImageName: c.PostForm("imageName"),
		UserID:    c.GetInt64("userID"),
		UserRole:  c.GetString("userRole"),
	}

	if err := verifyCreateArtPicture(json, c); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	artPicture, err2 := createArtPictureService(json, c)
	if err2 != nil {
		c.JSON(err2.StatusCode, gin.H{"error": err2.Message})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"art_picture": artPicture})
}
