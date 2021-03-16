package chart

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
)

var headers = map[string]string{
	"X-Space": "1",
	"X-User":  "1",
}

var invalidData = map[string]interface{}{
	"title":    "Pi",
	"theme_id": 0,
}

var data = map[string]interface{}{
	"title": "Pie",
	"slug":  "pie",
	"description": `{
		"data": [
			{
			"type": "articles",
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
			"self": "http://example.com/articles?page[number]=3&page[size]=1",
			"first": "http://example.com/articles?page[number]=1&page[size]=1",
			"prev": "http://example.com/articles?page[number]=2&page[size]=1",
			"next": "http://example.com/articles?page[number]=4&page[size]=1",
			"last": "http://example.com/articles?page[number]=13&page[size]=1"
		  }
	}`,
	"status":             "available",
	"featured_medium":    "data:image/jpeg;base64,/9j/2wCEAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRofHh0aHBwgJC4nICIsIxwcKDcpLDAxNDQ0Hyc5PTgyPC4zNDIBCQkJDAsMGA0NGDIhHCEyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMv/AABEIAZABkAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/APn+iiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKK1NE8Pan4huWg063MpTBdiQqoD3JPSlKSirt6DjFydktTLorvJ/hLr8VuZI7jT53H/LKOVgx/wC+lA/WuJurWexupLa5iaKaNiro4wQazp1qdT4HcupRqU/jViGiiitTMKKKKACiiigAooooAKKmtbWe9uorW2iaWeVgqIoyWPoKm1LSb/SJkh1C1ktpHXeqyDBK5Iz+YNK6vYdna5TooopiCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAK9l8HP/Ynwsn1O2jX7QyT3GSOrLlVz6gben19a8ar2bSv+SKH/ryuP/Q5K4MfrGEejkjuwGkpy6qLOK8I+JtZk8Yack+pXU0dxcLHIkkpZWDHB4PHetD4uWiReIrS5QANPbDf7lWIz+WB+FcJaXU1jdw3VtIY54XDxuOqsDkGtY6hqXizWtPttSvHneSVIEZsDaGYDsPetZULVlVWiSszNV70XSerbK9h4c1rVIDPY6Xd3EIOPMjiJXPpnpVG6s7mxuGt7u3lgmT70cqFWH4GvY/H3ia88J2+lWujeVArBwQYwwCLtCqAfqao/E+3jvvCml6wY1FxvRSwHO10LYz6AgY9MmsaWNlJxco2jK9u/wAz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igAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAqwb67PW6m/wC/hqvRRa4JtE/227/5+pv+/ho+3Xf/AD9Tf9/DUFFKyHzPuT/brv8A5+pv+/ho+23f/PzN/wB/DUFFFkHM+5P9tuv+fmb/AL7NH226/wCfmb/vs1BRRZBzPuT/AG27/wCfmb/v4aPtt3/z9Tf9/DUFFFkHM+4pJYkk5J6mkoopiCiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAP/Z",
	"featured_medium_id": uint(1),
	"theme_id":           uint(1),
	"is_public":          true,
	"published_date":     time.Time{},
	"category_ids":       []int{1},
	"tag_ids":            []int{1},
}

var byteDescriptionData, _ = json.Marshal(data["description"])
var byteConfigData, _ = json.Marshal(data["config"])

var dataWithoutSlug = map[string]interface{}{
	"title":     "Pie",
	"slug":      "",
	"is_public": true,
	"description": `{
		"data": [
			{
			"type": "articles",
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
			"self": "http://example.com/articles?page[number]=3&page[size]=1",
			"first": "http://example.com/articles?page[number]=1&page[size]=1",
			"prev": "http://example.com/articles?page[number]=2&page[size]=1",
			"next": "http://example.com/articles?page[number]=4&page[size]=1",
			"last": "http://example.com/articles?page[number]=13&page[size]=1"
		  }
	}`,
	"status":             "available",
	"featured_medium_id": uint(1),
	"theme_id":           uint(1),
	"published_date":     time.Time{},
	"category_ids":       []int{1},
	"tag_ids":            []int{1},
}

var tag = map[string]interface{}{
	"name":        "Elections",
	"slug":        "elections",
	"description": "desc",
}
var category = map[string]interface{}{
	"name":        "Elections",
	"slug":        "elections",
	"description": "desc",
}

var theme = map[string]interface{}{
	"name": "Light theme",
	"config": `{"image": { 
        "src": "Images/Sun.png",
        "name": "sun1",
        "hOffset": 250,
        "vOffset": 250,
        "alignment": "center"
    }}`,
}

var medium = map[string]interface{}{
	"name": "Politics",
	"slug": "politics",
	"type": "jpg",
	"url": postgres.Jsonb{
		RawMessage: []byte(`{"raw":"http://testimage.com/test.jpg"}`),
	},
}

var byteThemeData, _ = json.Marshal(theme["config"])

var columns = []string{
	"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "title", "slug", "description", "data_url", "config", "status", "featured_medium_id", "theme_id", "published_date", "space_id"}

var selectQuery = `SELECT (.+) FROM "bi_chart"`
var tagQuery = regexp.QuoteMeta(`SELECT * FROM "bi_tag"`)
var categoryQuery = regexp.QuoteMeta(`SELECT * FROM "bi_category"`)
var themeQuery = regexp.QuoteMeta(`SELECT * FROM "bi_theme"`)
var mediumQuery = regexp.QuoteMeta(`SELECT * FROM "bi_medium"`)
var deleteQuery = regexp.QuoteMeta(`UPDATE "bi_chart" SET "deleted_at"=`)
var countQuery = regexp.QuoteMeta(`SELECT count(1) FROM "bi_chart"`)
var paginationQuery = `SELECT \* FROM "bi_chart" (.+) LIMIT 1 OFFSET 1`

var basePath = "/charts"
var path = "/charts/{chart_id}"

func validateAssociations(result *httpexpect.Object) {
	result.Value("medium").
		Object().
		ContainsMap(medium)

	result.Value("theme").
		Object().
		ContainsMap(theme)
}

func recordNotFoundMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows(columns))

}
func SelectMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, data["title"], data["slug"], byteDescriptionData,
				data["data_url"], byteConfigData, data["status"], data["featured_medium_id"], data["theme_id"], time.Time{}, 1))
}

func selectAfterUpdate(mock sqlmock.Sqlmock, chart map[string]interface{}) {
	description, _ := json.Marshal(chart["description"])
	config, _ := json.Marshal(chart["config"])
	mock.ExpectQuery(selectQuery).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, chart["title"], chart["slug"], description,
				chart["data_url"], config, chart["status"], chart["featured_medium_id"], chart["theme_id"], time.Time{}, 1))

	chartPreloadMock(mock)
}

func chartTagUpdate(mock sqlmock.Sqlmock, err error) {
	tagQueryMock(mock)
	mediumQueryMock(mock)
	themeQueryMock(mock)

	if err != nil {
		mock.ExpectQuery(`INSERT INTO "bi_tag"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, "Elections", "elections", "", 1, 1).
			WillReturnError(err)
		mock.ExpectExec(`INSERT INTO "bi_chart_tag"`).
			WithArgs(1, 1).
			WillReturnError(err)
		mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "bi_chart_tag"`)).
			WithArgs(1, 1).
			WillReturnError(err)
	} else {
		mock.ExpectQuery(`INSERT INTO "bi_tag"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, "Elections", "elections", "", 1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectExec(`INSERT INTO "bi_chart_tag"`).
			WithArgs(1, 1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "bi_chart_tag"`)).
			WithArgs(1, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}
}

func chartCategoryUpdate(mock sqlmock.Sqlmock, err error) {
	categoryQueryMock(mock)
	mediumQueryMock(mock)
	themeQueryMock(mock)

	if err != nil {
		mock.ExpectQuery(`INSERT INTO "bi_category"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, "Elections", "elections", "", 1, 1).
			WillReturnError(err)
		mock.ExpectExec(`INSERT INTO "bi_chart_category"`).
			WithArgs(1, 1).
			WillReturnError(err)

		mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "bi_chart_category"`)).
			WithArgs(1, 1).
			WillReturnError(err)
	} else {
		mock.ExpectQuery(`INSERT INTO "bi_category"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, "Elections", "elections", "", 1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectExec(`INSERT INTO "bi_chart_category"`).
			WithArgs(1, 1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "bi_chart_category"`)).
			WithArgs(1, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}
}

func chartUpdateMock(mock sqlmock.Sqlmock, chart map[string]interface{}) {
	description, _ := json.Marshal(chart["description"])
	config, _ := json.Marshal(chart["config"])

	mock.ExpectExec(`UPDATE \"bi_chart\"`).
		WithArgs(true, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mediumQueryMock(mock)
	themeQueryMock(mock)
	mock.ExpectExec(`UPDATE \"bi_chart\"`).
		WithArgs(test.AnyTime{}, 1, chart["title"], chart["slug"], description, chart["data_url"], config, chart["status"], chart["featured_medium_id"], chart["theme_id"], 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func chartInsertMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(mediumQuery).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "organisation_id", "name", "slug", "type", "url"}).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, 1, medium["name"], medium["slug"], medium["type"], medium["url"]))
	mock.ExpectQuery(themeQuery).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "organisation_id", "name", "config"}).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, 1, theme["name"], byteThemeData))

	mock.ExpectQuery(`INSERT INTO "bi_chart"`).
		WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, data["title"], data["slug"], byteDescriptionData,
			data["data_url"], byteConfigData, data["status"], data["is_public"], data["theme_id"], test.AnyTime{}, 1, data["featured_medium_id"]).
		WillReturnRows(sqlmock.NewRows([]string{"featured_medium_id", "id"}).AddRow(1, 1))

	mock.ExpectQuery(`INSERT INTO "bi_tag"`).
		WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, "Elections", "elections", "", 1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectExec(`INSERT INTO "bi_chart_tag"`).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectQuery(`INSERT INTO "bi_category"`).
		WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, "Elections", "elections", "", 1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectExec(`INSERT INTO "bi_chart_category"`).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
}

func chartPreloadMock(mock sqlmock.Sqlmock, args ...driver.Value) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_chart_category"`)).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows(append([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}, []string{"category_id", "chart_id"}...)).
			AddRow(1, time.Now(), time.Now(), nil, "title1", "slug1", 1, 1))

	categoryQueryMock(mock)

	mediumQueryMock(mock)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_chart_tag"`)).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows(append([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}, []string{"tag_id", "chart_id"}...)).
			AddRow(1, time.Now(), time.Now(), nil, "title1", "slug1", 1, 1))

	tagQueryMock(mock)

	themeQueryMock(mock)
}

func tagQueryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(tagQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "space_id", "name", "slug"}).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, 1, tag["name"], tag["slug"]))
}

func categoryQueryMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(categoryQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "space_id", "name", "slug"}).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, 1, category["name"], category["slug"]))
}

func mediumQueryMock(mock sqlmock.Sqlmock, args ...driver.Value) {
	mock.ExpectQuery(mediumQuery).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "organisation_id", "name", "slug", "type", "url"}).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, 1, medium["name"], medium["slug"], medium["type"], medium["url"]))
}

func themeQueryMock(mock sqlmock.Sqlmock, args ...driver.Value) {
	mock.ExpectQuery(themeQuery).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "organisation_id", "name", "config"}).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, 1, theme["name"], byteThemeData))
}

func slugCheckMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT slug, space_id FROM "bi_chart"`)).
		WithArgs(fmt.Sprint(data["slug"], "%"), 1).
		WillReturnRows(sqlmock.NewRows([]string{"space_id", "slug"}))
}
