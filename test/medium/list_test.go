package medium

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

func TestMediumList(t *testing.T) {
	mock := test.SetupMockDB()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	mediumlist := []map[string]interface{}{
		{"name": "Test Medium 1",
			"type": "png",
			"slug": "test-medium-1",
			"url": `{"image": { 
			"src": "Images/Sun.png",
			"name": "sun1",
			"hOffset": 250,
			"vOffset": 250,
			"alignment": "center"
		}}`},
		{"name": "Test Medium 2",
			"type": "jpg",
			"slug": "test-medium-2",
			"url": `{"image": { 
			"src": "Images/Sun.png",
			"name": "sun2",
			"hOffset": 250,
			"vOffset": 250,
			"alignment": "center"
		}}`},
	}

	byteData0, _ := json.Marshal(mediumlist[0]["url"])
	byteData1, _ := json.Marshal(mediumlist[1]["url"])

	t.Run("get empty list of media", func(t *testing.T) {

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
		mediumCountQuery(mock, len(mediumlist))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, mediumlist[0]["name"], mediumlist[0]["slug"], mediumlist[0]["type"], byteData0).
				AddRow(2, time.Now(), time.Now(), nil, 1, mediumlist[1]["name"], mediumlist[1]["slug"], mediumlist[1]["type"], byteData1))

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
		mediumCountQuery(mock, len(mediumlist))

		mock.ExpectQuery(paginationQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(2, time.Now(), time.Now(), nil, 1, mediumlist[1]["name"], mediumlist[1]["slug"], mediumlist[1]["type"], byteData1))

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
