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

	"github.com/stretchr/testify/require"
)

func TestAccountType(t *testing.T) {

	// Test Data
	d1 := randomAccountType()
	d1.Uuid = util.SetUUID("65eacb0c-180b-43e5-a499-d28caf2d812c")

	d2 := randomAccountType()
	d2.Uuid = util.SetUUID("bbe88bf6-a771-46a1-9c98-7317b1f7bbef")

	del, _ := testQueriesAccount.GetAccountTypebyUuid(context.Background(), d1.Uuid)
	testDeleteAccountType(t, del.Id)
	del, _ = testQueriesAccount.GetAccountTypebyUuid(context.Background(), d2.Uuid)
	testDeleteAccountType(t, del.Id)

	// Test Create
	CreatedD1 := createTestAccountType(t, d1)
	CreatedD2 := createTestAccountType(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountType(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.AccountType, getData1.AccountType)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.GroupId, getData1.GroupId)
	// require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.NormalBalance, getData1.NormalBalance)
	require.Equal(t, d1.Isgl, getData1.Isgl)
	require.Equal(t, d1.Active, getData1.Active)
	require.Equal(t, d1.FilterType, getData1.FilterType)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetAccountType(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Uuid, getData2.Uuid)
	require.Equal(t, d2.CentralOfficeId, getData2.CentralOfficeId)
	require.Equal(t, d2.Code, getData2.Code)
	require.Equal(t, d2.AccountType, getData2.AccountType)
	require.Equal(t, d2.ProductId, getData2.ProductId)
	require.Equal(t, d2.GroupId, getData2.GroupId)
	// require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.NormalBalance, getData2.NormalBalance)
	require.Equal(t, d2.Isgl, getData2.Isgl)
	require.Equal(t, d2.Active, getData2.Active)
	require.Equal(t, d2.FilterType, getData2.FilterType)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetAccountTypebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)

	getData, err = testQueriesAccount.GetAccountTypebyName(context.Background(), CreatedD1.AccountType)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)

	fmt.Printf("%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.AccountType = updateD2.AccountType + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountType(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Uuid, updatedD1.Uuid)
	require.Equal(t, updateD2.CentralOfficeId, updatedD1.CentralOfficeId)
	require.Equal(t, updateD2.Code, updatedD1.Code)
	require.Equal(t, updateD2.AccountType, updatedD1.AccountType)
	require.Equal(t, updateD2.ProductId, updatedD1.ProductId)
	require.Equal(t, updateD2.GroupId, updatedD1.GroupId)
	// require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.NormalBalance, updatedD1.NormalBalance)
	require.Equal(t, updateD2.Isgl, updatedD1.Isgl)
	require.Equal(t, updateD2.Active, updatedD1.Active)
	require.Equal(t, updateD2.FilterType, updatedD1.FilterType)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteAccountType(t, getData1.Id)
	testDeleteAccountType(t, getData2.Id)
}

func TestListAccountType(t *testing.T) {

	prod, _ := testQueriesAccount.GetProductbyName(context.Background(), "Loan")

	arg := ListAccountTypeParams{
		ProductId: prod.Id,
		Limit:     5,
		Offset:    0,
	}

	accountType, err := testQueriesAccount.ListAccountType(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountType)
	require.NotEmpty(t, accountType)

}

func randomAccountType() AccountTypeRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	prod, _ := testQueriesAccount.GetProductbyName(context.Background(), util.RandomProduct())
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")
	grp, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccountTypeGroup", 0, "Microfinance")

	arg := AccountTypeRequest{
		CentralOfficeId: ofc.Id,
		Code:            util.RandomInt(1, 100),
		AccountType:     util.RandomString(20),
		ProductId:       prod.Id,
		GroupId:         grp.Id,
		// Iiid:            sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		NormalBalance: true,
		Isgl:          true,
		Active:        true,

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestAccountType(
	t *testing.T,
	d1 AccountTypeRequest) model.AccountType {

	getData1, err := testQueriesAccount.CreateAccountType(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.AccountType, getData1.AccountType)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.GroupId, getData1.GroupId)
	// require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.NormalBalance, getData1.NormalBalance)
	require.Equal(t, d1.Isgl, getData1.Isgl)
	require.Equal(t, d1.Active, getData1.Active)
	require.Equal(t, d1.FilterType, getData1.FilterType)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestAccountType(
	t *testing.T,
	d1 AccountTypeRequest) model.AccountType {

	getData1, err := testQueriesAccount.UpdateAccountType(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.AccountType, getData1.AccountType)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.GroupId, getData1.GroupId)
	// require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.NormalBalance, getData1.NormalBalance)
	require.Equal(t, d1.Isgl, getData1.Isgl)
	require.Equal(t, d1.Active, getData1.Active)
	require.Equal(t, d1.FilterType, getData1.FilterType)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAccountType(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteAccountType(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountType(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
