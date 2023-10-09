package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createCustomerGroup = `-- name: CreateCustomerGroup: one
INSERT INTO Customer_Group (
	Central_Office_Id, Code, Type_Id, Group_Name, Short_Name, 
	Date_Stablished, Meeting_Day, Office_Id, Officer_Id, Parent_Id, Alternate_Id, 
	Address_Detail, Address_URL, Geography_Id, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, 
$13, $14, $15
) RETURNING Id, UUId, Central_Office_Id, Code, Type_Id, Group_Name, Short_Name, 
	Date_Stablished, Meeting_Day, Office_Id, Officer_Id, Parent_Id, Alternate_Id, 
	Address_Detail, Address_URL, Geography_Id, Other_Info
`

type CustomerGroupRequest struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	CentralOfficeId int64          `json:"centralOfficeId"`
	Code            string         `json:"code"`
	TypeId          int64          `json:"typeId"`
	GroupName       string         `json:"group_name"`
	ShortName       string         `json:"short_name"`
	DateStablished  sql.NullTime   `json:"date_stablished"`
	MeetingDay      sql.NullInt16  `json:"meeting_day"`
	OfficeId        int64          `json:"officeId"`
	OfficerId       sql.NullInt64  `json:"officerId"`
	ParentId        sql.NullInt64  `json:"parentId"`
	AlternateId     sql.NullString `json:"alternateId"`
	AddressDetail   sql.NullString `json:"address_detail"`
	AddressUrl      sql.NullString `json:"address_url"`
	GeographyId     sql.NullInt64  `json:"geographyId"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

func (q *QueriesCustomer) CreateCustomerGroup(ctx context.Context, arg CustomerGroupRequest) (model.CustomerGroup, error) {
	row := q.db.QueryRowContext(ctx, createCustomerGroup,
		arg.CentralOfficeId,
		arg.Code,
		arg.TypeId,
		arg.GroupName,
		arg.ShortName,
		arg.DateStablished,
		arg.MeetingDay,
		arg.OfficeId,
		arg.OfficerId,
		arg.ParentId,
		arg.AlternateId,
		arg.AddressDetail,
		arg.AddressUrl,
		arg.GeographyId,
		arg.OtherInfo,
	)
	var i model.CustomerGroup
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.Code,
		&i.TypeId,
		&i.GroupName,
		&i.ShortName,
		&i.DateStablished,
		&i.MeetingDay,
		&i.OfficeId,
		&i.OfficerId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteCustomerGroup = `-- name: DeleteCustomerGroup :exec
DELETE FROM Customer_Group
WHERE id = $1
`

func (q *QueriesCustomer) DeleteCustomerGroup(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustomerGroup, id)
	return err
}

type CustomerGroupInfo struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	CentralOfficeId int64          `json:"centralOfficeId"`
	Code            string         `json:"code"`
	TypeId          int64          `json:"typeId"`
	GroupType       string         `json:"group_type"`
	GroupName       string         `json:"group_name"`
	ShortName       string         `json:"short_name"`
	DateStablished  sql.NullTime   `json:"date_stablished"`
	MeetingDay      sql.NullInt16  `json:"meeting_day"`
	OfficeId        int64          `json:"officeId"`
	OfficerId       sql.NullInt64  `json:"officerId"`
	ParentId        sql.NullInt64  `json:"parentId"`
	AlternateId     sql.NullString `json:"alternateId"`
	AddressDetail   sql.NullString `json:"address_detail"`
	AddressUrl      sql.NullString `json:"address_url"`
	GeographyId     sql.NullInt64  `json:"geographyId"`
	OtherInfo       sql.NullString `json:"otherInfo"`
	ModCtr          int64          `json:"modCtr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

const getCustomerGroup = `-- name: GetCustomerGroup :one
SELECT 
	d.Id, mr.UUId, Central_Office_Id, d.Code, d.Type_Id, r.Title Group_Type, Group_Name, 
	d.Short_Name, Date_Stablished, Meeting_Day, Office_Id, Officer_Id, d.Parent_Id, 
	Alternate_Id, Address_Detail, Address_URL, Geography_Id, d.Other_Info,
	mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Group d 
INNER JOIN Main_Record mr on mr.UUId = d.UUId
INNER JOIN Reference r on r.Id = d.Type_Id
WHERE d.id = $1 LIMIT 1
`

func (q *QueriesCustomer) GetCustomerGroup(ctx context.Context, id int64) (CustomerGroupInfo, error) {
	row := q.db.QueryRowContext(ctx, getCustomerGroup, id)
	var i CustomerGroupInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.Code,
		&i.TypeId,
		&i.GroupType,
		&i.GroupName,
		&i.ShortName,
		&i.DateStablished,
		&i.MeetingDay,
		&i.OfficeId,
		&i.OfficerId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getCustomerGroupbyUuId = `-- name: GetCustomerGroupbyUuId :one
SELECT 
	d.Id, mr.UUId, Central_Office_Id, d.Code, d.Type_Id, r.Title Group_Type, Group_Name, 
	d.Short_Name, Date_Stablished, Meeting_Day, Office_Id, Officer_Id, d.Parent_Id, 
	Alternate_Id, Address_Detail, Address_URL, Geography_Id, d.Other_Info,
	mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Group d 
INNER JOIN Main_Record mr on mr.UUId = d.UUId
INNER JOIN Reference r on r.Id = d.Type_Id
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesCustomer) GetCustomerGroupbyUuId(ctx context.Context, uuid uuid.UUID) (CustomerGroupInfo, error) {
	row := q.db.QueryRowContext(ctx, getCustomerGroupbyUuId, uuid)
	var i CustomerGroupInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.Code,
		&i.TypeId,
		&i.GroupType,
		&i.GroupName,
		&i.ShortName,
		&i.DateStablished,
		&i.MeetingDay,
		&i.OfficeId,
		&i.OfficerId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getCustomerGroupbyAltId = `-- name: GetCustomerGroupbyAltId :one
SELECT 
	d.Id, mr.UUId, Central_Office_Id, d.Code, d.Type_Id, r.Title Group_Type, Group_Name, 
	d.Short_Name, Date_Stablished, Meeting_Day, Office_Id, Officer_Id, d.Parent_Id, 
	Alternate_Id, Address_Detail, Address_URL, Geography_Id, d.Other_Info,
	mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Group d 
INNER JOIN Main_Record mr on mr.UUId = d.UUId
INNER JOIN Reference r on r.Id = d.Type_Id
WHERE d.Alternate_Id = $1 
LIMIT 1
`

func (q *QueriesCustomer) GetCustomerGroupbyAltId(ctx context.Context,
	altId string) (CustomerGroupInfo, error) {
	row := q.db.QueryRowContext(ctx, getCustomerGroupbyAltId, altId)
	var i CustomerGroupInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.Code,
		&i.TypeId,
		&i.GroupType,
		&i.GroupName,
		&i.ShortName,
		&i.DateStablished,
		&i.MeetingDay,
		&i.OfficeId,
		&i.OfficerId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getCustomerGroupbyCode = `-- name: GetCustomerGroupbyCode :one
SELECT 
	d.Id, mr.UUId, Central_Office_Id, d.Code, d.Type_Id, r.Title Group_Type, Group_Name, 
	d.Short_Name, Date_Stablished, Meeting_Day, Office_Id, Officer_Id, d.Parent_Id, 
	Alternate_Id, Address_Detail, Address_URL, Geography_Id, d.Other_Info,
	mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Group d 
INNER JOIN Main_Record mr on mr.UUId = d.UUId
INNER JOIN Reference r on r.Id = d.Type_Id
WHERE lower(r.Title) = lower($1) and Central_Office_Id = $2 and lower(d.Code) = $3 LIMIT 1
`

func (q *QueriesCustomer) GetCustomerGroupbyCode(ctx context.Context,
	groupType string, centralOfficeId int64, code string) (CustomerGroupInfo, error) {
	row := q.db.QueryRowContext(ctx, getCustomerGroupbyCode, groupType, centralOfficeId, code)
	var i CustomerGroupInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.Code,
		&i.TypeId,
		&i.GroupType,
		&i.GroupName,
		&i.ShortName,
		&i.DateStablished,
		&i.MeetingDay,
		&i.OfficeId,
		&i.OfficerId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listCustomerGroup = `-- name: ListCustomerGroup:many
SELECT 
	d.Id, mr.UUId, Central_Office_Id, d.Code, d.Type_Id, r.Title Group_Type, Group_Name, 
	d.Short_Name, Date_Stablished, Meeting_Day, Office_Id, Officer_Id, d.Parent_Id, 
	Alternate_Id, Address_Detail, Address_URL, Geography_Id, d.Other_Info,
	mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Group d 
INNER JOIN Main_Record mr on mr.UUId = d.UUId
INNER JOIN Reference r on r.Id = d.Type_Id
WHERE lower(r.Title) = lower($1) and Central_Office_Id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListCustomerGroupParams struct {
	GroupType       string `json:"group_type"`
	CentralOfficeId int64  `json:"centralOfficeId"`
	Limit           int32  `json:"limit"`
	Offset          int32  `json:"offset"`
}

func (q *QueriesCustomer) ListCustomerGroup(ctx context.Context, arg ListCustomerGroupParams) ([]CustomerGroupInfo, error) {
	rows, err := q.db.QueryContext(ctx, listCustomerGroup, arg.GroupType, arg.CentralOfficeId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CustomerGroupInfo{}
	for rows.Next() {
		var i CustomerGroupInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.CentralOfficeId,
			&i.Code,
			&i.TypeId,
			&i.GroupType,
			&i.GroupName,
			&i.ShortName,
			&i.DateStablished,
			&i.MeetingDay,
			&i.OfficeId,
			&i.OfficerId,
			&i.ParentId,
			&i.AlternateId,
			&i.AddressDetail,
			&i.AddressUrl,
			&i.GeographyId,
			&i.OtherInfo,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCustomerGroup = `-- name: UpdateCustomerGroup :one
UPDATE Customer_Group SET 
	Central_Office_Id = $2,
	Code = $3,
	Type_Id = $4,
	Group_Name = $5,
	Short_Name = $6,
	Date_Stablished = $7,
	Meeting_Day = $8,
	Office_Id = $9,
	Officer_Id = $10,
	Parent_Id = $11,
	Alternate_Id = $12,
	Address_Detail = $13,
	Address_URL = $14,
	Geography_Id = $15,
	Other_Info = $16
WHERE id = $1
RETURNING Id, UUId, Central_Office_Id, Code, Type_Id, Group_Name, Short_Name, 
Date_Stablished, Meeting_Day, Office_Id, Officer_Id, Parent_Id, Alternate_Id, 
Address_Detail, Address_URL, Geography_Id, Other_Info
`

func (q *QueriesCustomer) UpdateCustomerGroup(ctx context.Context, arg CustomerGroupRequest) (model.CustomerGroup, error) {
	row := q.db.QueryRowContext(ctx, updateCustomerGroup,
		arg.Id,
		arg.CentralOfficeId,
		arg.Code,
		arg.TypeId,
		arg.GroupName,
		arg.ShortName,
		arg.DateStablished,
		arg.MeetingDay,
		arg.OfficeId,
		arg.OfficerId,
		arg.ParentId,
		arg.AlternateId,
		arg.AddressDetail,
		arg.AddressUrl,
		arg.GeographyId,
		arg.OtherInfo,
	)
	var i model.CustomerGroup
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.Code,
		&i.TypeId,
		&i.GroupName,
		&i.ShortName,
		&i.DateStablished,
		&i.MeetingDay,
		&i.OfficeId,
		&i.OfficerId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.OtherInfo,
	)
	return i, err
}
