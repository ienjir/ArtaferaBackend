package picture

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strconv"
)

var BucketName = "pictures"

func CreatePicture(c *gin.Context) {
	var intPriority int

	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	priority := c.PostForm("priority")

	fmt.Printf("Priority: %s \n", priority)

	if priority != "" {
		parsedPriority, err := strconv.Atoi(priority)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid artID format"})
			return
		}

		intPriority = parsedPriority
	}

	json := models.CreatePictureRequest{
		ImageName: c.PostForm("imageName"),
		UserID:    c.GetInt64("userID"),
		UserRole:  c.GetString("userRole"),
		Priority:  &intPriority,
	}

	if err := verifyCreatePicture(json, c); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	artPicture, err2 := createPictureService(json, c)
	if err2 != nil {
		c.JSON(err2.StatusCode, gin.H{"error": err2.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"art_picture": artPicture})
}
