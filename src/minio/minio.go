package miniobucket

import (
	"context"
	"fmt"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
	"strconv"
)

var MinioClient *minio.Client

func InitMinIO() error {
	var err error

	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL, err := strconv.ParseBool(os.Getenv("MINIO_USE_SSL"))
	if err != nil {
		log.Fatalln("Error while getting use ssl bool from env")
	}

	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln("Error initialize minio client: ", err)
		return err
	}

	log.Println("Successfully initialized minio client")
	return nil
}

func CreateMinioBuckets() error {
	ctx := context.Background()

	buckets := []models.MinioBucket{
		{BucketName: "pictures-p", Location: "CH-CENTER-1", IsPublic: false},
		{BucketName: "pictures", Location: "CH-CENTER-1", IsPublic: true},
	}

	for _, bucket := range buckets {
		// Check if bucket exists
		exists, err := MinioClient.BucketExists(ctx, bucket.BucketName)
		if err != nil {
			log.Printf("Error checking bucket %s: %v\n", bucket.BucketName, err)
			return err
		}

		if !exists {
			err = MinioClient.MakeBucket(ctx, bucket.BucketName, minio.MakeBucketOptions{Region: bucket.Location})
			if err != nil {
				log.Printf("Failed to create bucket %s: %v\n", bucket.BucketName, err)
				return err
			}
			log.Printf("Successfully created bucket: %s\n", bucket.BucketName)
		} else {
			log.Printf("Bucket already exists: %s\n", bucket.BucketName)
		}

		if bucket.IsPublic {
			err = setPublicPolicy(ctx, bucket.BucketName)
			if err != nil {
				log.Printf("Failed to set public policy for bucket %s: %v\n", bucket.BucketName, err)
				return err
			}
			log.Printf("Public access enabled for bucket: %s\n", bucket.BucketName)
		}
	}

	log.Println("All MinIO buckets processed successfully")
	return nil
}

func setPublicPolicy(ctx context.Context, bucketName string) error {
	policy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": {"AWS": ["*"]},
				"Action": ["s3:GetObject"],
				"Resource": ["arn:aws:s3:::%s/*"]
			}
		]
	}`, bucketName)

	return MinioClient.SetBucketPolicy(ctx, bucketName, policy)
}

func DeleteAllBuckets() error {
	ctx := context.Background()

	buckets, err := MinioClient.ListBuckets(ctx)
	if err != nil {
		log.Printf("Failed to list buckets: %v\n", err)
		return err
	}

	for _, bucket := range buckets {
		log.Printf("Deleting all objects from bucket: %s\n", bucket.Name)

		err := clearBucket(ctx, bucket.Name)
		if err != nil {
			log.Printf("Failed to clear bucket %s: %v\n", bucket.Name, err)
			return err
		}

		err = MinioClient.RemoveBucket(ctx, bucket.Name)
		if err != nil {
			log.Printf("Failed to delete bucket %s: %v\n", bucket.Name, err)
			return err
		}

		log.Printf("Successfully deleted bucket: %s\n", bucket.Name)
	}

	log.Println("All buckets deleted successfully")
	return nil
}

func ClearAllBuckets() error {
	ctx := context.Background()

	buckets, err := MinioClient.ListBuckets(ctx)
	if err != nil {
		log.Printf("Failed to list buckets: %v\n", err)
		return err
	}

	for _, bucket := range buckets {
		log.Printf("Clearing bucket: %s\n", bucket.Name)

		err := clearBucket(ctx, bucket.Name)
		if err != nil {
			log.Printf("Failed to clear bucket %s: %v\n", bucket.Name, err)
			return err
		}

		log.Printf("Successfully cleared bucket: %s\n", bucket.Name)
	}

	log.Println("All buckets cleared successfully")
	return nil
}

func clearBucket(ctx context.Context, bucketName string) error {
	objectCh := MinioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Recursive: true})

	for object := range objectCh {
		if object.Err != nil {
			log.Printf("Error listing object %s in bucket %s: %v\n", object.Key, bucketName, object.Err)
			return object.Err
		}

		err := MinioClient.RemoveObject(ctx, bucketName, object.Key, minio.RemoveObjectOptions{})
		if err != nil {
			log.Printf("Failed to delete object %s from bucket %s: %v\n", object.Key, bucketName, err)
			return err
		}

		log.Printf("Deleted object: %s from bucket: %s\n", object.Key, bucketName)
	}

	return nil
}
