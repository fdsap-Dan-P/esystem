package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createAccountParamNumber = `-- name: CreateAccountParamNumber: one
INSERT INTO Account_Param_Number 
  (Param_Id, Item_ID, Value, Value2, Measure_Id) 
VALUES 
  ($1, $2, $3, $4, $5) 
ON CONFLICT( Param_Id, Item_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Param_Id, Item_Code, Item_ID, Value, Value2, Measure_Id
`

type AccountParamNumberRequest struct {
	Uuid      uuid.UUID       `json:"uuid"`
	ParamId   int64           `json:"ParamId"`
	ItemId    int64           `json:"itemid"`
	Value     decimal.Decimal `json:"value"`
	Value2    decimal.Decimal `json:"value2"`
	MeasureId sql.NullInt64   `json:"measureId"`
}

func (q *QueriesAccount) CreateAccountParamNumber(ctx context.Context, arg AccountParamNumberRequest) (model.AccountParamNumber, error) {
	row := q.db.QueryRowContext(ctx, createAccountParamNumber,
		arg.ParamId,
		arg.ItemId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.AccountParamNumber
	err := row.Scan(
		&i.Uuid,
		&i.ParamId,
		&i.ItemCode,
		&i.ItemId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const updateAccountParamNumber = `-- name: UpdateAccountParamNumber :one
UPDATE Account_Param_Number SET 
  Param_Id = $2,
  Item_ID = $3,
  Value = $4,
  Value2 = $5,
  Measure_Id = $6
WHERE uuid = $1
RETURNING UUID, Param_Id, Item_Code, Item_ID, Value, Value2, Measure_Id
`

func (q *QueriesAccount) UpdateAccountParamNumber(ctx context.Context, arg AccountParamNumberRequest) (model.AccountParamNumber, error) {
	row := q.db.QueryRowContext(ctx, updateAccountParamNumber,
		arg.Uuid,
		arg.ParamId,
		arg.ItemId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.AccountParamNumber
	err := row.Scan(
		&i.Uuid,
		&i.ParamId,
		&i.ItemCode,
		&i.ItemId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const deleteAccountParamNumber = `-- name: DeleteAccountParamNumber :exec
DELETE FROM Account_Param_Number
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteAccountParamNumber(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountParamNumber, uuid)
	return err
}

type AccountParamNumberInfo struct {
	Uuid            uuid.UUID       `json:"uuid"`
	ParamId         int64           `json:"ParamId"`
	ItemCode        string          `json:"ItemCode"`
	ItemId          int64           `json:"itemid"`
	Item            string          `json:"item"`
	ItemShortName   string          `json:"itemShortName"`
	ItemDescription sql.NullString  `json:"itemDescription"`
	Value           decimal.Decimal `json:"value"`
	Value2          decimal.Decimal `json:"value2"`
	MeasureId       sql.NullInt64   `json:"measureId"`
	Measure         sql.NullString  `json:"measure"`
	MeasureUnit     sql.NullString  `json:"measureUnit"`
	ModCtr          int64           `json:"mod_ctr"`
	Created         sql.NullTime    `json:"created"`
	Updated         sql.NullTime    `json:"updated"`
}

func populateAccountParamNumber(q *QueriesAccount, ctx context.Context, sql string) (AccountParamNumberInfo, error) {
	var i AccountParamNumberInfo
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
		&i.MeasureId,
		&i.Measure,
		&i.MeasureUnit,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccountParamNumber2(q *QueriesAccount, ctx context.Context, sql string) ([]AccountParamNumberInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []AccountParamNumberInfo{}
	for rows.Next() {
		var i AccountParamNumberInfo
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
			&i.MeasureId,
			&i.Measure,
			&i.MeasureUnit,

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

const accountParamNumberSQL = `-- name: accountParamNumberSQL
SELECT 
  mr.UUID, d.Param_ID, d.Item_Code, d.Item_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2, d.Measure_Id, mea.Title Measure, mea.Short_Name Measure_Unit, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Param_Number d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Item_ID
LEFT JOIN Reference mea on mea.ID = Measure_Id`

func (q *QueriesAccount) GetAccountParamNumber(ctx context.Context, ParamId int64, ItemId int64) (AccountParamNumberInfo, error) {
	return populateAccountParamNumber(q, ctx, fmt.Sprintf("%s WHERE d.Param_Id = %v and d.Item_ID = %v",
		accountParamNumberSQL, ParamId, ItemId))
}

func (q *QueriesAccount) GetAccountParamNumberbyUuid(ctx context.Context, uuid uuid.UUID) (AccountParamNumberInfo, error) {
	return populateAccountParamNumber(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", accountParamNumberSQL, uuid))
}

type ListAccountParamNumberParams struct {
	ParamId int64 `json:"ParamId"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountParamNumber(ctx context.Context, arg ListAccountParamNumberParams) ([]AccountParamNumberInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Param_Id = %v LIMIT %d OFFSET %d",
			accountParamNumberSQL, arg.ParamId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Param_Id = %v ", accountParamNumberSQL, arg.ParamId)
	}
	return populateAccountParamNumber2(q, ctx, sql)
}
