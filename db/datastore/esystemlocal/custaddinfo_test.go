package db

import (
	"context"
	"database/sql"
	"log"
	"simplebank/util"
	"time"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestCSVCustAddInfo(t *testing.T) {
	err := testQueriesLocal.CustAddInfoCSV(context.Background(), "csv/CustAddInfo.csv")
	require.NoError(t, err)
	if err != nil {
		log.Fatal(err)
	}
}

func TestCustAddInfo(t *testing.T) {

	// Test Data
	d1 := randomCustAddInfo()
	d2 := randomCustAddInfo()
	d2.InfoCode = d2.InfoCode + 1

	err := createTestCustAddInfo(t, d1)
	require.NoError(t, err)

	err = createTestCustAddInfo(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetCustAddInfo(context.Background(), d1.CID, d1.InfoDate, d1.InfoCode)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CID, getData1.CID)
	require.Equal(t, d1.InfoDate.Format(`2006-01-02`), getData1.InfoDate.Format(`2006-01-02`))
	require.Equal(t, d1.InfoCode, getData1.InfoCode)
	require.Equal(t, d1.Info, getData1.Info)
	require.Equal(t, d1.InfoValue, getData1.InfoValue)

	getData2, err2 := testQueriesLocal.GetCustAddInfo(context.Background(), d2.CID, d2.InfoDate, d2.InfoCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CID, getData2.CID)
	require.Equal(t, d2.InfoDate.Format(`2006-01-02`), getData2.InfoDate.Format(`2006-01-02`))
	require.Equal(t, d2.InfoCode, getData2.InfoCode)
	require.Equal(t, d2.Info, getData2.Info)
	require.Equal(t, d2.InfoValue, getData2.InfoValue)

	// Update Data
	updateD2 := d2
	updateD2.InfoCode = getData2.InfoCode
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestCustAddInfo(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesLocal.GetCustAddInfo(context.Background(), updateD2.CID, updateD2.InfoDate, updateD2.InfoCode)
	require.NoError(t, err1)

	require.Equal(t, updateD2.CID, getData1.CID)
	require.Equal(t, updateD2.InfoDate.Format(`2006-01-02`), getData1.InfoDate.Format(`2006-01-02`))
	require.Equal(t, updateD2.InfoCode, getData1.InfoCode)
	require.Equal(t, updateD2.Info, getData1.Info)
	require.Equal(t, updateD2.InfoValue, getData1.InfoValue)

	testListCustAddInfo(t, ListCustAddInfoParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteCustAddInfo(t, d1.CID, d1.InfoCode, d1.InfoDate)
	testDeleteCustAddInfo(t, d1.CID, d2.InfoCode, d2.InfoDate)
}

func testListCustAddInfo(t *testing.T, arg ListCustAddInfoParams) {

	CustAddInfo, err := testQueriesLocal.ListCustAddInfo(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", CustAddInfo)
	require.NotEmpty(t, CustAddInfo)

}

func randomCustAddInfo() CustAddInfoRequest {

	arg := CustAddInfoRequest{
		CID:       400001,
		InfoDate:  util.RandomDate(),
		InfoCode:  2,
		Info:      "Testing Info",
		InfoValue: 1,
	}
	return arg
}

func createTestCustAddInfo(
	t *testing.T,
	req CustAddInfoRequest) error {

	err1 := testQueriesLocal.CreateCustAddInfo(context.Background(), req)
	// fmt.Printf("Get by createTestCustAddInfo%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetCustAddInfo(context.Background(), req.CID, req.InfoDate, req.InfoCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.CID, getData.CID)
	require.Equal(t, req.InfoDate.Format(`2006-01-02`), getData.InfoDate.Format(`2006-01-02`))
	require.Equal(t, req.InfoCode, getData.InfoCode)
	require.Equal(t, req.Info, getData.Info)
	require.Equal(t, req.InfoValue, getData.InfoValue)

	return err2
}

func updateTestCustAddInfo(
	t *testing.T,
	d1 CustAddInfoRequest) error {

	err := testQueriesLocal.UpdateCustAddInfo(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteCustAddInfo(t *testing.T, cid int64, infoCode int64, infoDate time.Time) {
	err := testQueriesLocal.DeleteCustAddInfo(context.Background(), cid, infoDate, infoCode)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetCustAddInfo(context.Background(), cid, infoDate, infoCode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
