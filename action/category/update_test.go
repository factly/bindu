package category

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

func TestCategoryUpdate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.GenerateOrganisation).Put("/categories/{category_id}", update)

	ts := httptest.NewServer(r)
	defer ts.Close()

	category := &model.Category{
		Name:           "Test",
		Slug:           "test",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Category{}).Create(&category)

	t.Run("invalid category id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "PUT", "/categories/invalid_id", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("category record not found", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "PUT", "/categories/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("update category", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		var jsonStr = []byte(`
		{
			"name": "Category-1",
			"slug": "test"
		}`)

		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/categories/", category.Base.ID), bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Category-1" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Category-1")
		}

	})

	t.Run("update category by id with empty slug", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		var jsonStr = []byte(`
		{
			"name": "Category-2",
			"slug": ""
		}`)

		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/categories/", category.Base.ID), bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Category-2" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Category-2")
		}

	})

	t.Run("update category with different slug", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		var jsonStr = []byte(`
		{
			"name": "Category-test",
			"slug": "testing"
		}`)

		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/categories/", category.Base.ID), bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Category-test" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Category-test")
		}

	})

}
