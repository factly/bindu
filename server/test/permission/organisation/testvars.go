package organisation

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
	"organisation_id": 1,
	"spaces":          5,
}

var invalidData = map[string]interface{}{
	"organisationid": 1,
	"spes":           5,
}

var undecodableData = map[string]interface{}{
	"organisation_id": "1",
	"spes":            5,
}

var columns = []string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "organisation_id", "spaces"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_organisation_permission"`)
var countQuery = regexp.QuoteMeta(`SELECT count(1) FROM "bi_organisation_permission"`)
var deleteQuery = regexp.QuoteMeta(`UPDATE "bi_organisation_permission" SET "deleted_at"=`)

var basePath = "/permissions/organisations"
var path = "/permissions/organisations/{permission_id}"
var mypath = "/permissions/organisations/my"

func SelectQuery(mock sqlmock.Sqlmock, args ...driver.Value) {
	mock.ExpectQuery(selectQuery).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, Data["organisation_id"], Data["spaces"]))
}

func spaceSelectQuery(mock sqlmock.Sqlmock, args ...driver.Value) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_space"`)).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "name", "slug", "site_title", "tag_line", "description", "site_address", "logo_id", "logo_mobile_id", "fav_icon_id", "mobile_icon_id", "verification_codes", "social_media_urls", "contact_info", "organisation_id"}).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, "name", "slug", "site_title", "tag_line", "description", "site_address", 1, 1, 1, 1, nil, nil, nil, 1))
}
