package db

import (
	"context"
	"database/sql"
	"encoding/json"

	// "log"
	"testing"

	"simplebank/util"

	common "simplebank/db/common"
	"simplebank/model"

	"github.com/stretchr/testify/require"
)

func TestReferenceType(t *testing.T) {

	// Test Data
	d1 := randomReferenceType()
	createD1 := ReferenceTypeRequest{
		Code:        d1.Code,
		Title:       d1.Title,
		Description: d1.Description,
		OtherInfo:   d1.OtherInfo,
	}
	d2 := randomReferenceType()
	createD2 := ReferenceTypeRequest{
		Code:        d2.Code,
		Title:       d2.Title,
		Description: d2.Description,
		OtherInfo:   d2.OtherInfo,
	}

	// Test Create
	CreatedD1 := createTestReferenceType(t, createD1)
	CreatedD2 := createTestReferenceType(t, createD2)

	// Get Data
	getData1, err1 := testQueriesReference.GetReferenceType(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.Description, getData1.Description)

	getData2, err2 := testQueriesReference.GetReferenceTypeInfo(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Title, getData2.Title)
	require.Equal(t, d2.Description, getData2.Description)

	getData, err := testQueriesReference.GetReferenceTypeInfobyCode(context.Background(), CreatedD1.Code)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Title, getData.Title)
	require.Equal(t, CreatedD1.Description, getData.Description)

	getData3, err3 := testQueriesReference.GetReferenceTypeInfobyTitle(context.Background(), CreatedD1.Title)
	require.NoError(t, err3)
	require.NotEmpty(t, getData3)
	require.Equal(t, CreatedD1.Title, getData3.Title)
	require.Equal(t, CreatedD1.Description, getData3.Description)

	// getData, errtestQueriesReference.GetReferenceTypeInfobyUuId(context.Background(), getData2.Uuid)
	// require.NotEmpty(t, getData)
	// require.Equal(t, d2.Title, getData.Title)
	// require.Equal(t, d2.Description, getData.Description)

	// Update Data
	updateD2 := ReferenceTypeRequest{
		Id:          getData2.Id,
		Code:        getData2.Code,
		Title:       getData2.Title + "Edited",
		Description: getData2.Description,
		OtherInfo:   getData2.OtherInfo,
	}

	// log.Println(updateD2)
	updatedD1 := updateTestReferenceType(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Title, updatedD1.Title)
	require.Equal(t, updateD2.Description, updatedD1.Description)

	// Delete Data
	testDeleteReferenceType(t, getData1.Id)
	testDeleteReferenceType(t, getData2.Id)
}

func TestListReferenceType(t *testing.T) {
	// var lastAccount Account
	// for i := 0; i < 10; i++ {
	// 	lastAccount = createRandomAccount(t)
	// }

	arg := ListReferenceTypeParams{
		Limit:  5,
		Offset: 0,
	}

	referenceType, err := testQueriesReference.ListReferenceType(context.Background(), arg)
	require.NoError(t, err)
	// log.Println(referenceType)
	require.NotEmpty(t, referenceType)

	for _, referenceType := range referenceType {
		require.NotEmpty(t, referenceType)
		// require.Equal(t, lastAccount.Owner, account.Owner)
	}
}

func randomReferenceType() model.ReferenceType {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	arg := model.ReferenceType{
		Code:        sql.NullString(sql.NullString{String: util.RandomString(4), Valid: true}),
		Title:       util.RandomString(20),
		Description: sql.NullString(sql.NullString{String: util.RandomString(100), Valid: true}),
		OtherInfo:   sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestReferenceType(
	t *testing.T,
	createData ReferenceTypeRequest) model.ReferenceType {

	referenceType, err := testQueriesReference.CreateReferenceType(context.Background(), createData)
	require.NoError(t, err)
	require.NotEmpty(t, referenceType)

	require.Equal(t, createData.Code, referenceType.Code)
	require.Equal(t, createData.Title, referenceType.Title)
	require.Equal(t, createData.Description, referenceType.Description)
	return referenceType
}

func updateTestReferenceType(
	t *testing.T,
	updateData ReferenceTypeRequest) model.ReferenceType {

	referenceType, err := testQueriesReference.UpdateReferenceType(context.Background(), updateData)
	require.NoError(t, err)
	require.NotEmpty(t, referenceType)

	require.Equal(t, updateData.Code, referenceType.Code)
	require.Equal(t, updateData.Title, referenceType.Title)
	require.Equal(t, updateData.Description, referenceType.Description)
	return referenceType
}

func testDeleteReferenceType(t *testing.T, id int64) {
	err := testQueriesReference.DeleteReferenceType(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesReference.GetReferenceType(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
