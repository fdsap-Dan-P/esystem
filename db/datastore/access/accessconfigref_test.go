package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"simplebank/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAccessConfigRef(t *testing.T) {

	// Test Data
	d1 := randomAccessConfigRef()
	obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccessObjectType", 0, "Tab")
	d1.ConfigId = obj.Id

	d2 := randomAccessConfigRef()
	obj, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccessObjectType", 0, "Window")
	d2.ConfigId = obj.Id

	// Test Create
	CreatedD1 := createTestAccessConfigRef(t, d1)
	CreatedD2 := createTestAccessConfigRef(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccess.GetAccessConfigRef(context.Background(), CreatedD1.RoleId, CreatedD1.ConfigId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ConfigId, getData1.ConfigId)
	require.Equal(t, d1.RefId, getData1.RefId)

	getData2, err2 := testQueriesAccess.GetAccessConfigRef(context.Background(), CreatedD2.RoleId, CreatedD2.ConfigId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.RoleId, getData2.RoleId)
	require.Equal(t, d2.ConfigId, getData2.ConfigId)
	require.Equal(t, d2.RefId, getData2.RefId)

	getData, err := testQueriesAccess.GetAccessConfigRefbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccessConfigRef(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.RoleId, updatedD1.RoleId)
	require.Equal(t, updateD2.ConfigId, updatedD1.ConfigId)
	require.Equal(t, updateD2.RefId, updatedD1.RefId)

	testListAccessConfigRef(t, ListAccessConfigRefParams{
		RoleId: updatedD1.RoleId,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteAccessConfigRef(t, CreatedD1.Uuid)
	testDeleteAccessConfigRef(t, CreatedD2.Uuid)
}

func testListAccessConfigRef(t *testing.T, arg ListAccessConfigRefParams) {

	accessConfigRef, err := testQueriesAccess.ListAccessConfigRef(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accessConfigRef)
	require.NotEmpty(t, accessConfigRef)

}

func randomAccessConfigRef() AccessConfigRefRequest {

	role, _ := testQueriesAccess.GetAccessRolebyName(context.Background(), "Bookkeeper")
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeetype", 0, "Employed")

	arg := AccessConfigRefRequest{
		RoleId: role.Id,
		RefId:  typ.Id,
	}
	return arg
}

func createTestAccessConfigRef(
	t *testing.T,
	d1 AccessConfigRefRequest) model.AccessConfigRef {

	getData1, err := testQueriesAccess.CreateAccessConfigRef(context.Background(), d1)
	// fmt.Printf("Get by createTestAccessConfigRef%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ConfigId, getData1.ConfigId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func updateTestAccessConfigRef(
	t *testing.T,
	d1 AccessConfigRefRequest) model.AccessConfigRef {

	getData1, err := testQueriesAccess.UpdateAccessConfigRef(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ConfigId, getData1.ConfigId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func testDeleteAccessConfigRef(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccess.DeleteAccessConfigRef(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccess.GetAccessConfigRefbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
