package space

import (
	"database/sql/driver"
	"encoding/json"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/test/medium"
	"github.com/factly/bindu-server/util/test"
	"github.com/jinzhu/gorm/dialects/postgres"
)

func nilJsonb() postgres.Jsonb {
	ba, _ := json.Marshal(nil)
	return postgres.Jsonb{
		RawMessage: ba,
	}
}

var Data map[string]interface{} = map[string]interface{}{
	"name":               "Test Space",
	"slug":               "test-space",
	"site_title":         "Test site title",
	"tag_line":           "Test tagline",
	"description":        "Test Description",
	"site_address":       "testaddress.com",
	"logo_id":            1,
	"logo_mobile_id":     1,
	"fav_icon_id":        1,
	"mobile_icon_id":     1,
	"verification_codes": nilJsonb(),
	"social_media_urls":  nilJsonb(),
	"contact_info":       nilJsonb(),
	"organisation_id":    1,
}

var resData map[string]interface{} = map[string]interface{}{
	"name":               "Test Space",
	"slug":               "test-space",
	"site_title":         "Test site title",
	"tag_line":           "Test tagline",
	"description":        "Test Description",
	"site_address":       "testaddress.com",
	"verification_codes": nilJsonb(),
	"social_media_urls":  nilJsonb(),
	"contact_info":       nilJsonb(),
	"organisation_id":    1,
}

var invalidData map[string]interface{} = map[string]interface{}{
	"nam":             "Te",
	"slug":            "test-space",
	"organisation_id": 0,
}

var Columns = []string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "name", "slug", "site_title", "tag_line", "description", "site_address", "logo_id", "logo_mobile_id", "fav_icon_id", "mobile_icon_id", "verification_codes", "social_media_urls", "contact_info", "organisation_id"}

var selectQuery string = regexp.QuoteMeta(`SELECT * FROM "bi_space"`)
var deleteQuery string = regexp.QuoteMeta(`UPDATE "bi_space" SET "deleted_at"=`)
var countQuery string = regexp.QuoteMeta(`SELECT count(1) FROM "bi_space"`)

const path string = "/spaces/{space_id}"
const basePath string = "/spaces"

func SelectQuery(mock sqlmock.Sqlmock, args ...driver.Value) {
	mock.ExpectQuery(selectQuery).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows(Columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, Data["name"], Data["slug"], Data["site_title"], Data["tag_line"], Data["description"], Data["site_address"], Data["logo_id"], Data["logo_mobile_id"], Data["fav_icon_id"], Data["mobile_icon_id"], Data["verification_codes"], Data["social_media_urls"], Data["contact_info"], Data["organisation_id"]))
}

func slugCheckMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(Data["slug"].(string) + "%").
		WillReturnRows(sqlmock.NewRows(Columns))
}

func organisationPermissionSelectQuery(mock sqlmock.Sqlmock, args ...driver.Value) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_organisation_permission"`)).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "created_by_id", "updated_by_id", "organisation_id", "spaces"}).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, 1, 5))
}

func insertMock(mock sqlmock.Sqlmock) {
	organisationPermissionSelectQuery(mock, 1)

	mock.ExpectQuery(countQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	slugCheckMock(mock)
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "bi_space"`).
		WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, Data["name"], Data["slug"], Data["site_title"], Data["tag_line"], Data["description"], Data["site_address"], Data["verification_codes"], Data["social_media_urls"], Data["contact_info"], Data["organisation_id"]).
		WillReturnRows(sqlmock.
			NewRows([]string{"fav_icon_id", "mobile_icon_id", "logo_id", "logo_mobile_id", "id"}).
			AddRow(1, 1, 1, 1, 1))

	spacePermissionCreateQuery(mock)
}

func mediumNotFound(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_medium"`)).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug", "type", "title", "description", "caption", "alt_text", "file_size", "url", "dimensions", "space_id"}))
}

func updateMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(Columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, "name", "slug", "site_title", "tag_line", "description", "site_address", 1, 1, 1, 1, nilJsonb(), nilJsonb(), nilJsonb(), 1))

	mock.ExpectBegin()
	slugCheckMock(mock)
	medium.SelectMock(mock)
	medium.SelectMock(mock)
	medium.SelectMock(mock)
	medium.SelectMock(mock)
	mock.ExpectExec(`UPDATE \"bi_space\"`).
		WithArgs(test.AnyTime{}, 1, Data["name"], Data["slug"], Data["site_title"], Data["tag_line"], Data["description"], Data["site_address"], Data["logo_id"], Data["logo_mobile_id"], Data["fav_icon_id"], Data["mobile_icon_id"], Data["verification_codes"], Data["social_media_urls"], Data["contact_info"], 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	SelectQuery(mock, 1, 1)
	medium.SelectMock(mock)
	medium.SelectMock(mock)
	medium.SelectMock(mock)
	medium.SelectMock(mock)
}

func oneMediaIDZeroMock(mock sqlmock.Sqlmock, updateargs ...driver.Value) {
	mock.ExpectQuery(selectQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(Columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, 1, "name", "slug", "site_title", "tag_line", "description", "site_address", 1, 1, 1, 1, nilJsonb(), nilJsonb(), nilJsonb(), 1))

	mock.ExpectBegin()
	medium.SelectMock(mock)
	medium.SelectMock(mock)
	medium.SelectMock(mock)

	mock.ExpectExec(`UPDATE \"bi_space\"`).
		WithArgs(nil, test.AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	slugCheckMock(mock)
	medium.SelectMock(mock)
	medium.SelectMock(mock)
	medium.SelectMock(mock)

	mock.ExpectExec(`UPDATE \"bi_space\"`).
		WithArgs(updateargs...).
		WillReturnResult(sqlmock.NewResult(1, 1))
	SelectQuery(mock, 1, 1)
	medium.SelectMock(mock)
	medium.SelectMock(mock)
	medium.SelectMock(mock)
	medium.SelectMock(mock)
	mock.ExpectCommit()
}

func spacePermissionCreateQuery(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(`INSERT INTO "bi_space_permission"`).
		WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, 1, -1).
		WillReturnRows(sqlmock.
			NewRows([]string{"id"}).
			AddRow(1))
}
