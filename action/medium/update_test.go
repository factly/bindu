package medium

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/test"
	"github.com/go-chi/chi"
)

func TestMediumUpdate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.GenerateOrganisation).Put("/media/{medium_id}", update)

	ts := httptest.NewServer(r)
	defer ts.Close()

	medium := &model.Medium{
		Name:           "Chart",
		Slug:           "chart",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Medium{}).Create(&medium)

	t.Run("invalid medium id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "PUT", "/media/invalid_id", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("medium record not found", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "PUT", "/media/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("update medium", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		var jsonStr = []byte(`
		{
			"name": "Chart sample image",
			"slug": "chart"
		}`)

		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/media/", medium.Base.ID), bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Chart sample image" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Chart sample image")
		}

	})

	t.Run("update medium by id with empty slug", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		var jsonStr = []byte(`
		{
			"name": "Chart sample image",
			"slug": ""
		}`)

		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/media/", medium.Base.ID), bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Chart sample image" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Chart sample image")
		}

	})

	t.Run("update medium with different slug", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		var jsonStr = []byte(`
		{
			"name": "Chart sample image",
			"slug": "chart-image"
		}`)

		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/media/", medium.Base.ID), bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Chart sample image" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Chart sample image")
		}

	})

}
