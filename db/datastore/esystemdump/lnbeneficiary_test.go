package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"
	"simplebank/util"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestLnBeneficiary(t *testing.T) {

	// Test Data
	d1 := randomLnBeneficiary()
	d2 := randomLnBeneficiary()

	err := createTestLnBeneficiary(t, d1)
	require.NoError(t, err)

	err = createTestLnBeneficiary(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetLnBeneficiary(context.Background(), d1.BrCode, d1.Acc)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.BDay.Format(`2006-01-02`), getData1.BDay.Format(`2006-01-02`))
	require.Equal(t, d1.EducLvl, getData1.EducLvl)
	require.Equal(t, d1.Gender, getData1.Gender)
	require.Equal(t, d1.LastName, getData1.LastName)
	require.Equal(t, d1.FirstName, getData1.FirstName)
	require.Equal(t, d1.MiddleName, getData1.MiddleName)
	require.Equal(t, d1.Remarks, getData1.Remarks)

	getData2, err2 := testQueriesDump.GetLnBeneficiary(context.Background(), d2.BrCode, d2.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.BDay.Format(`2006-01-02`), getData2.BDay.Format(`2006-01-02`))
	require.Equal(t, d2.EducLvl, getData2.EducLvl)
	require.Equal(t, d2.Gender, getData2.Gender)
	require.Equal(t, d2.LastName, getData2.LastName)
	require.Equal(t, d2.FirstName, getData2.FirstName)
	require.Equal(t, d2.MiddleName, getData2.MiddleName)
	require.Equal(t, d2.Remarks, getData2.Remarks)

	// Update Data
	updateD2 := d2
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestLnBeneficiary(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetLnBeneficiary(context.Background(), updateD2.BrCode, updateD2.Acc)
	require.NoError(t, err1)

	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.Equal(t, updateD2.BDay.Format(`2006-01-02`), getData1.BDay.Format(`2006-01-02`))
	require.Equal(t, updateD2.EducLvl, getData1.EducLvl)
	require.Equal(t, updateD2.Gender, getData1.Gender)
	require.Equal(t, updateD2.LastName, getData1.LastName)
	require.Equal(t, updateD2.FirstName, getData1.FirstName)
	require.Equal(t, updateD2.MiddleName, getData1.MiddleName)
	require.Equal(t, updateD2.Remarks, getData1.Remarks)
	testListLnBeneficiary(t, ListLnBeneficiaryParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteLnBeneficiary(t, d1.BrCode, d1.Acc)
	testDeleteLnBeneficiary(t, d2.BrCode, d2.Acc)
}

func testListLnBeneficiary(t *testing.T, arg ListLnBeneficiaryParams) {

	LnBeneficiary, err := testQueriesDump.ListLnBeneficiary(context.Background(), 0)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", LnBeneficiary)
	require.NotEmpty(t, LnBeneficiary)

}

func randomLnBeneficiary() model.LnBeneficiary {

	arg := model.LnBeneficiary{
		ModCtr:     1,
		BrCode:     "E3",
		Acc:        "0101-4041-0157460",
		BDay:       util.SetDate("2021-12-31"),
		EducLvl:    "GRD10",
		Gender:     true,
		LastName:   util.SetNullString("Mercado"),
		FirstName:  util.SetNullString("Roderick"),
		MiddleName: util.SetNullString("G"),
		Remarks:    util.SetNullString(""),
	}
	return arg
}

func createTestLnBeneficiary(
	t *testing.T,
	req model.LnBeneficiary) error {

	err1 := testQueriesDump.CreateLnBeneficiary(context.Background(), req)
	// fmt.Printf("Get by createTestLnBeneficiary%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetLnBeneficiary(context.Background(), req.BrCode, req.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.Acc, getData.Acc)
	require.Equal(t, req.BDay.Format(`2006-01-02`), getData.BDay.Format(`2006-01-02`))
	require.Equal(t, req.EducLvl, getData.EducLvl)
	require.Equal(t, req.Gender, getData.Gender)
	require.Equal(t, req.LastName, getData.LastName)
	require.Equal(t, req.FirstName, getData.FirstName)
	require.Equal(t, req.MiddleName, getData.MiddleName)
	require.Equal(t, req.Remarks, getData.Remarks)

	return err2
}

func updateTestLnBeneficiary(
	t *testing.T,
	d1 model.LnBeneficiary) error {

	err := testQueriesDump.UpdateLnBeneficiary(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteLnBeneficiary(t *testing.T, brCode string, acc string) {
	err := testQueriesDump.DeleteLnBeneficiary(context.Background(), brCode, acc)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetLnBeneficiary(context.Background(), brCode, acc)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
