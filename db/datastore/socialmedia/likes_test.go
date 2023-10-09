package db

import (
	"context"
	"database/sql"

	"testing"

	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestLikes(t *testing.T) {

	// Test Data
	d1 := randomLikes()
	d1.Uuid = uuid.MustParse("2af90d74-3bee-48c5-8935-443edafb8f5a")

	d2 := randomLikes()
	d2.Uuid = uuid.MustParse("26dfab18-f80b-46cf-9c54-be79d4fc5d23")

	// Test Create
	CreatedD1 := createTestLikes(t, d1)
	CreatedD2 := createTestLikes(t, d2)

	// Get Data
	getData1, err1 := testQueriesSocialMedia.GetLikes(context.Background(), CreatedD1.Uuid, CreatedD1.UserId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.Mood, getData1.Mood)
	require.Equal(t, d1.DateLiked.Format("2006-01-02"), getData1.DateLiked.Format("2006-01-02"))

	getData2, err2 := testQueriesSocialMedia.GetLikes(context.Background(), CreatedD2.Uuid, CreatedD2.UserId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.Equal(t, d2.Mood, getData2.Mood)
	require.Equal(t, d2.DateLiked.Format("2006-01-02"), getData2.DateLiked.Format("2006-01-02"))

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	updateD2.Mood = updateD2.Mood + 2

	// log.Println(updateD2)
	updatedD1 := updateTestLikes(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.Equal(t, updateD2.Mood, updatedD1.Mood)
	require.Equal(t, updateD2.DateLiked.Format("2006-01-02"), updatedD1.DateLiked.Format("2006-01-02"))

	testListLikes(t, ListLikesParams{
		Uuid:   updatedD1.Uuid,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteLikes(t, CreatedD1.Uuid, CreatedD1.UserId)
	testDeleteLikes(t, CreatedD2.Uuid, CreatedD2.UserId)
}

func testListLikes(t *testing.T, arg ListLikesParams) {

	likes, err := testQueriesSocialMedia.ListLikes(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", likes)
	require.NotEmpty(t, likes)

}

func randomLikes() LikesRequest {

	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")

	arg := LikesRequest{
		UserId:    usr.Id,
		Mood:      util.RandomInt32(1, 100),
		DateLiked: util.RandomDate(),
	}
	return arg
}

func createTestLikes(
	t *testing.T,
	d1 LikesRequest) model.Likes {

	getData1, err := testQueriesSocialMedia.CreateLikes(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.Mood, getData1.Mood)
	require.Equal(t, d1.DateLiked.Format("2006-01-02"), getData1.DateLiked.Format("2006-01-02"))

	return getData1
}

func updateTestLikes(
	t *testing.T,
	d1 LikesRequest) model.Likes {

	getData1, err := testQueriesSocialMedia.UpdateLikes(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.UserId, getData1.UserId)
	require.Equal(t, d1.Mood, getData1.Mood)
	require.Equal(t, d1.DateLiked.Format("2006-01-02"), getData1.DateLiked.Format("2006-01-02"))

	return getData1
}

func testDeleteLikes(t *testing.T, uuid uuid.UUID, userId int64) {
	err := testQueriesSocialMedia.DeleteLikes(context.Background(), uuid, userId)
	require.NoError(t, err)

	ref1, err := testQueriesSocialMedia.GetLikes(context.Background(), uuid, userId)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
