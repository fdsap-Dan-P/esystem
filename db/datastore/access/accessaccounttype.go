package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccessAccountType = `-- name: CreateAccessAccountType: one
INSERT INTO Access_Account_Type 
 (Role_Id, Type_Id, Allow, Other_Info) 
VALUES ($1, $2, $3, $4) 
ON CONFLICT(Role_Id, Type_Id)
DO UPDATE SET
  Allow = EXCLUDED.Allow,
  Other_Info = EXCLUDED.Other_Info
RETURNING UUID, Role_Id, Type_Id, Allow, Other_Info
`

type AccessAccountTypeRequest struct {
	Uuid      uuid.UUID      `json:"uuid"`
	RoleId    int64          `json:"roleId"`
	TypeId    int64          `json:"typeId"`
	Allow     sql.NullBool   `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`
}

func (q *QueriesAccess) CreateAccessAccountType(ctx context.Context, arg AccessAccountTypeRequest) (model.AccessAccountType, error) {
	row := q.db.QueryRowContext(ctx, createAccessAccountType,
		arg.RoleId,
		arg.TypeId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.AccessAccountType
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.TypeId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccessAccountType = `-- name: DeleteAccessAccountType :exec
DELETE FROM Access_Account_Type
WHERE Uuid = $1
`

func (q *QueriesAccess) DeleteAccessAccountType(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccessAccountType, uuid)
	return err
}

type AccessAccountTypeInfo struct {
	Uuid      uuid.UUID      `json:"uuid"`
	RoleId    int64          `json:"roleId"`
	TypeId    int64          `json:"typeId"`
	Allow     sql.NullBool   `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`
	ModCtr    int64          `json:"modCtr"`
	Created   sql.NullTime   `json:"created"`
	Updated   sql.NullTime   `json:"updated"`
}

const getAccessAccountType = `-- name: GetAccessAccountType :one
SELECT 
  mr.UUId, Role_Id, Type_Id, Allow, Other_Info,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Account_Type d 
INNER JOIN Main_Record mr on mr.Uuid = d.Uuid
WHERE mr.Uuid = $1 LIMIT 1
`

func (q *QueriesAccess) GetAccessAccountType(ctx context.Context, uuid uuid.UUID) (AccessAccountTypeInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccessAccountType, uuid)
	var i AccessAccountTypeInfo
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.TypeId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAccessAccountTypebyUuId = `-- name: GetAccessAccountTypebyUuId :one
SELECT 
mr.UUId, Role_Id, Type_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Account_Type d INNER JOIN Main_Record mr on mr.Uuid = d.UUId
WHERE mr.Uuid = $1 LIMIT 1
`

func (q *QueriesAccess) GetAccessAccountTypebyUuId(ctx context.Context, uuid uuid.UUID) (AccessAccountTypeInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccessAccountTypebyUuId, uuid)
	var i AccessAccountTypeInfo
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.TypeId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAccessAccountTypebyName = `-- name: GetAccessAccountTypebyName :one
SELECT 
mr.UUId, Role_Id, Type_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Account_Type d INNER JOIN Main_Record mr on mr.Uuid = d.UUId
WHERE Title = $1 LIMIT 1
`

func (q *QueriesAccess) GetAccessAccountTypebyName(ctx context.Context, name string) (AccessAccountTypeInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccessAccountTypebyName, name)
	var i AccessAccountTypeInfo
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.TypeId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listAccessAccountType = `-- name: ListAccessAccountType:many
SELECT 
  mr.UUId, Role_Id, Type_Id, Allow, Other_Info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Account_Type d 
INNER JOIN Main_Record mr on mr.Uuid = d.UUId
WHERE Role_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListAccessAccountTypeParams struct {
	RoleId int64 `json:"roleId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccess) ListAccessAccountType(ctx context.Context, arg ListAccessAccountTypeParams) ([]AccessAccountTypeInfo, error) {
	rows, err := q.db.QueryContext(ctx, listAccessAccountType, arg.RoleId, arg.Limit, arg.Offset)
	log.Printf("ListAccessAccountType: %v", arg)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccessAccountTypeInfo{}
	for rows.Next() {
		var i AccessAccountTypeInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.RoleId,
			&i.TypeId,
			&i.Allow,
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

const updateAccessAccountType = `-- name: UpdateAccessAccountType :one
UPDATE Access_Account_Type SET 
Role_Id = $2,
Type_Id = $3,
Allow = $4,
Other_Info = $5
WHERE Uuid = $1
RETURNING UUId, Role_Id, Type_Id, Allow, Other_Info
`

func (q *QueriesAccess) UpdateAccessAccountType(ctx context.Context, arg AccessAccountTypeRequest) (model.AccessAccountType, error) {
	row := q.db.QueryRowContext(ctx, updateAccessAccountType,
		arg.Uuid,
		arg.RoleId,
		arg.TypeId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.AccessAccountType
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.TypeId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}
