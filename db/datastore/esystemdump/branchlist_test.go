package db

import (
	"context"
	"database/sql"
	"simplebank/util"
	"time"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestBranchList(t *testing.T) {

	// Test Data
	d1 := randomBranchList()
	d2 := randomBranchList()
	d2.BrCode = "01"

	err := createTestBranchList(t, d1)
	require.NoError(t, err)

	err = createTestBranchList(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetBranchList(context.Background(), d1.BrCode)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.EbSysDate.Format(`2006-01-02`), getData1.EbSysDate.Format(`2006-01-02`))
	require.Equal(t, d1.RunState, getData1.RunState)
	require.Equal(t, d1.OrgAddress, getData1.OrgAddress)
	require.Equal(t, d1.TaxInfo, getData1.TaxInfo)
	require.Equal(t, d1.DefCity, getData1.DefCity)
	require.Equal(t, d1.DefProvince, getData1.DefProvince)
	require.Equal(t, d1.DefCountry, getData1.DefCountry)
	require.Equal(t, d1.DefZip, getData1.DefZip)
	require.Equal(t, d1.WaivableInt, getData1.WaivableInt)
	require.Equal(t, d1.DBVersion, getData1.DBVersion)
	require.Equal(t, d1.ESystemVer, getData1.ESystemVer)
	require.Equal(t, d1.NewBrCode, getData1.NewBrCode)

	getData2, err2 := testQueriesDump.GetBranchList(context.Background(), d2.BrCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.EbSysDate.Format(`2006-01-02`), getData2.EbSysDate.Format(`2006-01-02`))
	require.Equal(t, d2.RunState, getData2.RunState)
	require.Equal(t, d2.OrgAddress, getData2.OrgAddress)
	require.Equal(t, d2.TaxInfo, getData2.TaxInfo)
	require.Equal(t, d2.DefCity, getData2.DefCity)
	require.Equal(t, d2.DefProvince, getData2.DefProvince)
	require.Equal(t, d2.DefCountry, getData2.DefCountry)
	require.Equal(t, d2.DefZip, getData2.DefZip)
	require.Equal(t, d2.WaivableInt, getData2.WaivableInt)
	require.Equal(t, d2.DBVersion, getData2.DBVersion)
	require.Equal(t, d2.ESystemVer, getData2.ESystemVer)
	require.Equal(t, d2.NewBrCode, getData2.NewBrCode)

	// Update Data
	updateD2 := BranchListLight{
		BrCode:    d2.BrCode,
		EbSysDate: util.RandomDate(),
		RunState:  10,
	}
	updateD2.EbSysDate = util.RandomDate()
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestBranchList(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetBranchList(context.Background(), updateD2.BrCode)
	require.NoError(t, err1)
	require.Equal(t, updateD2.EbSysDate.Format(`2006-01-02`), getData1.EbSysDate.Format(`2006-01-02`))
	require.Equal(t, updateD2.RunState, getData1.RunState)

	testListBranchList(t, ListBranchListParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	// testDeleteBranchList(t, d1.BrCode)
	// testDeleteBranchList(t, d2.BrCode)
}

func testListBranchList(t *testing.T, arg ListBranchListParams) {

	branchList, err := testQueriesDump.ListBranchList(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", branchList)
	require.NotEmpty(t, branchList)

}

func randomBranchList() BranchList {

	arg := BranchList{
		BrCode:         "E3",
		EbSysDate:      util.SetDate("2022-10-28"),
		RunState:       99,
		OrgAddress:     "Zone 5, Tagbong, Pili, Camarines Sur",
		TaxInfo:        "",
		DefCity:        "Naga City",
		DefProvince:    "Camarines Sur",
		DefCountry:     "Philippines",
		DefZip:         "4400",
		WaivableInt:    true,
		DBVersion:      "4.21.0.0",
		ESystemVer:     []byte("¡ò=ñ( åèéc+ê i  È © =Ä ^"),
		NewBrCode:      1077,
		LastConnection: time.Now(),
	}
	return arg
}

func createTestBranchList(
	t *testing.T,
	req BranchList) error {

	err1 := testQueriesDump.CreateBranchList(context.Background(), req)
	// fmt.Printf("Get by createTestBranchList%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetBranchList(context.Background(), req.BrCode)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.EbSysDate.Format(`2006-01-02`), getData.EbSysDate.Format(`2006-01-02`))
	require.Equal(t, req.RunState, getData.RunState)
	require.Equal(t, req.OrgAddress, getData.OrgAddress)
	require.Equal(t, req.TaxInfo, getData.TaxInfo)
	require.Equal(t, req.DefCity, getData.DefCity)
	require.Equal(t, req.DefProvince, getData.DefProvince)
	require.Equal(t, req.DefCountry, getData.DefCountry)
	require.Equal(t, req.DefZip, getData.DefZip)
	require.Equal(t, req.WaivableInt, getData.WaivableInt)
	require.Equal(t, req.DBVersion, getData.DBVersion)
	require.Equal(t, req.ESystemVer, getData.ESystemVer)
	require.Equal(t, req.NewBrCode, getData.NewBrCode)

	return err2
}

func updateTestBranchList(
	t *testing.T,
	d1 BranchListLight) error {

	err := testQueriesDump.UpdateBranchList(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteBranchList(t *testing.T, brCode string) {
	err := testQueriesDump.DeleteBranchList(context.Background(), brCode)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetBranchList(context.Background(), brCode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
