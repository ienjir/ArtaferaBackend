package utils

import (
	"github.com/gin-gonic/gin"
	miniobucket "github.com/ienjir/ArtaferaBackend/src/minio"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"net/http"
)

func UploadFileToMinio(file multipart.FileHeader, bucketName string, fileName string, context *gin.Context) (*minio.UploadInfo, *models.ServiceError) {
	openFile, err := file.Open()
	if err != nil {
		return nil,
			&models.ServiceError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error while opening file",
			}
	}
	defer openFile.Close()

	info, err := miniobucket.MinioClient.PutObject(context, bucketName, fileName, openFile, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		return nil,
			&models.ServiceError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error while saving file",
			}
	}

	return &info, nil
}

func GetFileFromMinio(bucketName string, fileName string, context *gin.Context) (*minio.Object, *models.ServiceError) {
	_, err := miniobucket.MinioClient.StatObject(context, bucketName, fileName, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "File not found",
			}
		}
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while getting file",
		}
	}

	object, err := miniobucket.MinioClient.GetObject(context, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving file",
		}
	}

	return object, nil
}
