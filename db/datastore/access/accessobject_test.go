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

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAccessObject(t *testing.T) {

	// Test Data
	obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccessObjectType", 0, "Tab")
	d1 := randomAccessObject()
	d1.ObjectId = obj.Id
	// fmt.Printf("Access Object1%+v\n", obj)
	obj, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccessObjectType", 0, "Window")
	d2 := randomAccessObject()
	d2.ObjectId = obj.Id
	// fmt.Printf("Access Object2%+v\n", obj)

	// Test Create
	CreatedD1 := createTestAccessObject(t, d1)
	CreatedD2 := createTestAccessObject(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccess.GetAccessObject(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ObjectId, getData1.ObjectId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.Equal(t, d1.MaxValue.String(), getData1.MaxValue.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccess.GetAccessObject(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.RoleId, getData2.RoleId)
	require.Equal(t, d2.ObjectId, getData2.ObjectId)
	require.Equal(t, d2.Allow, getData2.Allow)
	require.Equal(t, d2.MaxValue.String(), getData2.MaxValue.String())
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccess.GetAccessObjectbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// getData, err = testQueriesAccess.GetAccessObjectbyName(context.Background(), CreatedD1.Name)
	// require.NoError(t, err)
	// require.NotEmpty(t, getData)
	// require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	// fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccessObject(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.RoleId, updatedD1.RoleId)
	require.Equal(t, updateD2.ObjectId, updatedD1.ObjectId)
	require.Equal(t, updateD2.Allow, updatedD1.Allow)
	require.Equal(t, updateD2.MaxValue.String(), updatedD1.MaxValue.String())
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteAccessObject(t, CreatedD1.Uuid)
	testDeleteAccessObject(t, CreatedD2.Uuid)
}

func TestListAccessObject(t *testing.T) {

	role, _ := testQueriesAccess.GetAccessRolebyName(context.Background(), "Bookkeeper")

	arg := ListAccessObjectParams{
		RoleId: role.Id,
		Limit:  5,
		Offset: 0,
	}

	accessObject, err := testQueriesAccess.ListAccessObject(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accessObject)
	require.NotEmpty(t, accessObject)

}

func randomAccessObject() AccessObjectRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	role, _ := testQueriesAccess.GetAccessRolebyName(context.Background(), "Bookkeeper")
	// p := util.RandomAccountType()
	arg := AccessObjectRequest{
		RoleId: role.Id,
		// ObjectId: obj,Id,
		Allow:    sql.NullBool(sql.NullBool{Bool: true, Valid: true}),
		MaxValue: util.RandomMoney(),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestAccessObject(
	t *testing.T,
	d1 AccessObjectRequest) model.AccessObject {
	getData1, err := testQueriesAccess.CreateAccessObject(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ObjectId, getData1.ObjectId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.Equal(t, d1.MaxValue.String(), getData1.MaxValue.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestAccessObject(
	t *testing.T,
	d1 AccessObjectRequest) model.AccessObject {

	getData1, err := testQueriesAccess.UpdateAccessObject(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ObjectId, getData1.ObjectId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.Equal(t, d1.MaxValue.String(), getData1.MaxValue.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAccessObject(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccess.DeleteAccessObject(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccess.GetAccessObject(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
