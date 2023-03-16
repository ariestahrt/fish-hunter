package s3

import (
	"context"
	"fish-hunter/util"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3config struct {
	Region string
	Bucket string
	SecretKey string
	AccessKey string
}
type S3 struct {
	client *s3.Client
	bucket string
}

type AWS_S3 interface {
	UploadFile(file *os.File, key string) error
	DownloadFile(file *os.File, key string) error
}

func NewAWS_S3() AWS_S3 {
	creds := credentials.NewStaticCredentialsProvider(util.GetConfig("AWS_ACCESS_KEY"), util.GetConfig("AWS_SECRET_KEY"), "")

	// Load config from environment
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(util.GetConfig("AWS_REGION")))
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// Create S3 service client
	client := s3.NewFromConfig(cfg)
	return &S3{
		client: client,
		bucket: util.GetConfig("AWS_BUCKET"),
	}
}

func (s *S3) UploadFile(file *os.File, key string) error {
	uploader := manager.NewUploader(s.client)
	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}

func (s *S3) DownloadFile(file *os.File, key string) error {
	fmt.Println("Downloading file from s3", key)
	downloader := manager.NewDownloader(s.client)
	_, err := downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})

	return err
}