package organisation

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

func TestOrganisationPermissionUpdate(t *testing.T) {
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
		e.PUT(path).
			WithPath("permission_id", "invalid_id").
			WithHeaders(headers).
			WithJSON(Data).
			Expect().
			Status(http.StatusBadRequest)
		test.ExpectationsMet(t, mock)
	})

	t.Run("undecodable permission", func(t *testing.T) {
		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("permission_id", "1").
			WithHeaders(headers).
			WithJSON(undecodableData).
			Expect().
			Status(http.StatusUnprocessableEntity)
		test.ExpectationsMet(t, mock)
	})

	t.Run("unprocessable permission", func(t *testing.T) {
		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("permission_id", "1").
			WithHeaders(headers).
			WithJSON(invalidData).
			Expect().
			Status(http.StatusUnprocessableEntity)
		test.ExpectationsMet(t, mock)
	})

	t.Run("permission record does not exist", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns))

		e.PUT(path).
			WithPath("permission_id", "1").
			WithHeaders(headers).
			WithJSON(Data).
			Expect().
			Status(http.StatusNotFound)
		test.ExpectationsMet(t, mock)
	})

	t.Run("update permission", func(t *testing.T) {
		test.CheckSpace(mock)

		SelectQuery(mock)

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE \"bi_organisation_permission\"`).
			WithArgs(test.AnyTime{}, 1, Data["spaces"], 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		SelectQuery(mock, 1, 1)
		mock.ExpectCommit()

		e.PUT(path).
			WithPath("permission_id", "1").
			WithHeaders(headers).
			WithJSON(Data).
			Expect().
			Status(http.StatusOK)
		test.ExpectationsMet(t, mock)
	})
}
