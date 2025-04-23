package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	miniobucket "github.com/ienjir/ArtaferaBackend/src/minio"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/minio/minio-go/v7"
	"io"
	"mime/multipart"
	"net/http"
)

func UploadMultipartFileToMinio(file multipart.FileHeader, bucketName string, fileName string, context *gin.Context) (*minio.UploadInfo, *models.ServiceError) {
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

func UploadReaderToMinio(reader io.Reader, size int64, contentType, bucketName, fileName string, context *gin.Context) (*minio.UploadInfo, *models.ServiceError) {
	info, err := miniobucket.MinioClient.PutObject(context, bucketName, fileName, reader, size, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while uploading file to target bucket",
		}
	}

	return &info, nil
}

func GetFileFromMinio(bucketName string, fileName string, context *gin.Context) (*minio.Object, *models.ServiceError) {
	_, err := miniobucket.MinioClient.StatObject(context, bucketName, fileName, minio.StatObjectOptions{})
	if err != nil {
		fmt.Printf("Error: %s \n", err.Error())
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

func DeleteFile(filename string, bucketName string, context *gin.Context) (*string, *models.ServiceError) {
	err := miniobucket.MinioClient.RemoveObject(context, bucketName, filename, minio.RemoveObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "File not found",
			}
		}
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while deleting file",
		}
	}

	success := "File successfully deleted"
	return &success, nil
}

func TransferAndRenameFile(originalFileName string, targetFileName string, originalBucketName string, targetBucketName string, context *gin.Context) *models.ServiceError {
	// Define the source object
	src := minio.CopySrcOptions{
		Bucket: originalBucketName,
		Object: originalFileName,
	}

	fmt.Printf("Original: %s \n", src.Bucket)

	// Define the destination object
	dst := minio.CopyDestOptions{
		Bucket: targetBucketName,
		Object: targetFileName,
	}

	_, err := miniobucket.MinioClient.CopyObject(context, dst, src)
	if err != nil {
		fmt.Printf("Error: %s \n", err.Error())
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "File not found",
			}
		}
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while getting file",
		}
	}

	_, errDelete := DeleteFile(originalFileName, originalBucketName, context)
	if errDelete != nil {
		return errDelete
	}

	return nil
}
