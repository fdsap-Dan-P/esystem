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

func TestAccountInterest(t *testing.T) {

	// Test Data
	d1 := randomAccountInterest()
	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")
	d1.AccountId = acc.Id
	// fmt.Printf("Get by AccountId%+v\n", acc.AlternateAcc)

	d2 := randomAccountInterest()
	acc, _ = testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000002")
	d2.AccountId = acc.Id

	// Test Create
	CreatedD1 := createTestAccountInterest(t, d1)
	CreatedD2 := createTestAccountInterest(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountInterest(context.Background(), CreatedD1.AccountId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.Interest.String(), getData1.Interest.String())
	require.Equal(t, d1.EffectiveRate, getData1.EffectiveRate)
	require.Equal(t, d1.InterestRate, getData1.InterestRate)
	require.Equal(t, d1.Credit.String(), getData1.Credit.String())
	require.Equal(t, d1.Debit.String(), getData1.Debit.String())
	require.Equal(t, d1.Accruals.String(), getData1.Accruals.String())
	require.Equal(t, d1.WaivedInt.String(), getData1.WaivedInt.String())
	require.Equal(t, d1.LastAccruedDate.Time.Format("2006-01-02"), getData1.LastAccruedDate.Time.Format("2006-01-02"))

	getData2, err2 := testQueriesAccount.GetAccountInterest(context.Background(), CreatedD2.AccountId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AccountId, getData2.AccountId)
	require.Equal(t, d2.Interest.String(), getData2.Interest.String())
	require.Equal(t, d2.EffectiveRate, getData2.EffectiveRate)
	require.Equal(t, d2.InterestRate, getData2.InterestRate)
	require.Equal(t, d2.Credit.String(), getData2.Credit.String())
	require.Equal(t, d2.Debit.String(), getData2.Debit.String())
	require.Equal(t, d2.Accruals.String(), getData2.Accruals.String())
	require.Equal(t, d2.WaivedInt.String(), getData2.WaivedInt.String())
	require.Equal(t, d2.LastAccruedDate.Time.Format("2006-01-02"), getData2.LastAccruedDate.Time.Format("2006-01-02"))

	getData, err := testQueriesAccount.GetAccountInterestbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountInterest(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.AccountId, updatedD1.AccountId)
	require.Equal(t, updateD2.Interest.String(), updatedD1.Interest.String())
	require.Equal(t, updateD2.EffectiveRate, updatedD1.EffectiveRate)
	require.Equal(t, updateD2.InterestRate, updatedD1.InterestRate)
	require.Equal(t, updateD2.Credit.String(), updatedD1.Credit.String())
	require.Equal(t, updateD2.Debit.String(), updatedD1.Debit.String())
	require.Equal(t, updateD2.Accruals.String(), updatedD1.Accruals.String())
	require.Equal(t, updateD2.WaivedInt.String(), updatedD1.WaivedInt.String())
	require.Equal(t, updateD2.LastAccruedDate.Time.Format("2006-01-02"), updatedD1.LastAccruedDate.Time.Format("2006-01-02"))

	testListAccountInterest(t, ListAccountInterestParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteAccountInterest(t, CreatedD1.AccountId)
	testDeleteAccountInterest(t, CreatedD2.AccountId)
}

func testListAccountInterest(t *testing.T, arg ListAccountInterestParams) {

	accountInterest, err := testQueriesAccount.ListAccountInterest(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountInterest)
	require.NotEmpty(t, accountInterest)

}

func randomAccountInterest() AccountInterestRequest {

	arg := AccountInterestRequest{
		// AccountId:       acc.Id,
		Interest:        util.RandomMoney(),
		EffectiveRate:   util.RandomMoney(),
		InterestRate:    util.RandomMoney(),
		Credit:          util.RandomMoney(),
		Debit:           util.RandomMoney(),
		Accruals:        util.RandomMoney(),
		WaivedInt:       util.RandomMoney(),
		LastAccruedDate: sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
	}
	return arg
}

func createTestAccountInterest(
	t *testing.T,
	d1 AccountInterestRequest) model.AccountInterest {

	getData1, err := testQueriesAccount.CreateAccountInterest(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.Interest.String(), getData1.Interest.String())
	require.Equal(t, d1.EffectiveRate, getData1.EffectiveRate)
	require.Equal(t, d1.InterestRate, getData1.InterestRate)
	require.Equal(t, d1.Credit.String(), getData1.Credit.String())
	require.Equal(t, d1.Debit.String(), getData1.Debit.String())
	require.Equal(t, d1.Accruals.String(), getData1.Accruals.String())
	require.Equal(t, d1.WaivedInt.String(), getData1.WaivedInt.String())
	require.Equal(t, d1.LastAccruedDate.Time.Format("2006-01-02"), getData1.LastAccruedDate.Time.Format("2006-01-02"))

	return getData1
}

func updateTestAccountInterest(
	t *testing.T,
	d1 AccountInterestRequest) model.AccountInterest {

	getData1, err := testQueriesAccount.UpdateAccountInterest(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.Interest.String(), getData1.Interest.String())
	require.Equal(t, d1.EffectiveRate, getData1.EffectiveRate)
	require.Equal(t, d1.InterestRate, getData1.InterestRate)
	require.Equal(t, d1.Credit.String(), getData1.Credit.String())
	require.Equal(t, d1.Debit.String(), getData1.Debit.String())
	require.Equal(t, d1.Accruals.String(), getData1.Accruals.String())
	require.Equal(t, d1.WaivedInt.String(), getData1.WaivedInt.String())
	require.Equal(t, d1.LastAccruedDate.Time.Format("2006-01-02"), getData1.LastAccruedDate.Time.Format("2006-01-02"))

	return getData1
}

func testDeleteAccountInterest(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteAccountInterest(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountInterest(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
