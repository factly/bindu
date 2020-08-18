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

func TestMediumUpdate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/media", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	medium := &model.Medium{
		Name:           "Agri",
		Slug:           "agriculture",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Medium{}).Create(&medium)

	t.Run("invalid medium id", func(t *testing.T) {
		e.PUT("/media/{medium_id}").
			WithPath("medium_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("medium record not found", func(t *testing.T) {

		e.PUT("/media/{medium_id}").
			WithPath("medium_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("update medium", func(t *testing.T) {

		body := model.Medium{
			Name: "Agriculture",
			Slug: "agriculture",
		}

		resObj := e.PUT("/media/{medium_id}").
			WithPath("medium_id", medium.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("name").String().Equal("Agriculture")

	})

	t.Run("update medium by id with empty slug", func(t *testing.T) {

		body := model.Medium{
			Name: "Crop",
			Slug: "",
		}

		resObj := e.PUT("/media/{medium_id}").
			WithPath("medium_id", medium.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("name").String().Equal("Crop")
		resObj.Value("slug").String().Equal("crop")

	})

	t.Run("update medium with different slug", func(t *testing.T) {

		body := model.Medium{
			Name: "Crop test",
			Slug: "crop-test",
		}

		resObj := e.PUT("/media/{medium_id}").
			WithPath("medium_id", medium.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("name").String().Equal("Crop test")
		resObj.Value("slug").String().Equal("crop-test")

	})

}
