package category

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/test"
	"github.com/go-chi/chi"
	"gopkg.in/h2non/gock.v1"
)

func TestCategoryCreate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/categories", Router())

	var jsonStr = []byte(`
	{
		"name": "Politics",
		"slug": "politics"
	}`)

	ts := httptest.NewServer(r)
	gock.New(ts.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer ts.Close()

	t.Run("Unprocessable category", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "POST", "/categories", nil, headers)

		if statusCode != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnprocessableEntity)
		}
	})

	t.Run("create category", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "POST", "/categories", bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusCreated)
		}

		if respBody["name"] != "Politics" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Politics")
		}

	})

	t.Run("create categories with slug is empty", func(t *testing.T) {

		jsonStr = []byte(`
		{
			"name": "Crime"
		}`)
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "POST", "/categories", bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusCreated)
		}

		if respBody["slug"] != "crime" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["slug"], "crime")
		}

	})

}
