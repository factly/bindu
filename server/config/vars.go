package config

import (
	"log"

	"github.com/spf13/viper"
)

// SetupVars setups all the config variables to run application
func SetupVars() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetEnvPrefix("bindu")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("config file not found...")
	}

	if !viper.IsSet("database_host") {
		log.Fatal("please provide database_host config param")
	}

	if !viper.IsSet("database_user") {
		log.Fatal("please provide database_user config param")
	}

	if !viper.IsSet("database_name") {
		log.Fatal("please provide database_name config param")
	}

	if !viper.IsSet("database_password") {
		log.Fatal("please provide database_password config param")
	}

	if !viper.IsSet("database_port") {
		log.Fatal("please provide database_port config param")
	}

	if !viper.IsSet("database_ssl_mode") {
		log.Fatal("please provide database_ssl_mode config param")
	}

	if !viper.IsSet("kavach_url") {
		log.Fatal("please provide kavach_url in config file")
	}

	if !viper.IsSet("keto_url") {
		log.Fatal("please provide keto_url in config file")
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

	if !viper.IsSet("meili_url") {
		log.Fatal("please provide meili_url in config file")
	}

	if !viper.IsSet("meili_key") {
		log.Fatal("please provide meili_key in config file")
	}

	if !viper.IsSet("templates_path") {
		log.Fatal("please provide templates_path in config file")
	}

	if Sqlite() {
		if !viper.IsSet("sqlite_db_path") {
			log.Fatal("please provide sqlite_db_path config param")
		}
	}
}

func Sqlite() bool {
	return viper.IsSet("use_sqlite") && viper.GetBool("use_sqlite")
}
