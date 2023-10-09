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

func TestUserProduct(t *testing.T) {

	// Test Data
	d1 := randomUserProduct()
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")
	d1.UserId = usr.Id

	d2 := randomUserProduct()
	usr, _ = testQueriesUser.GetUserbyName(context.Background(), "olive.mercado0609@gmail.com")
	d2.UserId = usr.Id

	// Test Create
	CreatedD1 := createTestUserProduct(t, d1)
	CreatedD2 := createTestUserProduct(t, d2)

	// Get Data
	getData1, err1 := testQueriesUser.GetUserProduct(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesUser.GetUserProduct(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.ProductId, getData2.ProductId)
	require.Equal(t, d2.Allow, getData2.Allow)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesUser.GetUserProductbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestUserProduct(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.ProductId, updatedD1.ProductId)
	require.Equal(t, updateD2.Allow, updatedD1.Allow)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListUserProduct(t, ListUserProductParams{
		UserId: updatedD1.UserId,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteUserProduct(t, CreatedD1.Uuid)
	testDeleteUserProduct(t, CreatedD2.Uuid)
}

func testListUserProduct(t *testing.T, arg ListUserProductParams) {

	userProduct, err := testQueriesUser.ListUserProduct(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", userProduct)
	require.NotEmpty(t, userProduct)

}

func randomUserProduct() UserProductRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	p := util.RandomProduct()
	prod, _ := testQueriesAccount.GetProductbyName(context.Background(), p)

	arg := UserProductRequest{
		UserId:    util.RandomInt(1, 100),
		ProductId: prod.Id,
		Allow:     true,

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestUserProduct(
	t *testing.T,
	d1 UserProductRequest) model.UserProduct {

	getData1, err := testQueriesUser.CreateUserProduct(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestUserProduct(
	t *testing.T,
	d1 UserProductRequest) model.UserProduct {

	getData1, err := testQueriesUser.UpdateUserProduct(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteUserProduct(t *testing.T, uuid uuid.UUID) {
	err := testQueriesUser.DeleteUserProduct(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesUser.GetUserProduct(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
