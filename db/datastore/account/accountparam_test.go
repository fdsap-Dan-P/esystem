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

	"github.com/stretchr/testify/require"
)

func TestAccountParam(t *testing.T) {

	// Test Data
	d1 := randomAccountParam()
	d2 := randomAccountParam()
	d2.DateImplemented = util.SetDate("2021-01-02")
	d2.Uuid = util.SetUUID("43de929e-65e4-4db3-9c40-6e1a6d12cc8b")

	// Test Create
	CreatedD1 := createTestAccountParam(t, d1)
	CreatedD2 := createTestAccountParam(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountParam(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.DateImplemented.Format("2006-01-02"), getData1.DateImplemented.Format("2006-01-02"))
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetAccountParam(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Uuid, getData2.Uuid)
	require.Equal(t, d2.AccountTypeId, getData2.AccountTypeId)
	require.Equal(t, d2.DateImplemented.Format("2006-01-02"), getData2.DateImplemented.Format("2006-01-02"))
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetAccountParambyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// getData, err = testQueriesAccount.GetAccountParambyName(context.Background(), CreatedD1.Name)
	// require.NoError(t, err)
	// require.NotEmpty(t, getData)
	// require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	// fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.Uuid = getData2.Uuid
	updateD2.DateImplemented = util.RandomDate()

	// log.Println(updateD2)
	updatedD1 := updateTestAccountParam(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Uuid, updatedD1.Uuid)
	require.Equal(t, updateD2.AccountTypeId, updatedD1.AccountTypeId)
	require.Equal(t, updateD2.DateImplemented.Format("2006-01-02"), updatedD1.DateImplemented.Format("2006-01-02"))
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListAccountParam(t, ListAccountParamParams{
		AccountTypeId: CreatedD1.AccountTypeId,
		Limit:         5,
		Offset:        0,
	})
	// Delete Data
	testDeleteAccountParam(t, CreatedD1.Id)
	testDeleteAccountParam(t, CreatedD2.Id)
}

func testListAccountParam(t *testing.T, arg ListAccountParamParams) {

	accountParam, err := testQueriesAccount.ListAccountParam(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountParam)
	require.NotEmpty(t, accountParam)

}

func randomAccountParam() AccountParamRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	// p := util.RandomAccountType()
	p := "MF - Sikap"
	accType, _ := testQueriesAccount.GetAccountTypebyName(context.Background(), p)

	arg := AccountParamRequest{
		Uuid:            util.SetUUID("013b1beb-1635-4a1d-a187-bb192e8aeaf2"),
		AccountTypeId:   accType.Id,
		DateImplemented: util.SetDate("2021-01-01"),
		OtherInfo:       sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	log.Printf("randomAccountParam: %v", arg)
	return arg
}

func createTestAccountParam(
	t *testing.T,
	d1 AccountParamRequest) model.AccountParam {

	getData1, err := testQueriesAccount.CreateAccountParam(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.DateImplemented.Format("2006-01-02"), getData1.DateImplemented.Format("2006-01-02"))
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)
	return getData1
}

func updateTestAccountParam(
	t *testing.T,
	d1 AccountParamRequest) model.AccountParam {

	getData1, err := testQueriesAccount.UpdateAccountParam(context.Background(), d1)
	log.Printf("updateTestAccountParam: d1 %v", d1)
	log.Printf("updateTestAccountParam: d1 %v", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.DateImplemented.Format("2006-01-02"), getData1.DateImplemented.Format("2006-01-02"))
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAccountParam(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteAccountParam(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountParam(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
