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

func TestLedgerDetails(t *testing.T) {

	// Test Data
	d1 := randomLedgerDetails()
	d2 := randomLedgerDetails()
	d2.Acc = "5-0-02-04-06-03-04"

	err := createTestLedgerDetails(t, d1)
	require.NoError(t, err)

	err = createTestLedgerDetails(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetLedgerDetails(context.Background(), d1.BrCode, d1.TrnDate, d1.Acc)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d1.Acc, getData1.Acc)
	require.True(t, d1.Balance.Equal(getData1.Balance))

	getData2, err2 := testQueriesDump.GetLedgerDetails(context.Background(), d2.BrCode, d2.TrnDate, d2.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TrnDate.Format(`2006-01-02`), getData2.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d2.Acc, getData2.Acc)
	require.True(t, d2.Balance.Equal(getData2.Balance))

	// Update Data
	updateD2 := d2
	updateD2.Balance = util.RandomMoney()
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestLedgerDetails(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetLedgerDetails(context.Background(), updateD2.BrCode, updateD2.TrnDate, updateD2.Acc)
	require.NoError(t, err1)

	require.Equal(t, updateD2.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.True(t, updateD2.Balance.Equal(getData1.Balance))

	testListLedgerDetails(t, ListLedgerDetailsParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteLedgerDetails(t, d1.BrCode, d1.TrnDate, d1.Acc)
	testDeleteLedgerDetails(t, d2.BrCode, d2.TrnDate, d2.Acc)
}

func testListLedgerDetails(t *testing.T, arg ListLedgerDetailsParams) {

	LedgerDetails, err := testQueriesDump.ListLedgerDetails(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", LedgerDetails)
	require.NotEmpty(t, LedgerDetails)

}

func randomLedgerDetails() model.LedgerDetails {

	arg := model.LedgerDetails{
		ModCtr:  1,
		BrCode:  "01",
		TrnDate: util.DateValue("2022-01-01"),
		Acc:     "1-1-04-06-00-00",
		Balance: util.RandomMoney(),
	}
	return arg
}

func createTestLedgerDetails(
	t *testing.T,
	req model.LedgerDetails) error {

	err1 := testQueriesDump.CreateLedgerDetails(context.Background(), req)
	// fmt.Printf("Get by createTestLedgerDetails%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetLedgerDetails(context.Background(), req.BrCode, req.TrnDate, req.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.TrnDate.Format(`2006-01-02`), getData.TrnDate.Format(`2006-01-02`))
	require.Equal(t, req.Acc, getData.Acc)
	require.True(t, req.Balance.Equal(getData.Balance))

	return err2
}

func updateTestLedgerDetails(
	t *testing.T,
	d1 model.LedgerDetails) error {

	err := testQueriesDump.UpdateLedgerDetails(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteLedgerDetails(t *testing.T, brCode string, trnDate time.Time, acc string) {
	err := testQueriesDump.DeleteLedgerDetails(context.Background(), brCode, trnDate, acc)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetLedgerDetails(context.Background(), brCode, trnDate, acc)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
