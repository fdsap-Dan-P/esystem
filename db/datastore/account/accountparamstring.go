package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccountParamString = `-- name: CreateAccountParamString: one
INSERT INTO Account_Param_String 
  (Param_ID, Item_ID, Value) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Param_ID, Item_ID ) DO UPDATE SET
  Value = excluded.Value
RETURNING 
  UUID, Param_ID, Item_Code, Item_ID, Value
`

type AccountParamStringRequest struct {
	Uuid    uuid.UUID `json:"uuid"`
	ParamId int64     `json:"ParamId"`
	ItemId  int64     `json:"itemid"`
	Value   string    `json:"value"`
}

func (q *QueriesAccount) CreateAccountParamString(ctx context.Context, arg AccountParamStringRequest) (model.AccountParamString, error) {
	row := q.db.QueryRowContext(ctx, createAccountParamString,
		arg.ParamId,
		arg.ItemId,
		arg.Value,
	)
	var i model.AccountParamString
	err := row.Scan(
		&i.Uuid,
		&i.ParamId,
		&i.ItemCode,
		&i.ItemId,
		&i.Value,
	)
	return i, err
}

const updateAccountParamString = `-- name: UpdateAccountParamString :one
UPDATE Account_Param_String SET 
Param_ID = $2,
  Item_ID = $3,
  Value = $4
WHERE uuid = $1
RETURNING UUID, Param_ID, Item_Code, Item_ID, Value
`

func (q *QueriesAccount) UpdateAccountParamString(ctx context.Context, arg AccountParamStringRequest) (model.AccountParamString, error) {
	row := q.db.QueryRowContext(ctx, updateAccountParamString,
		arg.Uuid,
		arg.ParamId,
		arg.ItemId,
		arg.Value,
	)
	var i model.AccountParamString
	err := row.Scan(
		&i.Uuid,
		&i.ParamId,
		&i.ItemCode,
		&i.ItemId,
		&i.Value,
	)
	return i, err
}

const deleteAccountParamString = `-- name: DeleteAccountParamString :exec
DELETE FROM Account_Param_String
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteAccountParamString(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountParamString, uuid)
	return err
}

type AccountParamStringInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	ParamId         int64          `json:"ParamId"`
	ItemCode        string         `json:"ItemCode"`
	ItemId          int64          `json:"itemid"`
	Item            string         `json:"item"`
	ItemShortName   string         `json:"itemShortName"`
	ItemDescription sql.NullString `json:"itemDescription"`
	Value           string         `json:"value"`
	ModCtr          int64          `json:"mod_ctr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

func populateAccountParamtring(q *QueriesAccount, ctx context.Context, sql string) (AccountParamStringInfo, error) {
	var i AccountParamStringInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.ParamId,
		&i.ItemCode,
		&i.ItemId,
		&i.Item,
		&i.ItemShortName,
		&i.ItemDescription,
		&i.Value,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccountParamtring2(q *QueriesAccount, ctx context.Context, sql string) ([]AccountParamStringInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []AccountParamStringInfo{}
	for rows.Next() {
		var i AccountParamStringInfo
		err := rows.Scan(
			&i.Uuid,
			&i.ParamId,
			&i.ItemCode,
			&i.ItemId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
			&i.Value,

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

const accountParamStringSQL = `-- name: accountParamStringSQL
SELECT 
  mr.UUID, d.Param_ID, d.Item_Code, d.Item_ID, ref.Title, ref.Short_Name Item_Short_Name, 
  ref.Remark Item_Description, d.Value,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Param_String d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Item_ID
`

func (q *QueriesAccount) GetAccountParamString(ctx context.Context, ParamId int64, ItemId int64) (AccountParamStringInfo, error) {
	return populateAccountParamtring(q, ctx, fmt.Sprintf("%s WHERE d.Param_ID = %v and d.Item_ID = %v",
		accountParamStringSQL, ParamId, ItemId))
}

func (q *QueriesAccount) GetAccountParamStringbyUuid(ctx context.Context, uuid uuid.UUID) (AccountParamStringInfo, error) {
	return populateAccountParamtring(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", accountParamStringSQL, uuid))
}

type ListAccountParamStringParams struct {
	ParamId int64 `json:"ParamId"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountParamString(ctx context.Context, arg ListAccountParamStringParams) ([]AccountParamStringInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Param_ID = %v LIMIT %d OFFSET %d",
			accountParamStringSQL, arg.ParamId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Param_ID = %v ", accountParamStringSQL, arg.ParamId)
	}
	return populateAccountParamtring2(q, ctx, sql)
}
