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

func TestTicketActionConditionDate(t *testing.T) {

	// Test Data
	d1 := randomTicketActionConditionDate(t)
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateManufactured")
	d1.ItemId = item.Id

	d2 := randomTicketActionConditionDate(t)
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateExpired")
	d2.ItemId = item.Id
	d2.Uuid = util.ToUUID("f478a4fc-afc6-4091-b0ef-80dc097073ec")

	// Test Create
	CreatedD1 := createTestTicketActionConditionDate(t, d1)
	CreatedD2 := createTestTicketActionConditionDate(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketActionConditionDate(context.Background(), CreatedD1.TicketTypeStatusId, CreatedD1.ItemId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TicketTypeStatusId, getData1.TicketTypeStatusId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	getData2, err2 := testQueriesTransaction.GetTicketActionConditionDate(context.Background(), CreatedD2.TicketTypeStatusId, CreatedD2.ItemId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TicketTypeStatusId, getData2.TicketTypeStatusId)
	require.Equal(t, d2.ItemId, getData2.ItemId)
	require.Equal(t, d2.Value.Format("2006-01-02"), getData2.Value.Format("2006-01-02"))
	require.Equal(t, d2.Value2.Format("2006-01-02"), getData2.Value2.Format("2006-01-02"))

	getData, err := testQueriesTransaction.GetTicketActionConditionDatebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTicketActionConditionDate(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TicketTypeStatusId, updatedD1.TicketTypeStatusId)
	require.Equal(t, updateD2.ItemId, updatedD1.ItemId)
	require.Equal(t, updateD2.Value.Format("2006-01-02"), updatedD1.Value.Format("2006-01-02"))
	require.Equal(t, updateD2.Value2.Format("2006-01-02"), updatedD1.Value2.Format("2006-01-02"))

	testListTicketActionConditionDate(t, ListTicketActionConditionDateParams{
		TicketTypeStatusId: updatedD1.TicketTypeStatusId,
		Limit:              5,
		Offset:             0,
	})

	// Delete Data
	testDeleteTicketActionConditionDate(t, CreatedD1.Uuid)
	testDeleteTicketActionConditionDate(t, CreatedD2.Uuid)
}

func testListTicketActionConditionDate(t *testing.T, arg ListTicketActionConditionDateParams) {

	TicketActionConditionDate, err := testQueriesTransaction.ListTicketActionConditionDate(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", TicketActionConditionDate)
	require.NotEmpty(t, TicketActionConditionDate)
}

func randomTicketActionConditionDate(t *testing.T) TicketActionConditionDateRequest {

	trn, _ := testQueriesTransaction.CreateTicketTypeStatus(context.Background(), RandomTicketTypeStatus(t))
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "LogicalCondition", 0, "Equal")

	arg := TicketActionConditionDateRequest{
		Uuid:               util.ToUUID("c64db986-e574-44a3-a4e1-4ee292eea211"),
		TicketTypeStatusId: trn.Id,
		ConditionId:        item.Id,
		// ItemId: sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomDate(),
		Value2: util.RandomDate(),
	}
	return arg
}

func createTestTicketActionConditionDate(
	t *testing.T,
	d1 TicketActionConditionDateRequest) model.TicketActionConditionDate {

	getData1, err := testQueriesTransaction.CreateTicketActionConditionDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketTypeStatusId, getData1.TicketTypeStatusId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func updateTestTicketActionConditionDate(
	t *testing.T,
	d1 TicketActionConditionDateRequest) model.TicketActionConditionDate {

	getData1, err := testQueriesTransaction.UpdateTicketActionConditionDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketTypeStatusId, getData1.TicketTypeStatusId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func testDeleteTicketActionConditionDate(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketActionConditionDate(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketActionConditionDatebyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
