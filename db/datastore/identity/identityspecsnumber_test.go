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

func TestIdentitySpecsNumber(t *testing.T) {

	// Test Data
	d1 := randomIdentitySpecsNumber()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomIdentitySpecsNumber()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestIdentitySpecsNumber(t, d1)
	CreatedD2 := createTestIdentitySpecsNumber(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetIdentitySpecsNumber(context.Background(), CreatedD1.Iiid, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	getData2, err2 := testQueriesIdentity.GetIdentitySpecsNumber(context.Background(), CreatedD2.Iiid, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.String(), getData2.Value.String())
	require.Equal(t, d2.Value2.String(), getData2.Value2.String())

	getData, err := testQueriesIdentity.GetIdentitySpecsNumberbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestIdentitySpecsNumber(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.String(), updatedD1.Value.String())
	require.Equal(t, updateD2.Value2.String(), updatedD1.Value2.String())

	testListIdentitySpecsNumber(t, ListIdentitySpecsNumberParams{
		Iiid:   updatedD1.Iiid,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteIdentitySpecsNumber(t, CreatedD1.Uuid)
	testDeleteIdentitySpecsNumber(t, CreatedD2.Uuid)
}

func testListIdentitySpecsNumber(t *testing.T, arg ListIdentitySpecsNumberParams) {

	identitySpecsNumber, err := testQueriesIdentity.ListIdentitySpecsNumber(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", identitySpecsNumber)
	require.NotEmpty(t, identitySpecsNumber)

}

func randomIdentitySpecsNumber() IdentitySpecsNumberRequest {

	accQtl, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "101")

	arg := IdentitySpecsNumberRequest{
		Iiid: accQtl.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomMoney(),
		Value2: util.RandomMoney(),
	}
	return arg
}

func createTestIdentitySpecsNumber(
	t *testing.T,
	d1 IdentitySpecsNumberRequest) model.IdentitySpecsNumber {

	getData1, err := testQueriesIdentity.CreateIdentitySpecsNumber(context.Background(), d1)
	// fmt.Printf("Get by createTestIdentitySpecsNumber%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func updateTestIdentitySpecsNumber(
	t *testing.T,
	d1 IdentitySpecsNumberRequest) model.IdentitySpecsNumber {

	getData1, err := testQueriesIdentity.UpdateIdentitySpecsNumber(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func testDeleteIdentitySpecsNumber(t *testing.T, uuid uuid.UUID) {
	err := testQueriesIdentity.DeleteIdentitySpecsNumber(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetIdentitySpecsNumberbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
