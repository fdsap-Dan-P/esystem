package db

import (
	"context"
	"database/sql"
	"log"

	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSubjectEvent(t *testing.T) {

	// Test Data

	d1 := randomSubjectEvent(t, "English", util.ToUUID("87e3362e-842b-46a4-900a-a9c243f1de66"))
	d2 := randomSubjectEvent(t, "Mathematics", util.ToUUID("33d4b023-3e3b-4b7f-bbf3-65a4828a9077"))
	d2.Uuid = util.ToUUID("67c9b220-1f6c-4718-8549-687bb1d86db3")

	// Test Create
	CreatedD1 := createTestSubjectEvent(t, d1)
	CreatedD2 := createTestSubjectEvent(t, d2)

	infoUuid, err1 := testQueriesSchool.GetSubjectEventbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err1)
	compareSubjectEventInfo(t, CreatedD1, infoUuid)

	// Update Data
	updateD2 := d2
	updateD2.Id = CreatedD2.Id
	updateD2.ItemCount = util.RandomMoney()

	updateTestSubjectEvent(t, updateD2)

	// log.Println(updateD2)

	testListSubjectEvent(t, ListSubjectEventParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteSubjectEvent(t, CreatedD1.Uuid)
	testDeleteSubjectEvent(t, CreatedD2.Uuid)
}

func testListSubjectEvent(t *testing.T, arg ListSubjectEventParams) {

	subjectEvent, err := testQueriesSchool.ListSubjectEvent(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", subjectEvent)
	require.NotEmpty(t, subjectEvent)
}

func compareSubjectEventRequest(t *testing.T, mod model.SubjectEvent, req SubjectEventRequest) {
	require.NotEmpty(t, mod)
	require.Equal(t, req.Uuid, mod.Uuid)
	require.Equal(t, req.TypeId, mod.TypeId)
	require.Equal(t, req.TicketItemId, mod.TicketItemId)
	require.Equal(t, req.Iiid, mod.Iiid)
	require.Equal(t, req.SectionSubjectId, mod.SectionSubjectId)
	require.Equal(t, req.EventDate.Format("2006-01-02"), mod.EventDate.Format("2006-01-02"))
	require.Equal(t, req.GradingPeriod, mod.GradingPeriod)
	require.Equal(t, req.ItemCount.String(), mod.ItemCount.String())
	require.Equal(t, req.StatusId, mod.StatusId)
	require.Equal(t, req.Remarks, mod.Remarks)
	require.JSONEq(t, req.OtherInfo.String, mod.OtherInfo.String)
}

func compareSubjectEventInfo(t *testing.T, mod model.SubjectEvent, info SubjectEventInfo) {
	require.NotEmpty(t, mod)
	require.Equal(t, info.Uuid, mod.Uuid)
	require.Equal(t, info.TypeId, mod.TypeId)
	require.Equal(t, info.TicketItemId, mod.TicketItemId)
	require.Equal(t, info.Iiid, mod.Iiid)
	require.Equal(t, info.SectionSubjectId, mod.SectionSubjectId)
	require.Equal(t, info.EventDate.Format("2006-01-02"), mod.EventDate.Format("2006-01-02"))
	require.Equal(t, info.GradingPeriod, mod.GradingPeriod)
	require.Equal(t, info.ItemCount.String(), mod.ItemCount.String())
	require.Equal(t, info.StatusId, mod.StatusId)
	require.Equal(t, info.Remarks, mod.Remarks)
	require.JSONEq(t, info.OtherInfo.String, mod.OtherInfo.String)
}
func createTestSubjectEvent(
	t *testing.T,
	req SubjectEventRequest) model.SubjectEvent {

	mod, err := testQueriesSchool.CreateSubjectEvent(context.Background(), req)
	require.NoError(t, err)
	compareSubjectEventRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetSubjectEvent(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareSubjectEventInfo(t, mod, info)

	return mod
}

func updateTestSubjectEvent(
	t *testing.T,
	req SubjectEventRequest) model.SubjectEvent {

	mod, err := testQueriesSchool.UpdateSubjectEvent(context.Background(), req)
	require.NoError(t, err)
	compareSubjectEventRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetSubjectEvent(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareSubjectEventInfo(t, mod, info)

	return mod
}

func testDeleteSubjectEvent(t *testing.T, uuid uuid.UUID) {
	err := testQueriesSchool.DeleteSubjectEvent(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesSchool.GetSubjectEventbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func randomSubjectEvent(t *testing.T, subject string, uuid uuid.UUID) SubjectEventRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "1059")

	secSubj, _ := testQueriesSchool.CreateSectionSubject(context.Background(), randomSectionSubject(t, subject, uuid))

	// subj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectEvent", 0, "English")
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectEvent", 0, "Written Works")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectStatus", 0, "Current")
	d := RandomTicketItem()
	tic, er := testQueriesTransaction.CreateTicketItem(context.Background(), d)

	log.Printf("testQueriesTransaction: Error >>> %v", er)
	log.Printf("testQueriesTransaction: d-- %v", d)
	log.Printf("Ticket: %v", tic)

	// Section_Subject

	arg := SubjectEventRequest{
		Uuid:             util.ToUUID("c1606f72-b59e-42f5-b917-ec073eae3489"),
		TypeId:           typ.Id,
		TicketItemId:     tic.Id,
		Iiid:             ii.Id,
		SectionSubjectId: secSubj.Id,
		EventDate:        util.SetDate("2021-01-21"),
		GradingPeriod:    1,
		ItemCount:        util.SetDecimal("100"),
		StatusId:         stat.Id,
		Remarks:          util.RandomNullString(10),
		OtherInfo:        sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}
