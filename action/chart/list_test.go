package chart

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

func TestChartList(t *testing.T) {
	mock := test.SetupMockDB()

	testServer := httptest.NewServer(Routes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	chartlist := []map[string]interface{}{
		{
			"title": "Chart Test 1",
			"slug":  "chart-test-1",
			"description": `{
				"data": [
					{
					"type": "charts",
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
			"data_url": "http://data.com/crime?page[number]=3&page[size]=1",
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
		},
		{
			"title": "Chart Test 2",
			"slug":  "chart-test-2",
			"description": `{
				"data": [
					{
					"type": "pie",
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
			"data_url": "http://data.com/crime?page[number]=3&page[size]=1",
			"config": `{
				"links": {
					"self": "http://example.com/pie?page[number]=3&page[size]=1",
				  }
			}`,
			"status":             "available",
			"featured_medium_id": uint(1),
			"theme_id":           uint(1),
			"published_date":     time.Time{},
		},
	}
	byteConfigDataOne, _ := json.Marshal(chartlist[0]["config"])
	byteConfigDataTwo, _ := json.Marshal(chartlist[1]["config"])

	byteDescriptionDataOne, _ := json.Marshal(chartlist[0]["description"])
	byteDescriptionDataTwo, _ := json.Marshal(chartlist[1]["description"])

	t.Run("get empty list of categories", func(t *testing.T) {

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("0"))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns))

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_tag" INNER JOIN "bi_chart_tag"`)).
			WillReturnRows(sqlmock.NewRows(append([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}, []string{"tag_id", "chart_id"}...)))
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_category" INNER JOIN "bi_chart_category"`)).
			WillReturnRows(sqlmock.NewRows(append([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}, []string{"category_id", "chart_id"}...)))

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

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(chartlist)))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, chartlist[0]["title"], chartlist[0]["slug"], byteDescriptionDataOne,
					chartlist[0]["data_url"], byteConfigDataOne, chartlist[0]["status"], chartlist[0]["featured_medium_id"], chartlist[0]["theme_id"], time.Time{}, 1).
				AddRow(2, time.Now(), time.Now(), nil, chartlist[1]["title"], chartlist[1]["slug"], byteDescriptionDataTwo,
					chartlist[1]["data_url"], byteConfigDataTwo, chartlist[1]["status"], chartlist[1]["featured_medium_id"], chartlist[1]["theme_id"], time.Time{}, 1))

		mock.ExpectQuery(mediumQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug", "type", "url"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, medium["name"], medium["slug"], medium["type"], byteMediumData))
		mock.ExpectQuery(themeQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "config"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, theme["name"], byteThemeData))
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_tag" INNER JOIN "bi_chart_tag"`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows(append([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}, []string{"tag_id", "chart_id"}...)).
				AddRow(1, time.Now(), time.Now(), nil, "title1", "slug1", 1, 1))
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_category" INNER JOIN "bi_chart_category"`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows(append([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}, []string{"category_id", "chart_id"}...)).
				AddRow(1, time.Now(), time.Now(), nil, "title1", "slug1", 1, 1))

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

	t.Run("get categories with pagination", func(t *testing.T) {
		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(chartlist)))

		mock.ExpectQuery(paginationQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(2, time.Now(), time.Now(), nil, chartlist[1]["title"], chartlist[1]["slug"], byteDescriptionDataTwo,
					chartlist[1]["data_url"], byteConfigDataTwo, chartlist[1]["status"], chartlist[1]["featured_medium_id"], chartlist[1]["theme_id"], time.Time{}, 1))

		mock.ExpectQuery(mediumQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug", "type", "url"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, medium["name"], medium["slug"], medium["type"], byteMediumData))
		mock.ExpectQuery(themeQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "config"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, theme["name"], byteThemeData))
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_tag" INNER JOIN "bi_chart_tag"`)).
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows(append([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}, []string{"tag_id", "chart_id"}...)).
				AddRow(1, time.Now(), time.Now(), nil, "title1", "slug1", 1, 1))
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_category" INNER JOIN "bi_chart_category"`)).
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows(append([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}, []string{"category_id", "chart_id"}...)).
				AddRow(1, time.Now(), time.Now(), nil, "title1", "slug1", 1, 1))

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
