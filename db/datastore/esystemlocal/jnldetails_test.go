package db

import (
	"context"
	"database/sql"
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
	getData1, err1 := testQueriesLocal.GetJnlDetails(context.Background(), d1.Trn, d1.Acc)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.Trn, getData1.Trn)
	require.Equal(t, d1.Series, getData1.Series)
	require.True(t, d1.Debit.Decimal.Equal(getData1.Debit.Decimal))
	require.True(t, d1.Credit.Decimal.Equal(getData1.Credit.Decimal))

	getData2, err2 := testQueriesLocal.GetJnlDetails(context.Background(), d2.Trn, d2.Acc)
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

	getData1, err1 = testQueriesLocal.GetJnlDetails(context.Background(), updateD2.Trn, updateD2.Acc)
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
	testDeleteJnlDetails(t, d1.Trn, d1.Acc)
	testDeleteJnlDetails(t, d2.Trn, d2.Acc)
}

func testListJnlDetails(t *testing.T, arg ListJnlDetailsParams) {

	JnlDetails, err := testQueriesLocal.ListJnlDetails(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", JnlDetails)
	require.NotEmpty(t, JnlDetails)

}

func randomJnlDetails() JnlDetailsRequest {

	arg := JnlDetailsRequest{
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
	req JnlDetailsRequest) error {

	err1 := testQueriesLocal.CreateJnlDetails(context.Background(), req)
	// fmt.Printf("Get by createTestJnlDetails%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetJnlDetails(context.Background(), req.Trn, req.Acc)
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
	d1 JnlDetailsRequest) error {

	err := testQueriesLocal.UpdateJnlDetails(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteJnlDetails(t *testing.T, trn string, acc string) {
	err := testQueriesLocal.DeleteJnlDetails(context.Background(), trn, acc)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetJnlDetails(context.Background(), trn, acc)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
