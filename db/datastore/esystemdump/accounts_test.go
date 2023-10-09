package db

import (
	"context"
	"database/sql"
	"fmt"
	model "simplebank/db/datastore/esystemlocal"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccounts(t *testing.T) {

	// Test Data
	d1 := randomAccounts()
	d2 := randomAccounts()
	d2.Acc = d2.Acc + "-"

	err := createTestAccounts(t, d1)
	require.NoError(t, err)

	err = createTestAccounts(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetAccounts(context.Background(), d1.BrCode, d1.Acc)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.Category, getData1.Category)
	require.Equal(t, d1.Type, getData1.Type)
	require.Equal(t, d1.MainCD, getData1.MainCD)
	require.Equal(t, d1.Parent, getData1.Parent)

	getData2, err2 := testQueriesDump.GetAccounts(context.Background(), d2.BrCode, d2.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.Title, getData2.Title)
	require.Equal(t, d2.Category, getData2.Category)
	require.Equal(t, d2.Type, getData2.Type)
	require.Equal(t, d2.MainCD, getData2.MainCD)
	require.Equal(t, d2.Parent, getData2.Parent)

	// Update Data
	updateD2 := d2
	updateD2.Acc = getData2.Acc
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestAccounts(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetAccounts(context.Background(), updateD2.BrCode, updateD2.Acc)
	require.NoError(t, err1)

	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.Equal(t, updateD2.Title, getData1.Title)
	require.Equal(t, updateD2.Category, getData1.Category)
	require.Equal(t, updateD2.Type, getData1.Type)
	require.Equal(t, updateD2.MainCD, getData1.MainCD)
	require.Equal(t, updateD2.Parent, getData1.Parent)

	testListAccounts(t, ListAccountsParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteAccounts(t, d1.BrCode, d1.Acc)
	testDeleteAccounts(t, d2.BrCode, d2.Acc)
}

func testListAccounts(t *testing.T, arg ListAccountsParams) {

	Accounts, err := testQueriesDump.ListAccounts(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Accounts)
	require.NotEmpty(t, Accounts)

}

func randomAccounts() model.Accounts {

	arg := model.Accounts{
		ModCtr:   1,
		BrCode:   "01",
		Acc:      "111",
		Title:    "dddd",
		Category: 1,
		Type:     "DETAIL",
		MainCD:   sql.NullString{String: "5-01-47", Valid: true},
		Parent:   sql.NullString{String: "", Valid: true},
	}
	return arg
}

func createTestAccounts(
	t *testing.T,
	req model.Accounts) error {

	err1 := testQueriesDump.CreateAccounts(context.Background(), req)
	fmt.Printf("Get by createTestAccounts%+v\n", req)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetAccounts(context.Background(), req.BrCode, req.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.Acc, getData.Acc)
	require.Equal(t, req.Title, getData.Title)
	require.Equal(t, req.Category, getData.Category)
	require.Equal(t, req.Type, getData.Type)
	require.Equal(t, req.MainCD, getData.MainCD)
	require.Equal(t, req.Parent, getData.Parent)

	return err2
}

func updateTestAccounts(
	t *testing.T,
	d1 model.Accounts) error {

	err := testQueriesDump.UpdateAccounts(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteAccounts(t *testing.T, brCode string, Acc string) {
	err := testQueriesDump.DeleteAccounts(context.Background(), brCode, Acc)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetAccounts(context.Background(), brCode, Acc)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
