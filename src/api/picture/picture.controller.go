package picture

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"io"
	"net/http"
	"strconv"
)

var publicBucket = "pictures"
var privateBucket = "pictures-p"

func GetPictureByID(c *gin.Context) {
	var json models.GetPictureByIDRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PictureID is wrong"})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	if err := verifyGetPictureByIDRequest(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	json.PublicBucket = publicBucket
	json.PrivateBucket = privateBucket

	picture, minioFile, err := getPictureByIDService(json, c)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	buf := new(bytes.Buffer)
	_, err2 := io.Copy(buf, minioFile)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"data": picture,
	})
}

func GetPictureByName(c *gin.Context) {
	var json models.GetPictureByNameRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyGetPictureByNameRequest(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	json.PublicBucket = publicBucket

}

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

	json.PublicBucket = publicBucket
	json.PrivateBucket = privateBucket

	pictureDB, srvErr := createPictureService(json, c)
	if srvErr != nil {
		c.JSON(srvErr.StatusCode, gin.H{"error": srvErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"picture": pictureDB})
}
