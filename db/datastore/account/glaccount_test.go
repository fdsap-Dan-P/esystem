package db

import (
	"context"
	"database/sql"
	"log"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestGlAccount(t *testing.T) {

	// Test Data
	d1 := randomGlAccount()
	part, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "FundSource", 0, "GSB")
	d1.PartitionId = sql.NullInt64(sql.NullInt64{Int64: part.Id, Valid: true})

	d2 := randomGlAccount()
	part, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "FundSource", 0, "CGAP")
	d2.PartitionId = sql.NullInt64(sql.NullInt64{Int64: part.Id, Valid: true})

	// Test Create
	CreatedD1 := createTestGlAccount(t, d1)
	CreatedD2 := createTestGlAccount(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetGlAccount(context.Background(), CreatedD2.Id)

	log.Printf("GetGlAccount %+v: %+v", CreatedD1, getData1)
	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, getData1.OfficeId, CreatedD2.OfficeId)
	require.Equal(t, getData1.CoaId, CreatedD2.CoaId)
	require.Equal(t, getData1.Balance.String(), CreatedD2.Balance.String())
	require.Equal(t, getData1.PendingTrnAmt.String(), CreatedD2.PendingTrnAmt.String())
	require.Equal(t, getData1.AccountTypeId, CreatedD2.AccountTypeId)
	require.Equal(t, getData1.Currency, CreatedD2.Currency)
	require.Equal(t, getData1.PartitionId, CreatedD2.PartitionId)
	require.Equal(t, getData1.Remark, CreatedD2.Remark)
	require.JSONEq(t, getData1.OtherInfo.String, CreatedD2.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetGlAccount(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.OfficeId, getData2.OfficeId)
	require.Equal(t, d2.CoaId, getData2.CoaId)
	require.Equal(t, d2.Balance.String(), getData2.Balance.String())
	require.Equal(t, d2.PendingTrnAmt.String(), getData2.PendingTrnAmt.String())
	require.Equal(t, d2.AccountTypeId, getData2.AccountTypeId)
	require.Equal(t, d2.Currency, getData2.Currency)
	require.Equal(t, d2.PartitionId, getData2.PartitionId)
	require.Equal(t, d2.Remark, getData2.Remark)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetGlAccountbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestGlAccount(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.OfficeId, updatedD1.OfficeId)
	require.Equal(t, updateD2.CoaId, updatedD1.CoaId)
	require.Equal(t, updateD2.Balance.String(), updatedD1.Balance.String())
	require.Equal(t, updateD2.PendingTrnAmt.String(), updatedD1.PendingTrnAmt.String())
	require.Equal(t, updateD2.AccountTypeId, updatedD1.AccountTypeId)
	require.Equal(t, updateD2.Currency, updatedD1.Currency)
	require.Equal(t, updateD2.PartitionId, updatedD1.PartitionId)
	require.Equal(t, updateD2.Remark, updatedD1.Remark)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListGlAccount(t, ListGlAccountParams{
		OfficeId: updatedD1.OfficeId,
		Limit:    5,
		Offset:   0,
	})

	// Delete Data
	testDeleteGlAccount(t, getData1.Id)
	testDeleteGlAccount(t, getData2.Id)
}

func testListGlAccount(t *testing.T, arg ListGlAccountParams) {

	glAccount, err := testQueriesAccount.ListGlAccount(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", glAccount)
	require.NotEmpty(t, glAccount)

}

func randomGlAccount() GlAccountRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")
	accType, _ := testQueriesAccount.GetAccountTypebyName(context.Background(), "Sikap 1")
	coa, _ := testQueriesReference.GetChartofAccountbyTitle(context.Background(), "Cash on Hand")

	arg := GlAccountRequest{
		OfficeId:      ofc.Id,
		CoaId:         sql.NullInt64(sql.NullInt64{Int64: coa.Id, Valid: true}),
		Balance:       util.RandomMoney(),
		PendingTrnAmt: util.RandomMoney(),
		AccountTypeId: sql.NullInt64(sql.NullInt64{Int64: accType.Id, Valid: true}),
		Currency:      sql.NullString(sql.NullString{String: util.RandomCurrency(), Valid: true}),
		// PartitionId:   sql.NullInt64(sql.NullInt64{Int64: part.Id, Valid: true}),
		Remark: sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestGlAccount(
	t *testing.T,
	d1 GlAccountRequest) model.GlAccount {

	getData1, err := testQueriesAccount.CreateGlAccount(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.CoaId, getData1.CoaId)
	require.Equal(t, d1.Balance.String(), getData1.Balance.String())
	require.Equal(t, d1.PendingTrnAmt.String(), getData1.PendingTrnAmt.String())
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.Equal(t, d1.PartitionId, getData1.PartitionId)
	require.Equal(t, d1.Remark, getData1.Remark)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestGlAccount(
	t *testing.T,
	d1 GlAccountRequest) model.GlAccount {

	getData1, err := testQueriesAccount.UpdateGlAccount(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.CoaId, getData1.CoaId)
	require.Equal(t, d1.Balance.String(), getData1.Balance.String())
	require.Equal(t, d1.PendingTrnAmt.String(), getData1.PendingTrnAmt.String())
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.Equal(t, d1.PartitionId, getData1.PartitionId)
	require.Equal(t, d1.Remark, getData1.Remark)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteGlAccount(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteGlAccount(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetGlAccount(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
