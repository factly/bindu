package template

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/test/category"
	"github.com/factly/bindu-server/test/medium"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect"
	"gopkg.in/h2non/gock.v1"
)

func TestTemplateCreate(t *testing.T) {

	mock := test.SetupMockDB()
	test.SetEnv()
	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("cannot decode template", func(t *testing.T) {
		test.CheckSpace(mock)

		e.POST(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)
		test.ExpectationsMet(t, mock)
	})

	t.Run("create a template", func(t *testing.T) {
		test.CheckSpace(mock)

		slugCheckMock(mock)

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "bi_template"`).
			WithArgs(sqlmock.AnyArg(), test.AnyTime{}, test.AnyTime{}, nil, 1, 1, data["title"], data["slug"], data["spec"], data["properties"], data["category_id"], data["is_default"], 1, data["medium_id"]).
			WillReturnRows(sqlmock.
				NewRows([]string{"medium_id"}).
				AddRow(1))

		SelectMock(mock)
		category.SelectMock(mock)
		medium.SelectMock(mock)
		mock.ExpectCommit()

		e.POST(basePath).
			WithHeaders(headers).
			WithJSON(data).
			Expect().
			Status(http.StatusCreated).JSON().Object().ContainsMap(data)
		test.ExpectationsMet(t, mock)
	})

}
