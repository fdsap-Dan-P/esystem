package db

import (
	"context"
	"database/sql"
	"simplebank/util"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestReactivateWriteoff(t *testing.T) {

	// Test Data
	d1 := randomReactivateWriteoff()
	d2 := randomReactivateWriteoff()

	err := createTestReactivateWriteoff(t, d1)
	require.NoError(t, err)

	err = createTestReactivateWriteoff(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetReactivateWriteoffbyCID(context.Background(), d1.CID)
	d1.ID = getData1.ID

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CID, getData1.CID)
	require.Equal(t, d1.DeactivateBy, getData1.DeactivateBy)
	require.Equal(t, d1.ReactivateBy, getData1.ReactivateBy)
	require.Equal(t, d1.Status, getData1.Status)
	require.Equal(t, d1.StatusDate, getData1.StatusDate)

	getData2, err2 := testQueriesLocal.GetReactivateWriteoffbyCID(context.Background(), d2.CID)
	d2.ID = getData2.ID
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CID, getData2.CID)
	require.Equal(t, d2.DeactivateBy, getData2.DeactivateBy)
	require.Equal(t, d2.ReactivateBy, getData2.ReactivateBy)
	require.Equal(t, d2.Status, getData2.Status)
	require.Equal(t, d2.StatusDate, getData2.StatusDate)

	// Update Data
	updateD2 := d2
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestReactivateWriteoff(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesLocal.GetReactivateWriteoffbyCID(context.Background(), updateD2.CID)
	require.NoError(t, err1)
	require.Equal(t, updateD2.CID, getData1.CID)
	require.Equal(t, updateD2.DeactivateBy, getData1.DeactivateBy)
	require.Equal(t, updateD2.ReactivateBy, getData1.ReactivateBy)
	require.Equal(t, updateD2.Status, getData1.Status)
	require.Equal(t, updateD2.StatusDate, getData1.StatusDate)

	testListReactivateWriteoff(t, ListReactivateWriteoffParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteReactivateWriteoff(t, d1.ID)
	testDeleteReactivateWriteoff(t, d2.ID)
}

func testListReactivateWriteoff(t *testing.T, arg ListReactivateWriteoffParams) {

	ReactivateWriteoff, err := testQueriesLocal.ListReactivateWriteoff(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", ReactivateWriteoff)
	require.NotEmpty(t, ReactivateWriteoff)

}

func randomReactivateWriteoff() ReactivateWriteoffRequest {

	arg := ReactivateWriteoffRequest{
		CID:          int64(19858200),
		DeactivateBy: util.SetNullString("sa"),
		ReactivateBy: util.SetNullString("sa"),
		Status:       1,
		StatusDate:   util.SetDate("2021-01-01"),
	}
	return arg
}

func createTestReactivateWriteoff(
	t *testing.T,
	req ReactivateWriteoffRequest) error {

	err1 := testQueriesLocal.CreateReactivateWriteoff(context.Background(), req)
	// fmt.Printf("Get by createTestReactivateWriteoff%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetReactivateWriteoffbyCID(context.Background(), req.CID)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.CID, getData.CID)
	require.Equal(t, req.DeactivateBy, getData.DeactivateBy)
	require.Equal(t, req.ReactivateBy, getData.ReactivateBy)
	require.Equal(t, req.Status, getData.Status)
	require.Equal(t, req.StatusDate, getData.StatusDate)

	return err2
}

func updateTestReactivateWriteoff(
	t *testing.T,
	d1 ReactivateWriteoffRequest) error {

	err := testQueriesLocal.UpdateReactivateWriteoff(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteReactivateWriteoff(t *testing.T, id int64) {
	err := testQueriesLocal.DeleteReactivateWriteoff(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetReactivateWriteoffbyCID(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
