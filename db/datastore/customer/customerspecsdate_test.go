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

func TestCustomerSpecsDate(t *testing.T) {

	// Test Data
	d1 := randomCustomerSpecsDate()
	itemParent, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "ComputerSpecs")
	item, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateManufactured")
	d1.SpecsId = item.Id

	d2 := randomCustomerSpecsDate()
	item, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", itemParent.Id, "DateExpired")
	d2.SpecsId = item.Id

	// Test Create
	CreatedD1 := createTestCustomerSpecsDate(t, d1)
	CreatedD2 := createTestCustomerSpecsDate(t, d2)

	// Get Data
	getData1, err1 := testQueriesCustomer.GetCustomerSpecsDate(context.Background(), CreatedD1.CustomerId, CreatedD1.SpecsId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	getData2, err2 := testQueriesCustomer.GetCustomerSpecsDate(context.Background(), CreatedD2.CustomerId, CreatedD2.SpecsId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CustomerId, getData2.CustomerId)
	require.Equal(t, d2.SpecsId, getData2.SpecsId)
	require.Equal(t, d2.Value.Format("2006-01-02"), getData2.Value.Format("2006-01-02"))
	require.Equal(t, d2.Value2.Format("2006-01-02"), getData2.Value2.Format("2006-01-02"))

	getData, err := testQueriesCustomer.GetCustomerSpecsDatebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUID%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestCustomerSpecsDate(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.CustomerId, updatedD1.CustomerId)
	require.Equal(t, updateD2.SpecsId, updatedD1.SpecsId)
	require.Equal(t, updateD2.Value.Format("2006-01-02"), updatedD1.Value.Format("2006-01-02"))
	require.Equal(t, updateD2.Value2.Format("2006-01-02"), updatedD1.Value2.Format("2006-01-02"))

	testListCustomerSpecsDate(t, ListCustomerSpecsDateParams{
		CustomerId: updatedD1.CustomerId,
		Limit:      5,
		Offset:     0,
	})

	// Delete Data
	testDeleteCustomerSpecsDate(t, CreatedD1.Uuid)
	testDeleteCustomerSpecsDate(t, CreatedD2.Uuid)
}

func testListCustomerSpecsDate(t *testing.T, arg ListCustomerSpecsDateParams) {

	customerSpecsDate, err := testQueriesCustomer.ListCustomerSpecsDate(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", customerSpecsDate)
	require.NotEmpty(t, customerSpecsDate)
}

func randomCustomerSpecsDate() CustomerSpecsDateRequest {

	cust, _ := testQueriesCustomer.GetCustomerbyAltId(context.Background(), "10001")

	arg := CustomerSpecsDateRequest{
		CustomerId: cust.Id,
		// SpecsId: sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Value:  util.RandomDate(),
		Value2: util.RandomDate(),
	}
	return arg
}

func createTestCustomerSpecsDate(
	t *testing.T,
	d1 CustomerSpecsDateRequest) model.CustomerSpecsDate {

	getData1, err := testQueriesCustomer.CreateCustomerSpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func updateTestCustomerSpecsDate(
	t *testing.T,
	d1 CustomerSpecsDateRequest) model.CustomerSpecsDate {

	getData1, err := testQueriesCustomer.UpdateCustomerSpecsDate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.SpecsId, getData1.SpecsId)
	require.Equal(t, d1.Value.Format("2006-01-02"), getData1.Value.Format("2006-01-02"))
	require.Equal(t, d1.Value2.Format("2006-01-02"), getData1.Value2.Format("2006-01-02"))

	return getData1
}

func testDeleteCustomerSpecsDate(t *testing.T, uuid uuid.UUID) {
	err := testQueriesCustomer.DeleteCustomerSpecsDate(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesCustomer.GetCustomerSpecsDatebyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
