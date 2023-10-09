package db

import (
	"context"
	"database/sql"
	"log"
	"simplebank/util"
	"time"

	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestTrnMaster(t *testing.T) {

	// Test Data
	d1 := randomTrnMaster()
	d2 := randomTrnMaster()
	// d2.Trn = d2.Trn + int64(1)

	err := createTestTrnMaster(t, d1)
	require.NoError(t, err)

	err = createTestTrnMaster(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetTrnMaster(context.Background(), d1.TrnDate, d1.Trn)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d1.Trn, getData1.Trn)
	require.Equal(t, d1.TrnType, getData1.TrnType)
	require.Equal(t, d1.OrNo, getData1.OrNo)
	require.True(t, d1.Prin.Equal(getData1.Prin))
	require.True(t, d1.IntR.Equal(getData1.IntR))
	require.True(t, d1.WaivedInt.Equal(getData1.WaivedInt))
	require.Equal(t, d1.RefNo, getData1.RefNo)
	require.Equal(t, d1.UserName, getData1.UserName)
	require.Equal(t, d1.Particular, getData1.Particular)

	getData2, err2 := testQueriesLocal.GetTrnMaster(context.Background(), d2.TrnDate, d2.Trn)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.TrnDate.Format(`2006-01-02`), getData2.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d2.Trn, getData2.Trn)
	require.Equal(t, d2.TrnType, getData2.TrnType)
	require.Equal(t, d2.OrNo, getData2.OrNo)
	require.True(t, d2.Prin.Equal(getData2.Prin))
	require.True(t, d2.IntR.Equal(getData2.IntR))
	require.True(t, d2.WaivedInt.Equal(getData2.WaivedInt))
	require.Equal(t, d2.RefNo, getData2.RefNo)
	require.Equal(t, d2.UserName, getData2.UserName)
	require.Equal(t, d2.Particular, getData2.Particular)

	// Update Data
	updateD2 := d2
	updateD2.Prin.Add(decimal.NewFromFloat(100))

	// log.Println(updateD2)
	err3 := updateTestTrnMaster(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesLocal.GetTrnMaster(context.Background(), updateD2.TrnDate, updateD2.Trn)
	require.NoError(t, err1)

	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.Equal(t, updateD2.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, updateD2.Trn, getData1.Trn)
	require.Equal(t, updateD2.TrnType, getData1.TrnType)
	require.Equal(t, updateD2.OrNo, getData1.OrNo)
	require.True(t, updateD2.Prin.Equal(getData1.Prin))
	require.True(t, updateD2.IntR.Equal(getData1.IntR))
	require.True(t, updateD2.WaivedInt.Equal(getData1.WaivedInt))
	require.Equal(t, updateD2.RefNo, getData1.RefNo)
	require.Equal(t, updateD2.UserName, getData1.UserName)
	require.Equal(t, updateD2.Particular, getData1.Particular)

	testListTrnMaster(t, ListTrnMasterParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteTrnMaster(t, d1.TrnDate, d1.Trn)
	testDeleteTrnMaster(t, d2.TrnDate, d2.Trn)
}

func testListTrnMaster(t *testing.T, arg ListTrnMasterParams) {

	TrnMaster, err := testQueriesLocal.ListTrnMaster(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", TrnMaster)
	require.NotEmpty(t, TrnMaster)

}

func randomTrnMaster() TrnMasterRequest {

	arg := TrnMasterRequest{
		Acc:        "03E3-4002-2153782",
		TrnDate:    util.DateValue("1999-01-02"),
		Trn:        10001,
		TrnType:    sql.NullInt64{Int64: 1, Valid: true},
		OrNo:       sql.NullInt64{Int64: 105, Valid: true},
		Prin:       decimal.NewFromInt(101),
		IntR:       decimal.NewFromInt(102),
		WaivedInt:  decimal.NewFromInt(103),
		RefNo:      sql.NullString{String: "dsdff", Valid: true},
		UserName:   sql.NullString{String: "sa", Valid: true},
		Particular: sql.NullString{String: "Particular", Valid: true},
	}
	return arg
}

func createTestTrnMaster(
	t *testing.T,
	req TrnMasterRequest) error {

	err1 := testQueriesLocal.CreateTrnMaster(context.Background(), req)
	log.Printf("Get by createTestTrnMaster trn:%v", req)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetTrnMaster(context.Background(), req.TrnDate, req.Trn)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.Acc, getData.Acc)
	require.Equal(t, req.TrnDate.Format(`2006-01-02`), getData.TrnDate.Format(`2006-01-02`))
	require.Equal(t, req.Trn, getData.Trn)
	require.Equal(t, req.TrnType, getData.TrnType)
	require.Equal(t, req.OrNo, getData.OrNo)
	require.Equal(t, req.Prin.String(), getData.Prin.String())
	require.True(t, req.Prin.Equal(getData.Prin))
	require.True(t, req.IntR.Equal(getData.IntR))
	require.True(t, req.WaivedInt.Equal(getData.WaivedInt))
	require.Equal(t, req.RefNo, getData.RefNo)
	require.Equal(t, req.UserName, getData.UserName)
	require.Equal(t, req.Particular, getData.Particular)

	return err2
}

func updateTestTrnMaster(
	t *testing.T,
	d1 TrnMasterRequest) error {

	err := testQueriesLocal.UpdateTrnMaster(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteTrnMaster(t *testing.T, trnDate time.Time, trn int64) {
	err := testQueriesLocal.DeleteTrnMaster(context.Background(), trnDate, trn)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetTrnMaster(context.Background(), trnDate, trn)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
