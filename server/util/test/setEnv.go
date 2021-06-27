package test

import "github.com/spf13/viper"

//SetEnv - to set .env
func SetEnv() {
	viper.Set("kavach_url", "http://kavach:5000")
	viper.Set("keto_url", "http://keto:4466")
	viper.Set("minio_url", "minio:9000")
	viper.Set("minio_bucket", "dega")
	viper.Set("minio_key", "miniokey")
	viper.Set("minio_secret", "miniosecret")

	viper.Set("create_super_organisation", true)
}
