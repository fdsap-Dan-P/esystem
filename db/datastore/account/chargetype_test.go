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

	"github.com/stretchr/testify/require"
)

func TestChargeType(t *testing.T) {

	// Test Data
	d1 := randomChargeType()
	d2 := randomChargeType()
	d2.Uuid = util.ToUUID("0c29d9bb-1427-4bb3-834d-0b9f176c1cfa")

	// Test Create
	CreatedD1 := createTestChargeType(t, d1)
	CreatedD2 := createTestChargeType(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetChargeType(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.ChargeType, getData1.ChargeType)
	require.Equal(t, d1.UnrealizedId, getData1.UnrealizedId)
	require.Equal(t, d1.RealizedId, getData1.RealizedId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetChargeType(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.ChargeType, getData2.ChargeType)
	require.Equal(t, d2.UnrealizedId, getData2.UnrealizedId)
	require.Equal(t, d2.RealizedId, getData2.RealizedId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetChargeTypebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	getData, err = testQueriesAccount.GetChargeTypebyName(context.Background(), CreatedD1.ChargeType)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.ChargeType = updateD2.ChargeType + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestChargeType(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.ChargeType, updatedD1.ChargeType)
	require.Equal(t, updateD2.UnrealizedId, updatedD1.UnrealizedId)
	require.Equal(t, updateD2.RealizedId, updatedD1.RealizedId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteChargeType(t, getData1.Id)
	testDeleteChargeType(t, getData2.Id)
}

func TestListChargeType(t *testing.T) {

	arg := ListChargeTypeParams{
		Limit:  5,
		Offset: 0,
	}

	chargeType, err := testQueriesAccount.ListChargeType(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", chargeType)
	require.NotEmpty(t, chargeType)

}

func randomChargeType() ChargeTypeRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	unRel, _ := testQueriesReference.GetChartofAccountbyTitle(context.Background(), "Service Charges/Fees")
	rel, _ := testQueriesReference.GetChartofAccountbyTitle(context.Background(), "Service Charges/Fees")
	// stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "LoanStatus", 0, "Current")

	arg := ChargeTypeRequest{
		Uuid:         util.ToUUID("8c13da86-afa2-49af-9463-1d5aedc7d055"),
		ChargeType:   util.RandomString(10),
		UnrealizedId: unRel.Id,
		RealizedId:   rel.Id,

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestChargeType(
	t *testing.T,
	d1 ChargeTypeRequest) model.ChargeType {

	getData1, err := testQueriesAccount.CreateChargeType(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ChargeType, getData1.ChargeType)
	require.Equal(t, d1.UnrealizedId, getData1.UnrealizedId)
	require.Equal(t, d1.RealizedId, getData1.RealizedId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestChargeType(
	t *testing.T,
	d1 ChargeTypeRequest) model.ChargeType {

	getData1, err := testQueriesAccount.UpdateChargeType(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ChargeType, getData1.ChargeType)
	require.Equal(t, d1.UnrealizedId, getData1.UnrealizedId)
	require.Equal(t, d1.RealizedId, getData1.RealizedId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteChargeType(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteChargeType(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetChargeType(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
