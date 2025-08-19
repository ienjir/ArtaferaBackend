package picture

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"github.com/minio/minio-go/v7"
	"path/filepath"
	"strconv"
)

var wrong = false

func getPictureByIDService(data models.GetPictureByIDRequest, context *gin.Context) (*models.Picture, *minio.Object, *models.ServiceError) {
	var returnMinioFile *minio.Object
	var bucketName string

	picture, err := database.Repositories.Picture.GetByID(data.TargetID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, nil, utils.NewPictureNotFoundError()
		}
		return nil, nil, err
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

	return picture, returnMinioFile, nil
}

func getPictureByNameService(data models.GetPictureByNameRequest, context *gin.Context) (*models.Picture, *minio.Object, *models.ServiceError) {
	var returnMinioFile *minio.Object
	var bucketName string

	picture, err := database.Repositories.Picture.FindByField("name", data.Name)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, nil, utils.NewPictureNotFoundError()
		}
		return nil, nil, err
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

	return picture, returnMinioFile, nil
}

func listPictureService(data models.ListPictureRequest, context *gin.Context) (*[]models.Picture, *[]minio.Object, *int64, *models.ServiceError) {
	var returnMinioFiles []minio.Object

	pictures, err := database.Repositories.Picture.List(int(data.Offset*5), 5)
	if err != nil {
		return nil, nil, nil, err
	}

	count, err := database.Repositories.Picture.Count()
	if err != nil {
		return nil, nil, nil, err
	}

	for _, picture := range *pictures {
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

	return pictures, &returnMinioFiles, count, nil
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

	if err := database.Repositories.Picture.Create(&picture); err != nil {
		return nil, err
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
	picture, err := database.Repositories.Picture.GetByID(data.TargetID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewPictureNotFoundError()
		}
		return nil, err
	}

	var originalFileName string
	var newFileName string
	var originalBucketName string
	var newBucketName string

	originalFileName = picture.Name + "__" + strconv.Itoa(int(picture.ID)) + picture.Type

	if data.Name != nil {
		picture.Name = *data.Name
		newFileName = *data.Name + "__" + strconv.Itoa(int(picture.ID)) + picture.Type
	} else {
		newFileName = picture.Name + "__" + strconv.Itoa(int(picture.ID)) + picture.Type
	}

	if data.IsPublic != nil && *data.IsPublic != picture.IsPublic {
		if *data.IsPublic != picture.IsPublic {
			if *data.IsPublic == true {
				picture.IsPublic = true
				originalBucketName = PrivateBucket
				newBucketName = PublicBucket
			} else {
				picture.IsPublic = false
				originalBucketName = PublicBucket
				newBucketName = PrivateBucket
			}
		}
	} else {
		if picture.IsPublic == true {
			originalBucketName = PublicBucket
			newBucketName = PublicBucket
		} else {
			originalBucketName = PrivateBucket
			newBucketName = PrivateBucket
		}
	}

	if data.Name != nil || data.IsPublic != nil {
		if err := utils.TransferAndRenameFile(originalFileName, newFileName, originalBucketName, newBucketName, context); err != nil {
			return nil, err
		}
	}

	if err := database.Repositories.Picture.Update(picture); err != nil {
		return nil, err
	}

	return picture, nil
}

func deletePictureService(data models.DeletePictureRequest, context *gin.Context) *models.ServiceError {
	var bucketName string

	picture, err := database.Repositories.Picture.GetByID(data.TargetID)
	if err != nil {
		if err.StatusCode == 404 {
			return utils.NewPictureNotFoundError()
		}
		return err
	}

	if picture.IsPublic {
		bucketName = PublicBucket
	} else {
		bucketName = PrivateBucket
	}

	fileName := picture.Name + "__" + strconv.Itoa(int(picture.ID)) + picture.Type
	_, deleteErr := utils.DeleteFile(fileName, bucketName, context)
	if deleteErr != nil {
		return deleteErr
	}

	if err := database.Repositories.Picture.DeleteEntity(picture); err != nil {
		return err
	}

	return nil
}
