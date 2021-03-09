package medium

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

func TestMediumList(t *testing.T) {
	mock := test.SetupMockDB()
	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	mediumlist := []map[string]interface{}{
		{"name": "Sample Medium 1", "slug": "test-medium-1"},
		{"name": "Sample Medium 2", "slug": "test-medium-2"},
	}

	t.Run("get empty list of media", func(t *testing.T) {
		test.CheckSpace(mock)

		mediumCountQuery(mock, 0)

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

	t.Run("get non-empty list of media", func(t *testing.T) {
		test.CheckSpace(mock)
		mediumCountQuery(mock, len(mediumlist))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, 1, mediumlist[0]["name"], mediumlist[0]["slug"], mediumlist[0]["type"], mediumlist[0]["title"], mediumlist[0]["description"], mediumlist[0]["caption"], mediumlist[0]["alt_text"], mediumlist[0]["file_size"], mediumlist[0]["url"], mediumlist[0]["dimensions"], 1).
				AddRow(2, time.Now(), time.Now(), nil, 1, 1, mediumlist[1]["name"], mediumlist[1]["slug"], mediumlist[1]["type"], mediumlist[1]["title"], mediumlist[1]["description"], mediumlist[1]["caption"], mediumlist[1]["alt_text"], mediumlist[1]["file_size"], mediumlist[1]["url"], mediumlist[1]["dimensions"], 1))

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": len(mediumlist)}).
			Value("nodes").
			Array().
			Element(0).
			Object().
			ContainsMap(mediumlist[0])

		test.ExpectationsMet(t, mock)
	})

	t.Run("get media with pagination", func(t *testing.T) {
		test.CheckSpace(mock)
		mediumCountQuery(mock, len(mediumlist))

		mock.ExpectQuery(paginationQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(2, time.Now(), time.Now(), nil, 1, 1, mediumlist[1]["name"], mediumlist[1]["slug"], mediumlist[1]["type"], mediumlist[1]["title"], mediumlist[1]["description"], mediumlist[1]["caption"], mediumlist[1]["alt_text"], mediumlist[1]["file_size"], mediumlist[1]["url"], mediumlist[1]["dimensions"], 1))

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
			ContainsMap(map[string]interface{}{"total": len(mediumlist)}).
			Value("nodes").
			Array().
			Element(0).
			Object().
			ContainsMap(mediumlist[1])

		test.ExpectationsMet(t, mock)

	})
}
