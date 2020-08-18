package tag

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

func TestTagCreate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/tags", Router())

	tagOne := model.Tag{
		Name: "Elections",
		Slug: "elections",
	}

	// tag with slug empty
	tagTwo := model.Tag{
		Name: "Football",
	}

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("Unprocessable tag", func(t *testing.T) {
		e.POST("/tags").
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("create tag", func(t *testing.T) {
		resObj := e.POST("/tags").
			WithHeaders(headers).
			WithJSON(tagOne).
			Expect().
			Status(http.StatusCreated).JSON().Object()

		resObj.Value("name").String().Equal("Elections")
		resObj.Value("slug").String().Equal("elections")

	})

	t.Run("create tag with slug is empty", func(t *testing.T) {

		resObj := e.POST("/tags").
			WithHeaders(headers).
			WithJSON(tagTwo).
			Expect().
			Status(http.StatusCreated).JSON().Object()

		resObj.Value("name").String().Equal("Football")
		resObj.Value("slug").String().Equal("football")
	})

}
