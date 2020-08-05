package theme

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

func TestThemeDelete(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.GenerateOrganisation).Delete("/themes/{theme_id}", delete)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("invalid theme id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "DELETE", "/themes/invalid_id", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("theme record not found", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "DELETE", "/themes/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("check theme associated with other entity", func(t *testing.T) {

		theme := &model.Theme{
			Name:           "Sample Theme",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Theme{}).Create(&theme)

		chart := &model.Chart{
			Title:          "Sample chart",
			OrganisationID: 1,
			ThemeID:        theme.Base.ID,
		}

		config.DB.Model(&model.Chart{}).Create(&chart)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "DELETE", fmt.Sprint("/themes/", theme.Base.ID), nil, headers)

		if statusCode != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnprocessableEntity)
		}
	})

	t.Run("theme record deleted", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		result := &model.Theme{
			Name:           "testing",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Theme{}).Create(&result)

		_, statusCode := test.Request(t, ts, "DELETE", fmt.Sprint("/themes/", result.Base.ID), nil, headers)

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

}
