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

func TestChartList(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/charts", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("get charts with pagination", func(t *testing.T) {
		theme := &model.Theme{
			Name:           "Theme sample",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Theme{}).Create(&theme)

		chartOne := &model.Chart{
			Title:          "Sample-1",
			ThemeID:        theme.Base.ID,
			OrganisationID: 1,
		}
		chartTwo := &model.Chart{
			Title:          "Sample-2",
			ThemeID:        theme.Base.ID,
			OrganisationID: 1,
		}
		charts := []model.Chart{}
		total := 0

		config.DB.Model(&model.Chart{}).Create(&chartOne)
		config.DB.Model(&model.Chart{}).Create(&chartTwo)
		config.DB.Model(&model.Chart{}).Where(&model.Chart{
			OrganisationID: 1,
		}).Count(&total).Order("id desc").Offset(1).Limit(1).Find(&charts)

		resObj := e.GET("/charts").
			WithQueryObject(map[string]string{
				"limit": "1",
				"page":  "2",
			}).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("nodes").Array().Element(0).Object().Value("title").String().Equal(charts[0].Title)
		resObj.Value("total").Number().Equal(total)

	})

}
