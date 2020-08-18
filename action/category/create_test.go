package category

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-chi/chi"
	"gopkg.in/h2non/gock.v1"
)

func TestCategoryCreate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/categories", Router())

	categoryOne := model.Category{
		Name: "Politics",
		Slug: "politics",
	}

	// category with slug empty
	categoryTwo := model.Category{
		Name: "Cricket",
	}

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("Unprocessable category", func(t *testing.T) {
		e.POST("/categories").
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("create category", func(t *testing.T) {
		resObj := e.POST("/categories").
			WithHeaders(headers).
			WithJSON(categoryOne).
			Expect().
			Status(http.StatusCreated).JSON().Object()

		resObj.Value("name").String().Equal("Politics")
		resObj.Value("slug").String().Equal("politics")

	})

	t.Run("create categories with slug is empty", func(t *testing.T) {

		resObj := e.POST("/categories").
			WithHeaders(headers).
			WithJSON(categoryTwo).
			Expect().
			Status(http.StatusCreated).JSON().Object()

		resObj.Value("name").String().Equal("Cricket")
		resObj.Value("slug").String().Equal("cricket")
	})

}
