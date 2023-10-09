package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	ref "simplebank/db/datastore/reference"
	"testing"

	"simplebank/util"

	_ "github.com/lib/pq"
)

var testQueriesIdentity *QueriesIdentity
var testQueriesReference *ref.QueriesReference
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)

	testDB, err = sql.Open(config.DBDriver, postgresqlDbInfo)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueriesIdentity = New(testDB)
	testQueriesReference = ref.New(testDB)

	os.Exit(m.Run())
}
