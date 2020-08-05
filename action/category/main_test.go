package category

import (
	"os"
	"testing"

	"github.com/factly/bindu-server/util/test"
)

func TestMain(m *testing.M) {

	test.Init()

	exitValue := m.Run()

	//test.CleanTables()

	os.Exit(exitValue)
}
