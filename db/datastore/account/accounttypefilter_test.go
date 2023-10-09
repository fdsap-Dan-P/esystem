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

func TestAccountTypeFilter(t *testing.T) {

	// Test Data
	d1 := randomAccountTypeFilter(t)
	d1.Uuid = util.SetUUID("95046e2f-4615-4240-8f68-851f94b11aad")
	accType, _ := testQueriesAccount.GetAccountTypebyName(context.Background(), "Sipag")
	d1.AccountTypeId = accType.Id

	d2 := randomAccountTypeFilter(t)
	d2.Uuid = util.SetUUID("46ccb26d-fb1d-4af5-936c-9e1a98e9d250")
	accType, _ = testQueriesAccount.GetAccountTypebyName(context.Background(), "Sikap 1")
	d2.AccountTypeId = accType.Id

	// Test Create
	CreatedD1 := createTestAccountTypeFilter(t, d1)
	CreatedD2 := createTestAccountTypeFilter(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountTypeFilter(context.Background(), CreatedD1.CentralOfficeId, CreatedD1.AccountTypeId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d2.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetAccountTypeFilter(context.Background(), CreatedD2.CentralOfficeId, CreatedD2.AccountTypeId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AccountTypeId, getData2.AccountTypeId)
	require.Equal(t, d2.CentralOfficeId, getData2.CentralOfficeId)
	require.Equal(t, d2.Allow, getData2.Allow)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetAccountTypeFilterbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	// updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountTypeFilter(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.AccountTypeId, updatedD1.AccountTypeId)
	require.Equal(t, updateD2.CentralOfficeId, updatedD1.CentralOfficeId)
	require.Equal(t, updateD2.Allow, updatedD1.Allow)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListAccountTypeFilter(t, ListAccountTypeFilterParams{
		CentralOfficeId: updatedD1.CentralOfficeId,
		Limit:           5,
		Offset:          0,
	})

	// Delete Data
	testDeleteAccountTypeFilter(t, CreatedD1.Uuid)
	testDeleteAccountTypeFilter(t, CreatedD2.Uuid)
}

func testListAccountTypeFilter(t *testing.T, arg ListAccountTypeFilterParams) {

	accountTypeFilter, err := testQueriesAccount.ListAccountTypeFilter(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountTypeFilter)
	require.NotEmpty(t, accountTypeFilter)

}

func randomAccountTypeFilter(t *testing.T) AccountTypeFilterRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")

	arg := AccountTypeFilterRequest{
		CentralOfficeId: ofc.Id,
		Allow:           true,
		OtherInfo:       sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestAccountTypeFilter(
	t *testing.T,
	d1 AccountTypeFilterRequest) model.AccountTypeFilter {

	getData1, err := testQueriesAccount.CreateAccountTypeFilter(context.Background(), d1)
	// fmt.Printf("Get by createTestAccountTypeFilter%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestAccountTypeFilter(
	t *testing.T,
	d1 AccountTypeFilterRequest) model.AccountTypeFilter {

	getData1, err := testQueriesAccount.UpdateAccountTypeFilter(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAccountTypeFilter(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteAccountTypeFilter(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountTypeFilterbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
