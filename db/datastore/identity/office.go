package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createOffice = `-- name: CreateOffice: one
INSERT INTO Office (
  Code, Short_Name, Office_Name, Date_Stablished, Type_ID, 
  Parent_Id, Alternate_Id, Address_Detail, Address_URL, Geography_Id, 
  Cid_Sequence, Other_Info
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
RETURNING 
  Id, UUId, Code, Short_Name, Office_Name, Date_Stablished, Type_ID, 
  Parent_Id, Alternate_Id, Address_Detail, Address_URL, Geography_Id, 
  Cid_Sequence, Other_Info
`

type OfficeRequest struct {
	Id             int64          `json:"id"`
	Uuid           uuid.UUID      `json:"uuid"`
	Code           string         `json:"code"`
	ShortName      string         `json:"shortName"`
	OfficeName     string         `json:"officeName"`
	DateStablished sql.NullTime   `json:"dateStablished"`
	TypeId         int64          `json:"typeId"`
	ParentId       sql.NullInt64  `json:"parentId"`
	AlternateId    sql.NullString `json:"alternateId"`
	AddressDetail  sql.NullString `json:"addressDetail"`
	AddressUrl     sql.NullString `json:"addressUrl"`
	GeographyId    sql.NullInt64  `json:"geographyId"`
	CidSequence    sql.NullInt64  `json:"CidSequence"`
	OtherInfo      sql.NullString `json:"otherInfo"`
}

func (q *QueriesIdentity) CreateOffice(ctx context.Context, arg OfficeRequest) (model.Office, error) {
	row := q.db.QueryRowContext(ctx, createOffice,
		arg.Code,
		arg.ShortName,
		arg.OfficeName,
		arg.DateStablished,
		arg.TypeId,
		arg.ParentId,
		arg.AlternateId,
		arg.AddressDetail,
		arg.AddressUrl,
		arg.GeographyId,
		arg.CidSequence,
		arg.OtherInfo,
	)
	var i model.Office
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.OfficeName,
		&i.DateStablished,
		&i.TypeId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.CidSequence,
		&i.OtherInfo,
	)
	return i, err
}

const deleteOffice = `-- name: DeleteOffice :exec
DELETE FROM Office
WHERE id = $1
`

func (q *QueriesIdentity) DeleteOffice(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteOffice, id)
	return err
}

type OfficeInfo struct {
	Id             int64          `json:"id"`
	Uuid           uuid.UUID      `json:"uuid"`
	Code           string         `json:"code"`
	ShortName      string         `json:"shortName"`
	OfficeName     string         `json:"officeName"`
	DateStablished sql.NullTime   `json:"dateStablished"`
	TypeId         int64          `json:"typeId"`
	ParentId       sql.NullInt64  `json:"parentId"`
	AlternateId    sql.NullString `json:"alternateId"`
	AddressDetail  sql.NullString `json:"addressDetail"`
	AddressUrl     sql.NullString `json:"addressUrl"`
	GeographyId    sql.NullInt64  `json:"geographyId"`
	CidSequence    sql.NullInt64  `json:"CidSequence"`
	OtherInfo      sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getOffice = `-- name: GetOffice :one
SELECT 
Id, mr.UUId, Code, Short_Name, 
Office_Name, Date_Stablished, Type_ID, Parent_Id, Alternate_Id, 
Address_Detail, Address_URL, Geography_Id, Cid_Sequence, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesIdentity) GetOffice(ctx context.Context, id int64) (OfficeInfo, error) {
	row := q.db.QueryRowContext(ctx, getOffice, id)
	var i OfficeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.OfficeName,
		&i.DateStablished,
		&i.TypeId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.CidSequence,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getOfficebyUuId = `-- name: GetOfficebyUuId :one
SELECT 
Id, mr.UUId, Code, Short_Name, 
Office_Name, Date_Stablished, Type_ID, Parent_Id, Alternate_Id, 
Address_Detail, Address_URL, Geography_Id, Cid_Sequence, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesIdentity) GetOfficebyUuId(ctx context.Context, uuid uuid.UUID) (OfficeInfo, error) {
	row := q.db.QueryRowContext(ctx, getOfficebyUuId, uuid)
	var i OfficeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.OfficeName,
		&i.DateStablished,
		&i.TypeId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.CidSequence,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getOfficebyShortName = `-- name: GetOfficebyShortName :one
SELECT 
Id, mr.UUId, Code, Short_Name, 
Office_Name, Date_Stablished, Type_ID, Parent_Id, Alternate_Id, 
Address_Detail, Address_URL, Geography_Id, Cid_Sequence, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE COALESCE(Parent_Id,0) = $1 and lower(Short_Name) = lower($2) LIMIT 1
`

func (q *QueriesIdentity) GetOfficebyShortName(ctx context.Context, parentId int64, shortName string) (OfficeInfo, error) {
	row := q.db.QueryRowContext(ctx, getOfficebyShortName, parentId, shortName)
	var i OfficeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.OfficeName,
		&i.DateStablished,
		&i.TypeId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.CidSequence,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getOfficebyCode = `-- name: GetOfficebyCode :one
SELECT 
Id, mr.UUId, Code, Short_Name, 
Office_Name, Date_Stablished, Type_ID, Parent_Id, Alternate_Id, 
Address_Detail, Address_URL, Geography_Id, Cid_Sequence, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE COALESCE(Parent_Id,0) = $1 and lower(Code) = lower($2) LIMIT 1
`

func (q *QueriesIdentity) GetOfficebyCode(ctx context.Context, parentId int64, code string) (OfficeInfo, error) {
	row := q.db.QueryRowContext(ctx, getOfficebyCode, parentId, code)
	var i OfficeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.OfficeName,
		&i.DateStablished,
		&i.TypeId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.CidSequence,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getOfficebyAltId = `-- name: GetOfficebyAltId :one
SELECT 
Id, mr.UUId, Code, Short_Name, 
Office_Name, Date_Stablished, Type_ID, Parent_Id, Alternate_Id, 
Address_Detail, Address_URL, Geography_Id, Cid_Sequence, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Alternate_Id = $1 LIMIT 1
`

func (q *QueriesIdentity) GetOfficebyAltId(ctx context.Context, altId string) (OfficeInfo, error) {
	row := q.db.QueryRowContext(ctx, getOfficebyAltId, altId)
	var i OfficeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.OfficeName,
		&i.DateStablished,
		&i.TypeId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.CidSequence,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listOffice = `-- name: ListOffice:many
SELECT 
Id, mr.UUId, Code, Short_Name, 
Office_Name, Date_Stablished, Type_ID, Parent_Id, Alternate_Id, 
Address_Detail, Address_URL, Geography_Id, Cid_Sequence, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Type_ID = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListOfficeParams struct {
	TypeId int64 `json:"typeId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListOffice(ctx context.Context, arg ListOfficeParams) ([]OfficeInfo, error) {
	rows, err := q.db.QueryContext(ctx, listOffice, arg.TypeId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OfficeInfo{}
	for rows.Next() {
		var i OfficeInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Code,
			&i.ShortName,
			&i.OfficeName,
			&i.DateStablished,
			&i.TypeId,
			&i.ParentId,
			&i.AlternateId,
			&i.AddressDetail,
			&i.AddressUrl,
			&i.GeographyId,
			&i.CidSequence,
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

const updateOffice = `-- name: UpdateOffice :one
UPDATE Office SET 
Code = $2,
Short_Name = $3,
Office_Name = $4,
Date_Stablished = $5,
Type_ID = $6,
Parent_Id = $7,
Alternate_Id = $8,
Address_Detail = $9,
Address_URL = $10,
Geography_Id = $11,
Cid_Sequence = $12,
Other_Info = $13
WHERE id = $1
RETURNING Id, UUId, Code, Short_Name, Office_Name, Date_Stablished, Type_ID, 
Parent_Id, Alternate_Id, Address_Detail, Address_URL, Geography_Id, 
Cid_Sequence, Other_Info
`

func (q *QueriesIdentity) UpdateOffice(ctx context.Context, arg OfficeRequest) (model.Office, error) {
	row := q.db.QueryRowContext(ctx, updateOffice,
		arg.Id,
		arg.Code,
		arg.ShortName,
		arg.OfficeName,
		arg.DateStablished,
		arg.TypeId,
		arg.ParentId,
		arg.AlternateId,
		arg.AddressDetail,
		arg.AddressUrl,
		arg.GeographyId,
		arg.CidSequence,
		arg.OtherInfo,
	)
	var i model.Office
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.OfficeName,
		&i.DateStablished,
		&i.TypeId,
		&i.ParentId,
		&i.AlternateId,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.CidSequence,
		&i.OtherInfo,
	)
	return i, err
}
