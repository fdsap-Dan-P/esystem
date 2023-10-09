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

func TestTicketTypeStatus(t *testing.T) {

	// Test Data
	d1 := RandomTicketTypeStatus(t)
	d2 := RandomTicketTypeStatus(t)
	d2.Uuid = util.ToUUID("b165be7d-e2e1-4b3a-8931-bf660579c09a")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ticketstatus", 0, "Pending")
	d2.StatusId = stat.Id

	// Test Create
	CreatedD1 := createTestTicketTypeStatus(t, d1)
	CreatedD2 := createTestTicketTypeStatus(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTicketTypeStatus(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.ProductTicketTypeId, getData1.ProductTicketTypeId)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.TicketTypeActionArray, getData1.TicketTypeActionArray)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetTicketTypeStatus(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Uuid, getData2.Uuid)
	require.Equal(t, d2.ProductTicketTypeId, getData2.ProductTicketTypeId)
	require.Equal(t, d2.StatusId, getData2.StatusId)
	require.Equal(t, d2.TicketTypeActionArray, getData2.TicketTypeActionArray)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetTicketTypeStatusbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	log.Println(getData2)
	log.Println(updateD2)
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	updatedD1 := updateTestTicketTypeStatus(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Uuid, updatedD1.Uuid)
	require.Equal(t, updateD2.ProductTicketTypeId, updatedD1.ProductTicketTypeId)
	require.Equal(t, updateD2.StatusId, updatedD1.StatusId)
	require.Equal(t, updateD2.TicketTypeActionArray, updatedD1.TicketTypeActionArray)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListTicketTypeStatus(t, ListTicketTypeStatusParams{
		ProductTicketTypeId: updatedD1.ProductTicketTypeId,
		Limit:               5,
		Offset:              0,
	})

	// Delete Data
	testDeleteTicketTypeStatus(t, getData1.Uuid)
	testDeleteTicketTypeStatus(t, getData2.Uuid)
}

func testListTicketTypeStatus(t *testing.T, arg ListTicketTypeStatusParams) {

	TicketTypeStatus, err := testQueriesTransaction.ListTicketTypeStatus(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", TicketTypeStatus)
	require.NotEmpty(t, TicketTypeStatus)

}

func RandomTicketTypeStatus(t *testing.T) TicketTypeStatusRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ticketstatus", 0, "Completed")
	prodTic, _ := testQueriesTransaction.CreateProductTicketType(context.Background(), RandomProductTicketType())

	arg := TicketTypeStatusRequest{
		Uuid:                  util.ToUUID("38c6a971-463f-48fc-8a2e-cb7000c892ba"),
		ProductTicketTypeId:   prodTic.Id,
		StatusId:              stat.Id,
		TicketTypeActionArray: []int64{1, 1},
		OtherInfo:             sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestTicketTypeStatus(
	t *testing.T,
	d1 TicketTypeStatusRequest) model.TicketTypeStatus {

	getData1, err := testQueriesTransaction.CreateTicketTypeStatus(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.ProductTicketTypeId, getData1.ProductTicketTypeId)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.TicketTypeActionArray, getData1.TicketTypeActionArray)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestTicketTypeStatus(
	t *testing.T,
	d1 TicketTypeStatusRequest) model.TicketTypeStatus {

	getData1, err := testQueriesTransaction.UpdateTicketTypeStatus(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.ProductTicketTypeId, getData1.ProductTicketTypeId)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.TicketTypeActionArray, getData1.TicketTypeActionArray)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteTicketTypeStatus(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTicketTypeStatus(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTicketTypeStatusbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
