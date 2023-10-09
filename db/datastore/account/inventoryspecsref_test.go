package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"simplebank/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestInventorySpecsRef(t *testing.T) {

	// Test Data
	d1 := randomInventorySpecsRef()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomInventorySpecsRef()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestInventorySpecsRef(t, d1)
	CreatedD2 := createTestInventorySpecsRef(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetInventorySpecsRef(context.Background(), CreatedD1.InventoryItemId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	getData2, err2 := testQueriesAccount.GetInventorySpecsRef(context.Background(), CreatedD2.InventoryItemId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.InventoryItemId, getData2.InventoryItemId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.RefId, getData2.RefId)

	getData, err := testQueriesAccount.GetInventorySpecsRefbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestInventorySpecsRef(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.InventoryItemId, updatedD1.InventoryItemId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.RefId, updatedD1.RefId)

	testListInventorySpecsRef(t, ListInventorySpecsRefParams{
		InventoryItemId: updatedD1.InventoryItemId,
		Limit:           5,
		Offset:          0,
	})

	// Delete Data
	testDeleteInventorySpecsRef(t, CreatedD1.Uuid)
	testDeleteInventorySpecsRef(t, CreatedD2.Uuid)
}

func testListInventorySpecsRef(t *testing.T, arg ListInventorySpecsRefParams) {

	inventorySpecsRef, err := testQueriesAccount.ListInventorySpecsRef(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", inventorySpecsRef)
	require.NotEmpty(t, inventorySpecsRef)

}

func randomInventorySpecsRef() InventorySpecsRefRequest {

	accQtl, _ := testQueriesAccount.GetAccountInventorybyUuid(context.Background(), uuid.MustParse("c3476afe-bd50-49e6-8de3-074555a8e1bd"))
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeetype", 0, "Employed")

	arg := InventorySpecsRefRequest{
		InventoryItemId: accQtl.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		RefId: typ.Id,
	}
	return arg
}

func createTestInventorySpecsRef(
	t *testing.T,
	d1 InventorySpecsRefRequest) model.InventorySpecsRef {

	getData1, err := testQueriesAccount.CreateInventorySpecsRef(context.Background(), d1)
	// fmt.Printf("Get by createTestInventorySpecsRef%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func updateTestInventorySpecsRef(
	t *testing.T,
	d1 InventorySpecsRefRequest) model.InventorySpecsRef {

	getData1, err := testQueriesAccount.UpdateInventorySpecsRef(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func testDeleteInventorySpecsRef(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteInventorySpecsRef(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetInventorySpecsRefbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
