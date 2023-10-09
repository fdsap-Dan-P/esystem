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

func TestCustomer(t *testing.T) {

	// Test Data
	d1 := randomCustomer()
	d2 := randomCustomer()

	// Test Create
	CreatedD1 := createTestCustomer(t, d1)
	CreatedD2 := createTestCustomer(t, d2)

	// Get Data
	getData1, err1 := testQueriesCustomer.GetCustomer(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.Cid, getData1.Cid)
	require.Equal(t, d1.CustomerAltId, getData1.CustomerAltId)
	require.Equal(t, d1.DebitLimit.String(), getData1.DebitLimit.String())
	require.Equal(t, d1.CreditLimit.String(), getData1.CreditLimit.String())
	require.Equal(t, d1.DateEntry.Time.Format("2006-01-02"), getData1.DateEntry.Time.Format("2006-01-02"))
	require.Equal(t, d1.LastActivityDate.Time.Format("2006-01-02"), getData1.LastActivityDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.Dosri, getData1.Dosri)
	require.Equal(t, d1.ClassificationId, getData1.ClassificationId)
	require.Equal(t, d1.CustomerGroupId, getData1.CustomerGroupId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.RestrictionId, getData1.RestrictionId)
	require.Equal(t, d1.RiskClassId, getData1.RiskClassId)
	require.Equal(t, d1.StatusCode, getData1.StatusCode)
	require.Equal(t, d1.SubClassificationId, getData1.SubClassificationId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesCustomer.GetCustomer(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.Iiid, getData2.Iiid)
	require.Equal(t, d2.CentralOfficeId, getData2.CentralOfficeId)
	require.Equal(t, d2.Cid, getData2.Cid)
	require.Equal(t, d2.CustomerAltId, getData2.CustomerAltId)
	require.Equal(t, d2.DebitLimit.String(), getData2.DebitLimit.String())
	require.Equal(t, d2.CreditLimit.String(), getData2.CreditLimit.String())
	require.Equal(t, d2.DateEntry.Time.Format("2006-01-02"), getData2.DateEntry.Time.Format("2006-01-02"))
	require.Equal(t, d2.LastActivityDate.Time.Format("2006-01-02"), getData2.LastActivityDate.Time.Format("2006-01-02"))
	require.Equal(t, d2.Dosri, getData2.Dosri)
	require.Equal(t, d2.ClassificationId, getData2.ClassificationId)
	require.Equal(t, d2.CustomerGroupId, getData2.CustomerGroupId)
	require.Equal(t, d2.OfficeId, getData2.OfficeId)
	require.Equal(t, d2.RestrictionId, getData2.RestrictionId)
	require.Equal(t, d2.RiskClassId, getData2.RiskClassId)
	require.Equal(t, d2.StatusCode, getData2.StatusCode)
	require.Equal(t, d2.SubClassificationId, getData2.SubClassificationId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesCustomer.GetCustomerbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	getData, err = testQueriesCustomer.GetCustomerbyAltId(context.Background(), CreatedD1.CustomerAltId.String)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	// updateD2.Location = updateD2.Location + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestCustomer(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.Iiid, updatedD1.Iiid)
	require.Equal(t, updateD2.CentralOfficeId, updatedD1.CentralOfficeId)
	require.Equal(t, updateD2.Cid, updatedD1.Cid)
	require.Equal(t, updateD2.CustomerAltId, updatedD1.CustomerAltId)
	require.Equal(t, updateD2.DebitLimit.String(), updatedD1.DebitLimit.String())
	require.Equal(t, updateD2.CreditLimit.String(), updatedD1.CreditLimit.String())
	require.Equal(t, updateD2.DateEntry.Time.Format("2006-01-02"), updatedD1.DateEntry.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.LastActivityDate.Time.Format("2006-01-02"), updatedD1.LastActivityDate.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.Dosri, updatedD1.Dosri)
	require.Equal(t, updateD2.ClassificationId, updatedD1.ClassificationId)
	require.Equal(t, updateD2.CustomerGroupId, updatedD1.CustomerGroupId)
	require.Equal(t, updateD2.OfficeId, updatedD1.OfficeId)
	require.Equal(t, updateD2.RestrictionId, updatedD1.RestrictionId)
	require.Equal(t, updateD2.RiskClassId, updatedD1.RiskClassId)
	require.Equal(t, updateD2.StatusCode, updatedD1.StatusCode)
	require.Equal(t, updateD2.SubClassificationId, updatedD1.SubClassificationId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListCustomer(t, ListCustomerParams{
		Iiid:   updatedD1.Iiid,
		Limit:  5,
		Offset: 0,
	})

	// Delete Data
	testDeleteCustomer(t, getData1.Id)
	testDeleteCustomer(t, getData2.Id)
}

func testListCustomer(t *testing.T, arg ListCustomerParams) {

	customer, err := testQueriesCustomer.ListCustomer(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", customer)
	require.NotEmpty(t, customer)

}

func randomCustomer() CustomerRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)

	grp, _ := testQueriesCustomer.GetCustomerGroupbyAltId(context.Background(), "10000")
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")
	cls, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "CustomerClass", 0, "Member")
	subcls, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "CustomerSubClass", cls.Id, "Regular Member")
	rest, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "CustomerRestriction", 0, "Not Restricted")
	risk, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "CustomerRiskClass", 0, "Low Risk")
	stat, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "CustomerStatus", 0, "Active")

	// ofcr, _ := testQueriesIdentity.GetEmployeebyEmpNo(context.Background(), 2, "97-0114")
	ii, _ := testQueriesIdentity.GetIdentityInfobyAltId(context.Background(), "100")

	arg := CustomerRequest{
		Iiid:                ii.Id,
		CentralOfficeId:     ofc.Id,
		Cid:                 util.RandomInt(1, 100000),
		CustomerAltId:       sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		DebitLimit:          util.RandomMoney(),
		CreditLimit:         util.RandomMoney(),
		DateEntry:           sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		LastActivityDate:    sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		Dosri:               true,
		ClassificationId:    sql.NullInt64(sql.NullInt64{Int64: cls.Id, Valid: true}),
		CustomerGroupId:     sql.NullInt64(sql.NullInt64{Int64: grp.Id, Valid: true}),
		OfficeId:            ofc.Id,
		RestrictionId:       sql.NullInt64(sql.NullInt64{Int64: rest.Id, Valid: true}),
		RiskClassId:         sql.NullInt64(sql.NullInt64{Int64: risk.Id, Valid: true}),
		StatusCode:          stat.Code,
		SubClassificationId: sql.NullInt64(sql.NullInt64{Int64: subcls.Id, Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestCustomer(
	t *testing.T,
	d1 CustomerRequest) model.Customer {

	getData1, err := testQueriesCustomer.CreateCustomer(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.Cid, getData1.Cid)
	require.Equal(t, d1.CustomerAltId, getData1.CustomerAltId)
	require.Equal(t, d1.DebitLimit.String(), getData1.DebitLimit.String())
	require.Equal(t, d1.CreditLimit.String(), getData1.CreditLimit.String())
	require.Equal(t, d1.DateEntry.Time.Format("2006-01-02"), getData1.DateEntry.Time.Format("2006-01-02"))
	require.Equal(t, d1.LastActivityDate.Time.Format("2006-01-02"), getData1.LastActivityDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.Dosri, getData1.Dosri)
	require.Equal(t, d1.ClassificationId, getData1.ClassificationId)
	require.Equal(t, d1.CustomerGroupId, getData1.CustomerGroupId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.RestrictionId, getData1.RestrictionId)
	require.Equal(t, d1.RiskClassId, getData1.RiskClassId)
	require.Equal(t, d1.StatusCode, getData1.StatusCode)
	require.Equal(t, d1.SubClassificationId, getData1.SubClassificationId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestCustomer(
	t *testing.T,
	d1 CustomerRequest) model.Customer {

	getData1, err := testQueriesCustomer.UpdateCustomer(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.Iiid, getData1.Iiid)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.Cid, getData1.Cid)
	require.Equal(t, d1.CustomerAltId, getData1.CustomerAltId)
	require.Equal(t, d1.DebitLimit.String(), getData1.DebitLimit.String())
	require.Equal(t, d1.CreditLimit.String(), getData1.CreditLimit.String())
	require.Equal(t, d1.DateEntry.Time.Format("2006-01-02"), getData1.DateEntry.Time.Format("2006-01-02"))
	require.Equal(t, d1.LastActivityDate.Time.Format("2006-01-02"), getData1.LastActivityDate.Time.Format("2006-01-02"))
	require.Equal(t, d1.Dosri, getData1.Dosri)
	require.Equal(t, d1.ClassificationId, getData1.ClassificationId)
	require.Equal(t, d1.CustomerGroupId, getData1.CustomerGroupId)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.RestrictionId, getData1.RestrictionId)
	require.Equal(t, d1.RiskClassId, getData1.RiskClassId)
	require.Equal(t, d1.StatusCode, getData1.StatusCode)
	require.Equal(t, d1.SubClassificationId, getData1.SubClassificationId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteCustomer(t *testing.T, id int64) {
	err := testQueriesCustomer.DeleteCustomer(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesCustomer.GetCustomer(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
