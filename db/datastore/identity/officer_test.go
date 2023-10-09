package db

import (
	"context"
	"database/sql"

	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestOfficer(t *testing.T) {

	// Test Data
	d1 := randomOfficer()
	d2 := randomOfficer()
	d2.Uuid = uuid.MustParse("347e3717-1875-4e54-beb7-953d5fa6c89b")

	// Test Create
	CreatedD1 := createTestOfficer(t, d1)
	CreatedD2 := createTestOfficer(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetOfficerbyUuid(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.Position, getData1.Position)
	require.Equal(t, d1.PeriodStart.Format("2006-01-02"), getData1.PeriodStart.Format("2006-01-02"))
	require.Equal(t, d1.PeriodEnd.Time.Format("2006-01-02"), getData1.PeriodEnd.Time.Format("2006-01-02"))
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesIdentity.GetOfficerbyUuid(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.OfficeId, getData2.OfficeId)
	require.Equal(t, d2.Position, getData2.Position)
	require.Equal(t, d2.PeriodStart.Format("2006-01-02"), getData2.PeriodStart.Format("2006-01-02"))
	require.Equal(t, d2.PeriodEnd.Time.Format("2006-01-02"), getData2.PeriodEnd.Time.Format("2006-01-02"))
	require.Equal(t, d2.StatusId, getData2.StatusId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	// Update Data
	updateD2 := d2
	updateD2.Position.String = "Test Edit"

	// log.Println(updateD2)
	updatedD1 := updateTestOfficer(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.OfficeId, updatedD1.OfficeId)
	require.Equal(t, updateD2.Position, updatedD1.Position)
	require.Equal(t, updateD2.PeriodStart.Format("2006-01-02"), updatedD1.PeriodStart.Format("2006-01-02"))
	require.Equal(t, updateD2.PeriodEnd.Time.Format("2006-01-02"), updatedD1.PeriodEnd.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.StatusId, updatedD1.StatusId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListOfficer(t, ListOfficerParams{
		OfficeId: updatedD1.OfficeId,
		Limit:    5,
		Offset:   0,
	})
	// Delete Data
	testDeleteOfficer(t, getData1.Uuid)
	testDeleteOfficer(t, getData2.Uuid)
}

func testListOfficer(t *testing.T, arg ListOfficerParams) {

	officer, err := testQueriesIdentity.ListOfficer(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", officer)
	require.NotEmpty(t, officer)

}

func randomOfficer() OfficerRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "EmployeeStatus", 0, "Active")

	// obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "OfficerType", 0, "Institution")

	arg := OfficerRequest{
		Uuid:        uuid.MustParse("d7383aa1-5ca7-459b-b0d8-4e55718b9525"),
		OfficeId:    ofc.Id,
		Position:    util.SetNullString("Position"),
		PeriodStart: util.RandomDate(),
		PeriodEnd:   util.RandomNullDate(),
		StatusId:    stat.Id,
		OtherInfo:   sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestOfficer(
	t *testing.T,
	d1 OfficerRequest) model.Officer {

	getData1, err := testQueriesIdentity.CreateOfficer(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.Position, getData1.Position)
	require.Equal(t, d1.PeriodStart.Format("2006-01-02"), getData1.PeriodStart.Format("2006-01-02"))
	require.Equal(t, d1.PeriodEnd.Time.Format("2006-01-02"), getData1.PeriodEnd.Time.Format("2006-01-02"))
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestOfficer(
	t *testing.T,
	d1 OfficerRequest) model.Officer {

	getData1, err := testQueriesIdentity.UpdateOfficer(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.Position, getData1.Position)
	require.Equal(t, d1.PeriodStart.Format("2006-01-02"), getData1.PeriodStart.Format("2006-01-02"))
	require.Equal(t, d1.PeriodEnd.Time.Format("2006-01-02"), getData1.PeriodEnd.Time.Format("2006-01-02"))
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteOfficer(t *testing.T, uuid uuid.UUID) {
	err := testQueriesIdentity.DeleteOfficer(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetOfficerbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
