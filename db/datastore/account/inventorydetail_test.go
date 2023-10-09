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

// var store StoreAccount

//	func init() {
//		store = NewStoreAccount(testDB)
//	}
func TestInventoryDetail(t *testing.T) {

	// Test Data

	// store := NewStoreAccount(testDB)
	d1 := randomInventoryDetail()
	d2 := randomInventoryDetailFull()
	d2.Uuid = uuid.MustParse("419f4b7f-c049-470d-b609-af465d0e8ba4")

	fmt.Printf("Get by UUId%+v\n", d1)
	// Test Create
	CreatedD1 := createTestInventoryDetail(t, d1)
	CreatedD2 := createTestInventoryDetailFull(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetInventoryDetail(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AccountInventoryId, getData1.AccountInventoryId)
	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SupplierId, getData1.SupplierId)
	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
	require.Equal(t, d1.BookValue.String(), getData1.BookValue.String())
	require.Equal(t, d1.Unit.String(), getData1.Unit.String())
	require.Equal(t, d1.MeasureId, getData1.MeasureId)
	require.Equal(t, d1.BatchNumber, getData1.BatchNumber)
	require.Equal(t, d1.DateManufactured.Time.Format("2006-01-02"), getData1.DateManufactured.Time.Format("2006-01-02"))
	require.Equal(t, d1.BatchNumber, getData1.BatchNumber)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetInventoryDetail(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AccountInventoryId, getData2.AccountInventoryId)

	InvItem, err := testQueriesAccount.GetInventoryItem(context.Background(), getData2.InventoryItemId)
	require.NoError(t, err)

	require.Equal(t, d2.InventoryItem.BarCode, InvItem.BarCode)
	require.Equal(t, d2.InventoryItem.ItemName, InvItem.ItemName)
	require.Equal(t, d2.InventoryItem.UniqueVariation, InvItem.UniqueVariation)
	require.Equal(t, d2.InventoryItem.ParentId.Int64, InvItem.ParentId.Int64)
	require.Equal(t, d2.InventoryItem.GenericNameId.Int64, InvItem.GenericNameId.Int64)
	require.Equal(t, d2.InventoryItem.BrandNameId.Int64, InvItem.BrandNameId.Int64)
	require.Equal(t, d2.InventoryItem.MeasureId, InvItem.MeasureId)
	require.Equal(t, d2.InventoryItem.Remarks, InvItem.Remarks)
	require.JSONEq(t, d2.InventoryItem.OtherInfo.String, InvItem.OtherInfo.String)

	require.Equal(t, d2.SupplierId, getData2.SupplierId)
	require.Equal(t, d2.UnitPrice.String(), getData2.UnitPrice.String())
	require.Equal(t, d2.BookValue.String(), getData2.BookValue.String())
	require.Equal(t, d2.Unit.String(), getData2.Unit.String())
	require.Equal(t, d2.MeasureId, getData2.MeasureId)
	require.Equal(t, d2.BatchNumber, getData2.BatchNumber)
	require.Equal(t, d2.DateManufactured.Time.Format("2006-01-02"), getData2.DateManufactured.Time.Format("2006-01-02"))
	require.Equal(t, d2.BatchNumber, getData2.BatchNumber)
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetInventoryDetailbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d1
	updateD2.Id = getData1.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestInventoryDetail(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updatedD1.AccountInventoryId, updateD2.AccountInventoryId)
	require.Equal(t, updatedD1.InventoryItemId, updateD2.InventoryItemId)
	require.Equal(t, updatedD1.SupplierId, updateD2.SupplierId)
	require.Equal(t, updatedD1.UnitPrice.String(), updateD2.UnitPrice.String())
	require.Equal(t, updatedD1.BookValue.String(), updateD2.BookValue.String())
	require.Equal(t, updatedD1.Unit.String(), updateD2.Unit.String())
	require.Equal(t, updatedD1.MeasureId, updateD2.MeasureId)
	require.Equal(t, updatedD1.BatchNumber, updateD2.BatchNumber)
	require.Equal(t, updatedD1.DateManufactured.Time.Format("2006-01-02"), updateD2.DateManufactured.Time.Format("2006-01-02"))
	require.Equal(t, updatedD1.BatchNumber, updateD2.BatchNumber)
	require.Equal(t, updatedD1.Remarks, updateD2.Remarks)
	require.JSONEq(t, updatedD1.OtherInfo.String, updateD2.OtherInfo.String)

	testListInventoryDetail(t, ListInventoryDetailParams{
		AccountInventoryId: updatedD1.AccountInventoryId,
		Limit:              5,
		Offset:             0,
	})

	// Delete Data
	//testDeleteInventoryDetail(t, getData1.Id)
	testDeleteInventoryDetail(t, getData2.Id)
}

func testListInventoryDetail(t *testing.T, arg ListInventoryDetailParams) {
	// store := NewStoreAccount(testDB)
	InventoryDetail, err := testQueriesAccount.ListInventoryDetail(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", InventoryDetail)
	require.NotEmpty(t, InventoryDetail)
}

func randomInventoryDetail() InventoryDetailRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	// acc, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "LoanClass", 0, "Current")
	accInv, _ := testQueriesAccount.GetAccountInventorybyUuid(context.Background(), util.ToUUID("de5e9bff-4fa4-4470-92ca-d9776268230c"))
	invItem, _ := testQueriesAccount.GetInventoryItembyUuid(context.Background(), util.ToUUID("090db518-587c-41a3-9baa-9dc70dae58f8"))
	measure, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "UnitMeasure", 0, "Milliliter")

	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")

	arg := InventoryDetailRequest{
		// AccountId:     util.RandomInt(1, 100),
		AccountInventoryId: accInv.Id,
		InventoryItemId:    invItem.Id,
		SupplierId:         sql.NullInt64(sql.NullInt64{Int64: ii.Id, Valid: true}),
		UnitPrice:          util.RandomMoney(),
		BookValue:          util.RandomMoney(),
		Unit:               util.RandomMoney(),
		MeasureId:          measure.Id,
		BatchNumber:        sql.NullString{String: util.RandomString(30), Valid: true},
		DateManufactured:   sql.NullTime{Time: util.RandomDate(), Valid: true},
		DateExpired:        sql.NullTime{Time: util.RandomDate(), Valid: true},
		Remarks:            util.RandomString(10),
		OtherInfo:          sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}

	return arg
}

func randomInventoryDetailFull() InventoryDetailFullRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	// acc, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "LoanClass", 0, "Current")
	repo, _ := testQueriesAccount.GetInventoryRepositorybyUuid(context.Background(), util.ToUUID("04b80e91-a1ef-4b3b-abe5-f40c158d1c6e"))
	accInv, _ := testQueriesAccount.GetAccountInventorybyUuid(context.Background(), util.ToUUID("de5e9bff-4fa4-4470-92ca-d9776268230c"))
	measure, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "UnitMeasure", 0, "Milliliter")

	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")

	arg := InventoryDetailFullRequest{
		// AccountId:     util.RandomInt(1, 100),
		Uuid:               uuid.MustParse("24e733fa-1119-4096-a434-37b9728521c5"),
		AccountInventoryId: accInv.Id,
		RepositoryId:       util.SetNullInt64(repo.Id),
		InventoryItem:      randomInventoryItem(),
		SupplierId:         util.SetNullInt64(ii.Id),
		UnitPrice:          util.RandomMoney(),
		BookValue:          util.RandomMoney(),
		Unit:               util.RandomMoney(),
		MeasureId:          measure.Id,
		BatchNumber:        sql.NullString{String: util.RandomString(30), Valid: true},
		DateManufactured:   sql.NullTime{Time: util.RandomDate(), Valid: true},
		DateExpired:        sql.NullTime{Time: util.RandomDate(), Valid: true},
		Remarks:            util.RandomString(10),
		OtherInfo:          sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}

	return arg
}

func createTestInventoryDetail(
	t *testing.T,
	d1 InventoryDetailRequest) model.InventoryDetail {
	// store := NewStoreAccount(testDB)

	getData1, err := testQueriesAccount.CreateInventoryDetail(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountInventoryId, getData1.AccountInventoryId)
	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SupplierId, getData1.SupplierId)
	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
	require.Equal(t, d1.BookValue.String(), getData1.BookValue.String())
	require.Equal(t, d1.Unit.String(), getData1.Unit.String())
	require.Equal(t, d1.MeasureId, getData1.MeasureId)
	require.Equal(t, d1.BatchNumber, getData1.BatchNumber)
	require.Equal(t, d1.DateManufactured.Time.Format("2006-01-02"), getData1.DateManufactured.Time.Format("2006-01-02"))
	require.Equal(t, d1.BatchNumber, getData1.BatchNumber)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func createTestInventoryDetailFull(
	t *testing.T,
	d1 InventoryDetailFullRequest) model.InventoryDetail {
	// store := NewStoreAccount(testDB)

	getData1, err := testQueriesAccount.CreateInventoryDetailFull(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountInventoryId, getData1.AccountInventoryId)

	InvItem, err := testQueriesAccount.GetInventoryItem(context.Background(), getData1.InventoryItemId)
	require.NoError(t, err)

	require.Equal(t, d1.InventoryItem.BarCode, InvItem.BarCode)
	require.Equal(t, d1.InventoryItem.ItemName, InvItem.ItemName)
	require.Equal(t, d1.InventoryItem.UniqueVariation, InvItem.UniqueVariation)
	require.Equal(t, d1.InventoryItem.ParentId.Int64, InvItem.ParentId.Int64)
	require.Equal(t, d1.InventoryItem.GenericNameId.Int64, InvItem.GenericNameId.Int64)
	require.Equal(t, d1.InventoryItem.BrandNameId.Int64, InvItem.BrandNameId.Int64)
	require.Equal(t, d1.InventoryItem.MeasureId, InvItem.MeasureId)
	require.Equal(t, d1.InventoryItem.Remarks, InvItem.Remarks)
	require.JSONEq(t, d1.InventoryItem.OtherInfo.String, InvItem.OtherInfo.String)

	require.Equal(t, d1.SupplierId, getData1.SupplierId)
	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
	require.Equal(t, d1.BookValue.String(), getData1.BookValue.String())
	require.Equal(t, d1.Unit.String(), getData1.Unit.String())
	require.Equal(t, d1.MeasureId, getData1.MeasureId)
	require.Equal(t, d1.BatchNumber, getData1.BatchNumber)
	require.Equal(t, d1.DateManufactured.Time.Format("2006-01-02"), getData1.DateManufactured.Time.Format("2006-01-02"))
	require.Equal(t, d1.BatchNumber, getData1.BatchNumber)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestInventoryDetail(
	t *testing.T,
	d1 InventoryDetailRequest) model.InventoryDetail {

	getData1, err := testQueriesAccount.UpdateInventoryDetail(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountInventoryId, getData1.AccountInventoryId)
	require.Equal(t, d1.InventoryItemId, getData1.InventoryItemId)
	require.Equal(t, d1.SupplierId, getData1.SupplierId)
	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
	require.Equal(t, d1.BookValue.String(), getData1.BookValue.String())
	require.Equal(t, d1.Unit.String(), getData1.Unit.String())
	require.Equal(t, d1.MeasureId, getData1.MeasureId)
	require.Equal(t, d1.BatchNumber, getData1.BatchNumber)
	require.Equal(t, d1.DateManufactured.Time.Format("2006-01-02"), getData1.DateManufactured.Time.Format("2006-01-02"))
	require.Equal(t, d1.BatchNumber, getData1.BatchNumber)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteInventoryDetail(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteInventoryDetail(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetInventoryDetail(context.Background(), id)
	require.Error(t, err)
	// require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
