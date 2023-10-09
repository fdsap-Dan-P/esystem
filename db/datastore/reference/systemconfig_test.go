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

func TestSystemConfig(t *testing.T) {

	// Test Data
	d1 := randomSystemConfig()
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")
	d1.OfficeId = ofc.Id

	d2 := randomSystemConfig()
	ofc, _ = testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "0000")
	d2.OfficeId = ofc.Id

	// Test Create
	CreatedD1 := createTestSystemConfig(t, d1)
	CreatedD2 := createTestSystemConfig(t, d2)

	// Get Data
	getData1, err1 := testQueriesReference.GetSystemConfig(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.GlDate.Format("2006-01-02"), getData1.GlDate.Format("2006-01-02"))
	require.Equal(t, d1.LastAccruals.Format("2006-01-02"), getData1.LastAccruals.Format("2006-01-02"))
	require.Equal(t, d1.LastMonthEnd.Format("2006-01-02"), getData1.LastMonthEnd.Format("2006-01-02"))
	require.Equal(t, d1.NextMonthEnd.Format("2006-01-02"), getData1.NextMonthEnd.Format("2006-01-02"))
	require.Equal(t, d1.SystemDate.Format("2006-01-02"), getData1.SystemDate.Format("2006-01-02"))
	require.Equal(t, d1.RunState, getData1.RunState)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesReference.GetSystemConfig(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.OfficeId, getData2.OfficeId)
	require.Equal(t, d2.GlDate.Format("2006-01-02"), getData2.GlDate.Format("2006-01-02"))
	require.Equal(t, d2.LastAccruals.Format("2006-01-02"), getData2.LastAccruals.Format("2006-01-02"))
	require.Equal(t, d2.LastMonthEnd.Format("2006-01-02"), getData2.LastMonthEnd.Format("2006-01-02"))
	require.Equal(t, d2.NextMonthEnd.Format("2006-01-02"), getData2.NextMonthEnd.Format("2006-01-02"))
	require.Equal(t, d2.SystemDate.Format("2006-01-02"), getData2.SystemDate.Format("2006-01-02"))
	require.Equal(t, d2.RunState, getData2.RunState)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesReference.GetSystemConfigbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// getData, err = testQueriesReference.GetSystemConfigbyName(context.Background(), CreatedD1.Name)
	// require.NoError(t, err)
	// require.NotEmpty(t, getData)
	// require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	// fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestSystemConfig(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.OfficeId, updatedD1.OfficeId)
	require.Equal(t, updateD2.GlDate.Format("2006-01-02"), updatedD1.GlDate.Format("2006-01-02"))
	require.Equal(t, updateD2.LastAccruals.Format("2006-01-02"), updatedD1.LastAccruals.Format("2006-01-02"))
	require.Equal(t, updateD2.LastMonthEnd.Format("2006-01-02"), updatedD1.LastMonthEnd.Format("2006-01-02"))
	require.Equal(t, updateD2.NextMonthEnd.Format("2006-01-02"), updatedD1.NextMonthEnd.Format("2006-01-02"))
	require.Equal(t, updateD2.SystemDate.Format("2006-01-02"), updatedD1.SystemDate.Format("2006-01-02"))
	require.Equal(t, updateD2.RunState, updatedD1.RunState)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListSystemConfig(t, ListSystemConfigParams{
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteSystemConfig(t, CreatedD1.Uuid)
	testDeleteSystemConfig(t, CreatedD2.Uuid)
}

func testListSystemConfig(t *testing.T, arg ListSystemConfigParams) {

	systemConfig, err := testQueriesReference.ListSystemConfig(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", systemConfig)
	require.NotEmpty(t, systemConfig)

}

func randomSystemConfig() SystemConfigRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")

	arg := &SystemConfigRequest{
		OfficeId:     ofc.Id,
		GlDate:       util.RandomDate(),
		LastAccruals: util.RandomDate(),
		LastMonthEnd: util.RandomDate(),
		NextMonthEnd: util.RandomDate(),
		SystemDate:   util.RandomDate(),
		RunState:     int16(util.RandomInt32(1, 100)),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return *arg
}

func createTestSystemConfig(
	t *testing.T,
	d1 SystemConfigRequest) model.SystemConfig {

	getData1, err := testQueriesReference.CreateSystemConfig(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.GlDate.Format("2006-01-02"), getData1.GlDate.Format("2006-01-02"))
	require.Equal(t, d1.LastAccruals.Format("2006-01-02"), getData1.LastAccruals.Format("2006-01-02"))
	require.Equal(t, d1.LastMonthEnd.Format("2006-01-02"), getData1.LastMonthEnd.Format("2006-01-02"))
	require.Equal(t, d1.NextMonthEnd.Format("2006-01-02"), getData1.NextMonthEnd.Format("2006-01-02"))
	require.Equal(t, d1.SystemDate.Format("2006-01-02"), getData1.SystemDate.Format("2006-01-02"))
	require.Equal(t, d1.RunState, getData1.RunState)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)
	return getData1
}

func updateTestSystemConfig(
	t *testing.T,
	d1 SystemConfigRequest) model.SystemConfig {

	getData1, err := testQueriesReference.UpdateSystemConfig(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.GlDate.Format("2006-01-02"), getData1.GlDate.Format("2006-01-02"))
	require.Equal(t, d1.LastAccruals.Format("2006-01-02"), getData1.LastAccruals.Format("2006-01-02"))
	require.Equal(t, d1.LastMonthEnd.Format("2006-01-02"), getData1.LastMonthEnd.Format("2006-01-02"))
	require.Equal(t, d1.NextMonthEnd.Format("2006-01-02"), getData1.NextMonthEnd.Format("2006-01-02"))
	require.Equal(t, d1.SystemDate.Format("2006-01-02"), getData1.SystemDate.Format("2006-01-02"))
	require.Equal(t, d1.RunState, getData1.RunState)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteSystemConfig(t *testing.T, uuid uuid.UUID) {
	err := testQueriesReference.DeleteSystemConfig(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesReference.GetSystemConfig(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
