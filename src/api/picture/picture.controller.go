package picture

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

var PublicBucket = "pictures"
var PrivateBucket = "pictures-p"

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

	json.PublicBucket = PublicBucket
	json.PrivateBucket = PrivateBucket

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

	// Encode the picture as base64
	base64Picture := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.JSON(http.StatusOK, gin.H{
		"metadata": picture,
		"picture":  base64Picture,
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

	json.PublicBucket = PublicBucket
	json.PrivateBucket = PrivateBucket
	json.Name = strings.TrimSuffix(json.Name, filepath.Ext(json.Name))

	picture, minioFile, err := getPictureByNameService(json, c)
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

	// Encode the picture as base64
	base64Picture := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.JSON(http.StatusOK, gin.H{
		"metadata": picture,
		"picture":  base64Picture,
	})
}

func ListPicture(c *gin.Context) {
	var json models.ListPictureRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListPicture(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	json.PublicBucket = PublicBucket
	json.PrivateBucket = PrivateBucket

	pictures, minioFiles, count, err := listPictureService(json, c)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	var base64Pictures []string

	for _, minioFile := range *minioFiles {
		buf := new(bytes.Buffer)
		_, err2 := io.Copy(buf, &minioFile)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
			return
		}

		// Encode the pictures as base64
		base64Picture := base64.StdEncoding.EncodeToString(buf.Bytes())

		base64Pictures = append(base64Pictures, base64Picture)
	}

	c.JSON(http.StatusOK, gin.H{
		"metadata": pictures,
		"count":    count,
		"picture":  base64Pictures,
	})
	return
}

func CreatePicture(c *gin.Context) {
	var json models.CreatePictureRequest

	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	picture, err := c.FormFile("picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Picture is required"})
		return
	}

	json.Picture = *picture
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if priority := c.PostForm("priority"); priority != "" {
		if priorityInt, err := strconv.ParseInt(priority, 10, 8); err == nil {
			json.Priority = &priorityInt
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid priority format"})
			return
		}
	}

	if name := c.PostForm("name"); name != "" {
		trimmedName := strings.TrimSuffix(*json.Name, filepath.Ext(json.Picture.Filename))
		json.Name = &trimmedName
	} else {
		filenameWithoutExt := strings.TrimSuffix(json.Picture.Filename, filepath.Ext(json.Picture.Filename))
		json.Name = &filenameWithoutExt
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

	json.PublicBucket = PublicBucket
	json.PrivateBucket = PrivateBucket

	pictureDB, srvErr := createPictureService(json, c)
	if srvErr != nil {
		c.JSON(srvErr.StatusCode, gin.H{"error": srvErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"picture": pictureDB})
}

func UpdatePicture(c *gin.Context) {
	var json models.UpdatePictureRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not convert ID"})
		return
	}

	json.TargetID = targetID
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyUpdatePicture(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	picture, err := updatePictureService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"picture": picture})
}
