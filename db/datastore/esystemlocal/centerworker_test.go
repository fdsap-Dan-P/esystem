package db

import (
	"context"
	"database/sql"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestCenterWorker(t *testing.T) {

	// Test Data
	d1 := randomCenterWorker()
	d2 := randomCenterWorker()
	d2.AOID = d2.AOID + 1

	err := createTestCenterWorker(t, d1)
	require.NoError(t, err)

	err = createTestCenterWorker(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetCenterWorker(context.Background(), d1.AOID)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AOID, getData1.AOID)
	require.Equal(t, d1.Lname, getData1.Lname)
	require.Equal(t, d1.FName, getData1.FName)
	require.Equal(t, d1.Mname, getData1.Mname)
	require.Equal(t, d1.PhoneNumber, getData1.PhoneNumber)

	getData2, err2 := testQueriesLocal.GetCenterWorker(context.Background(), d2.AOID)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AOID, getData2.AOID)
	require.Equal(t, d2.AOID, getData2.AOID)
	require.Equal(t, d2.Lname, getData2.Lname)
	require.Equal(t, d2.FName, getData2.FName)
	require.Equal(t, d2.Mname, getData2.Mname)
	require.Equal(t, d2.PhoneNumber, getData2.PhoneNumber)

	// Update Data
	updateD2 := d2
	updateD2.AOID = getData2.AOID
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestCenterWorker(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesLocal.GetCenterWorker(context.Background(), updateD2.AOID)
	require.NoError(t, err1)

	require.Equal(t, updateD2.AOID, getData1.AOID)
	require.Equal(t, updateD2.Lname, getData1.Lname)
	require.Equal(t, updateD2.FName, getData1.FName)
	require.Equal(t, updateD2.Mname, getData1.Mname)
	require.Equal(t, updateD2.PhoneNumber, getData1.PhoneNumber)

	testListCenterWorker(t, ListCenterWorkerParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteCenterWorker(t, d1.AOID)
	testDeleteCenterWorker(t, d2.AOID)
}

func testListCenterWorker(t *testing.T, arg ListCenterWorkerParams) {

	CenterWorker, err := testQueriesLocal.ListCenterWorker(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", CenterWorker)
	require.NotEmpty(t, CenterWorker)

}

func randomCenterWorker() CenterWorkerRequest {

	arg := CenterWorkerRequest{
		AOID:        69,
		Lname:       sql.NullString{String: "lname", Valid: true},
		FName:       sql.NullString{String: "fname", Valid: true},
		Mname:       sql.NullString{String: "d", Valid: true},
		PhoneNumber: sql.NullString{String: "phone", Valid: true},
	}
	return arg
}

func createTestCenterWorker(
	t *testing.T,
	req CenterWorkerRequest) error {

	err1 := testQueriesLocal.CreateCenterWorker(context.Background(), req)
	// fmt.Printf("Get by createTestCenterWorker%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetCenterWorker(context.Background(), req.AOID)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.AOID, getData.AOID)
	require.Equal(t, req.Lname, getData.Lname)
	require.Equal(t, req.FName, getData.FName)
	require.Equal(t, req.Mname, getData.Mname)
	require.Equal(t, req.PhoneNumber, getData.PhoneNumber)

	return err2
}

func updateTestCenterWorker(
	t *testing.T,
	d1 CenterWorkerRequest) error {

	err := testQueriesLocal.UpdateCenterWorker(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteCenterWorker(t *testing.T, AOID int64) {
	err := testQueriesLocal.DeleteCenterWorker(context.Background(), AOID)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetCenterWorker(context.Background(), AOID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
