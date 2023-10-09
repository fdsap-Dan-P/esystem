package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAccountSpecsString(t *testing.T) {

	// Test Data
	d1 := randomAccountSpecsString()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomAccountSpecsString()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestAccountSpecsString(t, d1)
	CreatedD2 := createTestAccountSpecsString(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountSpecsString(context.Background(), CreatedD1.AccountId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	getData2, err2 := testQueriesAccount.GetAccountSpecsString(context.Background(), CreatedD2.AccountId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AccountId, getData2.AccountId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value, getData2.Value)

	getData, err := testQueriesAccount.GetAccountSpecsStringbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountSpecsString(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.AccountId, updatedD1.AccountId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value, updatedD1.Value)

	testListAccountSpecsString(t, ListAccountSpecsStringParams{
		AccountId: updatedD1.AccountId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteAccountSpecsString(t, CreatedD1.Uuid)
	testDeleteAccountSpecsString(t, CreatedD2.Uuid)
}

func testListAccountSpecsString(t *testing.T, arg ListAccountSpecsStringParams) {

	accountSpecsString, err := testQueriesAccount.ListAccountSpecsString(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountSpecsString)
	require.NotEmpty(t, accountSpecsString)

}

func randomAccountSpecsString() AccountSpecsStringRequest {

	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")

	arg := AccountSpecsStringRequest{
		AccountId: acc.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value: util.RandomString(20),
	}
	return arg
}

func createTestAccountSpecsString(
	t *testing.T,
	d1 AccountSpecsStringRequest) model.AccountSpecsString {

	getData1, err := testQueriesAccount.CreateAccountSpecsString(context.Background(), d1)
	// fmt.Printf("Get by createTestAccountSpecsString%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func updateTestAccountSpecsString(
	t *testing.T,
	d1 AccountSpecsStringRequest) model.AccountSpecsString {

	getData1, err := testQueriesAccount.UpdateAccountSpecsString(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func testDeleteAccountSpecsString(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteAccountSpecsString(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountSpecsStringbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
