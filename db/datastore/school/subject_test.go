package db

import (
	"context"
	"database/sql"

	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSubject(t *testing.T) {

	// Test Data

	d1 := randomSubject("English", util.ToUUID("68fc7186-4dae-488e-9e0f-0ceb7ea2c4e0"))
	d2 := randomSubject("Mathematics", util.ToUUID("76f7cc6d-e498-447c-a3ab-45b81b4650d3"))

	// Test Create
	CreatedD1 := createTestSubject(t, d1)
	CreatedD2 := createTestSubject(t, d2)

	infoUuid, err1 := testQueriesSchool.GetSubjectbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err1)
	compareSubjectInfo(t, CreatedD1, infoUuid)

	// Update Data
	updateD2 := d2
	updateD2.Id = CreatedD2.Id
	updateD2.SubjectTitle = updateD2.SubjectTitle + "-Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestSubject(t, updateD2)

	testListSubject(t, ListSubjectParams{
		SubjectRefId: updatedD1.SubjectRefId,
		Limit:        5,
		Offset:       0,
	})

	// Delete Data
	testDeleteSubject(t, CreatedD1.Id)
	testDeleteSubject(t, CreatedD2.Id)
}

func testListSubject(t *testing.T, arg ListSubjectParams) {

	subject, err := testQueriesSchool.ListSubject(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", subject)
	require.NotEmpty(t, subject)
}

func compareSubjectRequest(t *testing.T, mod model.Subject, req SubjectRequest) {
	require.NotEmpty(t, mod)
	require.Equal(t, req.SubjectTitle, mod.SubjectTitle)
	require.Equal(t, req.SubjectRefId, mod.SubjectRefId)
	require.Equal(t, req.TypeId, mod.TypeId)
	require.Equal(t, req.Remarks, mod.Remarks)
	require.JSONEq(t, req.OtherInfo.String, mod.OtherInfo.String)
}

func compareSubjectInfo(t *testing.T, mod model.Subject, info SubjectInfo) {
	require.NotEmpty(t, mod)
	require.Equal(t, info.SubjectTitle, mod.SubjectTitle)
	require.Equal(t, info.SubjectRefId, mod.SubjectRefId)
	require.Equal(t, info.TypeId, mod.TypeId)
	require.Equal(t, info.Remarks, mod.Remarks)
	require.JSONEq(t, info.OtherInfo.String, mod.OtherInfo.String)
}
func createTestSubject(
	t *testing.T,
	req SubjectRequest) model.Subject {

	mod, err := testQueriesSchool.CreateSubject(context.Background(), req)
	require.NoError(t, err)
	compareSubjectRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetSubject(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareSubjectInfo(t, mod, info)

	return mod
}

func updateTestSubject(
	t *testing.T,
	req SubjectRequest) model.Subject {

	mod, err := testQueriesSchool.UpdateSubject(context.Background(), req)
	require.NoError(t, err)
	compareSubjectRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetSubject(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareSubjectInfo(t, mod, info)

	return mod
}

func testDeleteSubject(t *testing.T, id int64) {
	err := testQueriesSchool.DeleteSubject(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesSchool.GetSubject(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func randomSubject(subject string, uuid uuid.UUID) SubjectRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	subj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Subject", 0, subject)
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectType", 0, "Regular")

	arg := SubjectRequest{
		Uuid:         uuid,
		SubjectTitle: subject,
		SubjectRefId: subj.Id,
		TypeId:       util.SetNullInt64(typ.Id),
		Remarks:      util.RandomString(10),
		OtherInfo:    sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

// func testSubjectInfo(t *testing.T, sub string, uuid uuid.UUID) SubjectInfo {
// 	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
// 	info, _ := json.Marshal(otherInfo)
// 	subj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Subject", 0, sub)
// 	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectType", 0, "Regular")

// 	arg := SubjectRequest{
// 		Uuid:         uuid,
// 		SubjectTitle: sub,
// 		SubjectRefId: subj.Id,
// 		TypeId:       util.SetNullInt64(typ.Id),
// 		Remarks:      "Elementary" + sub,
// 		OtherInfo:    sql.NullString(sql.NullString{String: string(info), Valid: true}),
// 	}

// 	subjInfo, _ := testQueriesSchool.GetSubjectbyUuid(context.Background(), arg.Uuid)
// 	// require.NoError(t, err)

// 	if subjInfo.Id == 0 {
// 		syl := createTestSubject(t, arg)
// 		subjInfo, _ = testQueriesSchool.GetSubjectbyUuid(context.Background(), syl.Uuid)
// 	}

// 	return subjInfo
// }
