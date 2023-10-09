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

func TestTrnHeadSpecsDate(t *testing.T) {

	// Test Data
	d1 := randomTrnHeadSpecsDate()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateManufactured")
	d1.SpecsId = item.Id

	d2 := randomTrnHeadSpecsDate()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateExpired")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestTrnHeadSpecsDate(t, d1)
	CreatedD2 := createTestTrnHeadSpecsDate(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTrnHeadSpecsDate(context.Background(), CreatedD1.TrnHeadId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	getData2, err2 := testQueriesTransaction.GetTrnHeadSpecsDate(context.Background(), CreatedD2.TrnHeadId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TrnHeadId, getData2.TrnHeadId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.Format("2006-01-02"), getData2.Value.Format("2006-01-02"))
	require.Equal(t, d2.Value2.Format("2006-01-02"), getData2.Value2.Format("2006-01-02"))

	getData, err := testQueriesTransaction.GetTrnHeadSpecsDatebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTrnHeadSpecsDate(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TrnHeadId, updatedD1.TrnHeadId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.Format("2006-01-02"), updatedD1.Value.Format("2006-01-02"))
	require.Equal(t, updateD2.Value2.Format("2006-01-02"), updatedD1.Value2.Format("2006-01-02"))

	testListTrnHeadSpecsDate(t, ListTrnHeadSpecsDateParams{
		TrnHeadId: updatedD1.TrnHeadId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteTrnHeadSpecsDate(t, CreatedD1.Uuid)
	testDeleteTrnHeadSpecsDate(t, CreatedD2.Uuid)
}

func testListTrnHeadSpecsDate(t *testing.T, arg ListTrnHeadSpecsDateParams) {

	trnHeadSpecsDate, err := testQueriesTransaction.ListTrnHeadSpecsDate(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", trnHeadSpecsDate)
	require.NotEmpty(t, trnHeadSpecsDate)
}

func randomTrnHeadSpecsDate() TrnHeadSpecsDateRequest {

	trn, _ := testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("2af90d74-3bee-48c5-8935-443edafb8f5a"))

	arg := TrnHeadSpecsDateRequest{
		TrnHeadId: trn.Id,
		// SpecsId: sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomDate(),
		Value2: util.RandomDate(),
	}
	return arg
}

func createTestTrnHeadSpecsDate(
	t *testing.T,
	d1 TrnHeadSpecsDateRequest) model.TrnHeadSpecsDate {

	getData1, err := testQueriesTransaction.CreateTrnHeadSpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func updateTestTrnHeadSpecsDate(
	t *testing.T,
	d1 TrnHeadSpecsDateRequest) model.TrnHeadSpecsDate {

	getData1, err := testQueriesTransaction.UpdateTrnHeadSpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func testDeleteTrnHeadSpecsDate(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTrnHeadSpecsDate(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTrnHeadSpecsDatebyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
