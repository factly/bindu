package template

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm/dialects/postgres"
)

func nilJsonb() postgres.Jsonb {
	ba, _ := json.Marshal(nil)
	return postgres.Jsonb{
		RawMessage: ba,
	}
}

var headers = map[string]string{
	"X-Space": "1",
	"X-User":  "1",
}

var data = map[string]interface{}{
	"title":      "test title",
	"slug":       "test-slug",
	"schema":     nilJsonb(),
	"properties": nilJsonb(),
	"medium_id":  1,
}

var templateList = []map[string]interface{}{
	{
		"title":      "test title 1",
		"slug":       "test-slug-1",
		"schema":     nilJsonb(),
		"properties": nilJsonb(),
		"medium_id":  1,
	},
	{
		"title":      "test title 2",
		"slug":       "test-slug-2",
		"schema":     nilJsonb(),
		"properties": nilJsonb(),
		"medium_id":  1,
	},
}

var columns = []string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "title", "slug", "schema", "properties", "medium_id", "space_id"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_template"`)
var countQuery = regexp.QuoteMeta(`SELECT count(1) FROM "bi_template"`)

var basePath = "/templates"
var path = "/templates/{template_id}"

func slugCheckMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT slug, space_id FROM "bi_template"`)).
		WithArgs(fmt.Sprint(data["slug"], "%"), 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "space_id", "name", "slug"}))
}

func SelectMock(mock sqlmock.Sqlmock, args ...driver.Value) {
	mock.ExpectQuery(selectQuery).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, data["title"], data["slug"], data["schema"], data["properties"], data["medium_id"], 1))
}
