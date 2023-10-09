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

func TestEmployeeEvent(t *testing.T) {

	// Test Data
	d1 := randomEmployeeEvent()
	d2 := randomEmployeeEvent()

	// Test Create
	CreatedD1 := createTestEmployeeEvent(t, d1)
	CreatedD2 := createTestEmployeeEvent(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetEmployeeEvent(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.EmployeeId, getData1.EmployeeId)
	require.Equal(t, d1.TicketId, getData1.TicketId)
	require.Equal(t, d1.EventTypeId, getData1.EventTypeId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.PositionId, getData1.PositionId)
	require.Equal(t, d1.BasicPay.String(), getData1.BasicPay.String())
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.JobGrade, getData1.JobGrade)
	require.Equal(t, d1.JobStep, getData1.JobStep)
	require.Equal(t, d1.LevelId, getData1.LevelId)
	require.Equal(t, d1.EmployeeTypeId, getData1.EmployeeTypeId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetEmployeeEvent(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.EmployeeId, getData2.EmployeeId)
	require.Equal(t, d2.TicketId, getData2.TicketId)
	require.Equal(t, d2.EventTypeId, getData2.EventTypeId)
	require.Equal(t, d2.OfficeId, getData2.OfficeId)
	require.Equal(t, d2.PositionId, getData2.PositionId)
	require.Equal(t, d2.BasicPay.String(), getData2.BasicPay.String())
	require.Equal(t, d2.StatusId, getData2.StatusId)
	require.Equal(t, d2.JobGrade, getData2.JobGrade)
	require.Equal(t, d2.JobStep, getData2.JobStep)
	require.Equal(t, d2.LevelId, getData2.LevelId)
	require.Equal(t, d2.EmployeeTypeId, getData2.EmployeeTypeId)
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetEmployeeEventbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestEmployeeEvent(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.EmployeeId, updatedD1.EmployeeId)
	require.Equal(t, updateD2.TicketId, updatedD1.TicketId)
	require.Equal(t, updateD2.EventTypeId, updatedD1.EventTypeId)
	require.Equal(t, updateD2.OfficeId, updatedD1.OfficeId)
	require.Equal(t, updateD2.PositionId, updatedD1.PositionId)
	require.Equal(t, updateD2.BasicPay.String(), updatedD1.BasicPay.String())
	require.Equal(t, updateD2.StatusId, updatedD1.StatusId)
	require.Equal(t, updateD2.JobGrade, updatedD1.JobGrade)
	require.Equal(t, updateD2.JobStep, updatedD1.JobStep)
	require.Equal(t, updateD2.LevelId, updatedD1.LevelId)
	require.Equal(t, updateD2.EmployeeTypeId, updatedD1.EmployeeTypeId)
	require.Equal(t, updateD2.Remarks, updatedD1.Remarks)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListEmployeeEvent(t, ListEmployeeEventParams{
		EmployeeId: updatedD1.EmployeeId,
		Limit:      5,
		Offset:     0,
	})

	// Delete Data
	testDeleteEmployeeEvent(t, CreatedD1.Uuid)
	testDeleteEmployeeEvent(t, CreatedD2.Uuid)
}

func testListEmployeeEvent(t *testing.T, arg ListEmployeeEventParams) {

	employeeEvent, err := testQueriesTransaction.ListEmployeeEvent(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", employeeEvent)
	require.NotEmpty(t, employeeEvent)

}

func randomEmployeeEvent() EmployeeEventRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "0000")
	evnt, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeeevent", 0, "Regularization")
	pos, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "position", 0, "IT Officer")
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeetype", 0, "Employed")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeeeventstatus", 0, "Current")

	emp, _ := testQueriesIdentity.GetEmployeebyEmpNo(context.Background(), ofc.Id, "97-0114")
	tkt, _ := testQueriesTransaction.GetTicketbyUuid(context.Background(), uuid.MustParse("da970ce4-dc2f-44af-b1a8-49a987148922"))

	arg := EmployeeEventRequest{
		EmployeeId:     emp.Id,
		TicketId:       tkt.Id,
		EventTypeId:    evnt.Id,
		OfficeId:       ofc.Id,
		PositionId:     pos.Id,
		BasicPay:       util.RandomMoney(),
		StatusId:       stat.Id,
		JobGrade:       int16(util.RandomInt32(1, 100)),
		JobStep:        int16(util.RandomInt32(1, 100)),
		LevelId:        sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		EmployeeTypeId: sql.NullInt64(sql.NullInt64{Int64: typ.Id, Valid: true}),
		Remarks:        sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestEmployeeEvent(
	t *testing.T,
	d1 EmployeeEventRequest) model.EmployeeEvent {

	getData1, err := testQueriesTransaction.CreateEmployeeEvent(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.EmployeeId, getData1.EmployeeId)
	require.Equal(t, d1.TicketId, getData1.TicketId)
	require.Equal(t, d1.EventTypeId, getData1.EventTypeId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.PositionId, getData1.PositionId)
	require.Equal(t, d1.BasicPay.String(), getData1.BasicPay.String())
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.JobGrade, getData1.JobGrade)
	require.Equal(t, d1.JobStep, getData1.JobStep)
	require.Equal(t, d1.LevelId, getData1.LevelId)
	require.Equal(t, d1.EmployeeTypeId, getData1.EmployeeTypeId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestEmployeeEvent(
	t *testing.T,
	d1 EmployeeEventRequest) model.EmployeeEvent {

	getData1, err := testQueriesTransaction.UpdateEmployeeEvent(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.EmployeeId, getData1.EmployeeId)
	require.Equal(t, d1.TicketId, getData1.TicketId)
	require.Equal(t, d1.EventTypeId, getData1.EventTypeId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.PositionId, getData1.PositionId)
	require.Equal(t, d1.BasicPay.String(), getData1.BasicPay.String())
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.JobGrade, getData1.JobGrade)
	require.Equal(t, d1.JobStep, getData1.JobStep)
	require.Equal(t, d1.LevelId, getData1.LevelId)
	require.Equal(t, d1.EmployeeTypeId, getData1.EmployeeTypeId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteEmployeeEvent(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteEmployeeEvent(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetEmployeeEvent(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
