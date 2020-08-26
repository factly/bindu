package medium

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/test"
	"github.com/factly/x/loggerx"
	"github.com/go-chi/chi"
	"gopkg.in/h2non/gock.v1"
)

var headers = map[string]string{
	"X-Organisation": "1",
	"X-User":         "1",
}
var invalidData = map[string]interface{}{
	"name": "po",
}

var data = map[string]interface{}{
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

var mediumWithoutSlug = map[string]interface{}{
	"name": "Politics",
	"slug": "",
	"type": "jpg",
	"url": `{"image": { 
        "src": "Images/Sun.png",
        "name": "sun1",
        "hOffset": 250,
        "vOffset": 250,
        "alignment": "center"
    }}`,
}

var byteData, _ = json.Marshal(data["url"])

var columns = []string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug", "type", "url"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_medium"`)
var chartQuery = regexp.QuoteMeta(`SELECT count(*) FROM "bi_chart"`)
var deleteQuery = regexp.QuoteMeta(`UPDATE "bi_medium" SET "deleted_at"=`)
var paginationQuery = `SELECT \* FROM "bi_medium" (.+) LIMIT 1 OFFSET 1`

var basePath = "/media"
var path = "/media/{medium_id}"

func mediumSelectMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, data["name"], data["slug"], data["type"], byteData))
}

func mediumChartExpect(mock sqlmock.Sqlmock, count int) {
	mock.ExpectQuery(chartQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}

//check medium exits or not
func recordNotFoundMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(100, 1).
		WillReturnRows(sqlmock.NewRows(columns))
}

func slugCheckMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT slug, organisation_id FROM "bi_medium"`)).
		WithArgs(fmt.Sprint(data["slug"], "%"), 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}))
}

func mediumInsertMock(mock sqlmock.Sqlmock) {
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "bi_medium"`).
		WithArgs(test.AnyTime{}, test.AnyTime{}, nil, data["name"], data["slug"], data["type"], byteData, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
}

func mediumCountQuery(mock sqlmock.Sqlmock, count int) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "bi_medium"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}

func mediumUpdateMock(mock sqlmock.Sqlmock, medium map[string]interface{}) {
	var urlByteData, _ = json.Marshal(medium["url"])
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE \"bi_medium\" SET (.+)  WHERE (.+) \"bi_medium\".\"id\" = `).
		WithArgs(medium["name"], medium["slug"], medium["type"], test.AnyTime{}, urlByteData, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
}

func Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(loggerx.Init())
	r.With(util.CheckUser, util.CheckOrganisation).Mount(basePath, Router())
	return r
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
