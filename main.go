package main

import (
	"log"
	"net/http"
	"os"

	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/config"
	"github.com/joho/godotenv"
)

// @title Bindu API
// @version 1.0
// @description Bindu API

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7000
// @BasePath /
func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("cant load .env")
	}

	DSN := os.Getenv("DSN")

	// db setup
	config.SetupDB(DSN)

	// register routes
	r := action.RegisterRoutes()

	err = http.ListenAndServe(":8000", r)

	if err != nil {
		log.Fatal(err)
	}
}
