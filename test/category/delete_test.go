package category

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

func TestCategoryDelete(t *testing.T) {
	mock := test.SetupMockDB()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid category id", func(t *testing.T) {
		e.DELETE(path).
			WithPath("category_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)

	})

	t.Run("category record not found", func(t *testing.T) {
		recordNotFoundMock(mock)

		e.DELETE(path).
			WithPath("category_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("check category associated with other entity", func(t *testing.T) {
		categorySelectMock(mock)

		categoryChartExpect(mock, 1)

		e.DELETE(path).
			WithPath("category_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)
	})

	t.Run("category record deleted", func(t *testing.T) {
		categorySelectMock(mock)

		categoryChartExpect(mock, 0)

		mock.ExpectBegin()
		mock.ExpectExec(deleteQuery).
			WithArgs(test.AnyTime{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e.DELETE(path).
			WithPath("category_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)
	})

}
