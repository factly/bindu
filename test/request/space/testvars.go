package space

import (
	"database/sql/driver"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm/dialects/postgres"
)

var headers = map[string]string{
	"X-Space": "1",
	"X-User":  "1",
}

var Data = map[string]interface{}{
	"title": "Test Title",
	"description": postgres.Jsonb{
		RawMessage: []byte(`{"type":"description"}`),
	},
	"status":   "pending",
	"space_id": 1,
	"charts":   10,
}

var requestList = []map[string]interface{}{
	{
		"title": "Test Title 1",
		"description": postgres.Jsonb{
			RawMessage: []byte(`{"type":"description1"}`),
		},
		"status":   "pending",
		"space_id": 1,
		"charts":   10,
	},
	{
		"title": "Test Title 2",
		"description": postgres.Jsonb{
			RawMessage: []byte(`{"type":"description2"}`),
		},
		"status":   "pending",
		"space_id": 1,
		"charts":   20,
	},
}

var invalidData = map[string]interface{}{
	"title": "aa",
}

var Columns = []string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "title", "description", "status", "charts", "space_id"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_space_permission_request"`)
var countQuery = regexp.QuoteMeta(`SELECT count(1) FROM "bi_space_permission_request"`)

var basePath = "/requests/spaces"
var path = "/requests/spaces/{request_id}"
var approvePath = "/requests/spaces/{request_id}/approve"
var rejectPath = "/requests/spaces/{request_id}/reject"
var myPath = "/requests/spaces/my"

func SelectQuery(mock sqlmock.Sqlmock, args ...driver.Value) {
	mock.ExpectQuery(selectQuery).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows(Columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, Data["title"], Data["description"], Data["status"], Data["charts"], Data["space_id"]))
}
