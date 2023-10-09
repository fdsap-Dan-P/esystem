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

func TestTicketItemSpecsDate(t *testing.T) {

	// Test Data
	d1 := randomTicketItemSpecsDate()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateManufactured")
	d1.SpecsId = item.Id

	d2 := randomTicketItemSpecsDate()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateExpired")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestTicketItemSpecsDate(t, d1)
	CreatedD2 := createTestTicketItemSpecsDate(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketItemSpecsDate(context.Background(), CreatedD1.TicketItemId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	getData2, err2 := testQueriesTransaction.GetTicketItemSpecsDate(context.Background(), CreatedD2.TicketItemId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TicketItemId, getData2.TicketItemId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.Format("2006-01-02"), getData2.Value.Format("2006-01-02"))
	require.Equal(t, d2.Value2.Format("2006-01-02"), getData2.Value2.Format("2006-01-02"))

	getData, err := testQueriesTransaction.GetTicketItemSpecsDatebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTicketItemSpecsDate(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TicketItemId, updatedD1.TicketItemId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.Format("2006-01-02"), updatedD1.Value.Format("2006-01-02"))
	require.Equal(t, updateD2.Value2.Format("2006-01-02"), updatedD1.Value2.Format("2006-01-02"))

	testListTicketItemSpecsDate(t, ListTicketItemSpecsDateParams{
		TicketItemId: updatedD1.TicketItemId,
		Limit:        5,
		Offset:       0,
	})

	// Delete Data
	testDeleteTicketItemSpecsDate(t, CreatedD1.Uuid)
	testDeleteTicketItemSpecsDate(t, CreatedD2.Uuid)
}

func testListTicketItemSpecsDate(t *testing.T, arg ListTicketItemSpecsDateParams) {

	TicketItemSpecsDate, err := testQueriesTransaction.ListTicketItemSpecsDate(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", TicketItemSpecsDate)
	require.NotEmpty(t, TicketItemSpecsDate)
}

func randomTicketItemSpecsDate() TicketItemSpecsDateRequest {

	trn, _ := testQueriesTransaction.CreateTicketItem(context.Background(), RandomTicketItem())

	arg := TicketItemSpecsDateRequest{
		TicketItemId: trn.Id,
		// SpecsId: sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomDate(),
		Value2: util.RandomDate(),
	}
	return arg
}

func createTestTicketItemSpecsDate(
	t *testing.T,
	d1 TicketItemSpecsDateRequest) model.TicketItemSpecsDate {

	getData1, err := testQueriesTransaction.CreateTicketItemSpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func updateTestTicketItemSpecsDate(
	t *testing.T,
	d1 TicketItemSpecsDateRequest) model.TicketItemSpecsDate {

	getData1, err := testQueriesTransaction.UpdateTicketItemSpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func testDeleteTicketItemSpecsDate(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketItemSpecsDate(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketItemSpecsDatebyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
