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

func TestTrnHead(t *testing.T) {

	// Test Data
	d1 := RandomTrnHead()
	d2 := RandomTrnHead()
	d2.Uuid = uuid.MustParse("1d5a0ddf-a5ce-46a9-97db-e0157d5633f0")
	d2.TrnSerial = "1d5a0ddf-a5ce-46a9-97db-e0157d5633f0"

	// Test Create
	CreatedD1 := createTestTrnHead(t, d1)
	CreatedD2 := createTestTrnHead(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTrnHead(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TrnSerial, getData1.TrnSerial)
	require.Equal(t, d1.TicketId, getData1.TicketId)
	require.Equal(t, d1.TrnDate.Format("2006-01-02"), getData1.TrnDate.Format("2006-01-02"))
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.TransactingIiid, getData1.TransactingIiid)
	require.Equal(t, d1.Orno, getData1.Orno)
	require.Equal(t, d1.Isfinal, getData1.Isfinal)
	require.Equal(t, d1.Ismanual, getData1.Ismanual)
	require.Equal(t, d1.AlternateTrn, getData1.AlternateTrn)
	require.Equal(t, d1.Reference, getData1.Reference)
	require.Equal(t, d1.Particular, getData1.Particular)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetTrnHead(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TicketId, getData2.TicketId)
	require.Equal(t, d2.TrnDate.Format("2006-01-02"), getData2.TrnDate.Format("2006-01-02"))
	require.Equal(t, d2.TrnSerial, getData2.TrnSerial)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.OfficeId, getData2.OfficeId)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.TransactingIiid, getData2.TransactingIiid)
	require.Equal(t, d2.Orno, getData2.Orno)
	require.Equal(t, d2.Isfinal, getData2.Isfinal)
	require.Equal(t, d2.Ismanual, getData2.Ismanual)
	require.Equal(t, d2.AlternateTrn, getData2.AlternateTrn)
	require.Equal(t, d2.Reference, getData2.Reference)
	require.Equal(t, d2.Particular, getData2.Particular)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTrnHead(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TicketId, updatedD1.TicketId)
	require.Equal(t, updateD2.TrnDate.Format("2006-01-02"), updatedD1.TrnDate.Format("2006-01-02"))
	require.Equal(t, updateD2.TrnSerial, updatedD1.TrnSerial)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.OfficeId, updatedD1.OfficeId)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.TransactingIiid, updatedD1.TransactingIiid)
	require.Equal(t, updateD2.Orno, updatedD1.Orno)
	require.Equal(t, updateD2.Isfinal, updatedD1.Isfinal)
	require.Equal(t, updateD2.Ismanual, updatedD1.Ismanual)
	require.Equal(t, updateD2.AlternateTrn, updatedD1.AlternateTrn)
	require.Equal(t, updateD2.Reference, updatedD1.Reference)
	require.Equal(t, updateD2.Particular, updatedD1.Particular)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListTrnHead(t, ListTrnHeadParams{
		TicketId: updatedD1.TicketId,
		Limit:    5,
		Offset:   0,
	})

	// Delete Data
	// testDeleteTrnHead(t, getData1.Id)
	testDeleteTrnHead(t, getData2.Id)
}

func testListTrnHead(t *testing.T, arg ListTrnHeadParams) {

	trnHead, err := testQueriesTransaction.ListTrnHead(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", trnHead)
	require.NotEmpty(t, trnHead)

}

func RandomTrnHead() TrnHeadRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	// tkt, _ := testQueriesTransaction.GetTicketbyUuid(context.Background(), uuid.MustParse("da970ce4-dc2f-44af-b1a8-49a987148922"))

	tkt, _ := testQueriesTransaction.CreateTicket(context.Background(), RandomTicket())

	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "TrnType", 0, "S/A Deposit")
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "olive.mercado0609@gmail.com")
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "101")
	arg := TrnHeadRequest{
		Uuid:            uuid.MustParse("2af90d74-3bee-48c5-8935-443edafb8f5a"),
		TicketId:        tkt.Id,
		TrnDate:         util.RandomDate(),
		TrnSerial:       "2af90d74-3bee-48c5-8935-443edafb8f5a",
		TypeId:          typ.Id,
		OfficeId:        ofc.Id,
		UserId:          usr.Id,
		TransactingIiid: util.SetNullInt64(ii.Id),
		Orno:            sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		Isfinal:         sql.NullBool(sql.NullBool{Bool: true, Valid: true}),
		Ismanual:        sql.NullBool(sql.NullBool{Bool: true, Valid: true}),
		AlternateTrn:    sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		Reference:       util.RandomString(10),
		Particular:      sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestTrnHead(
	t *testing.T,
	d1 TrnHeadRequest) model.TrnHead {

	getData1, err := testQueriesTransaction.CreateTrnHead(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketId, getData1.TicketId)
	require.Equal(t, d1.TrnDate.Format("2006-01-02"), getData1.TrnDate.Format("2006-01-02"))
	require.Equal(t, d1.TrnSerial, getData1.TrnSerial)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.TransactingIiid, getData1.TransactingIiid)
	require.Equal(t, d1.Orno, getData1.Orno)
	require.Equal(t, d1.Isfinal, getData1.Isfinal)
	require.Equal(t, d1.Ismanual, getData1.Ismanual)
	require.Equal(t, d1.AlternateTrn, getData1.AlternateTrn)
	require.Equal(t, d1.Reference, getData1.Reference)
	require.Equal(t, d1.Particular, getData1.Particular)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestTrnHead(
	t *testing.T,
	d1 TrnHeadRequest) model.TrnHead {

	getData1, err := testQueriesTransaction.UpdateTrnHead(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TicketId, getData1.TicketId)
	require.Equal(t, d1.TrnDate.Format("2006-01-02"), getData1.TrnDate.Format("2006-01-02"))
	require.Equal(t, d1.TrnSerial, getData1.TrnSerial)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.TransactingIiid, getData1.TransactingIiid)
	require.Equal(t, d1.Orno, getData1.Orno)
	require.Equal(t, d1.Isfinal, getData1.Isfinal)
	require.Equal(t, d1.Ismanual, getData1.Ismanual)
	require.Equal(t, d1.AlternateTrn, getData1.AlternateTrn)
	require.Equal(t, d1.Reference, getData1.Reference)
	require.Equal(t, d1.Particular, getData1.Particular)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteTrnHead(t *testing.T, id int64) {
	err := testQueriesTransaction.DeleteTrnHead(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTrnHead(context.Background(), id)
	require.Error(t, err)
	// require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
