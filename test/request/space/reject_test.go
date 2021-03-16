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

func TestSpaceRequestReject(t *testing.T) {
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

	t.Run("invalid request id", func(t *testing.T) {
		test.CheckSpace(mock)
		e.POST(rejectPath).
			WithPath("request_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusBadRequest)
		test.ExpectationsMet(t, mock)
	})

	t.Run("request record not found", func(t *testing.T) {
		test.CheckSpace(mock)
		mock.ExpectQuery(selectQuery).
			WithArgs("pending", 1).
			WillReturnRows(sqlmock.NewRows(Columns))

		e.POST(rejectPath).
			WithPath("request_id", "1").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
		test.ExpectationsMet(t, mock)
	})

	t.Run("request rejected", func(t *testing.T) {
		test.CheckSpace(mock)
		SelectQuery(mock)

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE \"bi_space_permission_request\"`).
			WithArgs(test.AnyTime{}, 1, "rejected", 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		e.POST(rejectPath).
			WithPath("request_id", "1").
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)
		test.ExpectationsMet(t, mock)
	})
}
