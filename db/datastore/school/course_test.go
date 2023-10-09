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

func TestCourse(t *testing.T) {

	// Test Data
	d1 := randomCourse()
	d2 := randomCourse()

	// Test Create
	CreatedD1 := createTestCourse(t, d1)
	CreatedD2 := createTestCourse(t, d2)

	infoUuid, err1 := testQueriesSchool.GetCoursebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err1)
	compareCourseInfo(t, CreatedD1, infoUuid)

	// Update Data
	updateD2 := d2
	updateD2.Id = CreatedD2.Id
	updateD2.CourseTitle = updateD2.CourseTitle + "-Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestCourse(t, updateD2)

	testListCourse(t, ListCourseParams{
		SchoolId: updatedD1.SchoolId,
		Limit:    5,
		Offset:   0,
	})

	// Delete Data
	testDeleteCourse(t, CreatedD1.Id)
	testDeleteCourse(t, CreatedD2.Id)
}

func testListCourse(t *testing.T, arg ListCourseParams) {

	course, err := testQueriesSchool.ListCourse(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", course)
	require.NotEmpty(t, course)
}

func compareCourseRequest(t *testing.T, mod model.Course, req CourseRequest) {
	require.NotEmpty(t, mod)
	require.Equal(t, req.SchoolId, mod.SchoolId)
	require.Equal(t, req.CourseTitle, mod.CourseTitle)
	require.Equal(t, req.CourseRefId, mod.CourseRefId)
	require.Equal(t, req.StatusId, mod.StatusId)
	require.Equal(t, req.Remarks, mod.Remarks)
	require.JSONEq(t, req.OtherInfo.String, mod.OtherInfo.String)
}

func compareCourseInfo(t *testing.T, mod model.Course, info CourseInfo) {
	require.NotEmpty(t, mod)
	require.Equal(t, info.SchoolId, mod.SchoolId)
	require.Equal(t, info.CourseTitle, mod.CourseTitle)
	require.Equal(t, info.CourseRefId, mod.CourseRefId)
	require.Equal(t, info.StatusId, mod.StatusId)
	require.Equal(t, info.Remarks, mod.Remarks)
	require.JSONEq(t, info.OtherInfo.String, mod.OtherInfo.String)
}
func createTestCourse(
	t *testing.T,
	req CourseRequest) model.Course {

	mod, err := testQueriesSchool.CreateCourse(context.Background(), req)
	require.NoError(t, err)
	compareCourseRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetCourse(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareCourseInfo(t, mod, info)

	return mod
}

func updateTestCourse(
	t *testing.T,
	req CourseRequest) model.Course {

	mod, err := testQueriesSchool.UpdateCourse(context.Background(), req)
	require.NoError(t, err)
	compareCourseRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetCourse(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareCourseInfo(t, mod, info)

	return mod
}

func testDeleteCourse(t *testing.T, id int64) {
	err := testQueriesSchool.DeleteCourse(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesSchool.GetCourse(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func randomCourse() CourseRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")

	corsP, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "CourseType", 0, "Pre-College")
	cors, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Courses", corsP.Id, "Elementary")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "CourseStatus", 0, "Current")

	arg := CourseRequest{
		SchoolId:    ofc.Id,
		CourseTitle: util.RandomString(10),
		CourseRefId: cors.Id,
		StatusId:    stat.Id,
		Remarks:     util.RandomString(10),
		OtherInfo:   sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	log.Printf("CourseRequest corsP: %v", corsP)
	log.Printf("CourseRequest: %v", arg)
	return arg
}

func testCourseData() CourseRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")

	corsP, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "CourseType", 0, "Pre-College")
	cors, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Courses", corsP.Id, "Elementary")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "CourseStatus", 0, "Current")

	arg := CourseRequest{
		Uuid:        util.ToUUID("8845d46a-c910-4467-a40b-37008c78a288"),
		SchoolId:    ofc.Id,
		CourseTitle: util.RandomString(10),
		CourseRefId: cors.Id,
		StatusId:    stat.Id,
		Remarks:     util.RandomString(10),
		OtherInfo:   sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}
