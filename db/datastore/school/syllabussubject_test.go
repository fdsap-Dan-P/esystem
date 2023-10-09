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

func TestSyllabusSubject(t *testing.T) {

	// Test Data
	d1 := randomSyllabusSubject(t, "English", util.ToUUID("68fc7186-4dae-488e-9e0f-0ceb7ea2c4e0"))
	d2 := randomSyllabusSubject(t, "Mathematics", util.ToUUID("76f7cc6d-e498-447c-a3ab-45b81b4650d3"))
	d2.Uuid = util.ToUUID("610336d8-51ff-4e1b-82e7-dc03958e75c8")

	// Test Create
	CreatedD1 := createTestSyllabusSubject(t, d1)
	CreatedD2 := createTestSyllabusSubject(t, d2)

	infoUuid, err1 := testQueriesSchool.GetSyllabusSubjectbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err1)

	log.Printf("GetSyllabusSubjectbyUuid CreatedD1 %v", CreatedD1)
	log.Printf("GetSyllabusSubjectbyUuid infoUuid %v", infoUuid)

	compareSyllabusSubjectInfo(t, CreatedD1, infoUuid)

	// Update Data
	updateD2 := d2
	updateD2.Id = CreatedD2.Id
	updateD2.Remarks = updateD2.Remarks + "-Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestSyllabusSubject(t, updateD2)

	testListSyllabusSubject(t, ListSyllabusSubjectParams{
		SyllabusId: updatedD1.SyllabusId,
		Limit:      5,
		Offset:     0,
	})

	// testSyllabusSubjectInfo(t)

	// Delete Data
	testDeleteSyllabusSubject(t, CreatedD1.Uuid)
	testDeleteSyllabusSubject(t, CreatedD2.Uuid)
}

func testListSyllabusSubject(t *testing.T, arg ListSyllabusSubjectParams) {

	SyllabusSubject, err := testQueriesSchool.ListSyllabusSubject(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", SyllabusSubject)
	require.NotEmpty(t, SyllabusSubject)
}

func compareSyllabusSubjectRequest(t *testing.T, mod model.SyllabusSubject, req SyllabusSubjectRequest) {
	require.NotEmpty(t, mod)
	require.Equal(t, req.Uuid, mod.Uuid)
	require.Equal(t, req.SyllabusId, mod.SyllabusId)
	require.Equal(t, req.SubjectId, mod.SubjectId)
	require.Equal(t, req.Units, mod.Units)
	require.Equal(t, req.TypeId, mod.TypeId)
	require.Equal(t, req.StatusId, mod.StatusId)
	require.Equal(t, req.Remarks, mod.Remarks)
	require.JSONEq(t, req.OtherInfo.String, mod.OtherInfo.String)
}

func compareSyllabusSubjectInfo(t *testing.T, mod model.SyllabusSubject, info SyllabusSubjectInfo) {
	require.NotEmpty(t, mod)
	require.Equal(t, info.Uuid, mod.Uuid)
	require.Equal(t, info.SyllabusId, mod.SyllabusId)
	require.Equal(t, info.SubjectId, mod.SubjectId)
	require.Equal(t, info.Units, mod.Units)
	require.Equal(t, info.TypeId, mod.TypeId)
	require.Equal(t, info.StatusId, mod.StatusId)
	require.Equal(t, info.Remarks, mod.Remarks)
	require.JSONEq(t, info.OtherInfo.String, mod.OtherInfo.String)
}
func createTestSyllabusSubject(
	t *testing.T,
	req SyllabusSubjectRequest) model.SyllabusSubject {

	mod, err := testQueriesSchool.CreateSyllabusSubject(context.Background(), req)
	require.NoError(t, err)
	compareSyllabusSubjectRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetSyllabusSubject(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareSyllabusSubjectInfo(t, mod, info)

	return mod
}

func updateTestSyllabusSubject(
	t *testing.T,
	req SyllabusSubjectRequest) model.SyllabusSubject {

	mod, err := testQueriesSchool.UpdateSyllabusSubject(context.Background(), req)
	require.NoError(t, err)
	compareSyllabusSubjectRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetSyllabusSubject(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareSyllabusSubjectInfo(t, mod, info)

	return mod
}

func testDeleteSyllabusSubject(t *testing.T, uuid uuid.UUID) {
	err := testQueriesSchool.DeleteSyllabusSubject(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesSchool.GetSyllabusSubjectbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func randomSyllabusSubject(t *testing.T, subject string, uuid uuid.UUID) SyllabusSubjectRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectType", 0, "Regular")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectStatus", 0, "Current")
	// adv, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SchoolSemister", 0, "Full Year")

	subj, _ := testQueriesSchool.CreateSubject(context.Background(), randomSubject(subject, uuid))
	syl, _ := testQueriesSchool.CreateSyllabus(context.Background(), randomSyllabus(t))

	log.Printf("subject %v %v", subj, subject)
	log.Printf("syl %v", syl)

	arg := SyllabusSubjectRequest{
		Uuid:       util.ToUUID("729fbb22-4a5c-4da0-a566-c68364788a35"),
		SubjectId:  subj.Id,
		SyllabusId: syl.Id,
		Units:      100,
		TypeId:     typ.Id,
		StatusId:   stat.Id,
		Remarks:    util.RandomString(10),
		OtherInfo:  sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}

	log.Printf("ARG %v", arg)
	log.Printf("subjectId %v", arg.SubjectId)
	log.Printf("SyllabusId %v", arg.SyllabusId)
	return arg
}

// func testSyllabusSubjectInfo(t *testing.T) SyllabusSubjectInfo {
// 	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
// 	info, _ := json.Marshal(otherInfo)
// 	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectType", 0, "Regular")
// 	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectStatus", 0, "Current")
// 	// adv, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SchoolSemister", 0, "Full Year")

// 	sec := testSchoolSectionData(t)
// 	secInfo, _ := testQueriesSchool.GetSchoolSectionbyUuid(context.Background(), sec.Uuid)
// 	// Test Create
// 	if secInfo.Id == 0 {
// 		syl := createTestSchoolSection(t, sec)
// 		secInfo, _ = testQueriesSchool.GetSchoolSectionbyUuid(context.Background(), syl.Uuid)
// 	}

// 	subj := testSubjectInfo(t, "English", util.ToUUID("3bdd00c6-dffb-4b7b-9b02-bf358205a690"))

// 	arg := SyllabusSubjectRequest{
// 		Uuid:      util.ToUUID("ed6a63f9-44bd-4aef-87ed-8e5538f392f8"),
// 		SubjectId: subj.Id,
// 		TypeId:    util.SetNullInt64(typ.Id),
// 		StatusId:  stat.Id,
// 		Remarks:   util.RandomString(10),
// 		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
// 	}

// 	syl := testSyllabusData(t)
// 	secSubjInfo, _ := testQueriesSchool.CreateSyllabus(context.Background(), syl)

// 	arg.SyllabusId = secSubjInfo.Id
// 	return arg
// }
