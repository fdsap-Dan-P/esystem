package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createAccountSpecsNumber = `-- name: CreateAccountSpecsNumber: one
INSERT INTO Account_Specs_Number 
  (Account_Id, Specs_ID, Value, Value2, measure_id) 
VALUES 
  ($1, $2, $3, $4, $5) 
ON CONFLICT( Account_Id, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Account_Id, Specs_Code, Specs_ID, Value, Value2, measure_id
`

type AccountSpecsNumberRequest struct {
	Uuid      uuid.UUID       `json:"uuid"`
	AccountId int64           `json:"accountId"`
	SpecsId   int64           `json:"specsId"`
	Value     decimal.Decimal `json:"value"`
	Value2    decimal.Decimal `json:"value2"`
	MeasureId sql.NullInt64   `json:"measureId"`
}

func (q *QueriesAccount) CreateAccountSpecsNumber(ctx context.Context, arg AccountSpecsNumberRequest) (model.AccountSpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, createAccountSpecsNumber,
		arg.AccountId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.AccountSpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const updateAccountSpecsNumber = `-- name: UpdateAccountSpecsNumber :one
UPDATE Account_Specs_Number SET 
  Account_Id = $2,
  Specs_ID = $3,
  Value = $4,
  Value2 = $5,
  Measure_Id = $6
WHERE uuid = $1
RETURNING UUID, Account_Id, Specs_Code, Specs_ID, Value, Value2, Measure_Id
`

func (q *QueriesAccount) UpdateAccountSpecsNumber(ctx context.Context, arg AccountSpecsNumberRequest) (model.AccountSpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, updateAccountSpecsNumber,
		arg.Uuid,
		arg.AccountId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.AccountSpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const deleteAccountSpecsNumber = `-- name: DeleteAccountSpecsNumber :exec
DELETE FROM Account_Specs_Number
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteAccountSpecsNumber(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountSpecsNumber, uuid)
	return err
}

type AccountSpecsNumberInfo struct {
	Uuid            uuid.UUID       `json:"uuid"`
	AccountId       int64           `json:"accountId"`
	SpecsCode       string          `json:"specsCode"`
	SpecsId         int64           `json:"specsId"`
	Item            string          `json:"item"`
	ItemShortName   string          `json:"itemShortName"`
	ItemDescription string          `json:"itemDescription"`
	Value           decimal.Decimal `json:"value"`
	Value2          decimal.Decimal `json:"value2"`
	MeasureId       sql.NullInt64   `json:"measureId"`
	Measure         sql.NullString  `json:"measure"`
	MeasureUnit     sql.NullString  `json:"measureUnit"`
	ModCtr          int64           `json:"mod_ctr"`
	Created         sql.NullTime    `json:"created"`
	Updated         sql.NullTime    `json:"updated"`
}

func populateIdentitySpecNumber(q *QueriesAccount, ctx context.Context, sql string) (AccountSpecsNumberInfo, error) {
	var i AccountSpecsNumberInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.SpecsCode,
		&i.SpecsId,
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

func populateIdentitySpecNumber2(q *QueriesAccount, ctx context.Context, sql string) ([]AccountSpecsNumberInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []AccountSpecsNumberInfo{}
	for rows.Next() {
		var i AccountSpecsNumberInfo
		err := rows.Scan(
			&i.Uuid,
			&i.AccountId,
			&i.SpecsCode,
			&i.SpecsId,
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

const accountSpecsNumberSQL = `-- name: accountSpecsNumberSQL
SELECT 
  mr.UUID, d.Account_Id, d.Specs_Code, d.Specs_ID, ref.Title Item, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2, d.Measure_Id, mea.Title Measure, mea.Short_Name Measure_Unit, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Specs_Number d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
LEFT JOIN Reference mea on mea.ID = Measure_Id`

func (q *QueriesAccount) GetAccountSpecsNumber(ctx context.Context, accountId int64, specsId int64) (AccountSpecsNumberInfo, error) {
	return populateIdentitySpecNumber(q, ctx, fmt.Sprintf("%s WHERE d.Account_Id = %v and d.Specs_ID = %v",
		accountSpecsNumberSQL, accountId, specsId))
}

func (q *QueriesAccount) GetAccountSpecsNumberbyUuid(ctx context.Context, uuid uuid.UUID) (AccountSpecsNumberInfo, error) {
	return populateIdentitySpecNumber(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", accountSpecsNumberSQL, uuid))
}

type ListAccountSpecsNumberParams struct {
	AccountId int64 `json:"AccountId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountSpecsNumber(ctx context.Context, arg ListAccountSpecsNumberParams) ([]AccountSpecsNumberInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Account_Id = %v LIMIT %d OFFSET %d",
			accountSpecsNumberSQL, arg.AccountId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Account_Id = %v ", accountSpecsNumberSQL, arg.AccountId)
	}
	return populateIdentitySpecNumber2(q, ctx, sql)
}
