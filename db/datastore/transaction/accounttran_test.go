package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAccountTran(t *testing.T) {

	// Test Data
	d1 := randomAccountTran()
	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")
	d1.AccountId = acc.Id

	d2 := randomAccountTran()
	acc, _ = testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000002")
	d2.AccountId = acc.Id

	// Test Create
	CreatedD1 := createTestAccountTran(t, d1)
	CreatedD2 := createTestAccountTran(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetAccountTran(context.Background(), CreatedD1.TrnHeadId, CreatedD1.Series)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.ValueDate.Format("2006-01-02"), getData1.ValueDate.Format("2006-01-02"))
	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.PassbookPosted, getData1.PassbookPosted)
	require.Equal(t, d1.TrnPrin.String(), getData1.TrnPrin.String())
	require.Equal(t, d1.TrnInt.String(), getData1.TrnInt.String())
	require.Equal(t, d1.BalPrin.String(), getData1.BalPrin.String())
	require.Equal(t, d1.BalInt.String(), getData1.BalInt.String())
	require.Equal(t, d1.Cancelled, getData1.Cancelled)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetAccountTran(context.Background(), CreatedD2.TrnHeadId, CreatedD2.Series)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TrnHeadId, getData2.TrnHeadId)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.ValueDate.Format("2006-01-02"), getData2.ValueDate.Format("2006-01-02"))
	require.Equal(t, d2.AccountId, getData2.AccountId)
	require.Equal(t, d2.Currency, getData2.Currency)
	require.Equal(t, d2.ItemId, getData2.ItemId)
	require.Equal(t, d2.PassbookPosted, getData2.PassbookPosted)
	require.Equal(t, d2.TrnPrin.String(), getData2.TrnPrin.String())
	require.Equal(t, d2.TrnInt.String(), getData2.TrnInt.String())
	require.Equal(t, d2.BalPrin.String(), getData2.BalPrin.String())
	require.Equal(t, d2.BalInt.String(), getData2.BalInt.String())
	require.Equal(t, d2.Cancelled, getData2.Cancelled)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetAccountTranbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountTran(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TrnHeadId, updatedD1.TrnHeadId)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.ValueDate.Format("2006-01-02"), updatedD1.ValueDate.Format("2006-01-02"))
	require.Equal(t, updateD2.AccountId, updatedD1.AccountId)
	require.Equal(t, updateD2.Currency, updatedD1.Currency)
	require.Equal(t, updateD2.ItemId, updatedD1.ItemId)
	require.Equal(t, updateD2.PassbookPosted, updatedD1.PassbookPosted)
	require.Equal(t, updateD2.TrnPrin.String(), updatedD1.TrnPrin.String())
	require.Equal(t, updateD2.TrnInt.String(), updatedD1.TrnInt.String())
	require.Equal(t, updateD2.BalPrin.String(), updatedD1.BalPrin.String())
	require.Equal(t, updateD2.BalInt.String(), updatedD1.BalInt.String())
	require.Equal(t, updateD2.Cancelled, updatedD1.Cancelled)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListAccountTran(t, ListAccountTranParams{
		TrnHeadId: updatedD1.TrnHeadId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteAccountTran(t, CreatedD1.Uuid)
	testDeleteAccountTran(t, CreatedD2.Uuid)
}

func testListAccountTran(t *testing.T, arg ListAccountTranParams) {

	accountTran, err := testQueriesTransaction.ListAccountTran(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountTran)
	require.NotEmpty(t, accountTran)

}

func randomAccountTran() AccountTranRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	trn, _ := testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("2af90d74-3bee-48c5-8935-443edafb8f5a"))

	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "TrnType", 0, "S/A Deposit")

	arg := AccountTranRequest{
		TrnHeadId:   trn.Id,
		Series:      util.RandomInt(1, 100),
		ValueDate:   util.RandomDate(),
		TrnTypeCode: typ.Code,
		// AccountId:      util.RandomInt(1, 100),
		Currency:       util.RandomCurrency(),
		ItemId:         sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		PassbookPosted: true,
		TrnPrin:        util.RandomMoney(),
		TrnInt:         util.RandomMoney(),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestAccountTran(
	t *testing.T,
	d1 AccountTranRequest) model.AccountTran {

	getData1, err := testQueriesTransaction.CreateAccountTran(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.ValueDate.Format("2006-01-02"), getData1.ValueDate.Format("2006-01-02"))
	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.PassbookPosted, getData1.PassbookPosted)
	require.Equal(t, d1.TrnPrin.String(), getData1.TrnPrin.String())
	require.Equal(t, d1.TrnInt.String(), getData1.TrnInt.String())
	require.Equal(t, d1.BalPrin.String(), getData1.BalPrin.String())
	require.Equal(t, d1.BalInt.String(), getData1.BalInt.String())
	require.Equal(t, d1.Cancelled, getData1.Cancelled)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestAccountTran(
	t *testing.T,
	d1 AccountTranRequest) model.AccountTran {

	getData1, err := testQueriesTransaction.UpdateAccountTran(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.ValueDate.Format("2006-01-02"), getData1.ValueDate.Format("2006-01-02"))
	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.PassbookPosted, getData1.PassbookPosted)
	require.Equal(t, d1.TrnPrin.String(), getData1.TrnPrin.String())
	require.Equal(t, d1.TrnInt.String(), getData1.TrnInt.String())
	require.Equal(t, d1.BalPrin.String(), getData1.BalPrin.String())
	require.Equal(t, d1.BalInt.String(), getData1.BalInt.String())
	require.Equal(t, d1.Cancelled, getData1.Cancelled)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAccountTran(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteAccountTran(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetAccountTranbyUuid(context.Background(), uuid)
	require.Error(t, err)
	// require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
