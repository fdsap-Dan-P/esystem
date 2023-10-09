package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createUserOffice = `-- name: CreateUserOffice: one
INSERT INTO User_Office (
User_Id, Office_Id, Allow, Other_Info) 
VALUES ($1, $2, $3, $4) 
ON CONFLICT(User_ID, Office_Id) DO UPDATE SET
  Allow = EXCLUDED.Allow,
  Other_Info = EXCLUDED.Other_Info
RETURNING UUId, User_Id, Office_Id, Allow, Other_Info
`

type UserOfficeRequest struct {
	Uuid      uuid.UUID      `json:"uuid"`
	UserId    int64          `json:"userId"`
	OfficeId  int64          `json:"officeId"`
	Allow     bool           `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`
}

func (q *QueriesUser) CreateUserOffice(ctx context.Context, arg UserOfficeRequest) (model.UserOffice, error) {
	row := q.db.QueryRowContext(ctx, createUserOffice,
		arg.UserId,
		arg.OfficeId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.UserOffice
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.OfficeId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}

const deleteUserOffice = `-- name: DeleteUserOffice :exec
DELETE FROM User_Office
WHERE uuid = $1
`

func (q *QueriesUser) DeleteUserOffice(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserOffice, uuid)
	return err
}

type UserOfficeInfo struct {
	Uuid      uuid.UUID      `json:"uuid"`
	UserId    int64          `json:"userId"`
	OfficeId  int64          `json:"officeId"`
	Allow     bool           `json:"allow"`
	OtherInfo sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getUserOffice = `-- name: GetUserOffice :one
SELECT 
mr.UUId, User_Id, Office_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Office d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesUser) GetUserOffice(ctx context.Context, uuid uuid.UUID) (UserOfficeInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserOffice, uuid)
	var i UserOfficeInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.OfficeId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUserOfficebyUuId = `-- name: GetUserOfficebyUuId :one
SELECT 
mr.UUId, User_Id, Office_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Office d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesUser) GetUserOfficebyUuId(ctx context.Context, uuid uuid.UUID) (UserOfficeInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserOfficebyUuId, uuid)
	var i UserOfficeInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.OfficeId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUserOfficebyName = `-- name: GetUserOfficebyName :one
SELECT 
mr.UUId, User_Id, Office_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Office d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Title = $1 LIMIT 1
`

func (q *QueriesUser) GetUserOfficebyName(ctx context.Context, name string) (UserOfficeInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserOfficebyName, name)
	var i UserOfficeInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.OfficeId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listUserOffice = `-- name: ListUserOffice:many
SELECT 
mr.UUId, User_Id, Office_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Office d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE User_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListUserOfficeParams struct {
	UserId int64 `json:"userId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesUser) ListUserOffice(ctx context.Context, arg ListUserOfficeParams) ([]UserOfficeInfo, error) {
	rows, err := q.db.QueryContext(ctx, listUserOffice, arg.UserId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserOfficeInfo{}
	for rows.Next() {
		var i UserOfficeInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.UserId,
			&i.OfficeId,
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

const updateUserOffice = `-- name: UpdateUserOffice :one
UPDATE User_Office SET 
User_Id = $2,
Office_Id = $3,
Allow = $4,
Other_Info = $5
WHERE uuid = $1
RETURNING UUId, User_Id, Office_Id, Allow, Other_Info
`

func (q *QueriesUser) UpdateUserOffice(ctx context.Context, arg UserOfficeRequest) (model.UserOffice, error) {
	row := q.db.QueryRowContext(ctx, updateUserOffice,

		arg.Uuid,
		arg.UserId,
		arg.OfficeId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.UserOffice
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.OfficeId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}
