package category

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

func TestCategoryList(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/categories", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("get categories with pagination", func(t *testing.T) {
		categoryOne := &model.Category{
			Name:           "Sample-1",
			OrganisationID: 1,
		}
		categoryTwo := &model.Category{
			Name:           "Sample-2",
			OrganisationID: 1,
		}
		categories := []model.Category{}
		total := 0

		config.DB.Model(&model.Category{}).Create(&categoryOne)
		config.DB.Model(&model.Category{}).Create(&categoryTwo)
		config.DB.Model(&model.Category{}).Where(&model.Category{
			OrganisationID: 1,
		}).Count(&total).Order("id desc").Offset(1).Limit(1).Find(&categories)

		resObj := e.GET("/categories").
			WithQueryObject(map[string]string{
				"limit": "1",
				"page":  "2",
			}).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("nodes").Array().Element(0).Object().Value("name").String().Equal(categories[0].Name)
		resObj.Value("total").Number().Equal(total)

	})

}
