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

func TestChartDelete(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.GenerateOrganisation).Delete("/charts/{chart_id}", delete)

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
		_, statusCode := test.Request(t, ts, "DELETE", "/charts/invalid_id", nil, headers)

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
		_, statusCode := test.Request(t, ts, "DELETE", "/charts/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("chart record deleted", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		chart := model.Chart{
			Title:          "Sample chart",
			ThemeID:        theme.Base.ID,
			OrganisationID: 1,
		}

		config.DB.Model(&model.Chart{}).Create(&chart)

		_, statusCode := test.Request(t, ts, "DELETE", fmt.Sprint("/charts/", chart.Base.ID), nil, headers)

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

}
