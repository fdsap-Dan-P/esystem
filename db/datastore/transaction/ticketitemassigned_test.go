package db

import (
	"context"
	"database/sql"
	"log"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestTicketItemAssigned(t *testing.T) {

	// Test Data
	d1 := RandomTicketItemAssigned()

	d2 := RandomTicketItemAssigned()
	// act, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ActionList", 0, "Posted")
	// d2.ActionId = act.Id
	d2.Uuid = util.ToUUID("779b9166-5122-413f-85bb-5b3459015e8d")

	// Test Create
	CreatedD1 := createTestTicketItemAssigned(t, d1)
	CreatedD2 := createTestTicketItemAssigned(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketItemAssignedbyUuid(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.AssignedById, getData1.AssignedById)
	require.Equal(t, d1.AssignedDate.Format("2006-01-02"), getData1.AssignedDate.Format("2006-01-02"))
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetTicketItemAssignedbyUuid(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TicketItemId, getData2.TicketItemId)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.AssignedById, getData2.AssignedById)
	require.Equal(t, d2.AssignedDate.Format("2006-01-02"), getData2.AssignedDate.Format("2006-01-02"))
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.Equal(t, d2.StatusId, getData2.StatusId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetTicketItemAssignedbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, getData.TicketItemId, CreatedD1.TicketItemId)
	require.Equal(t, getData.UserId, CreatedD1.UserId)
	require.Equal(t, getData.AssignedById, CreatedD1.AssignedById)
	require.Equal(t, getData.AssignedDate.Format("2006-01-02"), CreatedD1.AssignedDate.Format("2006-01-02"))
	require.Equal(t, getData.Remarks, CreatedD1.Remarks)
	require.Equal(t, getData.StatusId, CreatedD1.StatusId)
	require.JSONEq(t, getData.OtherInfo.String, CreatedD1.OtherInfo.String)

	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	// updateD2.Id = getData2.Id
	updateD2.Remarks = updateD2.Remarks + "-Edited"

	updatedD1 := updateTestTicketItemAssigned(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TicketItemId, updatedD1.TicketItemId)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.AssignedById, updatedD1.AssignedById)
	require.Equal(t, updateD2.AssignedDate.Format("2006-01-02"), updatedD1.AssignedDate.Format("2006-01-02"))
	require.Equal(t, updateD2.Remarks, updatedD1.Remarks)
	require.Equal(t, updateD2.StatusId, updatedD1.StatusId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListTicketItemAssigned(t, ListTicketItemAssignedParams{
		TicketItemId: updatedD1.TicketItemId,
		Limit:        5,
		Offset:       0,
	})

	// Delete Data
	testDeleteTicketItemAssigned(t, getData1.Uuid)
	testDeleteTicketItemAssigned(t, getData2.Uuid)
}

func testListTicketItemAssigned(t *testing.T, arg ListTicketItemAssignedParams) {

	TicketItemAssigned, err := testQueriesTransaction.ListTicketItemAssigned(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", TicketItemAssigned)
	require.NotEmpty(t, TicketItemAssigned)

}

func RandomTicketItemAssigned() TicketItemAssignedRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	tic, _ := testQueriesTransaction.CreateTicketItem(context.Background(), RandomTicketItem())
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "olive.mercado0609@gmail.com")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "TicketStatus", 0, "Open")

	arg := TicketItemAssignedRequest{
		Uuid:         util.ToUUID("a9a97b7d-26a6-4a0e-a618-8727b83a823d"),
		TicketItemId: tic.Id,
		UserId:       usr.Id,
		AssignedById: usr.Id,
		AssignedDate: util.RandomDate(),
		Remarks:      util.RandomString(10),
		StatusId:     stat.Id,
		OtherInfo:    sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestTicketItemAssigned(
	t *testing.T,
	d1 TicketItemAssignedRequest) model.TicketItemAssigned {

	getData1, err := testQueriesTransaction.CreateTicketItemAssigned(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.AssignedById, getData1.AssignedById)
	require.Equal(t, d1.AssignedDate.Format("2006-01-02"), getData1.AssignedDate.Format("2006-01-02"))
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestTicketItemAssigned(
	t *testing.T,
	d1 TicketItemAssignedRequest) model.TicketItemAssigned {

	log.Printf("updatedD1: %v", d1)
	getData1, err := testQueriesTransaction.UpdateTicketItemAssigned(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.AssignedById, getData1.AssignedById)
	require.Equal(t, d1.AssignedDate.Format("2006-01-02"), getData1.AssignedDate.Format("2006-01-02"))
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteTicketItemAssigned(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketItemAssigned(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketItemAssignedbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
