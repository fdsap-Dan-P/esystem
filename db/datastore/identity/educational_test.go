package db

import (
	"context"
	"database/sql"

	"encoding/json"
	"fmt"
	common "simplebank/db/common"
	ref "simplebank/db/datastore/reference"
	"simplebank/model"
	"simplebank/util"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestEducational(t *testing.T) {

	// Test Data
	d1 := randomEducational()
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")
	d1.Iiid = ii.Id
	lvl, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "educationallevel", 0, "Grade 6")
	typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "coursetype", 0, "Vocational")
	// crs, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "courses", 0, "Formal Sciences")
	d1.LevelId.Int64 = lvl.Id
	d1.LevelId.Valid = true
	d1.CourseTypeId.Int64 = typ.Id
	d1.CourseTypeId.Valid = true
	// d1.CourseId.Int64 = crs.Id
	d1.CourseId.Valid = false

	fmt.Printf("Get by UUId%+v\n", d1.CourseId)

	d2 := randomEducational()
	ii, _ = testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "1001")
	d2.Iiid = ii.Id
	lvl, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "educationallevel", 0, "Vocational Graduate")
	typ, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "coursetype", 0, "Vocational")
	// crs, _ = testQueriesReference.GetReferenceInfobyTitle(context.Background(), "courses", 0, "Vocational")
	d2.LevelId.Int64 = lvl.Id
	d2.LevelId.Valid = true
	d2.CourseTypeId.Int64 = typ.Id
	d2.CourseTypeId.Valid = true
	// d2.CourseId.Int64 = crs.Id
	d2.CourseId.Valid = false

	// Test Create
	CreatedD1 := createTestEducational(t, d1)
	CreatedD2 := createTestEducational(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetEducational(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.LevelId, getData1.LevelId)
	require.Equal(t, d1.CourseTypeId, getData1.CourseTypeId)
	require.Equal(t, d1.CourseId, getData1.CourseId)
	require.Equal(t, d1.Course, getData1.Course)
	require.Equal(t, d1.School, getData1.School)
	require.Equal(t, d1.AddressDetail, getData1.AddressDetail)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.Equal(t, d1.StartDate.Time.Format("2006-01-02"), getData1.StartDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.EndDate.Time.Format("2006-01-02"), getData1.EndDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.PeriodDate, getData1.PeriodDate)
	require.Equal(t, d1.Completed, getData1.Completed)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesIdentity.GetEducational(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.LevelId, getData2.LevelId)
	require.Equal(t, d2.CourseTypeId, getData2.CourseTypeId)
	require.Equal(t, d2.CourseId, getData2.CourseId)
	require.Equal(t, d2.Course, getData2.Course)
	require.Equal(t, d2.School, getData2.School)
	require.Equal(t, d2.AddressDetail, getData2.AddressDetail)
	require.Equal(t, d2.AddressUrl, getData2.AddressUrl)
	require.Equal(t, d2.GeographyId, getData2.GeographyId)
	require.Equal(t, d2.StartDate.Time.Format("2006-01-02"), getData2.StartDate.Time.Format("2006-01-02"))
	require.Equal(t, d2.EndDate.Time.Format("2006-01-02"), getData2.EndDate.Time.Format("2006-01-02"))
	require.Equal(t, d2.PeriodDate, getData2.PeriodDate)
	require.Equal(t, d2.Completed, getData2.Completed)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesIdentity.GetEducationalbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	updateD2.Course = updateD2.Course + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestEducational(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.LevelId, updatedD1.LevelId)
	require.Equal(t, updateD2.CourseTypeId, updatedD1.CourseTypeId)
	require.Equal(t, updateD2.CourseId, updatedD1.CourseId)
	require.Equal(t, updateD2.Course, updatedD1.Course)
	require.Equal(t, updateD2.School, updatedD1.School)
	require.Equal(t, updateD2.AddressDetail, updatedD1.AddressDetail)
	require.Equal(t, updateD2.AddressUrl, updatedD1.AddressUrl)
	require.Equal(t, updateD2.GeographyId, updatedD1.GeographyId)
	require.Equal(t, updateD2.StartDate.Time.Format("2006-01-02"), updatedD1.StartDate.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.EndDate.Time.Format("2006-01-02"), updatedD1.EndDate.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.PeriodDate, updatedD1.PeriodDate)
	require.Equal(t, updateD2.Completed, updatedD1.Completed)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListEducational(t, ListEducationalParams{
		Iiid:   updatedD1.Iiid,
		Limit:  5,
		Offset: 0,
	})
	// Delete Data
	testDeleteEducational(t, CreatedD1.Uuid)
	testDeleteEducational(t, CreatedD2.Uuid)
}

func testListEducational(t *testing.T, arg ListEducationalParams) {

	educational, err := testQueriesIdentity.ListEducational(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", educational)
	require.NotEmpty(t, educational)

}

func randomEducational() EducationalRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	geo, _ := testQueriesReference.SearchGeography(context.Background(),
		ref.SearchGeographyParams{
			SearchText: "Soledad San Pablo City, Laguna",
			Limit:      1,
			Offset:     0,
		})

	arg := EducationalRequest{
		// Iiid:          util.RandomInt(1, 100),
		Series: int16(util.RandomInt32(1, 100)),
		// LevelId:       sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		// CourseTypeId:  sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		// CourseId:      sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		Course:        util.RandomString(10),
		School:        util.RandomString(10),
		AddressDetail: sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		AddressUrl:    sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		GeographyId:   sql.NullInt64(sql.NullInt64{Int64: geo[0].Id, Valid: true}),
		StartDate:     sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		EndDate:       sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		PeriodDate:    sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		Completed:     sql.NullBool(sql.NullBool{Bool: true, Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestEducational(
	t *testing.T,
	d1 EducationalRequest) model.Educational {

	getData1, err := testQueriesIdentity.CreateEducational(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.LevelId, getData1.LevelId)
	require.Equal(t, d1.CourseTypeId, getData1.CourseTypeId)
	require.Equal(t, d1.CourseId, getData1.CourseId)
	require.Equal(t, d1.Course, getData1.Course)
	require.Equal(t, d1.School, getData1.School)
	require.Equal(t, d1.AddressDetail, getData1.AddressDetail)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.Equal(t, d1.StartDate.Time.Format("2006-01-02"), getData1.StartDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.EndDate.Time.Format("2006-01-02"), getData1.EndDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.PeriodDate, getData1.PeriodDate)
	require.Equal(t, d1.Completed, getData1.Completed)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestEducational(
	t *testing.T,
	d1 EducationalRequest) model.Educational {

	getData1, err := testQueriesIdentity.UpdateEducational(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.LevelId, getData1.LevelId)
	require.Equal(t, d1.CourseTypeId, getData1.CourseTypeId)
	require.Equal(t, d1.CourseId, getData1.CourseId)
	require.Equal(t, d1.Course, getData1.Course)
	require.Equal(t, d1.School, getData1.School)
	require.Equal(t, d1.AddressDetail, getData1.AddressDetail)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.Equal(t, d1.StartDate.Time.Format("2006-01-02"), getData1.StartDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.EndDate.Time.Format("2006-01-02"), getData1.EndDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.PeriodDate, getData1.PeriodDate)
	require.Equal(t, d1.Completed, getData1.Completed)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteEducational(t *testing.T, uuid uuid.UUID) {
	err := testQueriesIdentity.DeleteEducational(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetEducational(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
