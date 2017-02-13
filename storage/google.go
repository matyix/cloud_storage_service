package storage

import (
	"context"
	"errors"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"

	gstorage "cloud.google.com/go/storage"
)

type GoogleCloudStorage struct {
	Storage
	client *gstorage.Client
	bucket *gstorage.BucketHandle
}

func (gcs *GoogleCloudStorage) Setup() error {
	bktName := os.Getenv("GOOGLE_STORAGE_BUCKET")
	if bktName == "" {
		return errors.New("GOOGLE_STORAGE_BUCKET env must be set")
	}

	projectID := os.Getenv("GOOGLE_STORAGE_PROJECT_ID")
	if projectID == "" {
		return errors.New("GOOGLE_STORAGE_PROJECT_ID env must be set")
	}

	location := os.Getenv("GOOGLE_STORAGE_LOCATION")
	if location == "" {
		return errors.New("GOOGLE_STORAGE_LOCATION env must be set")
	}

	ctx := context.Background()
	client, err := gstorage.NewClient(ctx)
	if err != nil {
		return err
	}

	gcs.client = client

	bkt := client.Bucket(bktName)

	attrs := &gstorage.BucketAttrs{Location: location}
	err = bkt.Create(ctx, projectID, attrs)
	if err == nil {
		log.Printf("Created Google Cloud Storage bucket %s in %s",
			bktName, location)
	}

	if err != nil {
		if !strings.Contains(err.Error(), "You already own this bucket") {
			return err
		}

		log.Printf("Using existing Google Cloud Storage bucket %v", bktName)
	}

	gcs.bucket = bkt

	return nil
}

func (gcs *GoogleCloudStorage) Store(ctx context.Context, filename string, data []byte, metadata map[string]string) error {
	o := gcs.bucket.Object(filename)
	w := o.NewWriter(ctx)

	// w.Attrs().Metadata = metadata
	_, err := w.Write(data)
	if err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}