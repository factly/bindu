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

func TestThemeDetails(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Get("/themes/{theme_id}", details)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("invalid theme id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "GET", "/themes/invalid_id", nil, headers)

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
		_, statusCode := test.Request(t, ts, "GET", "/themes/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("get theme by id", func(t *testing.T) {
		theme := &model.Theme{
			Name:           "Theme sample",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Theme{}).Create(&theme)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		resp, statusCode := test.Request(t, ts, "GET", fmt.Sprint("/themes/", theme.Base.ID), nil, headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Theme sample" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Theme sample")
		}

	})

}
