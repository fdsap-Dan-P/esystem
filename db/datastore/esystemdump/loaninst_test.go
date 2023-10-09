package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"
	"simplebank/util"

	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestLoanInst(t *testing.T) {

	// Test Data
	d1 := randomLoanInst()
	d2 := randomLoanInst()
	d2.Dnum = d2.Dnum + 1

	err := createTestLoanInst(t, d1)
	require.NoError(t, err)

	err = createTestLoanInst(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetLoanInst(context.Background(), d1.BrCode, d1.Acc, d1.Dnum)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.Dnum, getData1.Dnum)
	require.Equal(t, d1.DueDate.Format(`2006-01-02`), getData1.DueDate.Format(`2006-01-02`))
	require.Equal(t, d1.InstFlag, getData1.InstFlag)
	require.Equal(t, d1.DuePrin.String(), getData1.DuePrin.String())
	require.Equal(t, d1.DueInt.String(), getData1.DueInt.String())
	require.Equal(t, d1.UpInt.String(), getData1.UpInt.String())

	getData2, err2 := testQueriesDump.GetLoanInst(context.Background(), d2.BrCode, d2.Acc, d2.Dnum)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.Dnum, getData2.Dnum)
	require.Equal(t, d2.DueDate.Format(`2006-01-02`), getData2.DueDate.Format(`2006-01-02`))
	require.Equal(t, d2.InstFlag, getData2.InstFlag)
	require.Equal(t, d2.DuePrin.String(), getData2.DuePrin.String())
	require.Equal(t, d2.DueInt.String(), getData2.DueInt.String())
	require.Equal(t, d2.UpInt.String(), getData2.UpInt.String())

	// Update Data
	updateD2 := d2
	updateD2.Acc = getData2.Acc
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestLoanInst(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetLoanInst(context.Background(), updateD2.BrCode, updateD2.Acc, updateD2.Dnum)
	require.NoError(t, err1)

	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.Equal(t, updateD2.Dnum, getData1.Dnum)
	require.Equal(t, updateD2.DueDate.Format(`2006-01-02`), getData1.DueDate.Format(`2006-01-02`))
	require.Equal(t, updateD2.InstFlag, getData1.InstFlag)
	require.Equal(t, updateD2.DuePrin.String(), getData1.DuePrin.String())
	require.Equal(t, updateD2.DueInt.String(), getData1.DueInt.String())
	require.Equal(t, updateD2.UpInt.String(), getData1.UpInt.String())

	testListLoanInst(t, ListLoanInstParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteLoanInst(t, d1.BrCode, d1.Acc, d1.Dnum)
	testDeleteLoanInst(t, d2.BrCode, d2.Acc, d2.Dnum)
}

func testListLoanInst(t *testing.T, arg ListLoanInstParams) {

	LoanInst, err := testQueriesDump.ListLoanInst(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", LoanInst)
	require.NotEmpty(t, LoanInst)

}

func randomLoanInst() model.LoanInst {

	arg := model.LoanInst{
		ModCtr:   1,
		BrCode:   "01",
		Acc:      "0101-4041-0157454",
		Dnum:     24,
		DueDate:  util.DateValue("2009-05-26"),
		InstFlag: 0,
		DuePrin:  decimal.NewFromFloat(100),
		DueInt:   decimal.NewFromFloat(100),
		UpInt:    decimal.NewFromFloat(100),
	}
	return arg
}

func createTestLoanInst(
	t *testing.T,
	req model.LoanInst) error {

	err1 := testQueriesDump.CreateLoanInst(context.Background(), req)
	// fmt.Printf("Get by createTestLoanInst%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetLoanInst(context.Background(), req.BrCode, req.Acc, req.Dnum)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.Acc, getData.Acc)
	require.Equal(t, req.Dnum, getData.Dnum)
	require.Equal(t, req.DueDate.Format(`2006-01-02`), getData.DueDate.Format(`2006-01-02`))
	require.Equal(t, req.InstFlag, getData.InstFlag)
	require.Equal(t, req.DuePrin.String(), getData.DuePrin.String())
	require.Equal(t, req.DueInt.String(), getData.DueInt.String())
	require.Equal(t, req.UpInt.String(), getData.UpInt.String())

	return err2
}

func updateTestLoanInst(
	t *testing.T,
	d1 model.LoanInst) error {

	err := testQueriesDump.UpdateLoanInst(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteLoanInst(t *testing.T, brCode string, Acc string, dNum int64) {
	err := testQueriesDump.DeleteLoanInst(context.Background(), brCode, Acc, dNum)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetLoanInst(context.Background(), brCode, Acc, dNum)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
