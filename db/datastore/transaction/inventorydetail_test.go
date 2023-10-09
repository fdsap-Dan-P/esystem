package db

// import (
// 	"context"
// 	"database/sql"

// 	"fmt"
// 	"testing"

// 	"encoding/json"
// 	common "simplebank/db/common"
// 	"simplebank/model"
// 	"simplebank/util"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// )

// func TestInventoryDetail(t *testing.T) {

// 	// Test Data
// 	d1 := randomInventoryDetail()

// 	// accQtl, _ := testQueriesAccount.GetAccountInventorybyUuid(context.Background(), uuid.MustParse("c3476afe-bd50-49e6-8de3-074555a8e1bd"))
// 	// d1.InventoryDetailId = accQtl.Id

// 	d2 := randomInventoryDetail()
// 	d2.Uuid = uuid.MustParse("419f4b7f-c049-470d-b609-af465d0e8ba4")

// 	acc, _ := testQueriesAccount.CreateAccountInventory(context.Background(), randomAccountInventory("1001-0001-0000002"))
// 	d2.AccountInventoryId = acc.Id

// 	// accQtl, _ = testQueriesAccount.GetAccountInventorybyUuid(context.Background(), uuid.MustParse("b35e39e8-885b-41a6-a070-a249c2a099e5"))
// 	// d2.InventoryDetailId = accQtl.Id

// 	// Test Create
// 	CreatedD1 := createTestInventoryDetail(t, d1)
// 	CreatedD2 := createTestInventoryDetail(t, d2)

// 	// Get Data
// 	getData1, err1 := testQueriesTransaction.GetInventoryDetailbyUuid(context.Background(), CreatedD1.Uuid)

// 	require.NoError(t, err1)
// 	require.NotEmpty(t, getData1)
// 	require.Equal(t, d1.Uuid, getData1.Uuid)
// 	require.Equal(t, d1.AccountInventoryId, getData1.AccountInventoryId)
// 	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
// 	require.Equal(t, d1.RepositoryId, getData1.RepositoryId)
// 	require.Equal(t, d1.SupplierId, getData1.SupplierId)
// 	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
// 	require.Equal(t, d1.BookValue.String(), getData1.BookValue.String())
// 	require.Equal(t, d1.Unit.String(), getData1.Unit.String())
// 	require.Equal(t, d1.MeasureId, getData1.MeasureId)
// 	require.Equal(t, d1.BatchNumber, getData1.BatchNumber)
// 	require.Equal(t, d1.DateManufactured.Time.Format("2006-01-02"), getData1.DateManufactured.Time.Format("2006-01-02"))
// 	require.Equal(t, d1.DateExpired.Time.Format("2006-01-02"), getData1.DateExpired.Time.Format("2006-01-02"))
// 	require.Equal(t, d1.Remarks, getData1.Remarks)
// 	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

// 	getData2, err2 := testQueriesTransaction.GetInventoryDetailbyUuid(context.Background(), CreatedD2.Uuid)
// 	require.NoError(t, err2)
// 	require.NotEmpty(t, getData2)
// 	require.Equal(t, d2.Uuid, getData2.Uuid)
// 	require.Equal(t, d2.AccountInventoryId, getData2.AccountInventoryId)
// 	require.Equal(t, d2.InventoryItemId, getData2.InventoryItemId)
// 	require.Equal(t, d2.RepositoryId, getData2.RepositoryId)
// 	require.Equal(t, d2.SupplierId, getData2.SupplierId)
// 	require.Equal(t, d2.UnitPrice.String(), getData2.UnitPrice.String())
// 	require.Equal(t, d2.BookValue.String(), getData2.BookValue.String())
// 	require.Equal(t, d2.Unit.String(), getData2.Unit.String())
// 	require.Equal(t, d2.MeasureId, getData2.MeasureId)
// 	require.Equal(t, d2.BatchNumber, getData2.BatchNumber)
// 	require.Equal(t, d2.DateManufactured.Time.Format("2006-01-02"), getData2.DateManufactured.Time.Format("2006-01-02"))
// 	require.Equal(t, d2.DateExpired.Time.Format("2006-01-02"), getData2.DateExpired.Time.Format("2006-01-02"))
// 	require.Equal(t, d2.Remarks, getData2.Remarks)
// 	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

// 	getData, err := testQueriesTransaction.GetInventoryDetailbyUuid(context.Background(), CreatedD1.Uuid)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, getData)
// 	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
// 	fmt.Printf("Get by UUid%+v\n", getData)

// 	// Update Data
// 	updateD2 := d2
// 	updateD2.Id = getData2.Id
// 	// updateD2.Location = updateD2.Location + "Edited"

// 	// log.Println(updateD2)
// 	updatedD1 := updateTestInventoryDetail(t, updateD2)
// 	require.NotEmpty(t, updatedD1)
// 	require.Equal(t, updateD2.Uuid, updatedD1.Uuid)
// 	require.Equal(t, updateD2.AccountInventoryId, updatedD1.AccountInventoryId)
// 	require.Equal(t, updateD2.InventoryItemId, updatedD1.InventoryItemId)
// 	require.Equal(t, updateD2.RepositoryId, updatedD1.RepositoryId)
// 	require.Equal(t, updateD2.SupplierId, updatedD1.SupplierId)
// 	require.Equal(t, updateD2.UnitPrice.String(), updatedD1.UnitPrice.String())
// 	require.Equal(t, updateD2.BookValue.String(), updatedD1.BookValue.String())
// 	require.Equal(t, updateD2.Unit.String(), updatedD1.Unit.String())
// 	require.Equal(t, updateD2.MeasureId, updatedD1.MeasureId)
// 	require.Equal(t, updateD2.BatchNumber, updatedD1.BatchNumber)
// 	require.Equal(t, updateD2.DateManufactured.Time.Format("2006-01-02"), updatedD1.DateManufactured.Time.Format("2006-01-02"))
// 	require.Equal(t, updateD2.DateExpired.Time.Format("2006-01-02"), updatedD1.DateExpired.Time.Format("2006-01-02"))
// 	require.Equal(t, updateD2.Remarks, updatedD1.Remarks)
// 	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

// 	testListInventoryDetail(t, ListInventoryDetailParams{
// 		Limit:  5,
// 		Offset: 0,
// 	})

// 	// Delete Data
// 	// testDeleteInventoryDetail(t, CreatedD1.Uuid)
// 	// testDeleteInventoryDetail(t, CreatedD2.Uuid)
// }

// func testListInventoryDetail(t *testing.T, arg ListInventoryDetailParams) {

// 	InventoryDetail, err := testQueriesTransaction.ListInventoryDetail(context.Background(), arg)
// 	require.NoError(t, err)
// 	// fmt.Printf("%+v\n", InventoryDetail)
// 	require.NotEmpty(t, InventoryDetail)

// }

// func randomInventoryDetail() InventoryDetailRequest {
// 	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
// 	info, _ := json.Marshal(otherInfo)

// 	meas, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "UnitMeasure", 0, "Item")
// 	acc, _ := testQueriesAccount.CreateAccountInventory(context.Background(), randomAccountInventory("1001-0001-0000001"))
// 	inv, _ := testQueriesAccount.CreateInventoryItem(context.Background(), randomInventoryItem())
// 	repo, _ := testQueriesAccount.CreateInventoryRepository(context.Background(), randomInventoryRepository())
// 	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")

// 	arg := InventoryDetailRequest{
// 		Uuid:               uuid.MustParse("24e733fa-1119-4096-a434-37b9728521c5"),
// 		AccountInventoryId: acc.Id,
// 		InventoryItemId:    inv.Id,
// 		RepositoryId:       util.SetNullInt64(repo.Id),
// 		SupplierId:         util.SetNullInt64(ii.Id),
// 		UnitPrice:          util.RandomMoney(),
// 		BookValue:          util.RandomMoney(),
// 		Unit:               util.RandomMoney(),
// 		MeasureId:          meas.Id,
// 		BatchNumber:        util.SetNullString("BatchNum"),
// 		DateManufactured:   util.RandomNullDate(),
// 		DateExpired:        util.RandomNullDate(),
// 		Remarks:            "String",
// 		OtherInfo:          sql.NullString(sql.NullString{String: string(info), Valid: true}),
// 	}
// 	return arg
// }

// func createTestInventoryDetail(
// 	t *testing.T,
// 	d1 InventoryDetailRequest) model.InventoryDetail {

// 	getData1, err := testQueriesTransaction.CreateInventoryDetail(context.Background(), d1)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, getData1)

// 	require.Equal(t, d1.Uuid, getData1.Uuid)
// 	require.Equal(t, d1.AccountInventoryId, getData1.AccountInventoryId)
// 	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
// 	require.Equal(t, d1.RepositoryId, getData1.RepositoryId)
// 	require.Equal(t, d1.SupplierId, getData1.SupplierId)
// 	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
// 	require.Equal(t, d1.BookValue.String(), getData1.BookValue.String())
// 	require.Equal(t, d1.Unit.String(), getData1.Unit.String())
// 	require.Equal(t, d1.MeasureId, getData1.MeasureId)
// 	require.Equal(t, d1.BatchNumber, getData1.BatchNumber)
// 	require.Equal(t, d1.DateManufactured.Time.Format("2006-01-02"), getData1.DateManufactured.Time.Format("2006-01-02"))
// 	require.Equal(t, d1.DateExpired.Time.Format("2006-01-02"), getData1.DateExpired.Time.Format("2006-01-02"))
// 	require.Equal(t, d1.Remarks, getData1.Remarks)
// 	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

// 	return getData1
// }

// func updateTestInventoryDetail(
// 	t *testing.T,
// 	d1 InventoryDetailRequest) model.InventoryDetail {

// 	getData1, err := testQueriesTransaction.UpdateInventoryDetail(context.Background(), d1)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, getData1)

// 	require.Equal(t, d1.Uuid, getData1.Uuid)
// 	require.Equal(t, d1.AccountInventoryId, getData1.AccountInventoryId)
// 	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
// 	require.Equal(t, d1.RepositoryId, getData1.RepositoryId)
// 	require.Equal(t, d1.SupplierId, getData1.SupplierId)
// 	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
// 	require.Equal(t, d1.BookValue.String(), getData1.BookValue.String())
// 	require.Equal(t, d1.Unit.String(), getData1.Unit.String())
// 	require.Equal(t, d1.MeasureId, getData1.MeasureId)
// 	require.Equal(t, d1.BatchNumber, getData1.BatchNumber)
// 	require.Equal(t, d1.DateManufactured.Time.Format("2006-01-02"), getData1.DateManufactured.Time.Format("2006-01-02"))
// 	require.Equal(t, d1.DateExpired.Time.Format("2006-01-02"), getData1.DateExpired.Time.Format("2006-01-02"))
// 	require.Equal(t, d1.Remarks, getData1.Remarks)
// 	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

// 	return getData1
// }

// func testDeleteInventoryDetail(t *testing.T, uuid uuid.UUID) {
// 	err := testQueriesTransaction.DeleteInventoryDetail(context.Background(), uuid)
// 	require.NoError(t, err)

// 	ref1, err := testQueriesTransaction.GetInventoryDetailbyUuid(context.Background(), uuid)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, ref1)
// }
