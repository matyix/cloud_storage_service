package storage

import (
	"os"
	"errors"
	"context"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSS3Storage struct {
	Storage
	bucketName string
	bucket     *s3.Bucket
}

type Credential struct {
	AccessKey, SecretKey, Token string
}

func EnvAuth() (credential Credential, err error) {
	credential.AccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	if credential.AccessKey == "" {
		credential.AccessKey = os.Getenv("AWS_ACCESS_KEY")
	}

	credential.SecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	if credential.SecretKey == "" {
		credential.SecretKey = os.Getenv("AWS_SECRET_KEY")
	}

	credential.Token = os.Getenv("AWS_SECURITY_TOKEN")

	if credential.AccessKey == "" {
		err = errors.New("AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY not found in environment")
	}
	if credential.SecretKey == "" {
		err = errors.New("AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY not found in environment")
	}
	return
}

func (aws *AWSS3Storage) Setup() error {
	//ctx := context.Background()
	return nil
}

func (aws *AWSS3Storage) PublicURL(filename string) string {
	return "Bucket URL" + *s3.Bucket.Name
}

func (aws *AWSS3Storage) Store(ctx context.Context, filename string, data []byte, metadata map[string]string) error {
	return nil
}

func (aws *AWSS3Storage) Delete(ctx context.Context, filename string) error {
	return nil
}
