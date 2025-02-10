package picture

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"net/http"
	"path/filepath"
	"time"
)

func createPictureService(data models.CreatePictureRequest, c *gin.Context) (*models.Picture, *models.ServiceError) {
	file, err := c.FormFile("image")
	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "File upload failed",
		}
	}

	fileExt := filepath.Ext(file.Filename)
	imageName := data.ImageName
	if imageName == "" {
		imageName = fmt.Sprintf("picture_%d%s", time.Now().Unix(), fileExt)
	}

	// Upload image to MinIO
	filePath, err := utils.UploadToMinio(file, imageName)
	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to upload image",
		}
	}

	// Create Picture record in the database
	picture := models.Picture{
		Name:        imageName,
		Priority:    data.Priority,
		PictureLink: filePath,
	}

	if err := database.DB.Create(&picture).Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to save Picture in the database",
		}
	}

	return &picture, nil
}

func CreatePictureFromRequest(c *gin.Context, requestData models.CreatePictureRequest) (*models.Picture, *models.ServiceError) {
	return createPictureService(requestData, c)
}
