package category

import (
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

func TestCategoryDetails(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Get("/categories/{category_id}", details)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("invalid category id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "GET", "/categories/invalid_id", nil, headers)

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
		_, statusCode := test.Request(t, ts, "GET", "/categories/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("get category by id", func(t *testing.T) {
		category := &model.Category{
			Name:           "Sports",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Category{}).Create(&category)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		resp, statusCode := test.Request(t, ts, "GET", fmt.Sprint("/categories/", category.Base.ID), nil, headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Sports" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Sports")
		}

	})

}
