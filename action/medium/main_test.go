package medium

import (
	"encoding/json"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/util/test"
	"github.com/joho/godotenv"
	"gopkg.in/h2non/gock.v1"
)

var headers = map[string]string{
	"X-Organisation": "1",
	"X-User":         "1",
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

var mediumProps = []string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug", "type", "url"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_medium"`)
var chartQuery = regexp.QuoteMeta(`SELECT count(*) FROM "bi_chart"`)
var deleteQuery = regexp.QuoteMeta(`UPDATE "bi_medium" SET "deleted_at"=`)
var countQuery = regexp.QuoteMeta(`SELECT count(*) FROM "bi_medium"`)
var paginationQuery = `SELECT \* FROM "bi_medium" (.+) LIMIT 1 OFFSET 1`

var url = "/media"
var urlWithPath = "/media/{medium_id}"

func mediumSelectMock(mock sqlmock.Sqlmock) {
	mock.ExpectQuery(selectQuery).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows(mediumProps).
			AddRow(1, time.Now(), time.Now(), nil, 1, data["name"], data["slug"], data["type"], byteData))
}

func mediumChartExpect(mock sqlmock.Sqlmock, count int) {
	mock.ExpectQuery(chartQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}

func TestMain(m *testing.M) {

	godotenv.Load("../../.env")

	// Mock kavach server and allowing persisted external traffic
	defer gock.Disable()
	test.MockServer()
	defer gock.DisableNetworking()

	exitValue := m.Run()

	os.Exit(exitValue)
}
