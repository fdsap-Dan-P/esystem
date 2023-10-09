package db

import (
	"context"
	"database/sql"

	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestLnChrgData(t *testing.T) {

	// Test Data
	d1 := randomLnChrgData()
	d2 := randomLnChrgData()
	d2.Acc = "0101-4041-0157454"

	err := createTestLnChrgData(t, d1)
	require.NoError(t, err)

	err = createTestLnChrgData(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetLnChrgData(context.Background(), d1.Acc, d1.ChrgCode, d1.RefAcc.String)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.ChrgCode, getData1.ChrgCode)
	require.Equal(t, d1.RefAcc, getData1.RefAcc)
	// require.True(t, d1.ChrAmnt.Equal(getData1.ChrAmnt))

	getData2, err2 := testQueriesLocal.GetLnChrgData(context.Background(), d2.Acc, d2.ChrgCode, d2.RefAcc.String)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.ChrgCode, getData2.ChrgCode)
	require.Equal(t, d2.RefAcc, getData2.RefAcc)
	// require.True(t, d2.ChrAmnt.Equal(getData2.ChrAmnt))

	// Update Data
	updateD2 := d2
	updateD2.ChrAmnt.Add(decimal.NewFromInt(10))
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestLnChrgData(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesLocal.GetLnChrgData(context.Background(), updateD2.Acc, updateD2.ChrgCode, updateD2.RefAcc.String)
	require.NoError(t, err1)

	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.Equal(t, updateD2.ChrgCode, getData1.ChrgCode)
	require.Equal(t, updateD2.RefAcc, getData1.RefAcc)
	// require.True(t, updateD2.ChrAmnt.Equal(getData1.ChrAmnt))

	testListLnChrgData(t, ListLnChrgDataParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteLnChrgData(t, d1.Acc, d1.ChrgCode, d1.RefAcc.String)
	testDeleteLnChrgData(t, d2.Acc, d2.ChrgCode, d2.RefAcc.String)
}

func testListLnChrgData(t *testing.T, arg ListLnChrgDataParams) {

	LnChrgData, err := testQueriesLocal.ListLnChrgData(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", LnChrgData)
	require.NotEmpty(t, LnChrgData)

}

func randomLnChrgData() LnChrgDataRequest {

	arg := LnChrgDataRequest{
		Acc:      "0101-4041-0157455",
		ChrgCode: 18,
		RefAcc:   sql.NullString{String: "dsdff", Valid: true},
		ChrAmnt:  decimal.NewFromInt(100),
	}
	return arg
}

func createTestLnChrgData(
	t *testing.T,
	req LnChrgDataRequest) error {

	err1 := testQueriesLocal.CreateLnChrgData(context.Background(), req)
	// fmt.Printf("Get by createTestLnChrgData%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetLnChrgData(context.Background(), req.Acc, req.ChrgCode, req.RefAcc.String)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.Acc, getData.Acc)
	require.Equal(t, req.ChrgCode, getData.ChrgCode)
	require.Equal(t, req.RefAcc, getData.RefAcc)
	// require.True(t, req.ChrAmnt.Equal(getData.ChrAmnt))

	return err2
}

func updateTestLnChrgData(
	t *testing.T,
	d1 LnChrgDataRequest) error {

	err := testQueriesLocal.UpdateLnChrgData(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteLnChrgData(t *testing.T, acc string, code int64, refAcc string) {
	err := testQueriesLocal.DeleteLnChrgData(context.Background(), acc, code, acc)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetLnChrgData(context.Background(), acc, code, acc)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
