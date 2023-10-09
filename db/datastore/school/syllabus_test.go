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

	"github.com/stretchr/testify/require"
)

func TestSyllabus(t *testing.T) {

	// Test Data
	d1 := randomSyllabus(t)
	d2 := randomSyllabus(t)
	d2.Uuid = util.ToUUID("de56bc93-bbba-4915-8efb-091e4cec0f37")

	// Test Create
	CreatedD1 := createTestSyllabus(t, d1)
	CreatedD2 := createTestSyllabus(t, d2)

	infoUuid, err1 := testQueriesSchool.GetSyllabusbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err1)
	compareSyllabusInfo(t, CreatedD1, infoUuid)

	// Update Data
	updateD2 := d2
	updateD2.Id = CreatedD2.Id
	updateD2.Remarks = updateD2.Remarks + "-Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestSyllabus(t, updateD2)

	testListSyllabus(t, ListSyllabusParams{
		CourseId: updatedD1.CourseId,
		Limit:    5,
		Offset:   0,
	})

	// Delete Data
	testDeleteSyllabus(t, CreatedD1.Id)
	testDeleteSyllabus(t, CreatedD2.Id)
}

func testListSyllabus(t *testing.T, arg ListSyllabusParams) {

	syllabus, err := testQueriesSchool.ListSyllabus(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", syllabus)
	require.NotEmpty(t, syllabus)
}

func compareSyllabusRequest(t *testing.T, mod model.Syllabus, req SyllabusRequest) {
	require.NotEmpty(t, mod)
	require.Equal(t, req.CourseId, mod.CourseId)
	require.Equal(t, req.Version, mod.Version)
	require.Equal(t, req.CourseYear, mod.CourseYear)
	require.Equal(t, req.SemisterId, mod.SemisterId)
	require.Equal(t, req.StatusId, mod.StatusId)
	require.Equal(t, req.DateImplement.Format("2006-01-02"), mod.DateImplement.Format("2006-01-02"))
	require.Equal(t, req.Remarks, mod.Remarks)
	require.JSONEq(t, req.OtherInfo.String, mod.OtherInfo.String)
}

func compareSyllabusInfo(t *testing.T, mod model.Syllabus, info SyllabusInfo) {
	require.NotEmpty(t, mod)
	require.Equal(t, info.CourseId, mod.CourseId)
	require.Equal(t, info.Version, mod.Version)
	require.Equal(t, info.CourseYear, mod.CourseYear)
	require.Equal(t, info.SemisterId, mod.SemisterId)
	require.Equal(t, info.StatusId, mod.StatusId)
	require.Equal(t, info.DateImplement.Format("2006-01-02"), mod.DateImplement.Format("2006-01-02"))
	require.Equal(t, info.Remarks, mod.Remarks)
	require.JSONEq(t, info.OtherInfo.String, mod.OtherInfo.String)
}
func createTestSyllabus(
	t *testing.T,
	req SyllabusRequest) model.Syllabus {

	mod, err := testQueriesSchool.CreateSyllabus(context.Background(), req)
	require.NoError(t, err)
	compareSyllabusRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetSyllabus(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareSyllabusInfo(t, mod, info)

	return mod
}

func updateTestSyllabus(
	t *testing.T,
	req SyllabusRequest) model.Syllabus {

	mod, err := testQueriesSchool.UpdateSyllabus(context.Background(), req)
	require.NoError(t, err)
	compareSyllabusRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetSyllabus(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareSyllabusInfo(t, mod, info)

	return mod
}

func testDeleteSyllabus(t *testing.T, id int64) {
	err := testQueriesSchool.DeleteSyllabus(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesSchool.GetSyllabus(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func randomSyllabus(t *testing.T) SyllabusRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	//	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectType", 0, "Regular")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectStatus", 0, "Current")
	sem, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SchoolSemister", 0, "Full Year")

	arg := SyllabusRequest{
		Uuid:          util.ToUUID("94b29ea1-430f-4194-ae6f-fbc6ec836278"),
		Version:       util.RandomString(10),
		CourseYear:    util.RandomInt32(1990, 2090),
		SemisterId:    sem.Id,
		StatusId:      stat.Id,
		DateImplement: util.RandomDate(),
		Remarks:       util.RandomString(10),
		OtherInfo:     sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}

	corsd1 := testCourseData()
	cors, _ := testQueriesSchool.GetCoursebyUuid(context.Background(), corsd1.Uuid)
	// Test Create
	if cors.Id == 0 {
		cors2 := createTestCourse(t, corsd1)
		arg.CourseId = cors2.Id
	} else {
		arg.CourseId = cors.Id
	}
	log.Printf("arg %v", arg)

	return arg
}

func testSyllabusData(t *testing.T) SyllabusRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	//	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectType", 0, "Regular")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectStatus", 0, "Current")
	sem, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SchoolSemister", 0, "Full Year")

	arg := SyllabusRequest{
		Uuid:          util.ToUUID("034ae2f0-ff93-405d-a7cf-0983750a6ed9"),
		Version:       util.RandomString(10),
		CourseYear:    util.RandomInt32(1990, 2090),
		SemisterId:    sem.Id,
		StatusId:      stat.Id,
		DateImplement: util.RandomDate(),
		Remarks:       util.RandomString(10),
		OtherInfo:     sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}

	corsd1 := testCourseData()
	cors, _ := testQueriesSchool.GetCoursebyUuid(context.Background(), corsd1.Uuid)
	// Test Create
	if cors.Id == 0 {
		cors2 := createTestCourse(t, corsd1)
		arg.CourseId = cors2.Id
	} else {
		arg.CourseId = cors.Id
	}

	return arg
}
