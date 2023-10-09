package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	// "log"
	"testing"

	"simplebank/util"

	common "simplebank/db/common"
	"simplebank/model"

	"github.com/stretchr/testify/require"
)

func TestReference(t *testing.T) {

	// Test Data
	d1 := RandomReference()
	d2 := RandomReference()

	log.Println(d1)
	// Test Createx
	CreatedD1 := createTestReference(t, d1)
	CreatedD2 := createTestReference(t, d2)

	// Get Data
	getData1, err1 := testQueriesReference.GetReference(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.ShortName, getData1.ShortName)

	getData2, err2 := testQueriesReference.GetReferenceInfo(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Title, getData2.Title)
	require.Equal(t, d2.ShortName, getData2.ShortName)

	getData2, err2 = testQueriesReference.GetReferenceInfobyUuId(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Title, getData2.Title)
	require.Equal(t, d2.ShortName, getData2.ShortName)

	fmt.Printf("%+v\n", getData2)

	// Update Data
	updateD2 := ReferenceRequest{
		Id:        getData2.Id,
		Code:      getData2.Code,
		ShortName: getData2.ShortName,
		Title:     getData2.Title + "Edited",
		ParentId:  getData2.ParentId,
		TypeId:    getData2.TypeId,
		Remark:    getData2.Remark,
		OtherInfo: getData2.OtherInfo,
	}

	// log.Println(updateD2)
	updatedD1 := updateTestReference(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Title, updatedD1.Title)
	require.Equal(t, updateD2.ShortName, updatedD1.ShortName)

	arg2 := FilterParams{
		Filter: "id in " + util.Int64List2String([]int64{getData1.Id, getData2.Id}),
		Limit:  5,
		Offset: 0,
	}

	reference, err := testQueriesReference.GetReferenceFilter(context.Background(), arg2)
	require.NoError(t, err)
	log.Printf("reference: %+v", reference)
	require.NotEmpty(t, reference)

	for _, ref := range reference {
		require.NotEmpty(t, ref)
		// require.NotEmpty(t, ref)
		// require.Equal(t, lastAccount.Owner, account.Owner)
	}

	arg3 := ReferenceSearchParams{
		Search: "parameter required interest",
		Limit:  2,
		Offset: 1,
	}

	reference, err = testQueriesReference.ReferenceSearch(context.Background(), arg3)
	require.NoError(t, err)

	for _, ref := range reference {
		require.NotEmpty(t, ref)
		// require.NotEmpty(t, ref)
		// require.Equal(t, lastAccount.Owner, account.Owner)
	}

	// Delete Data
	testDeleteReference(t, getData1.Id)
	testDeleteReference(t, getData2.Id)
}

func TestListReference(t *testing.T) {
	// var lastAccount Account
	// for i := 0; i < 10; i++ {
	// 	lastAccount = createRandomAccount(t)
	// }

	arg := ListReferenceParams{
		RefType: "ContactType",
		Limit:   5,
		Offset:  0,
	}

	reference, err := testQueriesReference.ListReference(context.Background(), arg)
	require.NoError(t, err)
	// log.Println(reference)
	require.NotEmpty(t, reference)

	for _, reference := range reference {
		require.NotEmpty(t, reference)
		// require.Equal(t, lastAccount.Owner, account.Owner)
	}

}

func RandomReference() ReferenceRequest {

	refType, _ := testQueriesReference.GetReferenceTypeInfobyTitle(context.Background(), "Parameter")

	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	arg := ReferenceRequest{
		Code:      util.RandomInt(1, 1000),
		ShortName: sql.NullString(sql.NullString{String: util.RandomString(20), Valid: true}),
		Title:     util.RandomString(50),
		ParentId:  sql.NullInt64(sql.NullInt64{Int64: util.RandomInt(1, 2), Valid: true}),
		TypeId:    refType.Id,
		Remark:    sql.NullString(sql.NullString{String: util.RandomString(100), Valid: true}),
		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestReference(
	t *testing.T,
	createData ReferenceRequest) model.Reference {

	reference, err := testQueriesReference.CreateReference(context.Background(), createData)
	require.NoError(t, err)
	require.NotEmpty(t, reference)

	require.Equal(t, createData.Code, reference.Code)
	require.Equal(t, createData.Title, reference.Title)
	require.Equal(t, createData.ShortName, reference.ShortName)
	return reference
}

func updateTestReference(
	t *testing.T,
	updateData ReferenceRequest) model.Reference {

	reference, err := testQueriesReference.UpdateReference(context.Background(), updateData)
	require.NoError(t, err)
	require.NotEmpty(t, reference)

	require.Equal(t, updateData.Code, reference.Code)
	require.Equal(t, updateData.Title, reference.Title)
	require.Equal(t, updateData.ShortName, reference.ShortName)
	return reference
}

func testDeleteReference(t *testing.T, id int64) {
	err := testQueriesReference.DeleteReference(context.Background(), id)
	require.NoError(t, err)

	data2, err := testQueriesReference.GetReference(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, data2)
}
