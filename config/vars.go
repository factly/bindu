package config

import (
	"log"

	"github.com/spf13/viper"
)

// SetupVars setups all the config variables to run application
func SetupVars() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("config file not found...")
	}

	if !viper.IsSet("dsn") {
		log.Fatal("please provide dsn in config file")
	}

	if !viper.IsSet("kavach_url") {
		log.Fatal("please provide kavach_url in config file")
	}

	if !viper.IsSet("minio_url") {
		log.Fatal("please provide minio_url in config file")
	}

	if !viper.IsSet("minio_key") {
		log.Fatal("please provide minio_key in config file")
	}

	if !viper.IsSet("minio_secret") {
		log.Fatal("please provide minio_secret in config file")
	}

	if !viper.IsSet("minio_bucket") {
		log.Fatal("please provide minio_bucket in config file")
	}
}
