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

func TestContacts(t *testing.T) {

	// Test Data
	d1 := randomContacts()
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")
	d1.Iiid = ii.Id
	d2 := randomContacts()
	ii, _ = testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "101")
	d2.Iiid = ii.Id
	// Test Create
	CreatedD1 := createTestContacts(t, d1)
	CreatedD2 := createTestContacts(t, d2)

	// Get Data
	getData1, err1 := testQueriesIdentity.GetContact(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Contact, getData1.Contact)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesIdentity.GetContact(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.Contact, getData2.Contact)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesIdentity.GetContactbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	getData, err = testQueriesIdentity.GetContactbyName(context.Background(), CreatedD1.Contact)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	updateD2.Contact = updateD2.Contact + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestContacts(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.Contact, updatedD1.Contact)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListContacts(t, ListContactParams{
		Iiid:   updatedD1.Iiid,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteContacts(t, CreatedD1.Uuid)
	testDeleteContacts(t, CreatedD2.Uuid)
}

func testListContacts(t *testing.T, arg ListContactParams) {

	Contact, err := testQueriesIdentity.ListContact(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Contact)
	require.NotEmpty(t, Contact)

}

func randomContacts() ContactRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ContactType", 0, "Cellphone")

	arg := ContactRequest{
		// Iiid:    util.RandomInt(1, 100),
		Series:  int16(util.RandomInt32(1, 100)),
		Contact: util.RandomString(10),
		TypeId:  obj.Id,

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestContacts(
	t *testing.T,
	d1 ContactRequest) model.Contact {

	getData1, err := testQueriesIdentity.CreateContact(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Contact, getData1.Contact)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestContacts(
	t *testing.T,
	d1 ContactRequest) model.Contact {

	getData1, err := testQueriesIdentity.UpdateContact(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.Contact, getData1.Contact)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteContacts(t *testing.T, uuid uuid.UUID) {
	err := testQueriesIdentity.DeleteContact(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesIdentity.GetContact(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
