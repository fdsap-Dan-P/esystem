package db

import (
	"context"
	"database/sql"
	model "simplebank/db/datastore/esystemlocal"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestReferencesDetails(t *testing.T) {

	// Test Data
	d1 := randomReferencesDetails()
	d2 := randomReferencesDetails()
	d2.ID = 6014
	d2.RefID = 1007

	err := createTestReferencesDetails(t, d1)
	require.NoError(t, err)

	err = createTestReferencesDetails(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetReferencesDetails(context.Background(), d1.BrCode, d1.ID)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.ID, getData1.ID)
	require.Equal(t, d1.RefID, getData1.RefID)
	require.Equal(t, d1.PurposeDescription, getData1.PurposeDescription)
	require.Equal(t, d1.ParentID, getData1.ParentID)
	require.Equal(t, d1.CodeID, getData1.CodeID)
	require.Equal(t, d1.Stat, getData1.Stat)

	getData2, err2 := testQueriesDump.GetReferencesDetails(context.Background(), d2.BrCode, d2.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.ID, getData2.ID)
	require.Equal(t, d2.RefID, getData2.RefID)
	require.Equal(t, d2.PurposeDescription, getData2.PurposeDescription)
	require.Equal(t, d2.ParentID, getData2.ParentID)
	require.Equal(t, d2.CodeID, getData2.CodeID)
	require.Equal(t, d2.Stat, getData2.Stat)

	// Update Data
	updateD2 := d2
	updateD2.ID = getData2.ID
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestReferencesDetails(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetReferencesDetails(context.Background(), updateD2.BrCode, updateD2.ID)
	require.NoError(t, err1)

	require.Equal(t, updateD2.ID, getData1.ID)
	require.Equal(t, updateD2.RefID, getData1.RefID)
	require.Equal(t, updateD2.PurposeDescription, getData1.PurposeDescription)
	require.Equal(t, updateD2.ParentID, getData1.ParentID)
	require.Equal(t, updateD2.CodeID, getData1.CodeID)
	require.Equal(t, updateD2.Stat, getData1.Stat)

	testListReferencesDetails(t, ListReferencesDetailsParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteReferencesDetails(t, d1.BrCode, d1.ID)
	testDeleteReferencesDetails(t, d2.BrCode, d2.ID)
}

func testListReferencesDetails(t *testing.T, arg ListReferencesDetailsParams) {

	ReferencesDetails, err := testQueriesDump.ListReferencesDetails(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", ReferencesDetails)
	require.NotEmpty(t, ReferencesDetails)

}

func randomReferencesDetails() model.ReferencesDetails {

	arg := model.ReferencesDetails{
		ModCtr:             1,
		BrCode:             "01",
		ID:                 6013,
		RefID:              1007,
		PurposeDescription: sql.NullString{String: "dsdff", Valid: true},
		ParentID:           sql.NullInt64{Int64: 100, Valid: true},
		CodeID:             sql.NullInt64{Int64: 100, Valid: true},
		Stat:               sql.NullInt64{Int64: 100, Valid: true},
	}
	return arg
}

func createTestReferencesDetails(
	t *testing.T,
	req model.ReferencesDetails) error {

	err1 := testQueriesDump.CreateReferencesDetails(context.Background(), req)
	// fmt.Printf("Get by createTestReferencesDetails%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetReferencesDetails(context.Background(), req.BrCode, req.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.ID, getData.ID)
	require.Equal(t, req.RefID, getData.RefID)
	require.Equal(t, req.PurposeDescription, getData.PurposeDescription)
	require.Equal(t, req.ParentID, getData.ParentID)
	require.Equal(t, req.CodeID, getData.CodeID)
	require.Equal(t, req.Stat, getData.Stat)

	return err2
}

func updateTestReferencesDetails(
	t *testing.T,
	d1 model.ReferencesDetails) error {

	err := testQueriesDump.UpdateReferencesDetails(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteReferencesDetails(t *testing.T, brCode string, ID int64) {
	err := testQueriesDump.DeleteReferencesDetails(context.Background(), brCode, ID)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetReferencesDetails(context.Background(), brCode, ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
