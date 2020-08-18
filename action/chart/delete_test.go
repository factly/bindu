package chart

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-chi/chi"
	"gopkg.in/h2non/gock.v1"
)

func TestChartDelete(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/charts", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	theme := &model.Theme{
		Name:           "Theme sample",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Theme{}).Create(&theme)

	t.Run("invalid chart id", func(t *testing.T) {
		e.DELETE("/charts/{chart_id}").
			WithPath("chart_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("chart record not found", func(t *testing.T) {
		e.DELETE("/charts/{chart_id}").
			WithPath("chart_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("chart record deleted", func(t *testing.T) {

		chart := model.Chart{
			Title:          "Sample chart",
			ThemeID:        theme.Base.ID,
			OrganisationID: 1,
		}
		config.DB.Model(&model.Theme{}).Create(&chart)

		e.DELETE("/charts/{chart_id}").
			WithPath("chart_id", chart.Base.ID).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)

	})

}
