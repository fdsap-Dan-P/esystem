package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestCenter(t *testing.T) {

	// Test Data
	d1 := randomCenter()
	d2 := randomCenter()
	d2.CenterCode = d2.CenterCode + "-1"

	log.Printf("%v", d1)
	err := createTestCenter(t, d1)
	require.NoError(t, err)

	err = createTestCenter(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetCenter(context.Background(), d1.CenterCode)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CenterCode, getData1.CenterCode)
	require.Equal(t, d1.CenterName, getData1.CenterName)
	require.Equal(t, d1.CenterAddress, getData1.CenterAddress)
	require.Equal(t, d1.MeetingDay, getData1.MeetingDay)
	require.Equal(t, d1.Unit, getData1.Unit)
	require.Equal(t, d1.DateEstablished.Time.Format("2006-01-02"), getData1.DateEstablished.Time.Format("2006-01-02"))

	getData2, err2 := testQueriesLocal.GetCenter(context.Background(), d2.CenterCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CenterCode, getData2.CenterCode)
	require.Equal(t, d2.CenterName, getData2.CenterName)
	require.Equal(t, d2.CenterAddress, getData2.CenterAddress)
	require.Equal(t, d2.MeetingDay, getData2.MeetingDay)
	require.Equal(t, d2.Unit, getData2.Unit)
	require.Equal(t, d2.DateEstablished.Time.Format("2006-01-02"), getData2.DateEstablished.Time.Format("2006-01-02"))

	// Update Data
	updateD2 := d2
	updateD2.CenterCode = getData2.CenterCode
	updateD2.CenterName.String = updateD2.CenterName.String + "Edited"

	// log.Println(updateD2)
	err3 := updateTestCenter(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesLocal.GetCenter(context.Background(), updateD2.CenterCode)
	require.NoError(t, err1)

	require.Equal(t, updateD2.CenterCode, getData1.CenterCode)
	require.Equal(t, updateD2.CenterName, getData1.CenterName)
	require.Equal(t, updateD2.CenterAddress, getData1.CenterAddress)
	require.Equal(t, updateD2.MeetingDay, getData1.MeetingDay)
	require.Equal(t, updateD2.Unit, getData1.Unit)
	require.Equal(t, updateD2.DateEstablished.Time.Format("2006-01-02"), getData1.DateEstablished.Time.Format("2006-01-02"))

	testListCenter(t, ListCenterParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteCenter(t, d1.CenterCode)
	testDeleteCenter(t, d2.CenterCode)
}

func testListCenter(t *testing.T, arg ListCenterParams) {

	Center, err := testQueriesLocal.ListCenter(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Center)
	require.NotEmpty(t, Center)

}

func randomCenter() CenterRequest {

	arg := CenterRequest{
		CenterCode:      "cen",
		CenterName:      sql.NullString{String: "Center Name", Valid: true},
		CenterAddress:   sql.NullString{String: "Address", Valid: true},
		MeetingDay:      sql.NullInt64{Int64: 1, Valid: true},
		Unit:            sql.NullInt64{Int64: 501, Valid: true},
		DateEstablished: sql.NullTime{Time: time.Now(), Valid: true},
		AOID:            sql.NullInt64{Int64: 16, Valid: true},
	}
	return arg
}

func createTestCenter(
	t *testing.T,
	req CenterRequest) error {

	err1 := testQueriesLocal.CreateCenter(context.Background(), req)
	// fmt.Printf("Get by createTestCenter%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetCenter(context.Background(), req.CenterCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.CenterCode, getData.CenterCode)
	require.Equal(t, req.CenterName, getData.CenterName)
	require.Equal(t, req.CenterAddress, getData.CenterAddress)
	require.Equal(t, req.MeetingDay, getData.MeetingDay)
	require.Equal(t, req.Unit, getData.Unit)
	require.Equal(t, req.DateEstablished.Time.Format("2006-01-02"), getData.DateEstablished.Time.Format("2006-01-02"))

	return err2
}

func updateTestCenter(
	t *testing.T,
	d1 CenterRequest) error {

	err := testQueriesLocal.UpdateCenter(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteCenter(t *testing.T, centerCode string) {
	err := testQueriesLocal.DeleteCenter(context.Background(), centerCode)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetCenter(context.Background(), centerCode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
