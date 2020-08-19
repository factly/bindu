package tag

import (
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
	"name": "Elections",
	"slug": "elections",
}

var tagWithoutSlug = map[string]interface{}{
	"name": "Politics",
	"slug": "",
}

var tagProps = []string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}

var selectQuery = regexp.QuoteMeta(`SELECT * FROM "bi_tag"`)
var chartQuery = regexp.QuoteMeta(`SELECT count(*) FROM "bi_chart" INNER JOIN "bi_chart_tag"`)
var deleteQuery = regexp.QuoteMeta(`UPDATE "bi_tag" SET "deleted_at"=`)
var countQuery = regexp.QuoteMeta(`SELECT count(*) FROM "bi_tag"`)
var paginationQuery = `SELECT \* FROM "bi_tag" (.+) LIMIT 1 OFFSET 1`

func TestMain(m *testing.M) {

	godotenv.Load("../../.env")

	// Mock kavach server and allowing persisted external traffic
	defer gock.Disable()
	test.MockServer()
	defer gock.DisableNetworking()

	exitValue := m.Run()

	os.Exit(exitValue)
}
