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

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

// var store StoreAccount

//	func init() {
//		store = NewStoreAccount(testDB)
//	}
func TestInventoryItem(t *testing.T) {

	// Test Data

	// store := NewStoreAccount(testDB)

	d1 := randomInventoryItem()
	d1.Uuid = uuid.MustParse("0f5cc4a6-0969-4352-b536-0ff54a289e63")
	d1.ItemCode = "Test01"

	// gen, _ := testQueriesAccount.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "Weight")
	// d1.GenericNameId = sql.NullInt64(sql.NullInt64{Int64: gen.Id, Valid: true})

	genWeight, _ := testQueriesAccount.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "Weight")
	genManufactured, _ := testQueriesAccount.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "DateManufactured")
	genCores, _ := testQueriesAccount.GetReferenceInfobyTitle(context.Background(), "InventorySpecs", 0, "Number of Cores")

	d1.SpecsNumberList = append(d1.SpecsNumberList, InventorySpecsNumberRequest{
		SpecsId: genWeight.Id,
		Value:   decimal.NewFromInt32(1),
		Value2:  decimal.NewFromInt32(11),
	})

	d1.SpecsDateList = append(d1.SpecsDateList, InventorySpecsDateRequest{
		SpecsId: genManufactured.Id,
		Value:   util.RandomDate(),
		Value2:  util.RandomDate(),
	})

	d1.SpecsStringList = append(d1.SpecsStringList, InventorySpecsStringRequest{
		SpecsId: genCores.Id,
		Value:   util.RandomString(4),
	})

	fmt.Printf("Get by d1%+v\n", d1.SpecsNumberList)

	d2 := randomInventoryItem()
	d2.Uuid = uuid.MustParse("d57a2efd-abe1-42c8-8be9-f7659d06e30e")
	d2.ItemCode = "Test02"
	// log.Println("Specs", d1.SpecsNumberList[0], d1.Id, gen.Id)

	// fmt.Printf("Get by UUId%+v\n", d1)
	// Test Create
	CreatedD1 := createTestInventoryItem(t, d1)
	CreatedD2 := createTestInventoryItem(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetInventoryItem(context.Background(), CreatedD1.Id)
	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.ItemCode, getData1.ItemCode)
	require.Equal(t, d1.BarCode, getData1.BarCode)
	require.Equal(t, d1.ItemName, getData1.ItemName)
	require.Equal(t, d1.UniqueVariation, getData1.UniqueVariation)
	require.Equal(t, d1.ParentId.Int64, getData1.ParentId.Int64)
	require.Equal(t, d1.GenericNameId.Int64, getData1.GenericNameId.Int64)
	require.Equal(t, d1.BrandNameId.Int64, getData1.BrandNameId.Int64)
	require.Equal(t, d1.MeasureId, getData1.MeasureId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetInventoryItem(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData1)
	require.Equal(t, d2.ItemCode, getData2.ItemCode)
	require.Equal(t, d2.BarCode, getData2.BarCode)
	require.Equal(t, d2.ItemName, getData2.ItemName)
	require.Equal(t, d2.UniqueVariation, getData2.UniqueVariation)
	require.Equal(t, d2.ParentId.Int64, getData2.ParentId.Int64)
	require.Equal(t, d2.GenericNameId.Int64, getData2.GenericNameId.Int64)
	require.Equal(t, d2.BrandNameId.Int64, getData2.BrandNameId.Int64)
	require.Equal(t, d2.MeasureId, getData2.MeasureId)
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetInventoryItembyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	// fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestInventoryItem(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, updatedD1.ItemCode, updateD2.ItemCode)
	require.Equal(t, updatedD1.BarCode, updateD2.BarCode)
	require.Equal(t, updatedD1.ItemName, updateD2.ItemName)
	require.Equal(t, updatedD1.UniqueVariation, updateD2.UniqueVariation)
	require.Equal(t, updatedD1.ParentId.Int64, updateD2.ParentId.Int64)
	require.Equal(t, updatedD1.GenericNameId.Int64, updateD2.GenericNameId.Int64)
	require.Equal(t, updatedD1.BrandNameId.Int64, updateD2.BrandNameId.Int64)
	require.Equal(t, updatedD1.MeasureId, updateD2.MeasureId)
	require.Equal(t, updatedD1.Remarks, updateD2.Remarks)
	require.JSONEq(t, updatedD1.OtherInfo.String, updateD2.OtherInfo.String)

	testListInventoryItembyBrand(t, ListInventoryItembyBrandParams{
		BrandId: updatedD1.BrandNameId.Int64,
		Limit:   5,
		Offset:  0,
	})

	testListInventoryItembyGeneric(t, ListInventoryItembyGenericParams{
		GenericId: updatedD1.GenericNameId.Int64,
		Limit:     5,
		Offset:    0,
	})

	arg3 := InventoryItemFilterParams{
		Filter: getData2.ItemName,
		Limit:  2,
		Offset: 0,
	}

	inventoryItem, err := testQueriesAccount.InventoryItemFilter(context.Background(), arg3)
	require.NoError(t, err)

	for _, ref := range inventoryItem {
		require.NotEmpty(t, ref)
		// require.NotEmpty(t, ref)
		// require.Equal(t, lastAccount.Owner, account.Owner)
	}

	// Delete Data
	//testDeleteInventoryItem(t, getData1.Id)
	testDeleteInventoryItem(t, getData2.Id)

	arg4 := InventoryItemSearchParams{
		SearchSpecsString: []SearchSpecsString{
			{
				ItemId:    d1.SpecsStringList[0].SpecsId,
				Condition: Contains,
				Value:     d1.SpecsStringList[0].Value,
				Weight:    .50,
			},
			{
				ItemId:    d1.SpecsStringList[0].SpecsId,
				Condition: Equal,
				Value:     d1.SpecsStringList[0].Value,
				Weight:    .50,
			},
		},
		SearchSpecsNumber: []SearchSpecsNumber{
			{
				ItemId:    d1.SpecsNumberList[0].SpecsId,
				Condition: Between,
				Value1:    d1.SpecsNumberList[0].Value,
				Value2:    d1.SpecsNumberList[0].Value2,
				Weight:    .10,
			},
			{
				ItemId:    d1.SpecsNumberList[0].SpecsId,
				Condition: Between,
				Value1:    d1.SpecsNumberList[0].Value,
				ValueUsed: Both,
				Weight:    .10,
			},
		},
		SearchSpecsDate: []SearchSpecsDate{
			{
				ItemId:    d1.SpecsDateList[0].SpecsId,
				Condition: Equal,
				Value1:    d1.SpecsDateList[0].Value,
				Weight:    .10,
			},
			{
				ItemId:    d1.SpecsDateList[0].SpecsId,
				Condition: Equal,
				Value1:    d1.SpecsDateList[0].Value,
				ValueUsed: Both,
				Weight:    .10,
			},
			{
				ItemId:    d1.SpecsDateList[0].SpecsId,
				Condition: Equal,
				Value1:    d1.SpecsDateList[0].Value2,
				ValueUsed: Value2,
				Weight:    .10,
			},
		},
		Limit:  2,
		Offset: 0,
	}

	searchList, er := testQueriesAccount.InventoryItemSearch(context.Background(), arg4)
	require.NoError(t, er)
	require.NotEmpty(t, searchList)
	fmt.Printf("%+v", searchList)

	// require.True(t, false)
}

func testListInventoryItembyBrand(t *testing.T, arg ListInventoryItembyBrandParams) {
	// store := NewStoreAccount(testDB)
	log.Printf("ListInventoryItembyBrandParams: %+v", arg)
	InventoryItem, err := testQueriesAccount.ListInventoryItembyBrand(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", InventoryItem)
	require.NotEmpty(t, InventoryItem)
}

func testListInventoryItembyGeneric(t *testing.T, arg ListInventoryItembyGenericParams) {
	// store := NewStoreAccount(testDB)
	log.Printf("ListInventoryItembyGenericParams: %+v", arg)
	InventoryItem, err := testQueriesAccount.ListInventoryItembyGeneric(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", InventoryItem)
	require.NotEmpty(t, InventoryItem)
}

func randomInventoryItem() InventoryItemRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	// acc, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "LoanClass", 0, "Current")
	gen, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "GenericName", 0, "Tooth Paste")
	brand, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "BrandName", 0, "Colgate")
	measure, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "UnitMeasure", 0, "Milliliter")

	log.Printf("randomInventoryItem: %+v", gen)
	arg := InventoryItemRequest{
		// AccountId:     util.RandomInt(1, 100),
		BarCode:         sql.NullString{String: "", Valid: false},
		ItemName:        util.RandomString(48),
		UniqueVariation: util.RandomString(48),
		ParentId:        sql.NullInt64(sql.NullInt64{Int64: 0, Valid: false}),
		GenericNameId:   sql.NullInt64(sql.NullInt64{Int64: gen.Id, Valid: true}),
		BrandNameId:     sql.NullInt64(sql.NullInt64{Int64: brand.Id, Valid: true}),
		MeasureId:       measure.Id,
		Remarks:         util.RandomString(10),
		OtherInfo:       sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestInventoryItem(
	t *testing.T,
	d1 InventoryItemRequest) model.InventoryItem {
	// store := NewStoreAccount(testDB)

	getData1, err := testQueriesAccount.CreateInventoryItemFull(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ItemCode, getData1.ItemCode)
	require.Equal(t, d1.BarCode, getData1.BarCode)
	require.Equal(t, d1.ItemName, getData1.ItemName)
	require.Equal(t, d1.UniqueVariation, getData1.UniqueVariation)
	require.Equal(t, d1.ParentId.Int64, getData1.ParentId.Int64)
	require.Equal(t, d1.GenericNameId.Int64, getData1.GenericNameId.Int64)
	require.Equal(t, d1.BrandNameId.Int64, getData1.BrandNameId.Int64)
	require.Equal(t, d1.MeasureId, getData1.MeasureId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	invSpecsNumber, errn := testQueriesAccount.ListInventorySpecsNumber(context.Background(),
		ListInventorySpecsNumberParams{
			InventoryItemId: getData1.Id,
			Limit:           5,
			Offset:          0,
		})
	require.NoError(t, errn)
	// fmt.Printf("Get by UUId%+v\n", invSpecsDate)

	invSpecsDate, errd := testQueriesAccount.ListInventorySpecsDate(context.Background(),
		ListInventorySpecsDateParams{
			InventoryItemId: getData1.Id,
			Limit:           5,
			Offset:          0,
		})
	require.NoError(t, errd)

	invSpecsString, errs := testQueriesAccount.ListInventorySpecsString(context.Background(),
		ListInventorySpecsStringParams{
			InventoryItemId: getData1.Id,
			Limit:           5,
			Offset:          0,
		})
	require.NoError(t, errs)

	for i, specs := range d1.SpecsNumberList {
		require.Equal(t, getData1.Id, invSpecsNumber[i].InventoryItemId)
		require.Equal(t, specs.SpecsId, invSpecsNumber[i].SpecsId)
		require.Equal(t, specs.Value.String(), invSpecsNumber[i].Value.String())
		require.Equal(t, specs.Value2.String(), invSpecsNumber[i].Value2.String())
	}

	require.NoError(t, errd)
	for i, specs := range d1.SpecsDateList {
		require.Equal(t, getData1.Id, invSpecsDate[i].InventoryItemId)
		require.Equal(t, specs.SpecsId, invSpecsDate[i].SpecsId)
		require.Equal(t, specs.Value.Format("2006-01-02"), invSpecsDate[i].Value.Format("2006-01-02"))
		require.Equal(t, specs.Value2.Format("2006-01-02"), invSpecsDate[i].Value2.Format("2006-01-02"))
	}

	require.NoError(t, errs)
	for i, specs := range d1.SpecsStringList {
		require.Equal(t, getData1.Id, invSpecsString[i].InventoryItemId)
		require.Equal(t, specs.SpecsId, invSpecsString[i].SpecsId)
		require.Equal(t, specs.Value, invSpecsString[i].Value)
	}

	return getData1
}

func updateTestInventoryItem(
	t *testing.T,
	d1 InventoryItemRequest) model.InventoryItem {

	getData1, err := testQueriesAccount.UpdateInventoryItem(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.ItemCode, getData1.ItemCode)
	require.Equal(t, d1.BarCode, getData1.BarCode)
	require.Equal(t, d1.ItemName, getData1.ItemName)
	require.Equal(t, d1.UniqueVariation, getData1.UniqueVariation)
	require.Equal(t, d1.ParentId.Int64, getData1.ParentId.Int64)
	require.Equal(t, d1.GenericNameId.Int64, getData1.GenericNameId.Int64)
	require.Equal(t, d1.BrandNameId.Int64, getData1.BrandNameId.Int64)
	require.Equal(t, d1.MeasureId, getData1.MeasureId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteInventoryItem(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteInventoryItem(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetInventoryItem(context.Background(), id)
	require.Error(t, err)
	// require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
