package minio_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang-boilerplate/config"
	"net/url"
)

type MinioClientService struct {
	MinioService *minio.Client
	Config       *config.Config
}

func NewMinioClientService(config *config.Config) *MinioClientService {
	minioClient, err := minio.New(config.Minio.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Minio.MinioAccessKey, config.Minio.MinioSecretKey, ""),
		Secure: config.Minio.MinioUseSSL,
	})

	if err != nil {
		panic(err)
	}

	return &MinioClientService{
		MinioService: minioClient,
		Config:       config,
	}
}

// GetPutPreSignedURL PreSigned PUT URL
func (mcs *MinioClientService) GetPutPreSignedURL(ctx *gin.Context, objectName string) (*url.URL, error) {
	return mcs.MinioService.PresignedPutObject(ctx, mcs.Config.Minio.MinioBucketName, objectName, 24*60*60*60*60*60*60)
}

// GetFileURL Get file url
func (mcs *MinioClientService) GetFileURL(objectName string) string {
	useSSL := mcs.Config.Minio.MinioUseSSL
	protocol := "https"
	if !useSSL {
		protocol = "http"
	}

	return fmt.Sprintf("%s://%s/%s/%s", protocol, mcs.Config.Minio.MinioEndpoint, mcs.Config.Minio.MinioBucketName, objectName)
}
