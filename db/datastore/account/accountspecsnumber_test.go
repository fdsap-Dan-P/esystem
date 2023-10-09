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

func TestAccountSpecsNumber(t *testing.T) {

	// Test Data
	d1 := randomAccountSpecsNumber()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomAccountSpecsNumber()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestAccountSpecsNumber(t, d1)
	CreatedD2 := createTestAccountSpecsNumber(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountSpecsNumber(context.Background(), CreatedD1.AccountId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	getData2, err2 := testQueriesAccount.GetAccountSpecsNumber(context.Background(), CreatedD2.AccountId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AccountId, getData2.AccountId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.String(), getData2.Value.String())
	require.Equal(t, d2.Value2.String(), getData2.Value2.String())

	getData, err := testQueriesAccount.GetAccountSpecsNumberbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountSpecsNumber(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.AccountId, updatedD1.AccountId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.String(), updatedD1.Value.String())
	require.Equal(t, updateD2.Value2.String(), updatedD1.Value2.String())

	testListAccountSpecsNumber(t, ListAccountSpecsNumberParams{
		AccountId: updatedD1.AccountId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteAccountSpecsNumber(t, CreatedD1.Uuid)
	// testDeleteAccountSpecsNumber(t, CreatedD2.Uuid)
}

func testListAccountSpecsNumber(t *testing.T, arg ListAccountSpecsNumberParams) {

	accountSpecsNumber, err := testQueriesAccount.ListAccountSpecsNumber(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountSpecsNumber)
	require.NotEmpty(t, accountSpecsNumber)

}

func randomAccountSpecsNumber() AccountSpecsNumberRequest {

	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")
	meas, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "UnitMeasure", 0, "Item")

	arg := AccountSpecsNumberRequest{
		AccountId: acc.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:     util.RandomMoney(),
		Value2:    util.RandomMoney(),
		MeasureId: util.SetNullInt64(meas.Id),
	}
	return arg
}

func createTestAccountSpecsNumber(
	t *testing.T,
	d1 AccountSpecsNumberRequest) model.AccountSpecsNumber {

	getData1, err := testQueriesAccount.CreateAccountSpecsNumber(context.Background(), d1)
	// fmt.Printf("Get by createTestAccountSpecsNumber%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func updateTestAccountSpecsNumber(
	t *testing.T,
	d1 AccountSpecsNumberRequest) model.AccountSpecsNumber {

	getData1, err := testQueriesAccount.UpdateAccountSpecsNumber(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func testDeleteAccountSpecsNumber(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteAccountSpecsNumber(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountSpecsNumberbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
