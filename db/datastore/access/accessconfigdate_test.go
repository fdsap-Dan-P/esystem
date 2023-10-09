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

func TestAccessConfigDate(t *testing.T) {

	// Test Data
	d1 := randomAccessConfigDate()
	obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccessObjectType", 0, "Tab")
	d1.ConfigId = obj.Id

	d2 := randomAccessConfigDate()
	obj, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccessObjectType", 0, "Window")
	d1.ConfigId = obj.Id

	// Test Create
	CreatedD1 := createTestAccessConfigDate(t, d1)
	CreatedD2 := createTestAccessConfigDate(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccess.GetAccessConfigDate(context.Background(), CreatedD1.RoleId, CreatedD1.ConfigId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ConfigId, getData1.ConfigId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	getData2, err2 := testQueriesAccess.GetAccessConfigDate(context.Background(), CreatedD2.RoleId, CreatedD2.ConfigId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.RoleId, getData2.RoleId)
	require.Equal(t, d2.ConfigId, getData2.ConfigId)
	require.Equal(t, d2.Value.Format("2006-01-02"), getData2.Value.Format("2006-01-02"))
	require.Equal(t, d2.Value2.Format("2006-01-02"), getData2.Value2.Format("2006-01-02"))

	getData, err := testQueriesAccess.GetAccessConfigDatebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// getData, err = testQueriesAccess.GetAccessConfigDatebyName(context.Background(), CreatedD1.Name)
	// require.NoError(t, err)
	// require.NotEmpty(t, getData)
	// require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	// fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccessConfigDate(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.RoleId, updatedD1.RoleId)
	require.Equal(t, updateD2.ConfigId, updatedD1.ConfigId)
	require.Equal(t, updateD2.Value.Format("2006-01-02"), updatedD1.Value.Format("2006-01-02"))
	require.Equal(t, updateD2.Value2.Format("2006-01-02"), updatedD1.Value2.Format("2006-01-02"))

	// Delete Data
	testDeleteAccessConfigDate(t, CreatedD1.Uuid)
	testDeleteAccessConfigDate(t, CreatedD2.Uuid)
}

func TestListAccessConfigDate(t *testing.T) {
	arg := ListAccessConfigDateParams{
		Limit:  5,
		Offset: 0,
	}

	accessConfigDate, err := testQueriesAccess.ListAccessConfigDate(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accessConfigDate)
	require.NotEmpty(t, accessConfigDate)
}

func randomAccessConfigDate() AccessConfigDateRequest {
	// otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}

	role, _ := testQueriesAccess.GetAccessRolebyName(context.Background(), "Bookkeeper")

	arg := AccessConfigDateRequest{
		RoleId:   role.Id,
		ConfigId: util.RandomInt(1, 100),
		Value:    util.RandomDate(),
		Value2:   util.RandomDate(),
	}
	return arg
}

func createTestAccessConfigDate(
	t *testing.T,
	d1 AccessConfigDateRequest) model.AccessConfigDate {

	getData1, err := testQueriesAccess.CreateAccessConfigDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ConfigId, getData1.ConfigId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func updateTestAccessConfigDate(
	t *testing.T,
	d1 AccessConfigDateRequest) model.AccessConfigDate {

	getData1, err := testQueriesAccess.UpdateAccessConfigDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ConfigId, getData1.ConfigId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func testDeleteAccessConfigDate(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccess.DeleteAccessConfigDate(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccess.GetAccessConfigDatebyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
