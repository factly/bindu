package test

import (
	"net/http"
	"os"
	"time"

	"gopkg.in/h2non/gock.v1"
)

// MockServer is created to intercept HTTP Calls outside this project. Mocking the external project servers helps with Unit Testing.
func MockServer() {
	// Creates a mock server for kavach URL with an appropriate dummy response.

	data := map[string]interface{}{
		"id":         1,
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"deleted_at": nil,
		"title":      "test org",
		"slug":       "tesing",
		"permission": map[string]interface{}{
			"id":         1,
			"created_at": time.Now(),
			"updated_at": time.Now(),
			"deleted_at": nil,
			"role":       "owner",
		},
	}

	res := make([]interface{}, 1)

	res[0] = data

	gock.New(os.Getenv("KAVACH_URL")).
		Get("/organisations").
		Persist().
		Reply(http.StatusOK).
		JSON(res)
}
