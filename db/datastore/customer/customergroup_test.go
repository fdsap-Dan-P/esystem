package db

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"encoding/json"
	common "simplebank/db/common"
	ref "simplebank/db/datastore/reference"
	"simplebank/model"
	"simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestCustomerGroup(t *testing.T) {

	// Test Data
	d1 := randomCustomerGroup()
	d2 := randomCustomerGroup()

	// Test Create
	CreatedD1 := createTestCustomerGroup(t, d1)
	CreatedD2 := createTestCustomerGroup(t, d2)

	// Get Data
	getData1, err1 := testQueriesCustomer.GetCustomerGroup(context.Background(), CreatedD1.Id)

	require.NoError(t, err1)
	require.NotEmpty(t, getData1)
	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.GroupName, getData1.GroupName)
	require.Equal(t, d1.ShortName, getData1.ShortName)
	require.Equal(t, d1.DateStablished.Time.Format("2006-01-02"), getData1.DateStablished.Time.Format("2006-01-02"))
	require.Equal(t, d1.MeetingDay, getData1.MeetingDay)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.OfficerId, getData1.OfficerId)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.Equal(t, d1.AlternateId, getData1.AlternateId)
	require.Equal(t, d1.AddressDetail, getData1.AddressDetail)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	getData2, err2 := testQueriesCustomer.GetCustomerGroup(context.Background(), CreatedD2.Id)
	require.NoError(t, err2)
	require.NotEmpty(t, getData2)
	require.Equal(t, d2.CentralOfficeId, getData2.CentralOfficeId)
	require.Equal(t, d2.Code, getData2.Code)
	require.Equal(t, d2.TypeId, getData2.TypeId)
	require.Equal(t, d2.GroupName, getData2.GroupName)
	require.Equal(t, d2.ShortName, getData2.ShortName)
	require.Equal(t, d2.DateStablished.Time.Format("2006-01-02"), getData2.DateStablished.Time.Format("2006-01-02"))
	require.Equal(t, d2.MeetingDay, getData2.MeetingDay)
	require.Equal(t, d2.OfficeId, getData2.OfficeId)
	require.Equal(t, d2.OfficerId, getData2.OfficerId)
	require.Equal(t, d2.ParentId, getData2.ParentId)
	require.Equal(t, d2.AlternateId, getData2.AlternateId)
	require.Equal(t, d2.AddressDetail, getData2.AddressDetail)
	require.Equal(t, d2.AddressUrl, getData2.AddressUrl)
	require.Equal(t, d2.GeographyId, getData2.GeographyId)
	require.JSONEq(t, d2.OtherInfo.String, getData2.OtherInfo.String)

	getData, err := testQueriesCustomer.GetCustomerGroupbyUuId(context.Background(), CreatedD1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by UUId%+v\n", getData)

	getData, err = testQueriesCustomer.GetCustomerGroupbyCode(
		context.Background(), getData1.GroupType, getData1.CentralOfficeId, getData1.Code)
	require.NoError(t, err)
	require.NotEmpty(t, getData)
	require.Equal(t, CreatedD1.Id, getData.Id)
	fmt.Printf("Get by Name%+v\n", getData)

	// Update Data
	updateD2 := d2
	updateD2.Id = getData2.Id
	updateD2.GroupName = updateD2.GroupName + "Edited"

	// log.Println(updateD2)
	updatedD1 := updateTestCustomerGroup(t, updateD2)
	require.NotEmpty(t, updatedD1)
	require.Equal(t, updateD2.CentralOfficeId, updatedD1.CentralOfficeId)
	require.Equal(t, updateD2.Code, updatedD1.Code)
	require.Equal(t, updateD2.TypeId, updatedD1.TypeId)
	require.Equal(t, updateD2.GroupName, updatedD1.GroupName)
	require.Equal(t, updateD2.ShortName, updatedD1.ShortName)
	require.Equal(t, updateD2.DateStablished.Time.Format("2006-01-02"), updatedD1.DateStablished.Time.Format("2006-01-02"))
	require.Equal(t, updateD2.MeetingDay, updatedD1.MeetingDay)
	require.Equal(t, updateD2.OfficeId, updatedD1.OfficeId)
	require.Equal(t, updateD2.OfficerId, updatedD1.OfficerId)
	require.Equal(t, updateD2.ParentId, updatedD1.ParentId)
	require.Equal(t, updateD2.AlternateId, updatedD1.AlternateId)
	require.Equal(t, updateD2.AddressDetail, updatedD1.AddressDetail)
	require.Equal(t, updateD2.AddressUrl, updatedD1.AddressUrl)
	require.Equal(t, updateD2.GeographyId, updatedD1.GeographyId)
	require.JSONEq(t, updateD2.OtherInfo.String, updatedD1.OtherInfo.String)

	testListCustomerGroup(t, ListCustomerGroupParams{
		GroupType:       getData1.GroupType,
		CentralOfficeId: getData1.CentralOfficeId,
		Limit:           5,
		Offset:          0,
	})

	// Delete Data
	testDeleteCustomerGroup(t, getData1.Id)
	testDeleteCustomerGroup(t, getData2.Id)
}

func testListCustomerGroup(t *testing.T, arg ListCustomerGroupParams) {
	Customer_Group, err := testQueriesCustomer.ListCustomerGroup(context.Background(), arg)
	require.NoError(t, err)
	// fmt.Printf("%+v\n", Customer_Group)
	require.NotEmpty(t, Customer_Group)
}

func randomCustomerGroup() CustomerGroupRequest {
	otherInfo := &common.TestOtherInfo{Greet: "Hello", Name: "World"}
	info, _ := json.Marshal(otherInfo)
	ofc, _ := testQueriesIdentity.GetOfficebyCode(context.Background(), 0, "9999")
	obj, _ := testQueriesReference.GetReferenceInfobyTitle(context.Background(), "CustomerGroupType", 0, "Center")

	ofcr, _ := testQueriesIdentity.GetEmployeebyEmpNo(context.Background(), 2, "97-0114")

	geo, _ := testQueriesReference.SearchGeography(context.Background(),
		ref.SearchGeographyParams{
			SearchText: "Soledad San Pablo City, Laguna",
			Limit:      1,
			Offset:     0,
		})

	arg := CustomerGroupRequest{
		CentralOfficeId: ofc.Id,
		Code:            util.RandomString(10),
		TypeId:          obj.Id,
		GroupName:       util.RandomString(10),
		ShortName:       util.RandomString(10),
		DateStablished:  sql.NullTime(sql.NullTime{Time: util.RandomDate(), Valid: true}),
		MeetingDay:      util.SetNullInt16(10),
		OfficeId:        ofc.Id,
		OfficerId:       sql.NullInt64(sql.NullInt64{Int64: ofcr.Id, Valid: true}),
		ParentId:        sql.NullInt64(sql.NullInt64{Int64: 0, Valid: false}),
		AlternateId:     sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		AddressDetail:   sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		AddressUrl:      sql.NullString(sql.NullString{String: util.RandomString(10), Valid: true}),
		GeographyId:     sql.NullInt64(sql.NullInt64{Int64: geo[0].Id, Valid: true}),

		OtherInfo: sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}
	return arg
}

func createTestCustomerGroup(
	t *testing.T,
	d1 CustomerGroupRequest) model.CustomerGroup {

	getData1, err := testQueriesCustomer.CreateCustomerGroup(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.GroupName, getData1.GroupName)
	require.Equal(t, d1.ShortName, getData1.ShortName)
	require.Equal(t, d1.DateStablished.Time.Format("2006-01-02"), getData1.DateStablished.Time.Format("2006-01-02"))
	require.Equal(t, d1.MeetingDay, getData1.MeetingDay)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.OfficerId, getData1.OfficerId)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.Equal(t, d1.AlternateId, getData1.AlternateId)
	require.Equal(t, d1.AddressDetail, getData1.AddressDetail)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func updateTestCustomerGroup(
	t *testing.T,
	d1 CustomerGroupRequest) model.CustomerGroup {

	getData1, err := testQueriesCustomer.UpdateCustomerGroup(context.Background(), d1)
	require.NoError(t, err)
	require.NotEmpty(t, getData1)

	require.Equal(t, d1.CentralOfficeId, getData1.CentralOfficeId)
	require.Equal(t, d1.Code, getData1.Code)
	require.Equal(t, d1.TypeId, getData1.TypeId)
	require.Equal(t, d1.GroupName, getData1.GroupName)
	require.Equal(t, d1.ShortName, getData1.ShortName)
	require.Equal(t, d1.DateStablished.Time.Format("2006-01-02"), getData1.DateStablished.Time.Format("2006-01-02"))
	require.Equal(t, d1.MeetingDay, getData1.MeetingDay)
	require.Equal(t, d1.OfficeId, getData1.OfficeId)
	require.Equal(t, d1.OfficerId, getData1.OfficerId)
	require.Equal(t, d1.ParentId, getData1.ParentId)
	require.Equal(t, d1.AlternateId, getData1.AlternateId)
	require.Equal(t, d1.AddressDetail, getData1.AddressDetail)
	require.Equal(t, d1.AddressUrl, getData1.AddressUrl)
	require.Equal(t, d1.GeographyId, getData1.GeographyId)
	require.JSONEq(t, d1.OtherInfo.String, getData1.OtherInfo.String)

	return getData1
}

func testDeleteCustomerGroup(t *testing.T, id int64) {
	err := testQueriesCustomer.DeleteCustomerGroup(context.Background(), id)
	require.NoError(t, err)

	ref1, err := testQueriesCustomer.GetCustomerGroup(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ref1)
}
