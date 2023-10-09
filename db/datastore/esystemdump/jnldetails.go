package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createJnlDetails = `-- name: CreateJnlDetails: one
INSERT INTO esystemdump.JnlDetails(
   ModCtr, BrCode, ModAction, Acc, Trn, Series, Debit, Credit )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (brCode, acc, trn, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	Series =  EXCLUDED.Series,
	Debit =  EXCLUDED.Debit,
	Credit =  EXCLUDED.Credit
`

func (q *QueriesDump) CreateJnlDetails(ctx context.Context, arg model.JnlDetails) error {
	_, err := q.db.ExecContext(ctx, createJnlDetails,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.Trn,
		arg.Series,
		arg.Debit,
		arg.Credit,
	)
	return err
}

const deleteJnlDetails = `-- name: DeleteJnlDetails :exec
DELETE FROM esystemdump.JnlDetails WHERE BrCode = $1 and Acc = $2 and Trn = $3
`

func (q *QueriesDump) DeleteJnlDetails(ctx context.Context, brCode string, acc string, trn string) error {
	_, err := q.db.ExecContext(ctx, deleteJnlDetails, brCode, acc, trn)
	return err
}

const getJnlDetails = `-- name: GetJnlDetails :one
SELECT
ModCtr, BrCode, ModAction, Acc, Trn, Series, Debit, Credit
FROM esystemdump.JnlDetails
`

func scanRowJnlDetails(row *sql.Row) (model.JnlDetails, error) {
	var i model.JnlDetails
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.Acc,
		&i.Trn,
		&i.Series,
		&i.Debit,
		&i.Credit,
	)
	return i, err
}

func scanRowsJnlDetails(rows *sql.Rows) ([]model.JnlDetails, error) {
	items := []model.JnlDetails{}
	for rows.Next() {
		var i model.JnlDetails
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.Trn,
			&i.Series,
			&i.Debit,
			&i.Credit,
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

func (q *QueriesDump) GetJnlDetails(ctx context.Context, brCode string, acc string, trn string) (model.JnlDetails, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and Acc = $2 and Trn = $3", getJnlDetails)
	row := q.db.QueryRowContext(ctx, sql, brCode, acc, trn)
	return scanRowJnlDetails(row)
}

type ListJnlDetailsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListJnlDetails(ctx context.Context, lastModCtr int64) ([]model.JnlDetails, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getJnlDetails)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsJnlDetails(rows)
}

const updateJnlDetails = `-- name: UpdateJnlDetails :one
UPDATE esystemdump.JnlDetails SET 
ModCtr = $1,
Series = $6,
Debit = $7,
Credit = $8
WHERE BrCode = $2 and Acc = $4 and Trn = $5 and ModAction = $3
`

func (q *QueriesDump) UpdateJnlDetails(ctx context.Context, arg model.JnlDetails) error {
	_, err := q.db.ExecContext(ctx, updateJnlDetails,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.Trn,
		arg.Series,
		arg.Debit,
		arg.Credit,
	)
	return err
}
