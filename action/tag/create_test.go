package tag

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/test"
	"github.com/go-chi/chi"
)

func TestTagCreate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Post("/tags", create)

	var jsonStr = []byte(`
	{
		"name": "Corruption",
		"slug": "corruption"
	}`)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("Unprocessable tag", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "POST", "/tags", nil, headers)

		if statusCode != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnprocessableEntity)
		}
	})

	t.Run("create tag", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "POST", "/tags", bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusCreated)
		}

		if respBody["name"] != "Corruption" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Corruption")
		}

	})

	t.Run("create tags with slug is empty", func(t *testing.T) {

		jsonStr = []byte(`
		{
			"name": "Murder"
		}`)
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "POST", "/tags", bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusCreated)
		}

		if respBody["slug"] != "murder" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["slug"], "murder")
		}

	})

}
