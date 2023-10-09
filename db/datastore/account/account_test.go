package db

import (
	"context"
	"database/sql"
	"fmt"

	"testing"

	"encoding/json"
	"simplebank/model"
	"simplebank/util"

	common "simplebank/db/common"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type NullString sql.NullString

func TestAccount(t *testing.T) {

	// Test Data
	d1 := randomAccount()
	d1.Acc = "1001-0001-0000001"
	d1.AlternateAcc = util.SetNullString("1001-0001-0000001")

	rancls := randomAccountClass(t)
	prod, _ := testQueriesAccount.GetProductbyName(context.Background(), "Rental")
	rancls.ProductId = prod.Id
	cls := createTestAccountClass(t, rancls)
	d1.ClassId = cls.Id

	createD1 := AccountRequest{
		Id:               d1.Id,
		Uuid:             d1.Uuid,
		CustomerId:       d1.CustomerId,
		Acc:              d1.Acc,
		AlternateAcc:     d1.AlternateAcc,
		AccountName:      d1.AccountName,
		Balance:          d1.Balance,
		NonCurrent:       d1.NonCurrent,
		ContractDate:     d1.ContractDate,
		Credit:           d1.Credit,
		Debit:            d1.Debit,
		Isbudget:         d1.Isbudget,
		LastActivityDate: d1.LastActivityDate,
		OpenDate:         d1.OpenDate,
		PassbookLine:     d1.PassbookLine,
		PendingTrnAmt:    d1.PendingTrnAmt,
		Principal:        d1.Principal,
		ClassId:          d1.ClassId,
		AccountTypeId:    d1.AccountTypeId,
		BudgetAccountId:  d1.BudgetAccountId,
		Currency:         d1.Currency,
		OfficeId:         d1.OfficeId,
		ReferredbyId:     d1.ReferredbyId,
		StatusCode:       d1.StatusCode,
		Remarks:          d1.Remarks,
		OtherInfo:        d1.OtherInfo,
	}

	d2 := randomAccount()
	d2.ClassId = cls.Id
	d2.Acc = "1001-0001-0000002"
	d2.AlternateAcc = util.SetNullString("1001-0001-0000002")

	createD2 := AccountRequest{
		Id:               d2.Id,
		Uuid:             d2.Uuid,
		CustomerId:       d2.CustomerId,
		Acc:              d2.Acc,
		AlternateAcc:     d2.AlternateAcc,
		AccountName:      d2.AccountName,
		Balance:          d2.Balance,
		NonCurrent:       d2.NonCurrent,
		ContractDate:     d2.ContractDate,
		Credit:           d2.Credit,
		Debit:            d2.Debit,
		Isbudget:         d2.Isbudget,
		LastActivityDate: d2.LastActivityDate,
		OpenDate:         d2.OpenDate,
		PassbookLine:     d2.PassbookLine,
		PendingTrnAmt:    d2.PendingTrnAmt,
		Principal:        d2.Principal,
		ClassId:          d2.ClassId,
		AccountTypeId:    d2.AccountTypeId,
		BudgetAccountId:  d2.BudgetAccountId,
		Currency:         d2.Currency,
		OfficeId:         d2.OfficeId,
		ReferredbyId:     d2.ReferredbyId,
		StatusCode:       d2.StatusCode,
		Remarks:          d2.Remarks,
		OtherInfo:        d2.OtherInfo,
	}

	// Test Create
	CreatedD1 := CreateTestAccount(t, createD1)
	CreatedD2 := CreateTestAccount(t, createD2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccount(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.ContractDate, getData1.ContractDate)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)

	getData2, err2 := testQueriesAccount.GetAccount(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.AccountTypeId, getData2.AccountTypeId)

	getData, err := testQueriesAccount.GetAccountbyAcc(context.Background(), []string{CreatedD1.Acc})
	// log.Println("test", err, getData)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Acc, getData[0].Acc)
	require.Equal(t, CreatedD1.AccountTypeId, getData[0].AccountTypeId)

	getDataStat, err := testQueriesAccount.GetAccountStat(context.Background(), []string{CreatedD1.Acc})
	// log.Println("test", err, getData)
	require.NoError(t, err)
	require.NotEmpty(t, getDataStat)
	require.Equal(t, CreatedD1.Acc, getDataStat[CreatedD1.Acc].Acc)
	require.Equal(t, CreatedD1.AccountTypeId, getDataStat[CreatedD1.Acc].AccountTypeId)

	// fmt.Printf("%+v\n", getData)

	// getData, errtestQueriesAccount.GetAccountbyUuid(context.Background(), getData2.Uuid)
	// require.NotEmpty(t, getData)
	// require.Equal(t, d2.Title, getData.Title)
	// require.Equal(t, d2.Description, getData.Description)

	// Update Data
	updateD2 := AccountRequest{
		Id:               getData2.Id,
		Uuid:             d2.Uuid,
		CustomerId:       d2.CustomerId,
		Acc:              d2.Acc,
		AlternateAcc:     d2.AlternateAcc,
		AccountName:      d2.AccountName + "Edited",
		Balance:          d2.Balance,
		NonCurrent:       d2.NonCurrent,
		ContractDate:     d2.ContractDate,
		Credit:           d2.Credit,
		Debit:            d2.Debit,
		Isbudget:         d2.Isbudget,
		LastActivityDate: d2.LastActivityDate,
		OpenDate:         d2.OpenDate,
		PassbookLine:     d2.PassbookLine,
		PendingTrnAmt:    d2.PendingTrnAmt,
		Principal:        d2.Principal,
		ClassId:          d2.ClassId,
		AccountTypeId:    d2.AccountTypeId,
		BudgetAccountId:  d2.BudgetAccountId,
		Currency:         d2.Currency,
		OfficeId:         d2.OfficeId,
		ReferredbyId:     d2.ReferredbyId,
		StatusCode:       d2.StatusCode,
		Remarks:          d2.Remarks,
		OtherInfo:        d2.OtherInfo,
	}

	// log.Println(updateD2)
	updatedD1 := updateTestAccount(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Acc, updatedD1.Acc)
	require.Equal(t, updateD2.AccountTypeId, updatedD1.AccountTypeId)

	testListAccount(t, ListAccountParams{
		CustomerId: updatedD1.CustomerId,
		Limit:      5,
		Offset:     0,
	})
	// Delete Data
	// testDeleteAccount(t, getData1.Id)
	// testDeleteAccount(t, getData2.Id)
}

func testListAccount(t *testing.T, arg ListAccountParams) {

	account, err := testQueriesAccount.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	// log.Println(account)
	fmt.Printf("%+v\n", account)
	require.NotEmpty(t, account)

	for _, account := range account {
		require.NotEmpty(t, account)
		// require.Equal(t, lastAccount.Owner, account.Owner)
	}
}

func randomAccount() AccountRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	bal := util.RandomMoney()
	Acc, _ := uuid.NewRandom()
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "LoanStatus", 0, "Current")

	cust, _ := testQueriesCustomer.GetCustomerbyAltId(context.Background(), "E3-400004")

	accType, _ := testQueriesAccount.GetAccountTypebyName(context.Background(), "MF - Sikap")
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")

	// log.Printf("test cls: %+v\n", cls)
	arg := AccountRequest{
		Balance:       bal,
		Currency:      util.RandomCurrency(),
		CustomerId:    cust.Id,
		AccountTypeId: accType.Id,
		// ClassId:      cls.Id,
		OfficeId:     ofc.Id,
		StatusCode:   stat.Id,
		AlternateAcc: sql.NullString(sql.NullString{String: Acc.String()[1:15], Valid: true}),
		OtherInfo:    sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}

	return arg
}

func CreateTestAccount(
	t *testing.T,
	createData AccountRequest) model.Account {

	account, err := testQueriesAccount.CreateAccount(context.Background(), createData)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.True(t, createData.Balance.Equal(account.Balance))
	require.Equal(t, createData.Acc, account.Acc)
	require.Equal(t, createData.AccountTypeId, account.AccountTypeId)
	return account
}

func updateTestAccount(
	t *testing.T,
	updateData AccountRequest) model.Account {

	account, err := testQueriesAccount.UpdateAccount(context.Background(), updateData)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.True(t, updateData.Balance.Equal(account.Balance))
	require.Equal(t, updateData.Balance.String(), account.Balance.String())
	require.Equal(t, updateData.Acc, account.Acc)
	require.Equal(t, updateData.AccountTypeId, account.AccountTypeId)

	return account
}

func testDeleteAccount(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteAccount(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccount(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
