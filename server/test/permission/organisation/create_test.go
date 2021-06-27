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

func TestOrganisationPermissionCreate(t *testing.T) {

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

	t.Run("Unprocessable permission", func(t *testing.T) {
		test.CheckSpace(mock)

		e.POST(basePath).
			WithJSON(invalidData).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

		test.ExpectationsMet(t, mock)
	})

	t.Run("Undecodable permission", func(t *testing.T) {
		test.CheckSpace(mock)

		e.POST(basePath).
			WithHeaders(headers).
			WithJSON(undecodableData).
			Expect().
			Status(http.StatusUnprocessableEntity)

		test.ExpectationsMet(t, mock)
	})

	t.Run("Organisation's permission already exist", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(countQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).
				AddRow(1))

		e.POST(basePath).
			WithJSON(Data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)
		test.ExpectationsMet(t, mock)
	})

	t.Run("Create organisation permission", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(countQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).
				AddRow(0))

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "bi_organisation_permission"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, Data["organisation_id"], Data["spaces"]).
			WillReturnRows(sqlmock.
				NewRows([]string{"id"}).
				AddRow(1))
		mock.ExpectCommit()

		e.POST(basePath).
			WithJSON(Data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusCreated)
		test.ExpectationsMet(t, mock)
	})
}
