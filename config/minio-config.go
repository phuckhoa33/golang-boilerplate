package config

import (
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
}

func LoadMinioConfig() MinioConfig {

	// Parse MINIO_USE_SSL
	useSSL, err := strconv.ParseBool(os.Getenv("MINIO_USE_SSL"))
	if err != nil {
		panic(err)
	}

	// Parse MINIO_PRESIGNED_URL_EXPIRES
	preSignedURLExpired, err := strconv.Atoi(os.Getenv("MINIO_PRESIGNED_URL_EXPIRES"))
	if err != nil {
		panic(err)
	}

	return MinioConfig{
		MinioEndpoint:            os.Getenv("MINIO_ENDPOINT"),
		MinioAccessKey:           os.Getenv("MINIO_ACCESS_KEY"),
		MinioSecretKey:           os.Getenv("MINIO_SECRET_KEY"),
		MinioBucketName:          os.Getenv("MINIO_BUCKET_NAME"),
		MinioUseSSL:              useSSL,
		MinioPreSignedURLExpired: time.Duration(preSignedURLExpired * 60 * 60),
	}
}
