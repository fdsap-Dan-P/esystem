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

func TestAddressList(t *testing.T) {

	// Test Data

	d1 := randomAddressList()
	geo, _ := testQueriesReference.SearchGeography(context.Background(),
		ref.SearchGeographyParams{
			SearchText: "West Triangle, Quezon City. Metro Manila",
			Limit:      1,
			Offset:     0,
		})
	d1.GeographyId = sql.NullInt64(sql.NullInt64{Int64: geo[0].Id, Valid: true})

	d2 := randomAddressList()
	geo, _ = testQueriesReference.SearchGeography(context.Background(),
		ref.SearchGeographyParams{
			SearchText: "Soledad San Pablo City, Laguna",
			Limit:      1,
			Offset:     0,
		})
	d2.GeographyId = sql.NullInt64(sql.NullInt64{Int64: geo[0].Id, Valid: true})

	// Test Create
	CreatedD1 := createTestAddressList(t, d1)
	CreatedD2 := createTestAddressList(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetAddressList(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Detail, getData1.Detail)
	require.Equal(t, d1.Url, getData1.Url)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesIdentity.GetAddressList(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.Detail, getData2.Detail)
	require.Equal(t, d2.Url, getData2.Url)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.GeographyId, getData2.GeographyId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesIdentity.GetAddressListbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// getData, err = testQueriesIdentity.GetAddressListbyName(context.Background(), CreatedD1.Name)
	// require.NoError(t, err)
	// require.NotEmpty(t, getData)
	// require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	// fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAddressList(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.Detail, updatedD1.Detail)
	require.Equal(t, updateD2.Url, updatedD1.Url)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.GeographyId, updatedD1.GeographyId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteAddressList(t, CreatedD1.Uuid)
	testDeleteAddressList(t, CreatedD2.Uuid)
}

func TestListAddressList(t *testing.T) {

	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")

	arg := ListAddressListParams{
		Iiid:   ii.Id,
		Limit:  5,
		Offset: 0,
	}

	addressList, err := testQueriesIdentity.ListAddressList(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", addressList)
	require.NotEmpty(t, addressList)

}

func randomAddressList() AddressListRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AddressType", 0, "Current Address")

	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")

	arg := AddressListRequest{
		Iiid:   ii.Id,
		Series: int16(util.RandomInt32(1, 100)),
		Detail: sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		Url:    sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		TypeId: obj.Id,
		// GeographyId: sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestAddressList(
	t *testing.T,
	d1 AddressListRequest) model.AddressList {

	getData1, err := testQueriesIdentity.CreateAddressList(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Detail, getData1.Detail)
	require.Equal(t, d1.Url, getData1.Url)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestAddressList(
	t *testing.T,
	d1 AddressListRequest) model.AddressList {

	getData1, err := testQueriesIdentity.UpdateAddressList(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Detail, getData1.Detail)
	require.Equal(t, d1.Url, getData1.Url)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteAddressList(t *testing.T, uuid uuid.UUID) {
	err := testQueriesIdentity.DeleteAddressList(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetAddressList(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
