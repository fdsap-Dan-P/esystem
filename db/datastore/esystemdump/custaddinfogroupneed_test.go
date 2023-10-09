package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustAddInfoGroupNeed(t *testing.T) {

	// Test Data
	d1 := randomCustAddInfoGroupNeed()
	d2 := randomCustAddInfoGroupNeed()
	d2.InfoCode = 4

	err := createTestCustAddInfoGroupNeed(t, d1)
	require.NoError(t, err)

	err = createTestCustAddInfoGroupNeed(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetCustAddInfoGroupNeed(context.Background(), d1.BrCode, d1.InfoGroup, d1.InfoCode)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.InfoGroup, getData1.InfoGroup)
	require.Equal(t, d1.InfoCode, getData1.InfoCode)
	require.Equal(t, d1.InfoProcess, getData1.InfoProcess)

	getData2, err2 := testQueriesDump.GetCustAddInfoGroupNeed(context.Background(), d2.BrCode, d2.InfoGroup, d2.InfoCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.InfoGroup, getData2.InfoGroup)
	require.Equal(t, d2.InfoCode, getData2.InfoCode)
	require.Equal(t, d2.InfoProcess, getData2.InfoProcess)

	// Update Data
	updateD2 := d2
	updateD2.InfoProcess.String = getData2.InfoProcess.String + "Edited"
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestCustAddInfoGroupNeed(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetCustAddInfoGroupNeed(context.Background(), updateD2.BrCode, updateD2.InfoGroup, updateD2.InfoCode)
	require.NoError(t, err1)

	require.Equal(t, updateD2.InfoGroup, getData1.InfoGroup)
	require.Equal(t, updateD2.InfoCode, getData1.InfoCode)
	require.Equal(t, updateD2.InfoProcess, getData1.InfoProcess)

	testListCustAddInfoGroupNeed(t, ListCustAddInfoGroupNeedParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteCustAddInfoGroupNeed(t, d1.BrCode, d1.InfoGroup, d1.InfoCode)
	testDeleteCustAddInfoGroupNeed(t, d2.BrCode, d2.InfoGroup, d2.InfoCode)
}

func testListCustAddInfoGroupNeed(t *testing.T, arg ListCustAddInfoGroupNeedParams) {

	CustAddInfoGroupNeed, err := testQueriesDump.ListCustAddInfoGroupNeed(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", CustAddInfoGroupNeed)
	require.NotEmpty(t, CustAddInfoGroupNeed)

}

func randomCustAddInfoGroupNeed() model.CustAddInfoGroupNeed {

	arg := model.CustAddInfoGroupNeed{
		ModCtr:      1,
		BrCode:      "01",
		InfoGroup:   2,
		InfoCode:    3,
		InfoProcess: sql.NullString{String: "", Valid: true},
	}
	return arg
}

func createTestCustAddInfoGroupNeed(
	t *testing.T,
	req model.CustAddInfoGroupNeed) error {

	err1 := testQueriesDump.CreateCustAddInfoGroupNeed(context.Background(), req)
	// fmt.Printf("Get by createTestCustAddInfoGroupNeed%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetCustAddInfoGroupNeed(context.Background(), req.BrCode, req.InfoGroup, req.InfoCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.InfoGroup, getData.InfoGroup)
	require.Equal(t, req.InfoCode, getData.InfoCode)
	require.Equal(t, req.InfoProcess, getData.InfoProcess)

	return err2
}

func updateTestCustAddInfoGroupNeed(
	t *testing.T,
	d1 model.CustAddInfoGroupNeed) error {

	err := testQueriesDump.UpdateCustAddInfoGroupNeed(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteCustAddInfoGroupNeed(t *testing.T, brCode string, infoGroup int64, infoCode int64) {
	err := testQueriesDump.DeleteCustAddInfoGroupNeed(context.Background(), brCode, infoGroup, infoCode)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetCustAddInfoGroupNeed(context.Background(), brCode, infoGroup, infoCode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
