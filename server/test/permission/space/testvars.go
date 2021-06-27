package space

import (
	"database/sql/driver"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

var headers = map[string]string{
	"X-Space": "1",
	"X-User":  "1",
}

var Data = map[string]interface{}{
	"space_id": 1,
	"charts":   10,
}

var invalidData = map[string]interface{}{
	"fact_check": 1,
}

var columns = []string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "space_id", "charts"}

var selectQuery = `SELECT (.+) FROM \"bi_space_permission\"`
var countQuery = regexp.QuoteMeta(`SELECT count(1) FROM "bi_space_permission"`)

var basePath = "/permissions/spaces"
var path = "/permissions/spaces/{permission_id}"
var mypath = "/permissions/spaces/my"

func SelectQuery(mock sqlmock.Sqlmock, args ...driver.Value) {
	mock.ExpectQuery(selectQuery).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, Data["space_id"], Data["charts"]))
}
