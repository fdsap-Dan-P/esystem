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

func TestActorGroupMember(t *testing.T) {

	// Test Data
	d1 := randomActorGroupMember()
	d2 := randomActorGroupMember()
	d2.Uuid = util.ToUUID("e6e514f0-62c9-4f8a-951f-5f0e30b6ce3c")
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "erick1421@gmail.com")
	d2.UserId = usr.Id

	// Test Create
	CreatedD1 := createTestActorGroupMember(t, d1)
	CreatedD2 := createTestActorGroupMember(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetActorGroupMember(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.ActorGroupId, getData1.ActorGroupId)
	require.Equal(t, d1.ActorGroup, getData1.ActorGroup)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetActorGroupMember(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Uuid, getData2.Uuid)
	require.Equal(t, d2.ActorGroupId, getData2.ActorGroupId)
	require.Equal(t, d2.ActorGroup, getData2.ActorGroup)
	require.Equal(t, d2.UserId, getData2.UserId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetActorGroupMemberbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestActorGroupMember(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Uuid, updatedD1.Uuid)
	require.Equal(t, updateD2.ActorGroupId, updatedD1.ActorGroupId)
	require.Equal(t, updateD2.ActorGroup, updatedD1.ActorGroup)
	require.Equal(t, updateD2.UserId, updatedD1.UserId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListActorGroupMember(t, ListActorGroupMemberParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteActorGroupMember(t, CreatedD1.Uuid)
	testDeleteActorGroupMember(t, CreatedD2.Uuid)
}

func testListActorGroupMember(t *testing.T, arg ListActorGroupMemberParams) {

	ActorGroupMember, err := testQueriesTransaction.ListActorGroupMember(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", ActorGroupMember)
	require.NotEmpty(t, ActorGroupMember)

}

func randomActorGroupMember() ActorGroupMemberRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	grp, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "actorgroup", 0, "Approver")
	usr, _ := testQueriesUser.GetUserbyName(context.Background(), "olive.mercado0609@gmail.com")

	arg := ActorGroupMemberRequest{
		Uuid: util.ToUUID("0786cbad-cd0f-47c6-8e57-e0d10f55ba73"),

		ActorGroupId: grp.Id,
		UserId:       usr.Id,
		OtherInfo:    sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestActorGroupMember(
	t *testing.T,
	d1 ActorGroupMemberRequest) model.ActorGroupMember {

	getData1, err := testQueriesTransaction.CreateActorGroupMember(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.ActorGroupId, getData1.ActorGroupId)
	require.Equal(t, d1.ActorGroup, getData1.ActorGroup)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)
	return getData1
}

func updateTestActorGroupMember(
	t *testing.T,
	d1 ActorGroupMemberRequest) model.ActorGroupMember {

	getData1, err := testQueriesTransaction.UpdateActorGroupMember(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.ActorGroupId, getData1.ActorGroupId)
	require.Equal(t, d1.ActorGroup, getData1.ActorGroup)
	require.Equal(t, d1.UserId, getData1.UserId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteActorGroupMember(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteActorGroupMember(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetActorGroupMemberbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
