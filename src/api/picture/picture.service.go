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
	"path/filepath"
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

	fileName := picture.Name + "__" + strconv.Itoa(int(picture.ID)) + picture.Type

	if minioFile, err := utils.GetFileFromMinio(bucketName, fileName, context); err != nil {
		return nil, nil, err
	} else {
		returnMinioFile = minioFile
	}

	return &picture, returnMinioFile, nil
}

func getPictureByNameService(data models.GetPictureByNameRequest, context *gin.Context) (*models.Picture, *minio.Object, *models.ServiceError) {
	var picture models.Picture
	var returnMinioFile *minio.Object
	var bucketName string

	if err := database.DB.Where("Name = ?", data.Name).First(&picture).Error; err != nil {
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

	fileName := picture.Name + "__" + strconv.Itoa(int(picture.ID)) + picture.Type

	if minioFile, err := utils.GetFileFromMinio(bucketName, fileName, context); err != nil {
		return nil, nil, err
	} else {
		returnMinioFile = minioFile
	}

	return &picture, returnMinioFile, nil
}

func listPictureService(data models.ListPictureRequest, context *gin.Context) (*[]models.Picture, *[]minio.Object, *int64, *models.ServiceError) {
	var pictures []models.Picture
	var returnMinioFiles []minio.Object
	var count int64

	if err := database.DB.Limit(5).Offset(int(data.Offset * 5)).Find(&pictures).Error; err != nil {
		return nil, nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving orders from pictures",
		}
	}

	if err := database.DB.Model(&models.Picture{}).Count(&count).Error; err != nil {
		return nil, nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while counting pictures in database",
		}
	}

	for _, picture := range pictures {
		var bucketName string
		fileName := picture.Name + "__" + strconv.Itoa(int(picture.ID)) + picture.Type

		if picture.IsPublic {
			bucketName = data.PublicBucket
		} else {
			bucketName = data.PrivateBucket
		}

		if minioFile, err := utils.GetFileFromMinio(bucketName, fileName, context); err != nil {
			return nil, nil, nil, err
		} else {
			returnMinioFiles = append(returnMinioFiles, *minioFile)
		}
	}

	return &pictures, &returnMinioFiles, &count, nil
}

func createPictureService(data models.CreatePictureRequest, context *gin.Context) (*models.Picture, *models.ServiceError) {
	var isPublic bool
	var bucketName string

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

	picture.Type = filepath.Ext(data.Picture.Filename)

	if db := database.DB.Create(&picture); db.Error != nil {
		return nil,
			&models.ServiceError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to save picture",
			}
	}

	fileExt := filepath.Ext(data.Picture.Filename)
	bucketFileName := *data.Name + "__" + strconv.Itoa(int(picture.ID)) + fileExt

	_, err := utils.UploadMultipartFileToMinio(data.Picture, bucketName, bucketFileName, context)
	if err != nil {
		return nil, err
	}

	return &picture, nil
}

func updatePictureService(data models.UpdatePictureRequest, context *gin.Context) (*models.Picture, *models.ServiceError) {
	var picture models.Picture

	if err := database.DB.First(&picture, data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{StatusCode: http.StatusNotFound, Message: "Picture not found"}
		}
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	if data.Name != nil {
		picture.Name = *data.Name
	}

	fmt.Printf("Priority: %s \n", data.Name)
	if data.Priority != nil {
		picture.Priority = data.Priority
		fmt.Printf("Priority: %n \n", data.Priority)
	}

	if data.IsPublic != nil {
		if *data.IsPublic != picture.IsPublic {
			picture.IsPublic = *data.IsPublic

			bucketFileName := picture.Name + "__" + strconv.Itoa(int(picture.ID)) + picture.Type

			if picture.IsPublic {
				utils.TransferFileBetweenBuckets(bucketFileName, data.PrivateBucket, data.PublicBucket, context)
			} else {
				utils.TransferFileBetweenBuckets(bucketFileName, data.PublicBucket, data.PrivateBucket, context)
			}
		}
	}

	if err := database.DB.Save(&picture).Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update picture",
		}
	}

	return &picture, nil
}
