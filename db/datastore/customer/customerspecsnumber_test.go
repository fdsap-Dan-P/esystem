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

func TestCustomerSpecsNumber(t *testing.T) {

	// Test Data
	d1 := randomCustomerSpecsNumber()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Storage")
	d1.SpecsId = item.Id

	d2 := randomCustomerSpecsNumber()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "Screen")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestCustomerSpecsNumber(t, d1)
	CreatedD2 := createTestCustomerSpecsNumber(t, d2)

	// Get Data
	getData1, err1 := testQueriesCustomer.GetCustomerSpecsNumber(context.Background(), CreatedD1.CustomerId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	getData2, err2 := testQueriesCustomer.GetCustomerSpecsNumber(context.Background(), CreatedD2.CustomerId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CustomerId, getData2.CustomerId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.String(), getData2.Value.String())
	require.Equal(t, d2.Value2.String(), getData2.Value2.String())

	getData, err := testQueriesCustomer.GetCustomerSpecsNumberbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestCustomerSpecsNumber(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.CustomerId, updatedD1.CustomerId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.String(), updatedD1.Value.String())
	require.Equal(t, updateD2.Value2.String(), updatedD1.Value2.String())

	testListCustomerSpecsNumber(t, ListCustomerSpecsNumberParams{
		CustomerId: updatedD1.CustomerId,
		Limit:      5,
		Offset:     0,
	})

	// Delete Data
	testDeleteCustomerSpecsNumber(t, CreatedD1.Uuid)
	testDeleteCustomerSpecsNumber(t, CreatedD2.Uuid)
}

func testListCustomerSpecsNumber(t *testing.T, arg ListCustomerSpecsNumberParams) {

	customerSpecsNumber, err := testQueriesCustomer.ListCustomerSpecsNumber(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", customerSpecsNumber)
	require.NotEmpty(t, customerSpecsNumber)

}

func randomCustomerSpecsNumber() CustomerSpecsNumberRequest {

	cust, _ := testQueriesCustomer.GetCustomerbyAltId(context.Background(), "10001")

	arg := CustomerSpecsNumberRequest{
		CustomerId: cust.Id,
		// SpecsId:             sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomMoney(),
		Value2: util.RandomMoney(),
	}
	return arg
}

func createTestCustomerSpecsNumber(
	t *testing.T,
	d1 CustomerSpecsNumberRequest) model.CustomerSpecsNumber {

	getData1, err := testQueriesCustomer.CreateCustomerSpecsNumber(context.Background(), d1)
	// fmt.Printf("Get by createTestCustomerSpecsNumber%+v\n", getData1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func updateTestCustomerSpecsNumber(
	t *testing.T,
	d1 CustomerSpecsNumberRequest) model.CustomerSpecsNumber {

	getData1, err := testQueriesCustomer.UpdateCustomerSpecsNumber(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.String(), getData1.Value.String())
	require.Equal(t, d1.Value2.String(), getData1.Value2.String())

	return getData1
}

func testDeleteCustomerSpecsNumber(t *testing.T, uuid uuid.UUID) {
	err := testQueriesCustomer.DeleteCustomerSpecsNumber(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesCustomer.GetCustomerSpecsNumberbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
