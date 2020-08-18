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

func TestTagDetails(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/tags", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid tag id", func(t *testing.T) {
		e.GET("/tags/{tag_id}").
			WithPath("tag_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("tag record not found", func(t *testing.T) {
		e.GET("/tags/{tag_id}").
			WithPath("tag_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("get tag by id", func(t *testing.T) {
		tag := &model.Tag{
			Name:           "Crime",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Tag{}).Create(&tag)

		resObj := e.GET("/tags/{tag_id}").
			WithPath("tag_id", tag.Base.ID).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("name").String().Equal("Crime")
	})

}
