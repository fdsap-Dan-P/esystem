package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"

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
	getData1, err1 := testQueriesDump.GetLnChrgData(context.Background(), d1.BrCode, d1.Acc, d1.ChrgCode, d1.RefAcc.String)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.ChrgCode, getData1.ChrgCode)
	require.Equal(t, d1.RefAcc.String, getData1.RefAcc.String)
	require.True(t, d1.ChrAmnt.Equal(getData1.ChrAmnt))

	getData2, err2 := testQueriesDump.GetLnChrgData(context.Background(), d2.BrCode, d2.Acc, d2.ChrgCode, d2.RefAcc.String)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.ChrgCode, getData2.ChrgCode)
	require.Equal(t, d2.RefAcc.String, getData2.RefAcc.String)
	require.True(t, d2.ChrAmnt.Equal(getData2.ChrAmnt))

	// Update Data
	updateD2 := d2
	updateD2.ChrAmnt.Add(decimal.NewFromInt(10))
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestLnChrgData(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetLnChrgData(context.Background(), updateD2.BrCode, updateD2.Acc, updateD2.ChrgCode, updateD2.RefAcc.String)
	require.NoError(t, err1)

	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.Equal(t, updateD2.ChrgCode, getData1.ChrgCode)
	require.Equal(t, updateD2.RefAcc.String, getData1.RefAcc.String)
	require.True(t, updateD2.ChrAmnt.Equal(getData1.ChrAmnt))

	testListLnChrgData(t, ListLnChrgDataParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteLnChrgData(t, d1.BrCode, d1.Acc, d1.ChrgCode, d1.RefAcc.String)
	testDeleteLnChrgData(t, d2.BrCode, d2.Acc, d2.ChrgCode, d2.RefAcc.String)
}

func testListLnChrgData(t *testing.T, arg ListLnChrgDataParams) {

	LnChrgData, err := testQueriesDump.ListLnChrgData(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", LnChrgData)
	require.NotEmpty(t, LnChrgData)

}

func randomLnChrgData() model.LnChrgData {

	arg := model.LnChrgData{
		ModCtr:   1,
		BrCode:   "01",
		Acc:      "0101-4041-0157455",
		ChrgCode: 18,
		RefAcc:   sql.NullString{String: "dsdff", Valid: true},
		ChrAmnt:  decimal.NewFromInt(100),
	}
	return arg
}

func createTestLnChrgData(
	t *testing.T,
	req model.LnChrgData) error {

	err1 := testQueriesDump.CreateLnChrgData(context.Background(), req)
	// fmt.Printf("Get by createTestLnChrgData%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetLnChrgData(context.Background(), req.BrCode, req.Acc, req.ChrgCode, req.RefAcc.String)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.Acc, getData.Acc)
	require.Equal(t, req.ChrgCode, getData.ChrgCode)
	require.Equal(t, req.RefAcc.String, getData.RefAcc.String)
	require.True(t, req.ChrAmnt.Equal(getData.ChrAmnt))

	return err2
}

func updateTestLnChrgData(
	t *testing.T,
	d1 model.LnChrgData) error {

	err := testQueriesDump.UpdateLnChrgData(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteLnChrgData(t *testing.T, brCode string, acc string, code int64, refAcc string) {
	err := testQueriesDump.DeleteLnChrgData(context.Background(), brCode, acc, code, refAcc)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetLnChrgData(context.Background(), brCode, acc, code, refAcc)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
