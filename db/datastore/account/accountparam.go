package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"
	"time"

	"github.com/google/uuid"
)

const createaccountParam = `-- name: CreateAccountParam: one
INSERT INTO Account_Param(
   uuid, account_type_id, date_implemented, other_info )
VALUES ($1, $2, $3, $4)
ON CONFLICT(uuid)
DO UPDATE SET
  account_type_id = excluded.account_type_id,
  date_implemented = excluded.date_implemented,
  other_info = excluded.other_info

RETURNING id, uuid, account_type_id, date_implemented, other_info`

type AccountParamRequest struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	AccountTypeId   int64          `json:"accountTypeId"`
	DateImplemented time.Time      `json:"dateImplemented"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

func (q *QueriesAccount) CreateAccountParam(ctx context.Context, arg AccountParamRequest) (model.AccountParam, error) {
	row := q.db.QueryRowContext(ctx, createaccountParam,
		arg.Uuid,
		arg.AccountTypeId,
		arg.DateImplemented,
		arg.OtherInfo,
	)
	var i model.AccountParam
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountTypeId,
		&i.DateImplemented,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccountParam = `-- name: DeleteAccountParam :exec
DELETE FROM Account_Param
WHERE id = $1
`

func (q *QueriesAccount) DeleteAccountParam(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccountParam, id)
	return err
}

type AccountParamInfo struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	AccountTypeId   int64          `json:"accountTypeId"`
	DateImplemented time.Time      `json:"dateImplemented"`
	OtherInfo       sql.NullString `json:"otherInfo"`
	ModCtr          int64          `json:"modCtr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

const accountParamSQL = `-- name: AccountParamSQL :one
SELECT
  Id, mr.UUID, account_type_id, date_implemented, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Param d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateAccountParam(q *QueriesAccount, ctx context.Context, sql string) (AccountParamInfo, error) {
	var i AccountParamInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountTypeId,
		&i.DateImplemented,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccountParams(q *QueriesAccount, ctx context.Context, sql string) ([]AccountParamInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountParamInfo{}
	for rows.Next() {
		var i AccountParamInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.AccountTypeId,
			&i.DateImplemented,
			&i.OtherInfo,
			&i.ModCtr,
			&i.Created,
			&i.Updated,
		)
		if err != nil {
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

func (q *QueriesAccount) GetAccountParam(ctx context.Context, id int64) (AccountParamInfo, error) {
	return populateAccountParam(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", accountParamSQL, id))
}

func (q *QueriesAccount) GetAccountParambyUuid(ctx context.Context, uuid uuid.UUID) (AccountParamInfo, error) {
	return populateAccountParam(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", accountParamSQL, uuid))
}

type ListAccountParamParams struct {
	AccountTypeId int64 `json:"accountTypeId"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountParam(ctx context.Context, arg ListAccountParamParams) ([]AccountParamInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d",
			accountParamSQL, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf(accountParamSQL)
	}
	return populateAccountParams(q, ctx, sql)
}

const updateAccountParam = `-- name: UpdateAccountParam :one
UPDATE Account_Param SET 
  uuid = $2,
  account_type_id = $3,
  date_implemented = $4,
  other_info = $5
WHERE id = $1
RETURNING id, uuid, account_type_id, date_implemented, other_info
`

func (q *QueriesAccount) UpdateAccountParam(ctx context.Context, arg AccountParamRequest) (model.AccountParam, error) {
	row := q.db.QueryRowContext(ctx, updateAccountParam,
		arg.Id,
		arg.Uuid,
		arg.AccountTypeId,
		arg.DateImplemented,
		arg.OtherInfo,
	)
	var i model.AccountParam
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountTypeId,
		&i.DateImplemented,
		&i.OtherInfo,
	)
	return i, err
}
