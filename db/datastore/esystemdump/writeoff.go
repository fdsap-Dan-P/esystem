package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createWriteoff = `-- name: CreateWriteoff: one
INSERT INTO esystemdump.Writeoff(
   ModCtr, BrCode, ModAction, Acc, DisbDate, Principal, Interest, BalPrin, BalInt, TrnDate, AcctType, Print, PostedBy, VerifiedBy )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
ON CONFLICT (brCode, acc, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	DisbDate =  EXCLUDED.DisbDate,
	Principal =  EXCLUDED.Principal,
	Interest =  EXCLUDED.Interest,
	BalPrin =  EXCLUDED.BalPrin,
	BalInt =  EXCLUDED.BalInt,
	TrnDate =  EXCLUDED.TrnDate,
	AcctType =  EXCLUDED.AcctType,
	Print =  EXCLUDED.Print,
	PostedBy =  EXCLUDED.PostedBy,
	VerifiedBy =  EXCLUDED.VerifiedBy
`

func (q *QueriesDump) CreateWriteoff(ctx context.Context, arg model.Writeoff) error {
	_, err := q.db.ExecContext(ctx, createWriteoff,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.DisbDate,
		arg.Principal,
		arg.Interest,
		arg.BalPrin,
		arg.BalInt,
		arg.TrnDate,
		arg.AcctType,
		arg.Print,
		arg.PostedBy,
		arg.VerifiedBy,
	)
	return err
}

const deleteWriteoff = `-- name: DeleteWriteoff :exec
DELETE FROM esystemdump.Writeoff WHERE BrCode = $1 and Acc = $2
`

func (q *QueriesDump) DeleteWriteoff(ctx context.Context, brCode string, acc string) error {
	_, err := q.db.ExecContext(ctx, deleteWriteoff, brCode, acc)
	return err
}

const getWriteoff = `-- name: GetWriteoff :one
SELECT
ModCtr, BrCode, ModAction, Acc, DisbDate, Principal, Interest, BalPrin, BalInt, TrnDate, AcctType, Print, PostedBy, VerifiedBy
FROM esystemdump.Writeoff
`

func scanRowWriteoff(row *sql.Row) (model.Writeoff, error) {
	var i model.Writeoff
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.Acc,
		&i.DisbDate,
		&i.Principal,
		&i.Interest,
		&i.BalPrin,
		&i.BalInt,
		&i.TrnDate,
		&i.AcctType,
		&i.Print,
		&i.PostedBy,
		&i.VerifiedBy,
	)
	return i, err
}

func scanRowsWriteoff(rows *sql.Rows) ([]model.Writeoff, error) {
	items := []model.Writeoff{}
	for rows.Next() {
		var i model.Writeoff
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.DisbDate,
			&i.Principal,
			&i.Interest,
			&i.BalPrin,
			&i.BalInt,
			&i.TrnDate,
			&i.AcctType,
			&i.Print,
			&i.PostedBy,
			&i.VerifiedBy,
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

func (q *QueriesDump) GetWriteoff(ctx context.Context, brCode string, acc string) (model.Writeoff, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and Acc = $2", getWriteoff)
	row := q.db.QueryRowContext(ctx, sql, brCode, acc)
	return scanRowWriteoff(row)
}

type ListWriteoffParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListWriteoff(ctx context.Context, lastModCtr int64) ([]model.Writeoff, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getWriteoff)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsWriteoff(rows)
}

const updateWriteoff = `-- name: UpdateWriteoff :one
UPDATE esystemdump.Writeoff SET 
	ModCtr = $1,
	DisbDate = $5,
	Principal = $6,
	Interest = $7,
	BalPrin = $8,
	BalInt = $9,
	TrnDate = $10,
	AcctType = $11,
	Print = $12,
	PostedBy = $13,
	VerifiedBy = $14
WHERE BrCode = $2 and Acc = $4 and modAction = $3
`

func (q *QueriesDump) UpdateWriteoff(ctx context.Context, arg model.Writeoff) error {
	_, err := q.db.ExecContext(ctx, updateWriteoff,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.DisbDate,
		arg.Principal,
		arg.Interest,
		arg.BalPrin,
		arg.BalInt,
		arg.TrnDate,
		arg.AcctType,
		arg.Print,
		arg.PostedBy,
		arg.VerifiedBy,
	)
	return err
}
