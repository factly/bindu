package theme

import (
	"encoding/json"
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

func TestThemeList(t *testing.T) {
	mock := test.SetupMockDB()
	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	themelist := []map[string]interface{}{
		{"name": "Test Theme 1", "config": `{"image": { 
			"src": "Images/Sun.png",
			"name": "sun1",
			"hOffset": 250,
			"vOffset": 250,
			"alignment": "center"
		}}`},
		{"name": "Test Theme 2", "config": `{"image": { 
			"src": "Images/Sun.png",
			"name": "sun2",
			"hOffset": 250,
			"vOffset": 250,
			"alignment": "center"
		}}`},
	}

	byteData0, _ := json.Marshal(themelist[0]["config"])
	byteData1, _ := json.Marshal(themelist[1]["config"])

	t.Run("get empty list of themes", func(t *testing.T) {
		test.CheckSpace(mock)
		themeCountQuery(mock, 0)

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

	t.Run("get non-empty list of themes", func(t *testing.T) {
		test.CheckSpace(mock)
		themeCountQuery(mock, len(themelist))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, 1, 1, themelist[0]["name"], byteData0, 1).
				AddRow(2, time.Now(), time.Now(), nil, 1, 1, 1, themelist[1]["name"], byteData1, 1))

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": len(themelist)}).
			Value("nodes").
			Array().
			Element(0).
			Object().
			ContainsMap(themelist[0])

		test.ExpectationsMet(t, mock)
	})

	t.Run("get themes with pagination", func(t *testing.T) {
		test.CheckSpace(mock)
		themeCountQuery(mock, len(themelist))

		mock.ExpectQuery(paginationQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(2, time.Now(), time.Now(), nil, 1, 1, 1, themelist[1]["name"], byteData1, 1))

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
			ContainsMap(map[string]interface{}{"total": len(themelist)}).
			Value("nodes").
			Array().
			Element(0).
			Object().
			ContainsMap(themelist[1])

		test.ExpectationsMet(t, mock)

	})
}
