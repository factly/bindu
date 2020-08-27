package config

import (
	"flag"
	"log"
)

// DSN dsn
var DSN string

// KavachURL keto server url
var KavachURL string

// SetupVars setups all the config variables to run application
func SetupVars() {
	var dsn string
	var kavach string
	flag.StringVar(&dsn, "dsn", "", "Database connection string")
	flag.StringVar(&kavach, "kavach", "", "Kavach connection string")
	flag.Parse()

	if dsn == "" {
		log.Fatal("Please pass dsn flag")
	}

	if kavach == "" {
		log.Fatal("Please pass kavach flag")
	}

	DSN = dsn
	KavachURL = kavach
}
