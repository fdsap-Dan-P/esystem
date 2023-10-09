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

func TestTicketActionConditionRef(t *testing.T) {

	// Test Data
	d1 := randomTicketActionConditionRef()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.ItemId = item.Id

	d2 := randomTicketActionConditionRef()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.ItemId = item.Id
	d2.Uuid = util.ToUUID("9eff8f43-73a1-4784-bb9b-a19b03096300")

	// Test Create
	CreatedD1 := createTestTicketActionConditionRef(t, d1)
	CreatedD2 := createTestTicketActionConditionRef(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketActionConditionRef(context.Background(), CreatedD1.TicketTypeStatusId, CreatedD1.ItemId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TicketTypeStatusId, getData1.TicketTypeStatusId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.RefId, getData1.RefId)

	getData2, err2 := testQueriesTransaction.GetTicketActionConditionRef(context.Background(), CreatedD2.TicketTypeStatusId, CreatedD2.ItemId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TicketTypeStatusId, getData2.TicketTypeStatusId)
	require.Equal(t, d2.ItemId, getData2.ItemId)
	require.Equal(t, d2.RefId, getData2.RefId)

	getData, err := testQueriesTransaction.GetTicketActionConditionRefbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTicketActionConditionRef(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TicketTypeStatusId, updatedD1.TicketTypeStatusId)
	require.Equal(t, updateD2.ItemId, updatedD1.ItemId)
	require.Equal(t, updateD2.RefId, updatedD1.RefId)

	testListTicketActionConditionRef(t, ListTicketActionConditionRefParams{
		TicketTypeStatusId: updatedD1.TicketTypeStatusId,
		Limit:              5,
		Offset:             0,
	})

	// Delete Data
	testDeleteTicketActionConditionRef(t, CreatedD1.Uuid)
	testDeleteTicketActionConditionRef(t, CreatedD2.Uuid)
}

func testListTicketActionConditionRef(t *testing.T, arg ListTicketActionConditionRefParams) {

	TicketActionConditionRef, err := testQueriesTransaction.ListTicketActionConditionRef(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", TicketActionConditionRef)
	require.NotEmpty(t, TicketActionConditionRef)

}

func randomTicketActionConditionRef() TicketActionConditionRefRequest {

	trn, _ := testQueriesTransaction.CreateTicketItem(context.Background(), RandomTicketItem())
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeetype", 0, "Employed")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "LogicalCondition", 0, "Equal")

	arg := TicketActionConditionRefRequest{
		Uuid:               util.ToUUID("d3fc3e95-1a89-4bd5-a266-653af659b973"),
		TicketTypeStatusId: trn.Id,
		ConditionId:        item.Id,
		// ItemId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		RefId: typ.Id,
	}
	return arg
}

func createTestTicketActionConditionRef(
	t *testing.T,
	d1 TicketActionConditionRefRequest) model.TicketActionConditionRef {

	getData1, err := testQueriesTransaction.CreateTicketActionConditionRef(context.Background(), d1)
	// fmt.Printf("Get by createTestTicketActionConditionRef%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketTypeStatusId, getData1.TicketTypeStatusId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func updateTestTicketActionConditionRef(
	t *testing.T,
	d1 TicketActionConditionRefRequest) model.TicketActionConditionRef {

	getData1, err := testQueriesTransaction.UpdateTicketActionConditionRef(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketTypeStatusId, getData1.TicketTypeStatusId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func testDeleteTicketActionConditionRef(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketActionConditionRef(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketActionConditionRefbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
