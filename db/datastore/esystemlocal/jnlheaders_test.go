package db

import (
	"context"
	"database/sql"
	"simplebank/util"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestJnlHeaders(t *testing.T) {

	// Test Data
	d1 := randomJnlHeaders()
	d2 := randomJnlHeaders()
	d2.Trn = d2.Trn + "-1"

	err := createTestJnlHeaders(t, d1)
	require.NoError(t, err)

	err = createTestJnlHeaders(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetJnlHeaders(context.Background(), d1.Trn)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Trn, getData1.Trn)
	require.Equal(t, d1.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d1.Particulars, getData1.Particulars)
	require.Equal(t, d1.UserName, getData1.UserName)
	require.Equal(t, d1.Code, getData1.Code)

	getData2, err2 := testQueriesLocal.GetJnlHeaders(context.Background(), d2.Trn)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Trn, getData2.Trn)
	require.Equal(t, d2.TrnDate.Format(`2006-01-02`), getData2.TrnDate.Format(`2006-01-02`))
	require.Equal(t, d2.Particulars, getData2.Particulars)
	require.Equal(t, d2.UserName, getData2.UserName)
	require.Equal(t, d2.Code, getData2.Code)

	// Update Data
	updateD2 := d2
	updateD2.Trn = getData2.Trn
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestJnlHeaders(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesLocal.GetJnlHeaders(context.Background(), updateD2.Trn)
	require.NoError(t, err1)

	require.Equal(t, updateD2.Trn, getData1.Trn)
	require.Equal(t, updateD2.TrnDate.Format(`2006-01-02`), getData1.TrnDate.Format(`2006-01-02`))
	require.Equal(t, updateD2.Particulars, getData1.Particulars)
	require.Equal(t, updateD2.UserName, getData1.UserName)
	require.Equal(t, updateD2.Code, getData1.Code)

	testListJnlHeaders(t, ListJnlHeadersParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteJnlHeaders(t, d1.Trn)
	testDeleteJnlHeaders(t, d2.Trn)
}

func testListJnlHeaders(t *testing.T, arg ListJnlHeadersParams) {

	JnlHeaders, err := testQueriesLocal.ListJnlHeaders(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", JnlHeaders)
	require.NotEmpty(t, JnlHeaders)

}

func randomJnlHeaders() JnlHeadersRequest {

	arg := JnlHeadersRequest{
		Trn:         util.RandomString(10),
		TrnDate:     util.RandomDate(),
		Particulars: util.RandomString(10),
		UserName:    util.SetNullString(util.RandomString(10)),
		Code:        1,
	}
	return arg
}

func createTestJnlHeaders(
	t *testing.T,
	req JnlHeadersRequest) error {

	err1 := testQueriesLocal.CreateJnlHeaders(context.Background(), req)
	// fmt.Printf("Get by createTestJnlHeaders%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesLocal.GetJnlHeaders(context.Background(), req.Trn)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.Trn, getData.Trn)
	require.Equal(t, req.TrnDate.Format(`2006-01-02`), getData.TrnDate.Format(`2006-01-02`))
	require.Equal(t, req.Particulars, getData.Particulars)
	require.Equal(t, req.UserName, getData.UserName)
	require.Equal(t, req.Code, getData.Code)

	return err2
}

func updateTestJnlHeaders(
	t *testing.T,
	d1 JnlHeadersRequest) error {

	err := testQueriesLocal.UpdateJnlHeaders(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteJnlHeaders(t *testing.T, trn string) {
	err := testQueriesLocal.DeleteJnlHeaders(context.Background(), trn)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetJnlHeaders(context.Background(), trn)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
