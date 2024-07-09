package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type MinioConfig struct {
	MinioEndpoint            string
	MinioAccessKey           string
	MinioSecretKey           string
	MinioBucketName          string
	MinioUseSSL              bool
	MinioPreSignedURLExpired time.Duration
	MinioURL                 string
}

func LoadMinioConfig() MinioConfig {

	// Parse MINIO_USE_SSL
	useSSL, err := strconv.ParseBool(os.Getenv("MINIO_USE_SSL"))
	if err != nil {
		panic(err)
	}

	// Parse MINIO_PRE_SIGNED_URL_EXPIRES
	preSignedURLExpired, err := strconv.Atoi(os.Getenv("MINIO_PRE_SIGNED_URL_EXPIRES"))
	if err != nil {
		panic(err)
	}

	protocol := "https"
	if !useSSL {
		protocol = "http"
	}

	return MinioConfig{
		MinioEndpoint:            os.Getenv("MINIO_ENDPOINT"),
		MinioAccessKey:           os.Getenv("MINIO_ACCESS_KEY"),
		MinioSecretKey:           os.Getenv("MINIO_SECRET_KEY"),
		MinioBucketName:          os.Getenv("MINIO_BUCKET_NAME"),
		MinioUseSSL:              useSSL,
		MinioPreSignedURLExpired: time.Duration(preSignedURLExpired * 60 * 60),
		MinioURL:                 fmt.Sprintf("%s://%s/%s", protocol, os.Getenv("MINIO_ENDPOINT"), os.Getenv("MINIO_BUCKET_NAME")),
	}
}
