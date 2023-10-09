package db

import (
	"context"
	"database/sql"
	"log"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUserAccountType(t *testing.T) {

	// Test Data
	d1 := randomUserAccountType()
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")
	d1.UserId = usr.Id
	d1.Uuid = util.SetUUID("156272d2-65d8-4a73-84df-38de0e118173")
	accType, _ := testQueriesAccount.GetAccountTypebyName(context.Background(), "Sipag")
	d1.AccountTypeId = accType.Id

	log.Printf("randomUserAccountType 1%v: ", usr)
	log.Printf("randomUserAccountType accType%v: ", accType)

	d2 := randomUserAccountType()
	usr, _ = testQueriesUser.GetUserbyName(context.Background(), "olive.mercado0609@gmail.com")
	d2.UserId = usr.Id
	d2.Uuid = util.SetUUID("532fd459-59b0-4648-91e9-ea4fdda7f881")

	accType, _ = testQueriesAccount.GetAccountTypebyName(context.Background(), "Sikap 1")
	d2.AccountTypeId = accType.Id

	log.Printf("randomUserAccountType 2%v: ", usr)
	log.Printf("randomUserAccountType accType%v: ", accType)

	// Test Create
	CreatedD1 := createTestUserAccountType(t, d1)
	CreatedD2 := createTestUserAccountType(t, d2)

	// Get Data
	getData1, err1 := testQueriesUser.GetUserAccountType(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesUser.GetUserAccountType(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.AccountTypeId, getData2.AccountTypeId)
	require.Equal(t, d2.Allow, getData2.Allow)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesUser.GetUserAccountTypebyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestUserAccountType(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.AccountTypeId, updatedD1.AccountTypeId)
	require.Equal(t, updateD2.Allow, updatedD1.Allow)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListUserAccountType(t, ListUserAccountTypeParams{
		UserId: updatedD1.UserId,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteUserAccountType(t, CreatedD1.Uuid)
	testDeleteUserAccountType(t, CreatedD2.Uuid)
}

func testListUserAccountType(t *testing.T, arg ListUserAccountTypeParams) {

	userAccountType, err := testQueriesUser.ListUserAccountType(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", userAccountType)
	require.NotEmpty(t, userAccountType)

}

func randomUserAccountType() UserAccountTypeRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := UserAccountTypeRequest{
		Allow: true,

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestUserAccountType(
	t *testing.T,
	d1 UserAccountTypeRequest) model.UserAccountType {

	getData1, err := testQueriesUser.CreateUserAccountType(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestUserAccountType(
	t *testing.T,
	d1 UserAccountTypeRequest) model.UserAccountType {

	getData1, err := testQueriesUser.UpdateUserAccountType(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteUserAccountType(t *testing.T, uuid uuid.UUID) {
	err := testQueriesUser.DeleteUserAccountType(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesUser.GetUserAccountType(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
