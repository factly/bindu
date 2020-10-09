package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

// SetupVars setups all the config variables to run application
func SetupVars() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config.yaml", "Config file path")
	flag.Parse()

	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("config file not found...")
	}

	if !viper.IsSet("postgres.dsn") {
		log.Fatal("please provide postgres.dsn in config file")
	}

	if !viper.IsSet("kavach.url") {
		log.Fatal("please provide kavach.url in config file")
	}

	if !viper.IsSet("minio.url") {
		log.Fatal("please provide minio.url in config file")
	}

	if !viper.IsSet("minio.key") {
		log.Fatal("please provide minio.key in config file")
	}

	if !viper.IsSet("minio.secret") {
		log.Fatal("please provide minio.secret in config file")
	}

	if !viper.IsSet("minio.bucket") {
		log.Fatal("please provide minio.bucket in config file")
	}
}
