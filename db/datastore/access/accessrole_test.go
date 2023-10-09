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

func TestAccessRole(t *testing.T) {

	// Test Data
	d1 := randomAccessRole()
	d2 := randomAccessRole()

	// Test Create
	CreatedD1 := createTestAccessRole(t, d1)
	CreatedD2 := createTestAccessRole(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccess.GetAccessRole(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AccessName, getData1.AccessName)
	require.Equal(t, d1.Description, getData1.Description)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccess.GetAccessRole(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AccessName, getData2.AccessName)
	require.Equal(t, d2.Description, getData2.Description)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccess.GetAccessRolebyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	getData, err = testQueriesAccess.GetAccessRolebyName(context.Background(), CreatedD1.AccessName)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.AccessName = updateD2.AccessName + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccessRole(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.AccessName, updatedD1.AccessName)
	require.Equal(t, updateD2.Description, updatedD1.Description)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteAccessRole(t, getData1.Id)
	testDeleteAccessRole(t, getData2.Id)
}

func TestListAccessRole(t *testing.T) {

	arg := ListAccessRoleParams{
		Limit:  5,
		Offset: 0,
	}

	accessRole, err := testQueriesAccess.ListAccessRole(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accessRole)
	require.NotEmpty(t, accessRole)

}

func randomAccessRole() AccessRoleRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := AccessRoleRequest{
		AccessName:  util.RandomString(10),
		Description: sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestAccessRole(
	t *testing.T,
	d1 AccessRoleRequest) model.AccessRole {

	getData1, err := testQueriesAccess.CreateAccessRole(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccessName, getData1.AccessName)
	require.Equal(t, d1.Description, getData1.Description)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestAccessRole(
	t *testing.T,
	d1 AccessRoleRequest) model.AccessRole {

	getData1, err := testQueriesAccess.UpdateAccessRole(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccessName, getData1.AccessName)
	require.Equal(t, d1.Description, getData1.Description)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAccessRole(t *testing.T, id int64) {
	err := testQueriesAccess.DeleteAccessRole(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccess.GetAccessRole(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
