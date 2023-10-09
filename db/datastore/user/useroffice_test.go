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

func TestUserOffice(t *testing.T) {

	// Test Data
	d1 := randomUserOffice()
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")
	d1.UserId = usr.Id

	d2 := randomUserOffice()
	usr, _ = testQueriesUser.GetUserbyName(context.Background(), "olive.mercado0609@gmail.com")
	d2.UserId = usr.Id

	// Test Create
	CreatedD1 := createTestUserOffice(t, d1)
	CreatedD2 := createTestUserOffice(t, d2)

	// Get Data
	getData1, err1 := testQueriesUser.GetUserOffice(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesUser.GetUserOffice(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.OfficeId, getData2.OfficeId)
	require.Equal(t, d2.Allow, getData2.Allow)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesUser.GetUserOfficebyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestUserOffice(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.OfficeId, updatedD1.OfficeId)
	require.Equal(t, updateD2.Allow, updatedD1.Allow)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListUserOffice(t, ListUserOfficeParams{
		UserId: updatedD1.UserId,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteUserOffice(t, CreatedD1.Uuid)
	testDeleteUserOffice(t, CreatedD2.Uuid)
}

func testListUserOffice(t *testing.T, arg ListUserOfficeParams) {

	userOffice, err := testQueriesUser.ListUserOffice(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", userOffice)
	require.NotEmpty(t, userOffice)

}

func randomUserOffice() UserOfficeRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	ofc, _ := testQueriesIdentity.GetOfficebyUuId(context.Background(), util.ToUUID("97a46847-b7d4-425a-a12a-63e5e8571cf8"))

	arg := UserOfficeRequest{
		// UserId:   util.RandomInt(1, 100),
		OfficeId: ofc.Id,
		Allow:    true,

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestUserOffice(
	t *testing.T,
	d1 UserOfficeRequest) model.UserOffice {

	getData1, err := testQueriesUser.CreateUserOffice(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestUserOffice(
	t *testing.T,
	d1 UserOfficeRequest) model.UserOffice {

	getData1, err := testQueriesUser.UpdateUserOffice(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteUserOffice(t *testing.T, uuid uuid.UUID) {
	err := testQueriesUser.DeleteUserOffice(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesUser.GetUserOffice(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
