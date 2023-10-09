package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	access "simplebank/db/datastore/access"
	account "simplebank/db/datastore/account"
	identity "simplebank/db/datastore/identity"
	ref "simplebank/db/datastore/reference"
	"testing"

	"simplebank/util"

	_ "github.com/lib/pq"
)

var testQueriesUser *QueriesUser
var testQueriesIdentity *identity.QueriesIdentity
var testQueriesAccess *access.QueriesAccess
var testQueriesAccount *account.QueriesAccount
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

	testQueriesUser = New(testDB)
	testQueriesIdentity = identity.New(testDB)
	testQueriesAccess = access.New(testDB)
	testQueriesReference = ref.New(testDB)
	testQueriesAccount = account.New(testDB)

	os.Exit(m.Run())
}
