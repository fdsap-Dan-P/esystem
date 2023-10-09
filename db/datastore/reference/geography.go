package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createGeography = `-- name: CreateGeography: one
INSERT INTO Geography (
Code, Short_Name, Location, Type_Id, Parent_Id, 
Zip_Code, Latitude, Longitude, Address_URL, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING Id, UUId, Code, Short_Name, Location, Type_Id, Parent_Id, 
Zip_Code, Latitude, Longitude, Address_URL, Simple_Location, Full_Location, 
Other_Info
`

type GeographyRequest struct {
	Id int64 `json:"id"`
	// Uuid       uuid.UUID       `json:"uuid"`
	Code       int64               `json:"code"`
	ShortName  sql.NullString      `json:"shortName"`
	Location   string              `json:"location"`
	TypeId     int64               `json:"typeId"`
	ParentId   sql.NullInt64       `json:"parentId"`
	ZipCode    sql.NullString      `json:"zipCode"`
	Latitude   decimal.NullDecimal `json:"latitude"`
	Longitude  decimal.NullDecimal `json:"longitude"`
	AddressUrl sql.NullString      `json:"addressUrl"`
	OtherInfo  sql.NullString      `json:"otherInfo"`
}

func (q *QueriesReference) CreateGeography(ctx context.Context, arg GeographyRequest) (model.Geography, error) {
	row := q.db.QueryRowContext(ctx, createGeography,
		arg.Code,
		arg.ShortName,
		arg.Location,
		arg.TypeId,
		arg.ParentId,
		arg.ZipCode,
		arg.Latitude,
		arg.Longitude,
		arg.AddressUrl,
		arg.OtherInfo,
	)
	var i model.Geography
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.Location,
		&i.TypeId,
		&i.ParentId,
		&i.ZipCode,
		&i.Latitude,
		&i.Longitude,
		&i.AddressUrl,
		&i.SimpleLocation,
		&i.FullLocation,
		&i.OtherInfo,
	)
	return i, err
}

const deleteGeography = `-- name: DeleteGeography :exec
DELETE FROM Geography
WHERE id = $1
`

func (q *QueriesReference) DeleteGeography(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteGeography, id)
	return err
}

type GeographyInfo struct {
	Id             int64               `json:"id"`
	Uuid           uuid.UUID           `json:"uuid"`
	Code           int64               `json:"code"`
	ShortName      sql.NullString      `json:"shortName"`
	Location       string              `json:"location"`
	TypeId         int64               `json:"typeId"`
	ParentId       sql.NullInt64       `json:"parentId"`
	ZipCode        sql.NullString      `json:"zipCode"`
	Latitude       decimal.NullDecimal `json:"latitude"`
	Longitude      decimal.NullDecimal `json:"longitude"`
	AddressUrl     sql.NullString      `json:"addressUrl"`
	SimpleLocation sql.NullString      `json:"simple_location"`
	FullLocation   sql.NullString      `json:"full_location"`
	OtherInfo      sql.NullString      `json:"otherInfo"`
	ModCtr         int64               `json:"modCtr"`
	Created        sql.NullTime        `json:"created"`
	Updated        sql.NullTime        `json:"updated"`
}

type GeographyInfoSearch struct {
	Id             int64               `json:"id"`
	Uuid           uuid.UUID           `json:"uuid"`
	Code           int64               `json:"code"`
	ShortName      sql.NullString      `json:"shortName"`
	Location       string              `json:"location"`
	TypeId         int64               `json:"typeId"`
	ParentId       sql.NullInt64       `json:"parentId"`
	ZipCode        sql.NullString      `json:"zipCode"`
	Latitude       decimal.NullDecimal `json:"latitude"`
	Longitude      decimal.NullDecimal `json:"longitude"`
	AddressUrl     sql.NullString      `json:"addressUrl"`
	SimpleLocation sql.NullString      `json:"simple_location"`
	FullLocation   sql.NullString      `json:"full_location"`
	RankFull       float64             `json:"rnkfull"`
	Rank           float64             `json:"rnk"`
	OtherInfo      sql.NullString      `json:"otherInfo"`
	ModCtr         int64               `json:"modCtr"`
	Created        sql.NullTime        `json:"created"`
	Updated        sql.NullTime        `json:"updated"`
}

const getGeography = `-- name: GetGeography :one
SELECT 
Id, mr.UUId, Code, 
Short_Name, Location, Type_Id, Parent_Id, Zip_Code, Latitude, 
Longitude, Address_URL, Simple_Location, Full_Location, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Geography d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesReference) GetGeography(ctx context.Context, id int64) (GeographyInfo, error) {
	row := q.db.QueryRowContext(ctx, getGeography, id)
	var i GeographyInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.Location,
		&i.TypeId,
		&i.ParentId,
		&i.ZipCode,
		&i.Latitude,
		&i.Longitude,
		&i.AddressUrl,
		&i.SimpleLocation,
		&i.FullLocation,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getGeographybyAcc = `-- name: GetGeographybyAcc :one
SELECT 
Id, mr.UUId, Code, 
Short_Name, Location, Type_Id, Parent_Id, Zip_Code, Latitude, 
Longitude, Address_URL, Simple_Location, Full_Location, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Geography d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE acc = $1 LIMIT 1
`

func (q *QueriesReference) GetGeographybyAcc(ctx context.Context, acc string) (GeographyInfo, error) {
	row := q.db.QueryRowContext(ctx, getGeographybyAcc, acc)
	var i GeographyInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.Location,
		&i.TypeId,
		&i.ParentId,
		&i.ZipCode,
		&i.Latitude,
		&i.Longitude,
		&i.AddressUrl,
		&i.SimpleLocation,
		&i.FullLocation,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getGeographybyLocation = `-- name: GetGeographybyLocation :one
SELECT 
Id, mr.UUId, Code, 
Short_Name, Location, Type_Id, Parent_Id, Zip_Code, Latitude, 
Longitude, Address_URL, Simple_Location, Full_Location, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Geography d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE 
  d.Type_Id = $1
  and COALESCE(d.Parent_Id,0) = $2
  and LOWER(d.Location) = LOWER($3) 
LIMIT 1
`

func (q *QueriesReference) GetGeographybyLocation(
	ctx context.Context, typeId int64,
	parentId int64, location string) (GeographyInfo, error) {
	row := q.db.QueryRowContext(ctx, getGeographybyLocation, typeId, parentId, location)
	var i GeographyInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.Location,
		&i.TypeId,
		&i.ParentId,
		&i.ZipCode,
		&i.Latitude,
		&i.Longitude,
		&i.AddressUrl,
		&i.SimpleLocation,
		&i.FullLocation,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getGeographybyUuId = `-- name: GetGeographybyUuId :one
SELECT 
Id, mr.UUId, Code, 
Short_Name, Location, Type_Id, Parent_Id, Zip_Code, Latitude, 
Longitude, Address_URL, Simple_Location, Full_Location, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Geography d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesReference) GetGeographybyUuId(ctx context.Context, uuid uuid.UUID) (GeographyInfo, error) {
	row := q.db.QueryRowContext(ctx, getGeographybyUuId, uuid)
	var i GeographyInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.Location,
		&i.TypeId,
		&i.ParentId,
		&i.ZipCode,
		&i.Latitude,
		&i.Longitude,
		&i.AddressUrl,
		&i.SimpleLocation,
		&i.FullLocation,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listGeography = `-- name: ListGeography:many
SELECT 
Id, mr.UUId, Code, 
Short_Name, Location, Type_Id, Parent_Id, Zip_Code, Latitude, 
Longitude, Address_URL, Simple_Location, Full_Location, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Geography d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Type_Id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListGeographyParams struct {
	TypeId int64 `json:"typeId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesReference) ListGeography(ctx context.Context, arg ListGeographyParams) ([]GeographyInfo, error) {
	rows, err := q.db.QueryContext(ctx, listGeography, arg.TypeId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GeographyInfo{}
	for rows.Next() {
		var i GeographyInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Code,
			&i.ShortName,
			&i.Location,
			&i.TypeId,
			&i.ParentId,
			&i.ZipCode,
			&i.Latitude,
			&i.Longitude,
			&i.AddressUrl,
			&i.SimpleLocation,
			&i.FullLocation,
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

const searchGeography = `-- name: ListGeography:many
SELECT 
d.Id, mr.UUId, Code, 
Short_Name, d.Location, Type_Id, Parent_Id, Zip_Code, Latitude, 
Longitude, Address_URL, Simple_Location, d.Full_Location, s.rnkfull, s.rnk, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM searchLocation($1,$2) s
INNER JOIN Geography d on d.Id = s.Id
INNER JOIN Main_Record mr on mr.UUId = d.UUId

ORDER BY rnkfull desc
LIMIT $2
OFFSET $3
`

type SearchGeographyParams struct {
	SearchText string `json:"search_text"`
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
}

func (q *QueriesReference) SearchGeography(ctx context.Context, arg SearchGeographyParams) ([]GeographyInfoSearch, error) {
	rows, err := q.db.QueryContext(ctx, searchGeography, arg.SearchText, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GeographyInfoSearch{}
	for rows.Next() {
		var i GeographyInfoSearch
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Code,
			&i.ShortName,
			&i.Location,
			&i.TypeId,
			&i.ParentId,
			&i.ZipCode,
			&i.Latitude,
			&i.Longitude,
			&i.AddressUrl,
			&i.SimpleLocation,
			&i.FullLocation,
			&i.RankFull,
			&i.Rank,
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

const updateGeography = `-- name: UpdateGeography :one
UPDATE Geography SET 
Code = $2,
Short_Name = $3,
Location = $4,
Type_Id = $5,
Parent_Id = $6,
Zip_Code = $7,
Latitude = $8,
Longitude = $9,
Address_URL = $10,
Other_Info = $11

WHERE id = $1
RETURNING Id, UUId, Code, Short_Name, Location, Type_Id, Parent_Id, 
Zip_Code, Latitude, Longitude, Address_URL, Simple_Location, Full_Location, 
Other_Info
`

func (q *QueriesReference) UpdateGeography(ctx context.Context, arg GeographyRequest) (model.Geography, error) {
	row := q.db.QueryRowContext(ctx, updateGeography,
		arg.Id,
		arg.Code,
		arg.ShortName,
		arg.Location,
		arg.TypeId,
		arg.ParentId,
		arg.ZipCode,
		arg.Latitude,
		arg.Longitude,
		arg.AddressUrl,
		arg.OtherInfo,
	)
	var i model.Geography
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Code,
		&i.ShortName,
		&i.Location,
		&i.TypeId,
		&i.ParentId,
		&i.ZipCode,
		&i.Latitude,
		&i.Longitude,
		&i.AddressUrl,
		&i.SimpleLocation,
		&i.FullLocation,
		&i.OtherInfo,
	)
	return i, err
}
