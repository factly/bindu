package tag

import (
	"os"
	"testing"

	"github.com/factly/bindu-server/util/test"
	"github.com/joho/godotenv"
	"gopkg.in/h2non/gock.v1"
)

func TestMain(m *testing.M) {

	test.Init()
	godotenv.Load("../../.env")

	// Mock kavach server and allowing persisted external traffic
	defer gock.Disable()
	test.MockServer()
	defer gock.DisableNetworking()

	exitValue := m.Run()

	test.CleanTables()

	os.Exit(exitValue)
}
