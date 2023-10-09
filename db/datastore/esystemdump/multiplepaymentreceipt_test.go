package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"
	"simplebank/util"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiplePaymentReceipt(t *testing.T) {

	// Test Data
	d1 := randomMultiplePaymentReceipt()
	d2 := randomMultiplePaymentReceipt()

	err := createTestMultiplePaymentReceipt(t, d1)
	require.NoError(t, err)

	err = createTestMultiplePaymentReceipt(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetMultiplePaymentReceipt(context.Background(), d1.BrCode, d1.OrNo)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.BrCode, getData1.BrCode)
	require.Equal(t, d1.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d1.OrNo, getData1.OrNo)
	require.Equal(t, d1.CID, getData1.CID)
	require.Equal(t, d1.PrNo, getData1.PrNo)
	require.Equal(t, d1.UserName, getData1.UserName)
	require.Equal(t, d1.TermId, getData1.TermId)
	require.Equal(t, d1.AmtPaid, getData1.AmtPaid)

	getData2, err2 := testQueriesDump.GetMultiplePaymentReceipt(context.Background(), d2.BrCode, d2.OrNo)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.BrCode, getData2.BrCode)
	require.Equal(t, d2.TrnDate.Format(`2006-01-02`), getData2.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d2.OrNo, getData2.OrNo)
	require.Equal(t, d2.CID, getData2.CID)
	require.Equal(t, d2.PrNo, getData2.PrNo)
	require.Equal(t, d2.UserName, getData2.UserName)
	require.Equal(t, d2.TermId, getData2.TermId)
	require.Equal(t, d2.AmtPaid, getData2.AmtPaid)
	// Update Data
	updateD2 := d2
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestMultiplePaymentReceipt(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetMultiplePaymentReceipt(context.Background(), updateD2.BrCode, updateD2.OrNo)
	require.NoError(t, err1)

	require.Equal(t, updateD2.BrCode, getData1.BrCode)
	require.Equal(t, updateD2.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, updateD2.OrNo, getData1.OrNo)
	require.Equal(t, updateD2.CID, getData1.CID)
	require.Equal(t, updateD2.PrNo, getData1.PrNo)
	require.Equal(t, updateD2.UserName, getData1.UserName)
	require.Equal(t, updateD2.TermId, getData1.TermId)
	require.Equal(t, updateD2.AmtPaid, getData1.AmtPaid)

	testListMultiplePaymentReceipt(t, ListMultiplePaymentReceiptParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteMultiplePaymentReceipt(t, d1.BrCode, d1.OrNo)
	testDeleteMultiplePaymentReceipt(t, d2.BrCode, d2.OrNo)
}

func testListMultiplePaymentReceipt(t *testing.T, arg ListMultiplePaymentReceiptParams) {

	MultiplePaymentReceipt, err := testQueriesDump.ListMultiplePaymentReceipt(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", MultiplePaymentReceipt)
	require.NotEmpty(t, MultiplePaymentReceipt)

}

func randomMultiplePaymentReceipt() model.MultiplePaymentReceipt {

	arg := model.MultiplePaymentReceipt{
		ModCtr:   1,
		BrCode:   "E3",
		TrnDate:  util.SetDate("2023-12-01"),
		OrNo:     100,
		CID:      100,
		PrNo:     0,
		UserName: "sa",
		TermId:   "",
		AmtPaid:  util.SetDecimal("100"),
	}
	return arg
}

func createTestMultiplePaymentReceipt(
	t *testing.T,
	req model.MultiplePaymentReceipt) error {

	err1 := testQueriesDump.CreateMultiplePaymentReceipt(context.Background(), req)
	// fmt.Printf("Get by createTestMultiplePaymentReceipt%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetMultiplePaymentReceipt(context.Background(), req.BrCode, req.OrNo)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.BrCode, getData.BrCode)
	require.Equal(t, req.TrnDate.Format(`2006-01-02`), getData.TrnDate.Format(`2006-01-02`))
	require.Equal(t, req.OrNo, getData.OrNo)
	require.Equal(t, req.CID, getData.CID)
	require.Equal(t, req.PrNo, getData.PrNo)
	require.Equal(t, req.UserName, getData.UserName)
	require.Equal(t, req.TermId, getData.TermId)
	require.Equal(t, req.AmtPaid, getData.AmtPaid)

	return err2
}

func updateTestMultiplePaymentReceipt(
	t *testing.T,
	d1 model.MultiplePaymentReceipt) error {

	err := testQueriesDump.UpdateMultiplePaymentReceipt(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteMultiplePaymentReceipt(t *testing.T, brCode string, orNo int64) {
	err := testQueriesDump.DeleteMultiplePaymentReceipt(context.Background(), brCode, orNo)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetMultiplePaymentReceipt(context.Background(), brCode, orNo)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
