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

func TestFxrate(t *testing.T) {

	// Test Data
	d1 := randomFxrate()
	d2 := randomFxrate()

	// Test Create
	CreatedD1 := createTestFxrate(t, d1)
	CreatedD2 := createTestFxrate(t, d2)

	// Get Data
	getData1, err1 := testQueriesReference.GetFxrate(context.Background(), CreatedD1.BaseCurrency, CreatedD1.Currency, CreatedD1.CutofDate)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.BuyRate.String(), getData1.BuyRate.String())
	require.True(t, d1.CutofDate.Equal(getData1.CutofDate))
	require.Equal(t, d1.SellRate.String(), getData1.SellRate.String())
	require.Equal(t, d1.BaseCurrency, getData1.BaseCurrency)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesReference.GetFxrate(context.Background(), CreatedD2.BaseCurrency, CreatedD2.Currency, CreatedD2.CutofDate)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.BuyRate.String(), getData2.BuyRate.String())
	require.True(t, d2.CutofDate.Equal(getData2.CutofDate))
	require.Equal(t, d2.SellRate.String(), getData2.SellRate.String())
	require.Equal(t, d2.BaseCurrency, getData2.BaseCurrency)
	require.Equal(t, d2.Currency, getData2.Currency)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesReference.GetFxratebyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)

	fmt.Printf("%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData.Uuid
	updateD2.CutofDate = util.DateValue("2021-04-15")

	// log.Println(updateD2)
	updatedD1 := updateTestFxrate(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.BuyRate.String(), updatedD1.BuyRate.String())
	require.True(t, updateD2.CutofDate.Equal(updatedD1.CutofDate))
	require.Equal(t, updateD2.SellRate.String(), updatedD1.SellRate.String())
	require.Equal(t, updateD2.BaseCurrency, updatedD1.BaseCurrency)
	require.Equal(t, updateD2.Currency, updatedD1.Currency)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)
	testListFxrate(t)
	// Delete Data
	testDeleteFxrate(t, getData1.Uuid)
	testDeleteFxrate(t, getData2.Uuid)
}

func testListFxrate(t *testing.T) {

	arg := ListFxrateParams{
		CutDate: util.DateValue("2021-04-14"),
		Limit:   5,
		Offset:  0,
	}

	fxrate, err := testQueriesReference.ListFxrate(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", fxrate)
	require.NotEmpty(t, fxrate)

}

func randomFxrate() FxrateRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	arg := FxrateRequest{
		BuyRate:      util.RandomMoney(),
		CutofDate:    util.DateValue("2021-04-14"),
		SellRate:     util.RandomMoney(),
		BaseCurrency: util.RandomCurrency(),
		Currency:     util.RandomCurrency(),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestFxrate(
	t *testing.T,
	d1 FxrateRequest) model.Fxrate {

	getData1, err := testQueriesReference.CreateFxrate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.BuyRate.String(), getData1.BuyRate.String())
	require.True(t, d1.CutofDate.Equal(getData1.CutofDate))
	require.Equal(t, d1.SellRate.String(), getData1.SellRate.String())
	require.Equal(t, d1.BaseCurrency, getData1.BaseCurrency)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestFxrate(
	t *testing.T,
	d1 FxrateRequest) model.Fxrate {

	getData1, err := testQueriesReference.UpdateFxrate(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.BuyRate.String(), getData1.BuyRate.String())
	require.True(t, d1.CutofDate.Equal(getData1.CutofDate))
	require.Equal(t, d1.SellRate.String(), getData1.SellRate.String())
	require.Equal(t, d1.BaseCurrency, getData1.BaseCurrency)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteFxrate(t *testing.T, uuid uuid.UUID) {
	err := testQueriesReference.DeleteFxrate(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesReference.GetFxratebyUuId(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
