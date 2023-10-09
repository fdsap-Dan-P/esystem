package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	common "simplebank/db/common"
	acc "simplebank/db/datastore/account"
	cust "simplebank/db/datastore/customer"
	identity "simplebank/db/datastore/identity"
	ref "simplebank/db/datastore/reference"
	trn "simplebank/db/datastore/transaction"
	usr "simplebank/db/datastore/user"
	"simplebank/util"

	_ "github.com/lib/pq"
)

var testQueriesSchool StoreSchool
var testQueriesReference *ref.QueriesReference
var testQueriesTransaction *trn.QueriesTransaction
var testQueriesCustomer *cust.QueriesCustomer
var testQueriesIdentity *identity.QueriesIdentity
var testQueriesAccount *acc.QueriesAccount
var testQueriesUser *usr.QueriesUser
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

	testQueriesSchool = NewStoreSchool(testDB)
	testQueriesReference = ref.New(testDB)
	testQueriesCustomer = cust.New(testDB)
	testQueriesIdentity = identity.New(testDB)
	testQueriesAccount = acc.New(testDB)
	testQueriesTransaction = trn.New(testDB)
	testQueriesUser = usr.New(testDB)
	// testStore = NewStoreSchool(testDB)

	os.Exit(m.Run())
}

func RandomTicket() trn.TicketRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "tickettype", 0, "Over the Counter")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ticketstatus", 0, "Completed")
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "0000")
	fmt.Printf("%+v\n", usr)
	arg := trn.TicketRequest{
		Uuid:            util.ToUUID("a9a97b7d-26a6-4a0e-a618-8727b83a823d"),
		CentralOfficeId: ofc.Id,
		TicketDate:      util.RandomDate(),
		TicketTypeId:    typ.Id,
		PostedbyId:      usr.Id,
		StatusId:        stat.Id,
		Remarks:         sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func RandomTicketItem() trn.TicketItemRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "TicketType", 0, "Over the Counter")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ticketstatus", 0, "Completed")

	tic, _ := testQueriesTransaction.CreateTicket(context.Background(), RandomTicket())

	arg := trn.TicketItemRequest{
		Uuid:      util.ToUUID("a9a97b7d-26a6-4a0e-a618-8727b83a823d"),
		TicketId:  tic.Id,
		ItemId:    item.Id,
		StatusId:  stat.Id,
		Remarks:   util.RandomString(10),
		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}
