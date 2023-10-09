package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"simplebank/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestTrnHeadSpecsRef(t *testing.T) {

	// Test Data
	d1 := randomTrnHeadSpecsRef()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomTrnHeadSpecsRef()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestTrnHeadSpecsRef(t, d1)
	CreatedD2 := createTestTrnHeadSpecsRef(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTrnHeadSpecsRef(context.Background(), CreatedD1.TrnHeadId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	getData2, err2 := testQueriesTransaction.GetTrnHeadSpecsRef(context.Background(), CreatedD2.TrnHeadId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TrnHeadId, getData2.TrnHeadId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.RefId, getData2.RefId)

	getData, err := testQueriesTransaction.GetTrnHeadSpecsRefbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTrnHeadSpecsRef(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TrnHeadId, updatedD1.TrnHeadId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.RefId, updatedD1.RefId)

	testListTrnHeadSpecsRef(t, ListTrnHeadSpecsRefParams{
		TrnHeadId: updatedD1.TrnHeadId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteTrnHeadSpecsRef(t, CreatedD1.Uuid)
	testDeleteTrnHeadSpecsRef(t, CreatedD2.Uuid)
}

func testListTrnHeadSpecsRef(t *testing.T, arg ListTrnHeadSpecsRefParams) {

	trnHeadSpecsRef, err := testQueriesTransaction.ListTrnHeadSpecsRef(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", trnHeadSpecsRef)
	require.NotEmpty(t, trnHeadSpecsRef)

}

func randomTrnHeadSpecsRef() TrnHeadSpecsRefRequest {

	trn, _ := testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("2af90d74-3bee-48c5-8935-443edafb8f5a"))
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeetype", 0, "Employed")

	arg := TrnHeadSpecsRefRequest{
		TrnHeadId: trn.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		RefId: typ.Id,
	}
	return arg
}

func createTestTrnHeadSpecsRef(
	t *testing.T,
	d1 TrnHeadSpecsRefRequest) model.TrnHeadSpecsRef {

	getData1, err := testQueriesTransaction.CreateTrnHeadSpecsRef(context.Background(), d1)
	// fmt.Printf("Get by createTestTrnHeadSpecsRef%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func updateTestTrnHeadSpecsRef(
	t *testing.T,
	d1 TrnHeadSpecsRefRequest) model.TrnHeadSpecsRef {

	getData1, err := testQueriesTransaction.UpdateTrnHeadSpecsRef(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func testDeleteTrnHeadSpecsRef(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTrnHeadSpecsRef(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTrnHeadSpecsRefbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
