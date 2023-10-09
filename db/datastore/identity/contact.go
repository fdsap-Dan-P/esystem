package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createContact = `-- name: CreateContact: one
INSERT INTO Contact (
  IIID, Series, Contact, Type_Id, Other_Info) 
VALUES ($1, $2, $3, $4, $5) 
ON CONFLICT (IIID, Type_ID, Contact)
DO UPDATE SET 
  Series = excluded.Series,
  Other_Info = excluded.Other_Info
RETURNING UUId, IIID, Series, Contact, Type_Id, Other_Info
`

type ContactRequest struct {
	Uuid      uuid.UUID      `json:"uuid"`
	Iiid      int64          `json:"iiid"`
	Series    int16          `json:"series"`
	Contact   string         `json:"contact"`
	TypeId    int64          `json:"typeId"`
	OtherInfo sql.NullString `json:"otherInfo"`
}

func (q *QueriesIdentity) CreateContact(ctx context.Context, arg ContactRequest) (model.Contact, error) {
	row := q.db.QueryRowContext(ctx, createContact,
		arg.Iiid,
		arg.Series,
		arg.Contact,
		arg.TypeId,
		arg.OtherInfo,
	)
	var i model.Contact
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Contact,
		&i.TypeId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteContact = `-- name: DeleteContact :exec
DELETE FROM Contact
WHERE uuid = $1
`

func (q *QueriesIdentity) DeleteContact(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteContact, uuid)
	return err
}

type ContactInfo struct {
	Uuid      uuid.UUID      `json:"uuid"`
	Iiid      int64          `json:"iiid"`
	Series    int16          `json:"series"`
	Contact   string         `json:"contact"`
	TypeId    int64          `json:"typeId"`
	OtherInfo sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getContact = `-- name: GetContact :one
SELECT 
mr.UUId, 
IIId, Series, Contact, Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Contact d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesIdentity) GetContact(ctx context.Context, uuid uuid.UUID) (ContactInfo, error) {
	row := q.db.QueryRowContext(ctx, getContact, uuid)
	var i ContactInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Contact,
		&i.TypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getContactbyUuId = `-- name: GetContactbyUuId :one
SELECT 
mr.UUId, 
IIId, Series, Contact, Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Contact d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesIdentity) GetContactbyUuId(ctx context.Context, uuid uuid.UUID) (ContactInfo, error) {
	row := q.db.QueryRowContext(ctx, getContactbyUuId, uuid)
	var i ContactInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Contact,
		&i.TypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getContactbyName = `-- name: GetContactbyName :one
SELECT 
mr.UUId, 
IIId, Series, Contact, Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Contact d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Contact = $1 LIMIT 1
`

func (q *QueriesIdentity) GetContactbyName(ctx context.Context, name string) (ContactInfo, error) {
	row := q.db.QueryRowContext(ctx, getContactbyName, name)
	var i ContactInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Contact,
		&i.TypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listContact = `-- name: ListContact:many
SELECT 
mr.UUId, 
IIId, Series, Contact, Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Contact d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Iiid = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListContactParams struct {
	Iiid   int64 `json:"iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListContact(ctx context.Context, arg ListContactParams) ([]ContactInfo, error) {
	rows, err := q.db.QueryContext(ctx, listContact, arg.Iiid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ContactInfo{}
	for rows.Next() {
		var i ContactInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.Iiid,
			&i.Series,
			&i.Contact,
			&i.TypeId,
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

const updateContact = `-- name: UpdateContact :one
UPDATE Contact SET 
IIId = $2,
Series = $3,
Contact = $4,
Type_Id = $5,
Other_Info = $6
WHERE uuid = $1
RETURNING UUId, IIId, Series, Contact, Type_Id, Other_Info
`

func (q *QueriesIdentity) UpdateContact(ctx context.Context, arg ContactRequest) (model.Contact, error) {
	row := q.db.QueryRowContext(ctx, updateContact,

		arg.Uuid,
		arg.Iiid,
		arg.Series,
		arg.Contact,
		arg.TypeId,
		arg.OtherInfo,
	)
	var i model.Contact
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Contact,
		&i.TypeId,
		&i.OtherInfo,
	)
	return i, err
}
