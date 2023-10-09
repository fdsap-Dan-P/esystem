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

func TestIdentitySpecsDate(t *testing.T) {

	// Test Data
	d1 := randomIdentitySpecsDate()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateManufactured")
	d1.SpecsId = item.Id

	d2 := randomIdentitySpecsDate()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateExpired")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestIdentitySpecsDate(t, d1)
	CreatedD2 := createTestIdentitySpecsDate(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetIdentitySpecsDate(context.Background(), CreatedD1.Iiid, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	getData2, err2 := testQueriesIdentity.GetIdentitySpecsDate(context.Background(), CreatedD2.Iiid, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.Format("2006-01-02"), getData2.Value.Format("2006-01-02"))
	require.Equal(t, d2.Value2.Format("2006-01-02"), getData2.Value2.Format("2006-01-02"))

	getData, err := testQueriesIdentity.GetIdentitySpecsDatebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestIdentitySpecsDate(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.Format("2006-01-02"), updatedD1.Value.Format("2006-01-02"))
	require.Equal(t, updateD2.Value2.Format("2006-01-02"), updatedD1.Value2.Format("2006-01-02"))

	testListIdentitySpecsDate(t, ListIdentitySpecsDateParams{
		Iiid:   updatedD1.Iiid,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteIdentitySpecsDate(t, CreatedD1.Uuid)
	testDeleteIdentitySpecsDate(t, CreatedD2.Uuid)
}

func testListIdentitySpecsDate(t *testing.T, arg ListIdentitySpecsDateParams) {

	identitySpecsDate, err := testQueriesIdentity.ListIdentitySpecsDate(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", identitySpecsDate)
	require.NotEmpty(t, identitySpecsDate)
}

func randomIdentitySpecsDate() IdentitySpecsDateRequest {

	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")

	arg := IdentitySpecsDateRequest{
		Iiid: ii.Id,
		// SpecsId: sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomDate(),
		Value2: util.RandomDate(),
	}
	return arg
}

func createTestIdentitySpecsDate(
	t *testing.T,
	d1 IdentitySpecsDateRequest) model.IdentitySpecsDate {

	getData1, err := testQueriesIdentity.CreateIdentitySpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func updateTestIdentitySpecsDate(
	t *testing.T,
	d1 IdentitySpecsDateRequest) model.IdentitySpecsDate {

	getData1, err := testQueriesIdentity.UpdateIdentitySpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func testDeleteIdentitySpecsDate(t *testing.T, uuid uuid.UUID) {
	err := testQueriesIdentity.DeleteIdentitySpecsDate(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetIdentitySpecsDatebyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
