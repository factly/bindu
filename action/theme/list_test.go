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

func TestThemeList(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/themes", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("get themes with pagination", func(t *testing.T) {
		themeOne := &model.Theme{
			Name:           "Light theme",
			OrganisationID: 1,
		}
		themeTwo := &model.Theme{
			Name:           "Dark theme",
			OrganisationID: 1,
		}
		themes := []model.Theme{}
		total := 0

		config.DB.Model(&model.Theme{}).Create(&themeOne)
		config.DB.Model(&model.Theme{}).Create(&themeTwo)
		config.DB.Model(&model.Theme{}).Where(&model.Theme{
			OrganisationID: 1,
		}).Count(&total).Order("id desc").Offset(1).Limit(1).Find(&themes)

		resObj := e.GET("/themes").
			WithQueryObject(map[string]string{
				"limit": "1",
				"page":  "2",
			}).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("nodes").Array().Element(0).Object().Value("name").String().Equal(themes[0].Name)
		resObj.Value("total").Number().Equal(total)

	})
}
