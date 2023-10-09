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

func TestCustomerSpecsRef(t *testing.T) {

	// Test Data
	d1 := randomCustomerSpecsRef()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomCustomerSpecsRef()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestCustomerSpecsRef(t, d1)
	CreatedD2 := createTestCustomerSpecsRef(t, d2)

	// Get Data
	getData1, err1 := testQueriesCustomer.GetCustomerSpecsRef(context.Background(), CreatedD1.CustomerId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	getData2, err2 := testQueriesCustomer.GetCustomerSpecsRef(context.Background(), CreatedD2.CustomerId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CustomerId, getData2.CustomerId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.RefId, getData2.RefId)

	getData, err := testQueriesCustomer.GetCustomerSpecsRefbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestCustomerSpecsRef(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.CustomerId, updatedD1.CustomerId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.RefId, updatedD1.RefId)

	testListCustomerSpecsRef(t, ListCustomerSpecsRefParams{
		CustomerId: updatedD1.CustomerId,
		Limit:      5,
		Offset:     0,
	})

	// Delete Data
	testDeleteCustomerSpecsRef(t, CreatedD1.Uuid)
	testDeleteCustomerSpecsRef(t, CreatedD2.Uuid)
}

func testListCustomerSpecsRef(t *testing.T, arg ListCustomerSpecsRefParams) {

	customerSpecsRef, err := testQueriesCustomer.ListCustomerSpecsRef(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", customerSpecsRef)
	require.NotEmpty(t, customerSpecsRef)

}

func randomCustomerSpecsRef() CustomerSpecsRefRequest {

	cust, _ := testQueriesCustomer.GetCustomerbyAltId(context.Background(), "10001")
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "employeetype", 0, "Employed")

	arg := CustomerSpecsRefRequest{
		CustomerId: cust.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		RefId: typ.Id,
	}
	return arg
}

func createTestCustomerSpecsRef(
	t *testing.T,
	d1 CustomerSpecsRefRequest) model.CustomerSpecsRef {

	getData1, err := testQueriesCustomer.CreateCustomerSpecsRef(context.Background(), d1)
	// fmt.Printf("Get by createTestCustomerSpecsRef%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func updateTestCustomerSpecsRef(
	t *testing.T,
	d1 CustomerSpecsRefRequest) model.CustomerSpecsRef {

	getData1, err := testQueriesCustomer.UpdateCustomerSpecsRef(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.RefId, getData1.RefId)

	return getData1
}

func testDeleteCustomerSpecsRef(t *testing.T, uuid uuid.UUID) {
	err := testQueriesCustomer.DeleteCustomerSpecsRef(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesCustomer.GetCustomerSpecsRefbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
