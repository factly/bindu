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

func TestTemplateUpdate(t *testing.T) {
	mock := test.SetupMockDB()
	test.SetEnv()
	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("undecodable template", func(t *testing.T) {
		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("template_id", "1").
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)
		test.ExpectationsMet(t, mock)
	})

	t.Run("template record not found", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns))

		e.PUT(path).
			WithPath("template_id", "1").
			WithJSON(data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
		test.ExpectationsMet(t, mock)
	})

	t.Run("update template", func(t *testing.T) {
		test.CheckSpace(mock)

		SelectMock(mock)
		mock.ExpectBegin()

		mock.ExpectExec(`UPDATE \"bi_template\"`).
			WithArgs(test.AnyTime{}, 1, data["title"], data["slug"], data["spec"], data["properties"], data["category_id"], data["medium_id"], data["mode"], data["description"], data["html_description"], "1").
			WillReturnResult(sqlmock.NewResult(1, 1))

		SelectMock(mock)
		category.SelectMock(mock)
		medium.SelectMock(mock)
		mock.ExpectCommit()

		e.PUT(path).
			WithPath("template_id", "1").
			WithJSON(data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(data)
		test.ExpectationsMet(t, mock)
	})

	t.Run("update template when medium_id = 0", func(t *testing.T) {
		test.CheckSpace(mock)

		SelectMock(mock)
		mock.ExpectBegin()

		mock.ExpectExec(`UPDATE \"bi_template\"`).
			WithArgs(nil, test.AnyTime{}, "1").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectExec(`UPDATE \"bi_template\"`).
			WithArgs(test.AnyTime{}, 1, data["title"], data["slug"], data["spec"], data["properties"], data["category_id"], data["mode"], data["description"], data["html_description"], "1").
			WillReturnResult(sqlmock.NewResult(1, 1))

		SelectMock(mock)
		category.SelectMock(mock)
		medium.SelectMock(mock)
		mock.ExpectCommit()

		data["medium_id"] = 0
		e.PUT(path).
			WithPath("template_id", "1").
			WithJSON(data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object()
		test.ExpectationsMet(t, mock)
		data["medium_id"] = 1
	})

}
