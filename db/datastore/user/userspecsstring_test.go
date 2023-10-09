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

func TestUserSpecsString(t *testing.T) {

	// Test Data
	d1 := randomUserSpecsString()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomUserSpecsString()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestUserSpecsString(t, d1)
	CreatedD2 := createTestUserSpecsString(t, d2)

	// Get Data
	getData1, err1 := testQueriesUser.GetUserSpecsString(context.Background(), CreatedD1.UserId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	getData2, err2 := testQueriesUser.GetUserSpecsString(context.Background(), CreatedD2.UserId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value, getData2.Value)

	getData, err := testQueriesUser.GetUserSpecsStringbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestUserSpecsString(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value, updatedD1.Value)

	testListUserSpecsString(t, ListUserSpecsStringParams{
		UserId: updatedD1.UserId,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteUserSpecsString(t, CreatedD1.Uuid)
	testDeleteUserSpecsString(t, CreatedD2.Uuid)
}

func testListUserSpecsString(t *testing.T, arg ListUserSpecsStringParams) {

	userSpecsString, err := testQueriesUser.ListUserSpecsString(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", userSpecsString)
	require.NotEmpty(t, userSpecsString)

}

func randomUserSpecsString() UserSpecsStringRequest {

	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")

	arg := UserSpecsStringRequest{
		UserId: usr.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value: util.RandomString(20),
	}
	return arg
}

func createTestUserSpecsString(
	t *testing.T,
	d1 UserSpecsStringRequest) model.UserSpecsString {

	getData1, err := testQueriesUser.CreateUserSpecsString(context.Background(), d1)
	// fmt.Printf("Get by createTestUserSpecsString%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func updateTestUserSpecsString(
	t *testing.T,
	d1 UserSpecsStringRequest) model.UserSpecsString {

	getData1, err := testQueriesUser.UpdateUserSpecsString(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func testDeleteUserSpecsString(t *testing.T, uuid uuid.UUID) {
	err := testQueriesUser.DeleteUserSpecsString(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesUser.GetUserSpecsStringbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
