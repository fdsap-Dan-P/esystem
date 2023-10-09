package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestIncomeSource(t *testing.T) {

	// Test Data
	d1 := randomIncomeSource()
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")
	d1.Iiid = ii.Id
	d2 := randomIncomeSource()
	ii, _ = testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "101")
	d2.Iiid = ii.Id

	// Test Create
	CreatedD1 := createTestIncomeSource(t, d1)
	CreatedD2 := createTestIncomeSource(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetIncomeSource(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Source, getData1.Source)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.MinIncome.String(), getData1.MinIncome.String())
	require.Equal(t, d1.MaxIncome.String(), getData1.MaxIncome.String())
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesIdentity.GetIncomeSource(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.Source, getData2.Source)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.MinIncome.String(), getData2.MinIncome.String())
	require.Equal(t, d2.MaxIncome.String(), getData2.MaxIncome.String())
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesIdentity.GetIncomeSourcebyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// getData, err = testQueriesIdentity.GetIncomeSourcebyName(context.Background(), CreatedD1.Source)
	// require.NoError(t, err)
	// require.NotEmpty(t, getData)
	// require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	// fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	updateD2.Source = updateD2.Source + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestIncomeSource(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.Source, updatedD1.Source)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.MinIncome.String(), updatedD1.MinIncome.String())
	require.Equal(t, updateD2.MaxIncome.String(), updatedD1.MaxIncome.String())
	require.Equal(t, updateD2.Remarks, updatedD1.Remarks)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListIncomeSource(t, ListIncomeSourceParams{
		Iiid:   updatedD1.Iiid,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteIncomeSource(t, CreatedD1.Uuid)
	testDeleteIncomeSource(t, CreatedD2.Uuid)
}

func testListIncomeSource(t *testing.T, arg ListIncomeSourceParams) {

	incomeSource, err := testQueriesIdentity.ListIncomeSource(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", incomeSource)
	require.NotEmpty(t, incomeSource)

}

func randomIncomeSource() IncomeSourceRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SourceofIncome", 0, "Salary")

	arg := IncomeSourceRequest{
		// Iiid:      ii.Id,
		Series:    int16(util.RandomInt32(1, 100)),
		Source:    util.RandomString(10),
		TypeId:    typ.Id,
		MinIncome: util.RandomMoney(),
		MaxIncome: util.RandomMoney(),
		Remarks:   sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestIncomeSource(
	t *testing.T,
	d1 IncomeSourceRequest) model.IncomeSource {

	getData1, err := testQueriesIdentity.CreateIncomeSource(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Source, getData1.Source)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.MinIncome.String(), getData1.MinIncome.String())
	require.Equal(t, d1.MaxIncome.String(), getData1.MaxIncome.String())
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestIncomeSource(
	t *testing.T,
	d1 IncomeSourceRequest) model.IncomeSource {

	getData1, err := testQueriesIdentity.UpdateIncomeSource(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Source, getData1.Source)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.MinIncome.String(), getData1.MinIncome.String())
	require.Equal(t, d1.MaxIncome.String(), getData1.MaxIncome.String())
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteIncomeSource(t *testing.T, uuid uuid.UUID) {
	err := testQueriesIdentity.DeleteIncomeSource(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetIncomeSource(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
