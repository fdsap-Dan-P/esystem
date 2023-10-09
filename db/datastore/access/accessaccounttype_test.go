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

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAccessAccountType(t *testing.T) {

	// Test Data
	d1 := randomAccessAccountType()
	log.Printf("randomAccessAccountType: %v", d1)

	d2 := randomAccessAccountType()

	// Test Create
	CreatedD1 := createTestAccessAccountType(t, d1)
	CreatedD2 := createTestAccessAccountType(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccess.GetAccessAccountType(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccess.GetAccessAccountType(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.RoleId, getData2.RoleId)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.Allow, getData2.Allow)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccess.GetAccessAccountTypebyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccessAccountType(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.RoleId, updatedD1.RoleId)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.Allow, updatedD1.Allow)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteAccessAccountType(t, CreatedD1.Uuid)
	testDeleteAccessAccountType(t, CreatedD2.Uuid)
}

func TestListAccessAccountType(t *testing.T) {

	role, _ := testQueriesAccess.GetAccessRolebyName(context.Background(), "Admin")
	arg := ListAccessAccountTypeParams{
		RoleId: role.Id,
		Limit:  5,
		Offset: 0,
	}

	accessAccountType, err := testQueriesAccess.ListAccessAccountType(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accessAccountType)
	require.NotEmpty(t, accessAccountType)

}

func randomAccessAccountType() AccessAccountTypeRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	role, _ := testQueriesAccess.GetAccessRolebyName(context.Background(), "Bookkeeper")
	p := util.RandomAccountType()
	accType, _ := testQueriesAccount.GetAccountTypebyName(context.Background(), p)

	// fmt.Printf("Get by UUId%+v\n", accType)

	arg := AccessAccountTypeRequest{
		RoleId: role.Id,
		TypeId: accType.Id,
		Allow:  sql.NullBool(sql.NullBool{Bool: true, Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestAccessAccountType(
	t *testing.T,
	d1 AccessAccountTypeRequest) model.AccessAccountType {

	getData1, err := testQueriesAccess.CreateAccessAccountType(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestAccessAccountType(
	t *testing.T,
	d1 AccessAccountTypeRequest) model.AccessAccountType {

	getData1, err := testQueriesAccess.UpdateAccessAccountType(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAccessAccountType(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccess.DeleteAccessAccountType(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccess.GetAccessAccountType(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
