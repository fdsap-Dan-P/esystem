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

func TestIds(t *testing.T) {

	// Test Data
	d1 := randomIds()
	d2 := randomIds()

	// Test Create
	CreatedD1 := createTestIds(t, d1)
	CreatedD2 := createTestIds(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetIds(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.IdNumber, getData1.IdNumber)
	require.Equal(t, d1.RegistrationDate.Time.Format("2006-01-02"), getData1.RegistrationDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.ValidityDate.Time.Format("2006-01-02"), getData1.ValidityDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesIdentity.GetIds(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.IdNumber, getData2.IdNumber)
	require.Equal(t, d2.RegistrationDate.Time.Format("2006-01-02"), getData2.RegistrationDate.Time.Format("2006-01-02"))
	require.Equal(t, d2.ValidityDate.Time.Format("2006-01-02"), getData2.ValidityDate.Time.Format("2006-01-02"))
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesIdentity.GetIdsbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	arg := GetbyIdsParams{
		IdNumber: CreatedD1.IdNumber,
		Limit:    5,
		Offset:   0,
	}

	idList, err2 := testQueriesIdentity.GetbyIds(context.Background(), arg)
	require.NoError(t, err2)
	require.NotEmpty(t, idList)
	var found bool = false
	for _, id := range idList {
		if id.IdNumber == CreatedD1.IdNumber {
			found = true
			break
		}
	}

	require.True(t, found)
	fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	updateD2.IdNumber = updateD2.IdNumber + "Ed"

	// log.Println(updateD2)
	updatedD1 := updateTestIds(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.IdNumber, updatedD1.IdNumber)
	require.Equal(t, updateD2.RegistrationDate.Time.Format("2006-01-02"), updatedD1.RegistrationDate.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.ValidityDate.Time.Format("2006-01-02"), updatedD1.ValidityDate.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteIds(t, CreatedD1.Uuid)
	testDeleteIds(t, CreatedD2.Uuid)
}

func TestListIds(t *testing.T) {

	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")

	arg := ListIdsParams{
		Iiid:   ii.Id,
		Limit:  5,
		Offset: 0,
	}

	ids, err := testQueriesIdentity.ListIds(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", ids)
	require.NotEmpty(t, ids)

}

func randomIds() IdsRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "1001")

	arg := IdsRequest{
		Iiid:             ii.Id,
		Series:           int16(util.RandomInt32(1, 100)),
		IdNumber:         util.RandomString(10),
		RegistrationDate: sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		ValidityDate:     sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		TypeId:           util.RandomInt(1, 100),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestIds(
	t *testing.T,
	d1 IdsRequest) model.Ids {

	getData1, err := testQueriesIdentity.CreateIds(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.IdNumber, getData1.IdNumber)
	require.Equal(t, d1.RegistrationDate.Time.Format("2006-01-02"), getData1.RegistrationDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.ValidityDate.Time.Format("2006-01-02"), getData1.ValidityDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestIds(
	t *testing.T,
	d1 IdsRequest) model.Ids {

	getData1, err := testQueriesIdentity.UpdateIds(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.IdNumber, getData1.IdNumber)
	require.Equal(t, d1.RegistrationDate.Time.Format("2006-01-02"), getData1.RegistrationDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.ValidityDate.Time.Format("2006-01-02"), getData1.ValidityDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteIds(t *testing.T, uuid uuid.UUID) {
	err := testQueriesIdentity.DeleteIds(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetIds(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
