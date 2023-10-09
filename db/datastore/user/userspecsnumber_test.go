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

func TestUserSpecsNumber(t *testing.T) {

	// Test Data
	d1 := randomUserSpecsNumber()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomUserSpecsNumber()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestUserSpecsNumber(t, d1)
	CreatedD2 := createTestUserSpecsNumber(t, d2)

	// Get Data
	getData1, err1 := testQueriesUser.GetUserSpecsNumber(context.Background(), CreatedD1.UserId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	getData2, err2 := testQueriesUser.GetUserSpecsNumber(context.Background(), CreatedD2.UserId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.String(), getData2.Value.String())
	require.Equal(t, d2.Value2.String(), getData2.Value2.String())

	getData, err := testQueriesUser.GetUserSpecsNumberbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestUserSpecsNumber(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.String(), updatedD1.Value.String())
	require.Equal(t, updateD2.Value2.String(), updatedD1.Value2.String())

	testListUserSpecsNumber(t, ListUserSpecsNumberParams{
		UserId: updatedD1.UserId,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteUserSpecsNumber(t, CreatedD1.Uuid)
	testDeleteUserSpecsNumber(t, CreatedD2.Uuid)
}

func testListUserSpecsNumber(t *testing.T, arg ListUserSpecsNumberParams) {

	userSpecsNumber, err := testQueriesUser.ListUserSpecsNumber(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", userSpecsNumber)
	require.NotEmpty(t, userSpecsNumber)

}

func randomUserSpecsNumber() UserSpecsNumberRequest {

	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")

	arg := UserSpecsNumberRequest{
		UserId: usr.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomMoney(),
		Value2: util.RandomMoney(),
	}
	return arg
}

func createTestUserSpecsNumber(
	t *testing.T,
	d1 UserSpecsNumberRequest) model.UserSpecsNumber {

	getData1, err := testQueriesUser.CreateUserSpecsNumber(context.Background(), d1)
	// fmt.Printf("Get by createTestUserSpecsNumber%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func updateTestUserSpecsNumber(
	t *testing.T,
	d1 UserSpecsNumberRequest) model.UserSpecsNumber {

	getData1, err := testQueriesUser.UpdateUserSpecsNumber(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func testDeleteUserSpecsNumber(t *testing.T, uuid uuid.UUID) {
	err := testQueriesUser.DeleteUserSpecsNumber(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesUser.GetUserSpecsNumberbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
