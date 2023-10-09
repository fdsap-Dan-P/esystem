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

func TestSchoolSection(t *testing.T) {

	// Test Data
	d1 := randomSchoolSection(t, util.ToUUID("0d4160d6-9ea8-4b55-8337-b2e2c63c13b3"))
	d2 := randomSchoolSection(t, util.ToUUID("f3fe830a-9b70-44f1-b42f-977998f289b8"))
	corsd1 := testCourseData()

	cors, _ := testQueriesSchool.GetCoursebyUuid(context.Background(), corsd1.Uuid)
	// Test Create
	if cors.Id == 0 {
		cors2 := createTestCourse(t, corsd1)
		d1.CourseId = cors2.Id
		d2.CourseId = cors2.Id
	} else {
		d1.CourseId = cors.Id
		d2.CourseId = cors.Id
	}

	// Test Create
	CreatedD1 := createTestSchoolSection(t, d1)
	CreatedD2 := createTestSchoolSection(t, d2)

	infoUuid, err1 := testQueriesSchool.GetSchoolSectionbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err1)
	compareSchoolSectionInfo(t, CreatedD1, infoUuid)

	// Update Data
	updateD2 := d2
	updateD2.Id = CreatedD2.Id
	updateD2.Remarks = updateD2.Remarks + "-Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestSchoolSection(t, updateD2)

	testListSchoolSection(t, ListSchoolSectionParams{
		SyllabusId: updatedD1.SyllabusId,
		Limit:      5,
		Offset:     0,
	})

	// Delete Data
	testDeleteSchoolSection(t, CreatedD1.Id)
	testDeleteSchoolSection(t, CreatedD2.Id)
}

func testListSchoolSection(t *testing.T, arg ListSchoolSectionParams) {

	schoolSection, err := testQueriesSchool.ListSchoolSection(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", schoolSection)
	require.NotEmpty(t, schoolSection)
}

func compareSchoolSectionRequest(t *testing.T, mod model.SchoolSection, req SchoolSectionRequest) {
	require.NotEmpty(t, mod)
	require.Equal(t, req.SyllabusId, mod.SyllabusId)
	require.Equal(t, req.SchoolId, mod.SchoolId)
	require.Equal(t, req.CourseId, mod.CourseId)
	require.Equal(t, req.StartDate.Time.Format("2006-01-02"), mod.StartDate.Time.Format("2006-01-02"))
	require.Equal(t, req.EndDate.Time.Format("2006-01-02"), mod.EndDate.Time.Format("2006-01-02"))
	require.Equal(t, req.AdviserId, mod.AdviserId)
	require.Equal(t, req.StatusId, mod.StatusId)
	require.Equal(t, req.SectionName, mod.SectionName)
	require.Equal(t, req.Remarks, mod.Remarks)
	require.JSONEq(t, req.OtherInfo.String, mod.OtherInfo.String)
}

func compareSchoolSectionInfo(t *testing.T, mod model.SchoolSection, info SchoolSectionInfo) {
	require.NotEmpty(t, mod)
	require.Equal(t, info.SyllabusId, mod.SyllabusId)
	require.Equal(t, info.SchoolId, mod.SchoolId)
	require.Equal(t, info.CourseId, mod.CourseId)
	require.Equal(t, info.StartDate.Time.Format("2006-01-02"), mod.StartDate.Time.Format("2006-01-02"))
	require.Equal(t, info.EndDate.Time.Format("2006-01-02"), mod.EndDate.Time.Format("2006-01-02"))
	require.Equal(t, info.AdviserId, mod.AdviserId)
	require.Equal(t, info.StatusId, mod.StatusId)
	require.Equal(t, info.SectionName, mod.SectionName)
	require.Equal(t, info.Remarks, mod.Remarks)
	require.JSONEq(t, info.OtherInfo.String, mod.OtherInfo.String)
}
func createTestSchoolSection(
	t *testing.T,
	req SchoolSectionRequest) model.SchoolSection {

	mod, err := testQueriesSchool.CreateSchoolSection(context.Background(), req)
	require.NoError(t, err)
	compareSchoolSectionRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetSchoolSection(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareSchoolSectionInfo(t, mod, info)

	return mod
}

func updateTestSchoolSection(
	t *testing.T,
	req SchoolSectionRequest) model.SchoolSection {

	mod, err := testQueriesSchool.UpdateSchoolSection(context.Background(), req)
	require.NoError(t, err)
	compareSchoolSectionRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetSchoolSection(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareSchoolSectionInfo(t, mod, info)

	return mod
}

func testDeleteSchoolSection(t *testing.T, id int64) {
	err := testQueriesSchool.DeleteSchoolSection(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesSchool.GetSchoolSection(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func randomSchoolSection(t *testing.T, uuid uuid.UUID) SchoolSectionRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	//	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectType", 0, "Regular")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectStatus", 0, "Current")
	// adv, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SchoolSemister", 0, "Full Year")

	syl := testSyllabusData(t)
	sylInfo, _ := testQueriesSchool.GetSyllabusbyUuid(context.Background(), syl.Uuid)
	// Test Create
	if sylInfo.Id == 0 {
		syl := createTestSyllabus(t, syl)
		sylInfo, _ = testQueriesSchool.GetSyllabusbyUuid(context.Background(), syl.Uuid)
	}
	fmt.Printf("randomSchoolSection: %+v\n", sylInfo)

	arg := SchoolSectionRequest{
		Uuid:       uuid,
		SyllabusId: sylInfo.Id,
		SchoolId:   sylInfo.SchoolId,
		CourseId:   sylInfo.CourseId,
		StartDate:  util.RandomNullDate(),
		EndDate:    util.RandomNullDate(),
		// AdviserId:   util.SetNullInt64(adv.Id),
		StatusId:    stat.Id,
		SectionName: util.RandomNullString(10),
		Remarks:     util.RandomString(10),
		OtherInfo:   sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}

	return arg
}

func testSchoolSectionData(t *testing.T) SchoolSectionRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	//	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectType", 0, "Regular")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectStatus", 0, "Current")
	// adv, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SchoolSemister", 0, "Full Year")

	syl := testSyllabusData(t)
	sylInfo, _ := testQueriesSchool.GetSyllabusbyUuid(context.Background(), syl.Uuid)
	// Test Create
	if sylInfo.Id == 0 {
		syl := createTestSyllabus(t, syl)
		sylInfo, _ = testQueriesSchool.GetSyllabusbyUuid(context.Background(), syl.Uuid)
	}
	fmt.Printf("randomSchoolSection: %+v\n", sylInfo)

	arg := SchoolSectionRequest{
		Uuid:       util.ToUUID("3bdd00c6-dffb-4b7b-9b02-bf358205a690"),
		SyllabusId: sylInfo.Id,
		SchoolId:   sylInfo.SchoolId,
		CourseId:   sylInfo.CourseId,
		StartDate:  util.RandomNullDate(),
		EndDate:    util.RandomNullDate(),
		// AdviserId:   util.SetNullInt64(adv.Id),
		StatusId:    stat.Id,
		SectionName: util.RandomNullString(10),
		Remarks:     util.RandomString(10),
		OtherInfo:   sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}

	// corsd1 := testCourseData()
	// cors, _ := testQueriesSchool.GetCoursebyUuid(context.Background(), corsd1.Uuid)
	// // Test Create
	// if cors.Id == 0 {
	// 	cors2 := createTestCourse(t, corsd1)
	// 	arg.CourseId = cors2.Id
	// } else {
	// 	arg.CourseId = cors.Id
	// }

	return arg
}
