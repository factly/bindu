package tag

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

func TestTagUpdate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Put("/tags/{tag_id}", update)

	ts := httptest.NewServer(r)
	defer ts.Close()

	tag := &model.Tag{
		Name:           "Test",
		Slug:           "test",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Tag{}).Create(&tag)

	t.Run("invalid tag id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "PUT", "/tags/invalid_id", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("tag record not found", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "PUT", "/tags/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("update tag", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		var jsonStr = []byte(`
		{
			"name": "Tag-1",
			"slug": "test"
		}`)

		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/tags/", tag.Base.ID), bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Tag-1" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Tag-1")
		}

	})

	t.Run("update tag by id with empty slug", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		var jsonStr = []byte(`
		{
			"name": "Tag-2",
			"slug": ""
		}`)

		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/tags/", tag.Base.ID), bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Tag-2" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Tag-2")
		}

	})

	t.Run("update tag with different slug", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		var jsonStr = []byte(`
		{
			"name": "Test",
			"slug": "testing"
		}`)

		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/tags/", tag.Base.ID), bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Test" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Test")
		}

	})

}
