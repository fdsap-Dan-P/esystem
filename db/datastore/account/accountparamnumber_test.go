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

func TestAccountParamNumber(t *testing.T) {

	// Test Data
	d1 := randomAccountParamNumber(t)
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Parameter", 0, "Interest")
	d1.ItemId = item.Id

	d2 := randomAccountParamNumber(t)
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Parameter", 0, "Charges")
	d2.ItemId = item.Id

	// Test Create
	CreatedD1 := createTestAccountParamNumber(t, d1)
	CreatedD2 := createTestAccountParamNumber(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountParamNumber(context.Background(), CreatedD1.ParamId, CreatedD1.ItemId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.ParamId, getData1.ParamId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	getData2, err2 := testQueriesAccount.GetAccountParamNumber(context.Background(), CreatedD2.ParamId, CreatedD2.ItemId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.ParamId, getData2.ParamId)
	require.Equal(t, d2.ItemId, getData2.ItemId)
	require.Equal(t, d2.Value.String(), getData2.Value.String())
	require.Equal(t, d2.Value2.String(), getData2.Value2.String())

	getData, err := testQueriesAccount.GetAccountParamNumberbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountParamNumber(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.ParamId, updatedD1.ParamId)
	require.Equal(t, updateD2.ItemId, updatedD1.ItemId)
	require.Equal(t, updateD2.Value.String(), updatedD1.Value.String())
	require.Equal(t, updateD2.Value2.String(), updatedD1.Value2.String())

	testListAccountParamNumber(t, ListAccountParamNumberParams{
		ParamId: updatedD1.ParamId,
		Limit:   5,
		Offset:  0,
	})

	// Delete Data
	testDeleteAccountParamNumber(t, CreatedD1.Uuid)
	testDeleteAccountParamNumber(t, CreatedD2.Uuid)
}

func testListAccountParamNumber(t *testing.T, arg ListAccountParamNumberParams) {

	accountParamNumber, err := testQueriesAccount.ListAccountParamNumber(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountParamNumber)
	require.NotEmpty(t, accountParamNumber)

}

func randomAccountParamNumber(t *testing.T) AccountParamNumberRequest {

	accParam := createTestAccountParam(t, randomAccountParam())
	arg := AccountParamNumberRequest{
		ParamId: accParam.Id,
		// ItemId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomMoney(),
		Value2: util.RandomMoney(),
	}
	return arg
}

func createTestAccountParamNumber(
	t *testing.T,
	d1 AccountParamNumberRequest) model.AccountParamNumber {

	getData1, err := testQueriesAccount.CreateAccountParamNumber(context.Background(), d1)
	// fmt.Printf("Get by createTestAccountParamNumber%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ParamId, getData1.ParamId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func updateTestAccountParamNumber(
	t *testing.T,
	d1 AccountParamNumberRequest) model.AccountParamNumber {

	getData1, err := testQueriesAccount.UpdateAccountParamNumber(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ParamId, getData1.ParamId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func testDeleteAccountParamNumber(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteAccountParamNumber(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountParamNumberbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
