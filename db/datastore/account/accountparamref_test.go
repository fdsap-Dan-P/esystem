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

func TestAccountParamRef(t *testing.T) {

	// Test Data
	d1 := randomAccountParamRef(t)
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Parameter", 0, "Interest")
	d1.ItemId = item.Id

	d2 := randomAccountParamRef(t)
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Parameter", 0, "Charges")
	d2.ItemId = item.Id

	// Test Create
	CreatedD1 := createTestAccountParamRef(t, d1)
	CreatedD2 := createTestAccountParamRef(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountParamRef(context.Background(), CreatedD1.ParamId, CreatedD1.ItemId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.ParamId, getData1.ParamId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.RefId, getData1.RefId)

	getData2, err2 := testQueriesAccount.GetAccountParamRef(context.Background(), CreatedD2.ParamId, CreatedD2.ItemId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.ParamId, getData2.ParamId)
	require.Equal(t, d2.ItemId, getData2.ItemId)
	require.Equal(t, d2.RefId, getData2.RefId)

	getData, err := testQueriesAccount.GetAccountParamRefbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountParamRef(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.ParamId, updatedD1.ParamId)
	require.Equal(t, updateD2.ItemId, updatedD1.ItemId)
	require.Equal(t, updateD2.RefId, updatedD1.RefId)

	testListAccountParamRef(t, ListAccountParamRefParams{
		ParamId: updatedD1.ParamId,
		Limit:   5,
		Offset:  0,
	})

	// Delete Data
	testDeleteAccountParamRef(t, CreatedD1.Uuid)
	testDeleteAccountParamRef(t, CreatedD2.Uuid)
}

func testListAccountParamRef(t *testing.T, arg ListAccountParamRefParams) {

	accountParamRef, err := testQueriesAccount.ListAccountParamRef(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountParamRef)
	require.NotEmpty(t, accountParamRef)

}

func randomAccountParamRef(t *testing.T) AccountParamRefRequest {

	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeetype", 0, "Employed")
	accParam := createTestAccountParam(t, randomAccountParam())

	arg := AccountParamRefRequest{
		ParamId: accParam.Id,
		// ItemId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		RefId: typ.Id,
	}
	return arg
}

func createTestAccountParamRef(
	t *testing.T,
	d1 AccountParamRefRequest) model.AccountParamRef {

	getData1, err := testQueriesAccount.CreateAccountParamRef(context.Background(), d1)
	// fmt.Printf("Get by createTestAccountParamRef%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ParamId, getData1.ParamId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func updateTestAccountParamRef(
	t *testing.T,
	d1 AccountParamRefRequest) model.AccountParamRef {

	getData1, err := testQueriesAccount.UpdateAccountParamRef(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ParamId, getData1.ParamId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func testDeleteAccountParamRef(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteAccountParamRef(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountParamRefbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
