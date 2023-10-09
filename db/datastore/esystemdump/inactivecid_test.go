package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"
	"simplebank/util"
	"time"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestInActiveCID(t *testing.T) {

	// Test Data
	d1 := randomInActiveCID()
	d2 := randomInActiveCID()

	err := createTestInActiveCID(t, d1)
	require.NoError(t, err)

	err = createTestInActiveCID(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetInActiveCID(context.Background(), d1.BrCode, d1.CID, d1.DateStart)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.BrCode, getData1.BrCode)
	require.Equal(t, d1.CID, getData1.CID)
	require.Equal(t, d1.InActive, getData1.InActive)
	require.Equal(t, d1.DateStart.Format(`2006-01-02`), getData1.DateStart.Format(`2006-01-02`))
	require.Equal(t, d1.DateEnd.Time.Format(`2006-01-02`), getData1.DateEnd.Time.Format(`2006-01-02`))
	require.Equal(t, d1.UserId, getData1.UserId)

	getData2, err2 := testQueriesDump.GetInActiveCID(context.Background(), d2.BrCode, d2.CID, d2.DateStart)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.BrCode, getData2.BrCode)
	require.Equal(t, d2.CID, getData2.CID)
	require.Equal(t, d2.InActive, getData2.InActive)
	require.Equal(t, d2.DateStart.Format(`2006-01-02`), getData2.DateStart.Format(`2006-01-02`))
	require.Equal(t, d2.DateEnd.Time.Format(`2006-01-02`), getData2.DateEnd.Time.Format(`2006-01-02`))
	require.Equal(t, d2.UserId, getData2.UserId)

	// Update Data
	updateD2 := d2
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestInActiveCID(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetInActiveCID(context.Background(), updateD2.BrCode, updateD2.CID, updateD2.DateStart)
	require.NoError(t, err1)

	require.Equal(t, updateD2.BrCode, getData1.BrCode)
	require.Equal(t, updateD2.CID, getData1.CID)
	require.Equal(t, updateD2.InActive, getData1.InActive)
	require.Equal(t, updateD2.DateStart.Format(`2006-01-02`), getData1.DateStart.Format(`2006-01-02`))
	require.Equal(t, updateD2.DateEnd.Time.Format(`2006-01-02`), getData1.DateEnd.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.UserId, getData1.UserId)

	testListInActiveCID(t, ListInActiveCIDParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteInActiveCID(t, d1.BrCode, d1.CID, d1.DateStart)
	testDeleteInActiveCID(t, d2.BrCode, d2.CID, d2.DateStart)
}

func testListInActiveCID(t *testing.T, arg ListInActiveCIDParams) {

	InActiveCID, err := testQueriesDump.ListInActiveCID(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", InActiveCID)
	require.NotEmpty(t, InActiveCID)

}

func randomInActiveCID() model.InActiveCID {

	arg := model.InActiveCID{
		ModCtr:    1,
		BrCode:    "E3",
		CID:       int64(19858200),
		InActive:  false,
		DateStart: util.SetDate("2021-01-01"),
		DateEnd:   util.SetNullDate("2021-12-31"),
		UserId:    "sa",
	}
	return arg
}

func createTestInActiveCID(
	t *testing.T,
	req model.InActiveCID) error {

	err1 := testQueriesDump.CreateInActiveCID(context.Background(), req)
	// fmt.Printf("Get by createTestInActiveCID%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetInActiveCID(context.Background(), req.BrCode, req.CID, req.DateStart)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.BrCode, getData.BrCode)
	require.Equal(t, req.CID, getData.CID)
	require.Equal(t, req.InActive, getData.InActive)
	require.Equal(t, req.DateStart.Format(`2006-01-02`), getData.DateStart.Format(`2006-01-02`))
	require.Equal(t, req.DateEnd.Time.Format(`2006-01-02`), getData.DateEnd.Time.Format(`2006-01-02`))
	require.Equal(t, req.UserId, getData.UserId)

	return err2
}

func updateTestInActiveCID(
	t *testing.T,
	d1 model.InActiveCID) error {

	err := testQueriesDump.UpdateInActiveCID(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteInActiveCID(t *testing.T, brCode string, cid int64, dateStart time.Time) {
	err := testQueriesDump.DeleteInActiveCID(context.Background(), brCode, cid, dateStart)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetInActiveCID(context.Background(), brCode, cid, dateStart)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
