package db

import (
	"context"
	"database/sql"
	"log"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnit(t *testing.T) {

	// Test Data
	d1 := randomUnit()
	d2 := randomUnit()
	d2.UnitCode = d2.UnitCode + 1

	log.Printf("d1: %v", d1)
	log.Printf("d2: %v", d2)

	err := createTestUnit(t, d1)
	require.NoError(t, err)

	err = createTestUnit(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetUnit(context.Background(), d1.UnitCode)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.UnitCode, getData1.UnitCode)
	require.Equal(t, d1.Unit, getData1.Unit)
	require.Equal(t, d1.AreaCode, getData1.AreaCode)
	require.Equal(t, d1.FName, getData1.FName)
	require.Equal(t, d1.LName, getData1.LName)
	require.Equal(t, d1.MName, getData1.MName)
	require.Equal(t, d1.VatReg, getData1.VatReg)
	require.Equal(t, d1.UnitAddress, getData1.UnitAddress)

	getData2, err2 := testQueriesLocal.GetUnit(context.Background(), d2.UnitCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.UnitCode, getData2.UnitCode)
	require.Equal(t, d2.Unit, getData2.Unit)
	require.Equal(t, d2.AreaCode, getData2.AreaCode)
	require.Equal(t, d2.FName, getData2.FName)
	require.Equal(t, d2.LName, getData2.LName)
	require.Equal(t, d2.MName, getData2.MName)
	require.Equal(t, d2.VatReg, getData2.VatReg)
	require.Equal(t, d2.UnitAddress, getData2.UnitAddress)

	// Update Data
	updateD2 := d2
	updateD2.UnitCode = getData2.UnitCode
	updateD2.Unit.String = updateD2.Unit.String + "Edited"

	// log.Println(updateD2)
	err3 := updateTestUnit(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesLocal.GetUnit(context.Background(), updateD2.UnitCode)
	require.NoError(t, err1)

	require.Equal(t, updateD2.UnitCode, getData1.UnitCode)
	require.Equal(t, updateD2.UnitCode, getData1.UnitCode)
	require.Equal(t, updateD2.Unit, getData1.Unit)
	require.Equal(t, updateD2.AreaCode, getData1.AreaCode)
	require.Equal(t, updateD2.FName, getData1.FName)
	require.Equal(t, updateD2.LName, getData1.LName)
	require.Equal(t, updateD2.MName, getData1.MName)
	require.Equal(t, updateD2.VatReg, getData1.VatReg)
	require.Equal(t, updateD2.UnitAddress, getData1.UnitAddress)

	testListUnit(t, ListUnitParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteUnit(t, d1.UnitCode)
	testDeleteUnit(t, d2.UnitCode)
}

func testListUnit(t *testing.T, arg ListUnitParams) {

	Unit, err := testQueriesLocal.ListUnit(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Unit)
	require.NotEmpty(t, Unit)

}

func randomUnit() UnitRequest {

	arg := UnitRequest{
		UnitCode: 111,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Unit:     sql.NullString{String: "Unit", Valid: true},
		AreaCode: sql.NullInt64{Int64: 101, Valid: true},
		FName:    sql.NullString{String: "FName", Valid: true},
		LName:    sql.NullString{String: "LName", Valid: true},
		MName:    sql.NullString{String: "M", Valid: true},
	}
	return arg
}

func createTestUnit(
	t *testing.T,
	req UnitRequest) error {

	err1 := testQueriesLocal.CreateUnit(context.Background(), req)
	// fmt.Printf("Get by createTestUnit%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetUnit(context.Background(), req.UnitCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.UnitCode, getData.UnitCode)
	require.Equal(t, req.UnitCode, getData.UnitCode)
	require.Equal(t, req.Unit, getData.Unit)
	require.Equal(t, req.AreaCode, getData.AreaCode)
	require.Equal(t, req.FName, getData.FName)
	require.Equal(t, req.LName, getData.LName)
	require.Equal(t, req.MName, getData.MName)
	require.Equal(t, req.VatReg, getData.VatReg)
	require.Equal(t, req.UnitAddress, getData.UnitAddress)

	return err2
}

func updateTestUnit(
	t *testing.T,
	d1 UnitRequest) error {

	err := testQueriesLocal.UpdateUnit(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteUnit(t *testing.T, UnitCode int64) {
	err := testQueriesLocal.DeleteUnit(context.Background(), UnitCode)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetUnit(context.Background(), UnitCode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
