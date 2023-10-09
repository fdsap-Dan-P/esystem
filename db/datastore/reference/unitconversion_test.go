package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestUnitConversion(t *testing.T) {

	// Test Data
	d1 := randomUnitConversion()
	fr, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "UnitMeasure", 0, "Meter")
	to, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "UnitMeasure", 0, "Centimeter")
	d1.FromId = fr.Id
	d1.ToId = to.Id
	d1.Value = decimal.NewFromInt(100)

	fmt.Printf("fr %+v\n", d1)
	// fmt.Printf("to %+v\n", to)

	d2 := randomUnitConversion()
	fr, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "UnitMeasure", 0, "Centimeter")
	to, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "UnitMeasure", 0, "Meter")
	d2.FromId = fr.Id
	d2.ToId = to.Id
	d2.Value = decimal.RequireFromString(".01")

	// Test Create
	CreatedD1 := createTestUnitConversion(t, d1)
	CreatedD2 := createTestUnitConversion(t, d2)

	// Get Data
	getData1, err1 := testQueriesReference.GetUnitConversion(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.FromId, getData1.FromId)
	require.Equal(t, d1.ToId, getData1.ToId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesReference.GetUnitConversion(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.FromId, getData2.FromId)
	require.Equal(t, d2.ToId, getData2.ToId)
	require.Equal(t, d2.Value.String(), getData2.Value.String())
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesReference.GetUnitConversionbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)

	fmt.Printf("%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.Value = updateD2.Value.Add(decimal.NewFromInt(1))

	// log.Println(updateD2)
	updatedD1 := updateTestUnitConversion(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.FromId, updatedD1.FromId)
	require.Equal(t, updateD2.ToId, updatedD1.ToId)
	require.Equal(t, updateD2.Value.String(), updatedD1.Value.String())
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	// testDeleteUnitConversion(t, getData1.Id)
	// testDeleteUnitConversion(t, getData2.Id)
}

func TestListUnitConversion(t *testing.T) {

	arg := ListUnitConversionParams{
		TypeId: 1027,
		Limit:  5,
		Offset: 0,
	}

	unitConversion, err := testQueriesReference.ListUnitConversion(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", unitConversion)
	require.NotEmpty(t, unitConversion)

}

func randomUnitConversion() UnitConversionRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ConversionType", 0, "Memory")

	arg := UnitConversionRequest{
		TypeId: typ.Id,
		// FromId: util.RandomInt(1, 100),
		// ToId:   util.RandomInt(1, 100),
		// Value:  util.RandomMoney(),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	// fmt.Printf("%+v\n", arg.ar)
	return arg
}

func createTestUnitConversion(
	t *testing.T,
	d1 UnitConversionRequest) model.UnitConversion {

	getData1, err := testQueriesReference.CreateUnitConversion(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	fmt.Printf("testing d1.OtherInfo.String: %+v\n", d1.OtherInfo.String)
	fmt.Printf("testing getData1.OtherInfo.String: %+v\n", getData1.OtherInfo.String)
	fmt.Printf("d1: %+v\n", d1)
	fmt.Printf("getData1: %+v\n", getData1)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.FromId, getData1.FromId)
	require.Equal(t, d1.ToId, getData1.ToId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestUnitConversion(
	t *testing.T,
	d1 UnitConversionRequest) model.UnitConversion {

	getData1, err := testQueriesReference.UpdateUnitConversion(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.FromId, getData1.FromId)
	require.Equal(t, d1.ToId, getData1.ToId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteUnitConversion(t *testing.T, id int64) {
	err := testQueriesReference.DeleteUnitConversion(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesReference.GetUnitConversion(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
