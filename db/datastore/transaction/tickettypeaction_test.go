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

func TestTicketTypeAction(t *testing.T) {

	// Test Data
	d1 := RandomTicketTypeAction()

	d2 := RandomTicketTypeAction()
	act, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ActionList", 0, "Posted")
	d2.ActionId = act.Id
	d2.Uuid = util.ToUUID("779b9166-5122-413f-85bb-5b3459015e8d")

	// Test Create
	CreatedD1 := createTestTicketTypeAction(t, d1)
	CreatedD2 := createTestTicketTypeAction(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketTypeAction(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.ProductTicketTypeId, getData1.ProductTicketTypeId)
	require.Equal(t, d1.ActionId, getData1.ActionId)
	require.Equal(t, d1.Actiondesc, getData1.Actiondesc)
	require.Equal(t, d1.ActionLinkId, getData1.ActionLinkId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetTicketTypeAction(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Uuid, getData2.Uuid)
	require.Equal(t, d2.ProductTicketTypeId, getData2.ProductTicketTypeId)
	require.Equal(t, d2.ActionId, getData2.ActionId)
	require.Equal(t, d2.Actiondesc, getData2.Actiondesc)
	require.Equal(t, d2.ActionLinkId, getData2.ActionLinkId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetTicketTypeActionbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTicketTypeAction(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Uuid, updatedD1.Uuid)
	require.Equal(t, updateD2.ProductTicketTypeId, updatedD1.ProductTicketTypeId)
	require.Equal(t, updateD2.ActionId, updatedD1.ActionId)
	require.Equal(t, updateD2.Actiondesc, updatedD1.Actiondesc)
	require.Equal(t, updateD2.ActionLinkId, updatedD1.ActionLinkId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListTicketTypeAction(t, ListTicketTypeActionParams{
		ProductTicketTypeId: updatedD1.ProductTicketTypeId,
		Limit:               5,
		Offset:              0,
	})

	// Delete Data
	testDeleteTicketTypeAction(t, getData1.Uuid)
	testDeleteTicketTypeAction(t, getData2.Uuid)
}

func testListTicketTypeAction(t *testing.T, arg ListTicketTypeActionParams) {

	ticketTypeAction, err := testQueriesTransaction.ListTicketTypeAction(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", ticketTypeAction)
	require.NotEmpty(t, ticketTypeAction)

}

func RandomTicketTypeAction() TicketTypeActionRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	// act, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ticketTypeActionstatus", 0, "Completed")

	prodTic, _ := testQueriesTransaction.CreateProductTicketType(context.Background(), RandomProductTicketType())
	act, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ActionList", 0, "Encoded")

	log.Printf("prodTic %v", prodTic)

	arg := TicketTypeActionRequest{
		Uuid:                util.ToUUID("a9a97b7d-26a6-4a0e-a618-8727b83a823d"),
		ProductTicketTypeId: prodTic.Id,
		ActionId:            act.Id,
		Actiondesc:          "Encoded",
		// ActionLinkId:
		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestTicketTypeAction(
	t *testing.T,
	d1 TicketTypeActionRequest) model.TicketTypeAction {

	getData1, err := testQueriesTransaction.CreateTicketTypeAction(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.ProductTicketTypeId, getData1.ProductTicketTypeId)
	require.Equal(t, d1.ActionId, getData1.ActionId)
	require.Equal(t, d1.Actiondesc, getData1.Actiondesc)
	require.Equal(t, d1.ActionLinkId, getData1.ActionLinkId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestTicketTypeAction(
	t *testing.T,
	d1 TicketTypeActionRequest) model.TicketTypeAction {

	getData1, err := testQueriesTransaction.UpdateTicketTypeAction(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.ProductTicketTypeId, getData1.ProductTicketTypeId)
	require.Equal(t, d1.ActionId, getData1.ActionId)
	require.Equal(t, d1.Actiondesc, getData1.Actiondesc)
	require.Equal(t, d1.ActionLinkId, getData1.ActionLinkId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteTicketTypeAction(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketTypeAction(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketTypeActionbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
