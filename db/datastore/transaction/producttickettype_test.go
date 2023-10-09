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

func TestProductTicketType(t *testing.T) {

	// Test Data
	d1 := RandomProductTicketType()
	d2 := RandomProductTicketType()
	d2.Uuid = util.ToUUID("324b3080-895a-4064-bdee-d6f6077732f2")
	ticTyp, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SavingStatus", 0, "New Account")
	d2.TicketTypeId = ticTyp.Id

	// Test Create
	CreatedD1 := createTestProductTicketType(t, d1)
	CreatedD2 := createTestProductTicketType(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetProductTicketType(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.TicketTypeId, getData1.TicketTypeId)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetProductTicketType(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Uuid, getData2.Uuid)
	require.Equal(t, d2.CentralOfficeId, getData2.CentralOfficeId)
	require.Equal(t, d2.ProductId, getData2.ProductId)
	require.Equal(t, d2.TicketTypeId, getData2.TicketTypeId)
	require.Equal(t, d2.StatusId, getData2.StatusId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetProductTicketTypebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestProductTicketType(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Uuid, updatedD1.Uuid)
	require.Equal(t, updateD2.CentralOfficeId, updatedD1.CentralOfficeId)
	require.Equal(t, updateD2.ProductId, updatedD1.ProductId)
	require.Equal(t, updateD2.TicketTypeId, updatedD1.TicketTypeId)
	require.Equal(t, updateD2.StatusId, updatedD1.StatusId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListProductTicketType(t, ListProductTicketTypeParams{
		ProductId: updatedD1.ProductId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteProductTicketType(t, getData1.Uuid)
	testDeleteProductTicketType(t, getData2.Uuid)
}

func testListProductTicketType(t *testing.T, arg ListProductTicketTypeParams) {

	ProductTicketType, err := testQueriesTransaction.ListProductTicketType(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", ProductTicketType)
	require.NotEmpty(t, ProductTicketType)

}

func RandomProductTicketType() ProductTicketTypeRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")
	prod, _ := testQueriesAccount.GetProductbyName(context.Background(), "Loan")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ticketstatus", 0, "Completed")
	ticTyp, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "ActionList", 0, "Encoded")

	arg := ProductTicketTypeRequest{
		Uuid:            util.ToUUID("cc2223c9-2f6f-408b-8740-28151190c98a"),
		CentralOfficeId: ofc.Id,
		ProductId:       prod.Id,
		TicketTypeId:    ticTyp.Id,
		StatusId:        stat.Id,
		OtherInfo:       sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestProductTicketType(
	t *testing.T,
	d1 ProductTicketTypeRequest) model.ProductTicketType {

	getData1, err := testQueriesTransaction.CreateProductTicketType(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.TicketTypeId, getData1.TicketTypeId)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestProductTicketType(
	t *testing.T,
	d1 ProductTicketTypeRequest) model.ProductTicketType {

	getData1, err := testQueriesTransaction.UpdateProductTicketType(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Uuid, getData1.Uuid)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.ProductId, getData1.ProductId)
	require.Equal(t, d1.TicketTypeId, getData1.TicketTypeId)
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteProductTicketType(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteProductTicketType(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetProductTicketTypebyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
