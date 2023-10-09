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

func TestQuestion(t *testing.T) {

	// Test Data
	d1 := randomQuestion()
	d2 := randomQuestion()
	d2.Uuid = util.ToUUID("bd5b3985-0c2b-492f-a5d6-371f06455c96")
	d2.Series = 2

	// Test Create
	CreatedD1 := createTestQuestion(t, d1)
	CreatedD2 := createTestQuestion(t, d2)

	// Update Data
	updateD2 := d2
	updateD2.Id = CreatedD2.Id

	// log.Println(updateD2)
	// updatedD1 := updateTestQuestion(t, updateD2)

	testListQuestion(t, ListQuestionParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteQuestion(t, CreatedD1.Uuid)
	testDeleteQuestion(t, CreatedD2.Uuid)
}

func testListQuestion(t *testing.T, arg ListQuestionParams) {

	Question, err := testQueriesSchool.ListQuestion(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Question)
	require.NotEmpty(t, Question)
}

func compareQuestionRequest(t *testing.T, mod model.Question, req QuestionRequest) {
	require.NotEmpty(t, mod)
	require.Equal(t, req.Uuid, mod.Uuid)
	require.Equal(t, req.QuestionaireId, mod.QuestionaireId)
	require.Equal(t, req.Series, mod.Series)
	require.Equal(t, req.TypeId, mod.TypeId)
	require.Equal(t, req.QuestionItem, mod.QuestionItem)
	require.JSONEq(t, req.Choices.String, mod.Choices.String)
	require.Equal(t, req.AnswerType, mod.AnswerType)
	require.Equal(t, req.ParentId, mod.ParentId)
	require.Equal(t, req.StatusId, mod.StatusId)
	require.Equal(t, req.Remarks, mod.Remarks)
	require.JSONEq(t, req.OtherInfo.String, mod.OtherInfo.String)
}

func compareQuestionInfo(t *testing.T, mod model.Question, info QuestionInfo) {
	require.NotEmpty(t, mod)
	require.Equal(t, info.Uuid, mod.Uuid)
	require.Equal(t, info.QuestionaireId, mod.QuestionaireId)
	require.Equal(t, info.Series, mod.Series)
	require.Equal(t, info.TypeId, mod.TypeId)
	require.Equal(t, info.QuestionItem, mod.QuestionItem)
	require.JSONEq(t, info.Choices.String, mod.Choices.String)
	require.Equal(t, info.AnswerType, mod.AnswerType)
	require.Equal(t, info.ParentId, mod.ParentId)
	require.Equal(t, info.StatusId, mod.StatusId)
	require.Equal(t, info.Remarks, mod.Remarks)
	require.JSONEq(t, info.OtherInfo.String, mod.OtherInfo.String)
}

func createTestQuestion(
	t *testing.T,
	req QuestionRequest) model.Question {

	mod, err := testQueriesSchool.CreateQuestion(context.Background(), req)
	require.NoError(t, err)
	compareQuestionRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetQuestion(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareQuestionInfo(t, mod, info)

	return mod
}

func updateTestQuestion(
	t *testing.T,
	req QuestionRequest) model.Question {

	mod, err := testQueriesSchool.UpdateQuestion(context.Background(), req)
	require.NoError(t, err)
	compareQuestionRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetQuestion(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareQuestionInfo(t, mod, info)

	return mod
}

func testDeleteQuestion(t *testing.T, uuid uuid.UUID) {
	err := testQueriesSchool.DeleteQuestion(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesSchool.GetQuestionbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func randomQuestion() QuestionRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectType", 0, "Regular")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectStatus", 0, "Current")

	corsP, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "QuestionType", 0, "Pre-College")

	ques, _ := testQueriesSchool.CreateQuestionaire(context.Background(), randomQuestionaire())

	arg := QuestionRequest{
		Uuid:           util.ToUUID("5bdb5178-9f2b-496d-9151-73e08aaafe64"),
		QuestionaireId: ques.Id,
		Series:         1,
		TypeId:         typ.Id,
		QuestionItem:   util.RandomString(20),
		Choices:        sql.NullString(sql.NullString{String: string(info), Valid: true}),
		AnswerType:     util.SetNullString("11"),
		// ParentId,
		StatusId: stat.Id,

		Remarks:   sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	log.Printf("QuestionRequest corsP: %v", corsP)
	log.Printf("QuestionRequest: %v", arg)
	return arg
}
