package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"testing"

	access "simplebank/db/datastore/access"
	account "simplebank/db/datastore/account"
	customer "simplebank/db/datastore/customer"
	identity "simplebank/db/datastore/identity"
	reference "simplebank/db/datastore/reference"
	user "simplebank/db/datastore/user"

	"simplebank/util"

	_ "github.com/lib/pq"
)

var (
	testQueriesDocument  *QueriesDocument
	testQueriesAccount   *account.QueriesAccount
	testQueriesReference *reference.QueriesReference
	testQueriesUser      *user.QueriesUser
	testQueriesIdentity  *identity.QueriesIdentity
	testQueriesCustomer  *customer.QueriesCustomer
	testQueriesAccess    *access.QueriesAccess
	testDB               *sql.DB
	sorsImgPath          string
	homePath             string
	targetPath           string
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	sorsImgPath = path.Join(config.HomeFolder, "static/uploads/images")
	homePath = config.HomeFolder //path.Join(config.HomeFolder, "app/images/")
	targetPath = "app/images"

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)

	testDB, err = sql.Open(config.DBDriver, postgresqlDbInfo)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueriesDocument = New(testDB)
	testQueriesAccount = account.New(testDB)
	testQueriesReference = reference.New(testDB)
	testQueriesUser = user.New(testDB)
	testQueriesAccess = access.New(testDB)
	testQueriesIdentity = identity.New(testDB)
	testQueriesCustomer = customer.New(testDB)

	os.Exit(m.Run())
}
