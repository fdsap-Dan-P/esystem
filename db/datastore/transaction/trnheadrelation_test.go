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

func TestTrnHeadRelation(t *testing.T) {

	// Test Data
	d1 := randomTrnHeadRelation()
	trn, _ := testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("2af90d74-3bee-48c5-8935-443edafb8f5a"))
	d1.TrnHeadId = trn.Id

	d2 := randomTrnHeadRelation()
	trn, _ = testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("26dfab18-f80b-46cf-9c54-be79d4fc5d23"))
	d2.TrnHeadId = trn.Id

	// Test Create
	CreatedD1 := createTestTrnHeadRelation(t, d1)
	CreatedD2 := createTestTrnHeadRelation(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetTrnHeadRelation(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.RelatedId, getData1.RelatedId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetTrnHeadRelation(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TrnHeadId, getData2.TrnHeadId)
	require.Equal(t, d2.RelatedId, getData2.RelatedId)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetTrnHeadRelationbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestTrnHeadRelation(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TrnHeadId, updatedD1.TrnHeadId)
	require.Equal(t, updateD2.RelatedId, updatedD1.RelatedId)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.Remarks, updatedD1.Remarks)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListTrnHeadRelation(t, ListTrnHeadRelationParams{
		TrnHeadId: updatedD1.TrnHeadId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteTrnHeadRelation(t, CreatedD1.Uuid)
	testDeleteTrnHeadRelation(t, CreatedD2.Uuid)
}

func testListTrnHeadRelation(t *testing.T, arg ListTrnHeadRelationParams) {

	trnHeadRelation, err := testQueriesTransaction.ListTrnHeadRelation(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", trnHeadRelation)
	require.NotEmpty(t, trnHeadRelation)

}

func randomTrnHeadRelation() TrnHeadRelationRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	trn, _ := testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("3793422c-eb9f-49f0-9ec6-e5cf80caac25"))
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "trnheadaction", 0, "Automated")

	arg := TrnHeadRelationRequest{
		// TrnHeadId: trn.Id,
		RelatedId: trn.Id,
		TypeId:    typ.Id,
		Remarks:   sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestTrnHeadRelation(
	t *testing.T,
	d1 TrnHeadRelationRequest) model.TrnHeadRelation {

	getData1, err := testQueriesTransaction.CreateTrnHeadRelation(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.RelatedId, getData1.RelatedId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestTrnHeadRelation(
	t *testing.T,
	d1 TrnHeadRelationRequest) model.TrnHeadRelation {

	getData1, err := testQueriesTransaction.UpdateTrnHeadRelation(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.RelatedId, getData1.RelatedId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteTrnHeadRelation(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteTrnHeadRelation(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetTrnHeadRelation(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
