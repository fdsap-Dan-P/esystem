package db

import (
	"context"
	"database/sql"
	"log"

	"testing"

	"simplebank/model"
	"simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestFollower(t *testing.T) {

	// Test Data
	d1 := randomFollower()
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "olive.mercado0609@gmail.com")

	d2 := randomFollower()
	usr2, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")
	d1.UserId = usr.Id
	d1.FollowerId = usr2.Id
	d2.UserId = usr2.Id
	d2.FollowerId = usr.Id
	d2.Uuid = util.ToUUID("40f0a5de-f1df-48a0-a059-164de43ca318")

	// Test Create
	CreatedD1 := createTestFollower(t, d1)
	CreatedD2 := createTestFollower(t, d2)

	// Get Data
	getData1, err1 := testQueriesSocialMedia.GetFollower(context.Background(), CreatedD1.UserId, CreatedD1.FollowerId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.FollowerId, getData1.FollowerId)
	require.Equal(t, d1.DateFollowed.Format("2006-01-02"), getData1.DateFollowed.Format("2006-01-02"))
	require.Equal(t, d1.IsFollower, getData1.IsFollower)

	getData2, err2 := testQueriesSocialMedia.GetFollower(context.Background(), CreatedD2.UserId, CreatedD2.FollowerId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.FollowerId, getData2.FollowerId)
	require.Equal(t, d2.DateFollowed.Format("2006-01-02"), getData2.DateFollowed.Format("2006-01-02"))
	require.Equal(t, d2.IsFollower, getData2.IsFollower)

	// Update Data
	updateD2 := d2
	updateD2.IsFollower = !(updateD2.IsFollower)

	// log.Println(updateD2)
	log.Printf("updateTestFollower %v", updateD2)
	updatedD1 := updateTestFollower(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.FollowerId, updatedD1.FollowerId)
	require.Equal(t, updateD2.DateFollowed.Format("2006-01-02"), updatedD1.DateFollowed.Format("2006-01-02"))
	require.Equal(t, updateD2.IsFollower, updatedD1.IsFollower)

	testListFollower(t, ListFollowerParams{
		UserId: updatedD1.UserId,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteFollower(t, CreatedD1.UserId, CreatedD1.FollowerId)
	testDeleteFollower(t, CreatedD2.UserId, CreatedD2.FollowerId)
}

func testListFollower(t *testing.T, arg ListFollowerParams) {

	follower, err := testQueriesSocialMedia.ListFollower(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", follower)
	require.NotEmpty(t, follower)

}

func randomFollower() FollowerRequest {

	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "Olive")

	arg := FollowerRequest{
		Uuid:         util.ToUUID("40e8f3ce-f77c-4dd3-b424-13220390ef50"),
		FollowerId:   usr.Id,
		DateFollowed: util.RandomDate(),
		IsFollower:   true,
	}
	return arg
}

func createTestFollower(
	t *testing.T,
	d1 FollowerRequest) model.Follower {

	getData1, err := testQueriesSocialMedia.CreateFollower(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.FollowerId, getData1.FollowerId)
	require.Equal(t, d1.DateFollowed.Format("2006-01-02"), getData1.DateFollowed.Format("2006-01-02"))
	require.Equal(t, d1.IsFollower, getData1.IsFollower)

	return getData1
}

func updateTestFollower(
	t *testing.T,
	d1 FollowerRequest) model.Follower {

	getData1, err := testQueriesSocialMedia.UpdateFollower(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.FollowerId, getData1.FollowerId)
	require.Equal(t, d1.DateFollowed.Format("2006-01-02"), getData1.DateFollowed.Format("2006-01-02"))
	require.Equal(t, d1.IsFollower, getData1.IsFollower)

	return getData1
}

func testDeleteFollower(t *testing.T, userId int64, followedId int64) {
	err := testQueriesSocialMedia.DeleteFollower(context.Background(), userId, followedId)
	require.NoError(t, err)

	ref1, err := testQueriesSocialMedia.GetFollower(context.Background(), userId, followedId)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
