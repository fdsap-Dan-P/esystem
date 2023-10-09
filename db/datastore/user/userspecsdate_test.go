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

func TestUserSpecsDate(t *testing.T) {

	// Test Data
	d1 := randomUserSpecsDate()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateManufactured")
	d1.SpecsId = item.Id

	d2 := randomUserSpecsDate()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateExpired")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestUserSpecsDate(t, d1)
	CreatedD2 := createTestUserSpecsDate(t, d2)

	// Get Data
	getData1, err1 := testQueriesUser.GetUserSpecsDate(context.Background(), CreatedD1.UserId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	getData2, err2 := testQueriesUser.GetUserSpecsDate(context.Background(), CreatedD2.UserId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.Format("2006-01-02"), getData2.Value.Format("2006-01-02"))
	require.Equal(t, d2.Value2.Format("2006-01-02"), getData2.Value2.Format("2006-01-02"))

	getData, err := testQueriesUser.GetUserSpecsDatebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestUserSpecsDate(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.Format("2006-01-02"), updatedD1.Value.Format("2006-01-02"))
	require.Equal(t, updateD2.Value2.Format("2006-01-02"), updatedD1.Value2.Format("2006-01-02"))

	testListUserSpecsDate(t, ListUserSpecsDateParams{
		UserId: updatedD1.UserId,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteUserSpecsDate(t, CreatedD1.Uuid)
	testDeleteUserSpecsDate(t, CreatedD2.Uuid)
}

func testListUserSpecsDate(t *testing.T, arg ListUserSpecsDateParams) {

	userSpecsDate, err := testQueriesUser.ListUserSpecsDate(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", userSpecsDate)
	require.NotEmpty(t, userSpecsDate)
}

func randomUserSpecsDate() UserSpecsDateRequest {

	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")

	arg := UserSpecsDateRequest{
		UserId: usr.Id,
		// SpecsId: sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomDate(),
		Value2: util.RandomDate(),
	}
	return arg
}

func createTestUserSpecsDate(
	t *testing.T,
	d1 UserSpecsDateRequest) model.UserSpecsDate {

	getData1, err := testQueriesUser.CreateUserSpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func updateTestUserSpecsDate(
	t *testing.T,
	d1 UserSpecsDateRequest) model.UserSpecsDate {

	getData1, err := testQueriesUser.UpdateUserSpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func testDeleteUserSpecsDate(t *testing.T, uuid uuid.UUID) {
	err := testQueriesUser.DeleteUserSpecsDate(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesUser.GetUserSpecsDatebyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
