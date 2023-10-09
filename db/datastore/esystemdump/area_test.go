package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"
	"simplebank/util"

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
	getData1, err1 := testQueriesDump.GetArea(context.Background(), d1.BrCode, d1.AreaCode)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AreaCode, getData1.AreaCode)
	require.Equal(t, d1.Area, getData1.Area)

	getData2, err2 := testQueriesDump.GetArea(context.Background(), d2.BrCode, d2.AreaCode)
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

	getData1, err1 = testQueriesDump.GetArea(context.Background(), updateD2.BrCode, updateD2.AreaCode)
	require.NoError(t, err1)

	require.Equal(t, updateD2.AreaCode, getData1.AreaCode)
	require.Equal(t, updateD2.Area, getData1.Area)

	testListArea(t, ListAreaParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteArea(t, d1.BrCode, d1.AreaCode)
	testDeleteArea(t, d2.BrCode, d2.AreaCode)
}

func testListArea(t *testing.T, arg ListAreaParams) {

	area, err := testQueriesDump.ListArea(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", area)
	require.NotEmpty(t, area)

}

func randomArea() model.Area {

	arg := model.Area{
		ModCtr:   1,
		BrCode:   "01",
		AreaCode: 111,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Area: util.SetNullString("Test Area"),
	}
	return arg
}

func createTestArea(
	t *testing.T,
	req model.Area) error {

	err1 := testQueriesDump.CreateArea(context.Background(), req)
	// fmt.Printf("Get by createTestArea%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetArea(context.Background(), req.BrCode, req.AreaCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.AreaCode, getData.AreaCode)
	require.Equal(t, req.Area, getData.Area)

	return err2
}

func updateTestArea(
	t *testing.T,
	d1 model.Area) error {

	err := testQueriesDump.UpdateArea(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteArea(t *testing.T, brCode string, areaCode int64) {
	err := testQueriesDump.DeleteArea(context.Background(), brCode, areaCode)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetArea(context.Background(), brCode, areaCode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
