package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createUserObject = `-- name: CreateUserObject: one
INSERT INTO User_Object (
User_Id, Object_Id, Allow, Other_Info) 
VALUES ($1, $2, $3, $4) 
ON CONFLICT(User_ID, Object_ID) DO UPDATE SET
  Allow = EXCLUDED.Allow,
  Other_Info = EXCLUDED.Other_Info
RETURNING UUId, User_Id, Object_Id, Allow, Other_Info

`

type UserObjectRequest struct {
	Uuid      uuid.UUID      `json:"uuid"`
	UserId    int64          `json:"userId"`
	ObjectId  int64          `json:"objectId"`
	Allow     bool           `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`
}

func (q *QueriesUser) CreateUserObject(ctx context.Context, arg UserObjectRequest) (model.UserObject, error) {
	row := q.db.QueryRowContext(ctx, createUserObject,
		arg.UserId,
		arg.ObjectId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.UserObject
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ObjectId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}

const deleteUserObject = `-- name: DeleteUserObject :exec
DELETE FROM User_Object
WHERE uuid = $1
`

func (q *QueriesUser) DeleteUserObject(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserObject, uuid)
	return err
}

type UserObjectInfo struct {
	Uuid      uuid.UUID      `json:"uuid"`
	UserId    int64          `json:"userId"`
	ObjectId  int64          `json:"objectId"`
	Allow     bool           `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getUserObject = `-- name: GetUserObject :one
SELECT 
mr.UUId, User_Id, Object_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Object d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesUser) GetUserObject(ctx context.Context, uuid uuid.UUID) (UserObjectInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserObject, uuid)
	var i UserObjectInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ObjectId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUserObjectbyUuId = `-- name: GetUserObjectbyUuId :one
SELECT 
mr.UUId, User_Id, Object_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Object d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesUser) GetUserObjectbyUuId(ctx context.Context, uuid uuid.UUID) (UserObjectInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserObjectbyUuId, uuid)
	var i UserObjectInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ObjectId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listUserObject = `-- name: ListUserObject:many
SELECT 
mr.UUId, User_Id, Object_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Object d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE User_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListUserObjectParams struct {
	UserId int64 `json:"userId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesUser) ListUserObject(ctx context.Context, arg ListUserObjectParams) ([]UserObjectInfo, error) {
	rows, err := q.db.QueryContext(ctx, listUserObject, arg.UserId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserObjectInfo{}
	for rows.Next() {
		var i UserObjectInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.UserId,
			&i.ObjectId,
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

const updateUserObject = `-- name: UpdateUserObject :one
UPDATE User_Object SET 
User_Id = $2,
Object_Id = $3,
Allow = $4,
Other_Info = $5
WHERE uuid = $1
RETURNING UUId, User_Id, Object_Id, Allow, Other_Info
`

func (q *QueriesUser) UpdateUserObject(ctx context.Context, arg UserObjectRequest) (model.UserObject, error) {
	row := q.db.QueryRowContext(ctx, updateUserObject,

		arg.Uuid,
		arg.UserId,
		arg.ObjectId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.UserObject
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.ObjectId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}
