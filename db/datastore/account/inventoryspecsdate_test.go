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

func TestInventorySpecsDate(t *testing.T) {

	// Test Data
	d1 := randomInventorySpecsDate()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateManufactured")
	d1.SpecsId = item.Id

	d2 := randomInventorySpecsDate()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateExpired")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestInventorySpecsDate(t, d1)
	CreatedD2 := createTestInventorySpecsDate(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetInventorySpecsDate(context.Background(), CreatedD1.InventoryItemId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))

	getData2, err2 := testQueriesAccount.GetInventorySpecsDate(context.Background(), CreatedD2.InventoryItemId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.InventoryItemId, getData2.InventoryItemId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.Format("2006-01-02"), getData2.Value.Format("2006-01-02"))
	require.Equal(t, d2.Value2.Format("2006-01-02"), getData2.Value2.Format("2006-01-02"))

	getData, err := testQueriesAccount.GetInventorySpecsDatebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestInventorySpecsDate(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.InventoryItemId, updatedD1.InventoryItemId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.Format("2006-01-02"), updatedD1.Value.Format("2006-01-02"))

	testListInventorySpecsDate(t, ListInventorySpecsDateParams{
		InventoryItemId: updatedD1.InventoryItemId,
		Limit:           5,
		Offset:          0,
	})

	// Delete Data
	testDeleteInventorySpecsDate(t, CreatedD1.Uuid)
	testDeleteInventorySpecsDate(t, CreatedD2.Uuid)
}

func testListInventorySpecsDate(t *testing.T, arg ListInventorySpecsDateParams) {

	inventorySpecsDate, err := testQueriesAccount.ListInventorySpecsDate(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", inventorySpecsDate)
	require.NotEmpty(t, inventorySpecsDate)

}

func randomInventorySpecsDate() InventorySpecsDateRequest {

	accQtl, _ := testQueriesAccount.GetAccountInventorybyUuid(context.Background(), uuid.MustParse("c3476afe-bd50-49e6-8de3-074555a8e1bd"))

	arg := InventorySpecsDateRequest{
		InventoryItemId: accQtl.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomDate(),
		Value2: util.RandomDate(),
	}
	return arg
}

func createTestInventorySpecsDate(
	t *testing.T,
	d1 InventorySpecsDateRequest) model.InventorySpecsDate {

	getData1, err := testQueriesAccount.CreateInventorySpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func updateTestInventorySpecsDate(
	t *testing.T,
	d1 InventorySpecsDateRequest) model.InventorySpecsDate {

	getData1, err := testQueriesAccount.UpdateInventorySpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func testDeleteInventorySpecsDate(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteInventorySpecsDate(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetInventorySpecsDatebyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
