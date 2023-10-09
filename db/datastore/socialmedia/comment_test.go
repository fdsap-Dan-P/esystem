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

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestComment(t *testing.T) {

	// Test Data
	d1 := randomComment()
	d1.RecordUuid = uuid.MustParse("2af90d74-3bee-48c5-8935-443edafb8f5a")

	d2 := randomComment()
	d2.RecordUuid = uuid.MustParse("26dfab18-f80b-46cf-9c54-be79d4fc5d23")

	// Test Create
	CreatedD1 := createTestComment(t, d1)
	CreatedD2 := createTestComment(t, d2)

	// Get Data
	getData1, err1 := testQueriesSocialMedia.GetComment(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.RecordUuid, getData1.RecordUuid)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.Comment, getData1.Comment)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesSocialMedia.GetComment(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.RecordUuid, getData2.RecordUuid)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.Comment, getData2.Comment)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesSocialMedia.GetCommentbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	updateD2.Comment = updateD2.Comment + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestComment(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.RecordUuid, updatedD1.RecordUuid)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.Comment, updatedD1.Comment)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListComment(t, ListCommentParams{
		RecordUuid: updatedD1.RecordUuid,
		Limit:      5,
		Offset:     0,
	})

	// Delete Data
	testDeleteComment(t, CreatedD1.Uuid)
	testDeleteComment(t, CreatedD2.Uuid)
}

func testListComment(t *testing.T, arg ListCommentParams) {

	comment, err := testQueriesSocialMedia.ListComment(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", comment)
	require.NotEmpty(t, comment)

}

func randomComment() CommentRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "olive.mercado0609@gmail.com")

	arg := CommentRequest{
		// RecordUuid: uuid.Newutil.Random(),
		UserId:  usr.Id,
		Comment: util.RandomString(10),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestComment(
	t *testing.T,
	d1 CommentRequest) model.Comment {

	getData1, err := testQueriesSocialMedia.CreateComment(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RecordUuid, getData1.RecordUuid)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.Comment, getData1.Comment)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestComment(
	t *testing.T,
	d1 CommentRequest) model.Comment {

	getData1, err := testQueriesSocialMedia.UpdateComment(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.RecordUuid, getData1.RecordUuid)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.Comment, getData1.Comment)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteComment(t *testing.T, uuid uuid.UUID) {
	err := testQueriesSocialMedia.DeleteComment(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesSocialMedia.GetComment(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
