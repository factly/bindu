package template

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
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

	t.Run("invalid template id", func(t *testing.T) {
		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("template_id", "invalid_id").
			WithHeaders(headers).
			WithJSON(data).
			Expect().
			Status(http.StatusBadRequest)
		test.ExpectationsMet(t, mock)
	})

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
			WithArgs(1, 1).
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

		SelectMock(mock, 1, 1)
		mock.ExpectBegin()

		mock.ExpectExec(`UPDATE \"bi_template\"`).
			WithArgs(test.AnyTime{}, 1, data["title"], data["slug"], data["schema"], data["properties"], data["medium_id"], 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		SelectMock(mock)
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

		SelectMock(mock, 1, 1)
		mock.ExpectBegin()

		mock.ExpectExec(`UPDATE \"bi_template\"`).
			WithArgs(nil, test.AnyTime{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectExec(`UPDATE \"bi_template\"`).
			WithArgs(test.AnyTime{}, 1, data["title"], data["slug"], data["schema"], data["properties"], 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		SelectMock(mock)
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
