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

func TestRelation(t *testing.T) {

	// Test Data
	d1 := randomRelation()
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")
	rel, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "1001")
	d1.Iiid = ii.Id
	d1.RelationIiid = rel.Id

	d2 := randomRelation()
	ii, _ = testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "101")
	rel, _ = testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "1002")
	d2.Iiid = ii.Id
	d2.RelationIiid = rel.Id

	// Test Create
	CreatedD1 := createTestRelation(t, d1)
	CreatedD2 := createTestRelation(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetRelation(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.RelationIiid, getData1.RelationIiid)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.RelationDate.Time.Format("2006-01-02"), getData1.RelationDate.Time.Format("2006-01-02"))
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesIdentity.GetRelation(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.RelationIiid, getData2.RelationIiid)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.RelationDate.Time.Format("2006-01-02"), getData2.RelationDate.Time.Format("2006-01-02"))
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesIdentity.GetRelationbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// getData, err = testQueriesIdentity.GetRelationbyName(context.Background(), CreatedD1.Name)
	// require.NoError(t, err)
	// require.NotEmpty(t, getData)
	// require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	// fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestRelation(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.RelationIiid, updatedD1.RelationIiid)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.RelationDate.Time.Format("2006-01-02"), updatedD1.RelationDate.Time.Format("2006-01-02"))
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListRelation(t, ListRelationParams{
		Iiid:   updatedD1.Iiid,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteRelation(t, CreatedD1.Uuid)
	testDeleteRelation(t, CreatedD2.Uuid)
}

func testListRelation(t *testing.T, arg ListRelationParams) {

	relation, err := testQueriesIdentity.ListRelation(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", relation)
	require.NotEmpty(t, relation)

}

func randomRelation() RelationRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "RelationshipType", 0, "Contact Person")

	arg := RelationRequest{
		// Iiid:           util.RandomInt(1, 100),
		Series: int16(util.RandomInt32(1, 100)),
		// RelationIiid:   util.RandomInt(1, 100),
		TypeId:       obj.Id,
		RelationDate: sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestRelation(
	t *testing.T,
	d1 RelationRequest) model.Relation {

	getData1, err := testQueriesIdentity.CreateRelation(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.RelationIiid, getData1.RelationIiid)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.RelationDate.Time.Format("2006-01-02"), getData1.RelationDate.Time.Format("2006-01-02"))
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestRelation(
	t *testing.T,
	d1 RelationRequest) model.Relation {

	getData1, err := testQueriesIdentity.UpdateRelation(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.RelationIiid, getData1.RelationIiid)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.RelationDate.Time.Format("2006-01-02"), getData1.RelationDate.Time.Format("2006-01-02"))
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteRelation(t *testing.T, uuid uuid.UUID) {
	err := testQueriesIdentity.DeleteRelation(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetRelation(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
