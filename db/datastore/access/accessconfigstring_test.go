package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAccessConfigString(t *testing.T) {

	// Test Data
	d1 := randomAccessConfigString()
	obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccessObjectType", 0, "Tab")
	d1.ConfigId = obj.Id

	d2 := randomAccessConfigString()
	obj, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccessObjectType", 0, "Window")
	d1.ConfigId = obj.Id

	// Test Create
	CreatedD1 := createTestAccessConfigString(t, d1)
	CreatedD2 := createTestAccessConfigString(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccess.GetAccessConfigString(context.Background(), CreatedD1.RoleId, CreatedD1.ConfigId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ConfigId, getData1.ConfigId)
	require.Equal(t, d1.Value, getData1.Value)

	getData2, err2 := testQueriesAccess.GetAccessConfigString(context.Background(), CreatedD2.RoleId, CreatedD2.ConfigId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.RoleId, getData2.RoleId)
	require.Equal(t, d2.ConfigId, getData2.ConfigId)
	require.Equal(t, d2.Value, getData2.Value)

	getData, err := testQueriesAccess.GetAccessConfigStringbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// getData, err = testQueriesAccess.GetAccessConfigStringbyName(context.Background(), CreatedD1.Name)
	// require.NoError(t, err)
	// require.NotEmpty(t, getData)
	// require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	// fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccessConfigString(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.RoleId, updatedD1.RoleId)
	require.Equal(t, updateD2.ConfigId, updatedD1.ConfigId)
	require.Equal(t, updateD2.Value, updatedD1.Value)

	// Delete Data
	testDeleteAccessConfigString(t, CreatedD1.Uuid)
	testDeleteAccessConfigString(t, CreatedD2.Uuid)
}

func TestListAccessConfigString(t *testing.T) {
	arg := ListAccessConfigStringParams{
		Limit:  5,
		Offset: 0,
	}

	accessConfigString, err := testQueriesAccess.ListAccessConfigString(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accessConfigString)
	require.NotEmpty(t, accessConfigString)
}

func randomAccessConfigString() AccessConfigStringRequest {
	// otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}

	role, _ := testQueriesAccess.GetAccessRolebyName(context.Background(), "Bookkeeper")

	arg := AccessConfigStringRequest{
		RoleId:   role.Id,
		ConfigId: util.RandomInt(1, 100),
		Value:    util.RandomString(10),
	}
	return arg
}

func createTestAccessConfigString(
	t *testing.T,
	d1 AccessConfigStringRequest) model.AccessConfigString {

	getData1, err := testQueriesAccess.CreateAccessConfigString(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ConfigId, getData1.ConfigId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func updateTestAccessConfigString(
	t *testing.T,
	d1 AccessConfigStringRequest) model.AccessConfigString {

	getData1, err := testQueriesAccess.UpdateAccessConfigString(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ConfigId, getData1.ConfigId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func testDeleteAccessConfigString(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccess.DeleteAccessConfigString(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccess.GetAccessConfigStringbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
