package chart

import (
	"encoding/json"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/factly/bindu-server/util/test"
	"github.com/joho/godotenv"
	"gopkg.in/h2non/gock.v1"
)

var headers = map[string]string{
	"X-Organisation": "1",
	"X-User":         "1",
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

var chartWithoutSlug = map[string]interface{}{
	"title": "Bar",
	"slug":  "",
}

var tag = map[string]interface{}{
	"name": "Elections",
	"slug": "elections",
}
var category = map[string]interface{}{
	"name": "Elections",
	"slug": "elections",
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

var byteThemeData, _ = json.Marshal(theme)
var byteMediumData, _ = json.Marshal(medium)

var chartColumns = []string{
	"id", "created_at", "updated_at", "deleted_at", "title", "slug", "description", "data_url", "config", "status", "featured_medium_id", "theme_id", "published_date", "organisation_id"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_chart"`)
var tagQuery = regexp.QuoteMeta(`SELECT * FROM "bi_tag"`)
var categoryQuery = regexp.QuoteMeta(`SELECT * FROM "bi_category"`)
var themeQuery = regexp.QuoteMeta(`SELECT * FROM "bi_theme"`)
var mediumQuery = regexp.QuoteMeta(`SELECT * FROM "bi_medium"`)
var deleteQuery = regexp.QuoteMeta(`UPDATE "bi_chart" SET "deleted_at"=`)
var countQuery = regexp.QuoteMeta(`SELECT count(*) FROM "bi_chart"`)
var paginationQuery = `SELECT \* FROM "bi_chart" (.+) LIMIT 1 OFFSET 1`

var url = "/charts"
var urlWithPath = "/charts/{chart_id}"

func TestMain(m *testing.M) {

	godotenv.Load("../../.env")

	// Mock kavach server and allowing persisted external traffic
	defer gock.Disable()
	test.MockServer()
	defer gock.DisableNetworking()

	exitValue := m.Run()

	os.Exit(exitValue)
}
