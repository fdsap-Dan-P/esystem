package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"
	"simplebank/util"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestMutualFund(t *testing.T) {

	// Test Data
	d1 := randomMutualFund()
	d2 := randomMutualFund()
	d2.CID = 400002

	err := createTestMutualFund(t, d1)
	require.NoError(t, err)

	err = createTestMutualFund(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetMutualFund(context.Background(), d1.BrCode, d1.CID, d1.OrNo.Int64)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CID, getData1.CID)
	require.Equal(t, d1.OrNo.Int64, getData1.OrNo.Int64)
	require.Equal(t, d1.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d1.TrnType, getData1.TrnType)
	require.True(t, d1.TrnAmt.Equal(getData1.TrnAmt))
	require.Equal(t, d1.UserName, getData1.UserName)

	getData2, err2 := testQueriesDump.GetMutualFund(context.Background(), d2.BrCode, d2.CID, d2.OrNo.Int64)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CID, getData2.CID)
	require.Equal(t, d2.OrNo.Int64, getData2.OrNo.Int64)
	require.Equal(t, d2.TrnDate.Format(`2006-01-02`), getData2.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d2.TrnType, getData2.TrnType)
	require.True(t, d2.TrnAmt.Equal(getData2.TrnAmt))
	require.Equal(t, d2.UserName, getData2.UserName)

	// Update Data
	updateD2 := d2
	updateD2.OrNo.Int64 = 529000224
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestMutualFund(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetMutualFund(context.Background(), updateD2.BrCode, updateD2.CID, updateD2.OrNo.Int64)
	require.NoError(t, err1)

	require.Equal(t, updateD2.CID, getData1.CID)
	require.Equal(t, updateD2.OrNo.Int64, getData1.OrNo.Int64)
	require.Equal(t, updateD2.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, updateD2.TrnType, getData1.TrnType)
	require.True(t, updateD2.TrnAmt.Equal(getData1.TrnAmt))
	require.Equal(t, updateD2.UserName, getData1.UserName)

	testListMutualFund(t, ListMutualFundParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteMutualFund(t, d1.BrCode, d1.CID, d1.OrNo.Int64)
	testDeleteMutualFund(t, d2.BrCode, d2.CID, d2.OrNo.Int64)
}

func testListMutualFund(t *testing.T, arg ListMutualFundParams) {

	MutualFund, err := testQueriesDump.ListMutualFund(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", MutualFund)
	require.NotEmpty(t, MutualFund)

}

func randomMutualFund() model.MutualFund {

	arg := model.MutualFund{
		ModCtr:   1,
		BrCode:   "01",
		CID:      400001,
		OrNo:     sql.NullInt64{Int64: 529000224, Valid: true},
		TrnDate:  util.DateValue("1999-09-18"),
		TrnType:  sql.NullString{String: "3001", Valid: true},
		TrnAmt:   util.RandomMoney(),
		UserName: sql.NullString{String: "sa", Valid: true},
	}
	return arg
}

func createTestMutualFund(
	t *testing.T,
	req model.MutualFund) error {

	err1 := testQueriesDump.CreateMutualFund(context.Background(), req)
	// fmt.Printf("Get by createTestMutualFund%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetMutualFund(context.Background(), req.BrCode, req.CID, req.OrNo.Int64)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.CID, getData.CID)
	require.Equal(t, req.OrNo.Int64, getData.OrNo.Int64)
	require.Equal(t, req.TrnDate.Format(`2006-01-02`), getData.TrnDate.Format(`2006-01-02`))
	require.Equal(t, req.TrnType, getData.TrnType)
	require.True(t, req.TrnAmt.Equal(getData.TrnAmt))
	require.Equal(t, req.UserName, getData.UserName)

	return err2
}

func updateTestMutualFund(
	t *testing.T,
	d1 model.MutualFund) error {

	err := testQueriesDump.UpdateMutualFund(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteMutualFund(t *testing.T, brCode string, cid int64, orno int64) {
	err := testQueriesDump.DeleteMutualFund(context.Background(), brCode, cid, orno)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetMutualFund(context.Background(), brCode, cid, orno)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
