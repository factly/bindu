package test

import "os"

//SetEnv - to set .env
func SetEnv() {
	os.Setenv("KAVACH_URL", "http://kavach:5000")
}
