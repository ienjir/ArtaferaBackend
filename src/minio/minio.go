package miniobucket

import (
	"context"
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

	return nil
}

func CreateMinioBuckets() error {
	ctx := context.Background()
	Buckets := []models.MinioBucket{
		{BucketName: "pictures", Location: "CH-CENTER-1"},
	}

	for i := 0; i < len(Buckets); i++ {
		err := MinioClient.MakeBucket(ctx, Buckets[i].BucketName, minio.MakeBucketOptions{Region: Buckets[i].Location})
		if err != nil {
			exists, errBucketExists := MinioClient.BucketExists(ctx, Buckets[i].BucketName)
			if errBucketExists == nil && exists {
				log.Printf("Bucket: %s already exists \ny", Buckets[i].BucketName)
			} else {
				log.Fatalln(err)
				return err
			}
		} else {
			log.Printf("Successfully created %s\n", Buckets[i].BucketName)
		}
	}

	return nil
}
