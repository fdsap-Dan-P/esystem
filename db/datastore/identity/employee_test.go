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

	"github.com/stretchr/testify/require"
)

func TestEmployee(t *testing.T) {

	// Test Data
	d1 := randomEmployee()
	d2 := randomEmployee()
	d2.Uuid = util.ToUUID("f154e698-cd38-46b6-bee2-9469dd381116")

	// Test Create
	CreatedD1 := createTestEmployee(t, d1)
	CreatedD2 := createTestEmployee(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetEmployee(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.CentralId, getData1.CentralId)
	require.Equal(t, d1.EmployeeNo, getData1.EmployeeNo)
	require.Equal(t, d1.BasicPay.String(), getData1.BasicPay.String())
	require.Equal(t, d1.DateHired.Time.Format("2006-01-02"), getData1.DateHired.Time.Format("2006-01-02"))
	require.Equal(t, d1.DateRegular.Time.Format("2006-01-02"), getData1.DateRegular.Time.Format("2006-01-02"))
	require.Equal(t, d1.JobGrade, getData1.JobGrade)
	require.Equal(t, d1.JobStep, getData1.JobStep)
	require.Equal(t, d1.LevelId, getData1.LevelId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.PositionId, getData1.PositionId)
	require.Equal(t, d1.StatusCode, getData1.StatusCode)
	require.Equal(t, d1.SuperiorId, getData1.SuperiorId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesIdentity.GetEmployee(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.CentralId, getData2.CentralId)
	require.Equal(t, d2.EmployeeNo, getData2.EmployeeNo)
	require.Equal(t, d2.BasicPay.String(), getData2.BasicPay.String())
	require.Equal(t, d2.DateHired.Time.Format("2006-01-02"), getData2.DateHired.Time.Format("2006-01-02"))
	require.Equal(t, d2.DateRegular.Time.Format("2006-01-02"), getData2.DateRegular.Time.Format("2006-01-02"))
	require.Equal(t, d2.JobGrade, getData2.JobGrade)
	require.Equal(t, d2.JobStep, getData2.JobStep)
	require.Equal(t, d2.LevelId, getData2.LevelId)
	require.Equal(t, d2.OfficeId, getData2.OfficeId)
	require.Equal(t, d2.PositionId, getData2.PositionId)
	require.Equal(t, d2.StatusCode, getData2.StatusCode)
	require.Equal(t, d2.SuperiorId, getData2.SuperiorId)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesIdentity.GetEmployeebyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	getData, err = testQueriesIdentity.GetEmployeebyEmpNo(context.Background(), CreatedD1.CentralId, CreatedD1.EmployeeNo)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestEmployee(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.CentralId, updatedD1.CentralId)
	require.Equal(t, updateD2.EmployeeNo, updatedD1.EmployeeNo)
	require.Equal(t, updateD2.BasicPay.String(), updatedD1.BasicPay.String())
	require.Equal(t, updateD2.DateHired.Time.Format("2006-01-02"), updatedD1.DateHired.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.DateRegular.Time.Format("2006-01-02"), updatedD1.DateRegular.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.JobGrade, updatedD1.JobGrade)
	require.Equal(t, updateD2.JobStep, updatedD1.JobStep)
	require.Equal(t, updateD2.LevelId, updatedD1.LevelId)
	require.Equal(t, updateD2.OfficeId, updatedD1.OfficeId)
	require.Equal(t, updateD2.PositionId, updatedD1.PositionId)
	require.Equal(t, updateD2.StatusCode, updatedD1.StatusCode)
	require.Equal(t, updateD2.SuperiorId, updatedD1.SuperiorId)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListEmployee(t, ListEmployeeParams{
		OfficeId: updatedD1.OfficeId,
		Limit:    5,
		Offset:   0,
	})
	// Delete Data
	testDeleteEmployee(t, getData1.Id)
	testDeleteEmployee(t, getData2.Id)
}

func testListEmployee(t *testing.T, arg ListEmployeeParams) {

	employee, err := testQueriesIdentity.ListEmployee(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", employee)
	require.NotEmpty(t, employee)

}

func randomEmployee() EmployeeRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")

	lvl, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeelevel", 0, "Officer")
	pos, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "position", 0, "IT Officer")
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeetype", 0, "Employed")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeestatus", typ.Id, "Regular")

	fmt.Printf("status:%+v\n", stat)

	arg := EmployeeRequest{
		Uuid:        util.ToUUID("ed6a63f9-44bd-4aef-87ed-8e5538f392f8"),
		Iiid:        ii.Id,
		CentralId:   ofc.Id,
		EmployeeNo:  util.RandomString(10),
		BasicPay:    util.SetDecimal("10000"),
		DateHired:   sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		DateRegular: sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		JobGrade:    int16(util.RandomInt32(1, 100)),
		JobStep:     int16(util.RandomInt32(1, 100)),
		LevelId:     sql.NullInt64(sql.NullInt64{Int64: lvl.Id, Valid: true}),
		OfficeId:    ofc.Id,
		PositionId:  pos.Id,
		StatusCode:  stat.Code,
		SuperiorId:  sql.NullInt64(sql.NullInt64{Int64: 0, Valid: false}),
		TypeId:      sql.NullInt64(sql.NullInt64{Int64: typ.Id, Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestEmployee(
	t *testing.T,
	d1 EmployeeRequest) model.Employee {

	getData1, err := testQueriesIdentity.CreateEmployee(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.CentralId, getData1.CentralId)
	require.Equal(t, d1.EmployeeNo, getData1.EmployeeNo)
	require.Equal(t, d1.BasicPay.String(), getData1.BasicPay.String())
	require.Equal(t, d1.DateHired.Time.Format("2006-01-02"), getData1.DateHired.Time.Format("2006-01-02"))
	require.Equal(t, d1.DateRegular.Time.Format("2006-01-02"), getData1.DateRegular.Time.Format("2006-01-02"))
	require.Equal(t, d1.JobGrade, getData1.JobGrade)
	require.Equal(t, d1.JobStep, getData1.JobStep)
	require.Equal(t, d1.LevelId, getData1.LevelId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.PositionId, getData1.PositionId)
	require.Equal(t, d1.StatusCode, getData1.StatusCode)
	require.Equal(t, d1.SuperiorId, getData1.SuperiorId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestEmployee(
	t *testing.T,
	d1 EmployeeRequest) model.Employee {

	getData1, err := testQueriesIdentity.UpdateEmployee(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.CentralId, getData1.CentralId)
	require.Equal(t, d1.EmployeeNo, getData1.EmployeeNo)
	require.Equal(t, d1.BasicPay.String(), getData1.BasicPay.String())
	require.Equal(t, d1.DateHired.Time.Format("2006-01-02"), getData1.DateHired.Time.Format("2006-01-02"))
	require.Equal(t, d1.DateRegular.Time.Format("2006-01-02"), getData1.DateRegular.Time.Format("2006-01-02"))
	require.Equal(t, d1.JobGrade, getData1.JobGrade)
	require.Equal(t, d1.JobStep, getData1.JobStep)
	require.Equal(t, d1.LevelId, getData1.LevelId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.PositionId, getData1.PositionId)
	require.Equal(t, d1.StatusCode, getData1.StatusCode)
	require.Equal(t, d1.SuperiorId, getData1.SuperiorId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteEmployee(t *testing.T, id int64) {
	err := testQueriesIdentity.DeleteEmployee(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetEmployee(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
