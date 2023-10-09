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

func TestTicketItemSpecsNumber(t *testing.T) {

	// Test Data
	d1 := randomTicketItemSpecsNumber()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomTicketItemSpecsNumber()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestTicketItemSpecsNumber(t, d1)
	CreatedD2 := createTestTicketItemSpecsNumber(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketItemSpecsNumber(context.Background(), CreatedD1.TicketItemId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	getData2, err2 := testQueriesTransaction.GetTicketItemSpecsNumber(context.Background(), CreatedD2.TicketItemId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TicketItemId, getData2.TicketItemId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.String(), getData2.Value.String())
	require.Equal(t, d2.Value2.String(), getData2.Value2.String())

	getData, err := testQueriesTransaction.GetTicketItemSpecsNumberbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTicketItemSpecsNumber(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TicketItemId, updatedD1.TicketItemId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.String(), updatedD1.Value.String())
	require.Equal(t, updateD2.Value2.String(), updatedD1.Value2.String())

	testListTicketItemSpecsNumber(t, ListTicketItemSpecsNumberParams{
		TicketItemId: updatedD1.TicketItemId,
		Limit:        5,
		Offset:       0,
	})

	// Delete Data
	testDeleteTicketItemSpecsNumber(t, CreatedD1.Uuid)
	testDeleteTicketItemSpecsNumber(t, CreatedD2.Uuid)
}

func testListTicketItemSpecsNumber(t *testing.T, arg ListTicketItemSpecsNumberParams) {

	TicketItemSpecsNumber, err := testQueriesTransaction.ListTicketItemSpecsNumber(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", TicketItemSpecsNumber)
	require.NotEmpty(t, TicketItemSpecsNumber)

}

func randomTicketItemSpecsNumber() TicketItemSpecsNumberRequest {

	trn, _ := testQueriesTransaction.CreateTicketItem(context.Background(), RandomTicketItem())

	arg := TicketItemSpecsNumberRequest{
		TicketItemId: trn.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomMoney(),
		Value2: util.RandomMoney(),
	}
	return arg
}

func createTestTicketItemSpecsNumber(
	t *testing.T,
	d1 TicketItemSpecsNumberRequest) model.TicketItemSpecsNumber {

	getData1, err := testQueriesTransaction.CreateTicketItemSpecsNumber(context.Background(), d1)
	// fmt.Printf("Get by createTestTicketItemSpecsNumber%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func updateTestTicketItemSpecsNumber(
	t *testing.T,
	d1 TicketItemSpecsNumberRequest) model.TicketItemSpecsNumber {

	getData1, err := testQueriesTransaction.UpdateTicketItemSpecsNumber(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func testDeleteTicketItemSpecsNumber(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketItemSpecsNumber(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketItemSpecsNumberbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
