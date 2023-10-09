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

func TestServer(t *testing.T) {

	// Test Data
	d1 := randomServer()
	d2 := randomServer()

	// Test Create
	CreatedD1 := createTestServer(t, d1)
	CreatedD2 := createTestServer(t, d2)

	// Get Data
	getData1, err1 := testQueriesDocument.GetServer(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.Connectivity, getData1.Connectivity)
	require.Equal(t, d1.NetAddress, getData1.NetAddress)
	require.Equal(t, d1.Certificate, getData1.Certificate)
	require.Equal(t, d1.Description, getData1.Description)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesDocument.GetServer(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Code, getData2.Code)
	require.Equal(t, d2.Connectivity, getData2.Connectivity)
	require.Equal(t, d2.NetAddress, getData2.NetAddress)
	require.Equal(t, d2.Certificate, getData2.Certificate)
	require.Equal(t, d2.Description, getData2.Description)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesDocument.GetServerbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	getData, err = testQueriesDocument.GetServerbyCode(context.Background(), CreatedD1.Code)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.Description.String = updateD2.Description.String + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestServer(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Code, updatedD1.Code)
	require.Equal(t, updateD2.Connectivity, updatedD1.Connectivity)
	require.Equal(t, updateD2.NetAddress, updatedD1.NetAddress)
	require.Equal(t, updateD2.Certificate, updatedD1.Certificate)
	require.Equal(t, updateD2.Description, updatedD1.Description)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListServer(t, ListServerParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteServer(t, getData1.Id)
	testDeleteServer(t, getData2.Id)
}

func testListServer(t *testing.T, arg ListServerParams) {

	Server, err := testQueriesDocument.ListServer(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Server)
	require.NotEmpty(t, Server)

}

func randomServer() ServerRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := ServerRequest{
		Code:         util.RandomString(30),
		Connectivity: model.Local,
		NetAddress:   util.RandomString(10),
		Certificate:  sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		Description:  sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestServer(
	t *testing.T,
	d1 ServerRequest) model.Server {

	getData1, err := testQueriesDocument.CreateServer(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.Connectivity, getData1.Connectivity)
	require.Equal(t, d1.NetAddress, getData1.NetAddress)
	require.Equal(t, d1.Certificate, getData1.Certificate)
	require.Equal(t, d1.Description, getData1.Description)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestServer(
	t *testing.T,
	d1 ServerRequest) model.Server {

	getData1, err := testQueriesDocument.UpdateServer(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.Connectivity, getData1.Connectivity)
	require.Equal(t, d1.NetAddress, getData1.NetAddress)
	require.Equal(t, d1.Certificate, getData1.Certificate)
	require.Equal(t, d1.Description, getData1.Description)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteServer(t *testing.T, id int64) {
	err := testQueriesDocument.DeleteServer(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesDocument.GetServer(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
