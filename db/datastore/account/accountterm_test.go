package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"simplebank/model"
	"simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestAccountTerm(t *testing.T) {

	// Test Data
	d1 := randomAccountTerm()
	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")
	d1.AccountId = acc.Id

	d2 := randomAccountTerm()
	acc, _ = testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000002")
	d2.AccountId = acc.Id

	// Test Create
	CreatedD1 := createTestAccountTerm(t, d1)
	CreatedD2 := createTestAccountTerm(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountTerm(context.Background(), CreatedD1.AccountId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Frequency, getData1.Frequency)
	require.Equal(t, d1.N, getData1.N)
	require.Equal(t, d1.PaidN, getData1.PaidN)
	require.Equal(t, d1.FixedDue.String(), getData1.FixedDue.String())
	require.Equal(t, d1.CummulativeDue.String(), getData1.CummulativeDue.String())
	require.Equal(t, d1.DateStart.Format("2006-01-02"), getData1.DateStart.Format("2006-01-02"))
	require.Equal(t, d1.Maturity.Format("2006-01-02"), getData1.Maturity.Format("2006-01-02"))

	getData2, err2 := testQueriesAccount.GetAccountTerm(context.Background(), CreatedD2.AccountId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Frequency, getData2.Frequency)
	require.Equal(t, d2.N, getData2.N)
	require.Equal(t, d2.PaidN, getData2.PaidN)
	require.Equal(t, d2.FixedDue.String(), getData2.FixedDue.String())
	require.Equal(t, d2.CummulativeDue.String(), getData2.CummulativeDue.String())
	require.Equal(t, d2.DateStart.Format("2006-01-02"), getData2.DateStart.Format("2006-01-02"))
	require.Equal(t, d2.Maturity.Format("2006-01-02"), getData2.Maturity.Format("2006-01-02"))

	getData, err := testQueriesAccount.GetAccountTermbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.AccountId, getData.AccountId)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.AccountId = getData2.AccountId
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountTerm(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Frequency, updatedD1.Frequency)
	require.Equal(t, updateD2.N, updatedD1.N)
	require.Equal(t, updateD2.PaidN, updatedD1.PaidN)
	require.Equal(t, updateD2.FixedDue.String(), updatedD1.FixedDue.String())
	require.Equal(t, updateD2.CummulativeDue.String(), updatedD1.CummulativeDue.String())
	require.Equal(t, updateD2.DateStart.Format("2006-01-02"), updatedD1.DateStart.Format("2006-01-02"))
	require.Equal(t, updateD2.Maturity.Format("2006-01-02"), updatedD1.Maturity.Format("2006-01-02"))

	testListAccountTerm(t, ListAccountTermParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteAccountTerm(t, getData1.AccountId)
	testDeleteAccountTerm(t, getData2.AccountId)
}

func testListAccountTerm(t *testing.T, arg ListAccountTermParams) {

	accountTerm, err := testQueriesAccount.ListAccountTerm(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountTerm)
	require.NotEmpty(t, accountTerm)

}

func randomAccountTerm() AccountTermRequest {
	arg := AccountTermRequest{
		Frequency:      int16(util.RandomInt32(1, 100)),
		N:              int16(util.RandomInt32(1, 100)),
		PaidN:          int16(util.RandomInt32(1, 100)),
		FixedDue:       util.RandomMoney(),
		CummulativeDue: util.RandomMoney(),
		DateStart:      util.RandomDate(),
		Maturity:       util.RandomDate(),
	}
	return arg
}

func createTestAccountTerm(
	t *testing.T,
	d1 AccountTermRequest) model.AccountTerm {

	getData1, err := testQueriesAccount.CreateAccountTerm(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Frequency, getData1.Frequency)
	require.Equal(t, d1.N, getData1.N)
	require.Equal(t, d1.PaidN, getData1.PaidN)
	require.Equal(t, d1.FixedDue.String(), getData1.FixedDue.String())
	require.Equal(t, d1.CummulativeDue.String(), getData1.CummulativeDue.String())
	require.Equal(t, d1.DateStart.Format("2006-01-02"), getData1.DateStart.Format("2006-01-02"))
	require.Equal(t, d1.Maturity.Format("2006-01-02"), getData1.Maturity.Format("2006-01-02"))

	return getData1
}

func updateTestAccountTerm(
	t *testing.T,
	d1 AccountTermRequest) model.AccountTerm {

	getData1, err := testQueriesAccount.UpdateAccountTerm(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Frequency, getData1.Frequency)
	require.Equal(t, d1.N, getData1.N)
	require.Equal(t, d1.PaidN, getData1.PaidN)
	require.Equal(t, d1.FixedDue.String(), getData1.FixedDue.String())
	require.Equal(t, d1.CummulativeDue.String(), getData1.CummulativeDue.String())
	require.Equal(t, d1.DateStart.Format("2006-01-02"), getData1.DateStart.Format("2006-01-02"))
	require.Equal(t, d1.Maturity.Format("2006-01-02"), getData1.Maturity.Format("2006-01-02"))

	return getData1
}

func testDeleteAccountTerm(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteAccountTerm(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountTerm(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
