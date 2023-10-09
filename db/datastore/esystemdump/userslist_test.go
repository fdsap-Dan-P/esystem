package db

import (
	"context"
	"database/sql"
	"log"
	model "simplebank/db/datastore/esystemlocal"
	"simplebank/util"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestUsersList(t *testing.T) {

	// Test Data
	d1 := randomUsersList()
	d2 := randomUsersList()
	d2.UserId = d2.UserId + "1"

	log.Printf("d1: %v", d1)
	log.Printf("d2: %v", d2)

	err := createTestUsersList(t, d1)
	require.NoError(t, err)

	err = createTestUsersList(t, d2)
	require.NoError(t, err)

	// Get Data
	getData1, err1 := testQueriesDump.GetUsersList(context.Background(), d1.BrCode, d1.UserId)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AccessCode, getData1.AccessCode)
	require.Equal(t, d1.LName, getData1.LName)
	require.Equal(t, d1.FName, getData1.FName)
	require.Equal(t, d1.MName, getData1.MName)
	require.Equal(t, d1.DateHired.Time.Format(`2006-01-02`), getData1.DateHired.Time.Format(`2006-01-02`))
	require.Equal(t, d1.BirthDay.Time.Format(`2006-01-02`), getData1.BirthDay.Time.Format(`2006-01-02`))
	require.Equal(t, d1.DateGiven.Time.Format(`2006-01-02`), getData1.DateGiven.Time.Format(`2006-01-02`))
	require.Equal(t, d1.DateExpired.Time.Format(`2006-01-02`), getData1.DateExpired.Time.Format(`2006-01-02`))
	require.Equal(t, d1.Address, getData1.Address)
	require.Equal(t, d1.Position, getData1.Position)
	require.Equal(t, d1.AreaCode, getData1.AreaCode)
	require.Equal(t, d1.ManCode, getData1.ManCode)
	require.Equal(t, d1.AddInfo, getData1.AddInfo)
	require.Equal(t, d1.Passwd, getData1.Passwd)
	require.Equal(t, d1.Attempt, getData1.Attempt)
	require.Equal(t, d1.DateLocked.Time.Format(`2006-01-02`), getData1.DateLocked.Time.Format(`2006-01-02`))
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.Equal(t, d1.Picture, getData1.Picture)
	require.Equal(t, d1.IsLoggedIn, getData1.IsLoggedIn)
	require.Equal(t, d1.AccountExpirationDt.Time.Format(`2006-01-02`), getData1.AccountExpirationDt.Time.Format(`2006-01-02`))

	getData2, err2 := testQueriesDump.GetUsersList(context.Background(), d2.BrCode, d2.UserId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AccessCode, getData2.AccessCode)
	require.Equal(t, d2.LName, getData2.LName)
	require.Equal(t, d2.FName, getData2.FName)
	require.Equal(t, d2.MName, getData2.MName)
	require.Equal(t, d2.DateHired.Time.Format(`2006-01-02`), getData2.DateHired.Time.Format(`2006-01-02`))
	require.Equal(t, d2.BirthDay.Time.Format(`2006-01-02`), getData2.BirthDay.Time.Format(`2006-01-02`))
	require.Equal(t, d2.DateGiven.Time.Format(`2006-01-02`), getData2.DateGiven.Time.Format(`2006-01-02`))
	require.Equal(t, d2.DateExpired.Time.Format(`2006-01-02`), getData2.DateExpired.Time.Format(`2006-01-02`))
	require.Equal(t, d2.Address, getData2.Address)
	require.Equal(t, d2.Position, getData2.Position)
	require.Equal(t, d2.AreaCode, getData2.AreaCode)
	require.Equal(t, d2.ManCode, getData2.ManCode)
	require.Equal(t, d2.AddInfo, getData2.AddInfo)
	require.Equal(t, d2.Passwd, getData2.Passwd)
	require.Equal(t, d2.Attempt, getData2.Attempt)
	require.Equal(t, d2.DateLocked.Time.Format(`2006-01-02`), getData2.DateLocked.Time.Format(`2006-01-02`))
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.Equal(t, d2.Picture, getData2.Picture)
	require.Equal(t, d2.IsLoggedIn, getData2.IsLoggedIn)
	require.Equal(t, d2.AccountExpirationDt.Time.Format(`2006-01-02`), getData2.AccountExpirationDt.Time.Format(`2006-01-02`))

	// Update Data
	updateD2 := d2
	updateD2.Remarks.String = updateD2.Remarks.String + "Edited"

	// log.Println(updateD2)
	err3 := updateTestUsersList(t, updateD2)
	require.NoError(t, err3)

	getData1, err1 = testQueriesDump.GetUsersList(context.Background(), updateD2.BrCode, updateD2.UserId)
	require.NoError(t, err1)

	require.Equal(t, updateD2.AccessCode, getData1.AccessCode)
	require.Equal(t, updateD2.LName, getData1.LName)
	require.Equal(t, updateD2.FName, getData1.FName)
	require.Equal(t, updateD2.MName, getData1.MName)
	require.Equal(t, updateD2.DateHired.Time.Format(`2006-01-02`), getData1.DateHired.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.BirthDay.Time.Format(`2006-01-02`), getData1.BirthDay.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.DateGiven.Time.Format(`2006-01-02`), getData1.DateGiven.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.DateExpired.Time.Format(`2006-01-02`), getData1.DateExpired.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.Address, getData1.Address)
	require.Equal(t, updateD2.Position, getData1.Position)
	require.Equal(t, updateD2.AreaCode, getData1.AreaCode)
	require.Equal(t, updateD2.ManCode, getData1.ManCode)
	require.Equal(t, updateD2.AddInfo, getData1.AddInfo)
	require.Equal(t, updateD2.Passwd, getData1.Passwd)
	require.Equal(t, updateD2.Attempt, getData1.Attempt)
	require.Equal(t, updateD2.DateLocked.Time.Format(`2006-01-02`), getData1.DateLocked.Time.Format(`2006-01-02`))
	require.Equal(t, updateD2.Remarks, getData1.Remarks)
	require.Equal(t, updateD2.Picture, getData1.Picture)
	require.Equal(t, updateD2.IsLoggedIn, getData1.IsLoggedIn)
	require.Equal(t, updateD2.AccountExpirationDt.Time.Format(`2006-01-02`), getData1.AccountExpirationDt.Time.Format(`2006-01-02`))

	testListUsersList(t, ListUsersListParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteUsersList(t, d1.BrCode, d1.UserId)
	testDeleteUsersList(t, d2.BrCode, d2.UserId)
}

func testListUsersList(t *testing.T, arg ListUsersListParams) {

	UsersList, err := testQueriesDump.ListUsersList(context.Background(), int64(0))
	require.NoError(t, err)
	// fmt.Printf("%+v\n", UsersList)
	require.NotEmpty(t, UsersList)

}

func randomUsersList() model.UsersList {

	arg := model.UsersList{
		ModCtr:              1,
		BrCode:              "01",
		UserId:              "User",
		AccessCode:          util.SetNullInt64(0),
		LName:               "lName",
		FName:               "fName",
		MName:               "M",
		DateHired:           util.RandomNullDate(),
		BirthDay:            util.RandomNullDate(),
		DateGiven:           util.RandomNullDate(),
		DateExpired:         util.RandomNullDate(),
		Address:             util.SetNullString("Address"),
		Position:            util.SetNullString("Position"),
		AreaCode:            util.SetNullInt64(101),
		ManCode:             util.SetNullInt64(10),
		AddInfo:             util.SetNullString("AddInfo"),
		Passwd:              []byte{},
		Attempt:             util.SetNullInt64(0),
		DateLocked:          util.RandomNullDate(),
		Remarks:             util.SetNullString("Remarks"),
		Picture:             []byte{},
		IsLoggedIn:          false,
		AccountExpirationDt: util.RandomNullDate(),
	}
	return arg
}

func createTestUsersList(
	t *testing.T,
	req model.UsersList) error {

	err1 := testQueriesDump.CreateUsersList(context.Background(), req)
	// fmt.Printf("Get by createTestUsersList%+v\n", getData1)
	require.NoError(t, err1)

	if err1 != nil {
		return err1
	}

	getData, err2 := testQueriesDump.GetUsersList(context.Background(), req.BrCode, req.UserId)
	require.NoError(t, err2)
	require.NotEmpty(t, getData)
	require.Equal(t, req.AccessCode, getData.AccessCode)
	require.Equal(t, req.LName, getData.LName)
	require.Equal(t, req.FName, getData.FName)
	require.Equal(t, req.MName, getData.MName)
	require.Equal(t, req.DateHired.Time.Format(`2006-01-02`), getData.DateHired.Time.Format(`2006-01-02`))
	require.Equal(t, req.BirthDay.Time.Format(`2006-01-02`), getData.BirthDay.Time.Format(`2006-01-02`))
	require.Equal(t, req.DateGiven.Time.Format(`2006-01-02`), getData.DateGiven.Time.Format(`2006-01-02`))
	require.Equal(t, req.DateExpired.Time.Format(`2006-01-02`), getData.DateExpired.Time.Format(`2006-01-02`))
	require.Equal(t, req.Address, getData.Address)
	require.Equal(t, req.Position, getData.Position)
	require.Equal(t, req.AreaCode, getData.AreaCode)
	require.Equal(t, req.ManCode, getData.ManCode)
	require.Equal(t, req.AddInfo, getData.AddInfo)
	require.Equal(t, req.Passwd, getData.Passwd)
	require.Equal(t, req.Attempt, getData.Attempt)
	require.Equal(t, req.DateLocked.Time.Format(`2006-01-02`), getData.DateLocked.Time.Format(`2006-01-02`))
	require.Equal(t, req.Remarks, getData.Remarks)
	require.Equal(t, req.Picture, getData.Picture)
	require.Equal(t, req.IsLoggedIn, getData.IsLoggedIn)
	require.Equal(t, req.AccountExpirationDt.Time.Format(`2006-01-02`), getData.AccountExpirationDt.Time.Format(`2006-01-02`))

	return err2
}

func updateTestUsersList(
	t *testing.T,
	d1 model.UsersList) error {

	err := testQueriesDump.UpdateUsersList(context.Background(), d1)
	require.NoError(t, err)

	return err
}

func testDeleteUsersList(t *testing.T, brCode string, userId string) {
	err := testQueriesDump.DeleteUsersList(context.Background(), brCode, userId)
	require.NoError(t, err)

	ref1, err := testQueriesDump.GetUsersList(context.Background(), brCode, userId)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
