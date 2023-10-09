package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"fmt"
	"testing"

	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAccountTypeGroup(t *testing.T) {

	// Test Data
	d1 := randomAccountTypeGroup(t)
	d1.AccountTypeGroup = "Microfinance"

	d2 := randomAccountTypeGroup(t)
	d2.Uuid = util.SetUUID("2a2df1f3-858c-4b9e-ad87-7f861545005d")
	d2.AccountTypeGroup = "Regular"

	// Test Create
	CreatedD1 := createTestAccountTypeGroup(t, d1)
	CreatedD2 := createTestAccountTypeGroup(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountTypeGroup(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.GroupId, getData1.GroupId)
	require.Equal(t, d1.AccountTypeGroup, getData1.AccountTypeGroup)
	// require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.NormalBalance, getData1.NormalBalance)
	require.Equal(t, d1.Isgl, getData1.Isgl)
	require.Equal(t, d1.Active, getData1.Active)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetAccountTypeGroup(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Uuid, getData2.Uuid)
	require.Equal(t, d2.ProductId, getData2.ProductId)
	require.Equal(t, d2.GroupId, getData2.GroupId)
	require.Equal(t, d2.AccountTypeGroup, getData2.AccountTypeGroup)
	// require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.NormalBalance, getData2.NormalBalance)
	require.Equal(t, d2.Isgl, getData2.Isgl)
	require.Equal(t, d2.Active, getData2.Active)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetAccountTypeGroupbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.AccountTypeGroup = updateD2.AccountTypeGroup + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountTypeGroup(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Uuid, updatedD1.Uuid)
	require.Equal(t, updateD2.ProductId, updatedD1.ProductId)
	require.Equal(t, updateD2.GroupId, updatedD1.GroupId)
	require.Equal(t, updateD2.AccountTypeGroup, updatedD1.AccountTypeGroup)
	// require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.NormalBalance, updatedD1.NormalBalance)
	require.Equal(t, updateD2.Isgl, updatedD1.Isgl)
	require.Equal(t, updateD2.Active, updatedD1.Active)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListAccountTypeGroup(t, ListAccountTypeGroupParams{
		ProductId: updatedD1.ProductId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteAccountTypeGroup(t, CreatedD1.Uuid)
	testDeleteAccountTypeGroup(t, CreatedD2.Uuid)
}

func testListAccountTypeGroup(t *testing.T, arg ListAccountTypeGroupParams) {

	accountTypeGroup, err := testQueriesAccount.ListAccountTypeGroup(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountTypeGroup)
	require.NotEmpty(t, accountTypeGroup)

}

func randomAccountTypeGroup(t *testing.T) AccountTypeGroupRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	prod, _ := testQueriesAccount.GetProductbyName(context.Background(), util.RandomProduct())
	grp, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccountTypeGroup", 0, "Microfinance")

	arg := AccountTypeGroupRequest{
		Uuid:             util.ToUUID("91788dad-3d4f-4117-9ed0-7b817da8ab12"),
		ProductId:        prod.Id,
		GroupId:          grp.Id,
		AccountTypeGroup: "Microfinance",
		NormalBalance:    true,
		Isgl:             true,
		Active:           true,
		OtherInfo:        sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestAccountTypeGroup(
	t *testing.T,
	d1 AccountTypeGroupRequest) model.AccountTypeGroup {

	getData1, err := testQueriesAccount.CreateAccountTypeGroup(context.Background(), d1)
	// fmt.Printf("Get by createTestAccountTypeGroup%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.GroupId, getData1.GroupId)
	require.Equal(t, d1.AccountTypeGroup, getData1.AccountTypeGroup)
	// require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.NormalBalance, getData1.NormalBalance)
	require.Equal(t, d1.Isgl, getData1.Isgl)
	require.Equal(t, d1.Active, getData1.Active)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestAccountTypeGroup(
	t *testing.T,
	d1 AccountTypeGroupRequest) model.AccountTypeGroup {

	getData1, err := testQueriesAccount.UpdateAccountTypeGroup(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.GroupId, getData1.GroupId)
	require.Equal(t, d1.AccountTypeGroup, getData1.AccountTypeGroup)
	// require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.NormalBalance, getData1.NormalBalance)
	require.Equal(t, d1.Isgl, getData1.Isgl)
	require.Equal(t, d1.Active, getData1.Active)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAccountTypeGroup(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteAccountTypeGroup(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountTypeGroupbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
