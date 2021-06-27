package category

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/util/test"
)

var headers = map[string]string{
	"X-Space": "1",
	"X-User":  "1",
}

var data = map[string]interface{}{
	"name":            "Politics",
	"slug":            "politics",
	"is_for_template": true,
}

var dataWithoutSlug = map[string]interface{}{
	"name":            "Politics",
	"slug":            "",
	"is_for_template": true,
}

var invalidData = map[string]interface{}{
	"name": "ab",
}

var columns = []string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "name", "slug", "is_for_template", "space_id"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_category"`)
var deleteQuery = regexp.QuoteMeta(`UPDATE "bi_category" SET "deleted_at"=`)
var paginationQuery = `SELECT \* FROM "bi_category" (.+) LIMIT 1 OFFSET 1`

var basePath = "/categories"
var path = "/categories/{category_id}"

func slugCheckMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT slug, space_id FROM "bi_category"`)).
		WithArgs(fmt.Sprint(data["slug"], "%"), 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "space_id", "name", "slug"}))
}

func categoryInsertMock(mock sqlmock.Sqlmock) {
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "bi_category"`).
		WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, data["name"], data["slug"], "", data["is_for_template"], 1).
		WillReturnRows(sqlmock.
			NewRows([]string{"id"}).
			AddRow(1))
	mock.ExpectCommit()
}

//check category exits or not
func recordNotFoundMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows(columns))
}

func SelectMock(mock sqlmock.Sqlmock, args ...driver.Value) {
	mock.ExpectQuery(selectQuery).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, data["name"], data["slug"], data["is_for_template"], 1))
}

// check category associated with any chart before deleting
func categoryChartExpect(mock sqlmock.Sqlmock, count int) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(1) FROM "bi_chart" JOIN "bi_chart_category"`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}

func categoryUpdateMock(mock sqlmock.Sqlmock, category map[string]interface{}) {
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE \"bi_category\"`).
		WithArgs(data["is_for_template"], 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`UPDATE \"bi_category\"`).
		WithArgs(test.AnyTime{}, 1, category["name"], category["slug"], 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func selectAfterUpdate(mock sqlmock.Sqlmock, category map[string]interface{}) {
	mock.ExpectQuery(selectQuery).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, category["name"], category["slug"], category["is_for_template"], 1))
}

func categoryCountQuery(mock sqlmock.Sqlmock, count int) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(1) FROM "bi_category"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}
