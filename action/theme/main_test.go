package theme

import (
	"encoding/json"
	"os"
	"regexp"
	"testing"

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
	"config": `{"image": { 
        "src": "Images/Sun.png",
        "name": "sun1",
        "hOffset": 250,
        "vOffset": 250,
        "alignment": "center"
    }}`,
}

var byteData, _ = json.Marshal(data["config"])

var themeProps = []string{"id", "created_at", "updated_at", "deleted_at", "name", "config"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_theme"`)
var chartQuery = regexp.QuoteMeta(`SELECT count(*) FROM "bi_chart"`)
var deleteQuery = regexp.QuoteMeta(`UPDATE "bi_theme" SET "deleted_at"=`)
var countQuery = regexp.QuoteMeta(`SELECT count(*) FROM "bi_theme"`)
var paginationQuery = `SELECT \* FROM "bi_theme" (.+) LIMIT 1 OFFSET 1`

var url = "/themes"
var urlWithPath = "/themes/{theme_id}"

func TestMain(m *testing.M) {

	godotenv.Load("../../.env")

	// Mock kavach server and allowing persisted external traffic
	defer gock.Disable()
	test.MockServer()
	defer gock.DisableNetworking()

	exitValue := m.Run()

	os.Exit(exitValue)
}
