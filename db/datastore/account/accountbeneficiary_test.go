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

func TestAccountBeneficiary(t *testing.T) {

	// Test Data
	d1 := randomAccountBeneficiary()
	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "E302U8-4059-1127199")
	d1.AccountId = acc.Id

	d2 := randomAccountBeneficiary()
	acc, _ = testQueriesAccount.GetAccountbyAltAcc(context.Background(), "E30304-4001-0000033")
	d2.AccountId = acc.Id

	// Test Create
	CreatedD1 := createTestAccountBeneficiary(t, d1)
	CreatedD2 := createTestAccountBeneficiary(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetAccountBeneficiary(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.BeneficiaryTypeId, getData1.BeneficiaryTypeId)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.RelationshipTypeId, getData1.RelationshipTypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetAccountBeneficiary(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AccountId, getData2.AccountId)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.BeneficiaryTypeId, getData2.BeneficiaryTypeId)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.RelationshipTypeId, getData2.RelationshipTypeId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetAccountBeneficiarybyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountBeneficiary(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.AccountId, updatedD1.AccountId)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.BeneficiaryTypeId, updatedD1.BeneficiaryTypeId)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.RelationshipTypeId, updatedD1.RelationshipTypeId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListAccountBeneficiary(t, ListAccountBeneficiaryParams{
		AccountId: updatedD1.AccountId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteAccountBeneficiary(t, CreatedD1.Uuid)
	testDeleteAccountBeneficiary(t, CreatedD2.Uuid)
}

func testListAccountBeneficiary(t *testing.T, arg ListAccountBeneficiaryParams) {

	accountBeneficiary, err := testQueriesAccount.ListAccountBeneficiary(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", accountBeneficiary)
	require.NotEmpty(t, accountBeneficiary)

}

func randomAccountBeneficiary() AccountBeneficiaryRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "BeneficiaryType", 0, "Revocable")
	rel, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "RelationshipType", 0, "Contact Person")
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")

	arg := AccountBeneficiaryRequest{
		// AccountId:      util.RandomInt(1, 100),
		Series:             int16(util.RandomInt32(1, 100)),
		BeneficiaryTypeId:  typ.Id,
		Iiid:               ii.Id,
		RelationshipTypeId: rel.Id,

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestAccountBeneficiary(
	t *testing.T,
	d1 AccountBeneficiaryRequest) model.AccountBeneficiary {

	getData1, err := testQueriesAccount.CreateAccountBeneficiary(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.BeneficiaryTypeId, getData1.BeneficiaryTypeId)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.RelationshipTypeId, getData1.RelationshipTypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestAccountBeneficiary(
	t *testing.T,
	d1 AccountBeneficiaryRequest) model.AccountBeneficiary {

	getData1, err := testQueriesAccount.UpdateAccountBeneficiary(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.BeneficiaryTypeId, getData1.BeneficiaryTypeId)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.RelationshipTypeId, getData1.RelationshipTypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAccountBeneficiary(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteAccountBeneficiary(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetAccountBeneficiary(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
