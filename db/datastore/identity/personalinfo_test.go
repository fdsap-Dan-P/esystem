package db

import (
	"context"
	"database/sql"
	"log"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestPersonalInfo(t *testing.T) {

	// Test Data
	d1 := randomPersonalInfo()
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")
	d1.Id = ii.Id
	d2 := randomPersonalInfo()
	ii, _ = testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "101")
	d2.Id = ii.Id

	log.Println("date1", d1.MarriageDate.Time.Format("2006-01-02"), d1.Id)

	// Test Create
	CreatedD1 := createTestPersonalInfo(t, d1)
	CreatedD2 := createTestPersonalInfo(t, d2)
	log.Println("date2", CreatedD1.MarriageDate.Time.Format("2006-01-02"), CreatedD1.Id)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetPersonalInfo(context.Background(), CreatedD1.Id)
	log.Println("date3", getData1.MarriageDate.Time.Format("2006-01-02"), getData1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Isadopted, getData1.Isadopted)
	require.Equal(t, d1.MarriageDate.Time.Format("2006-01-02"), getData1.MarriageDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.KnownLanguage, getData1.KnownLanguage)
	require.Equal(t, d1.IndustryId, getData1.IndustryId)
	require.Equal(t, d1.NationalityId, getData1.NationalityId)
	require.Equal(t, d1.OccupationId, getData1.OccupationId)
	require.Equal(t, d1.ReligionId, getData1.ReligionId)
	require.Equal(t, d1.SectorId, getData1.SectorId)
	require.Equal(t, d1.SourceIncomeId, getData1.SourceIncomeId)
	require.Equal(t, d1.DisabilityId, getData1.DisabilityId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesIdentity.GetPersonalInfo(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Isadopted, getData2.Isadopted)
	require.Equal(t, d2.MarriageDate.Time.Format("2006-01-02"), getData2.MarriageDate.Time.Format("2006-01-02"))
	require.Equal(t, d2.KnownLanguage, getData2.KnownLanguage)
	require.Equal(t, d2.IndustryId, getData2.IndustryId)
	require.Equal(t, d2.NationalityId, getData2.NationalityId)
	require.Equal(t, d2.OccupationId, getData2.OccupationId)
	require.Equal(t, d2.ReligionId, getData2.ReligionId)
	require.Equal(t, d2.SectorId, getData2.SectorId)
	require.Equal(t, d2.SourceIncomeId, getData2.SourceIncomeId)
	require.Equal(t, d2.DisabilityId, getData2.DisabilityId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesIdentity.GetPersonalInfobyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// getData, err = testQueriesIdentity.GetPersonalInfobyName(context.Background(), CreatedD1.Name)
	// require.NoError(t, err)
	// require.NotEmpty(t, getData)
	// require.Equal(t, CreatedD1.Id, getData.Id)
	// fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestPersonalInfo(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Isadopted, updatedD1.Isadopted)
	require.Equal(t, updateD2.MarriageDate.Time.Format("2006-01-02"), updatedD1.MarriageDate.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.KnownLanguage, updatedD1.KnownLanguage)
	require.Equal(t, updateD2.IndustryId, updatedD1.IndustryId)
	require.Equal(t, updateD2.NationalityId, updatedD1.NationalityId)
	require.Equal(t, updateD2.OccupationId, updatedD1.OccupationId)
	require.Equal(t, updateD2.ReligionId, updatedD1.ReligionId)
	require.Equal(t, updateD2.SectorId, updatedD1.SectorId)
	require.Equal(t, updateD2.SourceIncomeId, updatedD1.SourceIncomeId)
	require.Equal(t, updateD2.DisabilityId, updatedD1.DisabilityId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)
	testListPersonalInfo(t)

	// Delete Data
	testDeletePersonalInfo(t, getData1.Id)
	testDeletePersonalInfo(t, getData2.Id)
}

func testListPersonalInfo(t *testing.T) {

	arg := ListPersonalInfoParams{
		Limit:  5,
		Offset: 0,
	}

	personalInfo, err := testQueriesIdentity.ListPersonalInfo(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", personalInfo)
	require.NotEmpty(t, personalInfo)

}

func randomPersonalInfo() PersonalInfoRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := PersonalInfoRequest{
		// Id:             ii.Id,
		Isadopted:      sql.NullBool(sql.NullBool{Bool: true, Valid: true}),
		MarriageDate:   sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		KnownLanguage:  sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		IndustryId:     sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		NationalityId:  sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		OccupationId:   sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		ReligionId:     sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		SectorId:       sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		SourceIncomeId: sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),
		DisabilityId:   sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestPersonalInfo(
	t *testing.T,
	d1 PersonalInfoRequest) model.PersonalInfo {

	getData1, err := testQueriesIdentity.CreatePersonalInfo(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Isadopted, getData1.Isadopted)
	require.Equal(t, d1.MarriageDate.Time.Format("2006-01-02"), getData1.MarriageDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.KnownLanguage, getData1.KnownLanguage)
	require.Equal(t, d1.IndustryId, getData1.IndustryId)
	require.Equal(t, d1.NationalityId, getData1.NationalityId)
	require.Equal(t, d1.OccupationId, getData1.OccupationId)
	require.Equal(t, d1.ReligionId, getData1.ReligionId)
	require.Equal(t, d1.SectorId, getData1.SectorId)
	require.Equal(t, d1.SourceIncomeId, getData1.SourceIncomeId)
	require.Equal(t, d1.DisabilityId, getData1.DisabilityId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestPersonalInfo(
	t *testing.T,
	d1 PersonalInfoRequest) model.PersonalInfo {

	getData1, err := testQueriesIdentity.UpdatePersonalInfo(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Isadopted, getData1.Isadopted)
	require.Equal(t, d1.MarriageDate.Time.Format("2006-01-02"), getData1.MarriageDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.KnownLanguage, getData1.KnownLanguage)
	require.Equal(t, d1.IndustryId, getData1.IndustryId)
	require.Equal(t, d1.NationalityId, getData1.NationalityId)
	require.Equal(t, d1.OccupationId, getData1.OccupationId)
	require.Equal(t, d1.ReligionId, getData1.ReligionId)
	require.Equal(t, d1.SectorId, getData1.SectorId)
	require.Equal(t, d1.SourceIncomeId, getData1.SourceIncomeId)
	require.Equal(t, d1.DisabilityId, getData1.DisabilityId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeletePersonalInfo(t *testing.T, id int64) {
	err := testQueriesIdentity.DeletePersonalInfo(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetPersonalInfo(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
