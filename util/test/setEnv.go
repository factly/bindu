package test

import "github.com/spf13/viper"

//SetEnv - to set .env
func SetEnv() {
	viper.Set("kavach.url", "http://kavach:5000")
}
