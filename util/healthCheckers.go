package util

import (
	"errors"
	"net/http"

	"github.com/spf13/viper"
)

// KavachChecker checks if Kavach is ready
func KavachChecker() error {
	return GetRequest(viper.GetString("kavach_url") + "/health/ready")
}

// KetoChecker checks if keto is ready
func KetoChecker() error {
	return GetRequest(viper.GetString("keto_url") + "/health/ready")
}

// GetRequest returns error if error in status code
func GetRequest(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= 500 {
		return errors.New("cannot connect")
	}
	return nil
}
