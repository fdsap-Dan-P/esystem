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

func TestTicketActionConditionString(t *testing.T) {

	// Test Data
	d1 := randomTicketActionConditionString()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.ItemId = item.Id

	d2 := randomTicketActionConditionString()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.ItemId = item.Id
	d2.Uuid = util.ToUUID("f2f88c0b-016d-4846-abb1-76edd6f92ca9")

	// Test Create
	CreatedD1 := createTestTicketActionConditionString(t, d1)
	CreatedD2 := createTestTicketActionConditionString(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketActionConditionString(context.Background(), CreatedD1.TicketTypeStatusId, CreatedD1.ItemId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TicketTypeStatusId, getData1.TicketTypeStatusId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value, getData1.Value)

	getData2, err2 := testQueriesTransaction.GetTicketActionConditionString(context.Background(), CreatedD2.TicketTypeStatusId, CreatedD2.ItemId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TicketTypeStatusId, getData2.TicketTypeStatusId)
	require.Equal(t, d2.ItemId, getData2.ItemId)
	require.Equal(t, d2.Value, getData2.Value)

	getData, err := testQueriesTransaction.GetTicketActionConditionStringbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTicketActionConditionString(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TicketTypeStatusId, updatedD1.TicketTypeStatusId)
	require.Equal(t, updateD2.ItemId, updatedD1.ItemId)
	require.Equal(t, updateD2.Value, updatedD1.Value)

	testListTicketActionConditionString(t, ListTicketActionConditionStringParams{
		TicketTypeStatusId: updatedD1.TicketTypeStatusId,
		Limit:              5,
		Offset:             0,
	})

	// Delete Data
	testDeleteTicketActionConditionString(t, CreatedD1.Uuid)
	testDeleteTicketActionConditionString(t, CreatedD2.Uuid)
}

func testListTicketActionConditionString(t *testing.T, arg ListTicketActionConditionStringParams) {

	TicketActionConditionString, err := testQueriesTransaction.ListTicketActionConditionString(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", TicketActionConditionString)
	require.NotEmpty(t, TicketActionConditionString)

}

func randomTicketActionConditionString() TicketActionConditionStringRequest {

	trn, _ := testQueriesTransaction.CreateTicketItem(context.Background(), RandomTicketItem())
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "LogicalCondition", 0, "Equal")

	arg := TicketActionConditionStringRequest{
		Uuid:               util.ToUUID("29818f8d-1fbb-40c5-b12d-e5a0eb3cf434"),
		TicketTypeStatusId: trn.Id,
		ConditionId:        item.Id,
		// ItemId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value: util.RandomString(20),
	}
	return arg
}

func createTestTicketActionConditionString(
	t *testing.T,
	d1 TicketActionConditionStringRequest) model.TicketActionConditionString {

	getData1, err := testQueriesTransaction.CreateTicketActionConditionString(context.Background(), d1)
	// fmt.Printf("Get by createTestTicketActionConditionString%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketTypeStatusId, getData1.TicketTypeStatusId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func updateTestTicketActionConditionString(
	t *testing.T,
	d1 TicketActionConditionStringRequest) model.TicketActionConditionString {

	getData1, err := testQueriesTransaction.UpdateTicketActionConditionString(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketTypeStatusId, getData1.TicketTypeStatusId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func testDeleteTicketActionConditionString(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketActionConditionString(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketActionConditionStringbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
