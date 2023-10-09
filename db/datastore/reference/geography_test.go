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

	"github.com/stretchr/testify/require"
)

func TestGeography(t *testing.T) {

	// Test Data
	d1 := randomGeography()
	d2 := randomGeography()

	// Test Create
	CreatedD1 := createTestGeography(t, d1)

	d, e := testQueriesReference.GetGeographybyLocation(context.Background(), d1.TypeId, d1.ParentId.Int64, d1.Location)
	require.NoError(t, e)
	require.NotEmpty(t, d)
	require.Equal(t, d1.Code, d.Code)
	require.Equal(t, d1.ShortName, d.ShortName)
	require.Equal(t, d1.Location, d.Location)
	require.Equal(t, d1.TypeId, d.TypeId)
	require.Equal(t, d1.ParentId, d.ParentId)
	require.Equal(t, d1.ZipCode, d.ZipCode)
	require.Equal(t, d1.Latitude, d.Latitude)
	require.Equal(t, d1.Longitude, d.Longitude)
	require.Equal(t, d1.AddressUrl, d.AddressUrl)

	// log.Println(e)
	// log.Println("Random:", d1.TypeId, d2.TypeId)
	// log.Println(d1.ParentId.Int64, d2.ParentId.Int64)
	// log.Println(d1.Location, d2.Location)
	// fmt.Printf("check if exitst: %+v\n", d)

	CreatedD2 := createTestGeography(t, d2)

	// Get Data
	getData1, err1 := testQueriesReference.GetGeography(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.ShortName, getData1.ShortName)
	require.Equal(t, d1.Location, getData1.Location)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.Equal(t, d1.ZipCode, getData1.ZipCode)
	require.Equal(t, d1.Latitude, getData1.Latitude)
	require.Equal(t, d1.Longitude, getData1.Longitude)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	// require.Equal(t, d1.SimpleLocation, getData1.SimpleLocation)
	// require.Equal(t, d1.FullLocation, getData1.FullLocation)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesReference.GetGeography(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Code, getData2.Code)
	require.Equal(t, d2.ShortName, getData2.ShortName)
	require.Equal(t, d2.Location, getData2.Location)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.ParentId, getData2.ParentId)
	require.Equal(t, d2.ZipCode, getData2.ZipCode)
	require.Equal(t, d2.Latitude, getData2.Latitude)
	require.Equal(t, d2.Longitude, getData2.Longitude)
	require.Equal(t, d2.AddressUrl, getData2.AddressUrl)
	// require.Equal(t, d2.SimpleLocation, getData2.SimpleLocation)
	// require.Equal(t, d2.FullLocation, getData2.FullLocation)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesReference.GetGeographybyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)

	// fmt.Printf("%+v\n", getData)

	// Update Data
	updateD2 := d2

	updateD2.Id = getData2.Id
	updateD2.Location = updateD2.Location + "Edited"

	fmt.Printf("updateD2-->%+v\n", updateD2)
	// log.Println(updateD2)
	updatedD1 := updateTestGeography(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Code, updatedD1.Code)
	require.Equal(t, updateD2.ShortName, updatedD1.ShortName)
	require.Equal(t, updateD2.Location, updatedD1.Location)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.ParentId, updatedD1.ParentId)
	require.Equal(t, updateD2.ZipCode, updatedD1.ZipCode)
	require.Equal(t, updateD2.Latitude, updatedD1.Latitude)
	require.Equal(t, updateD2.Longitude, updatedD1.Longitude)
	require.Equal(t, updateD2.AddressUrl, updatedD1.AddressUrl)
	// require.Equal(t, updateD2.SimpleLocation, updatedD1.SimpleLocation)
	// require.Equal(t, updateD2.FullLocation, updatedD1.FullLocation)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteGeography(t, getData1.Id)
	testDeleteGeography(t, getData2.Id)
}

func TestListGeography(t *testing.T) {

	obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "LocationType", 0, "Country")

	arg := ListGeographyParams{
		TypeId: obj.Id,
		Limit:  5,
		Offset: 0,
	}

	geography, err := testQueriesReference.ListGeography(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", geography)
	require.NotEmpty(t, geography)

}

func TestGeographySearch(t *testing.T) {

	arg := SearchGeographyParams{
		SearchText: "West Triangle, Quezon City. Metro Manila",
		Limit:      1,
		Offset:     0,
	}

	geography, err := testQueriesReference.SearchGeography(context.Background(), arg)
	require.NoError(t, err)
	fmt.Printf("%+v\n", geography)
	require.NotEmpty(t, geography)

}

func randomGeography() GeographyRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := GeographyRequest{
		Code:       util.RandomInt(1, 100),
		ShortName:  sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		Location:   util.RandomString(10),
		TypeId:     util.RandomInt(1, 100),
		ParentId:   sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		ZipCode:    sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		Latitude:   util.SetNullDecimal("1.11"),
		Longitude:  util.SetNullDecimal("2.11"),
		AddressUrl: sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		// SimpleLocation: sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		// FullLocation:   sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	fmt.Printf("%+v\n", arg)
	return arg
}

func createTestGeography(
	t *testing.T,
	d1 GeographyRequest) model.Geography {

	getData1, err := testQueriesReference.CreateGeography(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.ShortName, getData1.ShortName)
	require.Equal(t, d1.Location, getData1.Location)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.Equal(t, d1.ZipCode, getData1.ZipCode)
	require.Equal(t, d1.Latitude, getData1.Latitude)
	require.Equal(t, d1.Longitude, getData1.Longitude)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	// require.Equal(t, d1.SimpleLocation, getData1.SimpleLocation)
	// require.Equal(t, d1.FullLocation, getData1.FullLocation)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestGeography(
	t *testing.T,
	d1 GeographyRequest) model.Geography {

	getData1, err := testQueriesReference.UpdateGeography(context.Background(), d1)
	fmt.Println(err)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.ShortName, getData1.ShortName)
	require.Equal(t, d1.Location, getData1.Location)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.Equal(t, d1.ZipCode, getData1.ZipCode)
	require.Equal(t, d1.Latitude, getData1.Latitude)
	require.Equal(t, d1.Longitude, getData1.Longitude)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	// require.Equal(t, d1.SimpleLocation, getData1.SimpleLocation)
	// require.Equal(t, d1.FullLocation, getData1.FullLocation)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteGeography(t *testing.T, id int64) {
	err := testQueriesReference.DeleteGeography(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesReference.GetGeography(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
