package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAddressList = `-- name: CreateAddressList: one
INSERT INTO Address_List (
IIID, Series, Detail, URL, Type_Id, 
Geography_Id, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7
) RETURNING UUId, IIId, Series, Detail, URL, Type_Id, 
Geography_Id, Other_Info
`

type AddressListRequest struct {
	Uuid        uuid.UUID      `json:"uuid"`
	Iiid        int64          `json:"iiid"`
	Series      int16          `json:"series"`
	Detail      sql.NullString `json:"detail"`
	Url         sql.NullString `json:"url"`
	TypeId      int64          `json:"typeId"`
	GeographyId sql.NullInt64  `json:"geographyId"`
	OtherInfo   sql.NullString `json:"otherInfo"`
}

func (q *QueriesIdentity) CreateAddressList(ctx context.Context, arg AddressListRequest) (model.AddressList, error) {
	row := q.db.QueryRowContext(ctx, createAddressList,
		arg.Iiid,
		arg.Series,
		arg.Detail,
		arg.Url,
		arg.TypeId,
		arg.GeographyId,
		arg.OtherInfo,
	)
	var i model.AddressList
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Detail,
		&i.Url,
		&i.TypeId,
		&i.GeographyId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAddressList = `-- name: DeleteAddressList :exec
DELETE FROM Address_List
WHERE uuid = $1
`

func (q *QueriesIdentity) DeleteAddressList(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAddressList, uuid)
	return err
}

type AddressListInfo struct {
	Uuid        uuid.UUID      `json:"uuid"`
	Iiid        int64          `json:"iiid"`
	Series      int16          `json:"series"`
	Detail      sql.NullString `json:"detail"`
	Url         sql.NullString `json:"url"`
	TypeId      int64          `json:"typeId"`
	GeographyId sql.NullInt64  `json:"geographyId"`
	OtherInfo   sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getAddressList = `-- name: GetAddressList :one
SELECT 
mr.UUId, IIId, Series, 
Detail, URL, Type_Id, Geography_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Address_List d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesIdentity) GetAddressList(ctx context.Context, uuid uuid.UUID) (AddressListInfo, error) {
	row := q.db.QueryRowContext(ctx, getAddressList, uuid)
	var i AddressListInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Detail,
		&i.Url,
		&i.TypeId,
		&i.GeographyId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAddressListbyUuId = `-- name: GetAddressListbyUuId :one
SELECT 
mr.UUId, IIId, Series, 
Detail, URL, Type_Id, Geography_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Address_List d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesIdentity) GetAddressListbyUuId(ctx context.Context, uuid uuid.UUID) (AddressListInfo, error) {
	row := q.db.QueryRowContext(ctx, getAddressListbyUuId, uuid)
	var i AddressListInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Detail,
		&i.Url,
		&i.TypeId,
		&i.GeographyId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAddressListbyName = `-- name: GetAddressListbyName :one
SELECT 
mr.UUId, IIId, Series, 
Detail, URL, Type_Id, Geography_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Address_List d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Title = $1 LIMIT 1
`

func (q *QueriesIdentity) GetAddressListbyName(ctx context.Context, name string) (AddressListInfo, error) {
	row := q.db.QueryRowContext(ctx, getAddressListbyName, name)
	var i AddressListInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Detail,
		&i.Url,
		&i.TypeId,
		&i.GeographyId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listAddressList = `-- name: ListAddressList:many
SELECT 
mr.UUId, IIId, Series, 
Detail, URL, Type_Id, Geography_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Address_List d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE IIID = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListAddressListParams struct {
	Iiid   int64 `json:"Iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListAddressList(ctx context.Context, arg ListAddressListParams) ([]AddressListInfo, error) {
	rows, err := q.db.QueryContext(ctx, listAddressList, arg.Iiid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AddressListInfo{}
	for rows.Next() {
		var i AddressListInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.Iiid,
			&i.Series,
			&i.Detail,
			&i.Url,
			&i.TypeId,
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

const updateAddressList = `-- name: UpdateAddressList :one
UPDATE Address_List SET 
IIId = $2,
Series = $3,
Detail = $4,
URL = $5,
Type_Id = $6,
Geography_Id = $7,
Other_Info = $8
WHERE uuid = $1
RETURNING UUId, IIId, Series, Detail, URL, Type_Id, 
Geography_Id, Other_Info
`

func (q *QueriesIdentity) UpdateAddressList(ctx context.Context, arg AddressListRequest) (model.AddressList, error) {
	row := q.db.QueryRowContext(ctx, updateAddressList,

		arg.Uuid,
		arg.Iiid,
		arg.Series,
		arg.Detail,
		arg.Url,
		arg.TypeId,
		arg.GeographyId,
		arg.OtherInfo,
	)
	var i model.AddressList
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Detail,
		&i.Url,
		&i.TypeId,
		&i.GeographyId,
		&i.OtherInfo,
	)
	return i, err
}
