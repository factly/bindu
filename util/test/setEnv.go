package test

import "github.com/factly/bindu-server/config"

//SetEnv - to set .env
func SetEnv() {
	config.KavachURL = "http://kavach:5000"
}
