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

func TestIdentitySpecsString(t *testing.T) {

	// Test Data
	d1 := randomIdentitySpecsString()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomIdentitySpecsString()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestIdentitySpecsString(t, d1)
	CreatedD2 := createTestIdentitySpecsString(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetIdentitySpecsString(context.Background(), CreatedD1.Iiid, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	getData2, err2 := testQueriesIdentity.GetIdentitySpecsString(context.Background(), CreatedD2.Iiid, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value, getData2.Value)

	getData, err := testQueriesIdentity.GetIdentitySpecsStringbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestIdentitySpecsString(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value, updatedD1.Value)

	testListIdentitySpecsString(t, ListIdentitySpecsStringParams{
		Iiid:   updatedD1.Iiid,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteIdentitySpecsString(t, CreatedD1.Uuid)
	testDeleteIdentitySpecsString(t, CreatedD2.Uuid)
}

func testListIdentitySpecsString(t *testing.T, arg ListIdentitySpecsStringParams) {

	identitySpecsString, err := testQueriesIdentity.ListIdentitySpecsString(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", identitySpecsString)
	require.NotEmpty(t, identitySpecsString)

}

func randomIdentitySpecsString() IdentitySpecsStringRequest {

	accQtl, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "101")

	arg := IdentitySpecsStringRequest{
		Iiid: accQtl.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value: util.RandomString(20),
	}
	return arg
}

func createTestIdentitySpecsString(
	t *testing.T,
	d1 IdentitySpecsStringRequest) model.IdentitySpecsString {

	getData1, err := testQueriesIdentity.CreateIdentitySpecsString(context.Background(), d1)
	// fmt.Printf("Get by createTestIdentitySpecsString%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func updateTestIdentitySpecsString(
	t *testing.T,
	d1 IdentitySpecsStringRequest) model.IdentitySpecsString {

	getData1, err := testQueriesIdentity.UpdateIdentitySpecsString(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func testDeleteIdentitySpecsString(t *testing.T, uuid uuid.UUID) {
	err := testQueriesIdentity.DeleteIdentitySpecsString(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetIdentitySpecsStringbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
