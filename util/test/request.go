package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
)

// Request - test request
func Request(t *testing.T, ts *httptest.Server, method, path string, body io.Reader, header map[string]string) (interface{}, int) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, http.StatusServiceUnavailable
	}

	req.Header = map[string][]string{
		"X-User":         {header["X-User"]},
		"X-Organisation": {header["X-Organisation"]},
	}

	req.Close = true

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, http.StatusServiceUnavailable
	}

	var respBody interface{}

	json.NewDecoder(resp.Body).Decode(&respBody)
	defer resp.Body.Close()

	return respBody, resp.StatusCode
}

// CleanTables - to clean tables in DB
func CleanTables() {
	config.DB.Model(&model.Chart{}).RemoveForeignKey("featured_medium_id", "bi_medium(id)")
	config.DB.Model(&model.Chart{}).RemoveForeignKey("theme_id", "bi_theme(id)")

	config.DB.DropTable("bi_chart_category")
	config.DB.DropTable("bi_chart_tag")
	config.DB.DropTable(&model.Medium{})
	config.DB.DropTable(&model.Theme{})
	config.DB.DropTable(&model.Tag{})
	config.DB.DropTable(&model.Category{})
	config.DB.DropTable(&model.Chart{})
}
