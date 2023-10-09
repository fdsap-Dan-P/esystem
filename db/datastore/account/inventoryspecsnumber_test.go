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

func TestInventorySpecsNumber(t *testing.T) {

	// Test Data
	d1 := randomInventorySpecsNumber()
	// itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "Storage")
	d1.SpecsId = item.Id

	d2 := randomInventorySpecsNumber()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestInventorySpecsNumber(t, d1)
	CreatedD2 := createTestInventorySpecsNumber(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetInventorySpecsNumber(context.Background(), CreatedD1.InventoryItemId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	getData2, err2 := testQueriesAccount.GetInventorySpecsNumber(context.Background(), CreatedD2.InventoryItemId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.InventoryItemId, getData2.InventoryItemId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.String(), getData2.Value.String())
	require.Equal(t, d2.Value2.String(), getData2.Value2.String())

	getData, err := testQueriesAccount.GetInventorySpecsNumberbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestInventorySpecsNumber(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.InventoryItemId, updatedD1.InventoryItemId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.String(), updatedD1.Value.String())
	require.Equal(t, updateD2.Value2.String(), updatedD1.Value2.String())

	testListInventorySpecsNumber(t, ListInventorySpecsNumberParams{
		InventoryItemId: updatedD1.InventoryItemId,
		Limit:           5,
		Offset:          0,
	})

	// Delete Data
	testDeleteInventorySpecsNumber(t, CreatedD1.Uuid)
	testDeleteInventorySpecsNumber(t, CreatedD2.Uuid)
}

func testListInventorySpecsNumber(t *testing.T, arg ListInventorySpecsNumberParams) {

	inventorySpecsNumber, err := testQueriesAccount.ListInventorySpecsNumber(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", inventorySpecsNumber)
	require.NotEmpty(t, inventorySpecsNumber)

}

func randomInventorySpecsNumber() InventorySpecsNumberRequest {

	accQtl, _ := testQueriesAccount.GetInventoryItembyUuid(context.Background(), uuid.MustParse("0df94671-3193-4440-bf0d-ec7f171b294e"))

	arg := InventorySpecsNumberRequest{
		InventoryItemId: accQtl.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomMoney(),
		Value2: util.RandomMoney(),
	}
	return arg
}

func createTestInventorySpecsNumber(
	t *testing.T,
	d1 InventorySpecsNumberRequest) model.InventorySpecsNumber {

	getData1, err := testQueriesAccount.CreateInventorySpecsNumber(context.Background(), d1)
	// fmt.Printf("Get by createTestInventorySpecsNumber%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func updateTestInventorySpecsNumber(
	t *testing.T,
	d1 InventorySpecsNumberRequest) model.InventorySpecsNumber {

	getData1, err := testQueriesAccount.UpdateInventorySpecsNumber(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func testDeleteInventorySpecsNumber(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteInventorySpecsNumber(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetInventorySpecsNumberbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
