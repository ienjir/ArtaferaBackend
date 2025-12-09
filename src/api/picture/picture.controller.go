package picture

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
)

var PublicBucket = "pictures"
var PrivateBucket = "pictures-p"

func GetPictureByID(c *gin.Context) {
	var json models.GetPictureByIDRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	if err := verifyGetPictureByIDRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	json.PublicBucket = PublicBucket
	json.PrivateBucket = PrivateBucket

	picture, minioFile, err := getPictureByIDService(json, c)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	buf := new(bytes.Buffer)
	_, err2 := io.Copy(buf, minioFile)
	if err2 != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, utils.ErrFileUpload)
		return
	}

	// Encode the picture as base64
	base64Picture := base64.StdEncoding.EncodeToString(buf.Bytes())

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"metadata": picture,
		"picture":  base64Picture,
	})
}

func GetPictureByName(c *gin.Context) {
	var json models.GetPictureByNameRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyGetPictureByNameRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	json.PublicBucket = PublicBucket
	json.PrivateBucket = PrivateBucket
	json.Name = strings.TrimSuffix(json.Name, filepath.Ext(json.Name))

	picture, minioFile, err := getPictureByNameService(json, c)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	buf := new(bytes.Buffer)
	_, err2 := io.Copy(buf, minioFile)
	if err2 != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, utils.ErrFileUpload)
		return
	}

	// Encode the picture as base64
	base64Picture := base64.StdEncoding.EncodeToString(buf.Bytes())

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"metadata": picture,
		"picture":  base64Picture,
	})
}

func ListPicture(c *gin.Context) {
	var json models.ListPictureRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListPicture(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	json.PublicBucket = PublicBucket
	json.PrivateBucket = PrivateBucket

	pictures, minioFiles, count, err := listPictureService(json, c)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	var base64Pictures []string

	for _, minioFile := range *minioFiles {
		buf := new(bytes.Buffer)
		_, err2 := io.Copy(buf, &minioFile)
		if err2 != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, utils.ErrFileUpload)
			return
		}

		// Encode the pictures as base64
		base64Picture := base64.StdEncoding.EncodeToString(buf.Bytes())

		base64Pictures = append(base64Pictures, base64Picture)
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"metadata": pictures,
		"count":    count,
		"picture":  base64Pictures,
	})
	return
}

func CreatePicture(c *gin.Context) {
	var json models.CreatePictureRequest

	picture, err := c.FormFile("picture")
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrPictureRequired)
		return
	}

	isPublicStr := c.PostForm("isPublic")
	if isPublicStr != "" {
		isPublic := isPublicStr == "true"
		json.IsPublic = &isPublic
	}

	json.Picture = *picture
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if isPublic := c.PostForm("isPublic"); isPublic != "" {
		if isPublicBool, err := strconv.ParseBool(isPublic); err == nil {
			json.IsPublic = &isPublicBool
		} else {
			utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidPublicFormat)
			return
		}
	}

	if err := verifyCreatePicture(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	json.PublicBucket = PublicBucket
	json.PrivateBucket = PrivateBucket

	pictureDB, srvErr := createPictureService(json, c)
	if srvErr != nil {
		utils.RespondWithServiceError(c, srvErr)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"picture": pictureDB})
}

func UpdatePicture(c *gin.Context) {
	var json models.UpdatePictureRequest

	if err := c.ShouldBind(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	if json.Name == nil && json.IsPublic == nil && json.Priority == nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrNoContentFound)
		return
	}

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.TargetID = targetID
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyUpdatePicture(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	picture, err := updatePictureService(json, c)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"picture": picture})
}

func DeletePicture(c *gin.Context) {
	var json models.DeletePictureRequest

	if err := c.ShouldBind(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.TargetID = targetID
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyDeletePicture(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	if err := deletePictureService(json, c); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"success": "Picture successfully deleted"})
	return
}
