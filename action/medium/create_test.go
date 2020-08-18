package medium

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

func TestMediumCreate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/media", Router())

	mediumOne := model.Medium{
		Name: "Bar graph",
		Slug: "bar-graph",
	}

	// medium with slug empty
	mediumTwo := model.Medium{
		Name: "Chart",
	}

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("Unprocessable medium", func(t *testing.T) {
		e.POST("/media").
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("create medium", func(t *testing.T) {
		resObj := e.POST("/media").
			WithHeaders(headers).
			WithJSON(mediumOne).
			Expect().
			Status(http.StatusCreated).JSON().Object()

		resObj.Value("name").String().Equal("Bar graph")
		resObj.Value("slug").String().Equal("bar-graph")

	})

	t.Run("create medium with slug is empty", func(t *testing.T) {

		resObj := e.POST("/media").
			WithHeaders(headers).
			WithJSON(mediumTwo).
			Expect().
			Status(http.StatusCreated).JSON().Object()

		resObj.Value("name").String().Equal("Chart")
		resObj.Value("slug").String().Equal("chart")
	})

}
