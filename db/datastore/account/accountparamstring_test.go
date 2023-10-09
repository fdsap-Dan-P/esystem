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

func TestAccountParamString(t *testing.T) {

	// Test Data
	d1 := randomAccountParamString(t)
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Parameter", 0, "Interest")
	d1.ItemId = item.Id

	d2 := randomAccountParamString(t)
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Parameter", 0, "Charges")
	d2.ItemId = item.Id

	// Test Create
	CreatedD1 := createTestAccountParamString(t, d1)
	CreatedD2 := createTestAccountParamString(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountParamString(context.Background(), CreatedD1.ParamId, CreatedD1.ItemId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.ParamId, getData1.ParamId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value, getData1.Value)

	getData2, err2 := testQueriesAccount.GetAccountParamString(context.Background(), CreatedD2.ParamId, CreatedD2.ItemId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.ParamId, getData2.ParamId)
	require.Equal(t, d2.ItemId, getData2.ItemId)
	require.Equal(t, d2.Value, getData2.Value)

	getData, err := testQueriesAccount.GetAccountParamStringbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountParamString(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.ParamId, updatedD1.ParamId)
	require.Equal(t, updateD2.ItemId, updatedD1.ItemId)
	require.Equal(t, updateD2.Value, updatedD1.Value)

	testListAccountParamString(t, ListAccountParamStringParams{
		ParamId: updatedD1.ParamId,
		Limit:   5,
		Offset:  0,
	})

	// Delete Data
	testDeleteAccountParamString(t, CreatedD1.Uuid)
	testDeleteAccountParamString(t, CreatedD2.Uuid)
}

func testListAccountParamString(t *testing.T, arg ListAccountParamStringParams) {

	accountParamString, err := testQueriesAccount.ListAccountParamString(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountParamString)
	require.NotEmpty(t, accountParamString)

}

func randomAccountParamString(t *testing.T) AccountParamStringRequest {

	accParam := createTestAccountParam(t, randomAccountParam())

	arg := AccountParamStringRequest{
		ParamId: accParam.Id,
		// ItemId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value: util.RandomString(20),
	}
	return arg
}

func createTestAccountParamString(
	t *testing.T,
	d1 AccountParamStringRequest) model.AccountParamString {

	getData1, err := testQueriesAccount.CreateAccountParamString(context.Background(), d1)
	// fmt.Printf("Get by createTestAccountParamString%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ParamId, getData1.ParamId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func updateTestAccountParamString(
	t *testing.T,
	d1 AccountParamStringRequest) model.AccountParamString {

	getData1, err := testQueriesAccount.UpdateAccountParamString(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ParamId, getData1.ParamId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func testDeleteAccountParamString(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteAccountParamString(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountParamStringbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
