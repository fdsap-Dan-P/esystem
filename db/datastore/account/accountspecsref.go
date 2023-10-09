package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccountSpecsRef = `-- name: CreateAccountSpecsRef: one
INSERT INTO Account_Specs_Ref 
  (Account_Id, Specs_ID, Ref_Id) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Account_Id, Specs_ID ) DO UPDATE SET
Ref_Id = excluded.Ref_Id
RETURNING 
  UUID, Account_Id, Specs_Code, Specs_ID, Ref_Id
`

type AccountSpecsRefRequest struct {
	Uuid      uuid.UUID `json:"uuid"`
	AccountId int64     `json:"accountId"`
	SpecsId   int64     `json:"specsId"`
	RefId     int64     `json:"refId"`
}

func (q *QueriesAccount) CreateAccountSpecsRef(ctx context.Context, arg AccountSpecsRefRequest) (model.AccountSpecsRef, error) {
	row := q.db.QueryRowContext(ctx, createAccountSpecsRef,
		arg.AccountId,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.AccountSpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const updateAccountSpecsRef = `-- name: UpdateAccountSpecsRef :one
UPDATE Account_Specs_Ref SET 
  Account_Id = $2,
  Specs_ID = $3,
  Ref_ID = $4
WHERE uuid = $1
RETURNING UUID, Account_Id, Specs_Code, Specs_ID, Ref_ID
`

func (q *QueriesAccount) UpdateAccountSpecsRef(ctx context.Context, arg AccountSpecsRefRequest) (model.AccountSpecsRef, error) {
	row := q.db.QueryRowContext(ctx, updateAccountSpecsRef,
		arg.Uuid,
		arg.AccountId,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.AccountSpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const deleteAccountSpecsRef = `-- name: DeleteAccountSpecsRef :exec
DELETE FROM Account_Specs_Ref
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteAccountSpecsRef(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountSpecsRef, uuid)
	return err
}

type AccountSpecsRefInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	AccountId       int64          `json:"accountId"`
	SpecsId         int64          `json:"specsId"`
	SpecsCode       string         `json:"specsCode"`
	Item            string         `json:"item"`
	ItemShortName   string         `json:"itemShortName"`
	ItemDescription string         `json:"itemDescription"`
	RefId           int64          `json:"refId"`
	MeasureId       sql.NullInt64  `json:"measureId"`
	Measure         sql.NullString `json:"measure"`
	MeasureUnit     sql.NullString `json:"measureUnit"`
	ModCtr          int64          `json:"mod_ctr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

func populateAccountSpecRef(q *QueriesAccount, ctx context.Context, sql string) (AccountSpecsRefInfo, error) {
	var i AccountSpecsRefInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.SpecsCode,
		&i.SpecsId,
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

func populateAccountSpecRef2(q *QueriesAccount, ctx context.Context, sql string) ([]AccountSpecsRefInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []AccountSpecsRefInfo{}
	for rows.Next() {
		var i AccountSpecsRefInfo
		err := rows.Scan(
			&i.Uuid,
			&i.AccountId,
			&i.SpecsCode,
			&i.SpecsId,
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

const accountSpecsRefSQL = `-- name: accountSpecsRefSQL
SELECT 
  mr.UUID, d.Account_Id, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, d.Ref_Id,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Specs_Ref d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesAccount) GetAccountSpecsRef(ctx context.Context, accountId int64, specsId int64) (AccountSpecsRefInfo, error) {
	return populateAccountSpecRef(q, ctx, fmt.Sprintf("%s WHERE d.Account_Id = %v and d.Specs_ID = %v",
		accountSpecsRefSQL, accountId, specsId))
}

func (q *QueriesAccount) GetAccountSpecsRefbyUuid(ctx context.Context, uuid uuid.UUID) (AccountSpecsRefInfo, error) {
	return populateAccountSpecRef(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", accountSpecsRefSQL, uuid))
}

type ListAccountSpecsRefParams struct {
	AccountId int64 `json:"AccountId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountSpecsRef(ctx context.Context, arg ListAccountSpecsRefParams) ([]AccountSpecsRefInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Account_Id = %v LIMIT %d OFFSET %d",
			accountSpecsRefSQL, arg.AccountId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Account_Id = %v ", accountSpecsRefSQL, arg.AccountId)
	}
	return populateAccountSpecRef2(q, ctx, sql)
}
