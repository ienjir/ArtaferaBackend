package picture

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

var bucketName = "pictures"

func CreatePicture(c *gin.Context) {
	var json models.CreatePictureRequest

	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	picture, err := c.FormFile("picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	json.Picture = *picture

	if err := verifyCreatePicture(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	json.BucketName = bucketName

	pictureDB, srvErr := createPictureService(json, c)
	if srvErr != nil {
		c.JSON(srvErr.StatusCode, gin.H{"error": srvErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"picture": pictureDB})
}
