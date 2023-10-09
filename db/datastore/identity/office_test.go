package db

import (
	"context"
	"database/sql"

	"testing"

	"encoding/json"
	common "simplebank/db/common"
	ref "simplebank/db/datastore/reference"
	"simplebank/model"
	"simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestOffice(t *testing.T) {

	// Test Data
	d1 := randomOffice()
	d2 := randomOffice()

	// Test Create
	CreatedD1 := createTestOffice(t, d1)
	CreatedD2 := createTestOffice(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetOffice(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.ShortName, getData1.ShortName)
	require.Equal(t, d1.OfficeName, getData1.OfficeName)
	require.Equal(t, d1.DateStablished.Time.Format("2006-01-02"), getData1.DateStablished.Time.Format("2006-01-02"))
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.Equal(t, d1.AlternateId, getData1.AlternateId)
	require.Equal(t, d1.AddressDetail, getData1.AddressDetail)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.Equal(t, d1.CidSequence, getData1.CidSequence)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesIdentity.GetOffice(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Code, getData2.Code)
	require.Equal(t, d2.ShortName, getData2.ShortName)
	require.Equal(t, d2.OfficeName, getData2.OfficeName)
	require.Equal(t, d2.DateStablished.Time.Format("2006-01-02"), getData2.DateStablished.Time.Format("2006-01-02"))
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.ParentId, getData2.ParentId)
	require.Equal(t, d2.AlternateId, getData2.AlternateId)
	require.Equal(t, d2.AddressDetail, getData2.AddressDetail)
	require.Equal(t, d2.AddressUrl, getData2.AddressUrl)
	require.Equal(t, d2.GeographyId, getData2.GeographyId)
	require.Equal(t, d2.CidSequence, getData2.CidSequence)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesIdentity.GetOfficebyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	// fmt.Printf("Get by UUId%+v\n", getData)

	getData, err = testQueriesIdentity.GetOfficebyShortName(context.Background(), CreatedD1.ParentId.Int64, CreatedD1.ShortName)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	// fmt.Printf("Get by Name%+v\n", getData)

	getData, err = testQueriesIdentity.GetOfficebyCode(context.Background(), CreatedD1.ParentId.Int64, CreatedD1.Code)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	// fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.OfficeName = updateD2.OfficeName + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestOffice(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Code, updatedD1.Code)
	require.Equal(t, updateD2.ShortName, updatedD1.ShortName)
	require.Equal(t, updateD2.OfficeName, updatedD1.OfficeName)
	require.Equal(t, updateD2.DateStablished.Time.Format("2006-01-02"), updatedD1.DateStablished.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.ParentId, updatedD1.ParentId)
	require.Equal(t, updateD2.AlternateId, updatedD1.AlternateId)
	require.Equal(t, updateD2.AddressDetail, updatedD1.AddressDetail)
	require.Equal(t, updateD2.AddressUrl, updatedD1.AddressUrl)
	require.Equal(t, updateD2.GeographyId, updatedD1.GeographyId)
	require.Equal(t, updateD2.CidSequence, updatedD1.CidSequence)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListOffice(t, ListOfficeParams{
		TypeId: updatedD1.TypeId,
		Limit:  5,
		Offset: 0,
	})
	// Delete Data
	testDeleteOffice(t, getData1.Id)
	testDeleteOffice(t, getData2.Id)
}

func testListOffice(t *testing.T, arg ListOfficeParams) {

	office, err := testQueriesIdentity.ListOffice(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", office)
	require.NotEmpty(t, office)

}

func randomOffice() OfficeRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	geo, _ := testQueriesReference.SearchGeography(context.Background(),
		ref.SearchGeographyParams{
			SearchText: "West Triangle, Quezon City. Metro Manila",
			Limit:      1,
			Offset:     0,
		})

	obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "OfficeType", 0, "Institution")

	arg := OfficeRequest{
		Code:           util.RandomString(10),
		ShortName:      util.RandomString(10),
		OfficeName:     util.RandomString(10),
		DateStablished: sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		TypeId:         obj.Id,
		ParentId:       sql.NullInt64(sql.NullInt64{Int64: 0, Valid: false}),
		AlternateId:    sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		AddressDetail:  sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		AddressUrl:     sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		GeographyId:    sql.NullInt64(sql.NullInt64{Int64: geo[0].Id, Valid: true}),
		CidSequence:    sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestOffice(
	t *testing.T,
	d1 OfficeRequest) model.Office {

	getData1, err := testQueriesIdentity.CreateOffice(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.ShortName, getData1.ShortName)
	require.Equal(t, d1.OfficeName, getData1.OfficeName)
	require.Equal(t, d1.DateStablished.Time.Format("2006-01-02"), getData1.DateStablished.Time.Format("2006-01-02"))
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.Equal(t, d1.AlternateId, getData1.AlternateId)
	require.Equal(t, d1.AddressDetail, getData1.AddressDetail)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.Equal(t, d1.CidSequence, getData1.CidSequence)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestOffice(
	t *testing.T,
	d1 OfficeRequest) model.Office {

	getData1, err := testQueriesIdentity.UpdateOffice(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.ShortName, getData1.ShortName)
	require.Equal(t, d1.OfficeName, getData1.OfficeName)
	require.Equal(t, d1.DateStablished.Time.Format("2006-01-02"), getData1.DateStablished.Time.Format("2006-01-02"))
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.Equal(t, d1.AlternateId, getData1.AlternateId)
	require.Equal(t, d1.AddressDetail, getData1.AddressDetail)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.Equal(t, d1.CidSequence, getData1.CidSequence)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteOffice(t *testing.T, id int64) {
	err := testQueriesIdentity.DeleteOffice(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetOffice(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
