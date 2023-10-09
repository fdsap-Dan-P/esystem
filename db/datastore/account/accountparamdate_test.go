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

func TestAccountParamDate(t *testing.T) {

	// Test Data
	d1 := randomAccountParamDate(t)
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Parameter", 0, "Interest")
	d1.ItemId = item.Id

	d2 := randomAccountParamDate(t)
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Parameter", 0, "Charges")
	d2.ItemId = item.Id

	// Test Create
	CreatedD1 := createTestAccountParamDate(t, d1)
	CreatedD2 := createTestAccountParamDate(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountParamDate(context.Background(), CreatedD1.ParamId, CreatedD1.ItemId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.ParamId, getData1.ParamId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	getData2, err2 := testQueriesAccount.GetAccountParamDate(context.Background(), CreatedD2.ParamId, CreatedD2.ItemId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.ParamId, getData2.ParamId)
	require.Equal(t, d2.ItemId, getData2.ItemId)
	require.Equal(t, d2.Value.Format("2006-01-02"), getData2.Value.Format("2006-01-02"))
	require.Equal(t, d2.Value2.Format("2006-01-02"), getData2.Value2.Format("2006-01-02"))

	getData, err := testQueriesAccount.GetAccountParamDatebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountParamDate(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.ParamId, updatedD1.ParamId)
	require.Equal(t, updateD2.ItemId, updatedD1.ItemId)
	require.Equal(t, updateD2.Value.Format("2006-01-02"), updatedD1.Value.Format("2006-01-02"))
	require.Equal(t, updateD2.Value2.Format("2006-01-02"), updatedD1.Value2.Format("2006-01-02"))

	testListAccountParamDate(t, ListAccountParamDateParams{
		ParamId: updatedD1.ParamId,
		Limit:   5,
		Offset:  0,
	})

	// Delete Data
	testDeleteAccountParamDate(t, CreatedD1.Uuid)
	testDeleteAccountParamDate(t, CreatedD2.Uuid)
}

func testListAccountParamDate(t *testing.T, arg ListAccountParamDateParams) {

	accountParamDate, err := testQueriesAccount.ListAccountParamDate(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountParamDate)
	require.NotEmpty(t, accountParamDate)
}

func randomAccountParamDate(t *testing.T) AccountParamDateRequest {

	accParam := createTestAccountParam(t, randomAccountParam())
	arg := AccountParamDateRequest{
		ParamId: accParam.Id,
		// ItemId: sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomDate(),
		Value2: util.RandomDate(),
	}
	return arg
}

func createTestAccountParamDate(
	t *testing.T,
	d1 AccountParamDateRequest) model.AccountParamDate {

	getData1, err := testQueriesAccount.CreateAccountParamDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ParamId, getData1.ParamId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func updateTestAccountParamDate(
	t *testing.T,
	d1 AccountParamDateRequest) model.AccountParamDate {

	getData1, err := testQueriesAccount.UpdateAccountParamDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ParamId, getData1.ParamId)
	require.Equal(t, d1.ItemId, getData1.ItemId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func testDeleteAccountParamDate(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteAccountParamDate(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountParamDatebyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
