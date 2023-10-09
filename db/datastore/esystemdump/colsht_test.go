package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"
	"simplebank/util"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestBulkInsertColSht(t *testing.T) {

	// Test Data
	d1 := randomColSht()
	d2 := randomColSht()
	d2.Acc = "E40304-1012-0000192"

	err1 := testQueriesDump.BulkInsertColSht(context.Background(), []model.ColSht{d1, d2})
	require.NoError(t, err1)

	// Delete Data
	testDeleteColSht(t, d1.BrCode)
	testDeleteColSht(t, d2.BrCode)
}

func TestColSht(t *testing.T) {

	// Test Data
	d1 := randomColSht()
	d2 := randomColSht()
	d2.Acc = "E40304-1012-0000192"

	err := createTestColSht(t, d1)
	require.NoError(t, err)

	err = createTestColSht(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetColSht(context.Background(), d1.Acc)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AppType, getData1.AppType)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.Status, getData1.Status)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.CID, getData1.CID)
	require.Equal(t, d1.UM, getData1.UM)
	require.Equal(t, d1.ClientName, getData1.ClientName)
	require.Equal(t, d1.CenterCode, getData1.CenterCode)
	require.Equal(t, d1.CenterName, getData1.CenterName)
	require.Equal(t, d1.ManCode, getData1.ManCode)
	require.Equal(t, d1.Unit, getData1.Unit)
	require.Equal(t, d1.AreaCode, getData1.AreaCode)
	require.Equal(t, d1.Area, getData1.Area)
	require.Equal(t, d1.StaffName, getData1.StaffName)
	require.Equal(t, d1.AcctType, getData1.AcctType)
	require.Equal(t, d1.AcctDesc, getData1.AcctDesc)
	require.Equal(t, d1.DisbDate.Format(`2006-01-02`), getData1.DisbDate.Format(`2006-01-02`))
	require.Equal(t, d1.DateStart.Format(`2006-01-02`), getData1.DateStart.Format(`2006-01-02`))
	require.Equal(t, d1.Maturity.Format(`2006-01-02`), getData1.Maturity.Format(`2006-01-02`))
	require.True(t, d1.Principal.Equal(getData1.Principal))
	require.True(t, d1.Interest.Equal(getData1.Interest))
	require.Equal(t, d1.Gives, getData1.Gives)
	require.True(t, d1.BalPrin.Equal(getData1.BalPrin))
	require.True(t, d1.BalInt.Equal(getData1.BalInt))
	require.True(t, d1.Amort.Equal(getData1.Amort))
	require.True(t, d1.DuePrin.Equal(getData1.DuePrin))
	require.True(t, d1.DueInt.Equal(getData1.DueInt))
	require.True(t, d1.LoanBal.Equal(getData1.LoanBal))
	require.True(t, d1.SaveBal.Equal(getData1.SaveBal))
	require.True(t, d1.WaivedInt.Equal(getData1.WaivedInt))
	require.Equal(t, d1.UnPaidCtr, getData1.UnPaidCtr)
	require.Equal(t, d1.WrittenOff, getData1.WrittenOff)
	require.Equal(t, d1.OrgName, getData1.OrgName)
	require.Equal(t, d1.OrgAddress, getData1.OrgAddress)
	require.Equal(t, d1.MeetingDate.Format(`2006-01-02`), getData1.MeetingDate.Format(`2006-01-02`))
	require.Equal(t, d1.MeetingDay, getData1.MeetingDay)
	require.True(t, d1.SharesOfStock.Equal(getData1.SharesOfStock))
	require.Equal(t, d1.DateEstablished.Format(`2006-01-02`), getData1.DateEstablished.Format(`2006-01-02`))
	require.Equal(t, d1.Classification, getData1.Classification)
	require.Equal(t, d1.WriteOff, getData1.WriteOff)

	getData2, err2 := testQueriesDump.GetColSht(context.Background(), d2.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AppType, getData2.AppType)
	require.Equal(t, d2.Code, getData2.Code)
	require.Equal(t, d2.Status, getData2.Status)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.CID, getData2.CID)
	require.Equal(t, d2.UM, getData2.UM)
	require.Equal(t, d2.ClientName, getData2.ClientName)
	require.Equal(t, d2.CenterCode, getData2.CenterCode)
	require.Equal(t, d2.CenterName, getData2.CenterName)
	require.Equal(t, d2.ManCode, getData2.ManCode)
	require.Equal(t, d2.Unit, getData2.Unit)
	require.Equal(t, d2.AreaCode, getData2.AreaCode)
	require.Equal(t, d2.Area, getData2.Area)
	require.Equal(t, d2.StaffName, getData2.StaffName)
	require.Equal(t, d2.AcctType, getData2.AcctType)
	require.Equal(t, d2.AcctDesc, getData2.AcctDesc)
	require.Equal(t, d2.DisbDate.Format(`2006-01-02`), getData2.DisbDate.Format(`2006-01-02`))
	require.Equal(t, d2.DateStart.Format(`2006-01-02`), getData2.DateStart.Format(`2006-01-02`))
	require.Equal(t, d2.Maturity.Format(`2006-01-02`), getData2.Maturity.Format(`2006-01-02`))
	require.True(t, d2.Principal.Equal(getData2.Principal))
	require.True(t, d2.Interest.Equal(getData2.Interest))
	require.Equal(t, d2.Gives, getData2.Gives)
	require.True(t, d2.BalPrin.Equal(getData2.BalPrin))
	require.True(t, d2.BalInt.Equal(getData2.BalInt))
	require.True(t, d2.Amort.Equal(getData2.Amort))
	require.True(t, d2.DuePrin.Equal(getData2.DuePrin))
	require.True(t, d2.DueInt.Equal(getData2.DueInt))
	require.True(t, d2.LoanBal.Equal(getData2.LoanBal))
	require.True(t, d2.SaveBal.Equal(getData2.SaveBal))
	require.True(t, d2.WaivedInt.Equal(getData2.WaivedInt))
	require.Equal(t, d2.UnPaidCtr, getData2.UnPaidCtr)
	require.Equal(t, d2.WrittenOff, getData2.WrittenOff)
	require.Equal(t, d2.OrgName, getData2.OrgName)
	require.Equal(t, d2.OrgAddress, getData2.OrgAddress)
	require.Equal(t, d2.MeetingDate.Format(`2006-01-02`), getData2.MeetingDate.Format(`2006-01-02`))
	require.Equal(t, d2.MeetingDay, getData2.MeetingDay)
	require.True(t, d2.SharesOfStock.Equal(getData2.SharesOfStock))
	require.Equal(t, d2.DateEstablished.Format(`2006-01-02`), getData2.DateEstablished.Format(`2006-01-02`))
	require.Equal(t, d2.Classification, getData2.Classification)
	require.Equal(t, d2.WriteOff, getData2.WriteOff)

	// Update Data
	updateD2 := d2
	updateD2.Principal = util.SetDecimal("143")
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestColSht(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetColSht(context.Background(), updateD2.Acc)
	require.NoError(t, err1)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, updateD2.AppType, getData1.AppType)
	require.Equal(t, updateD2.Code, getData1.Code)
	require.Equal(t, updateD2.Status, getData1.Status)
	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.Equal(t, updateD2.CID, getData1.CID)
	require.Equal(t, updateD2.UM, getData1.UM)
	require.Equal(t, updateD2.ClientName, getData1.ClientName)
	require.Equal(t, updateD2.CenterCode, getData1.CenterCode)
	require.Equal(t, updateD2.CenterName, getData1.CenterName)
	require.Equal(t, updateD2.ManCode, getData1.ManCode)
	require.Equal(t, updateD2.Unit, getData1.Unit)
	require.Equal(t, updateD2.AreaCode, getData1.AreaCode)
	require.Equal(t, updateD2.Area, getData1.Area)
	require.Equal(t, updateD2.StaffName, getData1.StaffName)
	require.Equal(t, updateD2.AcctType, getData1.AcctType)
	require.Equal(t, updateD2.AcctDesc, getData1.AcctDesc)
	require.Equal(t, updateD2.DisbDate.Format(`2006-01-02`), getData1.DisbDate.Format(`2006-01-02`))
	require.Equal(t, updateD2.DateStart.Format(`2006-01-02`), getData1.DateStart.Format(`2006-01-02`))
	require.Equal(t, updateD2.Maturity.Format(`2006-01-02`), getData1.Maturity.Format(`2006-01-02`))
	require.True(t, updateD2.Principal.Equal(getData1.Principal))
	require.True(t, updateD2.Interest.Equal(getData1.Interest))
	require.Equal(t, updateD2.Gives, getData1.Gives)
	require.True(t, updateD2.BalPrin.Equal(getData1.BalPrin))
	require.True(t, updateD2.BalInt.Equal(getData1.BalInt))
	require.True(t, updateD2.Amort.Equal(getData1.Amort))
	require.True(t, updateD2.DuePrin.Equal(getData1.DuePrin))
	require.True(t, updateD2.DueInt.Equal(getData1.DueInt))
	require.True(t, updateD2.LoanBal.Equal(getData1.LoanBal))
	require.True(t, updateD2.SaveBal.Equal(getData1.SaveBal))
	require.True(t, updateD2.WaivedInt.Equal(getData1.WaivedInt))
	require.Equal(t, updateD2.UnPaidCtr, getData1.UnPaidCtr)
	require.Equal(t, updateD2.WrittenOff, getData1.WrittenOff)
	require.Equal(t, updateD2.OrgName, getData1.OrgName)
	require.Equal(t, updateD2.OrgAddress, getData1.OrgAddress)
	require.Equal(t, updateD2.MeetingDate.Format(`2006-01-02`), getData1.MeetingDate.Format(`2006-01-02`))
	require.Equal(t, updateD2.MeetingDay, getData1.MeetingDay)
	require.True(t, updateD2.SharesOfStock.Equal(getData1.SharesOfStock))
	require.Equal(t, updateD2.DateEstablished.Format(`2006-01-02`), getData1.DateEstablished.Format(`2006-01-02`))
	require.Equal(t, updateD2.Classification, getData1.Classification)
	require.Equal(t, updateD2.WriteOff, getData1.WriteOff)

	testColShtPerCID(t, d1.CID)

	err1 = testQueriesDump.BulkInsertColSht(context.Background(), []model.ColSht{d1, d2})
	require.NoError(t, err1)

	// Delete Data
	testDeleteColSht(t, d1.BrCode)
	testDeleteColSht(t, d2.BrCode)
}

func testColShtPerCID(t *testing.T, cid int64) {

	ColSht, err := testQueriesDump.ColShtPerCID(context.Background(), cid)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", ColSht)
	require.NotEmpty(t, ColSht)

}

func randomColSht() model.ColSht {

	arg := model.ColSht{
		BrCode:          "E4",
		AppType:         0,
		Code:            0,
		Status:          1,
		Acc:             "E4MBA400198",
		CID:             400198,
		UM:              "De la Rama, Rosalin B.",
		ClientName:      "Belolo, Joselyn L.",
		CenterCode:      "18",
		CenterName:      "DEL ROSARIO",
		ManCode:         512,
		Unit:            "Milaor",
		AreaCode:        101,
		Area:            "Bicol 3",
		StaffName:       "Hernandez, Franklin P.",
		AcctType:        0,
		AcctDesc:        "MBA",
		DisbDate:        util.SetDate("2022-10-24"),
		DateStart:       util.SetDate("2022-10-24"),
		Maturity:        util.SetDate("2022-10-24"),
		Principal:       util.SetDecimal("0"),
		Interest:        util.SetDecimal("0"),
		Gives:           0,
		BalPrin:         util.SetDecimal("0"),
		BalInt:          util.SetDecimal("0"),
		Amort:           util.SetDecimal("0"),
		DuePrin:         util.SetDecimal("0"),
		DueInt:          util.SetDecimal("0"),
		LoanBal:         util.SetDecimal("0"),
		SaveBal:         util.SetDecimal("0"),
		WaivedInt:       util.SetDecimal("0"),
		UnPaidCtr:       0,
		WrittenOff:      0,
		OrgName:         "CARD, INC.",
		OrgAddress:      "Zone 1, Barangay Del Rosario, Milaor, Camarines Sur",
		MeetingDate:     util.SetDate("2022-10-24"),
		MeetingDay:      1,
		SharesOfStock:   util.SetDecimal("0"),
		DateEstablished: util.SetDate("2022-10-24"),
		Classification:  1555,
		WriteOff:        0,
	}
	return arg
}

func createTestColSht(
	t *testing.T,
	req model.ColSht) error {

	err1 := testQueriesDump.CreateColSht(context.Background(), req)
	// fmt.Printf("Get by createTestColSht%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetColSht(context.Background(), req.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.AppType, getData.AppType)
	require.Equal(t, req.Code, getData.Code)
	require.Equal(t, req.Status, getData.Status)
	require.Equal(t, req.Acc, getData.Acc)
	require.Equal(t, req.CID, getData.CID)
	require.Equal(t, req.UM, getData.UM)
	require.Equal(t, req.ClientName, getData.ClientName)
	require.Equal(t, req.CenterCode, getData.CenterCode)
	require.Equal(t, req.CenterName, getData.CenterName)
	require.Equal(t, req.ManCode, getData.ManCode)
	require.Equal(t, req.Unit, getData.Unit)
	require.Equal(t, req.AreaCode, getData.AreaCode)
	require.Equal(t, req.Area, getData.Area)
	require.Equal(t, req.StaffName, getData.StaffName)
	require.Equal(t, req.AcctType, getData.AcctType)
	require.Equal(t, req.AcctDesc, getData.AcctDesc)
	require.Equal(t, req.DisbDate.Format(`2006-01-02`), getData.DisbDate.Format(`2006-01-02`))
	require.Equal(t, req.DateStart.Format(`2006-01-02`), getData.DateStart.Format(`2006-01-02`))
	require.Equal(t, req.Maturity.Format(`2006-01-02`), getData.Maturity.Format(`2006-01-02`))
	require.True(t, req.Principal.Equal(getData.Principal))
	require.True(t, req.Interest.Equal(getData.Interest))
	require.Equal(t, req.Gives, getData.Gives)
	require.True(t, req.BalPrin.Equal(getData.BalPrin))
	require.True(t, req.BalInt.Equal(getData.BalInt))
	require.True(t, req.Amort.Equal(getData.Amort))
	require.True(t, req.DuePrin.Equal(getData.DuePrin))
	require.True(t, req.DueInt.Equal(getData.DueInt))
	require.True(t, req.LoanBal.Equal(getData.LoanBal))
	require.True(t, req.SaveBal.Equal(getData.SaveBal))
	require.True(t, req.WaivedInt.Equal(getData.WaivedInt))
	require.Equal(t, req.UnPaidCtr, getData.UnPaidCtr)
	require.Equal(t, req.WrittenOff, getData.WrittenOff)
	require.Equal(t, req.OrgName, getData.OrgName)
	require.Equal(t, req.OrgAddress, getData.OrgAddress)
	require.Equal(t, req.MeetingDate.Format(`2006-01-02`), getData.MeetingDate.Format(`2006-01-02`))
	require.Equal(t, req.MeetingDay, getData.MeetingDay)
	require.True(t, req.SharesOfStock.Equal(getData.SharesOfStock))
	require.Equal(t, req.DateEstablished.Format(`2006-01-02`), getData.DateEstablished.Format(`2006-01-02`))
	require.Equal(t, req.Classification, getData.Classification)
	require.Equal(t, req.WriteOff, getData.WriteOff)

	return err2
}

func updateTestColSht(
	t *testing.T,
	d1 model.ColSht) error {

	err := testQueriesDump.UpdateColSht(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteColSht(t *testing.T, brCode string) {
	err := testQueriesDump.DeleteColSht(context.Background(), brCode)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetColSht(context.Background(), brCode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
