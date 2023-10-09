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

	ref "simplebank/db/datastore/reference"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestInventoryRepository(t *testing.T) {

	// Test Data
	d1 := randomInventoryRepository()
	d1.Uuid = uuid.MustParse("04b80e91-a1ef-4b3b-abe5-f40c158d1c6e")
	d2 := randomInventoryRepository()

	// Test Create
	CreatedD1 := createTestInventoryRepository(t, d1)
	CreatedD2 := createTestInventoryRepository(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetInventoryRepository(context.Background(), CreatedD1.Id)

	log.Printf("GetInventoryRepository %+v: %+v", CreatedD1, getData1)
	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, getData1.CentralOfficeId, CreatedD1.CentralOfficeId)
	require.Equal(t, getData1.RepositoryCode, CreatedD1.RepositoryCode)
	require.Equal(t, getData1.Repository, CreatedD1.Repository)
	require.Equal(t, getData1.OfficeId, CreatedD1.OfficeId)
	require.Equal(t, getData1.CustodianId, CreatedD1.CustodianId)
	require.Equal(t, getData1.GeographyId, CreatedD1.GeographyId)
	require.Equal(t, getData1.LocationDescription, CreatedD1.LocationDescription)
	require.Equal(t, getData1.Remarks, CreatedD1.Remarks)
	require.JSONEq(t, getData1.OtherInfo.String, CreatedD2.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetInventoryRepository(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CentralOfficeId, CreatedD2.CentralOfficeId)
	require.Equal(t, d2.RepositoryCode, CreatedD2.RepositoryCode)
	require.Equal(t, d2.Repository, CreatedD2.Repository)
	require.Equal(t, d2.OfficeId, CreatedD2.OfficeId)
	require.Equal(t, d2.CustodianId, CreatedD2.CustodianId)
	require.Equal(t, d2.GeographyId, CreatedD2.GeographyId)
	require.Equal(t, d2.LocationDescription, CreatedD2.LocationDescription)
	require.Equal(t, d2.Remarks, CreatedD2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetInventoryRepositorybyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.LocationDescription.String = updateD2.LocationDescription.String + "-Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestInventoryRepository(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.CentralOfficeId, updatedD1.CentralOfficeId)
	require.Equal(t, updateD2.RepositoryCode, updatedD1.RepositoryCode)
	require.Equal(t, updateD2.Repository, updatedD1.Repository)
	require.Equal(t, updateD2.OfficeId, updatedD1.OfficeId)
	require.Equal(t, updateD2.CustodianId, updatedD1.CustodianId)
	require.Equal(t, updateD2.GeographyId, updatedD1.GeographyId)
	require.Equal(t, updateD2.LocationDescription, updatedD1.LocationDescription)
	require.Equal(t, updateD2.Remarks, updatedD1.Remarks)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListInventoryRepository(t, ListInventoryRepositoryParams{
		OfficeId: updatedD1.OfficeId,
		Limit:    5,
		Offset:   0,
	})

	// Delete Data
	testDeleteInventoryRepository(t, getData1.Id)
	// testDeleteInventoryRepository(t, getData2.Id)
}

func testListInventoryRepository(t *testing.T, arg ListInventoryRepositoryParams) {

	inventoryRepository, err := testQueriesAccount.ListInventoryRepository(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", inventoryRepository)
	require.NotEmpty(t, inventoryRepository)

}

func randomInventoryRepository() InventoryRepositoryRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")

	geoParam := ref.SearchGeographyParams{
		SearchText: "Soledad San Pablo City, Laguna",
		Limit:      1,
		Offset:     0,
	}

	geo, _ := testQueriesReference.SearchGeography(context.Background(), geoParam)
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")

	arg := InventoryRepositoryRequest{
		CentralOfficeId:     ofc.Id,
		RepositoryCode:      util.RandomString(10),
		Repository:          util.RandomString(10),
		OfficeId:            ofc.Id,
		CustodianId:         util.SetNullInt64(ii.Id),
		GeographyId:         util.SetNullInt64(geo[0].Id),
		LocationDescription: util.RandomNullString(10),
		Remarks:             util.RandomNullString(10),
		OtherInfo:           sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestInventoryRepository(
	t *testing.T,
	CreatedD1 InventoryRepositoryRequest) model.InventoryRepository {

	getData1, err := testQueriesAccount.CreateInventoryRepository(context.Background(), CreatedD1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, getData1.CentralOfficeId, CreatedD1.CentralOfficeId)
	require.Equal(t, getData1.RepositoryCode, CreatedD1.RepositoryCode)
	require.Equal(t, getData1.Repository, CreatedD1.Repository)
	require.Equal(t, getData1.OfficeId, CreatedD1.OfficeId)
	require.Equal(t, getData1.CustodianId, CreatedD1.CustodianId)
	require.Equal(t, getData1.GeographyId, CreatedD1.GeographyId)
	require.Equal(t, getData1.LocationDescription, CreatedD1.LocationDescription)
	require.Equal(t, getData1.Remarks, CreatedD1.Remarks)
	require.JSONEq(t, getData1.OtherInfo.String, CreatedD1.OtherInfo.String)

	return getData1
}

func updateTestInventoryRepository(
	t *testing.T,
	CreatedD1 InventoryRepositoryRequest) model.InventoryRepository {

	getData1, err := testQueriesAccount.UpdateInventoryRepository(context.Background(), CreatedD1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, getData1.CentralOfficeId, CreatedD1.CentralOfficeId)
	require.Equal(t, getData1.RepositoryCode, CreatedD1.RepositoryCode)
	require.Equal(t, getData1.Repository, CreatedD1.Repository)
	require.Equal(t, getData1.OfficeId, CreatedD1.OfficeId)
	require.Equal(t, getData1.CustodianId, CreatedD1.CustodianId)
	require.Equal(t, getData1.GeographyId, CreatedD1.GeographyId)
	require.Equal(t, getData1.LocationDescription, CreatedD1.LocationDescription)
	require.Equal(t, getData1.Remarks, CreatedD1.Remarks)
	require.JSONEq(t, getData1.OtherInfo.String, CreatedD1.OtherInfo.String)

	return getData1
}

func testDeleteInventoryRepository(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteInventoryRepository(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetInventoryRepository(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
