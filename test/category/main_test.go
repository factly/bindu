package category

import (
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/util/test"
	"gopkg.in/h2non/gock.v1"
)

var headers = map[string]string{
	"X-Organisation": "1",
	"X-User":         "1",
}

var data = map[string]interface{}{
	"name": "Politics",
	"slug": "politics",
}

var dataWithoutSlug = map[string]interface{}{
	"name": "Politics",
	"slug": "",
}

var invalidData = map[string]interface{}{
	"name": "ab",
}

var columns = []string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "name", "slug"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_category"`)
var deleteQuery = regexp.QuoteMeta(`UPDATE "bi_category" SET "deleted_at"=`)
var paginationQuery = `SELECT \* FROM "bi_category" (.+) LIMIT 1 OFFSET 1`

var basePath = "/categories"
var path = "/categories/{category_id}"

func slugCheckMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT slug, organisation_id FROM "bi_category"`)).
		WithArgs(fmt.Sprint(data["slug"], "%"), 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}))
}

func categoryInsertMock(mock sqlmock.Sqlmock) {
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "bi_category"`).
		WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, data["name"], data["slug"], "", 1).
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

func categorySelectMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, data["name"], data["slug"]))
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
		WithArgs(test.AnyTime{}, 1, category["name"], category["slug"], 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
}

func selectAfterUpdate(mock sqlmock.Sqlmock, category map[string]interface{}) {
	mock.ExpectQuery(selectQuery).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, category["name"], category["slug"]))
}

func categoryCountQuery(mock sqlmock.Sqlmock, count int) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(1) FROM "bi_category"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
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
