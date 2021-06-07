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
	"title":       "test title",
	"slug":        "test-slug",
	"spec":        nilJsonb(),
	"properties":  nilJsonb(),
	"medium_id":   1,
	"category_id": 1,
	"is_default":  false,
	"mode":        "vega",
}

var templateList = []map[string]interface{}{
	{
		"title":       "test title 1",
		"slug":        "test-slug-1",
		"spec":        nilJsonb(),
		"properties":  nilJsonb(),
		"medium_id":   1,
		"category_id": 1,
		"is_default":  false,
		"mode":        "vega",
	},
	{
		"title":       "test title 2",
		"slug":        "test-slug-2",
		"spec":        nilJsonb(),
		"properties":  nilJsonb(),
		"medium_id":   1,
		"category_id": 1,
		"is_default":  false,
		"mode":        "vega",
	},
}

var columns = []string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "title", "slug", "spec", "properties", "category_id", "medium_id", "is_default", "mode", "space_id"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_template"`)

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
			AddRow("1", time.Now(), time.Now(), nil, 1, 1, data["title"], data["slug"], data["spec"], data["properties"], data["category_id"], data["medium_id"], data["is_default"], data["mode"], 1))
}
