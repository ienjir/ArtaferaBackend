package picture

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var wrong = false

func getPictureByIDService(data models.GetPictureByIDRequest, context *gin.Context) (*models.Picture, *minio.Object, *models.ServiceError) {
	var picture models.Picture
	var returnMinioFile *minio.Object
	var bucketName string

	if err := database.DB.First(&picture, data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Picture not found",
			}
		} else {
			return nil, nil, &models.ServiceError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error while retrieving picture",
			}
		}
	}

	if picture.IsPublic {
		bucketName = data.PublicBucket
	} else {
		bucketName = data.PrivateBucket
	}

	if minioFile, err := utils.GetFileFromMinio(bucketName, strconv.FormatInt(picture.ID, 10), context); err != nil {
		return nil, nil, err
	} else {
		returnMinioFile = minioFile
	}

	return &picture, returnMinioFile, nil
}

func createPictureService(data models.CreatePictureRequest, context *gin.Context) (*models.Picture, *models.ServiceError) {
	var isPublic bool
	var bucketName string

	if data.Name != nil {
		data.Name = &data.Picture.Filename
	}

	switch {
	case data.IsPublic == nil:
		isPublic = false
		bucketName = data.PrivateBucket

	case *data.IsPublic:
		isPublic = true
		bucketName = data.PublicBucket

	default:
		isPublic = false
		bucketName = data.PrivateBucket
	}

	picture := models.Picture{
		Name:     *data.Name,
		Priority: data.Priority,
		IsPublic: isPublic,
	}

	fmt.Printf("1")

	if db := database.DB.Create(&picture); db.Error != nil {
		return nil,
			&models.ServiceError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to save picture",
			}
	}

	fmt.Printf("2")
	bucketFileName := strconv.Itoa(int(picture.ID))

	_, err := utils.UploadFileToMinio(data.Picture, bucketName, bucketFileName, context)
	if err != nil {
		return nil, err
	}

	fmt.Printf("3")
	return &picture, nil
}
