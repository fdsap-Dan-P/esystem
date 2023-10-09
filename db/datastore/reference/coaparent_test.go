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

func TestCoaParent(t *testing.T) {

	// Test Data
	d1 := randomCoaParent()
	d2 := randomCoaParent()

	// Test Create
	CreatedD1 := createTestCoaParent(t, d1)
	CreatedD2 := createTestCoaParent(t, d2)

	// Get Data
	getData1, err1 := testQueriesReference.GetCoaParent(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.CoaSeq, getData1.CoaSeq)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesReference.GetCoaParent(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.CoaSeq, getData2.CoaSeq)
	require.Equal(t, d2.Title, getData2.Title)
	require.Equal(t, d2.ParentId, getData2.ParentId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesReference.GetCoaParentbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)

	fmt.Printf("%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.Title = updateD2.Title + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestCoaParent(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Acc, updatedD1.Acc)
	require.Equal(t, updateD2.CoaSeq, updatedD1.CoaSeq)
	require.Equal(t, updateD2.Title, updatedD1.Title)
	require.Equal(t, updateD2.ParentId, updatedD1.ParentId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteCoaParent(t, getData1.Id)
	testDeleteCoaParent(t, getData2.Id)
}

func TestListCoaParent(t *testing.T) {

	arg := ListCoaParentParams{
		Limit:  5,
		Offset: 0,
	}

	coaParent, err := testQueriesReference.ListCoaParent(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", coaParent)
	require.NotEmpty(t, coaParent)

}

func randomCoaParent() CoaParentRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := CoaParentRequest{
		Acc:      util.RandomString(10),
		CoaSeq:   sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Title:    util.RandomString(10),
		ParentId: sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestCoaParent(
	t *testing.T,
	d1 CoaParentRequest) model.CoaParent {

	getData1, err := testQueriesReference.CreateCoaParent(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.CoaSeq, getData1.CoaSeq)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestCoaParent(
	t *testing.T,
	d1 CoaParentRequest) model.CoaParent {

	getData1, err := testQueriesReference.UpdateCoaParent(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.CoaSeq, getData1.CoaSeq)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteCoaParent(t *testing.T, id int64) {
	err := testQueriesReference.DeleteCoaParent(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesReference.GetCoaParent(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
