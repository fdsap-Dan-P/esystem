package db

import (
	"context"
	"database/sql"
	"log"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {

	// Test Data
	d1 := randomUser()
	d1.LoginName = "konek2CARD"

	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "Apps")
	d1.Iiid = ii.Id

	d2 := randomUser()
	ii, _ = testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "101")
	d2.Iiid = ii.Id

	// Test Create
	CreatedD1 := createTestUser(t, d1)
	CreatedD2 := createTestUser(t, d2)

	// Get Data
	getData1, err1 := testQueriesUser.GetUser(context.Background(), CreatedD1.Id)

	log.Println(testQueriesUser.ChangePass(context.Background(), "konek2CARD", "1234"))
	log.Println(testQueriesUser.ChangePass(context.Background(), "olive.mercado0609@gmail.com", "1234"))

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.LoginName, getData1.LoginName)
	require.Equal(t, d1.DisplayName, getData1.DisplayName)
	require.Equal(t, d1.AccessRoleId, getData1.AccessRoleId)
	require.Equal(t, d1.StatusCode, getData1.StatusCode)
	require.Equal(t, d1.DateGiven.Time.Format("2006-01-02"), getData1.DateGiven.Time.Format("2006-01-02"))
	require.Equal(t, d1.DateExpired.Time.Format("2006-01-02"), getData1.DateExpired.Time.Format("2006-01-02"))
	require.Equal(t, d1.DateLocked.Time.Format("2006-01-02"), getData1.DateLocked.Time.Format("2006-01-02"))
	require.Equal(t, d1.PasswordChangedAt.Time.Format("2006-01-02"), getData1.PasswordChangedAt.Time.Format("2006-01-02"))
	require.True(t, getData1.IsCorrectPassword(d1.Password))
	require.Equal(t, d1.Attempt, getData1.Attempt)
	require.Equal(t, d1.Isloggedin, getData1.Isloggedin)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesUser.GetUser(context.Background(), CreatedD2.Id)
	// log.Printf("hashpass: %v", getData2.HashedPassword)
	// require.NotEqual(t, d2.Iiid, getData2.Iiid)

	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.LoginName, getData2.LoginName)
	require.Equal(t, d2.DisplayName, getData2.DisplayName)
	require.Equal(t, d2.AccessRoleId, getData2.AccessRoleId)
	require.Equal(t, d2.StatusCode, getData2.StatusCode)
	require.Equal(t, d2.DateGiven.Time.Format("2006-01-02"), getData2.DateGiven.Time.Format("2006-01-02"))
	require.Equal(t, d2.DateExpired.Time.Format("2006-01-02"), getData2.DateExpired.Time.Format("2006-01-02"))
	require.Equal(t, d2.DateLocked.Time.Format("2006-01-02"), getData2.DateLocked.Time.Format("2006-01-02"))
	require.Equal(t, d2.PasswordChangedAt.Time.Format("2006-01-02"), getData2.PasswordChangedAt.Time.Format("2006-01-02"))
	require.True(t, getData2.IsCorrectPassword(d2.Password))
	require.Equal(t, d2.Attempt, getData2.Attempt)
	require.Equal(t, d2.Isloggedin, getData2.Isloggedin)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesUser.GetUserbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	getData, err = testQueriesUser.GetUserbyName(context.Background(), CreatedD1.LoginName)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestUser(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.LoginName, updatedD1.LoginName)
	require.Equal(t, updateD2.DisplayName, updatedD1.DisplayName)
	require.Equal(t, updateD2.AccessRoleId, updatedD1.AccessRoleId)
	require.Equal(t, updateD2.StatusCode, updatedD1.StatusCode)
	require.Equal(t, updateD2.DateGiven.Time.Format("2006-01-02"), updatedD1.DateGiven.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.DateExpired.Time.Format("2006-01-02"), updatedD1.DateExpired.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.DateLocked.Time.Format("2006-01-02"), updatedD1.DateLocked.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.PasswordChangedAt.Time.Format("2006-01-02"), updatedD1.PasswordChangedAt.Time.Format("2006-01-02"))
	require.True(t, updatedD1.IsCorrectPassword(updateD2.Password))
	require.Equal(t, updateD2.Attempt, updatedD1.Attempt)
	require.Equal(t, updateD2.Isloggedin, updatedD1.Isloggedin)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListUser(t, ListUserParams{
		Iiid:   updatedD1.Iiid,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	// testDeleteUser(t, getData1.Id)
	// testDeleteUser(t, getData2.Id)
}

func testListUser(t *testing.T, arg ListUserParams) {

	user, err := testQueriesUser.ListUser(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", user)
	require.NotEmpty(t, user)

}

func randomUser() UserRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	role, _ := testQueriesAccess.GetAccessRolebyName(context.Background(), "Admin")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "UserStatus", 0, "Active")
	name := util.RandomString(10)
	arg := UserRequest{
		Iiid:      util.RandomInt(1, 100),
		LoginName: name,
		// DisplayName:       sql.NullString(String: util.RandomDate(), Valid: true}),
		AccessRoleId:      role.Id,
		StatusCode:        stat.Code,
		DateGiven:         sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		DateExpired:       sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		DateLocked:        sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		PasswordChangedAt: sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		Password:          util.RandomString(10),
		Attempt:           int16(util.RandomInt32(0, 3)),
		Isloggedin:        sql.NullBool(sql.NullBool{Bool: true, Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	arg.DisplayName.String = name
	arg.DisplayName.Valid = true

	return arg
}

func createTestUser(
	t *testing.T,
	d1 UserRequest) model.Users {

	getData1, err := testQueriesUser.CreateUser(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.LoginName, getData1.LoginName)
	require.Equal(t, d1.DisplayName, getData1.DisplayName)
	require.Equal(t, d1.AccessRoleId, getData1.AccessRoleId)
	require.Equal(t, d1.StatusCode, getData1.StatusCode)
	require.Equal(t, d1.DateGiven.Time.Format("2006-01-02"), getData1.DateGiven.Time.Format("2006-01-02"))
	require.Equal(t, d1.DateExpired.Time.Format("2006-01-02"), getData1.DateExpired.Time.Format("2006-01-02"))
	require.Equal(t, d1.DateLocked.Time.Format("2006-01-02"), getData1.DateLocked.Time.Format("2006-01-02"))
	require.Equal(t, d1.PasswordChangedAt.Time.Format("2006-01-02"), getData1.PasswordChangedAt.Time.Format("2006-01-02"))
	require.True(t, getData1.IsCorrectPassword(d1.Password))
	require.Equal(t, d1.Attempt, getData1.Attempt)
	require.Equal(t, d1.Isloggedin, getData1.Isloggedin)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestUser(
	t *testing.T,
	d1 UserRequest) model.Users {

	getData1, err := testQueriesUser.UpdateUser(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.LoginName, getData1.LoginName)
	require.Equal(t, d1.DisplayName, getData1.DisplayName)
	require.Equal(t, d1.AccessRoleId, getData1.AccessRoleId)
	require.Equal(t, d1.StatusCode, getData1.StatusCode)
	require.Equal(t, d1.DateGiven.Time.Format("2006-01-02"), getData1.DateGiven.Time.Format("2006-01-02"))
	require.Equal(t, d1.DateExpired.Time.Format("2006-01-02"), getData1.DateExpired.Time.Format("2006-01-02"))
	require.Equal(t, d1.DateLocked.Time.Format("2006-01-02"), getData1.DateLocked.Time.Format("2006-01-02"))
	require.Equal(t, d1.PasswordChangedAt.Time.Format("2006-01-02"), getData1.PasswordChangedAt.Time.Format("2006-01-02"))
	require.True(t, getData1.IsCorrectPassword(d1.Password))
	require.Equal(t, d1.Attempt, getData1.Attempt)
	require.Equal(t, d1.Isloggedin, getData1.Isloggedin)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteUser(t *testing.T, id int64) {
	err := testQueriesUser.DeleteUser(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesUser.GetUser(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
