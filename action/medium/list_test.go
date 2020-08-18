package medium

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

func TestMediumList(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/media", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("get media with pagination", func(t *testing.T) {
		mediumOne := &model.Medium{
			Name:           "Bar chart",
			OrganisationID: 1,
		}
		mediumTwo := &model.Medium{
			Name:           "Pie chart",
			OrganisationID: 1,
		}
		media := []model.Medium{}
		total := 0

		config.DB.Model(&model.Medium{}).Create(&mediumOne)
		config.DB.Model(&model.Medium{}).Create(&mediumTwo)
		config.DB.Model(&model.Medium{}).Where(&model.Medium{
			OrganisationID: 1,
		}).Count(&total).Order("id desc").Offset(1).Limit(1).Find(&media)

		resObj := e.GET("/media").
			WithQueryObject(map[string]string{
				"limit": "1",
				"page":  "2",
			}).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("nodes").Array().Element(0).Object().Value("name").String().Equal(media[0].Name)
		resObj.Value("total").Number().Equal(total)

	})
}
