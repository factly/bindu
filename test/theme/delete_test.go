package theme

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

func TestThemeDelete(t *testing.T) {
	mock := test.SetupMockDB()
	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid theme id", func(t *testing.T) {
		test.CheckSpace(mock)

		e.DELETE(path).
			WithPath("theme_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusBadRequest)

	})

	t.Run("theme record not found", func(t *testing.T) {
		test.CheckSpace(mock)

		recordNotFoundMock(mock)

		e.DELETE(path).
			WithPath("theme_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("check theme associated with other entity", func(t *testing.T) {
		test.CheckSpace(mock)

		SelectMock(mock)

		themeChartExpect(mock, 1)

		e.DELETE(path).
			WithPath("theme_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)
	})

	t.Run("theme record deleted", func(t *testing.T) {
		test.CheckSpace(mock)

		SelectMock(mock)

		themeChartExpect(mock, 0)

		mock.ExpectBegin()
		mock.ExpectExec(deleteQuery).
			WithArgs(test.AnyTime{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e.DELETE(path).
			WithPath("theme_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)
	})

}
