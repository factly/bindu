package template

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect"
	"gopkg.in/h2non/gock.v1"
)

func TestTemplateDelete(t *testing.T) {
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

		e.DELETE(path).
			WithPath("template_id", "invalid").
			WithHeaders(headers).
			Expect().
			Status(http.StatusBadRequest)
		test.ExpectationsMet(t, mock)
	})

	t.Run("template record not found", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns))

		e.DELETE(path).
			WithPath("template_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
		test.ExpectationsMet(t, mock)
	})

	t.Run("delete template", func(t *testing.T) {
		test.CheckSpace(mock)

		SelectMock(mock)
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "bi_template" SET`)).
			WithArgs(test.AnyTime{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e.DELETE(path).
			WithPath("template_id", "1").
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)
		test.ExpectationsMet(t, mock)
	})

}
