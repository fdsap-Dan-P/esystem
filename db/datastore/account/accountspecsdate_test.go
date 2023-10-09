package db

import (
	"context"
	"database/sql"
	"log"

	"fmt"
	"testing"

	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAccountSpecsDate(t *testing.T) {

	// Test Data
	d1 := randomAccountSpecsDate()
	// --	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerItem")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "DateManufactured")
	d1.SpecsId = item.Id

	d2 := randomAccountSpecsDate()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "DateExpired")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestAccountSpecsDate(t, d1)
	CreatedD2 := createTestAccountSpecsDate(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountSpecsDate(context.Background(), CreatedD1.AccountId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	getData2, err2 := testQueriesAccount.GetAccountSpecsDate(context.Background(), CreatedD2.AccountId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AccountId, getData2.AccountId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.Format("2006-01-02"), getData2.Value.Format("2006-01-02"))
	require.Equal(t, d2.Value2.Format("2006-01-02"), getData2.Value2.Format("2006-01-02"))

	getData, err := testQueriesAccount.GetAccountSpecsDatebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountSpecsDate(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.AccountId, updatedD1.AccountId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.Format("2006-01-02"), updatedD1.Value.Format("2006-01-02"))
	require.Equal(t, updateD2.Value2.Format("2006-01-02"), updatedD1.Value2.Format("2006-01-02"))

	testListAccountSpecsDate(t, ListAccountSpecsDateParams{
		AccountId: updatedD1.AccountId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteAccountSpecsDate(t, CreatedD1.Uuid)
	testDeleteAccountSpecsDate(t, CreatedD2.Uuid)
}

func testListAccountSpecsDate(t *testing.T, arg ListAccountSpecsDateParams) {

	accountSpecsDate, err := testQueriesAccount.ListAccountSpecsDate(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountSpecsDate)
	require.NotEmpty(t, accountSpecsDate)
}

func randomAccountSpecsDate() AccountSpecsDateRequest {

	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")

	log.Printf("acc: %v", acc)

	arg := AccountSpecsDateRequest{
		AccountId: acc.Id,
		// SpecsId: sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomDate(),
		Value2: util.RandomDate(),
	}
	return arg
}

func createTestAccountSpecsDate(
	t *testing.T,
	d1 AccountSpecsDateRequest) model.AccountSpecsDate {

	getData1, err := testQueriesAccount.CreateAccountSpecsDate(context.Background(), d1)
	log.Printf("%v", d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func updateTestAccountSpecsDate(
	t *testing.T,
	d1 AccountSpecsDateRequest) model.AccountSpecsDate {

	getData1, err := testQueriesAccount.UpdateAccountSpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func testDeleteAccountSpecsDate(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteAccountSpecsDate(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountSpecsDatebyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
