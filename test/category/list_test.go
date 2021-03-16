package category

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

func TestCategoryList(t *testing.T) {
	mock := test.SetupMockDB()

	test.MockServers()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	categorylist := []map[string]interface{}{
		{"name": "Test Category 1", "slug": "test-category-1"},
		{"name": "Test Category 2", "slug": "test-category-2"},
	}

	t.Run("get empty list of categories", func(t *testing.T) {
		test.CheckSpace(mock)
		categoryCountQuery(mock, 0)

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns))

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": 0})

		test.ExpectationsMet(t, mock)
	})

	t.Run("get non-empty list of categories", func(t *testing.T) {

		test.CheckSpace(mock)
		categoryCountQuery(mock, len(categorylist))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, 1, categorylist[0]["name"], categorylist[0]["slug"], 1).
				AddRow(2, time.Now(), time.Now(), nil, 1, 1, categorylist[1]["name"], categorylist[1]["slug"], 1))

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": len(categorylist)}).
			Value("nodes").
			Array().
			Element(0).
			Object().
			ContainsMap(categorylist[0])

		test.ExpectationsMet(t, mock)
	})

	t.Run("get categories with pagination", func(t *testing.T) {
		test.CheckSpace(mock)
		categoryCountQuery(mock, len(categorylist))

		mock.ExpectQuery(paginationQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(2, time.Now(), time.Now(), nil, 1, 1, categorylist[1]["name"], categorylist[1]["slug"], 1))

		e.GET(basePath).
			WithQueryObject(map[string]interface{}{
				"limit": "1",
				"page":  "2",
			}).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": len(categorylist)}).
			Value("nodes").
			Array().
			Element(0).
			Object().
			ContainsMap(categorylist[1])

		test.ExpectationsMet(t, mock)

	})
}
