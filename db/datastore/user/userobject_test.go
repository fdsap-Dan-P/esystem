package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUserObject(t *testing.T) {

	// Test Data
	d1 := randomUserObject()
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")
	d1.UserId = usr.Id

	d2 := randomUserObject()
	usr, _ = testQueriesUser.GetUserbyName(context.Background(), "olive.mercado0609@gmail.com")
	d2.UserId = usr.Id

	// Test Create
	CreatedD1 := createTestUserObject(t, d1)
	CreatedD2 := createTestUserObject(t, d2)

	// Get Data
	getData1, err1 := testQueriesUser.GetUserObject(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.ObjectId, getData1.ObjectId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesUser.GetUserObject(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.ObjectId, getData2.ObjectId)
	require.Equal(t, d2.Allow, getData2.Allow)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesUser.GetUserObjectbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestUserObject(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.ObjectId, updatedD1.ObjectId)
	require.Equal(t, updateD2.Allow, updatedD1.Allow)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListUserObject(t, ListUserObjectParams{
		UserId: updatedD1.UserId,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteUserObject(t, CreatedD1.Uuid)
	testDeleteUserObject(t, CreatedD2.Uuid)
}

func testListUserObject(t *testing.T, arg ListUserObjectParams) {

	user_Object, err := testQueriesUser.ListUserObject(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", user_Object)
	require.NotEmpty(t, user_Object)

}

func randomUserObject() UserObjectRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccessObjectType", 0, "Tab")

	arg := UserObjectRequest{
		// UserId:   util.RandomInt(1, 100),
		ObjectId: obj.Id,
		Allow:    true,

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestUserObject(
	t *testing.T,
	d1 UserObjectRequest) model.UserObject {

	getData1, err := testQueriesUser.CreateUserObject(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.ObjectId, getData1.ObjectId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestUserObject(
	t *testing.T,
	d1 UserObjectRequest) model.UserObject {

	getData1, err := testQueriesUser.UpdateUserObject(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.ObjectId, getData1.ObjectId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteUserObject(t *testing.T, uuid uuid.UUID) {
	err := testQueriesUser.DeleteUserObject(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesUser.GetUserObject(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
