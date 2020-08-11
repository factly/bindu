package chart

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

func TestChartDetails(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.GenerateOrganisation).Get("/charts/{chart_id}", details)

	ts := httptest.NewServer(r)
	defer ts.Close()

	theme := &model.Theme{
		Name:           "Theme sample",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Theme{}).Create(&theme)

	t.Run("invalid chart id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "GET", "/charts/invalid_id", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("chart record not found", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "GET", "/charts/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("get chart by id", func(t *testing.T) {
		chart := model.Chart{
			Title:          "Sample chart",
			ThemeID:        theme.Base.ID,
			OrganisationID: 1,
		}

		config.DB.Model(&model.Chart{}).Create(&chart)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		resp, statusCode := test.Request(t, ts, "GET", fmt.Sprint("/charts/", chart.Base.ID), nil, headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["title"] != "Sample chart" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["title"], "Sample chart")
		}

	})

}
