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

func TestSectionSubject(t *testing.T) {

	// Test Data
	d1 := randomSectionSubject(t, "English", util.ToUUID("ed6a63f9-44bd-4aef-87ed-8e5538f392f8"))
	d2 := randomSectionSubject(t, "Mathematics", util.ToUUID("610336d8-51ff-4e1b-82e7-dc03958e75c8"))
	// subj2 := testSubjectInfo(t, "Filipino", util.ToUUID("e92ad175-eb89-449b-b3a1-7592c099df39"))
	// d2.SubjectId = subj2.Id

	// Test Create
	CreatedD1 := createTestSectionSubject(t, d1)
	CreatedD2 := createTestSectionSubject(t, d2)

	infoUuid, err1 := testQueriesSchool.GetSectionSubjectbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err1)

	log.Printf("GetSectionSubjectbyUuid CreatedD1 %v", CreatedD1)
	log.Printf("GetSectionSubjectbyUuid infoUuid %v", infoUuid)

	compareSectionSubjectInfo(t, CreatedD1, infoUuid)

	// Update Data
	updateD2 := d2
	updateD2.Id = CreatedD2.Id
	updateD2.Remarks = updateD2.Remarks + "-Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestSectionSubject(t, updateD2)

	testListSectionSubject(t, ListSectionSubjectParams{
		SchoolSectionId: updatedD1.SchoolSectionId,
		Limit:           5,
		Offset:          0,
	})

	// testSectionSubjectInfo(t)

	// Delete Data
	testDeleteSectionSubject(t, CreatedD1.Uuid)
	testDeleteSectionSubject(t, CreatedD2.Uuid)
}

func testListSectionSubject(t *testing.T, arg ListSectionSubjectParams) {

	SectionSubject, err := testQueriesSchool.ListSectionSubject(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", SectionSubject)
	require.NotEmpty(t, SectionSubject)
}

func compareSectionSubjectRequest(t *testing.T, mod model.SectionSubject, req SectionSubjectRequest) {
	require.NotEmpty(t, mod)
	require.Equal(t, req.Uuid, mod.Uuid)
	require.Equal(t, req.SchoolSectionId, mod.SchoolSectionId)
	require.Equal(t, req.SubjectId, mod.SubjectId)
	require.Equal(t, req.TeacherId, mod.TeacherId)
	require.Equal(t, req.TypeId, mod.TypeId)
	require.Equal(t, req.StatusId, mod.StatusId)
	require.Equal(t, req.ScheduleCode, mod.ScheduleCode)
	require.JSONEq(t, req.ScheduleJson.String, mod.ScheduleJson.String)
	require.Equal(t, req.Remarks, mod.Remarks)
	require.JSONEq(t, req.OtherInfo.String, mod.OtherInfo.String)
}

func compareSectionSubjectInfo(t *testing.T, mod model.SectionSubject, info SectionSubjectInfo) {
	require.NotEmpty(t, mod)
	require.Equal(t, info.Uuid, mod.Uuid)
	require.Equal(t, info.SchoolSectionId, mod.SchoolSectionId)
	require.Equal(t, info.SubjectId, mod.SubjectId)
	require.Equal(t, info.TeacherId, mod.TeacherId)
	require.Equal(t, info.TypeId, mod.TypeId)
	require.Equal(t, info.StatusId, mod.StatusId)
	require.Equal(t, info.ScheduleCode, mod.ScheduleCode)
	require.JSONEq(t, info.ScheduleJson.String, mod.ScheduleJson.String)
	require.Equal(t, info.Remarks, mod.Remarks)
	require.JSONEq(t, info.OtherInfo.String, mod.OtherInfo.String)
}
func createTestSectionSubject(
	t *testing.T,
	req SectionSubjectRequest) model.SectionSubject {

	mod, err := testQueriesSchool.CreateSectionSubject(context.Background(), req)
	require.NoError(t, err)
	compareSectionSubjectRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetSectionSubject(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareSectionSubjectInfo(t, mod, info)

	return mod
}

func updateTestSectionSubject(
	t *testing.T,
	req SectionSubjectRequest) model.SectionSubject {

	mod, err := testQueriesSchool.UpdateSectionSubject(context.Background(), req)
	require.NoError(t, err)
	compareSectionSubjectRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetSectionSubject(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareSectionSubjectInfo(t, mod, info)

	return mod
}

func testDeleteSectionSubject(t *testing.T, uuid uuid.UUID) {
	err := testQueriesSchool.DeleteSectionSubject(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesSchool.GetSectionSubjectbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func randomSectionSubject(t *testing.T, subject string, uuid uuid.UUID) SectionSubjectRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectType", 0, "Regular")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectStatus", 0, "Current")
	// adv, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SchoolSemister", 0, "Full Year")

	sec := testSchoolSectionData(t)
	secInfo, _ := testQueriesSchool.GetSchoolSectionbyUuid(context.Background(), sec.Uuid)
	// Test Create
	if secInfo.Id == 0 {
		syl := createTestSchoolSection(t, sec)
		secInfo, _ = testQueriesSchool.GetSchoolSectionbyUuid(context.Background(), syl.Uuid)
	}

	subj, _ := testQueriesSchool.CreateSubject(context.Background(), randomSubject(subject, uuid))

	emp, _ := testQueriesIdentity.GetEmployeebyEmpNo(context.Background(), 2, "97-0114")

	arg := SectionSubjectRequest{
		Uuid:            uuid,
		SubjectId:       subj.Id,
		SchoolSectionId: secInfo.Id,
		TeacherId:       util.SetNullInt64(emp.Id),
		TypeId:          typ.Id,
		StatusId:        stat.Id,
		ScheduleCode:    util.RandomNullString(10),
		ScheduleJson:    sql.NullString(sql.NullString{String: string(info), Valid: true}),
		Remarks:         util.RandomString(10),
		OtherInfo:       sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}

	return arg
}
