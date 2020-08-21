package tag

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

func TestTagList(t *testing.T) {
	mock := test.SetupMockDB()
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount(url, Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	taglist := []map[string]interface{}{
		{"name": "Test Tag 1", "slug": "test-tag-1"},
		{"name": "Test Tag 2", "slug": "test-tag-2"},
	}

	t.Run("get empty list of tags", func(t *testing.T) {

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("0"))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(tagProps))

		e.GET(url).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": 0})

		mock.ExpectationsWereMet()
	})

	t.Run("get non-empty list of tags", func(t *testing.T) {

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(taglist)))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(tagProps).
				AddRow(1, time.Now(), time.Now(), nil, taglist[0]["name"], taglist[0]["slug"]).
				AddRow(2, time.Now(), time.Now(), nil, taglist[1]["name"], taglist[1]["slug"]))

		e.GET(url).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": len(taglist)}).
			Value("nodes").
			Array().
			Element(0).
			Object().
			ContainsMap(taglist[0])

		mock.ExpectationsWereMet()
	})

	t.Run("get tags with pagination", func(t *testing.T) {
		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(taglist)))

		mock.ExpectQuery(paginationQuery).
			WillReturnRows(sqlmock.NewRows(tagProps).
				AddRow(2, time.Now(), time.Now(), nil, taglist[1]["name"], taglist[1]["slug"]))

		e.GET(url).
			WithQueryObject(map[string]interface{}{
				"limit": "1",
				"page":  "2",
			}).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": len(taglist)}).
			Value("nodes").
			Array().
			Element(0).
			Object().
			ContainsMap(taglist[1])

		mock.ExpectationsWereMet()

	})
}
