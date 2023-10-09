package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCustomerBeneficiary(t *testing.T) {

	// Test Data
	d1 := randomCustomerBeneficiary()
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")
	d1.Iiid = ii.Id
	d2 := randomCustomerBeneficiary()
	ii, _ = testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "101")
	d2.Iiid = ii.Id

	// Test Create
	CreatedD1 := createTestCustomerBeneficiary(t, d1)
	CreatedD2 := createTestCustomerBeneficiary(t, d2)

	// Get Data
	getData1, err1 := testQueriesCustomer.GetCustomerBeneficiary(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.RelationTypeId, getData1.RelationTypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesCustomer.GetCustomerBeneficiary(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CustomerId, getData2.CustomerId)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.RelationTypeId, getData2.RelationTypeId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesCustomer.GetCustomerBeneficiarybyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestCustomerBeneficiary(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.CustomerId, updatedD1.CustomerId)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.RelationTypeId, updatedD1.RelationTypeId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListCustomerBeneficiary(t, ListCustomerBeneficiaryParams{
		CustomerId: updatedD1.CustomerId,
		Limit:      5,
		Offset:     0,
	})

	// Delete Data
	testDeleteCustomerBeneficiary(t, CreatedD1.Uuid)
	testDeleteCustomerBeneficiary(t, CreatedD2.Uuid)
}

func testListCustomerBeneficiary(t *testing.T, arg ListCustomerBeneficiaryParams) {

	customerBeneficiary, err := testQueriesCustomer.ListCustomerBeneficiary(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", customerBeneficiary)
	require.NotEmpty(t, customerBeneficiary)

}

func randomCustomerBeneficiary() CustomerBeneficiaryRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	cust, _ := testQueriesCustomer.GetCustomerbyAltId(context.Background(), "10001")
	ben, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "BeneficiaryType", 0, "Revocable")
	rel, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "RelationshipType", 0, "Contact Person")

	arg := CustomerBeneficiaryRequest{
		CustomerId: cust.Id,
		Series:     int16(util.RandomInt(1, 100)),
		// Iiid:           ii.Id,
		TypeId:         ben.Id,
		RelationTypeId: rel.Id,

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestCustomerBeneficiary(
	t *testing.T,
	d1 CustomerBeneficiaryRequest) model.CustomerBeneficiary {

	getData1, err := testQueriesCustomer.CreateCustomerBeneficiary(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.RelationTypeId, getData1.RelationTypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestCustomerBeneficiary(
	t *testing.T,
	d1 CustomerBeneficiaryRequest) model.CustomerBeneficiary {

	getData1, err := testQueriesCustomer.UpdateCustomerBeneficiary(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.CustomerId, getData1.CustomerId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.RelationTypeId, getData1.RelationTypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteCustomerBeneficiary(t *testing.T, uuid uuid.UUID) {
	err := testQueriesCustomer.DeleteCustomerBeneficiary(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesCustomer.GetCustomerBeneficiary(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
