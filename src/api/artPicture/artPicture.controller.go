package artPicture

import (
	"context"
	"fmt"
	mino "github.com/ienjir/ArtaferaBackend/src/minio"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

const (
	UploadDir  = "./uploads"
	BucketName = "artpictures"
)

func Upload(c *gin.Context) {
	// Initialize MinIO client
	minioClient, err := mino.InitMinIO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize MinIO client"})
		return
	}

	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	// Ensure the bucket exists
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, BucketName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking bucket existence"})
		return
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, BucketName, minio.MakeBucketOptions{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bucket"})
			return
		}
	}

	// Upload the file to MinIO
	objectName := filepath.Base(file.Filename)
	_, err = minioClient.PutObject(ctx, BucketName, objectName, src, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to MinIO"})
		return
	}

	// Return the file URL
	fileURL := fmt.Sprintf("http://%s/%s/%s", minioClient.EndpointURL().Host, BucketName, objectName)
	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"url":      fileURL,
		"filename": objectName,
	})
}
