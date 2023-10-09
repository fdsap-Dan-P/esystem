package db

import (
	"context"
	"database/sql"
	"log"

	"testing"

	"encoding/json"
	common "simplebank/db/common"
	"simplebank/model"
	"simplebank/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAccountSubject(t *testing.T) {

	// Test Data
	d1 := randomAccountSubject(t, "English", util.ToUUID("87e3362e-842b-46a4-900a-a9c243f1de66"))
	d2 := randomAccountSubject(t, "Mathematics", util.ToUUID("33d4b023-3e3b-4b7f-bbf3-65a4828a9077"))

	d2.Uuid = util.ToUUID("bb3f0d39-600f-4db1-baaa-fcdc05cc090a")

	// subj1 := testSubjectInfo(t, "English", util.ToUUID("3bdd00c6-dffb-4b7b-9b02-bf358205a690"))
	// subj2 := testSubjectInfo(t, "Filipino", util.ToUUID("e92ad175-eb89-449b-b3a1-7592c099df39"))

	// d1.SubjectId = subj1.Id
	// d2.SubjectId = subj2.Id

	// Test Create
	CreatedD1 := createTestAccountSubject(t, d1)
	CreatedD2 := createTestAccountSubject(t, d2)

	infoUuid, err1 := testQueriesSchool.GetAccountSubjectbyUuid(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err1)
	compareAccountSubjectInfo(t, CreatedD1, infoUuid)

	// Update Data
	updateD2 := d2
	updateD2.Id = CreatedD2.Id
	updateD2.Remarks = updateD2.Remarks + "-Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestAccountSubject(t, updateD2)

	testListAccountSubject(t, ListAccountSubjectParams{
		AccountId: updatedD1.AccountId,
		Limit:     5,
		Offset:    0,
	})

	// testAccountSubjectInfo(t)

	// Delete Data
	testDeleteAccountSubject(t, CreatedD1.Uuid)
	testDeleteAccountSubject(t, CreatedD2.Uuid)
}

func testListAccountSubject(t *testing.T, arg ListAccountSubjectParams) {

	AccountSubject, err := testQueriesSchool.ListAccountSubject(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", AccountSubject)
	require.NotEmpty(t, AccountSubject)
}

func compareAccountSubjectRequest(t *testing.T, mod model.AccountSubject, req AccountSubjectRequest) {
	require.NotEmpty(t, mod)
	require.Equal(t, req.Uuid, mod.Uuid)
	require.Equal(t, req.AccountId, mod.AccountId)
	require.Equal(t, req.SectionSubjectId, mod.SectionSubjectId)
	require.Equal(t, req.SubjectId, mod.SubjectId)
	require.Equal(t, req.Ratings1stQtr.String(), mod.Ratings1stQtr.String())
	require.Equal(t, req.Ratings2ndQtr.String(), mod.Ratings2ndQtr.String())
	require.Equal(t, req.Ratings3rdQtr.String(), mod.Ratings3rdQtr.String())
	require.Equal(t, req.Ratings4thQtr.String(), mod.Ratings4thQtr.String())
	require.Equal(t, req.RatingsFinal.String(), mod.RatingsFinal.String())
	require.Equal(t, req.AttendanceCtr, mod.AttendanceCtr)
	require.Equal(t, req.AbsentCtr, mod.AbsentCtr)
	require.Equal(t, req.LateCtr, mod.LateCtr)
	require.Equal(t, req.StatusId, mod.StatusId)
	require.Equal(t, req.Remarks, mod.Remarks)
	require.JSONEq(t, req.OtherInfo.String, mod.OtherInfo.String)
}

func compareAccountSubjectInfo(t *testing.T, mod model.AccountSubject, info AccountSubjectInfo) {
	require.NotEmpty(t, mod)
	require.Equal(t, info.Uuid, mod.Uuid)
	require.Equal(t, info.AccountId, mod.AccountId)
	require.Equal(t, info.SectionSubjectId, mod.SectionSubjectId)
	require.Equal(t, info.SubjectId, mod.SubjectId)
	require.Equal(t, info.Ratings1stQtr.String(), mod.Ratings1stQtr.String())
	require.Equal(t, info.Ratings2ndQtr.String(), mod.Ratings2ndQtr.String())
	require.Equal(t, info.Ratings3rdQtr.String(), mod.Ratings3rdQtr.String())
	require.Equal(t, info.Ratings4thQtr.String(), mod.Ratings4thQtr.String())
	require.Equal(t, info.RatingsFinal.String(), mod.RatingsFinal.String())
	require.Equal(t, info.AttendanceCtr, mod.AttendanceCtr)
	require.Equal(t, info.AbsentCtr, mod.AbsentCtr)
	require.Equal(t, info.LateCtr, mod.LateCtr)
	require.Equal(t, info.StatusId, mod.StatusId)
	require.Equal(t, info.Remarks, mod.Remarks)
	require.JSONEq(t, info.OtherInfo.String, mod.OtherInfo.String)
}
func createTestAccountSubject(
	t *testing.T,
	req AccountSubjectRequest) model.AccountSubject {

	mod, err := testQueriesSchool.CreateAccountSubject(context.Background(), req)
	require.NoError(t, err)
	compareAccountSubjectRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetAccountSubject(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareAccountSubjectInfo(t, mod, info)

	return mod
}

func updateTestAccountSubject(
	t *testing.T,
	req AccountSubjectRequest) model.AccountSubject {

	mod, err := testQueriesSchool.UpdateAccountSubject(context.Background(), req)
	require.NoError(t, err)
	compareAccountSubjectRequest(t, mod, req)

	info, err1 := testQueriesSchool.GetAccountSubject(context.Background(), mod.Id)
	require.NoError(t, err1)
	compareAccountSubjectInfo(t, mod, info)

	return mod
}

func testDeleteAccountSubject(t *testing.T, uuid uuid.UUID) {
	err := testQueriesSchool.DeleteAccountSubject(context.Background(), uuid)
	require.NoError(t, err)

	ref1, err := testQueriesSchool.GetAccountSubjectbyUuid(context.Background(), uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}

func randomAccountSubject(t *testing.T, subject string, uuid uuid.UUID) AccountSubjectRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	acc, _ := testQueriesAccount.GetAccountbyAltAcc(context.Background(), "1001-0001-0000001")

	accSubj, err := testQueriesSchool.CreateSectionSubject(context.Background(), randomSectionSubject(t, subject, uuid))

	log.Printf("randomAccountSubject %v", accSubj)
	log.Printf("err %v", err)

	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectStatus", 0, "Current")
	arg := AccountSubjectRequest{
		Uuid:             util.ToUUID("3a7f6c08-8699-4e19-8017-c0770759343f"),
		AccountId:        acc.Id,
		SectionSubjectId: accSubj.Id,
		SubjectId:        accSubj.SubjectId,
		Ratings1stQtr:    util.RandomMoney(),
		Ratings2ndQtr:    util.RandomMoney(),
		Ratings3rdQtr:    util.RandomMoney(),
		Ratings4thQtr:    util.RandomMoney(),
		RatingsFinal:     util.RandomMoney(),
		AttendanceCtr:    util.RandomInt(1, 100),
		AbsentCtr:        util.RandomInt(1, 100),
		LateCtr:          util.RandomInt(1, 100),
		StatusId:         stat.Id,
		Remarks:          util.RandomString(10),
		OtherInfo:        sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}

	return arg
}

// func testAccountSubjectInfo(t *testing.T) AccountSubjectInfo {
// 	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
// 	info, _ := json.Marshal(otherInfo)
// 	// typ, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectType", 0, "Regular")
// 	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SubjectStatus", 0, "Current")
// 	// adv, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "SchoolSemister", 0, "Full Year")

// 	// sec := testAccountSubjectData(t)
// 	// secInfo, _ := testQueriesSchool.GetAccountSubjectbyUuid(context.Background(), sec.Uuid)
// 	// // Test Create
// 	// if secInfo.Id == 0 {
// 	// 	// syl := createTestAccountSubject(t, sec)
// 	// 	secInfo, _ = testQueriesSchool.GetAccountSubjectbyUuid(context.Background(), syl.Uuid)
// 	// }

// 	// subj := testSubjectInfo(t, "English", util.ToUUID("3bdd00c6-dffb-4b7b-9b02-bf358205a690"))

// 	arg := AccountSubjectRequest{
// 		Uuid: util.ToUUID("ed6a63f9-44bd-4aef-87ed-8e5538f392f8"),
// 		// AccountId:,
// 		// SectionSubjectId:d,
// 		// SubjectId: subj.Id,
// 		// Ratings1stQtr: d,
// 		// Ratings2ndQtr: d,
// 		// Ratings3rdQtr: d,
// 		// Ratings4thQtr: d,
// 		// RatingsFinal: d,
// 		// AttendanceCtr: d,
// 		// AbsentCtr: d,
// 		// LateCtr: d,
// 		StatusId:  stat.Id,
// 		Remarks:   util.RandomString(5),
// 		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
// 	}

// 	secSubjInfo, _ := testQueriesSchool.GetAccountSubjectbyUuid(context.Background(), arg.Uuid)
// 	// require.NoError(t, err)

// 	if secSubjInfo.Id == 0 {
// 		syl := createTestAccountSubject(t, arg)
// 		secSubjInfo, _ = testQueriesSchool.GetAccountSubjectbyUuid(context.Background(), syl.Uuid)
// 	}

// 	return secSubjInfo
// }
