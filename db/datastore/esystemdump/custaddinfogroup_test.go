package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustAddInfoGroup(t *testing.T) {

	// Test Data
	d1 := randomCustAddInfoGroup()
	d2 := randomCustAddInfoGroup()
	d2.InfoGroup = d2.InfoGroup + 1

	err := createTestCustAddInfoGroup(t, d1)
	require.NoError(t, err)

	err = createTestCustAddInfoGroup(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetCustAddInfoGroup(context.Background(), d1.BrCode, d1.InfoGroup)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.InfoGroup, getData1.InfoGroup)
	require.Equal(t, d1.GroupTitle, getData1.GroupTitle)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.Equal(t, d1.ReqOnEntry, getData1.ReqOnEntry)
	require.Equal(t, d1.ReqOnExit, getData1.ReqOnExit)
	require.Equal(t, d1.Link2Loan, getData1.Link2Loan)
	require.Equal(t, d1.Link2Save, getData1.Link2Save)

	getData2, err2 := testQueriesDump.GetCustAddInfoGroup(context.Background(), d2.BrCode, d2.InfoGroup)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.InfoGroup, getData2.InfoGroup)
	require.Equal(t, d2.GroupTitle, getData2.GroupTitle)
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.Equal(t, d2.ReqOnEntry, getData2.ReqOnEntry)
	require.Equal(t, d2.ReqOnExit, getData2.ReqOnExit)
	require.Equal(t, d2.Link2Loan, getData2.Link2Loan)
	require.Equal(t, d2.Link2Save, getData2.Link2Save)

	// Update Data
	updateD2 := d2
	updateD2.Remarks.String = updateD2.Remarks.String + "Edited"

	// log.Println(updateD2)
	err3 := updateTestCustAddInfoGroup(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetCustAddInfoGroup(context.Background(), updateD2.BrCode, updateD2.InfoGroup)
	require.NoError(t, err1)

	require.Equal(t, updateD2.InfoGroup, getData1.InfoGroup)
	require.Equal(t, updateD2.InfoGroup, getData1.InfoGroup)
	require.Equal(t, updateD2.GroupTitle, getData1.GroupTitle)
	require.Equal(t, updateD2.Remarks, getData1.Remarks)
	require.Equal(t, updateD2.ReqOnEntry, getData1.ReqOnEntry)
	require.Equal(t, updateD2.ReqOnExit, getData1.ReqOnExit)
	require.Equal(t, updateD2.Link2Loan, getData1.Link2Loan)
	require.Equal(t, updateD2.Link2Save, getData1.Link2Save)

	testListCustAddInfoGroup(t, ListCustAddInfoGroupParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteCustAddInfoGroup(t, d1.BrCode, d1.InfoGroup)
	testDeleteCustAddInfoGroup(t, d2.BrCode, d2.InfoGroup)
}

func testListCustAddInfoGroup(t *testing.T, arg ListCustAddInfoGroupParams) {

	CustAddInfoGroup, err := testQueriesDump.ListCustAddInfoGroup(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", CustAddInfoGroup)
	require.NotEmpty(t, CustAddInfoGroup)

}

func randomCustAddInfoGroup() model.CustAddInfoGroup {

	arg := model.CustAddInfoGroup{
		ModCtr:     1,
		BrCode:     "01",
		InfoGroup:  3,
		GroupTitle: sql.NullString{String: "Test Title", Valid: true},
		Remarks:    sql.NullString{String: "Test Remarks", Valid: true},
		ReqOnEntry: sql.NullBool{Bool: false, Valid: true},
		ReqOnExit:  sql.NullBool{Bool: false, Valid: true},
		Link2Loan:  sql.NullInt64{Int64: 0, Valid: true},
		Link2Save:  sql.NullInt64{Int64: 0, Valid: true},
	}
	return arg
}

func createTestCustAddInfoGroup(
	t *testing.T,
	req model.CustAddInfoGroup) error {

	err1 := testQueriesDump.CreateCustAddInfoGroup(context.Background(), req)
	// fmt.Printf("Get by createTestCustAddInfoGroup%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetCustAddInfoGroup(context.Background(), req.BrCode, req.InfoGroup)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.InfoGroup, getData.InfoGroup)
	require.Equal(t, req.GroupTitle, getData.GroupTitle)
	require.Equal(t, req.Remarks, getData.Remarks)
	require.Equal(t, req.ReqOnEntry, getData.ReqOnEntry)
	require.Equal(t, req.ReqOnExit, getData.ReqOnExit)
	require.Equal(t, req.Link2Loan, getData.Link2Loan)
	require.Equal(t, req.Link2Save, getData.Link2Save)

	return err2
}

func updateTestCustAddInfoGroup(
	t *testing.T,
	d1 model.CustAddInfoGroup) error {

	err := testQueriesDump.UpdateCustAddInfoGroup(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteCustAddInfoGroup(t *testing.T, brCode string, InfoGroup int64) {
	err := testQueriesDump.DeleteCustAddInfoGroup(context.Background(), brCode, InfoGroup)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetCustAddInfoGroup(context.Background(), brCode, InfoGroup)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
