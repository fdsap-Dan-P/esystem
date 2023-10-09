package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	// account "simplebank/db/datastore/account"
	// customer "simplebank/db/datastore/customer"
	// identity "simplebank/db/datastore/identity"
	// reference "simplebank/db/datastore/reference"
	// transaction "simplebank/db/datastore/transaction"
	// user "simplebank/db/datastore/user"

	"simplebank/util"

	_ "github.com/lib/pq"
)

// var testQueriesTransaction *transaction.QueriesTransaction
// var testQueriesAccount *account.QueriesAccount
// var testQueriesReference *reference.QueriesReference
// var testQueriesUser *user.QueriesUser
// var testQueriesIdentity *identity.QueriesIdentity
// var testQueriesCustomer *customer.QueriesCustomer
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// testQueriesTransaction = transaction.New(testDB)
	// testQueriesAccount = account.New(testDB)
	// testQueriesReference = reference.New(testDB)
	// testQueriesUser = user.New(testDB)
	// testQueriesIdentity = identity.New(testDB)
	// testQueriesCustomer = customer.New(testDB)

	os.Exit(m.Run())
}
