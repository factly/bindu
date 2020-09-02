package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type BucketDetails struct {
	BucketName string `json:"bucket_name"`
	Path       string `json:"path"`
	Limit      int    `json:"limit"`
}

type BucketConfig struct {
	Details BucketDetails
	Keys    map[string]string
}

// GetBucketConfig bucket config
func GetBucketConfig() (BucketConfig, []byte, error) {
	var bucketConfig BucketConfig
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(dir, "bucket.json")

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return bucketConfig, nil, err
	}

	err = json.Unmarshal(content, &bucketConfig)
	if err != nil {
		return bucketConfig, nil, err
	}

	bucketByte, err := json.Marshal(bucketConfig.Keys)

	if err != nil {
		return bucketConfig, nil, err
	}

	return bucketConfig, bucketByte, nil
}
