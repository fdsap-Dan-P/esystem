package main

import (
	"context"
	"database/sql"
	"log"
	"simplebank/db/datastoredb2/dwhcb"
	"simplebank/util"

	_ "github.com/ibmdb/go_ibm_db"
)

var testQueriesCustomer *dwhcb.QueriesCustomer

var testDB *sql.DB

func main() {
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open("go_ibm_db", config.DB2DWHCB)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueriesCustomer = dwhcb.New(testDB)

	var cid int64 = 1012099899
	getData2, err2 := testQueriesCustomer.GetCustomerInfo(context.Background(), cid)

	log.Printf("getData2 %+v:", getData2)
	log.Printf("err2 %+v:", err2)

}
