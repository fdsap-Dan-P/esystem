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

func TestQuestionaire(t *testing.T) {

	// Test Data
	d1 := randomQuestionaire()
	d2 := randomQuestionaire()
	d2.Uuid = util.ToUUID("c9fa1453-11fd-4eb4-948f-6a9d76466da7")

	// Test Create
	CreatedD1 := createTestQuestionaire(t, d1)
	CreatedD2 := createTestQuestionaire(t, d2)

	// Update Data
	updateD2 := d2
	updateD2.Id = CreatedD2.Id

	// log.Println(updateD2)
	// updatedD1 := updateTestQuestionaire(t, updateD2)

	testListQuestionaire(t, ListQuestionaireParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteQuestionaire(t, CreatedD1.Uuid)
	testDeleteQuestionaire(t, CreatedD2.Uuid)
}

func testListQuestionaire(t *testing.T, arg ListQuestionaireParams) {

	Questionaire, err := testQueriesSchool.ListQuestionaire(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Questionaire)
	require.NotEmpty(t, Questionaire)
}

func compareQuestionaireRequest(t *testing.T, mod model.Questionaire, req QuestionaireRequest) {
	require.NotEmpty(t, mod)
	require.Equal(t, req.Uuid, mod.Uuid)
	require.Equal(t, req.Code, mod.Code)
	require.Equal(t, req.Version, mod.Version)
	require.Equal(t, req.Title, mod.Title)
	require.Equal(t, req.TypeId, mod.TypeId)
	require.Equal(t, req.SubjectId, mod.SubjectId)
	require.Equal(t, req.DateRevised.Format("2006-01-02"), mod.DateRevised.Format("2006-01-02"))
	require.Equal(t, req.OfficeId, mod.OfficeId)
	require.Equal(t, req.AuthorId, mod.AuthorId)
	require.Equal(t, req.StatusId, mod.StatusId)
	require.JSONEq(t, req.PointEquivalent.String, mod.PointEquivalent.String)
	require.Equal(t, req.Remarks, mod.Remarks)
	require.JSONEq(t, req.OtherInfo.String, mod.OtherInfo.String)
}

func compareQuestionaireInfo(t *testing.T, mod model.Questionaire, info QuestionaireInfo) {
	require.NotEmpty(t, mod)
	require.Equal(t, info.Uuid, mod.Uuid)
	require.Equal(t, info.Code, mod.Code)
	require.Equal(t, info.Version, mod.Version)
	require.Equal(t, info.Title, mod.Title)
	require.Equal(t, info.TypeId, mod.TypeId)
	require.Equal(t, info.SubjectId, mod.SubjectId)
	require.Equal(t, info.DateRevised.Format("2006-01-02"), mod.DateRevised.Format("2006-01-02"))
	require.Equal(t, info.OfficeId, mod.OfficeId)
	require.Equal(t, info.AuthorId, mod.AuthorId)
	require.Equal(t, info.StatusId, mod.StatusId)
	require.JSONEq(t, info.PointEquivalent.String, mod.PointEquivalent.String)
	require.Equal(t, info.Remarks, mod.Remarks)
	require.JSONEq(t, info.OtherInfo.String, mod.OtherInfo.String)
}

func createTestQuestionaire(
	t *testing.T,
	req QuestionaireRequest) model.Questionaire {

	mod, err := testQueriesSchool.CreateQuestionaire(context.Background(), req)
	require.NoError(t, err)
	compareQuestionaireRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetQuestionaire(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareQuestionaireInfo(t, mod, info)

	return mod
}

func updateTestQuestionaire(
	t *testing.T,
	req QuestionaireRequest) model.Questionaire {

	mod, err := testQueriesSchool.UpdateQuestionaire(context.Background(), req)
	require.NoError(t, err)
	compareQuestionaireRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetQuestionaire(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareQuestionaireInfo(t, mod, info)

	return mod
}

func testDeleteQuestionaire(t *testing.T, uuid uuid.UUID) {
	err := testQueriesSchool.DeleteQuestionaire(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesSchool.GetQuestionairebyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func randomQuestionaire() QuestionaireRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectType", 0, "Regular")
	subj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Subject", 0, "English")
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectStatus", 0, "Current")

	corsP, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "QuestionaireType", 0, "Pre-College")
	// cors, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Questionaires", corsP.Id, "Elementary")
	// stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "QuestionaireStatus", 0, "Current")

	arg := QuestionaireRequest{
		Uuid:            util.ToUUID("0786cbad-cd0f-47c6-8e57-e0d10f55ba73"),
		Code:            util.RandomString(10),
		Version:         util.RandomInt(10, 100),
		Title:           util.RandomString(10),
		TypeId:          typ.Id,
		SubjectId:       sql.NullInt64{Int64: subj.Id, Valid: true},
		DateRevised:     util.RandomDate(),
		OfficeId:        sql.NullInt64{Int64: ofc.Id, Valid: true},
		AuthorId:        sql.NullInt64{Int64: ii.Id, Valid: true},
		StatusId:        stat.Id,
		PointEquivalent: sql.NullString(sql.NullString{String: string(info), Valid: true}),

		Remarks:   sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	log.Printf("QuestionaireRequest corsP: %v", corsP)
	log.Printf("QuestionaireRequest: %v", arg)
	return arg
}
