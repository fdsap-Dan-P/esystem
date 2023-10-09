package db

import (
	"context"
	"database/sql"

	"testing"

	"encoding/json"
	common "simplebank/db/common"

	"simplebank/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestDocumentUser(t *testing.T) {

	// Test Data
	d1 := RandomDocumentUser()
	d2 := RandomDocumentUser()
	d2.Uuid = uuid.MustParse("19d03a87-8a08-4065-8841-012c9fe8af40")

	// Test Create
	CreatedD1 := createTestDocumentUser(t, d1)
	CreatedD2 := createTestDocumentUser(t, d2)

	// Get Data
	getData1, err1 := testQueriesDocument.GetDocumentUserbyUuid(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.DocumentId, getData1.DocumentId)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.AccessCode, getData1.AccessCode)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesDocument.GetDocumentUserbyUuid(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.DocumentId, getData2.DocumentId)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.AccessCode, getData2.AccessCode)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Description.String = updateD2.Description.String + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestDocumentUser(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.DocumentId, updatedD1.DocumentId)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.AccessCode, updatedD1.AccessCode)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListDocumentUser(t, ListDocumentUserParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteDocumentUser(t, getData1.Uuid)
	testDeleteDocumentUser(t, getData2.Uuid)
}

func testListDocumentUser(t *testing.T, arg ListDocumentUserParams) {

	documentUser, err := testQueriesDocument.ListDocumentUser(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", documentUser)
	require.NotEmpty(t, documentUser)

}

func RandomDocumentUser() DocumentUserRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	doc, _ := testQueriesDocument.CreateDocument(context.Background(), RandomDocument())
	role, _ := testQueriesAccess.GetAccessRolebyName(context.Background(), "Admin")

	arg := DocumentUserRequest{
		Uuid:       uuid.MustParse("178ac86c-5d84-4fff-9cd9-e8cc681bfd7c"),
		DocumentId: doc.Id,
		UserId:     role.Id,
		AccessCode: "R",
		OtherInfo:  sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestDocumentUser(
	t *testing.T,
	d1 DocumentUserRequest) model.DocumentUser {

	getData1, err := testQueriesDocument.CreateDocumentUser(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.DocumentId, getData1.DocumentId)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.AccessCode, getData1.AccessCode)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestDocumentUser(
	t *testing.T,
	d1 DocumentUserRequest) model.DocumentUser {

	getData1, err := testQueriesDocument.UpdateDocumentUser(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.DocumentId, getData1.DocumentId)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.AccessCode, getData1.AccessCode)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteDocumentUser(t *testing.T, uuid uuid.UUID) {
	err := testQueriesDocument.DeleteDocumentUser(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesDocument.GetDocumentUserbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
