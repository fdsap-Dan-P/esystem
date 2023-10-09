package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createAccessObject = `-- name: CreateAccessObject: one
INSERT INTO Access_Object 
 (Role_Id, Object_Id, Allow, Max_Value, Other_Info) 
VALUES ($1, $2, $3, $4, $5) 
ON CONFLICT(Role_Id, Object_Id)
DO UPDATE SET
  Allow = EXCLUDED.Allow,
  Max_Value = EXCLUDED.Max_Value,
	Other_Info = EXCLUDED.Other_Info
RETURNING UUId, Role_Id, Object_Id, Allow, Max_Value, Other_Info
`

type AccessObjectRequest struct {
	Uuid      uuid.UUID       `json:"uuid"`
	RoleId    int64           `json:"RoleId"`
	ObjectId  int64           `json:"ObjectId"`
	Allow     sql.NullBool    `json:"allow"`
	MaxValue  decimal.Decimal `json:"maxValue"`
	OtherInfo sql.NullString  `json:"otherInfo"`
}

func (q *QueriesAccess) CreateAccessObject(ctx context.Context, arg AccessObjectRequest) (model.AccessObject, error) {
	row := q.db.QueryRowContext(ctx, createAccessObject,
		arg.RoleId,
		arg.ObjectId,
		arg.Allow,
		arg.MaxValue,
		arg.OtherInfo,
	)
	var i model.AccessObject
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ObjectId,
		&i.Allow,
		&i.MaxValue,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccessObject = `-- name: DeleteAccessObject :exec
DELETE FROM Access_Object
WHERE uuid = $1
`

func (q *QueriesAccess) DeleteAccessObject(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccessObject, uuid)
	return err
}

type AccessObjectInfo struct {
	Uuid      uuid.UUID       `json:"uuid"`
	RoleId    int64           `json:"RoleId"`
	ObjectId  int64           `json:"ObjectId"`
	Allow     sql.NullBool    `json:"allow"`
	MaxValue  decimal.Decimal `json:"maxValue"`
	OtherInfo sql.NullString  `json:"otherInfo"`
	ModCtr    int64           `json:"modCtr"`
	Created   sql.NullTime    `json:"created"`
	Updated   sql.NullTime    `json:"updated"`
}

const getAccessObject = `-- name: GetAccessObject :one
SELECT 
mr.UUId, 
Role_Id, Object_Id, Allow, Max_Value, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Object d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesAccess) GetAccessObject(ctx context.Context, uuid uuid.UUID) (AccessObjectInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccessObject, uuid)
	var i AccessObjectInfo
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ObjectId,
		&i.Allow,
		&i.MaxValue,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAccessObjectbyUuId = `-- name: GetAccessObjectbyUuId :one
SELECT 
mr.UUId, 
Role_Id, Object_Id, Allow, Max_Value, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Object d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccess) GetAccessObjectbyUuId(ctx context.Context, uuid uuid.UUID) (AccessObjectInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccessObjectbyUuId, uuid)
	var i AccessObjectInfo
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ObjectId,
		&i.Allow,
		&i.MaxValue,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listAccessObject = `-- name: ListAccessObject:many
SELECT 
mr.UUId, 
Role_Id, Object_Id, Allow, Max_Value, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Object d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Role_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListAccessObjectParams struct {
	RoleId int64 `json:"roleId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccess) ListAccessObject(ctx context.Context, arg ListAccessObjectParams) ([]AccessObjectInfo, error) {
	rows, err := q.db.QueryContext(ctx, listAccessObject, arg.RoleId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccessObjectInfo{}
	for rows.Next() {
		var i AccessObjectInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.RoleId,
			&i.ObjectId,
			&i.Allow,
			&i.MaxValue,
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

const updateAccessObject = `-- name: UpdateAccessObject :one
UPDATE Access_Object SET 
Role_Id = $2,
Object_Id = $3,
Allow = $4,
Max_Value = $5,
Other_Info = $6
WHERE uuid = $1
RETURNING UUId, Role_Id, Object_Id, Allow, Max_Value, Other_Info
`

func (q *QueriesAccess) UpdateAccessObject(ctx context.Context, arg AccessObjectRequest) (model.AccessObject, error) {
	row := q.db.QueryRowContext(ctx, updateAccessObject,

		arg.Uuid,
		arg.RoleId,
		arg.ObjectId,
		arg.Allow,
		arg.MaxValue,
		arg.OtherInfo,
	)
	var i model.AccessObject
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ObjectId,
		&i.Allow,
		&i.MaxValue,
		&i.OtherInfo,
	)
	return i, err
}
