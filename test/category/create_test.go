package category

import (
	"github.com/factly/bindu-server/action"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

func TestCategoryCreate(t *testing.T) {

	mock := test.SetupMockDB()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("cannot decode category", func(t *testing.T) {

		e.POST(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("Unprocessable category", func(t *testing.T) {

		e.POST(basePath).
			WithHeaders(headers).
			WithJSON(invalidData).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("create category", func(t *testing.T) {

		slugCheckMock(mock)

		categoryInsertMock(mock)

		e.POST(basePath).
			WithHeaders(headers).
			WithJSON(data).
			Expect().
			Status(http.StatusCreated).JSON().Object().ContainsMap(data)
		test.ExpectationsMet(t, mock)

	})

	t.Run("create category with slug is empty", func(t *testing.T) {

		slugCheckMock(mock)

		categoryInsertMock(mock)

		e.POST(basePath).
			WithHeaders(headers).
			WithJSON(dataWithoutSlug).
			Expect().
			Status(http.StatusCreated).JSON().Object().ContainsMap(data)

		test.ExpectationsMet(t, mock)
	})

}
