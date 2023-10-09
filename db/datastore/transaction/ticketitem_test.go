package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestTicketItem(t *testing.T) {

	// Test Data
	d1 := RandomTicketItem()
	d2 := RandomTicketItem()
	d1.Uuid = util.ToUUID("779b9166-5122-413f-85bb-5b3459015e8d")

	// Test Create
	CreatedD1 := createTestTicketItem(t, d1)
	CreatedD2 := createTestTicketItem(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketItem(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.TicketId, getData1.TicketId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Item, getData1.Item)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetTicketItem(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Uuid, getData2.Uuid)
	require.Equal(t, d2.TicketId, getData2.TicketId)
	require.Equal(t, d2.ItemId, getData2.ItemId)
	require.Equal(t, d2.Item, getData2.Item)
	require.Equal(t, d2.StatusId, getData2.StatusId)
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetTicketItembyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTicketItem(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Uuid, updatedD1.Uuid)
	require.Equal(t, updateD2.TicketId, updatedD1.TicketId)
	require.Equal(t, updateD2.ItemId, updatedD1.ItemId)
	require.Equal(t, updateD2.Item, updatedD1.Item)
	require.Equal(t, updateD2.StatusId, updatedD1.StatusId)
	require.Equal(t, updateD2.Remarks, updatedD1.Remarks)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListTicketItem(t, ListTicketItemParams{
		TicketId: updatedD1.TicketId,
		Limit:    5,
		Offset:   0,
	})

	// Delete Data
	testDeleteTicketItem(t, getData1.Uuid)
	testDeleteTicketItem(t, getData2.Uuid)
}

func testListTicketItem(t *testing.T, arg ListTicketItemParams) {

	ticketItem, err := testQueriesTransaction.ListTicketItem(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", ticketItem)
	require.NotEmpty(t, ticketItem)

}

func RandomTicketItem() TicketItemRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "TicketType", 0, "Over the Counter")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ticketstatus", 0, "Completed")

	tic, _ := testQueriesTransaction.CreateTicket(context.Background(), RandomTicket())

	arg := TicketItemRequest{
		Uuid:      util.ToUUID("a9a97b7d-26a6-4a0e-a618-8727b83a823d"),
		TicketId:  tic.Id,
		ItemId:    item.Id,
		StatusId:  stat.Id,
		Remarks:   util.RandomString(10),
		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestTicketItem(
	t *testing.T,
	d1 TicketItemRequest) model.TicketItem {

	getData1, err := testQueriesTransaction.CreateTicketItem(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.TicketId, getData1.TicketId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Item, getData1.Item)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestTicketItem(
	t *testing.T,
	d1 TicketItemRequest) model.TicketItem {

	getData1, err := testQueriesTransaction.UpdateTicketItem(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.TicketId, getData1.TicketId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Item, getData1.Item)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteTicketItem(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketItem(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketItembyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
