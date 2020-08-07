package main

import (
	"fmt"
	"net/http"

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

// @host localhost:8000
// @BasePath /
func main() {

	godotenv.Load()

	// db setup
	config.SetupDB()

	// register routes
	r := action.RegisterRoutes()

	fmt.Println("swagger-ui http://localhost:6620/swagger/index.html")
	http.ListenAndServe(":8000", r)
}
