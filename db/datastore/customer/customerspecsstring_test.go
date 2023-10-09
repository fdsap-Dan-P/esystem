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

func TestCustomerSpecsString(t *testing.T) {

	// Test Data
	d1 := randomCustomerSpecsString()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomCustomerSpecsString()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestCustomerSpecsString(t, d1)
	CreatedD2 := createTestCustomerSpecsString(t, d2)

	// Get Data
	getData1, err1 := testQueriesCustomer.GetCustomerSpecsString(context.Background(), CreatedD1.CustomerId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	getData2, err2 := testQueriesCustomer.GetCustomerSpecsString(context.Background(), CreatedD2.CustomerId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CustomerId, getData2.CustomerId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value, getData2.Value)

	getData, err := testQueriesCustomer.GetCustomerSpecsStringbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestCustomerSpecsString(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.CustomerId, updatedD1.CustomerId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value, updatedD1.Value)

	testListCustomerSpecsString(t, ListCustomerSpecsStringParams{
		CustomerId: updatedD1.CustomerId,
		Limit:      5,
		Offset:     0,
	})

	// Delete Data
	testDeleteCustomerSpecsString(t, CreatedD1.Uuid)
	testDeleteCustomerSpecsString(t, CreatedD2.Uuid)
}

func testListCustomerSpecsString(t *testing.T, arg ListCustomerSpecsStringParams) {

	customerSpecsString, err := testQueriesCustomer.ListCustomerSpecsString(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", customerSpecsString)
	require.NotEmpty(t, customerSpecsString)

}

func randomCustomerSpecsString() CustomerSpecsStringRequest {

	cust, _ := testQueriesCustomer.GetCustomerbyAltId(context.Background(), "10001")

	arg := CustomerSpecsStringRequest{
		CustomerId: cust.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value: util.RandomString(20),
	}
	return arg
}

func createTestCustomerSpecsString(
	t *testing.T,
	d1 CustomerSpecsStringRequest) model.CustomerSpecsString {

	getData1, err := testQueriesCustomer.CreateCustomerSpecsString(context.Background(), d1)
	// fmt.Printf("Get by createTestCustomerSpecsString%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func updateTestCustomerSpecsString(
	t *testing.T,
	d1 CustomerSpecsStringRequest) model.CustomerSpecsString {

	getData1, err := testQueriesCustomer.UpdateCustomerSpecsString(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value, getData1.Value)

	return getData1
}

func testDeleteCustomerSpecsString(t *testing.T, uuid uuid.UUID) {
	err := testQueriesCustomer.DeleteCustomerSpecsString(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesCustomer.GetCustomerSpecsStringbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
