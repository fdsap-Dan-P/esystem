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

	"github.com/stretchr/testify/require"
)

func TestProduct(t *testing.T) {

	// Test Data
	d1 := randomProduct()
	d2 := randomProduct()

	// Test Create
	CreatedD1 := createTestProduct(t, d1)
	CreatedD2 := createTestProduct(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetProduct(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.ProductName, getData1.ProductName)
	require.Equal(t, d1.Description, getData1.Description)
	require.Equal(t, d1.NormalBalance, getData1.NormalBalance)
	require.Equal(t, d1.Isgl, getData1.Isgl)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetProduct(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Code, getData2.Code)
	require.Equal(t, d2.ProductName, getData2.ProductName)
	require.Equal(t, d2.Description, getData2.Description)
	require.Equal(t, d2.NormalBalance, getData2.NormalBalance)
	require.Equal(t, d2.Isgl, getData2.Isgl)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetProductbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)

	fmt.Printf("%+v\n", getData)

	getData, err = testQueriesAccount.GetProductbyName(context.Background(), CreatedD1.ProductName)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.ProductName = updateD2.ProductName + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestProduct(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Code, updatedD1.Code)
	require.Equal(t, updateD2.ProductName, updatedD1.ProductName)
	require.Equal(t, updateD2.Description, updatedD1.Description)
	require.Equal(t, updateD2.NormalBalance, updatedD1.NormalBalance)
	require.Equal(t, updateD2.Isgl, updatedD1.Isgl)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteProduct(t, getData1.Id)
	testDeleteProduct(t, getData2.Id)
}

func TestListProduct(t *testing.T) {

	arg := ListProductParams{
		Limit:  5,
		Offset: 0,
	}

	product, err := testQueriesAccount.ListProduct(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", product)
	require.NotEmpty(t, product)

}

func randomProduct() ProductRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := ProductRequest{
		Code:          util.RandomInt(1, 100),
		ProductName:   util.RandomString(10),
		Description:   sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		NormalBalance: true,
		Isgl:          true,

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestProduct(
	t *testing.T,
	d1 ProductRequest) model.Product {

	getData1, err := testQueriesAccount.CreateProduct(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.ProductName, getData1.ProductName)
	require.Equal(t, d1.Description, getData1.Description)
	require.Equal(t, d1.NormalBalance, getData1.NormalBalance)
	require.Equal(t, d1.Isgl, getData1.Isgl)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestProduct(
	t *testing.T,
	d1 ProductRequest) model.Product {

	getData1, err := testQueriesAccount.UpdateProduct(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.ProductName, getData1.ProductName)
	require.Equal(t, d1.Description, getData1.Description)
	require.Equal(t, d1.NormalBalance, getData1.NormalBalance)
	require.Equal(t, d1.Isgl, getData1.Isgl)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteProduct(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteProduct(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetProduct(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
