package chart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/test"
	"github.com/factly/x/loggerx"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-chi/chi"
	"gopkg.in/h2non/gock.v1"
)

var headers = map[string]string{
	"X-Organisation": "1",
	"X-User":         "1",
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
	"featured_medium_id": uint(1),
	"theme_id":           uint(1),
	"published_date":     time.Time{},
	"category_ids":       []int{1},
	"tag_ids":            []int{1},
}

var byteDescriptionData, _ = json.Marshal(data["description"])
var byteConfigData, _ = json.Marshal(data["config"])

var dataWithoutSlug = map[string]interface{}{
	"title": "Pie",
	"slug":  "",
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
	"url": `{"image": { 
        "src": "Images/Sun.png",
        "name": "sun1",
        "hOffset": 250,
        "vOffset": 250,
        "alignment": "center"
    }}`,
}

var byteThemeData, _ = json.Marshal(theme["config"])
var byteMediumData, _ = json.Marshal(medium["url"])

var columns = []string{
	"id", "created_at", "updated_at", "deleted_at", "title", "slug", "description", "data_url", "config", "status", "featured_medium_id", "theme_id", "published_date", "organisation_id"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_chart"`)
var tagQuery = regexp.QuoteMeta(`SELECT * FROM "bi_tag"`)
var categoryQuery = regexp.QuoteMeta(`SELECT * FROM "bi_category"`)
var themeQuery = regexp.QuoteMeta(`SELECT * FROM "bi_theme"`)
var mediumQuery = regexp.QuoteMeta(`SELECT * FROM "bi_medium"`)
var deleteQuery = regexp.QuoteMeta(`UPDATE "bi_chart" SET "deleted_at"=`)
var countQuery = regexp.QuoteMeta(`SELECT count(*) FROM "bi_chart"`)
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
		WithArgs(100, 1).
		WillReturnRows(sqlmock.NewRows(columns))

}

func selectAfterUpdate(mock sqlmock.Sqlmock, chart map[string]interface{}) {
	description, _ := json.Marshal(chart["description"])
	config, _ := json.Marshal(chart["config"])
	mock.ExpectQuery(selectQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, chart["title"], chart["slug"], description,
				chart["data_url"], config, chart["status"], chart["featured_medium_id"], chart["theme_id"], time.Time{}, 1))

	mock.ExpectQuery(mediumQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug", "type", "url"}).
			AddRow(1, time.Now(), time.Now(), nil, 1, medium["name"], medium["slug"], medium["type"], byteMediumData))
	mock.ExpectQuery(themeQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "config"}).
			AddRow(1, time.Now(), time.Now(), nil, 1, theme["name"], byteThemeData))

	mock.ExpectQuery(tagQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
			AddRow(1, time.Now(), time.Now(), nil, 1, tag["name"], tag["slug"]))

	mock.ExpectQuery(categoryQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
			AddRow(1, time.Now(), time.Now(), nil, 1, category["name"], category["slug"]))
}

func chartUpdateMock(mock sqlmock.Sqlmock, chart map[string]interface{}) {
	description, _ := json.Marshal(chart["description"])
	config, _ := json.Marshal(chart["config"])

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
		WithArgs(config, chart["data_url"],
			description, chart["featured_medium_id"], chart["slug"],
			chart["status"], chart["theme_id"], chart["title"], test.AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(`UPDATE \"bi_tag\" SET (.+)  WHERE (.+) \"bi_tag\".\"id\" = `).
		WithArgs(test.AnyTime{}, test.AnyTime{}, nil, tag["name"], tag["slug"], "", 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(`INSERT INTO "bi_chart_tag"`).
		WithArgs(1, 1, 1, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec(`UPDATE \"bi_category\" SET (.+)  WHERE (.+) \"bi_category\".\"id\" = `).
		WithArgs(test.AnyTime{}, test.AnyTime{}, nil, category["name"], category["slug"], "", 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(`INSERT INTO "bi_chart_category"`).
		WithArgs(1, 1, 1, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()
}

func slugCheckMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT slug, organisation_id FROM "bi_chart"`)).
		WithArgs(fmt.Sprint(data["slug"], "%"), 1).
		WillReturnRows(sqlmock.NewRows([]string{"organisation_id", "slug"}))
}

func Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(loggerx.Init())
	r.With(util.CheckUser, util.CheckOrganisation).Mount(basePath, Router())
	return r
}

func TestMain(m *testing.M) {

	test.SetEnv()

	// Mock kavach server and allowing persisted external traffic
	defer gock.Disable()
	test.MockServer()
	defer gock.DisableNetworking()

	exitValue := m.Run()

	os.Exit(exitValue)
}
