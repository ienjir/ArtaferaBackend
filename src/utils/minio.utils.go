package utils

import (
	"context"
	"fmt"
	miniobucket "github.com/ienjir/ArtaferaBackend/src/minio"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

func UploadToMinio(file *multipart.FileHeader, fileName string) (string, error) {
	bucketName := "art-pictures"
	ctx := context.Background()

	// Ensure the bucket exists
	exists, err := miniobucket.MinioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return "", err
	}
	if !exists {
		_ = miniobucket.MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Upload file
	_, err = miniobucket.MinioClient.PutObject(ctx, bucketName, fileName, src, -1, minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://localhost:9000/%s/%s", bucketName, fileName), nil
}
