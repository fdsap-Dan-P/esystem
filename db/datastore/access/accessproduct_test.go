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

func TestAccessProduct(t *testing.T) {

	// Test Data
	d1 := randomAccessProduct()
	d2 := randomAccessProduct()

	// Test Create
	CreatedD1 := createTestAccessProduct(t, d1)
	CreatedD2 := createTestAccessProduct(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccess.GetAccessProduct(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccess.GetAccessProduct(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.RoleId, getData2.RoleId)
	require.Equal(t, d2.ProductId, getData2.ProductId)
	require.Equal(t, d2.Allow, getData2.Allow)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccess.GetAccessProductbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccessProduct(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.RoleId, updatedD1.RoleId)
	require.Equal(t, updateD2.ProductId, updatedD1.ProductId)
	require.Equal(t, updateD2.Allow, updatedD1.Allow)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteAccessProduct(t, CreatedD1.Uuid)
	testDeleteAccessProduct(t, CreatedD2.Uuid)
}

func TestListAccessProduct(t *testing.T) {

	arg := ListAccessProductParams{
		RoleId: 1,
		Limit:  5,
		Offset: 0,
	}

	accessProduct, err := testQueriesAccess.ListAccessProduct(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accessProduct)
	require.NotEmpty(t, accessProduct)

}

func randomAccessProduct() AccessProductRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	role, _ := testQueriesAccess.GetAccessRolebyName(context.Background(), "Bookkeeper")
	p := util.RandomProduct()
	prod, _ := testQueriesAccount.GetProductbyName(context.Background(), p)

	arg := AccessProductRequest{
		RoleId:    role.Id,
		ProductId: prod.Id,
		Allow:     sql.NullBool(sql.NullBool{Bool: true, Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestAccessProduct(
	t *testing.T,
	d1 AccessProductRequest) model.AccessProduct {

	getData1, err := testQueriesAccess.CreateAccessProduct(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestAccessProduct(
	t *testing.T,
	d1 AccessProductRequest) model.AccessProduct {

	getData1, err := testQueriesAccess.UpdateAccessProduct(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.Allow, getData1.Allow)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAccessProduct(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccess.DeleteAccessProduct(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccess.GetAccessProduct(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
