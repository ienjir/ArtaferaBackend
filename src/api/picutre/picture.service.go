package picture

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"net/http"
	"strconv"
)

func createPictureService(data models.CreatePictureRequest, context *gin.Context) (*models.Picture, *models.ServiceError) {
	if data.Name != "" {
		data.Name = data.Picture.Filename
	}

	picture := models.Picture{
		Name:     data.Name,
		Priority: data.Priority,
		IsPublic: data.IsPublic,
	}

	if db := database.DB.Create(&picture); db.Error != nil {
		return nil,
			&models.ServiceError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to save picture",
			}
	}

	bucketFileName := strconv.Itoa(int(picture.ID))

	_, err := utils.UploadFileToMinio(data.Picture, data.BucketName, bucketFileName, context)
	if err != nil {
		return nil, err
	}

	return &picture, nil
}
