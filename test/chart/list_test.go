package chart

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
	"github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/h2non/gock.v1"
)

func TestChartList(t *testing.T) {
	mock := test.SetupMockDB()
	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	chartlist := []map[string]interface{}{
		{
			"title": "Chart Test 1",
			"slug":  "chart-test-1",
			"mode":  "vega",
			"description": postgres.Jsonb{
				RawMessage: []byte(`{"time":1617039625490,"blocks":[{"type":"paragraph","data":{"text":"Test Description 1"}}],"version":"2.19.0"}`),
			},
			"html_description": "<p>Test Description 1</p>",
			"data_url":         "http://data.com/crime?page[number]=3&page[size]=1",
			"config": `{
				"links": {
					"self": "http://example.com/charts?page[number]=3&page[size]=1",
					"first": "http://example.com/charts?page[number]=1&page[size]=1",
					"prev": "http://example.com/charts?page[number]=2&page[size]=1",
					"next": "http://example.com/charts?page[number]=4&page[size]=1",
					"last": "http://example.com/charts?page[number]=13&page[size]=1"
				  }
			}`,
			"status":         "available",
			"published_date": time.Time{},
			"template_id":    "testtemplate1",
		},
		{
			"title": "Chart Test 2",
			"slug":  "chart-test-2",
			"description": postgres.Jsonb{
				RawMessage: []byte(`{"time":1617039625490,"blocks":[{"type":"paragraph","data":{"text":"Test Description 2"}}],"version":"2.19.0"}`),
			},
			"html_description": "<p>Test Description 2</p>",
			"data_url":         "http://data.com/crime?page[number]=3&page[size]=1",
			"config": `{
				"links": {
					"self": "http://example.com/pie?page[number]=3&page[size]=1",
				  }
			}`,
			"status":             "available",
			"featured_medium_id": uint(1),
			"theme_id":           uint(1),
			"published_date":     time.Time{},
			"template_id":        "testtemplate2",
		},
	}
	byteConfigDataOne, _ := json.Marshal(chartlist[0]["config"])
	byteConfigDataTwo, _ := json.Marshal(chartlist[1]["config"])

	byteDescriptionDataOne, _ := json.Marshal(chartlist[0]["description"])
	byteDescriptionDataTwo, _ := json.Marshal(chartlist[1]["description"])

	t.Run("get empty list of chart", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("0"))

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

	t.Run("get non-empty list of chart", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(chartlist)))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, 1, chartlist[0]["title"], chartlist[0]["slug"], byteDescriptionDataOne, chartlist[0]["html_description"], chartlist[0]["data_url"], byteConfigDataOne, chartlist[0]["status"], chartlist[0]["featured_medium_id"], chartlist[0]["template_id"], chartlist[0]["theme_id"], time.Time{}, chartlist[0]["mode"], 1).
				AddRow(2, time.Now(), time.Now(), nil, 1, 1, chartlist[1]["title"], chartlist[1]["slug"], byteDescriptionDataTwo, chartlist[1]["html_description"], chartlist[1]["data_url"], byteConfigDataTwo, chartlist[1]["status"], chartlist[1]["featured_medium_id"], chartlist[1]["template_id"], chartlist[1]["theme_id"], time.Time{}, chartlist[1]["mode"], 1))

		chartPreloadMock(mock)

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": len(chartlist)}).
			Value("nodes").
			Array().
			Element(0).
			Object().
			ContainsMap(chartlist[0])

		test.ExpectationsMet(t, mock)
	})

	t.Run("get chart with pagination", func(t *testing.T) {
		test.CheckSpace(mock)
		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(chartlist)))

		mock.ExpectQuery(paginationQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(2, time.Now(), time.Now(), nil, 1, 1, chartlist[1]["title"], chartlist[1]["slug"], byteDescriptionDataTwo, chartlist[1]["html_description"], chartlist[1]["data_url"], byteConfigDataTwo, chartlist[1]["status"], chartlist[1]["featured_medium_id"], chartlist[1]["template_id"], chartlist[1]["theme_id"], time.Time{}, chartlist[1]["mode"], 1))

		chartPreloadMock(mock)

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
			ContainsMap(map[string]interface{}{"total": len(chartlist)}).
			Value("nodes").
			Array().
			Element(0).
			Object().
			ContainsMap(chartlist[1])

		test.ExpectationsMet(t, mock)

	})
}
