package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestTrnHeadSpecsString(t *testing.T) {

	// Test Data
	d1 := randomTrnHeadSpecsString()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomTrnHeadSpecsString()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestTrnHeadSpecsString(t, d1)
	CreatedD2 := createTestTrnHeadSpecsString(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTrnHeadSpecsString(context.Background(), CreatedD1.TrnHeadId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	getData2, err2 := testQueriesTransaction.GetTrnHeadSpecsString(context.Background(), CreatedD2.TrnHeadId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TrnHeadId, getData2.TrnHeadId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value, getData2.Value)

	getData, err := testQueriesTransaction.GetTrnHeadSpecsStringbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTrnHeadSpecsString(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TrnHeadId, updatedD1.TrnHeadId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value, updatedD1.Value)

	testListTrnHeadSpecsString(t, ListTrnHeadSpecsStringParams{
		TrnHeadId: updatedD1.TrnHeadId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteTrnHeadSpecsString(t, CreatedD1.Uuid)
	testDeleteTrnHeadSpecsString(t, CreatedD2.Uuid)
}

func testListTrnHeadSpecsString(t *testing.T, arg ListTrnHeadSpecsStringParams) {

	trnHeadSpecsString, err := testQueriesTransaction.ListTrnHeadSpecsString(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", trnHeadSpecsString)
	require.NotEmpty(t, trnHeadSpecsString)

}

func randomTrnHeadSpecsString() TrnHeadSpecsStringRequest {

	trn, _ := testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("2af90d74-3bee-48c5-8935-443edafb8f5a"))

	arg := TrnHeadSpecsStringRequest{
		TrnHeadId: trn.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value: util.RandomString(20),
	}
	return arg
}

func createTestTrnHeadSpecsString(
	t *testing.T,
	d1 TrnHeadSpecsStringRequest) model.TrnHeadSpecsString {

	getData1, err := testQueriesTransaction.CreateTrnHeadSpecsString(context.Background(), d1)
	// fmt.Printf("Get by createTestTrnHeadSpecsString%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func updateTestTrnHeadSpecsString(
	t *testing.T,
	d1 TrnHeadSpecsStringRequest) model.TrnHeadSpecsString {

	getData1, err := testQueriesTransaction.UpdateTrnHeadSpecsString(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func testDeleteTrnHeadSpecsString(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTrnHeadSpecsString(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTrnHeadSpecsStringbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
