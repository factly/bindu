package space

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/test/space"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect"
	"gopkg.in/h2non/gock.v1"
)

func TestSpacePermissionList(t *testing.T) {
	mock := test.SetupMockDB()
	test.SetEnv()

	test.MockServers()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("get empty list of permissions", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(1) FROM "bi_space"`)).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_space"`)).
			WillReturnRows(sqlmock.NewRows(space.Columns))

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			Value("total").
			Number().
			Equal(0)

		test.ExpectationsMet(t, mock)
	})

	t.Run("get list of permissions", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(1) FROM "bi_space"`)).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		space.SelectQuery(mock)

		SelectQuery(mock)

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			Value("nodes").
			Array().
			Element(0).
			Object().
			Value("permission").
			Object().
			ContainsMap(Data)

		test.ExpectationsMet(t, mock)
	})
}
