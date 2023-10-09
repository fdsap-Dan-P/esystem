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

func TestOfficeAccountType(t *testing.T) {

	// Test Data
	d1 := randomOfficeAccountType()
	d2 := randomOfficeAccountType()

	// Test Create
	CreatedD1 := createTestOfficeAccountType(t, d1)
	CreatedD2 := createTestOfficeAccountType(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetOfficeAccountType(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.OfficeAccountType, getData1.OfficeAccountType)
	require.Equal(t, d1.CoaId, getData1.CoaId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetOfficeAccountType(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.OfficeAccountType, getData2.OfficeAccountType)
	require.Equal(t, d2.CoaId, getData2.CoaId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetOfficeAccountTypebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)

	fmt.Printf("%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.OfficeAccountType = updateD2.OfficeAccountType + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestOfficeAccountType(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.OfficeAccountType, updatedD1.OfficeAccountType)
	require.Equal(t, updateD2.CoaId, updatedD1.CoaId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	// Delete Data
	testDeleteOfficeAccountType(t, getData1.Id)
	testDeleteOfficeAccountType(t, getData2.Id)
}

func TestListOfficeAccountType(t *testing.T) {

	arg := ListOfficeAccountTypeParams{
		Limit:  5,
		Offset: 0,
	}

	officeAccountType, err := testQueriesAccount.ListOfficeAccountType(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", officeAccountType)
	require.NotEmpty(t, officeAccountType)

}

func randomOfficeAccountType() OfficeAccountTypeRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := OfficeAccountTypeRequest{
		OfficeAccountType: util.RandomString(10),
		CoaId:             util.RandomInt(1, 100),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestOfficeAccountType(
	t *testing.T,
	d1 OfficeAccountTypeRequest) model.OfficeAccountType {

	getData1, err := testQueriesAccount.CreateOfficeAccountType(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.OfficeAccountType, getData1.OfficeAccountType)
	require.Equal(t, d1.CoaId, getData1.CoaId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestOfficeAccountType(
	t *testing.T,
	d1 OfficeAccountTypeRequest) model.OfficeAccountType {

	getData1, err := testQueriesAccount.UpdateOfficeAccountType(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.OfficeAccountType, getData1.OfficeAccountType)
	require.Equal(t, d1.CoaId, getData1.CoaId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteOfficeAccountType(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteOfficeAccountType(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetOfficeAccountType(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
