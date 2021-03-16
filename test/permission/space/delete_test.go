package space

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

func TestSpacePermissionDelete(t *testing.T) {
	mock := test.SetupMockDB()
	test.SetEnv()

	test.MockServers()
	defer gock.DisableNetworking()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid permission id", func(t *testing.T) {
		test.CheckSpace(mock)

		e.DELETE(path).
			WithPath("permission_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusBadRequest)
		test.ExpectationsMet(t, mock)
	})

	t.Run("space permission record does not exist", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(selectQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows(columns))

		e.DELETE(path).
			WithPath("permission_id", "1").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
		test.ExpectationsMet(t, mock)
	})

	t.Run("delete space permission", func(t *testing.T) {
		test.CheckSpace(mock)

		SelectQuery(mock, 1)

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE \"bi_space_permission\"`).
			WithArgs(test.AnyTime{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e.DELETE(path).
			WithPath("permission_id", "1").
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)
		test.ExpectationsMet(t, mock)
	})

}
