package category

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-chi/chi"
	"gopkg.in/h2non/gock.v1"
)

func TestCategoryList(t *testing.T) {
	mock := test.SetupMockDB()
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/categories", Router())

	testServer := httptest.NewServer(r)
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

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("0"))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(categoryProps))

		e.GET("/categories").
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": 0})

		mock.ExpectationsWereMet()
	})

	t.Run("get non-empty list of categories", func(t *testing.T) {

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(categorylist)))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(categoryProps).
				AddRow(1, time.Now(), time.Now(), nil, categorylist[0]["name"], categorylist[0]["slug"]).
				AddRow(2, time.Now(), time.Now(), nil, categorylist[1]["name"], categorylist[1]["slug"]))

		e.GET("/categories").
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

		mock.ExpectationsWereMet()
	})

	t.Run("get categories with pagination", func(t *testing.T) {
		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(categorylist)))

		mock.ExpectQuery(paginationQuery).
			WillReturnRows(sqlmock.NewRows(categoryProps).
				AddRow(2, time.Now(), time.Now(), nil, categorylist[1]["name"], categorylist[1]["slug"]))

		e.GET("/categories").
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

		mock.ExpectationsWereMet()

	})
}
