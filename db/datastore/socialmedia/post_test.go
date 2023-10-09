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

func TestPost(t *testing.T) {

	// Test Data
	d1 := randomPost()
	d2 := randomPost()

	// Test Create
	CreatedD1 := createTestPost(t, d1)
	CreatedD2 := createTestPost(t, d2)

	// Get Data
	getData1, err1 := testQueriesSocialMedia.GetPost(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.Caption, getData1.Caption)
	require.Equal(t, d1.MessageBody, getData1.MessageBody)
	require.Equal(t, d1.Url, getData1.Url)
	require.Equal(t, d1.ImageUri, getData1.ImageUri)
	require.Equal(t, d1.ThumbnailUri, getData1.ThumbnailUri)
	require.Equal(t, d1.Keywords, getData1.Keywords)
	require.Equal(t, d1.Mood, getData1.Mood)
	require.Equal(t, d1.MoodEmoji, getData1.MoodEmoji)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesSocialMedia.GetPost(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.Caption, getData2.Caption)
	require.Equal(t, d2.MessageBody, getData2.MessageBody)
	require.Equal(t, d2.Url, getData2.Url)
	require.Equal(t, d2.ImageUri, getData2.ImageUri)
	require.Equal(t, d2.ThumbnailUri, getData2.ThumbnailUri)
	require.Equal(t, d2.Keywords, getData2.Keywords)
	require.Equal(t, d2.Mood, getData2.Mood)
	require.Equal(t, d2.MoodEmoji, getData2.MoodEmoji)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesSocialMedia.GetPostbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	updateD2.MessageBody = updateD2.MessageBody + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestPost(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.Caption, updatedD1.Caption)
	require.Equal(t, updateD2.MessageBody, updatedD1.MessageBody)
	require.Equal(t, updateD2.Url, updatedD1.Url)
	require.Equal(t, updateD2.ImageUri, updatedD1.ImageUri)
	require.Equal(t, updateD2.ThumbnailUri, updatedD1.ThumbnailUri)
	require.Equal(t, updateD2.Keywords, updatedD1.Keywords)
	require.Equal(t, updateD2.Mood, updatedD1.Mood)
	require.Equal(t, updateD2.MoodEmoji, updatedD1.MoodEmoji)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListPost(t, ListPostParams{
		UserId: updatedD1.UserId,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeletePost(t, CreatedD1.Uuid)
	testDeletePost(t, CreatedD2.Uuid)
}

func testListPost(t *testing.T, arg ListPostParams) {

	post, err := testQueriesSocialMedia.ListPost(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", post)
	require.NotEmpty(t, post)

}

func randomPost() PostRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "olive.mercado0609@gmail.com")

	arg := PostRequest{
		UserId:       usr.Id,
		Caption:      util.RandomString(10),
		MessageBody:  util.RandomString(10),
		Url:          util.RandomString(10),
		ImageUri:     util.RandomString(10),
		ThumbnailUri: util.RandomString(10),
		Keywords:     []string{util.RandomString(10)},
		Mood:         model.MoodStateHappy,
		MoodEmoji:    util.RandomString(10),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestPost(
	t *testing.T,
	d1 PostRequest) model.Post {

	getData1, err := testQueriesSocialMedia.CreatePost(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.Caption, getData1.Caption)
	require.Equal(t, d1.MessageBody, getData1.MessageBody)
	require.Equal(t, d1.Url, getData1.Url)
	require.Equal(t, d1.ImageUri, getData1.ImageUri)
	require.Equal(t, d1.ThumbnailUri, getData1.ThumbnailUri)
	require.Equal(t, d1.Keywords, getData1.Keywords)
	require.Equal(t, d1.Mood, getData1.Mood)
	require.Equal(t, d1.MoodEmoji, getData1.MoodEmoji)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestPost(
	t *testing.T,
	d1 PostRequest) model.Post {

	getData1, err := testQueriesSocialMedia.UpdatePost(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.Caption, getData1.Caption)
	require.Equal(t, d1.MessageBody, getData1.MessageBody)
	require.Equal(t, d1.Url, getData1.Url)
	require.Equal(t, d1.ImageUri, getData1.ImageUri)
	require.Equal(t, d1.ThumbnailUri, getData1.ThumbnailUri)
	require.Equal(t, d1.Keywords, getData1.Keywords)
	require.Equal(t, d1.Mood, getData1.Mood)
	require.Equal(t, d1.MoodEmoji, getData1.MoodEmoji)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeletePost(t *testing.T, uuid uuid.UUID) {
	err := testQueriesSocialMedia.DeletePost(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesSocialMedia.GetPost(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
