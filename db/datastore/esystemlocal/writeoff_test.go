package db

import (
	"context"
	"database/sql"
	"simplebank/util"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteoff(t *testing.T) {

	// Test Data
	d1 := randomWriteoff()
	d2 := randomWriteoff()
	d2.Acc = "01C4-4001-004894707"

	err := createTestWriteoff(t, d1)
	require.NoError(t, err)

	err = createTestWriteoff(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetWriteoff(context.Background(), d1.Acc)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Acc, getData1.Acc)
	require.Equal(t, d1.DisbDate.Format(`2006-01-02`), getData1.DisbDate.Format(`2006-01-02`))
	require.True(t, d1.Principal.Equal(getData1.Principal))
	require.True(t, d1.Interest.Equal(getData1.Interest))
	require.True(t, d1.BalPrin.Equal(getData1.BalPrin))
	require.True(t, d1.BalInt.Equal(getData1.BalInt))
	require.Equal(t, d1.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d1.AcctType, getData1.AcctType)
	require.Equal(t, d1.Print, getData1.Print)
	require.Equal(t, d1.PostedBy, getData1.PostedBy)
	require.Equal(t, d1.VerifiedBy, getData1.VerifiedBy)

	getData2, err2 := testQueriesLocal.GetWriteoff(context.Background(), d2.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Acc, getData2.Acc)
	require.Equal(t, d2.DisbDate.Format(`2006-01-02`), getData2.DisbDate.Format(`2006-01-02`))
	require.True(t, d2.Principal.Equal(getData2.Principal))
	require.True(t, d2.Interest.Equal(getData2.Interest))
	require.True(t, d2.BalPrin.Equal(getData2.BalPrin))
	require.True(t, d2.BalInt.Equal(getData2.BalInt))
	require.Equal(t, d2.TrnDate.Format(`2006-01-02`), getData2.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d2.AcctType, getData2.AcctType)
	require.Equal(t, d2.Print, getData2.Print)
	require.Equal(t, d2.PostedBy, getData2.PostedBy)
	require.Equal(t, d2.VerifiedBy, getData2.VerifiedBy)

	// Update Data
	updateD2 := d2
	updateD2.Acc = getData2.Acc
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestWriteoff(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesLocal.GetWriteoff(context.Background(), updateD2.Acc)
	require.NoError(t, err1)

	require.Equal(t, updateD2.Acc, getData1.Acc)
	require.Equal(t, updateD2.DisbDate.Format(`2006-01-02`), getData1.DisbDate.Format(`2006-01-02`))
	require.True(t, updateD2.Principal.Equal(getData1.Principal))
	require.True(t, updateD2.Interest.Equal(getData1.Interest))
	require.True(t, updateD2.BalPrin.Equal(getData1.BalPrin))
	require.True(t, updateD2.BalInt.Equal(getData1.BalInt))
	require.Equal(t, updateD2.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, updateD2.AcctType, getData1.AcctType)
	require.Equal(t, updateD2.Print, getData1.Print)
	require.Equal(t, updateD2.PostedBy, getData1.PostedBy)
	require.Equal(t, updateD2.VerifiedBy, getData1.VerifiedBy)

	testListWriteoff(t, ListWriteoffParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteWriteoff(t, d1.Acc)
	testDeleteWriteoff(t, d2.Acc)
}

func testListWriteoff(t *testing.T, arg ListWriteoffParams) {

	Writeoff, err := testQueriesLocal.ListWriteoff(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Writeoff)
	require.NotEmpty(t, Writeoff)

}

func randomWriteoff() WriteoffRequest {

	arg := WriteoffRequest{
		Acc:        "01C4-4001-004894703",
		DisbDate:   util.RandomDate(),
		Principal:  util.RandomMoney(),
		Interest:   util.RandomMoney(),
		BalPrin:    util.RandomMoney(),
		BalInt:     util.RandomMoney(),
		TrnDate:    util.RandomDate(),
		AcctType:   "301",
		Print:      sql.NullString{String: "Y", Valid: true},
		PostedBy:   sql.NullString{String: "dsdff", Valid: true},
		VerifiedBy: sql.NullString{String: "dsdff", Valid: true},
	}
	return arg
}

func createTestWriteoff(
	t *testing.T,
	req WriteoffRequest) error {

	err1 := testQueriesLocal.CreateWriteoff(context.Background(), req)
	// fmt.Printf("Get by createTestWriteoff%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetWriteoff(context.Background(), req.Acc)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.Acc, getData.Acc)
	require.Equal(t, req.DisbDate.Format(`2006-01-02`), getData.DisbDate.Format(`2006-01-02`))
	require.True(t, req.Principal.Equal(getData.Principal))
	require.True(t, req.Interest.Equal(getData.Interest))
	require.True(t, req.BalPrin.Equal(getData.BalPrin))
	require.True(t, req.BalInt.Equal(getData.BalInt))
	require.Equal(t, req.TrnDate.Format(`2006-01-02`), getData.TrnDate.Format(`2006-01-02`))
	require.Equal(t, req.AcctType, getData.AcctType)
	require.Equal(t, req.Print, getData.Print)
	require.Equal(t, req.PostedBy, getData.PostedBy)
	require.Equal(t, req.VerifiedBy, getData.VerifiedBy)

	return err2
}

func updateTestWriteoff(
	t *testing.T,
	d1 WriteoffRequest) error {

	err := testQueriesLocal.UpdateWriteoff(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteWriteoff(t *testing.T, Acc string) {
	err := testQueriesLocal.DeleteWriteoff(context.Background(), Acc)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetWriteoff(context.Background(), Acc)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
