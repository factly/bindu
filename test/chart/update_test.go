package chart

import (
	"encoding/json"
	"errors"
	"fmt"
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

var updateData = map[string]interface{}{
	"title": "Pie",
	"description": `{
		"data": [
			{
			"type": "sport",
			"id": "3",
			"attributes": {
				"title": "JSON:API paints my bikeshed!",
				"body": "The shortest article. Ever.",
				"created": "2015-05-22T14:56:29.000Z",
				"updated": "2015-05-22T14:56:28.000Z"
			}
			}
		]
		}`,
	"data_url": "http://data.com/sports?page[number]=3&page[size]=1",
	"config": `{
		"links": {
			"self": "http://example.com/sport?page[number]=3&page[size]=1",
			"first": "http://example.com/sport?page[number]=1&page[size]=1",
			"prev": "http://example.com/sport?page[number]=2&page[size]=1",
			"next": "http://example.com/sport?page[number]=4&page[size]=1",
			"last": "http://example.com/sport?page[number]=13&page[size]=1"
		  }
	}`,
	"status":             "unavailable",
	"featured_medium_id": uint(1),
	"theme_id":           uint(1),
	"published_date":     time.Time{},
	"category_ids":       []int{1},
	"tag_ids":            []int{1},
}

func TestChartUpdate(t *testing.T) {
	mock := test.SetupMockDB()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)
	res := map[string]interface{}{
		"title": "Pie",
		"description": `{
			"data": [
				{
				"type": "sport",
				"id": "3",
				"attributes": {
					"title": "JSON:API paints my bikeshed!",
					"body": "The shortest article. Ever.",
					"created": "2015-05-22T14:56:29.000Z",
					"updated": "2015-05-22T14:56:28.000Z"
				}
				}
			]
			}`,
		"data_url": "http://data.com/sports?page[number]=3&page[size]=1",
		"config": `{
			"links": {
				"self": "http://example.com/sport?page[number]=3&page[size]=1",
				"first": "http://example.com/sport?page[number]=1&page[size]=1",
				"prev": "http://example.com/sport?page[number]=2&page[size]=1",
				"next": "http://example.com/sport?page[number]=4&page[size]=1",
				"last": "http://example.com/sport?page[number]=13&page[size]=1"
			  }
		}`,
		"status":             "unavailable",
		"featured_medium_id": uint(1),
		"theme_id":           uint(1),
		"published_date":     time.Time{},
	}

	t.Run("invalid chart id", func(t *testing.T) {
		e.PUT(path).
			WithPath("chart_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusBadRequest)
	})
	t.Run("cannot decode chart", func(t *testing.T) {

		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("Unprocessable chart", func(t *testing.T) {

		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(invalidData).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("chart record not found", func(t *testing.T) {
		recordNotFoundMock(mock)

		e.PUT(path).
			WithPath("chart_id", "100").
			WithJSON(data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("update chart", func(t *testing.T) {
		updateChart := updateData
		updateChart["slug"] = "pie"

		chartSelectMock(mock)

		mock.ExpectBegin()
		chartTagUpdate(mock, nil)
		chartCategoryUpdate(mock, nil)

		mediumQueryMock(mock)
		themeQueryMock(mock)
		chartUpdateMock(mock, updateChart)

		res["slug"] = "pie"
		selectAfterUpdate(mock, res)

		mock.ExpectCommit()

		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(updateChart).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)
		test.ExpectationsMet(t, mock)
	})
	t.Run("update chart with different slug", func(t *testing.T) {
		updateChart := updateData
		updateChart["slug"] = "pie-test"

		chartSelectMock(mock)

		mock.ExpectQuery(`SELECT slug, organisation_id FROM "bi_chart"`).
			WithArgs(fmt.Sprint(updateChart["slug"], "%"), 1).
			WillReturnRows(sqlmock.NewRows([]string{"slug", "organisation_id"}))

		mock.ExpectBegin()
		chartTagUpdate(mock, nil)
		chartCategoryUpdate(mock, nil)

		mediumQueryMock(mock)
		themeQueryMock(mock)
		chartUpdateMock(mock, updateChart)

		res["slug"] = "pie-test"
		selectAfterUpdate(mock, res)
		mock.ExpectCommit()

		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(updateChart).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)

	})

	t.Run("update chart by id with empty slug", func(t *testing.T) {

		updateChart := updateData
		updateChart["slug"] = "pie"
		chartSelectMock(mock)

		slugCheckMock(mock)

		mock.ExpectBegin()
		chartTagUpdate(mock, nil)
		chartCategoryUpdate(mock, nil)

		mediumQueryMock(mock)
		themeQueryMock(mock)
		chartUpdateMock(mock, updateChart)

		res["slug"] = "pie"
		selectAfterUpdate(mock, res)
		mock.ExpectCommit()

		updateChart["slug"] = ""

		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(updateChart).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)
		updateChart["slug"] = "pie"
		test.ExpectationsMet(t, mock)
	})

	t.Run("update chart when medium id = 0", func(t *testing.T) {
		updateChart := updateData
		description, _ := json.Marshal(updateChart["description"])
		config, _ := json.Marshal(updateChart["config"])
		updateChart["slug"] = "pie"

		chartSelectMock(mock)

		mock.ExpectBegin()
		chartTagUpdate(mock, nil)
		chartCategoryUpdate(mock, nil)

		themeQueryMock(mock)
		mock.ExpectExec(`UPDATE \"bi_chart\"`).
			WithArgs(nil, test.AnyTime{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		themeQueryMock(mock)
		mock.ExpectExec(`UPDATE \"bi_chart\"`).
			WithArgs(test.AnyTime{}, 1, updateChart["title"], updateChart["slug"], description, updateChart["data_url"], config, updateChart["status"], updateChart["theme_id"], 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		res["slug"] = "pie"
		selectAfterUpdate(mock, res)

		mock.ExpectCommit()

		updateChart["featured_medium_id"] = 0
		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(updateChart).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)
		updateChart["featured_medium_id"] = 1
		test.ExpectationsMet(t, mock)
	})

	t.Run("update chart when theme id = 0", func(t *testing.T) {
		updateChart := updateData
		description, _ := json.Marshal(updateChart["description"])
		config, _ := json.Marshal(updateChart["config"])
		updateChart["slug"] = "pie"

		chartSelectMock(mock)

		mock.ExpectBegin()
		chartTagUpdate(mock, nil)
		chartCategoryUpdate(mock, nil)

		mediumQueryMock(mock)
		mock.ExpectExec(`UPDATE \"bi_chart\"`).
			WithArgs(nil, test.AnyTime{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mediumQueryMock(mock)
		mock.ExpectExec(`UPDATE \"bi_chart\"`).
			WithArgs(test.AnyTime{}, 1, updateChart["title"], updateChart["slug"], description, updateChart["data_url"], config, updateChart["status"], updateChart["featured_medium_id"], 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		res["slug"] = "pie"
		selectAfterUpdate(mock, res)

		mock.ExpectCommit()

		updateChart["theme_id"] = 0
		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(updateChart).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)
		updateChart["theme_id"] = 1
		test.ExpectationsMet(t, mock)
	})

	t.Run("updating chart tags fail", func(t *testing.T) {
		updateChart := updateData
		updateChart["slug"] = "pie"

		chartSelectMock(mock)

		mock.ExpectBegin()
		chartTagUpdate(mock, errors.New("cannot update chart tags"))

		mock.ExpectRollback()

		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(updateChart).
			Expect().
			Status(http.StatusInternalServerError)
		test.ExpectationsMet(t, mock)
	})

	t.Run("updating chart categories fail", func(t *testing.T) {
		updateChart := updateData
		updateChart["slug"] = "pie"

		chartSelectMock(mock)

		mock.ExpectBegin()
		chartTagUpdate(mock, nil)
		chartCategoryUpdate(mock, errors.New("cannot update chart categories"))

		mock.ExpectRollback()

		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(updateChart).
			Expect().
			Status(http.StatusInternalServerError)
		test.ExpectationsMet(t, mock)
	})
}
