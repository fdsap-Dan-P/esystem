package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"
	"time"

	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestCustomer(t *testing.T) {

	// Test Data
	d1 := randomCustomer()
	d2 := randomCustomer()
	d2.CID = d2.CID + int64(1)

	err := createTestCustomer(t, d1)
	require.NoError(t, err)

	err = createTestCustomer(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetCustomer(context.Background(), d1.BrCode, d1.CID)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CID, getData1.CID)
	require.Equal(t, d1.CenterCode, getData1.CenterCode)
	require.Equal(t, d1.Title, getData1.Title)
	require.Equal(t, d1.LName, getData1.LName)
	require.Equal(t, d1.FName, getData1.FName)
	require.Equal(t, d1.MName, getData1.MName)
	require.Equal(t, d1.MaidenFName, getData1.MaidenFName)
	require.Equal(t, d1.MaidenLName, getData1.MaidenLName)
	require.Equal(t, d1.MaidenMName, getData1.MaidenMName)
	require.Equal(t, d1.Sex, getData1.Sex)
	require.Equal(t, d1.BirthDate.Time.Format(`2006-01-02`), getData1.BirthDate.Time.Format(`2006-01-02`))
	require.Equal(t, d1.BirthPlace, getData1.BirthPlace)
	require.Equal(t, d1.CivilStatus, getData1.CivilStatus)
	require.Equal(t, d1.CustType, getData1.CustType)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.Equal(t, d1.Status, getData1.Status)
	require.Equal(t, d1.Classification, getData1.Classification)
	require.Equal(t, d1.DepoType, getData1.DepoType)
	require.Equal(t, d1.SubClassification, getData1.SubClassification)
	require.True(t, d1.PledgeAmount.Decimal.Equal(getData1.PledgeAmount.Decimal))
	require.True(t, d1.MutualAmount.Decimal.Equal(getData1.MutualAmount.Decimal))
	require.True(t, d1.PangarapAmount.Decimal.Equal(getData1.PangarapAmount.Decimal))
	require.True(t, d1.KatuparanAmount.Decimal.Equal(getData1.KatuparanAmount.Decimal))
	require.True(t, d1.InsuranceAmount.Decimal.Equal(getData1.InsuranceAmount.Decimal))
	require.True(t, d1.AccPledge.Decimal.Equal(getData1.AccPledge.Decimal))
	require.True(t, d1.AccMutual.Decimal.Equal(getData1.AccMutual.Decimal))
	require.True(t, d1.AccPang.Decimal.Equal(getData1.AccPang.Decimal))
	require.True(t, d1.AccKatuparan.Decimal.Equal(getData1.AccKatuparan.Decimal))
	require.True(t, d1.AccInsurance.Decimal.Equal(getData1.AccInsurance.Decimal))
	require.True(t, d1.LoanLimit.Decimal.Equal(getData1.LoanLimit.Decimal))
	require.True(t, d1.CreditLimit.Decimal.Equal(getData1.CreditLimit.Decimal))
	require.Equal(t, d1.DateRecognized.Time.Format(`2006-01-02`), getData1.DateRecognized.Time.Format(`2006-01-02`))
	require.Equal(t, d1.DateResigned.Time.Format(`2006-01-02`), getData1.DateResigned.Time.Format(`2006-01-02`))
	require.Equal(t, d1.DateEntry.Time.Format(`2006-01-02`), getData1.DateEntry.Time.Format(`2006-01-02`))
	require.Equal(t, d1.GoldenLifeDate.Time.Format(`2006-01-02`), getData1.GoldenLifeDate.Time.Format(`2006-01-02`))
	require.Equal(t, d1.Restricted, getData1.Restricted)
	require.Equal(t, d1.Borrower, getData1.Borrower)
	require.Equal(t, d1.CoMaker, getData1.CoMaker)
	require.Equal(t, d1.Guarantor, getData1.Guarantor)
	require.Equal(t, d1.DOSRI, getData1.DOSRI)
	require.Equal(t, d1.IDCode1, getData1.IDCode1)
	require.Equal(t, d1.IDNum1, getData1.IDNum1)
	require.Equal(t, d1.IDCode2, getData1.IDCode2)
	require.Equal(t, d1.IDNum2, getData1.IDNum2)
	require.Equal(t, d1.Contact1, getData1.Contact1)
	require.Equal(t, d1.Contact2, getData1.Contact2)
	require.Equal(t, d1.Phone1, getData1.Phone1)
	require.Equal(t, d1.Reffered1, getData1.Reffered1)
	require.Equal(t, d1.Reffered2, getData1.Reffered2)
	require.Equal(t, d1.Reffered3, getData1.Reffered3)
	require.Equal(t, d1.Education, getData1.Education)
	require.Equal(t, d1.Validity1.Time.Format(`2006-01-02`), getData1.Validity1.Time.Format(`2006-01-02`))
	require.Equal(t, d1.Validity2.Time.Format(`2006-01-02`), getData1.Validity2.Time.Format(`2006-01-02`))
	require.Equal(t, d1.BusinessType, getData1.BusinessType)
	require.Equal(t, d1.AccountNumber, getData1.AccountNumber)
	require.Equal(t, d1.IIID, getData1.IIID)
	require.Equal(t, d1.Religion, getData1.Religion)

	getData2, err2 := testQueriesDump.GetCustomer(context.Background(), d2.BrCode, d2.CID)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CID, getData2.CID)
	require.Equal(t, d2.CenterCode, getData2.CenterCode)
	require.Equal(t, d2.Title, getData2.Title)
	require.Equal(t, d2.LName, getData2.LName)
	require.Equal(t, d2.FName, getData2.FName)
	require.Equal(t, d2.MName, getData2.MName)
	require.Equal(t, d2.MaidenFName, getData2.MaidenFName)
	require.Equal(t, d2.MaidenLName, getData2.MaidenLName)
	require.Equal(t, d2.MaidenMName, getData2.MaidenMName)
	require.Equal(t, d2.Sex, getData2.Sex)
	require.Equal(t, d2.BirthDate.Time.Format(`2006-01-02`), getData2.BirthDate.Time.Format(`2006-01-02`))
	require.Equal(t, d2.BirthPlace, getData2.BirthPlace)
	require.Equal(t, d2.CivilStatus, getData2.CivilStatus)
	require.Equal(t, d2.CustType, getData2.CustType)
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.Equal(t, d2.Status, getData2.Status)
	require.Equal(t, d2.Classification, getData2.Classification)
	require.Equal(t, d2.DepoType, getData2.DepoType)
	require.Equal(t, d2.SubClassification, getData2.SubClassification)
	require.True(t, d2.PledgeAmount.Decimal.Equal(getData2.PledgeAmount.Decimal))
	require.True(t, d2.MutualAmount.Decimal.Equal(getData2.MutualAmount.Decimal))
	require.True(t, d2.PangarapAmount.Decimal.Equal(getData2.PangarapAmount.Decimal))
	require.True(t, d2.KatuparanAmount.Decimal.Equal(getData2.KatuparanAmount.Decimal))
	require.True(t, d2.InsuranceAmount.Decimal.Equal(getData2.InsuranceAmount.Decimal))
	require.True(t, d2.AccPledge.Decimal.Equal(getData2.AccPledge.Decimal))
	require.True(t, d2.AccMutual.Decimal.Equal(getData2.AccMutual.Decimal))
	require.True(t, d2.AccPang.Decimal.Equal(getData2.AccPang.Decimal))
	require.True(t, d2.AccKatuparan.Decimal.Equal(getData2.AccKatuparan.Decimal))
	require.True(t, d2.AccInsurance.Decimal.Equal(getData2.AccInsurance.Decimal))
	require.True(t, d2.LoanLimit.Decimal.Equal(getData2.LoanLimit.Decimal))
	require.True(t, d2.CreditLimit.Decimal.Equal(getData2.CreditLimit.Decimal))
	require.Equal(t, d2.DateRecognized.Time.Format(`2006-01-02`), getData2.DateRecognized.Time.Format(`2006-01-02`))
	require.Equal(t, d2.DateResigned.Time.Format(`2006-01-02`), getData2.DateResigned.Time.Format(`2006-01-02`))
	require.Equal(t, d2.DateEntry.Time.Format(`2006-01-02`), getData2.DateEntry.Time.Format(`2006-01-02`))
	require.Equal(t, d2.GoldenLifeDate.Time.Format(`2006-01-02`), getData2.GoldenLifeDate.Time.Format(`2006-01-02`))
	require.Equal(t, d2.Restricted, getData2.Restricted)
	require.Equal(t, d2.Borrower, getData2.Borrower)
	require.Equal(t, d2.CoMaker, getData2.CoMaker)
	require.Equal(t, d2.Guarantor, getData2.Guarantor)
	require.Equal(t, d2.DOSRI, getData2.DOSRI)
	require.Equal(t, d2.IDCode1, getData2.IDCode1)
	require.Equal(t, d2.IDNum1, getData2.IDNum1)
	require.Equal(t, d2.IDCode2, getData2.IDCode2)
	require.Equal(t, d2.IDNum2, getData2.IDNum2)
	require.Equal(t, d2.Contact1, getData2.Contact1)
	require.Equal(t, d2.Contact2, getData2.Contact2)
	require.Equal(t, d2.Phone1, getData2.Phone1)
	require.Equal(t, d2.Reffered1, getData2.Reffered1)
	require.Equal(t, d2.Reffered2, getData2.Reffered2)
	require.Equal(t, d2.Reffered3, getData2.Reffered3)
	require.Equal(t, d2.Education, getData2.Education)
	require.Equal(t, d2.Validity1.Time.Format(`2006-01-02`), getData2.Validity1.Time.Format(`2006-01-02`))
	require.Equal(t, d2.Validity2.Time.Format(`2006-01-02`), getData2.Validity2.Time.Format(`2006-01-02`))
	require.Equal(t, d2.BusinessType, getData2.BusinessType)
	require.Equal(t, d2.AccountNumber, getData2.AccountNumber)
	require.Equal(t, d2.IIID, getData2.IIID)
	require.Equal(t, d2.Religion, getData2.Religion)

	// Update Data
	updateD2 := d2
	updateD2.CID = getData2.CID
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestCustomer(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetCustomer(context.Background(), updateD2.BrCode, updateD2.CID)
	require.NoError(t, err1)

	require.Equal(t, updateD2.CID, getData1.CID)
	require.Equal(t, updateD2.CenterCode, getData1.CenterCode)
	require.Equal(t, updateD2.Title, getData1.Title)
	require.Equal(t, updateD2.LName, getData1.LName)
	require.Equal(t, updateD2.FName, getData1.FName)
	require.Equal(t, updateD2.MName, getData1.MName)
	require.Equal(t, updateD2.MaidenFName, getData1.MaidenFName)
	require.Equal(t, updateD2.MaidenLName, getData1.MaidenLName)
	require.Equal(t, updateD2.MaidenMName, getData1.MaidenMName)
	require.Equal(t, updateD2.Sex, getData1.Sex)
	require.Equal(t, updateD2.BirthDate.Time.Format(`2006-01-02`), getData1.BirthDate.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.BirthPlace, getData1.BirthPlace)
	require.Equal(t, updateD2.CivilStatus, getData1.CivilStatus)
	require.Equal(t, updateD2.CustType, getData1.CustType)
	require.Equal(t, updateD2.Remarks, getData1.Remarks)
	require.Equal(t, updateD2.Status, getData1.Status)
	require.Equal(t, updateD2.Classification, getData1.Classification)
	require.Equal(t, updateD2.DepoType, getData1.DepoType)
	require.Equal(t, updateD2.SubClassification, getData1.SubClassification)
	require.True(t, updateD2.PledgeAmount.Decimal.Equal(getData1.PledgeAmount.Decimal))
	require.True(t, updateD2.MutualAmount.Decimal.Equal(getData1.MutualAmount.Decimal))
	require.True(t, updateD2.PangarapAmount.Decimal.Equal(getData1.PangarapAmount.Decimal))
	require.True(t, updateD2.KatuparanAmount.Decimal.Equal(getData1.KatuparanAmount.Decimal))
	require.True(t, updateD2.InsuranceAmount.Decimal.Equal(getData1.InsuranceAmount.Decimal))
	require.True(t, updateD2.AccPledge.Decimal.Equal(getData1.AccPledge.Decimal))
	require.True(t, updateD2.AccMutual.Decimal.Equal(getData1.AccMutual.Decimal))
	require.True(t, updateD2.AccPang.Decimal.Equal(getData1.AccPang.Decimal))
	require.True(t, updateD2.AccKatuparan.Decimal.Equal(getData1.AccKatuparan.Decimal))
	require.True(t, updateD2.AccInsurance.Decimal.Equal(getData1.AccInsurance.Decimal))
	require.True(t, updateD2.LoanLimit.Decimal.Equal(getData1.LoanLimit.Decimal))
	require.True(t, updateD2.CreditLimit.Decimal.Equal(getData1.CreditLimit.Decimal))
	require.Equal(t, updateD2.DateRecognized.Time.Format(`2006-01-02`), getData1.DateRecognized.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.DateResigned.Time.Format(`2006-01-02`), getData1.DateResigned.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.DateEntry.Time.Format(`2006-01-02`), getData1.DateEntry.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.GoldenLifeDate.Time.Format(`2006-01-02`), getData1.GoldenLifeDate.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.Restricted, getData1.Restricted)
	require.Equal(t, updateD2.Borrower, getData1.Borrower)
	require.Equal(t, updateD2.CoMaker, getData1.CoMaker)
	require.Equal(t, updateD2.Guarantor, getData1.Guarantor)
	require.Equal(t, updateD2.DOSRI, getData1.DOSRI)
	require.Equal(t, updateD2.IDCode1, getData1.IDCode1)
	require.Equal(t, updateD2.IDNum1, getData1.IDNum1)
	require.Equal(t, updateD2.IDCode2, getData1.IDCode2)
	require.Equal(t, updateD2.IDNum2, getData1.IDNum2)
	require.Equal(t, updateD2.Contact1, getData1.Contact1)
	require.Equal(t, updateD2.Contact2, getData1.Contact2)
	require.Equal(t, updateD2.Phone1, getData1.Phone1)
	require.Equal(t, updateD2.Reffered1, getData1.Reffered1)
	require.Equal(t, updateD2.Reffered2, getData1.Reffered2)
	require.Equal(t, updateD2.Reffered3, getData1.Reffered3)
	require.Equal(t, updateD2.Education, getData1.Education)
	require.Equal(t, updateD2.Validity1.Time.Format(`2006-01-02`), getData1.Validity1.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.Validity2.Time.Format(`2006-01-02`), getData1.Validity2.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.BusinessType, getData1.BusinessType)
	require.Equal(t, updateD2.AccountNumber, getData1.AccountNumber)
	require.Equal(t, updateD2.IIID, getData1.IIID)
	require.Equal(t, updateD2.Religion, getData1.Religion)

	testListCustomer(t, ListCustomerParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteCustomer(t, d1.BrCode, d1.CID)
	testDeleteCustomer(t, d2.BrCode, d2.CID)
}

func testListCustomer(t *testing.T, arg ListCustomerParams) {

	Customer, err := testQueriesDump.ListCustomer(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Customer)
	require.NotEmpty(t, Customer)

}

func randomCustomer() model.Customer {

	arg := model.Customer{
		ModCtr:            1,
		BrCode:            "01",
		CID:               int64(19858200) + 1,
		CenterCode:        sql.NullString{String: "01", Valid: true},
		Title:             sql.NullInt64{Int64: 1, Valid: true},
		LName:             sql.NullString{String: "Last Name", Valid: true},
		FName:             sql.NullString{String: "First Name", Valid: true},
		MName:             sql.NullString{String: "M", Valid: true},
		MaidenFName:       sql.NullString{String: "", Valid: true},
		MaidenLName:       sql.NullString{String: "", Valid: true},
		MaidenMName:       sql.NullString{String: "", Valid: true},
		Sex:               sql.NullString{String: "M", Valid: true},
		BirthDate:         sql.NullTime{Time: time.Now(), Valid: true},
		BirthPlace:        sql.NullString{String: "", Valid: true},
		CivilStatus:       sql.NullInt64{Int64: 1, Valid: true},
		CustType:          sql.NullInt64{Int64: 1, Valid: true},
		Remarks:           sql.NullString{String: "", Valid: true},
		Status:            sql.NullInt64{Int64: 1, Valid: true},
		Classification:    sql.NullInt64{Int64: 1, Valid: true},
		DepoType:          sql.NullString{String: "", Valid: true},
		SubClassification: sql.NullInt64{Int64: 1, Valid: true},
		PledgeAmount:      decimal.NewNullDecimal(decimal.Zero),
		MutualAmount:      decimal.NewNullDecimal(decimal.Zero),
		PangarapAmount:    decimal.NewNullDecimal(decimal.Zero),
		KatuparanAmount:   decimal.NewNullDecimal(decimal.Zero),
		InsuranceAmount:   decimal.NewNullDecimal(decimal.Zero),
		AccPledge:         decimal.NewNullDecimal(decimal.Zero),
		AccMutual:         decimal.NewNullDecimal(decimal.Zero),
		AccPang:           decimal.NewNullDecimal(decimal.Zero),
		AccKatuparan:      decimal.NewNullDecimal(decimal.Zero),
		AccInsurance:      decimal.NewNullDecimal(decimal.Zero),
		LoanLimit:         decimal.NewNullDecimal(decimal.Zero),
		CreditLimit:       decimal.NewNullDecimal(decimal.Zero),
		DateRecognized:    sql.NullTime{Time: time.Now(), Valid: true},
		DateResigned:      sql.NullTime{Time: time.Now(), Valid: true},
		DateEntry:         sql.NullTime{Time: time.Now(), Valid: true},
		GoldenLifeDate:    sql.NullTime{Time: time.Now(), Valid: true},
		Restricted:        sql.NullString{String: "", Valid: true},
		Borrower:          sql.NullString{String: "", Valid: true},
		CoMaker:           sql.NullString{String: "", Valid: true},
		Guarantor:         sql.NullString{String: "", Valid: true},
		DOSRI:             sql.NullInt64{Int64: 1, Valid: true},
		IDCode1:           sql.NullInt64{Int64: 1, Valid: true},
		IDNum1:            sql.NullString{String: "", Valid: true},
		IDCode2:           sql.NullInt64{Int64: 1, Valid: true},
		IDNum2:            sql.NullString{String: "", Valid: true},
		Contact1:          sql.NullString{String: "", Valid: true},
		Contact2:          sql.NullString{String: "", Valid: true},
		Phone1:            sql.NullString{String: "", Valid: true},
		Reffered1:         sql.NullString{String: "", Valid: true},
		Reffered2:         sql.NullString{String: "", Valid: true},
		Reffered3:         sql.NullString{String: "", Valid: true},
		Education:         sql.NullInt64{Int64: 1, Valid: true},
		Validity1:         sql.NullTime{Time: time.Now(), Valid: true},
		Validity2:         sql.NullTime{Time: time.Now(), Valid: true},
		BusinessType:      sql.NullInt64{Int64: 1, Valid: true},
		AccountNumber:     sql.NullString{String: "", Valid: true},
		IIID:              sql.NullInt64{Int64: 1, Valid: true},
		Religion:          sql.NullInt64{Int64: 1, Valid: true},
	}
	return arg
}

func createTestCustomer(
	t *testing.T,
	req model.Customer) error {

	err1 := testQueriesDump.CreateCustomer(context.Background(), req)
	// fmt.Printf("Get by createTestCustomer%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetCustomer(context.Background(), req.BrCode, req.CID)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.CID, getData.CID)
	require.Equal(t, req.CenterCode, getData.CenterCode)
	require.Equal(t, req.Title, getData.Title)
	require.Equal(t, req.LName, getData.LName)
	require.Equal(t, req.FName, getData.FName)
	require.Equal(t, req.MName, getData.MName)
	require.Equal(t, req.MaidenFName, getData.MaidenFName)
	require.Equal(t, req.MaidenLName, getData.MaidenLName)
	require.Equal(t, req.MaidenMName, getData.MaidenMName)
	require.Equal(t, req.Sex, getData.Sex)
	require.Equal(t, req.BirthDate.Time.Format(`2006-01-02`), getData.BirthDate.Time.Format(`2006-01-02`))
	require.Equal(t, req.BirthPlace, getData.BirthPlace)
	require.Equal(t, req.CivilStatus, getData.CivilStatus)
	require.Equal(t, req.CustType, getData.CustType)
	require.Equal(t, req.Remarks, getData.Remarks)
	require.Equal(t, req.Status, getData.Status)
	require.Equal(t, req.Classification, getData.Classification)
	require.Equal(t, req.DepoType, getData.DepoType)
	require.Equal(t, req.SubClassification, getData.SubClassification)
	require.True(t, req.PledgeAmount.Decimal.Equal(getData.PledgeAmount.Decimal))
	require.True(t, req.MutualAmount.Decimal.Equal(getData.MutualAmount.Decimal))
	require.True(t, req.PangarapAmount.Decimal.Equal(getData.PangarapAmount.Decimal))
	require.True(t, req.KatuparanAmount.Decimal.Equal(getData.KatuparanAmount.Decimal))
	require.True(t, req.InsuranceAmount.Decimal.Equal(getData.InsuranceAmount.Decimal))
	require.True(t, req.AccPledge.Decimal.Equal(getData.AccPledge.Decimal))
	require.True(t, req.AccMutual.Decimal.Equal(getData.AccMutual.Decimal))
	require.True(t, req.AccPang.Decimal.Equal(getData.AccPang.Decimal))
	require.True(t, req.AccKatuparan.Decimal.Equal(getData.AccKatuparan.Decimal))
	require.True(t, req.AccInsurance.Decimal.Equal(getData.AccInsurance.Decimal))
	require.True(t, req.LoanLimit.Decimal.Equal(getData.LoanLimit.Decimal))
	require.True(t, req.CreditLimit.Decimal.Equal(getData.CreditLimit.Decimal))
	require.Equal(t, req.DateRecognized.Time.Format(`2006-01-02`), getData.DateRecognized.Time.Format(`2006-01-02`))
	require.Equal(t, req.DateResigned.Time.Format(`2006-01-02`), getData.DateResigned.Time.Format(`2006-01-02`))
	require.Equal(t, req.DateEntry.Time.Format(`2006-01-02`), getData.DateEntry.Time.Format(`2006-01-02`))
	require.Equal(t, req.GoldenLifeDate.Time.Format(`2006-01-02`), getData.GoldenLifeDate.Time.Format(`2006-01-02`))
	require.Equal(t, req.Restricted, getData.Restricted)
	require.Equal(t, req.Borrower, getData.Borrower)
	require.Equal(t, req.CoMaker, getData.CoMaker)
	require.Equal(t, req.Guarantor, getData.Guarantor)
	require.Equal(t, req.DOSRI, getData.DOSRI)
	require.Equal(t, req.IDCode1, getData.IDCode1)
	require.Equal(t, req.IDNum1, getData.IDNum1)
	require.Equal(t, req.IDCode2, getData.IDCode2)
	require.Equal(t, req.IDNum2, getData.IDNum2)
	require.Equal(t, req.Contact1, getData.Contact1)
	require.Equal(t, req.Contact2, getData.Contact2)
	require.Equal(t, req.Phone1, getData.Phone1)
	require.Equal(t, req.Reffered1, getData.Reffered1)
	require.Equal(t, req.Reffered2, getData.Reffered2)
	require.Equal(t, req.Reffered3, getData.Reffered3)
	require.Equal(t, req.Education, getData.Education)
	require.Equal(t, req.Validity1.Time.Format(`2006-01-02`), getData.Validity1.Time.Format(`2006-01-02`))
	require.Equal(t, req.Validity2.Time.Format(`2006-01-02`), getData.Validity2.Time.Format(`2006-01-02`))
	require.Equal(t, req.BusinessType, getData.BusinessType)
	require.Equal(t, req.AccountNumber, getData.AccountNumber)
	require.Equal(t, req.IIID, getData.IIID)
	require.Equal(t, req.Religion, getData.Religion)

	return err2
}

func updateTestCustomer(
	t *testing.T,
	d1 model.Customer) error {

	err := testQueriesDump.UpdateCustomer(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteCustomer(t *testing.T, brCode string, CID int64) {
	err := testQueriesDump.DeleteCustomer(context.Background(), brCode, CID)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetCustomer(context.Background(), brCode, CID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
