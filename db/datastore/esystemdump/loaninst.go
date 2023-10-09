package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createLoanInst = `-- name: CreateLoanInst: one
INSERT INTO esystemdump.LoanInst(
   ModCtr, BrCode, ModAction, Acc, Dnum, DueDate, InstFlag, DuePrin, DueInt, UpInt )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT (brCode, acc, dnum, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	DueDate =  EXCLUDED.DueDate,
	InstFlag =  EXCLUDED.InstFlag,
	DuePrin =  EXCLUDED.DuePrin,
	DueInt =  EXCLUDED.DueInt,
	UpInt =  EXCLUDED.UpInt
`

func (q *QueriesDump) CreateLoanInst(ctx context.Context, arg model.LoanInst) error {
	_, err := q.db.ExecContext(ctx, createLoanInst,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.Dnum,
		arg.DueDate,
		arg.InstFlag,
		arg.DuePrin,
		arg.DueInt,
		arg.UpInt,
	)
	return err
}

const deleteLoanInst = `-- name: DeleteLoanInst :exec
DELETE FROM esystemdump.LoanInst WHERE BrCode = $1 and Acc = $2 and Dnum = $3
`

func (q *QueriesDump) DeleteLoanInst(ctx context.Context, brCode string, acc string, dnum int64) error {
	_, err := q.db.ExecContext(ctx, deleteLoanInst, brCode, acc, dnum)
	return err
}

const getLoanInst = `-- name: GetLoanInst :one
SELECT
ModCtr, BrCode, ModAction, Acc, Dnum, DueDate, InstFlag, DuePrin, DueInt, UpInt
FROM esystemdump.LoanInst
`

func scanRowLoanInst(row *sql.Row) (model.LoanInst, error) {
	var i model.LoanInst
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.Acc,
		&i.Dnum,
		&i.DueDate,
		&i.InstFlag,
		&i.DuePrin,
		&i.DueInt,
		&i.UpInt,
	)
	return i, err
}

func scanRowsLoanInst(rows *sql.Rows) ([]model.LoanInst, error) {
	items := []model.LoanInst{}
	for rows.Next() {
		var i model.LoanInst
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.Dnum,
			&i.DueDate,
			&i.InstFlag,
			&i.DuePrin,
			&i.DueInt,
			&i.UpInt,
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

func (q *QueriesDump) GetLoanInst(ctx context.Context, brCode string, acc string, dnum int64) (model.LoanInst, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and Acc = $2 and Dnum = $3", getLoanInst)
	row := q.db.QueryRowContext(ctx, sql, brCode, acc, dnum)
	return scanRowLoanInst(row)
}

type ListLoanInstParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListLoanInst(ctx context.Context, lastModCtr int64) ([]model.LoanInst, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getLoanInst)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsLoanInst(rows)
}

const updateLoanInst = `-- name: UpdateLoanInst :one
UPDATE esystemdump.LoanInst SET 
	ModCtr = $1,
	DueDate = $6,
	InstFlag = $7,
	DuePrin = $8,
	DueInt = $9,
	UpInt = $10
WHERE BrCode = $2 and Acc = $4 and Dnum = $5 and ModAction = $3
`

func (q *QueriesDump) UpdateLoanInst(ctx context.Context, arg model.LoanInst) error {
	_, err := q.db.ExecContext(ctx, updateLoanInst,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.Dnum,
		arg.DueDate,
		arg.InstFlag,
		arg.DuePrin,
		arg.DueInt,
		arg.UpInt,
	)
	return err
}
