package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// var store StoreAccount

//	func init() {
//		store = NewStoreAccount(testDB)
//	}
func TestAccountInventory(t *testing.T) {

	// Test Data

	// store := NewStoreAccount(testDB)

	d1 := randomAccountInventory()

	d2 := randomAccountInventory()
	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000002")
	d2.AccountId = acc.Id
	d2.Uuid = uuid.MustParse("36d4f001-759f-4623-95d3-bd5112354512")

	fmt.Printf("Get by UUId%+v\n", d1)
	// Test Create
	CreatedD1 := createTestAccountInventory(t, d1)
	CreatedD2 := createTestAccountInventory(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountInventory(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.BarCode, getData1.BarCode)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.Quantity.String(), getData1.Quantity.String())
	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
	require.Equal(t, d1.BookValue.String(), getData1.BookValue.String())
	require.Equal(t, d1.Discount.String(), getData1.Discount.String())
	require.Equal(t, d1.TaxRate.String(), getData1.TaxRate.String())
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetAccountInventory(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AccountId, getData2.AccountId)
	require.Equal(t, d2.BarCode, getData2.BarCode)
	require.Equal(t, d2.Code, getData2.Code)
	require.Equal(t, d2.Quantity.String(), getData2.Quantity.String())
	require.Equal(t, d2.UnitPrice.String(), getData2.UnitPrice.String())
	require.Equal(t, d2.BookValue.String(), getData2.BookValue.String())
	require.Equal(t, d2.Discount.String(), getData2.Discount.String())
	require.Equal(t, d2.TaxRate.String(), getData2.TaxRate.String())
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetAccountInventorybyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountInventory(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updatedD1.AccountId, updateD2.AccountId)
	require.Equal(t, updatedD1.BarCode, updateD2.BarCode)
	require.Equal(t, updatedD1.Code, updateD2.Code)
	require.Equal(t, updatedD1.Quantity.String(), updateD2.Quantity.String())
	require.Equal(t, updatedD1.UnitPrice.String(), updateD2.UnitPrice.String())
	require.Equal(t, updatedD1.BookValue.String(), updateD2.BookValue.String())
	require.Equal(t, updatedD1.Discount.String(), updateD2.Discount.String())
	require.Equal(t, updatedD1.TaxRate.String(), updateD2.TaxRate.String())
	require.Equal(t, updatedD1.Remarks, updateD2.Remarks)
	require.JSONEq(t, updatedD1.OtherInfo.String, updateD2.OtherInfo.String)

	testListAccountInventory(t, ListAccountInventoryParams{
		AccountId: updatedD1.AccountId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	//testDeleteAccountInventory(t, getData1.Id)
	testDeleteAccountInventory(t, getData2.Id)
}

func testListAccountInventory(t *testing.T, arg ListAccountInventoryParams) {
	// store := NewStoreAccount(testDB)
	accountInventory, err := testQueriesAccount.ListAccountInventory(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountInventory)
	require.NotEmpty(t, accountInventory)
}

func randomAccountInventory() AccountInventoryRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")

	arg := AccountInventoryRequest{
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

func createTestAccountInventory(
	t *testing.T,
	d1 AccountInventoryRequest) model.AccountInventory {
	// store := NewStoreAccount(testDB)

	getData1, err := testQueriesAccount.CreateAccountInventory(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.BarCode, getData1.BarCode)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.Quantity.String(), getData1.Quantity.String())
	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
	require.Equal(t, d1.BookValue.String(), getData1.BookValue.String())
	require.Equal(t, d1.Discount.String(), getData1.Discount.String())
	require.Equal(t, d1.TaxRate.String(), getData1.TaxRate.String())
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestAccountInventory(
	t *testing.T,
	d1 AccountInventoryRequest) model.AccountInventory {

	getData1, err := testQueriesAccount.UpdateAccountInventory(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.BarCode, getData1.BarCode)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.Quantity.String(), getData1.Quantity.String())
	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
	require.Equal(t, d1.BookValue.String(), getData1.BookValue.String())
	require.Equal(t, d1.Discount.String(), getData1.Discount.String())
	require.Equal(t, d1.TaxRate.String(), getData1.TaxRate.String())
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAccountInventory(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteAccountInventory(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountInventory(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
