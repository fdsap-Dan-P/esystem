package db

import (
	"context"
	"database/sql"
	"log"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestActorGroupAction(t *testing.T) {

	// Test Data
	d1 := randomActorGroupAction()
	d2 := randomActorGroupAction()
	d2.Uuid = util.ToUUID("e6e514f0-62c9-4f8a-951f-5f0e30b6ce3c")
	grp, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "actorgroup", 0, "Checker")
	d2.ActorGroupId = grp.Id

	log.Printf("randomActorGroupAction %v", d1)

	// Test Create
	CreatedD1 := createTestActorGroupAction(t, d1)
	CreatedD2 := createTestActorGroupAction(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetActorGroupActionbyUuid(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.ActorGroupId, getData1.ActorGroupId)
	require.Equal(t, d1.TicketTypeActionId, getData1.TicketTypeActionId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetActorGroupActionbyUuid(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.ActorGroupId, getData2.ActorGroupId)
	require.Equal(t, d2.TicketTypeActionId, getData2.TicketTypeActionId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetActorGroupActionbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestActorGroupAction(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.ActorGroupId, updatedD1.ActorGroupId)
	require.Equal(t, updateD2.TicketTypeActionId, updatedD1.TicketTypeActionId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListActorGroupAction(t, ListActorGroupActionParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteActorGroupAction(t, CreatedD1.Uuid)
	testDeleteActorGroupAction(t, CreatedD2.Uuid)
}

func testListActorGroupAction(t *testing.T, arg ListActorGroupActionParams) {

	ActorGroupAction, err := testQueriesTransaction.ListActorGroupAction(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", ActorGroupAction)
	require.NotEmpty(t, ActorGroupAction)

}

func randomActorGroupAction() ActorGroupActionRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	grp, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "actorgroup", 0, "Approver")

	d1 := RandomTicketTypeAction()
	act, err := testQueriesTransaction.CreateTicketTypeAction(context.Background(), RandomTicketTypeAction())
	log.Printf("d1 %v", d1)
	log.Printf("act %v", act)
	log.Printf("error %v", err)

	arg := ActorGroupActionRequest{
		Uuid:         util.ToUUID("0786cbad-cd0f-47c6-8e57-e0d10f55ba73"),
		ActorGroupId: grp.Id,

		TicketTypeActionId: act.Id,

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestActorGroupAction(
	t *testing.T,
	d1 ActorGroupActionRequest) model.ActorGroupAction {

	getData1, err := testQueriesTransaction.CreateActorGroupAction(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ActorGroupId, getData1.ActorGroupId)
	require.Equal(t, d1.TicketTypeActionId, getData1.TicketTypeActionId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)
	return getData1
}

func updateTestActorGroupAction(
	t *testing.T,
	d1 ActorGroupActionRequest) model.ActorGroupAction {

	getData1, err := testQueriesTransaction.UpdateActorGroupAction(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ActorGroupId, getData1.ActorGroupId)
	require.Equal(t, d1.TicketTypeActionId, getData1.TicketTypeActionId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteActorGroupAction(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteActorGroupAction(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetActorGroupActionbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
