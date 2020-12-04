package util

import (
	"errors"

	"github.com/spf13/viper"
)

// KavachChecker checks if Kavach is ready
func KavachChecker() error {
	res, err := GetRequest(viper.GetString("kavach_url")+"/health/ready", nil)
	if err != nil {
		return err
	}

	if res.StatusCode >= 500 {
		return errors.New("cannot connect due to some error")
	}

	return nil
}

// KetoChecker checks if keto is ready
func KetoChecker() error {
	res, err := GetRequest(viper.GetString("keto_url")+"/health/ready", nil)
	if err != nil {
		return err
	}

	if res.StatusCode >= 500 {
		return errors.New("cannot connect due to some error")
	}

	return nil
}
