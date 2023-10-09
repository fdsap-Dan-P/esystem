package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	ref "simplebank/db/datastore/reference"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestEmployment(t *testing.T) {

	// Test Data
	d1 := randomEmployment()
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")
	d1.Iiid = ii.Id

	d2 := randomEmployment()
	ii, _ = testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "101")
	d2.Iiid = ii.Id

	// Test Create
	CreatedD1 := createTestEmployment(t, d1)
	CreatedD2 := createTestEmployment(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetEmployment(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Company, getData1.Company)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.AddressDetail, getData1.AddressDetail)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.Equal(t, d1.StartDate.Time.Format("2006-01-02"), getData1.StartDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.EndDate.Time.Format("2006-01-02"), getData1.EndDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.PeriodDate, getData1.PeriodDate)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesIdentity.GetEmployment(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.Company, getData2.Company)
	require.Equal(t, d2.Title, getData2.Title)
	require.Equal(t, d2.AddressDetail, getData2.AddressDetail)
	require.Equal(t, d2.AddressUrl, getData2.AddressUrl)
	require.Equal(t, d2.GeographyId, getData2.GeographyId)
	require.Equal(t, d2.StartDate.Time.Format("2006-01-02"), getData2.StartDate.Time.Format("2006-01-02"))
	require.Equal(t, d2.EndDate.Time.Format("2006-01-02"), getData2.EndDate.Time.Format("2006-01-02"))
	require.Equal(t, d2.PeriodDate, getData2.PeriodDate)
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesIdentity.GetEmploymentbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// getData, err = testQueriesIdentity.GetEmploymentbyName(context.Background(), CreatedD1.Name)
	// require.NoError(t, err)
	// require.NotEmpty(t, getData)
	// require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	// fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	updateD2.Company = updateD2.Company + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestEmployment(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.Company, updatedD1.Company)
	require.Equal(t, updateD2.Title, updatedD1.Title)
	require.Equal(t, updateD2.AddressDetail, updatedD1.AddressDetail)
	require.Equal(t, updateD2.AddressUrl, updatedD1.AddressUrl)
	require.Equal(t, updateD2.GeographyId, updatedD1.GeographyId)
	require.Equal(t, updateD2.StartDate.Time.Format("2006-01-02"), updatedD1.StartDate.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.EndDate.Time.Format("2006-01-02"), updatedD1.EndDate.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.PeriodDate, updatedD1.PeriodDate)
	require.Equal(t, updateD2.Remarks, updatedD1.Remarks)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListEmployment(t, ListEmploymentParams{
		Iiid:   updatedD1.Iiid,
		Limit:  5,
		Offset: 0,
	})
	// Delete Data
	testDeleteEmployment(t, CreatedD1.Uuid)
	testDeleteEmployment(t, CreatedD2.Uuid)
}

func testListEmployment(t *testing.T, arg ListEmploymentParams) {

	employment, err := testQueriesIdentity.ListEmployment(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", employment)
	require.NotEmpty(t, employment)

}

func randomEmployment() EmploymentRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	geo, _ := testQueriesReference.SearchGeography(context.Background(),
		ref.SearchGeographyParams{
			SearchText: "West Triangle, Quezon City. Metro Manila",
			Limit:      1,
			Offset:     0,
		})

	arg := EmploymentRequest{
		// Iiid:          util.RandomInt(1, 100),
		Series:        int16(util.RandomInt32(1, 100)),
		Company:       util.RandomString(10),
		Title:         util.RandomString(10),
		AddressDetail: sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		AddressUrl:    sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		GeographyId:   sql.NullInt64(sql.NullInt64{Int64: geo[0].Id, Valid: true}),
		StartDate:     sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		EndDate:       sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		PeriodDate:    sql.NullString(sql.NullString{String: "2010-2021", Valid: true}),
		Remarks:       sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestEmployment(
	t *testing.T,
	d1 EmploymentRequest) model.Employment {

	getData1, err := testQueriesIdentity.CreateEmployment(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Company, getData1.Company)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.AddressDetail, getData1.AddressDetail)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.Equal(t, d1.StartDate.Time.Format("2006-01-02"), getData1.StartDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.EndDate.Time.Format("2006-01-02"), getData1.EndDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.PeriodDate, getData1.PeriodDate)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestEmployment(
	t *testing.T,
	d1 EmploymentRequest) model.Employment {

	getData1, err := testQueriesIdentity.UpdateEmployment(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Company, getData1.Company)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.AddressDetail, getData1.AddressDetail)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.Equal(t, d1.StartDate.Time.Format("2006-01-02"), getData1.StartDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.EndDate.Time.Format("2006-01-02"), getData1.EndDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.PeriodDate, getData1.PeriodDate)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteEmployment(t *testing.T, uuid uuid.UUID) {
	err := testQueriesIdentity.DeleteEmployment(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetEmployment(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
