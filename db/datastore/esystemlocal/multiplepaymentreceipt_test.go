package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"simplebank/util"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiplePaymentReceipt(t *testing.T) {

	// Test Data
	d1 := randomMultiplePaymentReceipt()
	d2 := randomMultiplePaymentReceipt()
	d2.OrNo = d2.OrNo + 1

	err := createTestMultiplePaymentReceipt(t, d1)
	require.NoError(t, err)

	err = createTestMultiplePaymentReceipt(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetMultiplePaymentReceipt(context.Background(), d1.OrNo)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.OrNo, getData1.OrNo)
	require.Equal(t, d1.CID, getData1.CID)
	require.Equal(t, d1.PrNo, getData1.PrNo)
	require.Equal(t, d1.UserName, getData1.UserName)
	require.Equal(t, d1.TermId, getData1.TermId)

	getData2, err2 := testQueriesLocal.GetMultiplePaymentReceipt(context.Background(), d2.OrNo)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.OrNo, getData2.OrNo)
	require.Equal(t, d2.CID, getData2.CID)
	require.Equal(t, d2.PrNo, getData2.PrNo)
	require.Equal(t, d2.UserName, getData2.UserName)
	require.Equal(t, d2.TermId, getData2.TermId)

	// Update Data
	updateD2 := d2
	updateD2.TermId = "..edited"
	// updateD2.Location = updateD2.Location + "Edited"

	err3 := updateTestMultiplePaymentReceipt(t, updateD2)
	require.NoError(t, err3)

	log.Printf("test...: %v", updateD2)
	getData1, err1 = testQueriesLocal.GetMultiplePaymentReceipt(context.Background(), updateD2.OrNo)
	require.NoError(t, err1)

	require.Equal(t, updateD2.OrNo, getData1.OrNo)
	require.Equal(t, updateD2.CID, getData1.CID)
	require.Equal(t, updateD2.PrNo, getData1.PrNo)
	require.Equal(t, updateD2.UserName, getData1.UserName)
	require.Equal(t, updateD2.TermId, getData1.TermId)

	testListMultiplePaymentReceipt(t, ListMultiplePaymentReceiptParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteMultiplePaymentReceipt(t, d1.OrNo)
	testDeleteMultiplePaymentReceipt(t, d2.OrNo)
}

func testListMultiplePaymentReceipt(t *testing.T, arg ListMultiplePaymentReceiptParams) {

	MultiplePaymentReceipt, err := testQueriesLocal.ListMultiplePaymentReceipt(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", MultiplePaymentReceipt)
	require.NotEmpty(t, MultiplePaymentReceipt)

}

func randomMultiplePaymentReceipt() MultiplePaymentReceiptRequest {

	arg := MultiplePaymentReceiptRequest{
		TrnDate:  util.SetDate("2022-01-01"),
		OrNo:     529000224,
		CID:      400001,
		UserName: "sa",
		TermId:   "testServer",
		AmtPaid:  util.SetDecimal("10"),
	}
	return arg
}

func createTestMultiplePaymentReceipt(
	t *testing.T,
	req MultiplePaymentReceiptRequest) error {

	err1 := testQueriesLocal.CreateMultiplePaymentReceipt(context.Background(), req)
	fmt.Printf("Get by createTestMultiplePaymentReceipt%+v\n", req)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetMultiplePaymentReceipt(context.Background(), req.OrNo)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.OrNo, getData.OrNo)
	require.Equal(t, req.CID, getData.CID)
	require.Equal(t, req.PrNo, getData.PrNo)
	require.Equal(t, req.UserName, getData.UserName)
	require.Equal(t, req.TermId, getData.TermId)
	return err2
}

func updateTestMultiplePaymentReceipt(
	t *testing.T,
	d1 MultiplePaymentReceiptRequest) error {

	err := testQueriesLocal.UpdateMultiplePaymentReceipt(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteMultiplePaymentReceipt(t *testing.T, orno int64) {
	err := testQueriesLocal.DeleteMultiplePaymentReceipt(context.Background(), orno)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetMultiplePaymentReceipt(context.Background(), orno)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
