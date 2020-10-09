package test

import "github.com/spf13/viper"

//SetEnv - to set .env
func SetEnv() {
	viper.Set("kavach.url", "http://kavach:5000")
	viper.Set("minio.url", "minio:9000")
	viper.Set("minio.bucket", "dega")
	viper.Set("minio.key", "miniokey")
	viper.Set("minio.secret", "miniosecret")
}
