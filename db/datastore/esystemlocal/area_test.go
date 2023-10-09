package db

import (
	"context"
	"database/sql"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestArea(t *testing.T) {

	// Test Data
	d1 := randomArea()
	d2 := randomArea()
	d2.AreaCode = d2.AreaCode + 1

	err := createTestArea(t, d1)
	require.NoError(t, err)

	err = createTestArea(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetArea(context.Background(), d1.AreaCode)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AreaCode, getData1.AreaCode)
	require.Equal(t, d1.Area, getData1.Area)

	getData2, err2 := testQueriesLocal.GetArea(context.Background(), d2.AreaCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AreaCode, getData2.AreaCode)
	require.Equal(t, d2.Area, getData2.Area)

	// Update Data
	updateD2 := d2
	updateD2.AreaCode = getData2.AreaCode
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestArea(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesLocal.GetArea(context.Background(), updateD2.AreaCode)
	require.NoError(t, err1)

	require.Equal(t, updateD2.AreaCode, getData1.AreaCode)
	require.Equal(t, updateD2.Area, getData1.Area)

	testListArea(t, ListAreaParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteArea(t, d1.AreaCode)
	testDeleteArea(t, d2.AreaCode)
}

func testListArea(t *testing.T, arg ListAreaParams) {

	area, err := testQueriesLocal.ListArea(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", area)
	require.NotEmpty(t, area)

}

func randomArea() AreaRequest {

	arg := AreaRequest{
		AreaCode: 111,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Area: sql.NullString{String: "Test Area", Valid: true},
	}
	return arg
}

func createTestArea(
	t *testing.T,
	req AreaRequest) error {

	err1 := testQueriesLocal.CreateArea(context.Background(), req)
	// fmt.Printf("Get by createTestArea%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetArea(context.Background(), req.AreaCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.AreaCode, getData.AreaCode)
	require.Equal(t, req.Area, getData.Area)

	return err2
}

func updateTestArea(
	t *testing.T,
	d1 AreaRequest) error {

	err := testQueriesLocal.UpdateArea(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteArea(t *testing.T, areaCode int64) {
	err := testQueriesLocal.DeleteArea(context.Background(), areaCode)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetArea(context.Background(), areaCode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
