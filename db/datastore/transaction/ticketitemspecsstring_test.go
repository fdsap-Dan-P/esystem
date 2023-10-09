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

func TestTicketItemSpecsString(t *testing.T) {

	// Test Data
	d1 := randomTicketItemSpecsString()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomTicketItemSpecsString()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestTicketItemSpecsString(t, d1)
	CreatedD2 := createTestTicketItemSpecsString(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketItemSpecsString(context.Background(), CreatedD1.TicketItemId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	getData2, err2 := testQueriesTransaction.GetTicketItemSpecsString(context.Background(), CreatedD2.TicketItemId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TicketItemId, getData2.TicketItemId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value, getData2.Value)

	getData, err := testQueriesTransaction.GetTicketItemSpecsStringbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTicketItemSpecsString(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TicketItemId, updatedD1.TicketItemId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value, updatedD1.Value)

	testListTicketItemSpecsString(t, ListTicketItemSpecsStringParams{
		TicketItemId: updatedD1.TicketItemId,
		Limit:        5,
		Offset:       0,
	})

	// Delete Data
	testDeleteTicketItemSpecsString(t, CreatedD1.Uuid)
	testDeleteTicketItemSpecsString(t, CreatedD2.Uuid)
}

func testListTicketItemSpecsString(t *testing.T, arg ListTicketItemSpecsStringParams) {

	TicketItemSpecsString, err := testQueriesTransaction.ListTicketItemSpecsString(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", TicketItemSpecsString)
	require.NotEmpty(t, TicketItemSpecsString)

}

func randomTicketItemSpecsString() TicketItemSpecsStringRequest {

	trn, _ := testQueriesTransaction.CreateTicketItem(context.Background(), RandomTicketItem())

	arg := TicketItemSpecsStringRequest{
		TicketItemId: trn.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value: util.RandomString(20),
	}
	return arg
}

func createTestTicketItemSpecsString(
	t *testing.T,
	d1 TicketItemSpecsStringRequest) model.TicketItemSpecsString {

	getData1, err := testQueriesTransaction.CreateTicketItemSpecsString(context.Background(), d1)
	// fmt.Printf("Get by createTestTicketItemSpecsString%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func updateTestTicketItemSpecsString(
	t *testing.T,
	d1 TicketItemSpecsStringRequest) model.TicketItemSpecsString {

	getData1, err := testQueriesTransaction.UpdateTicketItemSpecsString(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func testDeleteTicketItemSpecsString(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketItemSpecsString(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketItemSpecsStringbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
