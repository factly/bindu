package theme

import (
	"encoding/json"
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

var columns = []string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "config"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_theme"`)
var deleteQuery = regexp.QuoteMeta(`UPDATE "bi_theme" SET "deleted_at"=`)
var paginationQuery = `SELECT \* FROM "bi_theme" (.+) LIMIT 1 OFFSET 1`

var basePath = "/themes"
var path = "/themes/{theme_id}"

//check theme exits or not
func recordNotFoundMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(100, 1).
		WillReturnRows(sqlmock.NewRows(columns))
}

func themeSelectMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, time.Now(), time.Now(), nil, 1, data["name"], byteData))

}

// check theme associated with any chart before deleting
func themeChartExpect(mock sqlmock.Sqlmock, count int) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "bi_chart"`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}

func themeCountQuery(mock sqlmock.Sqlmock, count int) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "bi_theme"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
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
