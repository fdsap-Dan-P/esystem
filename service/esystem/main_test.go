package service_test

import (
	"database/sql"
	"log"
	"os"
	"simplebank/util"
	"testing"

	service "simplebank/service/esystem"

	_ "github.com/lib/pq"
)

var testDumpStore *service.DumpServer
var testDB *sql.DB

// var test

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testDumpStore = service.NewDumpServer(testDB)

	// testStore = NewStoreAccount(testDB)

	os.Exit(m.Run())
}
