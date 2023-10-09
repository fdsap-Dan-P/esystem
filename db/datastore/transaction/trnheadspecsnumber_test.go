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

func TestTrnHeadSpecsNumber(t *testing.T) {

	// Test Data
	d1 := randomTrnHeadSpecsNumber()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomTrnHeadSpecsNumber()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestTrnHeadSpecsNumber(t, d1)
	CreatedD2 := createTestTrnHeadSpecsNumber(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTrnHeadSpecsNumber(context.Background(), CreatedD1.TrnHeadId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	getData2, err2 := testQueriesTransaction.GetTrnHeadSpecsNumber(context.Background(), CreatedD2.TrnHeadId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TrnHeadId, getData2.TrnHeadId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.String(), getData2.Value.String())
	require.Equal(t, d2.Value2.String(), getData2.Value2.String())

	getData, err := testQueriesTransaction.GetTrnHeadSpecsNumberbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTrnHeadSpecsNumber(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TrnHeadId, updatedD1.TrnHeadId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.String(), updatedD1.Value.String())
	require.Equal(t, updateD2.Value2.String(), updatedD1.Value2.String())

	testListTrnHeadSpecsNumber(t, ListTrnHeadSpecsNumberParams{
		TrnHeadId: updatedD1.TrnHeadId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteTrnHeadSpecsNumber(t, CreatedD1.Uuid)
	testDeleteTrnHeadSpecsNumber(t, CreatedD2.Uuid)
}

func testListTrnHeadSpecsNumber(t *testing.T, arg ListTrnHeadSpecsNumberParams) {

	trnHeadSpecsNumber, err := testQueriesTransaction.ListTrnHeadSpecsNumber(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", trnHeadSpecsNumber)
	require.NotEmpty(t, trnHeadSpecsNumber)

}

func randomTrnHeadSpecsNumber() TrnHeadSpecsNumberRequest {

	trn, _ := testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("2af90d74-3bee-48c5-8935-443edafb8f5a"))

	arg := TrnHeadSpecsNumberRequest{
		TrnHeadId: trn.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomMoney(),
		Value2: util.RandomMoney(),
	}
	return arg
}

func createTestTrnHeadSpecsNumber(
	t *testing.T,
	d1 TrnHeadSpecsNumberRequest) model.TrnHeadSpecsNumber {

	getData1, err := testQueriesTransaction.CreateTrnHeadSpecsNumber(context.Background(), d1)
	// fmt.Printf("Get by createTestTrnHeadSpecsNumber%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func updateTestTrnHeadSpecsNumber(
	t *testing.T,
	d1 TrnHeadSpecsNumberRequest) model.TrnHeadSpecsNumber {

	getData1, err := testQueriesTransaction.UpdateTrnHeadSpecsNumber(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func testDeleteTrnHeadSpecsNumber(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTrnHeadSpecsNumber(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTrnHeadSpecsNumberbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
