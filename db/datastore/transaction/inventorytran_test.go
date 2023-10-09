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

	dsAcc "simplebank/db/datastore/account"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestInventoryTran(t *testing.T) {

	// Test Data
	d1 := randomInventoryTran()
	// accQtl, _ := testQueriesAccount.GetAccountInventorybyUuid(context.Background(), uuid.MustParse("c3476afe-bd50-49e6-8de3-074555a8e1bd"))
	// d1.InventoryDetailId = accQtl.Id

	log.Printf("invDet1 %v:", d1)

	d2 := randomInventoryTran()
	// accQtl, _ = testQueriesAccount.GetAccountInventorybyUuid(context.Background(), uuid.MustParse("b35e39e8-885b-41a6-a070-a249c2a099e5"))
	// d2.InventoryDetailId = accQtl.Id
	d2.Uuid = uuid.MustParse("dda68542-d05e-4301-8d64-5f921bc88c2c")
	d2.Series = 2
	log.Printf("invDet2 %v:", d2)

	// Test Create
	CreatedD1 := createTestInventoryTran(t, d1)
	CreatedD2 := createTestInventoryTran(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetInventoryTranbyUuid(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.InventoryDetailId, getData1.InventoryDetailId)
	require.Equal(t, d1.RepositoryId, getData1.RepositoryId)
	require.Equal(t, d1.Quantity.String(), getData1.Quantity.String())
	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
	require.Equal(t, d1.Discount.String(), getData1.Discount.String())
	require.Equal(t, d1.TaxAmt.String(), getData1.TaxAmt.String())
	require.Equal(t, d1.NetTrnAmt.String(), getData1.NetTrnAmt.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetInventoryTranbyUuid(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TrnHeadId, getData2.TrnHeadId)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.InventoryDetailId, getData2.InventoryDetailId)
	require.Equal(t, d2.RepositoryId, getData2.RepositoryId)
	require.Equal(t, d2.Quantity.String(), getData2.Quantity.String())
	require.Equal(t, d2.UnitPrice.String(), getData2.UnitPrice.String())
	require.Equal(t, d2.Discount.String(), getData2.Discount.String())
	require.Equal(t, d2.TaxAmt.String(), getData2.TaxAmt.String())
	require.Equal(t, d2.NetTrnAmt.String(), getData2.NetTrnAmt.String())
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetInventoryTranbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUid%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestInventoryTran(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TrnHeadId, updatedD1.TrnHeadId)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.InventoryDetailId, updatedD1.InventoryDetailId)
	require.Equal(t, updateD2.RepositoryId, updatedD1.RepositoryId)
	require.Equal(t, updateD2.Quantity.String(), updatedD1.Quantity.String())
	require.Equal(t, updateD2.UnitPrice.String(), updatedD1.UnitPrice.String())
	require.Equal(t, updateD2.Discount.String(), updatedD1.Discount.String())
	require.Equal(t, updateD2.TaxAmt.String(), updatedD1.TaxAmt.String())
	require.Equal(t, updateD2.NetTrnAmt.String(), updatedD1.NetTrnAmt.String())
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListInventoryTran(t, ListInventoryTranParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteInventoryTran(t, CreatedD1.Uuid)
	testDeleteInventoryTran(t, CreatedD2.Uuid)
}

func testListInventoryTran(t *testing.T, arg ListInventoryTranParams) {

	InventoryTran, err := testQueriesTransaction.ListInventoryTran(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", InventoryTran)
	require.NotEmpty(t, InventoryTran)

}

func randomInventoryTran() InventoryTranRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	trn, _ := testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("2af90d74-3bee-48c5-8935-443edafb8f5a"))

	repo, _ := testQueriesAccount.CreateInventoryRepository(context.Background(), randomInventoryRepository())
	invDet, _ := testQueriesAccount.CreateInventoryDetail(context.Background(), randomInventoryDetail())

	arg := InventoryTranRequest{
		Uuid:              uuid.MustParse("1940daf6-0203-4837-ba92-04ced24648fd"),
		TrnHeadId:         trn.Id,
		Series:            1,
		InventoryDetailId: invDet.Id,
		RepositoryId:      repo.Id,
		Quantity:          util.RandomMoney(),
		UnitPrice:         util.RandomMoney(),
		Discount:          util.RandomMoney(),
		TaxAmt:            util.RandomMoney(),
		NetTrnAmt:         util.RandomMoney(),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}

	log.Printf("arg %v:", arg)

	return arg
}

func createTestInventoryTran(
	t *testing.T,
	d1 InventoryTranRequest) model.InventoryTran {

	getData1, err := testQueriesTransaction.CreateInventoryTran(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.InventoryDetailId, getData1.InventoryDetailId)
	require.Equal(t, d1.RepositoryId, getData1.RepositoryId)
	require.Equal(t, d1.Quantity.String(), getData1.Quantity.String())
	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
	require.Equal(t, d1.Discount.String(), getData1.Discount.String())
	require.Equal(t, d1.TaxAmt.String(), getData1.TaxAmt.String())
	require.Equal(t, d1.NetTrnAmt.String(), getData1.NetTrnAmt.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestInventoryTran(
	t *testing.T,
	d1 InventoryTranRequest) model.InventoryTran {

	getData1, err := testQueriesTransaction.UpdateInventoryTran(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.InventoryDetailId, getData1.InventoryDetailId)
	require.Equal(t, d1.RepositoryId, getData1.RepositoryId)
	require.Equal(t, d1.Quantity.String(), getData1.Quantity.String())
	require.Equal(t, d1.UnitPrice.String(), getData1.UnitPrice.String())
	require.Equal(t, d1.Discount.String(), getData1.Discount.String())
	require.Equal(t, d1.TaxAmt.String(), getData1.TaxAmt.String())
	require.Equal(t, d1.NetTrnAmt.String(), getData1.NetTrnAmt.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteInventoryTran(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteInventoryTran(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetInventoryTranbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func randomInventoryDetail() dsAcc.InventoryDetailRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	meas, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "UnitMeasure", 0, "Item")
	acc, _ := testQueriesAccount.CreateAccountInventory(context.Background(), randomAccountInventory("1001-0001-0000001"))
	inv, _ := testQueriesAccount.CreateInventoryItem(context.Background(), randomInventoryItem())
	repo, _ := testQueriesAccount.CreateInventoryRepository(context.Background(), randomInventoryRepository())
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")

	arg := dsAcc.InventoryDetailRequest{
		Uuid:               uuid.MustParse("24e733fa-1119-4096-a434-37b9728521c5"),
		AccountInventoryId: acc.Id,
		InventoryItemId:    inv.Id,
		RepositoryId:       util.SetNullInt64(repo.Id),
		SupplierId:         util.SetNullInt64(ii.Id),
		UnitPrice:          util.RandomMoney(),
		BookValue:          util.RandomMoney(),
		Unit:               util.RandomMoney(),
		MeasureId:          meas.Id,
		BatchNumber:        util.SetNullString("BatchNum"),
		DateManufactured:   util.RandomNullDate(),
		DateExpired:        util.RandomNullDate(),
		Remarks:            "String",
		OtherInfo:          sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}
