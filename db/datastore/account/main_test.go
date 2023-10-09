package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	cust "simplebank/db/datastore/customer"
	identity "simplebank/db/datastore/identity"
	ref "simplebank/db/datastore/reference"
	"simplebank/util"

	_ "github.com/lib/pq"
)

var testQueriesAccount StoreAccount
var testQueriesReference *ref.QueriesReference
var testQueriesCustomer *cust.QueriesCustomer
var testQueriesIdentity *identity.QueriesIdentity
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

	testQueriesAccount = NewStoreAccount(testDB)
	testQueriesReference = ref.New(testDB)
	testQueriesCustomer = cust.New(testDB)
	testQueriesIdentity = identity.New(testDB)
	// testStore = NewStoreAccount(testDB)

	os.Exit(m.Run())
}
