package artPicture

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

func createArtPictureService(data models.CreateArtPictureRequest, c *gin.Context) (*models.ArtPicture, *models.ServiceError) {
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
		imageName = fmt.Sprintf("%d_%d%s", data.ArtID, time.Now().Unix(), fileExt)
	}

	filePath, err := utils.UploadToMinio(file, imageName)
	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to upload image",
		}
	}

	artPicture := models.ArtPicture{
		ArtID:     data.ArtID,
		Name:      imageName,
		PictureID: 1,
	}

	fmt.Printf(filePath)

	if err := database.DB.Create(&artPicture).Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to save ArtPicture in the database",
		}
	}

	return &artPicture, nil
}
