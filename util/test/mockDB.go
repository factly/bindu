package test

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/config"
)

// AnyTime To match time for test sqlmock queries
type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

// SetupMockDB setups the mock sql db
func SetupMockDB() sqlmock.Sqlmock {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}

	config.SetupDB(db)

	return mock
}
