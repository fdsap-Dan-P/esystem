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

func TestDocumentAccess(t *testing.T) {

	// Test Data
	d1 := RandomDocumentAccess()
	d2 := RandomDocumentAccess()
	d2.Uuid = uuid.MustParse("6fedf148-efac-4677-9f4d-dc4e644a7f91")

	// Test Create
	CreatedD1 := createTestDocumentAccess(t, d1)
	CreatedD2 := createTestDocumentAccess(t, d2)

	// Get Data
	getData1, err1 := testQueriesDocument.GetDocumentAccessbyUuid(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.DocumentId, getData1.DocumentId)
	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.AccessCode, getData1.AccessCode)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesDocument.GetDocumentAccessbyUuid(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.DocumentId, getData2.DocumentId)
	require.Equal(t, d2.RoleId, getData2.RoleId)
	require.Equal(t, d2.AccessCode, getData2.AccessCode)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Description.String = updateD2.Description.String + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestDocumentAccess(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.DocumentId, updatedD1.DocumentId)
	require.Equal(t, updateD2.RoleId, updatedD1.RoleId)
	require.Equal(t, updateD2.AccessCode, updatedD1.AccessCode)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListDocumentAccess(t, ListDocumentAccessParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteDocumentAccess(t, getData1.Uuid)
	testDeleteDocumentAccess(t, getData2.Uuid)
}

func testListDocumentAccess(t *testing.T, arg ListDocumentAccessParams) {

	documentAccess, err := testQueriesDocument.ListDocumentAccess(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", documentAccess)
	require.NotEmpty(t, documentAccess)

}

func RandomDocumentAccess() DocumentAccessRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	doc, _ := testQueriesDocument.CreateDocument(context.Background(), RandomDocument())
	role, _ := testQueriesAccess.GetAccessRolebyName(context.Background(), "Admin")

	arg := DocumentAccessRequest{
		Uuid:       uuid.MustParse("3ec19a37-ae17-413a-a59c-000e69f61b8f"),
		DocumentId: doc.Id,
		RoleId:     role.Id,
		AccessCode: "R",
		OtherInfo:  sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestDocumentAccess(
	t *testing.T,
	d1 DocumentAccessRequest) model.DocumentAccess {

	getData1, err := testQueriesDocument.CreateDocumentAccess(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.DocumentId, getData1.DocumentId)
	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.AccessCode, getData1.AccessCode)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestDocumentAccess(
	t *testing.T,
	d1 DocumentAccessRequest) model.DocumentAccess {

	getData1, err := testQueriesDocument.UpdateDocumentAccess(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.DocumentId, getData1.DocumentId)
	require.Equal(t, d1.RoleId, getData1.RoleId)
	require.Equal(t, d1.AccessCode, getData1.AccessCode)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteDocumentAccess(t *testing.T, uuid uuid.UUID) {
	err := testQueriesDocument.DeleteDocumentAccess(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesDocument.GetDocumentAccessbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
