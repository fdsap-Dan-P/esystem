package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"
	"simplebank/util"
	"time"

	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestSaTrnMaster(t *testing.T) {

	// Test Data
	d1 := randomSaTrnMaster()
	d2 := randomSaTrnMaster()
	trn, _ := testQueriesDump.CodeIncrement(context.Background(), "saTrnMaster", `1999-01-02`)
	d2.Trn = trn

	err := createTestSaTrnMaster(t, d1)
	require.NoError(t, err)

	err = createTestSaTrnMaster(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetSaTrnMaster(context.Background(), d1.BrCode, d1.TrnDate, d1.Trn)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d1.Trn, getData1.Trn)
	require.Equal(t, d1.TrnType, getData1.TrnType)
	require.Equal(t, d1.OrNo, getData1.OrNo)
	require.True(t, d1.TrnAmt.Decimal.Equal(getData1.TrnAmt.Decimal))
	require.Equal(t, d1.RefNo, getData1.RefNo)
	require.Equal(t, d1.Particular, getData1.Particular)
	require.Equal(t, d1.PendApprove, getData1.PendApprove)

	getData2, err2 := testQueriesDump.GetSaTrnMaster(context.Background(), d2.BrCode, d2.TrnDate, d2.Trn)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.TrnDate.Format(`2006-01-02`), getData2.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d2.Trn, getData2.Trn)
	require.Equal(t, d2.TrnType, getData2.TrnType)
	require.Equal(t, d2.OrNo, getData2.OrNo)
	require.True(t, d2.TrnAmt.Decimal.Equal(getData2.TrnAmt.Decimal))
	require.Equal(t, d2.RefNo, getData2.RefNo)
	require.Equal(t, d2.Particular, getData2.Particular)
	require.Equal(t, d2.PendApprove, getData2.PendApprove)

	// Update Data
	updateD2 := d2
	// updateD2.Trn = getData2.Trn + 1
	updateD2.Particular = updateD2.Particular + "Edited"

	// log.Println(updateD2)
	err3 := updateTestSaTrnMaster(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetSaTrnMaster(context.Background(), updateD2.BrCode, updateD2.TrnDate, updateD2.Trn)
	require.NoError(t, err1)

	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.Equal(t, updateD2.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, updateD2.Trn, getData1.Trn)
	require.Equal(t, updateD2.TrnType, getData1.TrnType)
	require.Equal(t, updateD2.OrNo, getData1.OrNo)
	require.True(t, updateD2.TrnAmt.Decimal.Equal(getData1.TrnAmt.Decimal))
	require.Equal(t, updateD2.RefNo, getData1.RefNo)
	require.Equal(t, updateD2.Particular, getData1.Particular)
	require.Equal(t, updateD2.PendApprove, getData1.PendApprove)

	testListSaTrnMaster(t, ListSaTrnMasterParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteSaTrnMaster(t, d1.BrCode, d1.TrnDate, d1.Trn)
	testDeleteSaTrnMaster(t, d1.BrCode, d1.TrnDate, d2.Trn)
}

func testListSaTrnMaster(t *testing.T, arg ListSaTrnMasterParams) {

	SaTrnMaster, err := testQueriesDump.ListSaTrnMaster(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", SaTrnMaster)
	require.NotEmpty(t, SaTrnMaster)

}

func randomSaTrnMaster() model.SaTrnMaster {

	trn, _ := testQueriesDump.CodeIncrement(context.Background(), "saTrnMaster", `1999-01-02`)

	arg := model.SaTrnMaster{
		ModCtr:      1,
		BrCode:      "01",
		Acc:         "01C4-1012-06495651",
		TrnDate:     util.DateValue("1999-01-02"),
		Trn:         trn,
		TrnType:     sql.NullInt64{Int64: 1, Valid: true},
		OrNo:        sql.NullInt64{Int64: 100, Valid: true},
		TrnAmt:      decimal.NewNullDecimal(decimal.Zero),
		RefNo:       sql.NullString{String: "dsdff", Valid: true},
		Particular:  "dsdff",
		PendApprove: "A",
	}
	return arg
}

func createTestSaTrnMaster(
	t *testing.T,
	req model.SaTrnMaster) error {

	err1 := testQueriesDump.CreateSaTrnMaster(context.Background(), req)
	// fmt.Printf("Get by createTestSaTrnMaster%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetSaTrnMaster(context.Background(), req.BrCode, req.TrnDate, req.Trn)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.Acc, getData.Acc)
	require.Equal(t, req.TrnDate.Format(`2006-01-02`), getData.TrnDate.Format(`2006-01-02`))
	require.Equal(t, req.Trn, getData.Trn)
	require.Equal(t, req.TrnType, getData.TrnType)
	require.Equal(t, req.OrNo, getData.OrNo)
	require.True(t, req.TrnAmt.Decimal.Equal(getData.TrnAmt.Decimal))
	require.Equal(t, req.RefNo, getData.RefNo)
	require.Equal(t, req.Particular, getData.Particular)
	require.Equal(t, req.PendApprove, getData.PendApprove)

	return err2
}

func updateTestSaTrnMaster(
	t *testing.T,
	d1 model.SaTrnMaster) error {

	err := testQueriesDump.UpdateSaTrnMaster(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteSaTrnMaster(t *testing.T, brCode string, trnDate time.Time, trn int64) {
	err := testQueriesDump.DeleteSaTrnMaster(context.Background(), brCode, trnDate, trn)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetSaTrnMaster(context.Background(), brCode, trnDate, trn)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
