package db

import (
	"context"
	"database/sql"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddresses(t *testing.T) {

	// Test Data
	d1 := randomAddresses()
	d2 := randomAddresses()
	d2.SeqNum = d2.SeqNum + 1

	id, err := createTestAddresses(t, d1)
	d1.SeqNum = id
	require.NoError(t, err)

	id, err = createTestAddresses(t, d2)
	d2.SeqNum = id
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesLocal.GetAddresses(context.Background(), d1.SeqNum)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CID, getData1.CID)
	require.Equal(t, d1.SeqNum, getData1.SeqNum)
	require.Equal(t, d1.AddressDetails, getData1.AddressDetails)
	require.Equal(t, d1.Barangay, getData1.Barangay)
	require.Equal(t, d1.City, getData1.City)
	require.Equal(t, d1.Province, getData1.Province)
	require.Equal(t, d1.Phone1, getData1.Phone1)
	require.Equal(t, d1.Phone2, getData1.Phone2)
	require.Equal(t, d1.Phone3, getData1.Phone3)
	require.Equal(t, d1.Phone4, getData1.Phone4)

	getData2, err2 := testQueriesLocal.GetAddresses(context.Background(), d2.SeqNum)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CID, getData2.CID)
	require.Equal(t, d2.SeqNum, getData2.SeqNum)
	require.Equal(t, d2.AddressDetails, getData2.AddressDetails)
	require.Equal(t, d2.Barangay, getData2.Barangay)
	require.Equal(t, d2.City, getData2.City)
	require.Equal(t, d2.Province, getData2.Province)
	require.Equal(t, d2.Phone1, getData2.Phone1)
	require.Equal(t, d2.Phone2, getData2.Phone2)
	require.Equal(t, d2.Phone3, getData2.Phone3)
	require.Equal(t, d2.Phone4, getData2.Phone4)

	// Update Data
	updateD2 := d2
	updateD2.SeqNum = getData2.SeqNum
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	err3 := updateTestAddresses(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesLocal.GetAddresses(context.Background(), updateD2.SeqNum)
	require.NoError(t, err1)

	require.Equal(t, updateD2.CID, getData1.CID)
	require.Equal(t, updateD2.SeqNum, getData1.SeqNum)
	require.Equal(t, updateD2.AddressDetails, getData1.AddressDetails)
	require.Equal(t, updateD2.Barangay, getData1.Barangay)
	require.Equal(t, updateD2.City, getData1.City)
	require.Equal(t, updateD2.Province, getData1.Province)
	require.Equal(t, updateD2.Phone1, getData1.Phone1)
	require.Equal(t, updateD2.Phone2, getData1.Phone2)
	require.Equal(t, updateD2.Phone3, getData1.Phone3)
	require.Equal(t, updateD2.Phone4, getData1.Phone4)

	testListAddresses(t, ListAddressesParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteAddresses(t, d1.SeqNum)
	testDeleteAddresses(t, d2.SeqNum)
}

func testListAddresses(t *testing.T, arg ListAddressesParams) {

	Addresses, err := testQueriesLocal.ListAddresses(context.Background())
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Addresses)
	require.NotEmpty(t, Addresses)

}

func randomAddresses() AddressesRequest {

	arg := AddressesRequest{
		CID:            19858200,
		AddressDetails: sql.NullString{String: "dsdff", Valid: true},
		Barangay:       sql.NullString{String: "dsdff", Valid: true},
		City:           sql.NullString{String: "dsdff", Valid: true},
		Province:       sql.NullString{String: "dsdff", Valid: true},
		Phone1:         sql.NullString{String: "dsdff", Valid: true},
		Phone2:         sql.NullString{String: "dsdff", Valid: true},
		Phone3:         sql.NullString{String: "dsdff", Valid: true},
		Phone4:         sql.NullString{String: "dsdff", Valid: true}}
	return arg
}

func createTestAddresses(
	t *testing.T,
	req AddressesRequest) (int64, error) {

	id, err1 := testQueriesLocal.CreateAddresses(context.Background(), req)
	// fmt.Printf("Get by createTestAddresses%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return id, err1
	}

	getData, err2 := testQueriesLocal.GetAddresses(context.Background(), id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.CID, getData.CID)
	require.Equal(t, req.AddressDetails, getData.AddressDetails)
	require.Equal(t, req.Barangay, getData.Barangay)
	require.Equal(t, req.City, getData.City)
	require.Equal(t, req.Province, getData.Province)
	require.Equal(t, req.Phone1, getData.Phone1)
	require.Equal(t, req.Phone2, getData.Phone2)
	require.Equal(t, req.Phone3, getData.Phone3)
	require.Equal(t, req.Phone4, getData.Phone4)

	return id, err2
}

func updateTestAddresses(
	t *testing.T,
	d1 AddressesRequest) error {

	err := testQueriesLocal.UpdateAddresses(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteAddresses(t *testing.T, SeqNum int64) {
	err := testQueriesLocal.DeleteAddresses(context.Background(), SeqNum)
	require.NoError(t, err)

	ref1, err := testQueriesLocal.GetAddresses(context.Background(), SeqNum)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
