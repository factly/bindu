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

func TestTagUpdate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/tags", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	tag := &model.Tag{
		Name:           "Agri",
		Slug:           "agriculture",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Tag{}).Create(&tag)

	t.Run("invalid tag id", func(t *testing.T) {
		e.PUT("/tags/{tag_id}").
			WithPath("tag_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("tag record not found", func(t *testing.T) {

		e.PUT("/tags/{tag_id}").
			WithPath("tag_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("update tag", func(t *testing.T) {

		body := model.Tag{
			Name: "Agriculture",
			Slug: "agriculture",
		}

		resObj := e.PUT("/tags/{tag_id}").
			WithPath("tag_id", tag.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("name").String().Equal("Agriculture")

	})

	t.Run("update tag by id with empty slug", func(t *testing.T) {

		body := model.Tag{
			Name: "Crop",
			Slug: "",
		}

		resObj := e.PUT("/tags/{tag_id}").
			WithPath("tag_id", tag.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("name").String().Equal("Crop")
		resObj.Value("slug").String().Equal("crop")

	})

	t.Run("update tag with different slug", func(t *testing.T) {

		body := model.Tag{
			Name: "Crop test",
			Slug: "crop-test",
		}

		resObj := e.PUT("/tags/{tag_id}").
			WithPath("tag_id", tag.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("name").String().Equal("Crop test")
		resObj.Value("slug").String().Equal("crop-test")

	})

}
