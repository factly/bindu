package category

import (
	"os"
	"testing"

	"github.com/factly/bindu-server/util/test"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {

	test.Init()
	godotenv.Load("../../.env")

	exitValue := m.Run()

	test.CleanTables()

	os.Exit(exitValue)
}
