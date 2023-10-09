package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccountSpecsString = `-- name: CreateAccountSpecsString: one
INSERT INTO Account_Specs_String 
  (Account_ID, Specs_ID, Value) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Account_ID, Specs_ID ) DO UPDATE SET
  Value = excluded.Value
RETURNING 
  UUID, Account_ID, Specs_Code, Specs_ID, Value
`

type AccountSpecsStringRequest struct {
	Uuid      uuid.UUID `json:"uuid"`
	AccountId int64     `json:"accountId"`
	SpecsId   int64     `json:"specsId"`
	Value     string    `json:"value"`
}

func (q *QueriesAccount) CreateAccountSpecsString(ctx context.Context, arg AccountSpecsStringRequest) (model.AccountSpecsString, error) {
	row := q.db.QueryRowContext(ctx, createAccountSpecsString,
		arg.AccountId,
		arg.SpecsId,
		arg.Value,
	)
	var i model.AccountSpecsString
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const updateAccountSpecsString = `-- name: UpdateAccountSpecsString :one
UPDATE Account_Specs_String SET 
Account_ID = $2,
  Specs_ID = $3,
  Value = $4
WHERE uuid = $1
RETURNING UUID, Account_ID, Specs_Code, Specs_ID, Value
`

func (q *QueriesAccount) UpdateAccountSpecsString(ctx context.Context, arg AccountSpecsStringRequest) (model.AccountSpecsString, error) {
	row := q.db.QueryRowContext(ctx, updateAccountSpecsString,
		arg.Uuid,
		arg.AccountId,
		arg.SpecsId,
		arg.Value,
	)
	var i model.AccountSpecsString
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const deleteAccountSpecsString = `-- name: DeleteAccountSpecsString :exec
DELETE FROM Account_Specs_String
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteAccountSpecsString(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountSpecsString, uuid)
	return err
}

type AccountSpecsStringInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	AccountId       int64        `json:"accountId"`
	SpecsCode       string       `json:"specsCode"`
	SpecsId         int64        `json:"specsId"`
	Item            string       `json:"item"`
	ItemShortName   string       `json:"itemShortName"`
	ItemDescription string       `json:"itemDescription"`
	Value           string       `json:"value"`
	ModCtr          int64        `json:"mod_ctr"`
	Created         sql.NullTime `json:"created"`
	Updated         sql.NullTime `json:"updated"`
}

func populateAccountSpecString(q *QueriesAccount, ctx context.Context, sql string) (AccountSpecsStringInfo, error) {
	var i AccountSpecsStringInfo
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

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccountSpecString2(q *QueriesAccount, ctx context.Context, sql string) ([]AccountSpecsStringInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []AccountSpecsStringInfo{}
	for rows.Next() {
		var i AccountSpecsStringInfo
		err := rows.Scan(
			&i.Uuid,
			&i.AccountId,
			&i.SpecsCode,
			&i.SpecsId,
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

const accountSpecsStringSQL = `-- name: accountSpecsStringSQL
SELECT 
  mr.UUID, d.Account_ID, d.Specs_Code, d.Specs_Id, ref.Title, ref.Short_Name Item_Short_Name, 
  ref.Remark Item_Description, d.Value,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Specs_String d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesAccount) GetAccountSpecsString(ctx context.Context, accountId int64, specsId int64) (AccountSpecsStringInfo, error) {
	return populateAccountSpecString(q, ctx, fmt.Sprintf("%s WHERE d.Account_ID = %v and d.Specs_ID = %v",
		accountSpecsStringSQL, accountId, specsId))
}

func (q *QueriesAccount) GetAccountSpecsStringbyUuid(ctx context.Context, uuid uuid.UUID) (AccountSpecsStringInfo, error) {
	return populateAccountSpecString(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", accountSpecsStringSQL, uuid))
}

type ListAccountSpecsStringParams struct {
	AccountId int64 `json:"AccountId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountSpecsString(ctx context.Context, arg ListAccountSpecsStringParams) ([]AccountSpecsStringInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Account_ID = %v LIMIT %d OFFSET %d",
			accountSpecsStringSQL, arg.AccountId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Account_ID = %v ", accountSpecsStringSQL, arg.AccountId)
	}
	return populateAccountSpecString2(q, ctx, sql)
}
