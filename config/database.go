package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB - gorm DB
var DB *gorm.DB

// SetupDB is database setuo
func SetupDB(dsn interface{}) {

	var err error
	DB, err = gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	// Query log
	DB.LogMode(true)

	DB.SingularTable(true)

	// adding default prefix to table name
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "bi_" + defaultTableName
	}

	fmt.Println("connected to database ...")
}
