package chart

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
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

	testServer := httptest.NewServer(Routes())
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

	t.Run("chart record not found", func(t *testing.T) {
		mock.ExpectQuery(selectQuery).
			WithArgs(100, 1).
			WillReturnRows(sqlmock.NewRows(chartColumns))

		e.PUT(path).
			WithPath("chart_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("update chart", func(t *testing.T) {
		updateCategory := updateData
		updateCategory["slug"] = "pie"

		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows(chartColumns).
				AddRow(1, time.Now(), time.Now(), nil, data["title"], data["slug"], byteDescriptionData,
					data["data_url"], byteConfigData, data["status"], data["featured_medium_id"], data["theme_id"], time.Time{}, 1))
		mock.ExpectQuery(tagQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, tag["name"], tag["slug"]))

		mock.ExpectQuery(categoryQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, category["name"], category["slug"]))

		mock.ExpectQuery(tagQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, tag["name"], tag["slug"]))

		mock.ExpectQuery(categoryQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, category["name"], category["slug"]))

		chartUpdateMock(mock, updateCategory)
		res["slug"] = "pie"
		selectAfterUpdate(mock, res)

		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(updateCategory).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)

	})

	t.Run("update chart by id with empty slug", func(t *testing.T) {

		updateCategory := updateData
		updateCategory["slug"] = "pie"

		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows(chartColumns).
				AddRow(1, time.Now(), time.Now(), nil, data["title"], data["slug"], byteDescriptionData,
					data["data_url"], byteConfigData, data["status"], data["featured_medium_id"], data["theme_id"], time.Time{}, 1))
		mock.ExpectQuery(tagQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, tag["name"], tag["slug"]))

		mock.ExpectQuery(categoryQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, category["name"], category["slug"]))

		mock.ExpectQuery(tagQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, tag["name"], tag["slug"]))

		mock.ExpectQuery(categoryQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, category["name"], category["slug"]))

		slugCheckMock(mock)

		chartUpdateMock(mock, updateCategory)
		res["slug"] = "pie"
		selectAfterUpdate(mock, res)

		updateCategory["slug"] = ""

		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(updateCategory).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)

	})

	t.Run("update chart with different slug", func(t *testing.T) {
		updateCategory := updateData
		updateCategory["slug"] = "pie-test"

		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows(chartColumns).
				AddRow(1, time.Now(), time.Now(), nil, data["title"], data["slug"], byteDescriptionData,
					data["data_url"], byteConfigData, data["status"], data["featured_medium_id"], data["theme_id"], time.Time{}, 1))
		mock.ExpectQuery(tagQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, tag["name"], tag["slug"]))

		mock.ExpectQuery(categoryQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, category["name"], category["slug"]))

		mock.ExpectQuery(tagQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, tag["name"], tag["slug"]))

		mock.ExpectQuery(categoryQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, category["name"], category["slug"]))

		mock.ExpectQuery(`SELECT slug, organisation_id FROM "bi_chart"`).
			WithArgs(fmt.Sprint(updateCategory["slug"], "%"), 1).
			WillReturnRows(sqlmock.NewRows([]string{"slug", "organisation_id"}))

		chartUpdateMock(mock, updateCategory)
		res["slug"] = "pie-test"
		selectAfterUpdate(mock, res)

		e.PUT(path).
			WithPath("chart_id", 1).
			WithHeaders(headers).
			WithJSON(updateCategory).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)

	})

}
