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

func TestSchedule(t *testing.T) {

	// Test Data
	d1 := randomSchedule()
	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")
	d1.AccountId = acc.Id

	d2 := randomSchedule()
	acc, _ = testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000002")
	d2.AccountId = acc.Id

	// Test Create
	CreatedD1 := createTestSchedule(t, d1)
	CreatedD2 := createTestSchedule(t, d2)

	// Get Data
	getData1, err1 := testQueriesAccount.GetSchedule(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.DueDate.Format("2006-01-02"), getData1.DueDate.Format("2006-01-02"))
	require.Equal(t, d1.DuePrin.String(), getData1.DuePrin.String())
	require.Equal(t, d1.DueInt.String(), getData1.DueInt.String())
	require.Equal(t, d1.EndPrin.String(), getData1.EndPrin.String())
	require.Equal(t, d1.EndInt.String(), getData1.EndInt.String())
	require.Equal(t, d1.CarryingValue.String(), getData1.CarryingValue.String())
	require.Equal(t, d1.Realizable.String(), getData1.Realizable.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesAccount.GetSchedule(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.AccountId, getData2.AccountId)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.DueDate.Format("2006-01-02"), getData2.DueDate.Format("2006-01-02"))
	require.Equal(t, d2.DuePrin.String(), getData2.DuePrin.String())
	require.Equal(t, d2.DueInt.String(), getData2.DueInt.String())
	require.Equal(t, d2.EndPrin.String(), getData2.EndPrin.String())
	require.Equal(t, d2.EndInt.String(), getData2.EndInt.String())
	require.Equal(t, d2.CarryingValue.String(), getData2.CarryingValue.String())
	require.Equal(t, d2.Realizable.String(), getData2.Realizable.String())
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesAccount.GetSchedulebyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestSchedule(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.AccountId, updatedD1.AccountId)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.DueDate.Format("2006-01-02"), updatedD1.DueDate.Format("2006-01-02"))
	require.Equal(t, updateD2.DuePrin.String(), updatedD1.DuePrin.String())
	require.Equal(t, updateD2.DueInt.String(), updatedD1.DueInt.String())
	require.Equal(t, updateD2.EndPrin.String(), updatedD1.EndPrin.String())
	require.Equal(t, updateD2.EndInt.String(), updatedD1.EndInt.String())
	require.Equal(t, updateD2.CarryingValue.String(), updatedD1.CarryingValue.String())
	require.Equal(t, updateD2.Realizable.String(), updatedD1.Realizable.String())
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	schedule, err := testQueriesAccount.GetSchedulebyAccId(context.Background(), updatedD1.AccountId)
	require.NoError(t, err)
	require.NotEmpty(t, schedule)

	schedMap, err3 := testQueriesAccount.GetSchedulebyAcc(context.Background(), []string{"1001-0001-0000001", "1001-0001-0000001"})
	require.NoError(t, err3)
	require.NotEmpty(t, schedMap)

	// Delete Data
	testDeleteSchedule(t, getData1.Id)
	testDeleteSchedule(t, getData2.Id)
}

func randomSchedule() ScheduleRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := ScheduleRequest{
		// AccountId:     util.RandomInt(1, 100),
		Series:        int16(util.RandomInt32(1, 100)),
		DueDate:       util.RandomDate(),
		DuePrin:       util.RandomMoney(),
		DueInt:        util.RandomMoney(),
		EndPrin:       util.RandomMoney(),
		EndInt:        util.RandomMoney(),
		CarryingValue: util.RandomMoney(),
		Realizable:    util.RandomMoney(),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestSchedule(
	t *testing.T,
	d1 ScheduleRequest) model.Schedule {

	getData1, err := testQueriesAccount.CreateSchedule(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.DueDate.Format("2006-01-02"), getData1.DueDate.Format("2006-01-02"))
	require.Equal(t, d1.DuePrin.String(), getData1.DuePrin.String())
	require.Equal(t, d1.DueInt.String(), getData1.DueInt.String())
	require.Equal(t, d1.EndPrin.String(), getData1.EndPrin.String())
	require.Equal(t, d1.EndInt.String(), getData1.EndInt.String())
	require.Equal(t, d1.CarryingValue.String(), getData1.CarryingValue.String())
	require.Equal(t, d1.Realizable.String(), getData1.Realizable.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestSchedule(
	t *testing.T,
	d1 ScheduleRequest) model.Schedule {

	getData1, err := testQueriesAccount.UpdateSchedule(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.AccountId, getData1.AccountId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.DueDate.Format("2006-01-02"), getData1.DueDate.Format("2006-01-02"))
	require.Equal(t, d1.DuePrin.String(), getData1.DuePrin.String())
	require.Equal(t, d1.DueInt.String(), getData1.DueInt.String())
	require.Equal(t, d1.EndPrin.String(), getData1.EndPrin.String())
	require.Equal(t, d1.EndInt.String(), getData1.EndInt.String())
	require.Equal(t, d1.CarryingValue.String(), getData1.CarryingValue.String())
	require.Equal(t, d1.Realizable.String(), getData1.Realizable.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteSchedule(t *testing.T, id int64) {
	err := testQueriesAccount.DeleteSchedule(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesAccount.GetSchedule(context.Background(), id)
	require.Error(t, err)
	// require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
