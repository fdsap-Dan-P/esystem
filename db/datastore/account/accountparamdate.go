package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccountParamDate = `-- name: CreateAccountParamDate: one
INSERT INTO Account_Param_Date 
  (Param_Id, Item_ID, Value, Value2) 
VALUES 
  ($1, $2, $3, $4) 
ON CONFLICT( Param_Id, Item_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Param_Id, Item_Code, Item_ID, Value, Value2
`

type AccountParamDateRequest struct {
	Uuid    uuid.UUID `json:"uuid"`
	ParamId int64     `json:"ParamId"`
	ItemId  int64     `json:"itemid"`
	Value   time.Time `json:"value"`
	Value2  time.Time `json:"value2"`
}

func (q *QueriesAccount) CreateAccountParamDate(ctx context.Context, arg AccountParamDateRequest) (model.AccountParamDate, error) {
	row := q.db.QueryRowContext(ctx, createAccountParamDate,
		arg.ParamId,
		arg.ItemId,
		arg.Value,
		arg.Value2,
	)
	var i model.AccountParamDate
	err := row.Scan(
		&i.Uuid,
		&i.ParamId,
		&i.ItemCode,
		&i.ItemId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const updateAccountParamDate = `-- name: UpdateAccountParamDate :one
UPDATE Account_Param_Date SET 
Param_Id = $2,
	Item_ID = $3,
	Value = $4,
	Value2 = $5
WHERE uuid = $1
RETURNING UUID, Param_Id, Item_Code, Item_ID, Value, Value2
`

func (q *QueriesAccount) UpdateAccountParamDate(ctx context.Context, arg AccountParamDateRequest) (model.AccountParamDate, error) {
	row := q.db.QueryRowContext(ctx, updateAccountParamDate,
		arg.Uuid,
		arg.ParamId,
		arg.ItemId,
		arg.Value,
		arg.Value2,
	)
	var i model.AccountParamDate
	err := row.Scan(
		&i.Uuid,
		&i.ParamId,
		&i.ItemCode,
		&i.ItemId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const deleteAccountParamDate = `-- name: DeleteAccountParamDate :exec
DELETE FROM Account_Param_Date
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteAccountParamDate(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountParamDate, uuid)
	return err
}

type AccountParamDateInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	ParamId         int64          `json:"paramId"`
	ItemCode        string         `json:"itemCode"`
	ItemId          int64          `json:"itemid"`
	Item            string         `json:"item"`
	ItemShortName   string         `json:"itemShortName"`
	ItemDescription sql.NullString `json:"itemDescription"`
	Value           time.Time      `json:"value"`
	Value2          time.Time      `json:"value2"`
	ModCtr          int64          `json:"modCtr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

func populateAccountParamDate(q *QueriesAccount, ctx context.Context, sql string) (AccountParamDateInfo, error) {
	var i AccountParamDateInfo
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
		&i.Value2,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccountParamDate2(q *QueriesAccount, ctx context.Context, sql string) ([]AccountParamDateInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []AccountParamDateInfo{}
	for rows.Next() {
		var i AccountParamDateInfo
		err := rows.Scan(
			&i.Uuid,
			&i.ParamId,
			&i.ItemCode,
			&i.ItemId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
			&i.Value,
			&i.Value2,

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

const AccountParamDateSQL = `-- name: AccountParamDateSQL
SELECT 
  mr.UUID, d.Param_ID, d.Item_Code, d.Item_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Param_Date d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Item_ID
`

func (q *QueriesAccount) GetAccountParamDate(ctx context.Context, ParamId int64, ItemId int64) (AccountParamDateInfo, error) {
	return populateAccountParamDate(q, ctx, fmt.Sprintf("%s WHERE d.Param_ID = %v and d.Item_ID = %v",
		AccountParamDateSQL, ParamId, ItemId))
}

func (q *QueriesAccount) GetAccountParamDatebyUuid(ctx context.Context, uuid uuid.UUID) (AccountParamDateInfo, error) {
	return populateAccountParamDate(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", AccountParamDateSQL, uuid))
}

type ListAccountParamDateParams struct {
	ParamId int64 `json:"ParamId"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountParamDate(ctx context.Context, arg ListAccountParamDateParams) ([]AccountParamDateInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Param_ID = %v LIMIT %d OFFSET %d",
			AccountParamDateSQL, arg.ParamId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Param_ID = %v ", AccountParamDateSQL, arg.ParamId)
	}
	return populateAccountParamDate2(q, ctx, sql)
}
