package organisation

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/test/permission/space"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect"
	"gopkg.in/h2non/gock.v1"
)

func TestOrganisationPermissionDetails(t *testing.T) {
	mock := test.SetupMockDB()
	test.SetEnv()

	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("permission record not found", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns))

		e.GET(mypath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
		test.ExpectationsMet(t, mock)
	})

	t.Run("get my permission", func(t *testing.T) {
		test.CheckSpace(mock)

		SelectQuery(mock, 1)

		spaceSelectQuery(mock)
		space.SelectQuery(mock)

		e.GET(mypath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)
		test.ExpectationsMet(t, mock)
	})
}
