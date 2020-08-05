package medium

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/test"
	"github.com/go-chi/chi"
)

func TestCreateMedium(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.GenerateOrganisation).Post("/media", create)

	var jsonStr = []byte(`
	{
		"name": "pie-chart",
		"slug": "pie-chart"
	}`)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("media title required", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "POST", "/media", nil, headers)

		if statusCode != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnprocessableEntity)
		}
	})

	t.Run("create media", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "POST", "/media", bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusCreated)
		}

		if respBody["name"] != "pie-chart" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "pie-chart")
		}

	})

	t.Run("slug is empty", func(t *testing.T) {

		jsonStr = []byte(`
		{
			"name": "bar-chart"
		}`)
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "POST", "/media", bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusCreated)
		}

		if respBody["slug"] != "bar-chart" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["slug"], "bar-chart")
		}

	})

}
