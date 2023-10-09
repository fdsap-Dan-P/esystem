package db

import (
	"context"
	"log"
	"simplebank/util"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomer(t *testing.T) {
	var inaiiid int64 = 143

	// Get SearchCustomerCID
	cust, err := testQueriesKPlus.SearchCustomerCID(context.Background(), inaiiid)
	require.NoError(t, err)
	require.NotEmpty(t, cust)
	require.Equal(t, inaiiid, util.ToInt64(cust.INAIIID))

	custs, er := testQueriesKPlus.GetCustomersInfo(context.Background(), CustomersInfoParam{
		// INAIIID:    "143",
		SearchName: "Macaraig  Teresita G",
	})
	require.NoError(t, er)
	require.NotEmpty(t, custs)
	require.Equal(t, inaiiid, util.ToInt64(cust.INAIIID))

}

func TestSavings(t *testing.T) {
	var inaiiid int64 = 1435254

	// Get SearchCustomerCID
	dt, err := testQueriesKPlus.SavingsList(context.Background(), SavingsListParams{
		INAIIID: inaiiid,
		Limit:   5,
		Offset:  0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, dt)
	require.Equal(t, inaiiid, util.ToInt64(dt[0].INAIIID))
}

func TestTransaction(t *testing.T) {
	trnParams := TransactionParams{
		Acc:      "E30304-1012-3911659",
		DateFrom: util.SetDate("2019-09-01"),
		DateTo:   util.SetDate("2022-09-30"),
		Limit:    20,
		Offset:   0,
	}

	// Get SearchCustomerCID
	dt, err := testQueriesKPlus.Transaction(context.Background(), trnParams)

	for i, d := range dt {
		// log.Printf("TestTransaction: %v %v %v %v %v %v %v %v", i, d.TrnDate, d.Trn, d.Prin, d.Intr, d.BalPrin, d.BalInt, d.Balance, d.PaidPrin)
		log.Printf("TestTransaction: %v %v %v %v %v %v %v, %v", i, d.TrnDate, d.Trn, d.Prin, d.BalPrin, d.Balance, d.PaidPrin, d.NormalBalance)
	}
	require.NoError(t, err)
	require.NotEmpty(t, dt)
	require.Equal(t, trnParams.Acc, dt[0].Acc)
	// require.NotEqual(t, trnParams.Acc, dt[0].Acc)
}

func TestColSht(t *testing.T) {
	var inaiiid int64 = 143

	// Get SearchCustomerCID
	dt, err := testQueriesKPlus.ColSht(context.Background(), ColShtParams{
		INAIIID: inaiiid,
		Limit:   5,
		Offset:  0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, dt)
	require.Equal(t, inaiiid, util.ToInt64(dt[0].INAIIID))
}

func TestCallBackRef(t *testing.T) {
	var prNo string = "165312"

	// Get SearchCustomerCID
	dt, err := testQueriesKPlus.CallBackRef(context.Background(), prNo)
	require.NoError(t, err)
	require.NotEmpty(t, dt)
	require.Equal(t, "Transaction Exist!", dt.Message)

	dt, err = testQueriesKPlus.CallBackRef(context.Background(), "notexists")
	require.NoError(t, err)
	require.NotEmpty(t, dt)
	require.Equal(t, "Transaction not Exist!", dt.Message)
}

// E30101-4041-0528625
// E301C4-4001-004896338
// E302U8-4001-1149967
// E302U8-4001-1150165

func TestMultiplePayment(t *testing.T) {
	multiPay := MultiplePaymentRequest{
		RemitterINAIIID: "1435254",
		// CustomerId:nil,
		// IIID:nil,
		PrNumber:        "test-10001",
		SourceId:        1,
		OrNumber:        0,
		Username:        "konek2CARD",
		Trndate:         "2021-01-01",
		TotalCollection: 100,
		Particulars:     "",
		Payment: []Payment{
			{
				Acc:      "E30101-4041-0528625",
				Pay:      util.SetDecimal("2021.00"),
				Withdraw: util.SetDecimal("0"),
				AppType:  0,
				// Uuid:     "df0a5e20-b71a-40a8-a6d1-15c56267e27c",
			},
			// {
			// 	Acc:      "E302U8-4001-1149967",
			// 	Pay:      util.SetDecimal("11"),
			// 	Withdraw: util.SetDecimal("0"),
			// 	Type:     "1",
			// 	Uuid:     "df0a5e20-b71a-40a8-a6d1-15c56267e27c",
			// },
			// {
			// 	Acc:      "E301C4-4001-004896338",
			// 	Pay:      util.SetDecimal("12"),
			// 	Withdraw: util.SetDecimal("0"),
			// 	Type:     "1",
			// 	Uuid:     "df0a5e20-b71a-40a8-a6d1-15c56267e27c",
			// },
			// {
			// 	Acc:      "E302U8-4001-1150165",
			// 	Pay:      util.SetDecimal("13"),
			// 	Withdraw: util.SetDecimal("0"),
			// 	Type:     "1",
			// 	Uuid:     "df0a5e20-b71a-40a8-a6d1-15c56267e27c",
			// },
		},
	}

	// Get SearchCustomerCID
	dt, err := testQueriesKPlus.MultiplePayment(context.Background(), multiPay)
	require.Error(t, err)
	require.NotEmpty(t, dt)
	require.Equal(t, "Transaction Exist!", dt.Message)
}
