package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccountSpecsDate = `-- name: CreateAccountSpecsDate: one
INSERT INTO Account_Specs_Date 
  (Account_Id, Specs_ID, Value, Value2) 
VALUES 
  ($1, $2, $3, $4) 
ON CONFLICT( Account_Id, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Account_Id, Specs_Code, Specs_ID, Value, Value2
`

type AccountSpecsDateRequest struct {
	Uuid      uuid.UUID `json:"uuid"`
	AccountId int64     `json:"accountId"`
	SpecsId   int64     `json:"specsId"`
	Value     time.Time `json:"value"`
	Value2    time.Time `json:"value2"`
}

func (q *QueriesAccount) CreateAccountSpecsDate(ctx context.Context, arg AccountSpecsDateRequest) (model.AccountSpecsDate, error) {
	row := q.db.QueryRowContext(ctx, createAccountSpecsDate,
		arg.AccountId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.AccountSpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const updateAccountSpecsDate = `-- name: UpdateAccountSpecsDate :one
UPDATE Account_Specs_Date SET 
Account_Id = $2,
	Specs_ID = $3,
	Value = $4,
	Value2 = $5
WHERE uuid = $1
RETURNING UUID, Account_Id, Specs_Code, Specs_ID, Value, Value2
`

func (q *QueriesAccount) UpdateAccountSpecsDate(ctx context.Context, arg AccountSpecsDateRequest) (model.AccountSpecsDate, error) {
	row := q.db.QueryRowContext(ctx, updateAccountSpecsDate,
		arg.Uuid,
		arg.AccountId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.AccountSpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const deleteAccountSpecsDate = `-- name: DeleteAccountSpecsDate :exec
DELETE FROM Account_Specs_Date
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteAccountSpecsDate(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountSpecsDate, uuid)
	return err
}

type AccountSpecsDateInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	AccountId       int64          `json:"accountId"`
	SpecsCode       string         `json:"specsCode"`
	SpecsId         int64          `json:"specsId"`
	Item            string         `json:"item"`
	ItemShortName   string         `json:"itemShortName"`
	ItemDescription sql.NullString `json:"itemDescription"`
	Value           time.Time      `json:"value"`
	Value2          time.Time      `json:"value2"`
	ModCtr          int64          `json:"modCtr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

func populateIdentitySpecDate(q *QueriesAccount, ctx context.Context, sql string) (AccountSpecsDateInfo, error) {
	var i AccountSpecsDateInfo
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

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateIdentitySpecDate2(q *QueriesAccount, ctx context.Context, sql string) ([]AccountSpecsDateInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []AccountSpecsDateInfo{}
	for rows.Next() {
		var i AccountSpecsDateInfo
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

const AccountSpecsDateSQL = `-- name: AccountSpecsDateSQL
SELECT 
  mr.UUID, d.Account_ID, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Specs_Date d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesAccount) GetAccountSpecsDate(ctx context.Context, accountId int64, specsId int64) (AccountSpecsDateInfo, error) {
	return populateIdentitySpecDate(q, ctx, fmt.Sprintf("%s WHERE d.Account_ID = %v and d.Specs_ID = %v",
		AccountSpecsDateSQL, accountId, specsId))
}

func (q *QueriesAccount) GetAccountSpecsDatebyUuid(ctx context.Context, uuid uuid.UUID) (AccountSpecsDateInfo, error) {
	return populateIdentitySpecDate(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", AccountSpecsDateSQL, uuid))
}

type ListAccountSpecsDateParams struct {
	AccountId int64 `json:"AccountId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountSpecsDate(ctx context.Context, arg ListAccountSpecsDateParams) ([]AccountSpecsDateInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Account_ID = %v LIMIT %d OFFSET %d",
			AccountSpecsDateSQL, arg.AccountId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Account_ID = %v ", AccountSpecsDateSQL, arg.AccountId)
	}
	return populateIdentitySpecDate2(q, ctx, sql)
}
