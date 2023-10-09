package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccountParamRef = `-- name: CreateAccountParamRef: one
INSERT INTO Account_Param_Ref 
  (Param_Id, Item_ID, Ref_Id) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Param_Id, Item_ID ) DO UPDATE SET
Ref_Id = excluded.Ref_Id
RETURNING 
  UUID, Param_Id, Item_Code, Item_ID, Ref_Id
`

type AccountParamRefRequest struct {
	Uuid    uuid.UUID `json:"uuid"`
	ParamId int64     `json:"ParamId"`
	ItemId  int64     `json:"itemid"`
	RefId   int64     `json:"refId"`
}

func (q *QueriesAccount) CreateAccountParamRef(ctx context.Context, arg AccountParamRefRequest) (model.AccountParamRef, error) {
	row := q.db.QueryRowContext(ctx, createAccountParamRef,
		arg.ParamId,
		arg.ItemId,
		arg.RefId,
	)
	var i model.AccountParamRef
	err := row.Scan(
		&i.Uuid,
		&i.ParamId,
		&i.ItemCode,
		&i.ItemId,
		&i.RefId,
	)
	return i, err
}

const updateAccountParamRef = `-- name: UpdateAccountParamRef :one
UPDATE Account_Param_Ref SET 
  Param_Id = $2,
  Item_ID = $3,
  Ref_ID = $4
WHERE uuid = $1
RETURNING UUID, Param_Id, Item_Code, Item_ID, Ref_ID
`

func (q *QueriesAccount) UpdateAccountParamRef(ctx context.Context, arg AccountParamRefRequest) (model.AccountParamRef, error) {
	row := q.db.QueryRowContext(ctx, updateAccountParamRef,
		arg.Uuid,
		arg.ParamId,
		arg.ItemId,
		arg.RefId,
	)
	var i model.AccountParamRef
	err := row.Scan(
		&i.Uuid,
		&i.ParamId,
		&i.ItemCode,
		&i.ItemId,
		&i.RefId,
	)
	return i, err
}

const deleteAccountParamRef = `-- name: DeleteAccountParamRef :exec
DELETE FROM Account_Param_Ref
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteAccountParamRef(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountParamRef, uuid)
	return err
}

type AccountParamRefInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	ParamId         int64          `json:"ParamId"`
	ItemId          int64          `json:"itemid"`
	ItemCode        string         `json:"ItemCode"`
	Item            string         `json:"item"`
	ItemShortName   string         `json:"itemShortName"`
	ItemDescription sql.NullString `json:"itemDescription"`
	RefId           int64          `json:"refId"`
	MeasureId       sql.NullInt64  `json:"measureId"`
	Measure         sql.NullString `json:"measure"`
	MeasureUnit     sql.NullString `json:"measureUnit"`
	ModCtr          int64          `json:"mod_ctr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

func populateAccountParamRef(q *QueriesAccount, ctx context.Context, sql string) (AccountParamRefInfo, error) {
	var i AccountParamRefInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.ParamId,
		&i.ItemCode,
		&i.ItemId,
		&i.Item,
		&i.ItemShortName,
		&i.ItemDescription,
		&i.RefId,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccountParamRef2(q *QueriesAccount, ctx context.Context, sql string) ([]AccountParamRefInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []AccountParamRefInfo{}
	for rows.Next() {
		var i AccountParamRefInfo
		err := rows.Scan(
			&i.Uuid,
			&i.ParamId,
			&i.ItemCode,
			&i.ItemId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
			&i.RefId,

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

const accountParamRefSQL = `-- name: accountParamRefSQL
SELECT 
  mr.UUID, d.Param_ID, d.Item_Code, d.Item_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, d.Ref_Id,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Param_Ref d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Item_ID
`

func (q *QueriesAccount) GetAccountParamRef(ctx context.Context, ParamId int64, ItemId int64) (AccountParamRefInfo, error) {
	return populateAccountParamRef(q, ctx, fmt.Sprintf("%s WHERE d.Param_Id = %v and d.Item_ID = %v",
		accountParamRefSQL, ParamId, ItemId))
}

func (q *QueriesAccount) GetAccountParamRefbyUuid(ctx context.Context, uuid uuid.UUID) (AccountParamRefInfo, error) {
	return populateAccountParamRef(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", accountParamRefSQL, uuid))
}

type ListAccountParamRefParams struct {
	ParamId int64 `json:"ParamId"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountParamRef(ctx context.Context, arg ListAccountParamRefParams) ([]AccountParamRefInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Param_Id = %v LIMIT %d OFFSET %d",
			accountParamRefSQL, arg.ParamId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Param_Id = %v ", accountParamRefSQL, arg.ParamId)
	}
	return populateAccountParamRef2(q, ctx, sql)
}
