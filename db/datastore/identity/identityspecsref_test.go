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

func TestIdentitySpecsRef(t *testing.T) {

	// Test Data
	d1 := randomIdentitySpecsRef()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomIdentitySpecsRef()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestIdentitySpecsRef(t, d1)
	CreatedD2 := createTestIdentitySpecsRef(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetIdentitySpecsRef(context.Background(), CreatedD1.Iiid, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	getData2, err2 := testQueriesIdentity.GetIdentitySpecsRef(context.Background(), CreatedD2.Iiid, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.RefId, getData2.RefId)

	getData, err := testQueriesIdentity.GetIdentitySpecsRefbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestIdentitySpecsRef(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.RefId, updatedD1.RefId)

	testListIdentitySpecsRef(t, ListIdentitySpecsRefParams{
		Iiid:   updatedD1.Iiid,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteIdentitySpecsRef(t, CreatedD1.Uuid)
	testDeleteIdentitySpecsRef(t, CreatedD2.Uuid)
}

func testListIdentitySpecsRef(t *testing.T, arg ListIdentitySpecsRefParams) {

	identitySpecsRef, err := testQueriesIdentity.ListIdentitySpecsRef(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", identitySpecsRef)
	require.NotEmpty(t, identitySpecsRef)

}

func randomIdentitySpecsRef() IdentitySpecsRefRequest {

	accQtl, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "101")
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeetype", 0, "Employed")

	arg := IdentitySpecsRefRequest{
		Iiid: accQtl.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		RefId: typ.Id,
	}
	return arg
}

func createTestIdentitySpecsRef(
	t *testing.T,
	d1 IdentitySpecsRefRequest) model.IdentitySpecsRef {

	getData1, err := testQueriesIdentity.CreateIdentitySpecsRef(context.Background(), d1)
	// fmt.Printf("Get by createTestIdentitySpecsRef%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func updateTestIdentitySpecsRef(
	t *testing.T,
	d1 IdentitySpecsRefRequest) model.IdentitySpecsRef {

	getData1, err := testQueriesIdentity.UpdateIdentitySpecsRef(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func testDeleteIdentitySpecsRef(t *testing.T, uuid uuid.UUID) {
	err := testQueriesIdentity.DeleteIdentitySpecsRef(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetIdentitySpecsRefbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
