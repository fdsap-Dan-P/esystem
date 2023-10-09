package db

import (
	"context"
	"database/sql"
	"time"

	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestSaMaster(t *testing.T) {

	// Test Data
	d1 := randomSaMaster()
	d2 := randomSaMaster()
	d2.Acc = d2.Acc + "1"

	err := createTestSaMaster(t, d1)
	require.NoError(t, err)

	err = createTestSaMaster(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetSaMaster(context.Background(), d1.Acc)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.CID, getData1.CID)
	require.Equal(t, d1.Type, getData1.Type)
	require.True(t, d1.Balance.Decimal.Equal(getData1.Balance.Decimal))
	require.Equal(t, d1.DoLastTrn.Time.Format(`2006-01-02`), getData1.DoLastTrn.Time.Format(`2006-01-02`))
	require.Equal(t, d1.DoStatus.Time.Format(`2006-01-02`), getData1.DoStatus.Time.Format(`2006-01-02`))
	require.Equal(t, d1.Dopen.Time.Format(`2006-01-02`), getData1.Dopen.Time.Format(`2006-01-02`))
	require.Equal(t, d1.DoMaturity.Time.Format(`2006-01-02`), getData1.DoMaturity.Time.Format(`2006-01-02`))
	require.Equal(t, d1.Status, getData1.Status)

	getData2, err2 := testQueriesLocal.GetSaMaster(context.Background(), d2.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.CID, getData2.CID)
	require.Equal(t, d2.Type, getData2.Type)
	require.True(t, d2.Balance.Decimal.Equal(getData2.Balance.Decimal))
	require.Equal(t, d2.DoLastTrn.Time.Format(`2006-01-02`), getData2.DoLastTrn.Time.Format(`2006-01-02`))
	require.Equal(t, d2.DoStatus.Time.Format(`2006-01-02`), getData2.DoStatus.Time.Format(`2006-01-02`))
	require.Equal(t, d2.Dopen.Time.Format(`2006-01-02`), getData2.Dopen.Time.Format(`2006-01-02`))
	require.Equal(t, d2.DoMaturity.Time.Format(`2006-01-02`), getData2.DoMaturity.Time.Format(`2006-01-02`))
	require.Equal(t, d2.Status, getData2.Status)

	// Update Data
	updateD2 := d2
	updateD2.Acc = getData2.Acc
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestSaMaster(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesLocal.GetSaMaster(context.Background(), updateD2.Acc)
	require.NoError(t, err1)

	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.Equal(t, updateD2.CID, getData1.CID)
	require.Equal(t, updateD2.Type, getData1.Type)
	require.True(t, updateD2.Balance.Decimal.Equal(getData1.Balance.Decimal))
	require.Equal(t, updateD2.DoLastTrn.Time.Format(`2006-01-02`), getData1.DoLastTrn.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.DoStatus.Time.Format(`2006-01-02`), getData1.DoStatus.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.Dopen.Time.Format(`2006-01-02`), getData1.Dopen.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.DoMaturity.Time.Format(`2006-01-02`), getData1.DoMaturity.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.Status, getData1.Status)

	testListSaMaster(t, ListSaMasterParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteSaMaster(t, d1.Acc)
	testDeleteSaMaster(t, d2.Acc)
}

func testListSaMaster(t *testing.T, arg ListSaMasterParams) {

	SaMaster, err := testQueriesLocal.ListSaMaster(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", SaMaster)
	require.NotEmpty(t, SaMaster)

}

func randomSaMaster() SaMasterRequest {

	arg := SaMasterRequest{
		Acc:        "Acc Num",
		CID:        19858200,
		Type:       60,
		Balance:    decimal.NewNullDecimal(decimal.Zero),
		DoLastTrn:  sql.NullTime{Time: time.Now(), Valid: true},
		DoStatus:   sql.NullTime{Time: time.Now(), Valid: true},
		Dopen:      sql.NullTime{Time: time.Now(), Valid: true},
		DoMaturity: sql.NullTime{Time: time.Now(), Valid: true},
		Status:     sql.NullString{String: "dsdff", Valid: true},
	}
	return arg
}

func createTestSaMaster(
	t *testing.T,
	req SaMasterRequest) error {

	err1 := testQueriesLocal.CreateSaMaster(context.Background(), req)
	// fmt.Printf("Get by createTestSaMaster%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetSaMaster(context.Background(), req.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.Acc, getData.Acc)
	require.Equal(t, req.CID, getData.CID)
	require.Equal(t, req.Type, getData.Type)
	require.True(t, req.Balance.Decimal.Equal(getData.Balance.Decimal))
	require.Equal(t, req.DoLastTrn.Time.Format(`2006-01-02`), getData.DoLastTrn.Time.Format(`2006-01-02`))
	require.Equal(t, req.DoStatus.Time.Format(`2006-01-02`), getData.DoStatus.Time.Format(`2006-01-02`))
	require.Equal(t, req.Dopen.Time.Format(`2006-01-02`), getData.Dopen.Time.Format(`2006-01-02`))
	require.Equal(t, req.DoMaturity.Time.Format(`2006-01-02`), getData.DoMaturity.Time.Format(`2006-01-02`))
	require.Equal(t, req.Status, getData.Status)

	return err2
}

func updateTestSaMaster(
	t *testing.T,
	d1 SaMasterRequest) error {

	err := testQueriesLocal.UpdateSaMaster(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteSaMaster(t *testing.T, Acc string) {
	err := testQueriesLocal.DeleteSaMaster(context.Background(), Acc)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetSaMaster(context.Background(), Acc)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
