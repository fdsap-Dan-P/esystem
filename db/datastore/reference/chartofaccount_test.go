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

func TestChartofAccount(t *testing.T) {

	// Test Data
	d1 := randomChartofAccount()
	d2 := randomChartofAccount()

	// Test Create
	CreatedD1 := createTestChartofAccount(t, d1)
	CreatedD2 := createTestChartofAccount(t, d2)

	// Get Data
	getData1, err1 := testQueriesReference.GetChartofAccount(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.Active, getData1.Active)
	require.Equal(t, d1.ContraAccount, getData1.ContraAccount)
	require.Equal(t, d1.NormalBalance, getData1.NormalBalance)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.Equal(t, d1.ShortName, getData1.ShortName)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesReference.GetChartofAccount(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.Active, getData2.Active)
	require.Equal(t, d2.ContraAccount, getData2.ContraAccount)
	require.Equal(t, d2.NormalBalance, getData2.NormalBalance)
	require.Equal(t, d2.Title, getData2.Title)
	require.Equal(t, d2.ParentId, getData2.ParentId)
	require.Equal(t, d2.ShortName, getData2.ShortName)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesReference.GetChartofAccountbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("%+v\n", getData)

	getData, err = testQueriesReference.GetChartofAccountbyTitle(context.Background(), CreatedD1.Title)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.Title = updateD2.Title + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestChartofAccount(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Acc, updatedD1.Acc)
	require.Equal(t, updateD2.Active, updatedD1.Active)
	require.Equal(t, updateD2.ContraAccount, updatedD1.ContraAccount)
	require.Equal(t, updateD2.NormalBalance, updatedD1.NormalBalance)
	require.Equal(t, updateD2.Title, updatedD1.Title)
	require.Equal(t, updateD2.ParentId, updatedD1.ParentId)
	require.Equal(t, updateD2.ShortName, updatedD1.ShortName)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteChartofAccount(t, getData1.Id)
	testDeleteChartofAccount(t, getData2.Id)
}

func TestListChartofAccount(t *testing.T) {

	arg := ListChartofAccountParams{
		Limit:  5,
		Offset: 0,
	}

	chartofAccount, err := testQueriesReference.ListChartofAccount(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", chartofAccount)
	require.NotEmpty(t, chartofAccount)

}

func randomChartofAccount() ChartofAccountRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := ChartofAccountRequest{
		Acc:           util.RandomString(10),
		Active:        true,
		ContraAccount: true,
		NormalBalance: true,
		Title:         util.RandomString(10),
		ParentId:      util.RandomInt(1, 100),
		ShortName:     util.RandomString(10),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestChartofAccount(
	t *testing.T,
	d1 ChartofAccountRequest) model.ChartofAccount {

	getData1, err := testQueriesReference.CreateChartofAccount(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.Active, getData1.Active)
	require.Equal(t, d1.ContraAccount, getData1.ContraAccount)
	require.Equal(t, d1.NormalBalance, getData1.NormalBalance)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.Equal(t, d1.ShortName, getData1.ShortName)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestChartofAccount(
	t *testing.T,
	d1 ChartofAccountRequest) model.ChartofAccount {

	getData1, err := testQueriesReference.UpdateChartofAccount(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.Active, getData1.Active)
	require.Equal(t, d1.ContraAccount, getData1.ContraAccount)
	require.Equal(t, d1.NormalBalance, getData1.NormalBalance)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.Equal(t, d1.ShortName, getData1.ShortName)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteChartofAccount(t *testing.T, id int64) {
	err := testQueriesReference.DeleteChartofAccount(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesReference.GetChartofAccount(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
