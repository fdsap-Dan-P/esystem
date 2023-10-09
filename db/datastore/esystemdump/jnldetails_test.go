package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"
	"simplebank/util"

	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestJnlDetails(t *testing.T) {

	// Test Data
	d1 := randomJnlDetails()
	d2 := randomJnlDetails()
	d2.Acc = "5-0-02-04-06-03-04"

	err := createTestJnlDetails(t, d1)
	require.NoError(t, err)

	err = createTestJnlDetails(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetJnlDetails(context.Background(), d1.BrCode, d1.Acc, d1.Trn)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.Trn, getData1.Trn)
	require.Equal(t, d1.Series, getData1.Series)
	require.True(t, d1.Debit.Decimal.Equal(getData1.Debit.Decimal))
	require.True(t, d1.Credit.Decimal.Equal(getData1.Credit.Decimal))

	getData2, err2 := testQueriesDump.GetJnlDetails(context.Background(), d2.BrCode, d2.Acc, d2.Trn)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.Trn, getData2.Trn)
	require.Equal(t, d2.Series, getData2.Series)
	require.True(t, d2.Debit.Decimal.Equal(getData2.Debit.Decimal))
	require.True(t, d2.Credit.Decimal.Equal(getData2.Credit.Decimal))

	// Update Data
	updateD2 := d2
	updateD2.Credit = decimal.NewNullDecimal(decimal.Zero)
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestJnlDetails(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetJnlDetails(context.Background(), updateD2.BrCode, updateD2.Acc, updateD2.Trn)
	require.NoError(t, err1)

	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.Equal(t, updateD2.Trn, getData1.Trn)
	require.Equal(t, updateD2.Series, getData1.Series)
	require.True(t, updateD2.Debit.Decimal.Equal(getData1.Debit.Decimal))
	require.True(t, updateD2.Credit.Decimal.Equal(getData1.Credit.Decimal))

	testListJnlDetails(t, ListJnlDetailsParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteJnlDetails(t, d1.BrCode, d1.Acc, d1.Trn)
	testDeleteJnlDetails(t, d2.BrCode, d2.Acc, d2.Trn)
}

func testListJnlDetails(t *testing.T, arg ListJnlDetailsParams) {

	JnlDetails, err := testQueriesDump.ListJnlDetails(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", JnlDetails)
	require.NotEmpty(t, JnlDetails)

}

func randomJnlDetails() model.JnlDetails {

	arg := model.JnlDetails{
		ModCtr: 1,
		BrCode: "01",
		Acc:    "1-1-04-06-00-00",
		Trn:    "JV-2014-08-000004",
		Series: util.SetNullInt64(1),
		Debit:  decimal.NewNullDecimal(decimal.Zero),
		Credit: decimal.NewNullDecimal(decimal.Zero),
	}
	return arg
}

func createTestJnlDetails(
	t *testing.T,
	req model.JnlDetails) error {

	err1 := testQueriesDump.CreateJnlDetails(context.Background(), req)
	// fmt.Printf("Get by createTestJnlDetails%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetJnlDetails(context.Background(), req.BrCode, req.Acc, req.Trn)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.Acc, getData.Acc)
	require.Equal(t, req.Trn, getData.Trn)
	require.Equal(t, req.Series, getData.Series)
	require.True(t, req.Debit.Decimal.Equal(getData.Debit.Decimal))
	require.True(t, req.Credit.Decimal.Equal(getData.Credit.Decimal))

	return err2
}

func updateTestJnlDetails(
	t *testing.T,
	d1 model.JnlDetails) error {

	err := testQueriesDump.UpdateJnlDetails(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteJnlDetails(t *testing.T, brCode string, trn string, acc string) {
	err := testQueriesDump.DeleteJnlDetails(context.Background(), brCode, acc, trn)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetJnlDetails(context.Background(), brCode, acc, trn)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
