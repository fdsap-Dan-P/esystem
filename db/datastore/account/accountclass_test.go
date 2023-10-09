package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"

	"github.com/stretchr/testify/require"
)

func TestAccountClass(t *testing.T) {

	// Test Data
	product, _ := testQueriesAccount.GetProductbyName(context.Background(), "Rental")
	d1 := randomAccountClass(t)
	d1.ProductId = product.Id

	product, _ = testQueriesAccount.GetProductbyName(context.Background(), "Sales")
	d2 := randomAccountClass(t)
	d2.ProductId = product.Id
	// Test Create
	CreatedD1 := createTestAccountClass(t, d1)
	CreatedD2 := createTestAccountClass(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountClass(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.GroupId, getData1.GroupId)
	require.Equal(t, d1.ClassId, getData1.ClassId)
	require.Equal(t, d1.CurId, getData1.CurId)
	require.Equal(t, d1.NoncurId, getData1.NoncurId)
	require.Equal(t, d1.BsAccId, getData1.BsAccId)
	require.Equal(t, d1.IsAccId, getData1.IsAccId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetAccountClass(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.ProductId, getData2.ProductId)
	require.Equal(t, d2.GroupId, getData2.GroupId)
	require.Equal(t, d2.ClassId, getData2.ClassId)
	require.Equal(t, d2.CurId, getData2.CurId)
	require.Equal(t, d2.NoncurId, getData2.NoncurId)
	require.Equal(t, d2.BsAccId, getData2.BsAccId)
	require.Equal(t, d2.IsAccId, getData2.IsAccId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetAccountClassbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)

	getData, err = testQueriesAccount.GetAccountClassbyKeys(
		context.Background(), CreatedD1.ProductId, CreatedD1.GroupId, CreatedD1.ClassId)
	// , CreatedD1.ProductId, CreatedD1.GroupId, CreatedD1.ClassId
	// , productID int64, groupID, int64, classID int64, a int64
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)

	fmt.Printf("%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2. = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountClass(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.ProductId, updatedD1.ProductId)
	require.Equal(t, updateD2.GroupId, updatedD1.GroupId)
	require.Equal(t, updateD2.ClassId, updatedD1.ClassId)
	require.Equal(t, updateD2.CurId, updatedD1.CurId)
	require.Equal(t, updateD2.NoncurId, updatedD1.NoncurId)
	require.Equal(t, updateD2.BsAccId, updatedD1.BsAccId)
	require.Equal(t, updateD2.IsAccId, updatedD1.IsAccId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	// testDeleteAccountClass(t, getData1.Id)
	//	testDeleteAccountClass(t, getData2.Id)
}

func TestListAccountClass(t *testing.T) {

	arg := ListAccountClassParams{
		Limit:  5,
		Offset: 0,
	}

	accountClass, err := testQueriesAccount.ListAccountClass(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountClass)
	require.NotEmpty(t, accountClass)

}

func randomAccountClass(t *testing.T) AccountClassRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	// product := testQueriesAccount.GetProduct(context.Background(), CreatedD2.Id)
	// product, _ := testQueriesAccount.GetProductbyName(context.Background(), "Savings")
	group, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccountTypeGroup", 0, "Microfinance")
	class, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "LoanClass", 0, "Current")
	cur, _ := testQueriesReference.GetChartofAccountbyTitle(context.Background(), "Other Loan")
	nonCur, _ := testQueriesReference.GetChartofAccountbyTitle(context.Background(), "Other Loan")
	bsAcc, _ := testQueriesReference.GetChartofAccountbyTitle(context.Background(), "Other Loan")
	isAcc, _ := testQueriesReference.GetChartofAccountbyTitle(context.Background(), "Other Loan")

	arg := AccountClassRequest{
		// ProductId: product.Id,
		GroupId:   group.Id,
		ClassId:   class.Id,
		CurId:     cur.Id,
		NoncurId:  sql.NullInt64(sql.NullInt64{Int64: nonCur.Id, Valid: true}),
		BsAccId:   sql.NullInt64(sql.NullInt64{Int64: bsAcc.Id, Valid: true}),
		IsAccId:   sql.NullInt64(sql.NullInt64{Int64: isAcc.Id, Valid: true}),
		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestAccountClass(
	t *testing.T,
	d1 AccountClassRequest) model.AccountClass {

	getData1, err := testQueriesAccount.CreateAccountClass(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.GroupId, getData1.GroupId)
	require.Equal(t, d1.ClassId, getData1.ClassId)
	require.Equal(t, d1.CurId, getData1.CurId)
	require.Equal(t, d1.NoncurId, getData1.NoncurId)
	require.Equal(t, d1.BsAccId, getData1.BsAccId)
	require.Equal(t, d1.IsAccId, getData1.IsAccId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestAccountClass(
	t *testing.T,
	d1 AccountClassRequest) model.AccountClass {

	getData1, err := testQueriesAccount.UpdateAccountClass(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.GroupId, getData1.GroupId)
	require.Equal(t, d1.ClassId, getData1.ClassId)
	require.Equal(t, d1.CurId, getData1.CurId)
	require.Equal(t, d1.NoncurId, getData1.NoncurId)
	require.Equal(t, d1.BsAccId, getData1.BsAccId)
	require.Equal(t, d1.IsAccId, getData1.IsAccId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAccountClass(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteAccountClass(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountClass(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
