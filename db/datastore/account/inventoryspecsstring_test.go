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

func TestInventorySpecsString(t *testing.T) {

	// Test Data
	d1 := randomInventorySpecsString()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Number of Cores")
	d1.SpecsId = item.Id

	d2 := randomInventorySpecsString()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Number of Threads")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestInventorySpecsString(t, d1)
	CreatedD2 := createTestInventorySpecsString(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetInventorySpecsString(context.Background(), CreatedD1.InventoryItemId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	getData2, err2 := testQueriesAccount.GetInventorySpecsString(context.Background(), CreatedD2.InventoryItemId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.InventoryItemId, getData2.InventoryItemId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value, getData2.Value)

	getData, err := testQueriesAccount.GetInventorySpecsStringbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestInventorySpecsString(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.InventoryItemId, updatedD1.InventoryItemId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value, updatedD1.Value)

	testListInventorySpecsString(t, ListInventorySpecsStringParams{
		InventoryItemId: updatedD1.InventoryItemId,
		Limit:           5,
		Offset:          0,
	})

	// Delete Data
	testDeleteInventorySpecsString(t, CreatedD1.Uuid)
	testDeleteInventorySpecsString(t, CreatedD2.Uuid)
}

func testListInventorySpecsString(t *testing.T, arg ListInventorySpecsStringParams) {

	inventorySpecsString, err := testQueriesAccount.ListInventorySpecsString(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", inventorySpecsString)
	require.NotEmpty(t, inventorySpecsString)

}

func randomInventorySpecsString() InventorySpecsStringRequest {

	accQtl, _ := testQueriesAccount.GetAccountInventorybyUuid(context.Background(), uuid.MustParse("de5e9bff-4fa4-4470-92ca-d9776268230c"))

	arg := InventorySpecsStringRequest{
		InventoryItemId: accQtl.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value: util.RandomString(10),
	}
	return arg
}

func createTestInventorySpecsString(
	t *testing.T,
	d1 InventorySpecsStringRequest) model.InventorySpecsString {

	getData1, err := testQueriesAccount.CreateInventorySpecsString(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func updateTestInventorySpecsString(
	t *testing.T,
	d1 InventorySpecsStringRequest) model.InventorySpecsString {

	getData1, err := testQueriesAccount.UpdateInventorySpecsString(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func testDeleteInventorySpecsString(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteInventorySpecsString(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetInventorySpecsStringbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
