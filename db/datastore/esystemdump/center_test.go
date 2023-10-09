package db

import (
	"context"
	"database/sql"
	"log"
	model "simplebank/db/datastore/esystemlocal"
	"simplebank/util"
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
	getData1, err1 := testQueriesDump.GetCenter(context.Background(), d1.BrCode, d1.CenterCode)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CenterCode, getData1.CenterCode)
	require.Equal(t, d1.CenterName, getData1.CenterName)
	require.Equal(t, d1.CenterAddress, getData1.CenterAddress)
	require.Equal(t, d1.MeetingDay, getData1.MeetingDay)
	require.Equal(t, d1.Unit, getData1.Unit)
	require.Equal(t, d1.DateEstablished.Time.Format("2006-01-02"), getData1.DateEstablished.Time.Format("2006-01-02"))
	require.Equal(t, d1.AOID, getData1.AOID)

	getData2, err2 := testQueriesDump.GetCenter(context.Background(), d1.BrCode, d2.CenterCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CenterCode, getData2.CenterCode)
	require.Equal(t, d2.CenterName, getData2.CenterName)
	require.Equal(t, d2.CenterAddress, getData2.CenterAddress)
	require.Equal(t, d2.MeetingDay, getData2.MeetingDay)
	require.Equal(t, d2.Unit, getData2.Unit)
	require.Equal(t, d2.DateEstablished.Time.Format("2006-01-02"), getData2.DateEstablished.Time.Format("2006-01-02"))
	require.Equal(t, d2.AOID, getData2.AOID)

	// Update Data
	updateD2 := d2
	updateD2.CenterCode = getData2.CenterCode
	updateD2.CenterName.String = updateD2.CenterName.String + "Edited"

	// log.Println(updateD2)
	err3 := updateTestCenter(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetCenter(context.Background(), updateD2.BrCode, updateD2.CenterCode)
	require.NoError(t, err1)

	require.Equal(t, updateD2.CenterCode, getData1.CenterCode)
	require.Equal(t, updateD2.CenterName, getData1.CenterName)
	require.Equal(t, updateD2.CenterAddress, getData1.CenterAddress)
	require.Equal(t, updateD2.MeetingDay, getData1.MeetingDay)
	require.Equal(t, updateD2.Unit, getData1.Unit)
	require.Equal(t, updateD2.DateEstablished.Time.Format("2006-01-02"), getData1.DateEstablished.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.AOID, getData1.AOID)

	testListCenter(t, ListCenterParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteCenter(t, d1.BrCode, d1.CenterCode)
	testDeleteCenter(t, d2.BrCode, d2.CenterCode)
}

func testListCenter(t *testing.T, arg ListCenterParams) {

	Center, err := testQueriesDump.ListCenter(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Center)
	require.NotEmpty(t, Center)

}

func randomCenter() model.Center {

	arg := model.Center{
		ModCtr:          1,
		BrCode:          "01",
		CenterCode:      "cen",
		CenterName:      sql.NullString{String: "Center Name", Valid: true},
		CenterAddress:   sql.NullString{String: "Address", Valid: true},
		MeetingDay:      sql.NullInt64{Int64: 1, Valid: true},
		Unit:            sql.NullInt64{Int64: 501, Valid: true},
		DateEstablished: sql.NullTime{Time: time.Now(), Valid: true},
		AOID:            util.SetNullInt64(100),
	}
	return arg
}

func createTestCenter(
	t *testing.T,
	req model.Center) error {

	err1 := testQueriesDump.CreateCenter(context.Background(), req)
	// fmt.Printf("Get by createTestCenter%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetCenter(context.Background(), req.BrCode, req.CenterCode)
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
	d1 model.Center) error {

	err := testQueriesDump.UpdateCenter(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteCenter(t *testing.T, brCode string, centerCode string) {
	err := testQueriesDump.DeleteCenter(context.Background(), brCode, centerCode)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetCenter(context.Background(), brCode, centerCode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
