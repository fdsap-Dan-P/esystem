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

func TestTicketItemSpecsRef(t *testing.T) {

	// Test Data
	d1 := randomTicketItemSpecsRef()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomTicketItemSpecsRef()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestTicketItemSpecsRef(t, d1)
	CreatedD2 := createTestTicketItemSpecsRef(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketItemSpecsRef(context.Background(), CreatedD1.TicketItemId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	getData2, err2 := testQueriesTransaction.GetTicketItemSpecsRef(context.Background(), CreatedD2.TicketItemId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TicketItemId, getData2.TicketItemId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.RefId, getData2.RefId)

	getData, err := testQueriesTransaction.GetTicketItemSpecsRefbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTicketItemSpecsRef(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TicketItemId, updatedD1.TicketItemId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.RefId, updatedD1.RefId)

	testListTicketItemSpecsRef(t, ListTicketItemSpecsRefParams{
		TicketItemId: updatedD1.TicketItemId,
		Limit:        5,
		Offset:       0,
	})

	// Delete Data
	testDeleteTicketItemSpecsRef(t, CreatedD1.Uuid)
	testDeleteTicketItemSpecsRef(t, CreatedD2.Uuid)
}

func testListTicketItemSpecsRef(t *testing.T, arg ListTicketItemSpecsRefParams) {

	TicketItemSpecsRef, err := testQueriesTransaction.ListTicketItemSpecsRef(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", TicketItemSpecsRef)
	require.NotEmpty(t, TicketItemSpecsRef)

}

func randomTicketItemSpecsRef() TicketItemSpecsRefRequest {

	trn, _ := testQueriesTransaction.CreateTicketItem(context.Background(), RandomTicketItem())
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeetype", 0, "Employed")

	arg := TicketItemSpecsRefRequest{
		TicketItemId: trn.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		RefId: typ.Id,
	}
	return arg
}

func createTestTicketItemSpecsRef(
	t *testing.T,
	d1 TicketItemSpecsRefRequest) model.TicketItemSpecsRef {

	getData1, err := testQueriesTransaction.CreateTicketItemSpecsRef(context.Background(), d1)
	// fmt.Printf("Get by createTestTicketItemSpecsRef%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func updateTestTicketItemSpecsRef(
	t *testing.T,
	d1 TicketItemSpecsRefRequest) model.TicketItemSpecsRef {

	getData1, err := testQueriesTransaction.UpdateTicketItemSpecsRef(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func testDeleteTicketItemSpecsRef(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketItemSpecsRef(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketItemSpecsRefbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
