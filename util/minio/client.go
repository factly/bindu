package minio

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

// Client minio client
var Client *minio.Client

// SetupClient setups a minio client
func SetupClient() {
	var err error
	Client, err = minio.New(viper.GetString("minio_url"), &minio.Options{
		Creds:  credentials.NewStaticV4(viper.GetString("minio_key"), viper.GetString("minio_secret"), ""),
		Secure: false,
	})

	if err != nil {
		log.Fatal(err)
	}
}
