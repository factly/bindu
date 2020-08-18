package tag

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

func TestTagList(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/tags", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("get tags with pagination", func(t *testing.T) {
		tagOne := &model.Tag{
			Name:           "AP",
			OrganisationID: 1,
		}
		tagTwo := &model.Tag{
			Name:           "Telangana",
			OrganisationID: 1,
		}
		tags := []model.Tag{}
		total := 0

		config.DB.Model(&model.Tag{}).Create(&tagOne)
		config.DB.Model(&model.Tag{}).Create(&tagTwo)
		config.DB.Model(&model.Tag{}).Where(&model.Tag{
			OrganisationID: 1,
		}).Count(&total).Order("id desc").Offset(1).Limit(1).Find(&tags)

		resObj := e.GET("/tags").
			WithQueryObject(map[string]string{
				"limit": "1",
				"page":  "2",
			}).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("nodes").Array().Element(0).Object().Value("name").String().Equal(tags[0].Name)
		resObj.Value("total").Number().Equal(total)

	})
}
