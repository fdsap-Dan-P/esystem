package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestOfficeAccount(t *testing.T) {

	// Test Data
	d1 := randomOfficeAccount()
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")
	d1.OfficeId = ofc.Id

	d2 := randomOfficeAccount()
	ofc, _ = testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "0000")
	d2.OfficeId = ofc.Id

	// Test Create
	CreatedD1 := createTestOfficeAccount(t, d1)
	CreatedD2 := createTestOfficeAccount(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetOfficeAccount(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.Equal(t, d1.PartitionId, getData1.PartitionId)
	require.Equal(t, d1.Balance.String(), getData1.Balance.String())
	require.Equal(t, d1.PendingTrnAmt.String(), getData1.PendingTrnAmt.String())
	require.Equal(t, d1.Budget.String(), getData1.Budget.String())
	require.Equal(t, d1.LastActivityDate.Time.Format("2006-01-02"), getData1.LastActivityDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetOfficeAccount(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.OfficeId, getData2.OfficeId)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.Currency, getData2.Currency)
	require.Equal(t, d2.PartitionId, getData2.PartitionId)
	require.Equal(t, d2.Balance.String(), getData2.Balance.String())
	require.Equal(t, d2.PendingTrnAmt.String(), getData2.PendingTrnAmt.String())
	require.Equal(t, d2.Budget.String(), getData2.Budget.String())
	require.Equal(t, d2.LastActivityDate.Time.Format("2006-01-02"), getData2.LastActivityDate.Time.Format("2006-01-02"))
	require.Equal(t, d2.StatusId, getData2.StatusId)
	require.Equal(t, d2.Remarks, getData2.Remarks)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetOfficeAccountbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestOfficeAccount(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.OfficeId, updatedD1.OfficeId)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.Currency, updatedD1.Currency)
	require.Equal(t, updateD2.PartitionId, updatedD1.PartitionId)
	require.Equal(t, updateD2.Balance.String(), updatedD1.Balance.String())
	require.Equal(t, updateD2.PendingTrnAmt.String(), updatedD1.PendingTrnAmt.String())
	require.Equal(t, updateD2.Budget.String(), updatedD1.Budget.String())
	require.Equal(t, updateD2.LastActivityDate.Time.Format("2006-01-02"), updatedD1.LastActivityDate.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.StatusId, updatedD1.StatusId)
	require.Equal(t, updateD2.Remarks, updatedD1.Remarks)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListOfficeAccount(t, ListOfficeAccountParams{
		OfficeId: updatedD1.OfficeId,
		Limit:    5,
		Offset:   0,
	})

	// Delete Data
	testDeleteOfficeAccount(t, getData1.Id)
	testDeleteOfficeAccount(t, getData2.Id)
}

func testListOfficeAccount(t *testing.T, arg ListOfficeAccountParams) {

	officeAccount, err := testQueriesAccount.ListOfficeAccount(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", officeAccount)
	require.NotEmpty(t, officeAccount)

}

func randomOfficeAccount() OfficeAccountRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	typ, _ := testQueriesAccount.GetOfficeAccountTypebyName(context.Background(), "Cash")
	part, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "FundSource", 0, "GSB")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "AccountStatus", 0, "Active")

	arg := OfficeAccountRequest{
		// OfficeId:         ofc.Id,
		TypeId:           typ.Id,
		Currency:         util.RandomCurrency(),
		PartitionId:      sql.NullInt64(sql.NullInt64{Int64: part.Id, Valid: true}),
		Balance:          util.RandomMoney(),
		PendingTrnAmt:    util.RandomMoney(),
		Budget:           util.RandomMoney(),
		LastActivityDate: sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		StatusId:         stat.Id,
		Remarks:          sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestOfficeAccount(
	t *testing.T,
	d1 OfficeAccountRequest) model.OfficeAccount {

	getData1, err := testQueriesAccount.CreateOfficeAccount(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.Equal(t, d1.PartitionId, getData1.PartitionId)
	require.Equal(t, d1.Balance.String(), getData1.Balance.String())
	require.Equal(t, d1.PendingTrnAmt.String(), getData1.PendingTrnAmt.String())
	require.Equal(t, d1.Budget.String(), getData1.Budget.String())
	require.Equal(t, d1.LastActivityDate.Time.Format("2006-01-02"), getData1.LastActivityDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestOfficeAccount(
	t *testing.T,
	d1 OfficeAccountRequest) model.OfficeAccount {

	getData1, err := testQueriesAccount.UpdateOfficeAccount(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.Equal(t, d1.PartitionId, getData1.PartitionId)
	require.Equal(t, d1.Balance.String(), getData1.Balance.String())
	require.Equal(t, d1.PendingTrnAmt.String(), getData1.PendingTrnAmt.String())
	require.Equal(t, d1.Budget.String(), getData1.Budget.String())
	require.Equal(t, d1.LastActivityDate.Time.Format("2006-01-02"), getData1.LastActivityDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.StatusId, getData1.StatusId)
	require.Equal(t, d1.Remarks, getData1.Remarks)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteOfficeAccount(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteOfficeAccount(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetOfficeAccount(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
