package theme

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

func TestThemeDelete(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/themes", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid theme id", func(t *testing.T) {

		e.DELETE("/themes/{theme_id}").
			WithPath("theme_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)

	})

	t.Run("theme record not found", func(t *testing.T) {
		e.DELETE("/themes/{theme_id}").
			WithPath("theme_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("check theme associated with other entity", func(t *testing.T) {

		theme := model.Theme{
			Name:           "Entertainment",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Theme{}).Create(&theme)

		chart := &model.Chart{
			Title:          "Bar chart",
			OrganisationID: 1,
			ThemeID:        theme.Base.ID,
		}

		config.DB.Model(&model.Chart{}).Create(&chart)

		e.DELETE("/themes/{theme_id}").
			WithPath("theme_id", theme.Base.ID).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)
	})

	t.Run("theme record deleted", func(t *testing.T) {
		theme := &model.Theme{
			Name:           "History",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Theme{}).Create(&theme)

		e.DELETE("/themes/{theme_id}").
			WithPath("theme_id", theme.Base.ID).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)
	})

}
