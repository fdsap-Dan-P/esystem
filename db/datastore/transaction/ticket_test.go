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

func TestNewTicket(t *testing.T) {
	tic, er := testQueriesTransaction.NewDailyTicket(context.Background(),
		NewDailyTicketRequest{
			CentralOffice: "CI",
			TicketType:    "API",
			TicketDate:    util.SetDate("2022-01-01"),
			Postedby:      "konek2CARD",
			Status:        "10",
			Remarks:       util.SetNullString("remark"),
		})

	log.Printf("tic: %v er: %v", tic, er)
	require.NoError(t, er)
	//require.Error(t, er)
}

func TestTicket(t *testing.T) {

	// Test Data
	d1 := RandomTicket()
	d2 := RandomTicket()
	d1.Uuid = util.ToUUID("779b9166-5122-413f-85bb-5b3459015e8d")

	// Test Create
	CreatedD1 := createTestTicket(t, d1)
	CreatedD2 := createTestTicket(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicket(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TicketDate.Format("2006-01-02"), getData1.TicketDate.Format("2006-01-02"))
	require.Equal(t, d1.TicketTypeId, getData1.TicketTypeId)
	require.Equal(t, d1.PostedbyId, getData1.PostedbyId)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetTicket(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TicketDate.Format("2006-01-02"), getData2.TicketDate.Format("2006-01-02"))
	require.Equal(t, d2.TicketTypeId, getData2.TicketTypeId)
	require.Equal(t, d2.PostedbyId, getData2.PostedbyId)
	require.Equal(t, d2.StatusId, getData2.StatusId)
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetTicketbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTicket(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TicketDate.Format("2006-01-02"), updatedD1.TicketDate.Format("2006-01-02"))
	require.Equal(t, updateD2.TicketTypeId, updatedD1.TicketTypeId)
	require.Equal(t, updateD2.PostedbyId, updatedD1.PostedbyId)
	require.Equal(t, updateD2.StatusId, updatedD1.StatusId)
	require.Equal(t, updateD2.Remarks, updatedD1.Remarks)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListTicket(t, ListTicketParams{
		TicketDate: updatedD1.TicketDate,
		Limit:      5,
		Offset:     0,
	})

	// Delete Data
	testDeleteTicket(t, getData1.Uuid)
	testDeleteTicket(t, getData2.Uuid)
}

func testListTicket(t *testing.T, arg ListTicketParams) {

	ticket, err := testQueriesTransaction.ListTicket(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", ticket)
	require.NotEmpty(t, ticket)

}

func RandomTicket() TicketRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "tickettype", 0, "Over the Counter")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ticketstatus", 0, "Completed")
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "0000")
	fmt.Printf("%+v\n", usr)
	arg := TicketRequest{
		Uuid:            util.ToUUID("a9a97b7d-26a6-4a0e-a618-8727b83a823d"),
		CentralOfficeId: ofc.Id,
		TicketDate:      util.RandomDate(),
		TicketTypeId:    typ.Id,
		PostedbyId:      usr.Id,
		StatusId:        stat.Id,
		Remarks:         sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestTicket(
	t *testing.T,
	d1 TicketRequest) model.Ticket {

	getData1, err := testQueriesTransaction.CreateTicket(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketDate.Format("2006-01-02"), getData1.TicketDate.Format("2006-01-02"))
	require.Equal(t, d1.TicketTypeId, getData1.TicketTypeId)
	require.Equal(t, d1.PostedbyId, getData1.PostedbyId)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestTicket(
	t *testing.T,
	d1 TicketRequest) model.Ticket {

	getData1, err := testQueriesTransaction.UpdateTicket(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketDate.Format("2006-01-02"), getData1.TicketDate.Format("2006-01-02"))
	require.Equal(t, d1.TicketTypeId, getData1.TicketTypeId)
	require.Equal(t, d1.PostedbyId, getData1.PostedbyId)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteTicket(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicket(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
