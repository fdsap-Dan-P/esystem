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

func TestAccessConfigNumber(t *testing.T) {

	// Test Data
	d1 := randomAccessConfigNumber()
	obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccessObjectType", 0, "Tab")
	d1.ConfigId = obj.Id

	d2 := randomAccessConfigNumber()
	obj, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccessObjectType", 0, "Window")
	d2.ConfigId = obj.Id

	// Test Create
	CreatedD1 := createTestAccessConfigNumber(t, d1)
	CreatedD2 := createTestAccessConfigNumber(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccess.GetAccessConfigNumber(context.Background(), CreatedD1.RoleId, CreatedD1.ConfigId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ConfigId, getData1.ConfigId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	getData2, err2 := testQueriesAccess.GetAccessConfigNumber(context.Background(), CreatedD2.RoleId, CreatedD2.ConfigId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.RoleId, getData2.RoleId)
	require.Equal(t, d2.ConfigId, getData2.ConfigId)
	require.Equal(t, d2.Value.String(), getData2.Value.String())
	require.Equal(t, d2.Value2.String(), getData2.Value2.String())

	getData, err := testQueriesAccess.GetAccessConfigNumberbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// getData, err = testQueriesAccess.GetAccessConfigNumberbyName(context.Background(), CreatedD1.Name)
	// require.NoError(t, err)
	// require.NotEmpty(t, getData)
	// require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	// fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccessConfigNumber(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.RoleId, updatedD1.RoleId)
	require.Equal(t, updateD2.ConfigId, updatedD1.ConfigId)
	require.Equal(t, updateD2.Value.String(), updatedD1.Value.String())
	require.Equal(t, updateD2.Value2.String(), updatedD1.Value2.String())

	// Delete Data
	testDeleteAccessConfigNumber(t, CreatedD1.Uuid)
	testDeleteAccessConfigNumber(t, CreatedD2.Uuid)
}

func TestListAccessConfigNumber(t *testing.T) {
	arg := ListAccessConfigNumberParams{
		Limit:  5,
		Offset: 0,
	}

	accessConfigNumber, err := testQueriesAccess.ListAccessConfigNumber(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accessConfigNumber)
	require.NotEmpty(t, accessConfigNumber)
}

func randomAccessConfigNumber() AccessConfigNumberRequest {
	// otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}

	role, _ := testQueriesAccess.GetAccessRolebyName(context.Background(), "Bookkeeper")

	arg := AccessConfigNumberRequest{
		RoleId:   role.Id,
		ConfigId: util.RandomInt(1, 100),
		Value:    util.SetDecimal("23443"),
		Value2:   util.SetDecimal("0"),
	}
	return arg
}

func createTestAccessConfigNumber(
	t *testing.T,
	d1 AccessConfigNumberRequest) model.AccessConfigNumber {

	getData1, err := testQueriesAccess.CreateAccessConfigNumber(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ConfigId, getData1.ConfigId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func updateTestAccessConfigNumber(
	t *testing.T,
	d1 AccessConfigNumberRequest) model.AccessConfigNumber {

	getData1, err := testQueriesAccess.UpdateAccessConfigNumber(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ConfigId, getData1.ConfigId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func testDeleteAccessConfigNumber(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccess.DeleteAccessConfigNumber(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccess.GetAccessConfigNumberbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
