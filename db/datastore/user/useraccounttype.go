package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createUserAccountType = `-- name: CreateUserAccountType: one
INSERT INTO User_Account_Type (
UUID, User_Id, Account_Type_Id, Allow, Other_Info
) VALUES (
$1, $2, $3, $4, $5
) 
ON CONFLICT(UUID) DO UPDATE SET
  User_Id = EXCLUDED.User_Id,
  Account_Type_Id = EXCLUDED.Account_Type_Id,
  Allow = EXCLUDED.Allow,
  Other_Info = EXCLUDED.Other_Info

RETURNING UUID, User_Id, Account_Type_Id, Allow, Other_Info
`

type UserAccountTypeRequest struct {
	Uuid          uuid.UUID      `json:"uuid"`
	UserId        int64          `json:"userId"`
	AccountTypeId int64          `json:"accountTypeId"`
	Allow         bool           `json:"allow"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}

func (q *QueriesUser) CreateUserAccountType(ctx context.Context, arg UserAccountTypeRequest) (model.UserAccountType, error) {
	row := q.db.QueryRowContext(ctx, createUserAccountType,
		arg.Uuid,
		arg.UserId,
		arg.AccountTypeId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.UserAccountType
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.AccountTypeId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}

const deleteUserAccountType = `-- name: DeleteUserAccountType :exec
DELETE FROM User_Account_Type
WHERE uuid = $1
`

func (q *QueriesUser) DeleteUserAccountType(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserAccountType, uuid)
	return err
}

type UserAccountTypeInfo struct {
	Uuid          uuid.UUID      `json:"uuid"`
	UserId        int64          `json:"userId"`
	AccountTypeId int64          `json:"accountTypeId"`
	Allow         bool           `json:"allow"`
	OtherInfo     sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getUserAccountType = `-- name: GetUserAccountType :one
SELECT 
mr.UUId, User_Id, Account_Type_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Account_Type d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesUser) GetUserAccountType(ctx context.Context, uuid uuid.UUID) (UserAccountTypeInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserAccountType, uuid)
	var i UserAccountTypeInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.AccountTypeId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUserAccountTypebyUuId = `-- name: GetUserAccountTypebyUuId :one
SELECT 
mr.UUId, User_Id, Account_Type_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Account_Type d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesUser) GetUserAccountTypebyUuId(ctx context.Context, uuid uuid.UUID) (UserAccountTypeInfo, error) {
	row := q.db.QueryRowContext(ctx, getUserAccountTypebyUuId, uuid)
	var i UserAccountTypeInfo
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.AccountTypeId,
		&i.Allow,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listUserAccountType = `-- name: ListUserAccountType:many
SELECT 
mr.UUId, User_Id, Account_Type_Id, Allow, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM User_Account_Type d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE User_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListUserAccountTypeParams struct {
	UserId int64 `json:"userId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesUser) ListUserAccountType(ctx context.Context, arg ListUserAccountTypeParams) ([]UserAccountTypeInfo, error) {
	rows, err := q.db.QueryContext(ctx, listUserAccountType, arg.UserId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserAccountTypeInfo{}
	for rows.Next() {
		var i UserAccountTypeInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.UserId,
			&i.AccountTypeId,
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

const updateUserAccountType = `-- name: UpdateUserAccountType :one
UPDATE User_Account_Type SET 
User_Id = $2,
Account_Type_Id = $3,
Allow = $4,
Other_Info = $5
WHERE uuid = $1
RETURNING UUId, User_Id, Account_Type_Id, Allow, Other_Info
`

func (q *QueriesUser) UpdateUserAccountType(ctx context.Context, arg UserAccountTypeRequest) (model.UserAccountType, error) {
	row := q.db.QueryRowContext(ctx, updateUserAccountType,

		arg.Uuid,
		arg.UserId,
		arg.AccountTypeId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.UserAccountType
	err := row.Scan(
		&i.Uuid,
		&i.UserId,
		&i.AccountTypeId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}
