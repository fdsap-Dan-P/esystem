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

func TestActionLink(t *testing.T) {

	// Test Data
	d1 := randomActionLink()
	d2 := randomActionLink()
	d2.Uuid = util.ToUUID("e6e514f0-62c9-4f8a-951f-5f0e30b6ce3c")

	// Test Create
	CreatedD1 := createTestActionLink(t, d1)
	CreatedD2 := createTestActionLink(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetActionLink(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.EventName, getData1.EventName)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.EndPointCall, getData1.EndPointCall)
	require.Equal(t, d1.ServerId, getData1.ServerId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetActionLink(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Uuid, getData2.Uuid)
	require.Equal(t, d2.EventName, getData2.EventName)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.EndPointCall, getData2.EndPointCall)
	require.Equal(t, d2.ServerId, getData2.ServerId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetActionLinkbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestActionLink(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Uuid, updatedD1.Uuid)
	require.Equal(t, updateD2.EventName, updatedD1.EventName)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.EndPointCall, updatedD1.EndPointCall)
	require.Equal(t, updateD2.ServerId, updatedD1.ServerId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListActionLink(t, ListActionLinkParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteActionLink(t, CreatedD1.Uuid)
	testDeleteActionLink(t, CreatedD2.Uuid)
}

func testListActionLink(t *testing.T, arg ListActionLinkParams) {

	actionLink, err := testQueriesTransaction.ListActionLink(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", actionLink)
	require.NotEmpty(t, actionLink)

}

func randomActionLink() ActionLinkRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "actionlist", 0, "Posted")

	svr, _ := testQueriesDocument.CreateServer(context.Background(), randomServer())

	arg := ActionLinkRequest{
		Uuid:         util.ToUUID("0786cbad-cd0f-47c6-8e57-e0d10f55ba73"),
		EventName:    util.RandomString(20),
		TypeId:       typ.Id,
		EndPointCall: util.RandomNullString(20),
		ServerId:     util.SetNullInt64(svr.Id),
		OtherInfo:    sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestActionLink(
	t *testing.T,
	d1 ActionLinkRequest) model.ActionLink {

	getData1, err := testQueriesTransaction.CreateActionLink(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.EventName, getData1.EventName)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.EndPointCall, getData1.EndPointCall)
	require.Equal(t, d1.ServerId, getData1.ServerId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)
	return getData1
}

func updateTestActionLink(
	t *testing.T,
	d1 ActionLinkRequest) model.ActionLink {

	getData1, err := testQueriesTransaction.UpdateActionLink(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.EventName, getData1.EventName)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.EndPointCall, getData1.EndPointCall)
	require.Equal(t, d1.ServerId, getData1.ServerId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteActionLink(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteActionLink(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetActionLinkbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
