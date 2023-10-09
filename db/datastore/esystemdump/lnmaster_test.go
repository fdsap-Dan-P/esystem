package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"
	"time"

	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestLnMaster(t *testing.T) {

	// Test Data
	d1 := randomLnMaster()
	d2 := randomLnMaster()
	d2.Acc = d2.Acc + "-"

	err := createTestLnMaster(t, d1)
	require.NoError(t, err)

	err = createTestLnMaster(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetLnMaster(context.Background(), d1.BrCode, d1.Acc)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CID, getData1.CID)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.AcctType, getData1.AcctType)
	require.Equal(t, d1.DisbDate.Time.Format(`2006-01-02`), getData1.DisbDate.Time.Format(`2006-01-02`))
	require.True(t, d1.Principal.Decimal.Equal(getData1.Principal.Decimal))
	require.True(t, d1.Interest.Decimal.Equal(getData1.Interest.Decimal))
	require.True(t, d1.NetProceed.Decimal.Equal(getData1.NetProceed.Decimal))
	require.Equal(t, d1.Gives, getData1.Gives)
	require.Equal(t, d1.Frequency, getData1.Frequency)
	require.Equal(t, d1.AnnumDiv, getData1.AnnumDiv)
	require.True(t, d1.Prin.Decimal.Equal(getData1.Prin.Decimal))
	require.True(t, d1.IntR.Decimal.Equal(getData1.IntR.Decimal))
	require.True(t, d1.WaivedInt.Decimal.Equal(getData1.WaivedInt.Decimal))
	require.Equal(t, d1.WeeksPaid, getData1.WeeksPaid)
	require.Equal(t, d1.DoMaturity.Time.Format(`2006-01-02`), getData1.DoMaturity.Time.Format(`2006-01-02`))
	require.True(t, d1.ConIntRate.Decimal.Equal(getData1.ConIntRate.Decimal))
	require.Equal(t, d1.Status, getData1.Status)
	require.Equal(t, d1.Cycle, getData1.Cycle)
	require.Equal(t, d1.LNGrpCode, getData1.LNGrpCode)
	require.Equal(t, d1.Proff, getData1.Proff)
	require.Equal(t, d1.FundSource, getData1.FundSource)
	require.Equal(t, d1.DOSRI, getData1.DOSRI)
	require.Equal(t, d1.LnCategory, getData1.LnCategory)
	require.Equal(t, d1.OpenDate.Time.Format(`2006-01-02`), getData1.OpenDate.Time.Format(`2006-01-02`))
	require.Equal(t, d1.LastTrnDate.Time.Format(`2006-01-02`), getData1.LastTrnDate.Time.Format(`2006-01-02`))
	require.Equal(t, d1.DisbBy, getData1.DisbBy)

	getData2, err2 := testQueriesDump.GetLnMaster(context.Background(), d2.BrCode, d2.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CID, getData2.CID)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.AcctType, getData2.AcctType)
	require.Equal(t, d2.DisbDate.Time.Format(`2006-01-02`), getData2.DisbDate.Time.Format(`2006-01-02`))
	require.True(t, d2.Principal.Decimal.Equal(getData2.Principal.Decimal))
	require.True(t, d2.Interest.Decimal.Equal(getData2.Interest.Decimal))
	require.True(t, d2.NetProceed.Decimal.Equal(getData2.NetProceed.Decimal))
	require.Equal(t, d2.Gives, getData2.Gives)
	require.Equal(t, d2.Frequency, getData2.Frequency)
	require.Equal(t, d2.AnnumDiv, getData2.AnnumDiv)
	require.True(t, d2.Prin.Decimal.Equal(getData2.Prin.Decimal))
	require.True(t, d2.IntR.Decimal.Equal(getData2.IntR.Decimal))
	require.True(t, d2.WaivedInt.Decimal.Equal(getData2.WaivedInt.Decimal))
	require.Equal(t, d2.WeeksPaid, getData2.WeeksPaid)
	require.Equal(t, d2.DoMaturity.Time.Format(`2006-01-02`), getData2.DoMaturity.Time.Format(`2006-01-02`))
	require.True(t, d2.ConIntRate.Decimal.Equal(getData2.ConIntRate.Decimal))
	require.Equal(t, d2.Status, getData2.Status)
	require.Equal(t, d2.Cycle, getData2.Cycle)
	require.Equal(t, d2.LNGrpCode, getData2.LNGrpCode)
	require.Equal(t, d2.Proff, getData2.Proff)
	require.Equal(t, d2.FundSource, getData2.FundSource)
	require.Equal(t, d2.DOSRI, getData2.DOSRI)
	require.Equal(t, d2.LnCategory, getData2.LnCategory)
	require.Equal(t, d2.OpenDate.Time.Format(`2006-01-02`), getData2.OpenDate.Time.Format(`2006-01-02`))
	require.Equal(t, d2.LastTrnDate.Time.Format(`2006-01-02`), getData2.LastTrnDate.Time.Format(`2006-01-02`))
	require.Equal(t, d2.DisbBy, getData2.DisbBy)

	// Update Data
	updateD2 := d2
	updateD2.Acc = getData2.Acc
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestLnMaster(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetLnMaster(context.Background(), updateD2.BrCode, updateD2.Acc)
	require.NoError(t, err1)

	require.Equal(t, updateD2.CID, getData1.CID)
	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.Equal(t, updateD2.AcctType, getData1.AcctType)
	require.Equal(t, updateD2.DisbDate.Time.Format(`2006-01-02`), getData1.DisbDate.Time.Format(`2006-01-02`))
	require.True(t, updateD2.Principal.Decimal.Equal(getData1.Principal.Decimal))
	require.True(t, updateD2.Interest.Decimal.Equal(getData1.Interest.Decimal))
	require.True(t, updateD2.NetProceed.Decimal.Equal(getData1.NetProceed.Decimal))
	require.Equal(t, updateD2.Gives, getData1.Gives)
	require.Equal(t, updateD2.Frequency, getData1.Frequency)
	require.Equal(t, updateD2.AnnumDiv, getData1.AnnumDiv)
	require.True(t, updateD2.Prin.Decimal.Equal(getData1.Prin.Decimal))
	require.True(t, updateD2.IntR.Decimal.Equal(getData1.IntR.Decimal))
	require.True(t, updateD2.WaivedInt.Decimal.Equal(getData1.WaivedInt.Decimal))
	require.Equal(t, updateD2.WeeksPaid, getData1.WeeksPaid)
	require.Equal(t, updateD2.DoMaturity.Time.Format(`2006-01-02`), getData1.DoMaturity.Time.Format(`2006-01-02`))
	require.True(t, updateD2.ConIntRate.Decimal.Equal(getData1.ConIntRate.Decimal))
	require.Equal(t, updateD2.Status, getData1.Status)
	require.Equal(t, updateD2.Cycle, getData1.Cycle)
	require.Equal(t, updateD2.LNGrpCode, getData1.LNGrpCode)
	require.Equal(t, updateD2.Proff, getData1.Proff)
	require.Equal(t, updateD2.FundSource, getData1.FundSource)
	require.Equal(t, updateD2.DOSRI, getData1.DOSRI)
	require.Equal(t, updateD2.LnCategory, getData1.LnCategory)
	require.Equal(t, updateD2.OpenDate.Time.Format(`2006-01-02`), getData1.OpenDate.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.LastTrnDate.Time.Format(`2006-01-02`), getData1.LastTrnDate.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.DisbBy, getData1.DisbBy)

	testListLnMaster(t, ListLnMasterParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteLnMaster(t, d1.BrCode, d1.Acc)
	testDeleteLnMaster(t, d2.BrCode, d2.Acc)
}

func testListLnMaster(t *testing.T, arg ListLnMasterParams) {

	LnMaster, err := testQueriesDump.ListLnMaster(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", LnMaster)
	require.NotEmpty(t, LnMaster)

}

func randomLnMaster() model.LnMaster {

	arg := model.LnMaster{
		ModCtr:      1,
		BrCode:      "01",
		CID:         19858200,
		Acc:         "AccountNo",
		AcctType:    sql.NullInt64{Int64: 301, Valid: true},
		DisbDate:    sql.NullTime{Time: time.Now(), Valid: true},
		Principal:   decimal.NewNullDecimal(decimal.Zero),
		Interest:    decimal.NewNullDecimal(decimal.Zero),
		NetProceed:  decimal.NewNullDecimal(decimal.Zero),
		Gives:       sql.NullInt64{Int64: 100, Valid: true},
		Frequency:   sql.NullInt64{Int64: 100, Valid: true},
		AnnumDiv:    sql.NullInt64{Int64: 100, Valid: true},
		Prin:        decimal.NewNullDecimal(decimal.Zero),
		IntR:        decimal.NewNullDecimal(decimal.Zero),
		WaivedInt:   decimal.NewNullDecimal(decimal.Zero),
		WeeksPaid:   sql.NullInt64{Int64: 100, Valid: true},
		DoMaturity:  sql.NullTime{Time: time.Now(), Valid: true},
		ConIntRate:  decimal.NewNullDecimal(decimal.Zero),
		Status:      sql.NullString{String: "99", Valid: true},
		Cycle:       sql.NullInt64{Int64: 100, Valid: true},
		LNGrpCode:   sql.NullInt64{Int64: 1, Valid: true},
		Proff:       sql.NullInt64{Int64: 100, Valid: true},
		FundSource:  sql.NullString{String: "dsdff", Valid: true},
		DOSRI:       sql.NullBool{Bool: false, Valid: true},
		LnCategory:  sql.NullInt64{Int64: 1, Valid: true},
		OpenDate:    sql.NullTime{Time: time.Now(), Valid: true},
		LastTrnDate: sql.NullTime{Time: time.Now(), Valid: true},
		DisbBy:      sql.NullString{String: "user", Valid: true},
	}
	return arg
}

func createTestLnMaster(
	t *testing.T,
	req model.LnMaster) error {

	err1 := testQueriesDump.CreateLnMaster(context.Background(), req)
	// fmt.Printf("Get by createTestLnMaster%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetLnMaster(context.Background(), req.BrCode, req.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.CID, getData.CID)
	require.Equal(t, req.Acc, getData.Acc)
	require.Equal(t, req.AcctType, getData.AcctType)
	require.Equal(t, req.DisbDate.Time.Format(`2006-01-02`), getData.DisbDate.Time.Format(`2006-01-02`))
	require.True(t, req.Principal.Decimal.Equal(getData.Principal.Decimal))
	require.True(t, req.Interest.Decimal.Equal(getData.Interest.Decimal))
	require.True(t, req.NetProceed.Decimal.Equal(getData.NetProceed.Decimal))
	require.Equal(t, req.Gives, getData.Gives)
	require.Equal(t, req.Frequency, getData.Frequency)
	require.Equal(t, req.AnnumDiv, getData.AnnumDiv)
	require.True(t, req.Prin.Decimal.Equal(getData.Prin.Decimal))
	require.True(t, req.IntR.Decimal.Equal(getData.IntR.Decimal))
	require.True(t, req.WaivedInt.Decimal.Equal(getData.WaivedInt.Decimal))
	require.Equal(t, req.WeeksPaid, getData.WeeksPaid)
	require.Equal(t, req.DoMaturity.Time.Format(`2006-01-02`), getData.DoMaturity.Time.Format(`2006-01-02`))
	require.True(t, req.ConIntRate.Decimal.Equal(getData.ConIntRate.Decimal))
	require.Equal(t, req.Status, getData.Status)
	require.Equal(t, req.Cycle, getData.Cycle)
	require.Equal(t, req.LNGrpCode, getData.LNGrpCode)
	require.Equal(t, req.Proff, getData.Proff)
	require.Equal(t, req.FundSource, getData.FundSource)
	require.Equal(t, req.DOSRI, getData.DOSRI)
	require.Equal(t, req.LnCategory, getData.LnCategory)
	require.Equal(t, req.OpenDate.Time.Format(`2006-01-02`), getData.OpenDate.Time.Format(`2006-01-02`))
	require.Equal(t, req.LastTrnDate.Time.Format(`2006-01-02`), getData.LastTrnDate.Time.Format(`2006-01-02`))
	require.Equal(t, req.DisbBy, getData.DisbBy)

	return err2
}

func updateTestLnMaster(
	t *testing.T,
	d1 model.LnMaster) error {

	err := testQueriesDump.UpdateLnMaster(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteLnMaster(t *testing.T, brCode string, Acc string) {
	err := testQueriesDump.DeleteLnMaster(context.Background(), brCode, Acc)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetLnMaster(context.Background(), brCode, Acc)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
