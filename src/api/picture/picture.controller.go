package picture

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strconv"
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
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	json.Picture = *picture

	if priority := c.PostForm("priority"); priority != "" {
		if priorityInt, err := strconv.ParseInt(priority, 10, 8); err == nil {
			json.Priority = &priorityInt
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid priority format"})
			return
		}
	}

	if name := c.PostForm("name"); name != "" {
		json.Name = &name
	}

	if isPublic := c.PostForm("isPublic"); isPublic != "" {
		if isPublicBool, err := strconv.ParseBool(isPublic); err == nil {
			json.IsPublic = &isPublicBool
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid isPublic format"})
			return
		}
	}

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
