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

func TestMediumCreate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.GenerateOrganisation).Post("/media", create)

	var jsonStr = []byte(`
	{
		"name": "Pie chart",
		"slug": "pie-chart"
	}`)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("Unprocessable medium", func(t *testing.T) {
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

	t.Run("create medium", func(t *testing.T) {
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

		if respBody["name"] != "Pie chart" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Pie chart")
		}

	})

	t.Run("create medium with slug is empty", func(t *testing.T) {

		jsonStr = []byte(`
		{
			"name": "Bar"
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

		if respBody["slug"] != "bar" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["slug"], "bar")
		}

	})

}
