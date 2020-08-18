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

func TestThemeUpdate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/themes", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	theme := &model.Theme{
		Name:           "Light",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Theme{}).Create(&theme)

	t.Run("invalid theme id", func(t *testing.T) {
		e.PUT("/themes/{theme_id}").
			WithPath("theme_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("theme record not found", func(t *testing.T) {

		e.PUT("/themes/{theme_id}").
			WithPath("theme_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("update theme", func(t *testing.T) {

		body := model.Theme{
			Name: "Dark",
		}

		resObj := e.PUT("/themes/{theme_id}").
			WithPath("theme_id", theme.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("name").String().Equal("Dark")

	})

}
