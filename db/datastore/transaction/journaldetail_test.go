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

func TestJournalDetail(t *testing.T) {

	// Test Data
	d1 := randomJournalDetail()
	trn, _ := testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("2af90d74-3bee-48c5-8935-443edafb8f5a"))
	d1.TrnHeadId = trn.Id

	d2 := randomJournalDetail()
	trn, _ = testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("26dfab18-f80b-46cf-9c54-be79d4fc5d23"))
	d2.TrnHeadId = trn.Id

	// Test Create
	CreatedD1 := createTestJournalDetail(t, d1)
	CreatedD2 := createTestJournalDetail(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetJournalDetail(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.CoaId, getData1.CoaId)
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.Equal(t, d1.PartitionId, getData1.PartitionId)
	require.Equal(t, d1.TrnAmt.String(), getData1.TrnAmt.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetJournalDetail(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TrnHeadId, getData2.TrnHeadId)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.OfficeId, getData2.OfficeId)
	require.Equal(t, d2.CoaId, getData2.CoaId)
	require.Equal(t, d2.AccountTypeId, getData2.AccountTypeId)
	require.Equal(t, d2.Currency, getData2.Currency)
	require.Equal(t, d2.PartitionId, getData2.PartitionId)
	require.Equal(t, d2.TrnAmt.String(), getData2.TrnAmt.String())
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetJournalDetailbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestJournalDetail(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TrnHeadId, updatedD1.TrnHeadId)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.OfficeId, updatedD1.OfficeId)
	require.Equal(t, updateD2.CoaId, updatedD1.CoaId)
	require.Equal(t, updateD2.AccountTypeId, updatedD1.AccountTypeId)
	require.Equal(t, updateD2.Currency, updatedD1.Currency)
	require.Equal(t, updateD2.PartitionId, updatedD1.PartitionId)
	require.Equal(t, updateD2.TrnAmt.String(), updatedD1.TrnAmt.String())
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListJournalDetail(t, ListJournalDetailParams{
		TrnHeadId: updatedD1.TrnHeadId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteJournalDetail(t, CreatedD1.Uuid)
	testDeleteJournalDetail(t, CreatedD2.Uuid)
}

func testListJournalDetail(t *testing.T, arg ListJournalDetailParams) {

	journalDetail, err := testQueriesTransaction.ListJournalDetail(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", journalDetail)
	require.NotEmpty(t, journalDetail)

}

func randomJournalDetail() JournalDetailRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ofc, _ := testQueriesIdentity.GetOfficebyAltId(context.Background(), "10019")
	coa, _ := testQueriesReference.GetChartofAccountbyTitle(context.Background(), "Microfinance Loan")
	part, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "FundSource", 0, "GSB")

	acc, _ := testQueriesAccount.GetAccountTypebyName(context.Background(), "Sikap 1")

	arg := JournalDetailRequest{
		// TrnHeadId:     util.RandomInt(1, 100),
		Series:        util.RandomInt16(1, 100),
		OfficeId:      ofc.Id,
		CoaId:         sql.NullInt64(sql.NullInt64{Int64: coa.Id, Valid: true}),
		AccountTypeId: sql.NullInt64(sql.NullInt64{Int64: acc.Id, Valid: true}),
		Currency:      sql.NullString(sql.NullString{String: util.RandomCurrency(), Valid: true}),
		PartitionId:   sql.NullInt64(sql.NullInt64{Int64: part.Id, Valid: true}),
		TrnAmt:        util.RandomMoney(),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestJournalDetail(
	t *testing.T,
	d1 JournalDetailRequest) model.JournalDetail {

	getData1, err := testQueriesTransaction.CreateJournalDetail(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.CoaId, getData1.CoaId)
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.Equal(t, d1.PartitionId, getData1.PartitionId)
	require.Equal(t, d1.TrnAmt.String(), getData1.TrnAmt.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestJournalDetail(
	t *testing.T,
	d1 JournalDetailRequest) model.JournalDetail {

	getData1, err := testQueriesTransaction.UpdateJournalDetail(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.CoaId, getData1.CoaId)
	require.Equal(t, d1.AccountTypeId, getData1.AccountTypeId)
	require.Equal(t, d1.Currency, getData1.Currency)
	require.Equal(t, d1.PartitionId, getData1.PartitionId)
	require.Equal(t, d1.TrnAmt.String(), getData1.TrnAmt.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteJournalDetail(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteJournalDetail(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetJournalDetail(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
