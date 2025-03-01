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
