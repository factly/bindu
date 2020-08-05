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

func TestCategoryDelete(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.GenerateOrganisation).Delete("/categories/{category_id}", delete)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("invalid category id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "DELETE", "/categories/invalid_id", nil, headers)

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
		_, statusCode := test.Request(t, ts, "DELETE", "/categories/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("check category associated with other entity", func(t *testing.T) {

		category := model.Category{
			Name:           "sample",
			OrganisationID: 1,
		}

		theme := &model.Theme{
			Name: "Sample Theme",
		}

		config.DB.Model(&model.Category{}).Create(&category)
		config.DB.Model(&model.Theme{}).Create(&theme)

		chart := &model.Chart{
			Title:          "Sample chart",
			OrganisationID: 1,
			ThemeID:        theme.Base.ID,
			Categories:     []model.Category{category},
		}

		config.DB.Model(&model.Chart{}).Create(&chart)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "DELETE", fmt.Sprint("/categories/", category.Base.ID), nil, headers)

		if statusCode != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnprocessableEntity)
		}
	})

	t.Run("category record deleted", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		result := &model.Category{
			Name:           "Cricket",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Category{}).Create(&result)

		_, statusCode := test.Request(t, ts, "DELETE", fmt.Sprint("/categories/", result.Base.ID), nil, headers)

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

}
