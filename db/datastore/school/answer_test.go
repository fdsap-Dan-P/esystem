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

func TestAnswer(t *testing.T) {

	// Test Data
	d1 := randomAnswer(t, "English", util.ToUUID("87e3362e-842b-46a4-900a-a9c243f1de66"))
	d2 := randomAnswer(t, "Mathematics", util.ToUUID("33d4b023-3e3b-4b7f-bbf3-65a4828a9077"))

	d2.Uuid = util.ToUUID("c9fa1453-11fd-4eb4-948f-6a9d76466da7")

	// Test Create
	CreatedD1 := createTestAnswer(t, d1)
	CreatedD2 := createTestAnswer(t, d2)

	// Update Data
	updateD2 := d2
	updateD2.Id = CreatedD2.Id

	// log.Println(updateD2)
	// updatedD1 := updateTestAnswer(t, updateD2)

	testListAnswer(t, ListAnswerParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteAnswer(t, CreatedD1.Uuid)
	testDeleteAnswer(t, CreatedD2.Uuid)
}

func testListAnswer(t *testing.T, arg ListAnswerParams) {

	Answer, err := testQueriesSchool.ListAnswer(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Answer)
	require.NotEmpty(t, Answer)
}

func compareAnswerRequest(t *testing.T, mod model.Answer, req AnswerRequest) {
	require.NotEmpty(t, mod)
	require.Equal(t, req.Uuid, mod.Uuid)
	require.Equal(t, req.SubjectEventId, mod.SubjectEventId)
	require.Equal(t, req.QuestionId, mod.QuestionId)
	require.JSONEq(t, req.Answers.String, mod.Answers.String)
	require.Equal(t, req.Points.String(), mod.Points.String())
	require.Equal(t, req.Remarks, mod.Remarks)
	require.JSONEq(t, req.OtherInfo.String, mod.OtherInfo.String)
}

func compareAnswerInfo(t *testing.T, mod model.Answer, info AnswerInfo) {
	require.NotEmpty(t, mod)
	require.Equal(t, info.Uuid, mod.Uuid)
	require.Equal(t, info.SubjectEventId, mod.SubjectEventId)
	require.Equal(t, info.QuestionId, mod.QuestionId)
	require.JSONEq(t, info.Answers.String, mod.Answers.String)
	require.Equal(t, info.Points.String(), mod.Points.String())
	require.Equal(t, info.Remarks, mod.Remarks)
	require.JSONEq(t, info.OtherInfo.String, mod.OtherInfo.String)
}
func createTestAnswer(
	t *testing.T,
	req AnswerRequest) model.Answer {

	mod, err := testQueriesSchool.CreateAnswer(context.Background(), req)
	require.NoError(t, err)
	compareAnswerRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetAnswer(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareAnswerInfo(t, mod, info)

	return mod
}

func updateTestAnswer(
	t *testing.T,
	req AnswerRequest) model.Answer {

	mod, err := testQueriesSchool.UpdateAnswer(context.Background(), req)
	require.NoError(t, err)
	compareAnswerRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetAnswer(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareAnswerInfo(t, mod, info)

	return mod
}

func testDeleteAnswer(t *testing.T, uuid uuid.UUID) {
	err := testQueriesSchool.DeleteAnswer(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesSchool.GetAnswerbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func randomAnswer(t *testing.T, subject string, uuid uuid.UUID) AnswerRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	corsP, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AnswerType", 0, "Pre-College")
	// cors, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Answers", corsP.Id, "Elementary")
	// stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AnswerStatus", 0, "Current")
	subjEv, _ := testQueriesSchool.CreateSubjectEvent(context.Background(), randomSubjectEvent(t, subject, uuid))
	ques, _ := testQueriesSchool.CreateQuestion(context.Background(), randomQuestion())

	arg := AnswerRequest{
		Uuid:           util.ToUUID("0786cbad-cd0f-47c6-8e57-e0d10f55ba73"),
		SubjectEventId: subjEv.Id,
		QuestionId:     ques.Id,
		Answers:        sql.NullString(sql.NullString{String: string(info), Valid: true}),
		Points:         util.RandomMoney(),
		Remarks:        sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		OtherInfo:      sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	log.Printf("AnswerRequest corsP: %v", corsP)
	log.Printf("AnswerRequest: %v", arg)
	return arg
}
