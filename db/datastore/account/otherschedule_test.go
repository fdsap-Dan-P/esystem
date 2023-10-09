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

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestOtherSchedule(t *testing.T) {

	// Test Data
	d1 := randomOtherSchedule()
	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")
	d1.AccountId = acc.Id

	d2 := randomOtherSchedule()
	acc, _ = testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000002")
	d2.AccountId = acc.Id

	// Test Create
	CreatedD1 := createTestOtherSchedule(t, d1)
	CreatedD2 := createTestOtherSchedule(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetOtherSchedule(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.ChargeId, getData1.ChargeId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.DueDate.Format("2006-01-02"), getData1.DueDate.Format("2006-01-02"))
	require.Equal(t, d1.DueAmt.String(), getData1.DueAmt.String())
	require.Equal(t, d1.Realizable.String(), getData1.Realizable.String())
	require.Equal(t, d1.EndBal.String(), getData1.EndBal.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetOtherSchedule(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AccountId, getData2.AccountId)
	require.Equal(t, d2.ChargeId, getData2.ChargeId)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.DueDate.Format("2006-01-02"), getData2.DueDate.Format("2006-01-02"))
	require.Equal(t, d2.DueAmt.String(), getData2.DueAmt.String())
	require.Equal(t, d2.Realizable.String(), getData2.Realizable.String())
	require.Equal(t, d2.EndBal.String(), getData2.EndBal.String())
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetOtherSchedulebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestOtherSchedule(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.AccountId, updatedD1.AccountId)
	require.Equal(t, updateD2.ChargeId, updatedD1.ChargeId)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.DueDate.Format("2006-01-02"), updatedD1.DueDate.Format("2006-01-02"))
	require.Equal(t, updateD2.DueAmt.String(), updatedD1.DueAmt.String())
	require.Equal(t, updateD2.Realizable.String(), updatedD1.Realizable.String())
	require.Equal(t, updateD2.EndBal.String(), updatedD1.EndBal.String())
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListOtherSchedule(t, ListOtherScheduleParams{
		AccountId: updatedD1.AccountId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteOtherSchedule(t, CreatedD1.Uuid)
	testDeleteOtherSchedule(t, CreatedD2.Uuid)
}

func testListOtherSchedule(t *testing.T, arg ListOtherScheduleParams) {

	otherSchedule, err := testQueriesAccount.ListOtherSchedule(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", otherSchedule)
	require.NotEmpty(t, otherSchedule)

}

func randomOtherSchedule() OtherScheduleRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	chrg, _ := testQueriesAccount.GetChargeTypebyName(context.Background(), "Service Fee")

	arg := OtherScheduleRequest{
		// AccountId:  util.RandomInt(1, 100),
		ChargeId:   chrg.Id,
		Series:     int16(util.RandomInt32(1, 100)),
		DueDate:    util.RandomDate(),
		DueAmt:     util.RandomMoney(),
		Realizable: util.RandomMoney(),
		EndBal:     util.RandomMoney(),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestOtherSchedule(
	t *testing.T,
	d1 OtherScheduleRequest) model.OtherSchedule {

	getData1, err := testQueriesAccount.CreateOtherSchedule(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.ChargeId, getData1.ChargeId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.DueDate.Format("2006-01-02"), getData1.DueDate.Format("2006-01-02"))
	require.Equal(t, d1.DueAmt.String(), getData1.DueAmt.String())
	require.Equal(t, d1.Realizable.String(), getData1.Realizable.String())
	require.Equal(t, d1.EndBal.String(), getData1.EndBal.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestOtherSchedule(
	t *testing.T,
	d1 OtherScheduleRequest) model.OtherSchedule {

	getData1, err := testQueriesAccount.UpdateOtherSchedule(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.ChargeId, getData1.ChargeId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.DueDate.Format("2006-01-02"), getData1.DueDate.Format("2006-01-02"))
	require.Equal(t, d1.DueAmt.String(), getData1.DueAmt.String())
	require.Equal(t, d1.Realizable.String(), getData1.Realizable.String())
	require.Equal(t, d1.EndBal.String(), getData1.EndBal.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteOtherSchedule(t *testing.T, uuid uuid.UUID) {
	err := testQueriesAccount.DeleteOtherSchedule(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetOtherSchedule(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
