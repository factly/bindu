package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

// KetoAllowed is request object to check permissions of user
type KetoAllowed struct {
	Subject  string `json:"subject"`
	Action   string `json:"action"`
	Resource string `json:"resource"`
}

// CheckSpaceKetoPermission checks keto policy for operations on space
func CheckSpaceKetoPermission(action string, oID, uID uint) error {
	commonString := fmt.Sprint(":org:", oID, ":app:bindu:spaces")

	kresource := fmt.Sprint("resources", commonString)
	kaction := fmt.Sprint("actions", commonString, ":", action)

	result := KetoAllowed{}

	result.Action = kaction
	result.Resource = kresource
	result.Subject = fmt.Sprint(uID)

	resStatus, err := IsAllowed(result)
	if err != nil {
		return err
	}

	if resStatus != 200 {
		return errors.New("Permission not granted")
	}
	return nil
}

// IsAllowed checks if keto policy allows user to action on resource
// returns (status code, error)
func IsAllowed(result KetoAllowed) (int, error) {
	buf := new(bytes.Buffer)

	err := json.NewEncoder(buf).Encode(&result)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", viper.GetString("keto_url")+"/engines/acp/ory/regex/allowed", buf)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}
