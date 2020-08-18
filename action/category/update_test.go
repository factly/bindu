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

func TestCategoryUpdate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/categories", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	category := &model.Category{
		Name:           "Agri",
		Slug:           "agriculture",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Category{}).Create(&category)

	t.Run("invalid category id", func(t *testing.T) {
		e.PUT("/categories/{category_id}").
			WithPath("category_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("category record not found", func(t *testing.T) {

		e.PUT("/categories/{category_id}").
			WithPath("category_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("update category", func(t *testing.T) {

		body := model.Category{
			Name: "Agriculture",
			Slug: "agriculture",
		}

		resObj := e.PUT("/categories/{category_id}").
			WithPath("category_id", category.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("name").String().Equal("Agriculture")

	})

	t.Run("update category by id with empty slug", func(t *testing.T) {

		body := model.Category{
			Name: "Crop",
			Slug: "",
		}

		resObj := e.PUT("/categories/{category_id}").
			WithPath("category_id", category.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("name").String().Equal("Crop")
		resObj.Value("slug").String().Equal("crop")

	})

	t.Run("update category with different slug", func(t *testing.T) {

		body := model.Category{
			Name: "Crop test",
			Slug: "crop-test",
		}

		resObj := e.PUT("/categories/{category_id}").
			WithPath("category_id", category.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("name").String().Equal("Crop test")
		resObj.Value("slug").String().Equal("crop-test")

	})

}
