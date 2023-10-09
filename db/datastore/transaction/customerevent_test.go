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

func TestCustomerEvent(t *testing.T) {

	// Test Data
	d1 := randomCustomerEvent()
	cust, _ := testQueriesCustomer.GetCustomerbyAltId(context.Background(), "10001")
	d1.CustomerId = cust.Id
	d2 := randomCustomerEvent()
	cust, _ = testQueriesCustomer.GetCustomerbyAltId(context.Background(), "10002")
	d2.CustomerId = cust.Id

	// Test Create
	CreatedD1 := createTestCustomerEvent(t, d1)
	CreatedD2 := createTestCustomerEvent(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetCustomerEvent(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetCustomerEvent(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TrnHeadId, getData2.TrnHeadId)
	require.Equal(t, d2.CustomerId, getData2.CustomerId)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetCustomerEventbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestCustomerEvent(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TrnHeadId, updatedD1.TrnHeadId)
	require.Equal(t, updateD2.CustomerId, updatedD1.CustomerId)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.Remarks, updatedD1.Remarks)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListCustomerEvent(t, ListCustomerEventParams{
		CustomerId: updatedD1.CustomerId,
		Limit:      5,
		Offset:     0,
	})

	// Delete Data
	testDeleteCustomerEvent(t, CreatedD1.Uuid)
	testDeleteCustomerEvent(t, CreatedD2.Uuid)
}

func testListCustomerEvent(t *testing.T, arg ListCustomerEventParams) {

	customerEvent, err := testQueriesTransaction.ListCustomerEvent(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", customerEvent)
	require.NotEmpty(t, customerEvent)

}

func randomCustomerEvent() CustomerEventRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	trn, _ := testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("3793422c-eb9f-49f0-9ec6-e5cf80caac25"))
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "customereventtype", 0, "Recognize")

	arg := CustomerEventRequest{
		TrnHeadId: trn.Id,
		// CustomerId: util.RandomInt(1, 100),
		TypeId:  typ.Id,
		Remarks: util.RandomString(10),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestCustomerEvent(
	t *testing.T,
	d1 CustomerEventRequest) model.CustomerEvent {

	getData1, err := testQueriesTransaction.CreateCustomerEvent(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestCustomerEvent(
	t *testing.T,
	d1 CustomerEventRequest) model.CustomerEvent {

	getData1, err := testQueriesTransaction.UpdateCustomerEvent(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteCustomerEvent(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteCustomerEvent(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetCustomerEvent(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
