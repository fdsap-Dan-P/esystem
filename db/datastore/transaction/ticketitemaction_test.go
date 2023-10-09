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

func TestTicketItemAction(t *testing.T) {

	// Test Data
	d1 := RandomTicketItemAction()
	d2 := RandomTicketItemAction()
	d1.Uuid = util.ToUUID("779b9166-5122-413f-85bb-5b3459015e8d")

	// Test Create
	CreatedD1 := createTestTicketItemAction(t, d1)
	CreatedD2 := createTestTicketItemAction(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketItemAction(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.ActionId, getData1.ActionId)
	require.Equal(t, d1.ActionDate.Format("2006-01-02"), getData1.ActionDate.Format("2006-01-02"))
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetTicketItemAction(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Uuid, getData2.Uuid)
	require.Equal(t, d2.TicketItemId, getData2.TicketItemId)
	require.Equal(t, d2.TrnHeadId, getData2.TrnHeadId)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.ActionId, getData2.ActionId)
	require.Equal(t, d2.ActionDate.Format("2006-01-02"), getData2.ActionDate.Format("2006-01-02"))
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetTicketItemActionbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTicketItemAction(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Uuid, updatedD1.Uuid)
	require.Equal(t, updateD2.TicketItemId, updatedD1.TicketItemId)
	require.Equal(t, updateD2.TrnHeadId, updatedD1.TrnHeadId)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.ActionId, updatedD1.ActionId)
	require.Equal(t, updateD2.ActionDate.Format("2006-01-02"), updatedD1.ActionDate.Format("2006-01-02"))
	require.Equal(t, updateD2.Remarks, updatedD1.Remarks)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListTicketItemAction(t, ListTicketItemActionParams{
		TicketItemId: updatedD1.TicketItemId,
		Limit:        5,
		Offset:       0,
	})

	// Delete Data
	testDeleteTicketItemAction(t, getData1.Uuid)
	testDeleteTicketItemAction(t, getData2.Uuid)
}

func testListTicketItemAction(t *testing.T, arg ListTicketItemActionParams) {

	ticketItemAction, err := testQueriesTransaction.ListTicketItemAction(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", ticketItemAction)
	require.NotEmpty(t, ticketItemAction)

}

func RandomTicketItemAction() TicketItemActionRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ticItem, _ := testQueriesTransaction.CreateTicketItem(context.Background(), RandomTicketItem())
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")
	act, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "actionlist", 0, "Posted")

	trnHd, _ := testQueriesTransaction.CreateTrnHead(context.Background(), RandomTrnHead())

	arg := TicketItemActionRequest{
		Uuid: util.ToUUID("a9a97b7d-26a6-4a0e-a618-8727b83a823d"),

		TicketItemId: ticItem.Id,
		TrnHeadId:    trnHd.Id,
		UserId:       usr.Id,
		ActionId:     act.Id,
		ActionDate:   util.RandomDate(),
		Remarks:      sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		OtherInfo:    sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestTicketItemAction(
	t *testing.T,
	d1 TicketItemActionRequest) model.TicketItemAction {

	getData1, err := testQueriesTransaction.CreateTicketItemAction(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.ActionId, getData1.ActionId)
	require.Equal(t, d1.ActionDate.Format("2006-01-02"), getData1.ActionDate.Format("2006-01-02"))
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestTicketItemAction(
	t *testing.T,
	d1 TicketItemActionRequest) model.TicketItemAction {

	getData1, err := testQueriesTransaction.UpdateTicketItemAction(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.TicketItemId, getData1.TicketItemId)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.ActionId, getData1.ActionId)
	require.Equal(t, d1.ActionDate.Format("2006-01-02"), getData1.ActionDate.Format("2006-01-02"))
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteTicketItemAction(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketItemAction(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketItemActionbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
