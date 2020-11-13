package chart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
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

func deleteAssociationsMock(mock sqlmock.Sqlmock) {
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "bi_chart_tag"`)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "bi_chart_category"`)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
}

func chartTagUpdate(mock sqlmock.Sqlmock) {
	chartSelectMock(mock)

	chartTagMock(mock)
	chartCategoryMock(mock)

	tagQueryMock(mock)
	categoryQueryMock(mock)
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
			Status(http.StatusNotFound)
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

		chartTagUpdate(mock)

		deleteAssociationsMock(mock)

		chartUpdateMock(mock, updateChart)
		res["slug"] = "pie"
		selectAfterUpdate(mock, res)

		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(updateChart).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)

	})
	t.Run("update chart with different slug", func(t *testing.T) {
		updateChart := updateData
		updateChart["slug"] = "pie-test"

		chartTagUpdate(mock)

		mock.ExpectQuery(`SELECT slug, organisation_id FROM "bi_chart"`).
			WithArgs(fmt.Sprint(updateChart["slug"], "%"), 1).
			WillReturnRows(sqlmock.NewRows([]string{"slug", "organisation_id"}))

		deleteAssociationsMock(mock)

		chartUpdateMock(mock, updateChart)
		res["slug"] = "pie-test"
		selectAfterUpdate(mock, res)

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
		updateChart["category_ids"] = []uint{}
		updateChart["tag_ids"] = []uint{}

		chartSelectMock(mock)

		chartTagMock(mock)
		chartCategoryMock(mock)

		mock.ExpectQuery(tagQuery).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}))
		mock.ExpectQuery(categoryQuery).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}))

		slugCheckMock(mock)

		deleteAssociationsMock(mock)

		description, _ := json.Marshal(updateChart["description"])
		config, _ := json.Marshal(updateChart["config"])

		mock.ExpectBegin()

		mock.ExpectQuery(mediumQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug", "type", "url"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, medium["name"], medium["slug"], medium["type"], byteMediumData))

		mock.ExpectQuery(themeQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "config"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, theme["name"], byteThemeData))

		mock.ExpectExec(`UPDATE \"bi_chart\" SET (.+)  WHERE (.+) \"bi_chart\".\"id\" = `).
			WithArgs(config, updateChart["data_url"],
				description, updateChart["featured_medium_id"], updateChart["slug"],
				updateChart["status"], updateChart["theme_id"], updateChart["title"], test.AnyTime{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		res["slug"] = "pie"
		selectAfterUpdate(mock, res)

		updateChart["slug"] = ""

		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(updateChart).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)
		test.ExpectationsMet(t, mock)
	})

}
