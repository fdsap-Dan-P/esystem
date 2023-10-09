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
	account "simplebank/db/datastore/account"
	customer "simplebank/db/datastore/customer"
	document "simplebank/db/datastore/document"
	identity "simplebank/db/datastore/identity"
	ref "simplebank/db/datastore/reference"
	reference "simplebank/db/datastore/reference"
	user "simplebank/db/datastore/user"
	"simplebank/model"

	"simplebank/util"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var testQueriesTransaction *QueriesTransaction
var testQueriesAccount *account.QueriesAccount
var testQueriesReference *reference.QueriesReference
var testQueriesUser *user.QueriesUser
var testQueriesIdentity *identity.QueriesIdentity
var testQueriesCustomer *customer.QueriesCustomer
var testQueriesDocument *document.QueriesDocument
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

	testQueriesTransaction = New(testDB)
	testQueriesAccount = account.New(testDB)
	testQueriesReference = reference.New(testDB)
	testQueriesUser = user.New(testDB)
	testQueriesIdentity = identity.New(testDB)
	testQueriesCustomer = customer.New(testDB)
	testQueriesDocument = document.New(testDB)

	os.Exit(m.Run())
}
func randomAccountTypeGroup(t *testing.T) account.AccountTypeGroupRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	prod, _ := testQueriesAccount.GetProductbyName(context.Background(), util.RandomProduct())
	grp, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccountTypeGroup", 0, "Microfinance")

	arg := account.AccountTypeGroupRequest{
		Uuid:             util.ToUUID("91788dad-3d4f-4117-9ed0-7b817da8ab12"),
		ProductId:        prod.Id,
		GroupId:          grp.Id,
		AccountTypeGroup: "Microfinance",
		NormalBalance:    true,
		Isgl:             true,
		Active:           true,
		OtherInfo:        sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func randomServer() document.ServerRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := document.ServerRequest{
		Code:         util.RandomString(30),
		Connectivity: model.Connectivity(int16(util.RandomInt32(0, 2))),
		NetAddress:   util.RandomString(10),
		Certificate:  sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		Description:  sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func randomAccountInventory(acct string) account.AccountInventoryRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), acct)

	arg := account.AccountInventoryRequest{
		Uuid:      uuid.MustParse("de5e9bff-4fa4-4470-92ca-d9776268230c"),
		AccountId: acc.Id,
		BarCode:   util.RandomNullString(10),
		Code:      util.RandomString(48),
		Quantity:  util.RandomMoney(),
		UnitPrice: util.RandomMoney(),
		BookValue: util.RandomMoney(),
		Discount:  util.RandomMoney(),
		TaxRate:   util.RandomMoney(),
		Remarks:   util.RandomString(10),
		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func randomInventoryRepository() account.InventoryRepositoryRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")

	geoParam := ref.SearchGeographyParams{
		SearchText: "Soledad San Pablo City, Laguna",
		Limit:      1,
		Offset:     0,
	}

	geo, _ := testQueriesReference.SearchGeography(context.Background(), geoParam)
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")

	arg := account.InventoryRepositoryRequest{
		CentralOfficeId:     ofc.Id,
		RepositoryCode:      util.RandomString(10),
		Repository:          util.RandomString(10),
		OfficeId:            ofc.Id,
		CustodianId:         util.SetNullInt64(ii.Id),
		GeographyId:         util.SetNullInt64(geo[0].Id),
		LocationDescription: util.RandomNullString(10),
		Remarks:             util.RandomNullString(10),
		OtherInfo:           sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func randomInventoryItem() account.InventoryItemRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	// acc, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "LoanClass", 0, "Current")
	gen, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "GenericName", 0, "Tooth Paste")
	brand, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "BrandName", 0, "Colgate")
	measure, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "UnitMeasure", 0, "Milliliter")

	log.Printf("randomInventoryItem: %+v", gen)
	arg := account.InventoryItemRequest{
		// AccountId:     util.RandomInt(1, 100),
		Uuid:            uuid.MustParse("0f5cc4a6-0969-4352-b536-0ff54a289e63"),
		BarCode:         sql.NullString{String: "", Valid: false},
		ItemName:        util.RandomString(48),
		UniqueVariation: util.RandomString(48),
		ParentId:        sql.NullInt64(sql.NullInt64{Int64: 0, Valid: false}),
		GenericNameId:   sql.NullInt64(sql.NullInt64{Int64: gen.Id, Valid: true}),
		BrandNameId:     sql.NullInt64(sql.NullInt64{Int64: brand.Id, Valid: true}),
		MeasureId:       measure.Id,
		Remarks:         util.RandomString(10),
		OtherInfo:       sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}
