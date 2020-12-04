package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// GetRequest does get request to keto with empty body
func GetRequest(path string, body interface{}) (*http.Response, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(&body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", path, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
