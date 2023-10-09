package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustAddInfoList(t *testing.T) {

	// Test Data
	d1 := randomCustAddInfoList()
	d2 := randomCustAddInfoList()
	d2.InfoCode = d2.InfoCode + 1

	err := createTestCustAddInfoList(t, d1)
	require.NoError(t, err)

	err = createTestCustAddInfoList(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetCustAddInfoList(context.Background(), d1.BrCode, d1.InfoCode)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.InfoCode, getData1.InfoCode)
	require.Equal(t, d1.InfoCode, getData1.InfoCode)
	require.Equal(t, d1.InfoOrder, getData1.InfoOrder)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.InfoType, getData1.InfoType)
	require.Equal(t, d1.InfoLen, getData1.InfoLen)
	require.Equal(t, d1.InfoFormat, getData1.InfoFormat)
	require.Equal(t, d1.InputType, getData1.InputType)
	require.Equal(t, d1.InfoSource, getData1.InfoSource)

	getData2, err2 := testQueriesDump.GetCustAddInfoList(context.Background(), d2.BrCode, d2.InfoCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.InfoCode, getData2.InfoCode)
	require.Equal(t, d2.InfoCode, getData2.InfoCode)
	require.Equal(t, d2.InfoOrder, getData2.InfoOrder)
	require.Equal(t, d2.Title, getData2.Title)
	require.Equal(t, d2.InfoType, getData2.InfoType)
	require.Equal(t, d2.InfoLen, getData2.InfoLen)
	require.Equal(t, d2.InfoFormat, getData2.InfoFormat)
	require.Equal(t, d2.InputType, getData2.InputType)
	require.Equal(t, d2.InfoSource, getData2.InfoSource)

	// Update Data
	updateD2 := d2
	updateD2.InfoCode = getData2.InfoCode
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestCustAddInfoList(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetCustAddInfoList(context.Background(), updateD2.BrCode, updateD2.InfoCode)
	require.NoError(t, err1)

	require.Equal(t, updateD2.InfoCode, getData1.InfoCode)
	require.Equal(t, updateD2.InfoCode, getData1.InfoCode)
	require.Equal(t, updateD2.InfoOrder, getData1.InfoOrder)
	require.Equal(t, updateD2.Title, getData1.Title)
	require.Equal(t, updateD2.InfoType, getData1.InfoType)
	require.Equal(t, updateD2.InfoLen, getData1.InfoLen)
	require.Equal(t, updateD2.InfoFormat, getData1.InfoFormat)
	require.Equal(t, updateD2.InputType, getData1.InputType)
	require.Equal(t, updateD2.InfoSource, getData1.InfoSource)

	testListCustAddInfoList(t, ListCustAddInfoListParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteCustAddInfoList(t, d1.BrCode, d1.InfoCode)
	testDeleteCustAddInfoList(t, d2.BrCode, d2.InfoCode)
}

func testListCustAddInfoList(t *testing.T, arg ListCustAddInfoListParams) {

	CustAddInfoList, err := testQueriesDump.ListCustAddInfoList(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", CustAddInfoList)
	require.NotEmpty(t, CustAddInfoList)

}

func randomCustAddInfoList() model.CustAddInfoList {

	arg := model.CustAddInfoList{
		ModCtr:     1,
		BrCode:     "01",
		InfoCode:   24,
		InfoOrder:  "230",
		Title:      "Testing Code",
		InfoType:   "Integer",
		InfoLen:    4,
		InfoFormat: "",
		InputType:  20,
		InfoSource: "0,0 No;1,10 Yes",
	}
	return arg
}

func createTestCustAddInfoList(
	t *testing.T,
	req model.CustAddInfoList) error {

	err1 := testQueriesDump.CreateCustAddInfoList(context.Background(), req)
	// fmt.Printf("Get by createTestCustAddInfoList%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetCustAddInfoList(context.Background(), req.BrCode, req.InfoCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.InfoCode, getData.InfoCode)
	require.Equal(t, req.InfoCode, getData.InfoCode)
	require.Equal(t, req.InfoOrder, getData.InfoOrder)
	require.Equal(t, req.Title, getData.Title)
	require.Equal(t, req.InfoType, getData.InfoType)
	require.Equal(t, req.InfoLen, getData.InfoLen)
	require.Equal(t, req.InfoFormat, getData.InfoFormat)
	require.Equal(t, req.InputType, getData.InputType)
	require.Equal(t, req.InfoSource, getData.InfoSource)

	return err2
}

func updateTestCustAddInfoList(
	t *testing.T,
	d1 model.CustAddInfoList) error {

	err := testQueriesDump.UpdateCustAddInfoList(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteCustAddInfoList(t *testing.T, brCode string, InfoCode int64) {
	err := testQueriesDump.DeleteCustAddInfoList(context.Background(), brCode, InfoCode)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetCustAddInfoList(context.Background(), brCode, InfoCode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
