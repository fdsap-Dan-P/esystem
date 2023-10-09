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

func TestTicketActionConditionNumber(t *testing.T) {

	// Test Data
	d1 := randomTicketActionConditionNumber()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.ItemId = item.Id

	d2 := randomTicketActionConditionNumber()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.ItemId = item.Id
	d2.Uuid = util.ToUUID("a1c08d97-e550-4484-920e-6a12de82dcce")

	// Test Create
	CreatedD1 := createTestTicketActionConditionNumber(t, d1)
	CreatedD2 := createTestTicketActionConditionNumber(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketActionConditionNumber(context.Background(), CreatedD1.TicketTypeStatusId, CreatedD1.ItemId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TicketTypeStatusId, getData1.TicketTypeStatusId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	getData2, err2 := testQueriesTransaction.GetTicketActionConditionNumber(context.Background(), CreatedD2.TicketTypeStatusId, CreatedD2.ItemId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TicketTypeStatusId, getData2.TicketTypeStatusId)
	require.Equal(t, d2.ItemId, getData2.ItemId)
	require.Equal(t, d2.Value.String(), getData2.Value.String())
	require.Equal(t, d2.Value2.String(), getData2.Value2.String())

	getData, err := testQueriesTransaction.GetTicketActionConditionNumberbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTicketActionConditionNumber(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TicketTypeStatusId, updatedD1.TicketTypeStatusId)
	require.Equal(t, updateD2.ItemId, updatedD1.ItemId)
	require.Equal(t, updateD2.Value.String(), updatedD1.Value.String())
	require.Equal(t, updateD2.Value2.String(), updatedD1.Value2.String())

	testListTicketActionConditionNumber(t, ListTicketActionConditionNumberParams{
		TicketTypeStatusId: updatedD1.TicketTypeStatusId,
		Limit:              5,
		Offset:             0,
	})

	// Delete Data
	testDeleteTicketActionConditionNumber(t, CreatedD1.Uuid)
	testDeleteTicketActionConditionNumber(t, CreatedD2.Uuid)
}

func testListTicketActionConditionNumber(t *testing.T, arg ListTicketActionConditionNumberParams) {

	TicketActionConditionNumber, err := testQueriesTransaction.ListTicketActionConditionNumber(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", TicketActionConditionNumber)
	require.NotEmpty(t, TicketActionConditionNumber)

}

func randomTicketActionConditionNumber() TicketActionConditionNumberRequest {

	trn, _ := testQueriesTransaction.CreateTicketItem(context.Background(), RandomTicketItem())
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "LogicalCondition", 0, "Equal")

	log.Printf("randomTicketActionConditionNumber: %v", item)

	arg := TicketActionConditionNumberRequest{
		Uuid:               util.ToUUID("86df5636-a260-41bc-80ef-d6372b2e37dd"),
		TicketTypeStatusId: trn.Id,
		ConditionId:        item.Id,
		// ItemId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomMoney(),
		Value2: util.RandomMoney(),
	}
	return arg
}

func createTestTicketActionConditionNumber(
	t *testing.T,
	d1 TicketActionConditionNumberRequest) model.TicketActionConditionNumber {

	getData1, err := testQueriesTransaction.CreateTicketActionConditionNumber(context.Background(), d1)
	// fmt.Printf("Get by createTestTicketActionConditionNumber%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketTypeStatusId, getData1.TicketTypeStatusId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func updateTestTicketActionConditionNumber(
	t *testing.T,
	d1 TicketActionConditionNumberRequest) model.TicketActionConditionNumber {

	getData1, err := testQueriesTransaction.UpdateTicketActionConditionNumber(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketTypeStatusId, getData1.TicketTypeStatusId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func testDeleteTicketActionConditionNumber(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketActionConditionNumber(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketActionConditionNumberbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
