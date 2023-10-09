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

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestOfficeAccountTran(t *testing.T) {

	// Test Data
	d1 := randomOfficeAccountTran()
	trn, _ := testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("2af90d74-3bee-48c5-8935-443edafb8f5a"))
	d1.TrnHeadId = trn.Id

	d2 := randomOfficeAccountTran()
	trn, _ = testQueriesTransaction.GetTrnHeadbyUuid(context.Background(), uuid.MustParse("26dfab18-f80b-46cf-9c54-be79d4fc5d23"))
	d2.TrnHeadId = trn.Id

	// Test Create
	CreatedD1 := createTestOfficeAccountTran(t, d1)
	CreatedD2 := createTestOfficeAccountTran(t, d2)

	// Get Data
	getData1, err1 := testQueriesTransaction.GetOfficeAccountTran(context.Background(), CreatedD1.Uuid)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.OfficeAccountId, getData1.OfficeAccountId)
	require.Equal(t, d1.TrnAmt.String(), getData1.TrnAmt.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesTransaction.GetOfficeAccountTran(context.Background(), CreatedD2.Uuid)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.TrnHeadId, getData2.TrnHeadId)
	require.Equal(t, d2.Series, getData2.Series)
	require.Equal(t, d2.OfficeAccountId, getData2.OfficeAccountId)
	require.Equal(t, d2.TrnAmt.String(), getData2.TrnAmt.String())
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesTransaction.GetOfficeAccountTranbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Uuid, getData.Uuid)
	fmt.Printf("Get by UUId%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Uuid = getData2.Uuid
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestOfficeAccountTran(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.TrnHeadId, updatedD1.TrnHeadId)
	require.Equal(t, updateD2.Series, updatedD1.Series)
	require.Equal(t, updateD2.OfficeAccountId, updatedD1.OfficeAccountId)
	require.Equal(t, updateD2.TrnAmt.String(), updatedD1.TrnAmt.String())
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListOfficeAccountTran(t, ListOfficeAccountTranParams{
		TrnHeadId: updatedD1.TrnHeadId,
		Limit:     5,
		Offset:    0,
	})

	// Delete Data
	testDeleteOfficeAccountTran(t, CreatedD1.Uuid)
	testDeleteOfficeAccountTran(t, CreatedD2.Uuid)
}

func testListOfficeAccountTran(t *testing.T, arg ListOfficeAccountTranParams) {

	officeAccountTran, err := testQueriesTransaction.ListOfficeAccountTran(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", officeAccountTran)
	require.NotEmpty(t, officeAccountTran)

}

func randomOfficeAccountTran() OfficeAccountTranRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	ofc, _ := testQueriesIdentity.GetOfficebyAltId(context.Background(), "10019")
	typ, _ := testQueriesAccount.GetOfficeAccountTypebyName(context.Background(), "Cash")
	part, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "FundSource", 0, "GSB")

	acc, _ := testQueriesAccount.GetOfficeAccountbyCode(context.Background(),
		ofc.Id, typ.Id, "PHP", part.Id)
	log.Println(ofc.Id, typ.Id, "PHP", part.Id)
	fmt.Printf("%+v\n", acc)
	arg := OfficeAccountTranRequest{
		// TrnHeadId:       trn.Id,
		Series:          util.RandomInt16(1, 100),
		OfficeAccountId: acc.Id,
		TrnAmt:          util.RandomMoney(),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestOfficeAccountTran(
	t *testing.T,
	d1 OfficeAccountTranRequest) model.OfficeAccountTran {

	getData1, err := testQueriesTransaction.CreateOfficeAccountTran(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.OfficeAccountId, getData1.OfficeAccountId)
	require.Equal(t, d1.TrnAmt.String(), getData1.TrnAmt.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestOfficeAccountTran(
	t *testing.T,
	d1 OfficeAccountTranRequest) model.OfficeAccountTran {

	getData1, err := testQueriesTransaction.UpdateOfficeAccountTran(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.TrnHeadId, getData1.TrnHeadId)
	require.Equal(t, d1.Series, getData1.Series)
	require.Equal(t, d1.OfficeAccountId, getData1.OfficeAccountId)
	require.Equal(t, d1.TrnAmt.String(), getData1.TrnAmt.String())
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteOfficeAccountTran(t *testing.T, uuid uuid.UUID) {
	err := testQueriesTransaction.DeleteOfficeAccountTran(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesTransaction.GetOfficeAccountTran(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
