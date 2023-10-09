package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	ref "simplebank/db/datastore/reference"
	"simplebank/model"
	"simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestIdentityInfo(t *testing.T) {

	// Test Data
	d1 := randomIdentityInfo(t)
	d2 := randomIdentityInfo(t)

	// Test Create
	CreatedD1 := createTestIdentityInfo(t, d1)
	CreatedD2 := createTestIdentityInfo(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetIdentityInfo(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Isperson, getData1.Isperson)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.LastName, getData1.LastName)
	require.Equal(t, d1.FirstName, getData1.FirstName)
	require.Equal(t, d1.MiddleName, getData1.MiddleName)
	require.Equal(t, d1.MotherMaidenName, getData1.MotherMaidenName)
	require.Equal(t, d1.Birthday.Time.Format("2006-01-02"), getData1.Birthday.Time.Format("2006-01-02"))
	require.Equal(t, d1.Sex, getData1.Sex)
	require.Equal(t, d1.GenderId, getData1.GenderId)
	require.Equal(t, d1.CivilStatusId, getData1.CivilStatusId)
	require.Equal(t, d1.BirthPlaceId, getData1.BirthPlaceId)
	require.Equal(t, d1.ContactId, getData1.ContactId)
	require.Equal(t, d1.IdentityMapId, getData1.IdentityMapId)
	require.Equal(t, d1.AlternateId, getData1.AlternateId)
	require.Equal(t, d1.Phone, getData1.Phone)
	require.Equal(t, d1.Email, getData1.Email)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesIdentity.GetIdentityInfo(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Isperson, getData2.Isperson)
	require.Equal(t, d2.Title, getData2.Title)
	require.Equal(t, d2.LastName, getData2.LastName)
	require.Equal(t, d2.FirstName, getData2.FirstName)
	require.Equal(t, d2.MiddleName, getData2.MiddleName)
	require.Equal(t, d2.MotherMaidenName, getData2.MotherMaidenName)
	require.Equal(t, d2.Birthday.Time.Format("2006-01-02"), getData2.Birthday.Time.Format("2006-01-02"))
	require.Equal(t, d2.Sex, getData2.Sex)
	require.Equal(t, d2.GenderId, getData2.GenderId)
	require.Equal(t, d2.CivilStatusId, getData2.CivilStatusId)
	require.Equal(t, d2.BirthPlaceId, getData2.BirthPlaceId)
	require.Equal(t, d2.ContactId, getData2.ContactId)
	require.Equal(t, d2.IdentityMapId, getData2.IdentityMapId)
	require.Equal(t, d2.AlternateId, getData2.AlternateId)
	require.Equal(t, d2.Phone, getData2.Phone)
	require.Equal(t, d2.Email, getData2.Email)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesIdentity.GetIdentityInfobyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)

	fmt.Printf("%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.LastName = updateD2.LastName + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestIdentityInfo(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Isperson, updatedD1.Isperson)
	require.Equal(t, updateD2.Title, updatedD1.Title)
	require.Equal(t, updateD2.LastName, updatedD1.LastName)
	require.Equal(t, updateD2.FirstName, updatedD1.FirstName)
	require.Equal(t, updateD2.MiddleName, updatedD1.MiddleName)
	require.Equal(t, updateD2.MotherMaidenName, updatedD1.MotherMaidenName)
	require.Equal(t, updateD2.Birthday.Time.Format("2006-01-02"), updatedD1.Birthday.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.Sex, updatedD1.Sex)
	require.Equal(t, updateD2.GenderId, updatedD1.GenderId)
	require.Equal(t, updateD2.CivilStatusId, updatedD1.CivilStatusId)
	require.Equal(t, updateD2.BirthPlaceId, updatedD1.BirthPlaceId)
	require.Equal(t, updateD2.ContactId, updatedD1.ContactId)
	require.Equal(t, updateD2.IdentityMapId, updatedD1.IdentityMapId)
	require.Equal(t, updateD2.AlternateId, updatedD1.AlternateId)
	require.Equal(t, updateD2.Phone, updatedD1.Phone)
	require.Equal(t, updateD2.Email, updatedD1.Email)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteIdentityInfo(t, getData1.Id)
	testDeleteIdentityInfo(t, getData2.Id)
}

func TestListIdentityInfo(t *testing.T) {

	arg := ListIdentityInfoParams{
		Limit:  5,
		Offset: 0,
	}

	identityInfo, err := testQueriesIdentity.ListIdentityInfo(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", identityInfo)
	require.NotEmpty(t, identityInfo)

}

func randomIdentityInfo(t *testing.T) IdentityInfoRequest {
	otherInfo := &TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	birthPlace, err := testQueriesReference.SearchGeography(context.Background(),
		ref.SearchGeographyParams{
			SearchText: "West Triangle, Quezon City. Metro Manila",
			Limit:      1,
			Offset:     0,
		})

	require.NoError(t, err)
	fmt.Printf("%+v\n", birthPlace)
	require.NotEmpty(t, birthPlace)

	gender, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Gender", 0, "Male")
	civil, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "CivilStatus", 0, "Married")
	// contact, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "Gender", 0, "Male")
	// identityMap, _ := testQueriesIdentity.GetIdentityInfobyAltAcc(context.Background(), "Gender")

	arg := IdentityInfoRequest{
		Isperson:         true,
		Title:            util.SetNullString("Ms."),
		LastName:         util.RandomString(10),
		FirstName:        sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		MiddleName:       sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		MotherMaidenName: sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		Birthday:         sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		Sex:              sql.NullBool(sql.NullBool{Bool: true, Valid: true}),
		GenderId:         gender.ParentId,
		CivilStatusId:    sql.NullInt64(sql.NullInt64{Int64: civil.Id, Valid: true}),
		BirthPlaceId:     sql.NullInt64(sql.NullInt64{Int64: birthPlace[0].Id, Valid: true}),
		// ContactId:        sql.NullInt64(sql.NullInt64{Int64: contact.Id, Valid: true}),
		IdentityMapId: sql.NullInt64(sql.NullInt64{Int64: 0, Valid: false}),
		AlternateId:   sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		Phone:         sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		Email:         sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestIdentityInfo(
	t *testing.T,
	d1 IdentityInfoRequest) model.IdentityInfo {

	getData1, err := testQueriesIdentity.CreateIdentityInfo(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Isperson, getData1.Isperson)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.LastName, getData1.LastName)
	require.Equal(t, d1.FirstName, getData1.FirstName)
	require.Equal(t, d1.MiddleName, getData1.MiddleName)
	require.Equal(t, d1.MotherMaidenName, getData1.MotherMaidenName)
	require.Equal(t, d1.Birthday.Time.Format("2006-01-02"), getData1.Birthday.Time.Format("2006-01-02"))
	require.Equal(t, d1.Sex, getData1.Sex)
	require.Equal(t, d1.GenderId, getData1.GenderId)
	require.Equal(t, d1.CivilStatusId, getData1.CivilStatusId)
	require.Equal(t, d1.BirthPlaceId, getData1.BirthPlaceId)
	require.Equal(t, d1.ContactId, getData1.ContactId)
	require.Equal(t, d1.IdentityMapId, getData1.IdentityMapId)
	require.Equal(t, d1.AlternateId, getData1.AlternateId)
	require.Equal(t, d1.Phone, getData1.Phone)
	require.Equal(t, d1.Email, getData1.Email)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestIdentityInfo(
	t *testing.T,
	d1 IdentityInfoRequest) model.IdentityInfo {

	getData1, err := testQueriesIdentity.UpdateIdentityInfo(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Isperson, getData1.Isperson)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.LastName, getData1.LastName)
	require.Equal(t, d1.FirstName, getData1.FirstName)
	require.Equal(t, d1.MiddleName, getData1.MiddleName)
	require.Equal(t, d1.MotherMaidenName, getData1.MotherMaidenName)
	require.Equal(t, d1.Birthday.Time.Format("2006-01-02"), getData1.Birthday.Time.Format("2006-01-02"))
	require.Equal(t, d1.Sex, getData1.Sex)
	require.Equal(t, d1.GenderId, getData1.GenderId)
	require.Equal(t, d1.CivilStatusId, getData1.CivilStatusId)
	require.Equal(t, d1.BirthPlaceId, getData1.BirthPlaceId)
	require.Equal(t, d1.ContactId, getData1.ContactId)
	require.Equal(t, d1.IdentityMapId, getData1.IdentityMapId)
	require.Equal(t, d1.AlternateId, getData1.AlternateId)
	require.Equal(t, d1.Phone, getData1.Phone)
	require.Equal(t, d1.Email, getData1.Email)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteIdentityInfo(t *testing.T, id int64) {
	err := testQueriesIdentity.DeleteIdentityInfo(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetIdentityInfo(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
