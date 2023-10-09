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

func TestUserSpecsRef(t *testing.T) {

	// Test Data
	d1 := randomUserSpecsRef()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomUserSpecsRef()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestUserSpecsRef(t, d1)
	CreatedD2 := createTestUserSpecsRef(t, d2)

	// Get Data
	getData1, err1 := testQueriesUser.GetUserSpecsRef(context.Background(), CreatedD1.UserId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	getData2, err2 := testQueriesUser.GetUserSpecsRef(context.Background(), CreatedD2.UserId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.RefId, getData2.RefId)

	getData, err := testQueriesUser.GetUserSpecsRefbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestUserSpecsRef(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.RefId, updatedD1.RefId)

	testListUserSpecsRef(t, ListUserSpecsRefParams{
		UserId: updatedD1.UserId,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteUserSpecsRef(t, CreatedD1.Uuid)
	testDeleteUserSpecsRef(t, CreatedD2.Uuid)
}

func testListUserSpecsRef(t *testing.T, arg ListUserSpecsRefParams) {

	userSpecsRef, err := testQueriesUser.ListUserSpecsRef(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", userSpecsRef)
	require.NotEmpty(t, userSpecsRef)

}

func randomUserSpecsRef() UserSpecsRefRequest {

	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeetype", 0, "Employed")

	arg := UserSpecsRefRequest{
		UserId: usr.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		RefId: typ.Id,
	}
	return arg
}

func createTestUserSpecsRef(
	t *testing.T,
	d1 UserSpecsRefRequest) model.UserSpecsRef {

	getData1, err := testQueriesUser.CreateUserSpecsRef(context.Background(), d1)
	// fmt.Printf("Get by createTestUserSpecsRef%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func updateTestUserSpecsRef(
	t *testing.T,
	d1 UserSpecsRefRequest) model.UserSpecsRef {

	getData1, err := testQueriesUser.UpdateUserSpecsRef(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func testDeleteUserSpecsRef(t *testing.T, uuid uuid.UUID) {
	err := testQueriesUser.DeleteUserSpecsRef(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesUser.GetUserSpecsRefbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
