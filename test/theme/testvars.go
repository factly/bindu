package theme

import (
	"database/sql/driver"
	"encoding/json"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

var headers = map[string]string{
	"X-Space": "1",
	"X-User":  "1",
}

var data = map[string]interface{}{
	"name": "Light theme",
	"config": `{"image": { 
        "src": "Images/Sun.png",
        "name": "sun1",
        "hOffset": 250,
        "vOffset": 250,
        "alignment": "center"
    }}`,
}

var invalidData = map[string]interface{}{
	"name": "Li",
}

var byteData, _ = json.Marshal(data["config"])

var columns = []string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "organisation_id", "name", "config", "space_id"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_theme"`)
var deleteQuery = regexp.QuoteMeta(`UPDATE "bi_theme" SET "deleted_at"=`)
var paginationQuery = `SELECT \* FROM "bi_theme" (.+) LIMIT 1 OFFSET 1`

var basePath = "/themes"
var path = "/themes/{theme_id}"

//check theme exits or not
func recordNotFoundMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows(columns))
}

func SelectMock(mock sqlmock.Sqlmock, args ...driver.Value) {
	mock.ExpectQuery(selectQuery).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, 1, data["name"], byteData, 1))

}

// check theme associated with any chart before deleting
func themeChartExpect(mock sqlmock.Sqlmock, count int) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(1) FROM "bi_chart"`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}

func themeCountQuery(mock sqlmock.Sqlmock, count int) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(1) FROM "bi_theme"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}
