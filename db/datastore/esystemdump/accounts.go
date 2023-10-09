package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createAccounts = `-- name: CreateAccounts: one
INSERT INTO esystemdump.Accounts(
   ModCtr, BrCode, ModAction, Acc, Title, Category, Type, MainCD, Parent )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT (brCode, acc, ModAction)
DO UPDATE SET
  ModCtr =  EXCLUDED.ModCtr,
  Title =  EXCLUDED.Title,
  Category =  EXCLUDED.Category,
  Type =  EXCLUDED.Type,
  MainCD =  EXCLUDED.MainCD,
  Parent =  EXCLUDED.Parent;
`

func (q *QueriesDump) CreateAccounts(ctx context.Context, arg model.Accounts) error {
	_, err := q.db.ExecContext(ctx, createAccounts,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.Title,
		arg.Category,
		arg.Type,
		arg.MainCD,
		arg.Parent,
	)
	return err
}

const deleteAccounts = `-- name: DeleteAccounts :exec
DELETE FROM esystemdump.Accounts WHERE BrCode = $1 and Acc = $2
`

func (q *QueriesDump) DeleteAccounts(ctx context.Context, brCode string, acc string) error {
	_, err := q.db.ExecContext(ctx, deleteAccounts, brCode, acc)
	return err
}

const getAccounts = `-- name: GetAccounts :one
SELECT
ModCtr, BrCode, ModAction, Acc, Title, Category, Type, MainCD, Parent
FROM esystemdump.Accounts
`

func scanRowAccounts(row *sql.Row) (model.Accounts, error) {
	var i model.Accounts
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.Acc,
		&i.Title,
		&i.Category,
		&i.Type,
		&i.MainCD,
		&i.Parent,
	)
	return i, err
}

func scanRowsAccounts(rows *sql.Rows) ([]model.Accounts, error) {
	items := []model.Accounts{}
	for rows.Next() {
		var i model.Accounts
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.Title,
			&i.Category,
			&i.Type,
			&i.MainCD,
			&i.Parent,
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

func (q *QueriesDump) GetAccounts(ctx context.Context, brCode string, acc string) (model.Accounts, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and Acc = $2", getAccounts)
	row := q.db.QueryRowContext(ctx, sql, brCode, acc)
	return scanRowAccounts(row)
}

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListAccounts(ctx context.Context, lastModCtr int64) ([]model.Accounts, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getAccounts)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsAccounts(rows)
}

const updateAccounts = `-- name: UpdateAccounts :one
UPDATE esystemdump.Accounts SET 
  ModCtr = $1,
  Title = $5,
  Category = $6,
  Type = $7,
  MainCD = $8,
  Parent = $9
WHERE BrCode = $2 and ModAction = $3 and Acc = $4
`

func (q *QueriesDump) UpdateAccounts(ctx context.Context, arg model.Accounts) error {
	_, err := q.db.ExecContext(ctx, updateAccounts,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.Title,
		arg.Category,
		arg.Type,
		arg.MainCD,
		arg.Parent,
	)
	return err
}
