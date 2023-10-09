package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

const createLoanInst = `-- name: CreateLoanInst: one
INSERT INTO LoanInst (
	Acc, Dnum, DueDate, InstFlag, Prin, IntR, UpInt
) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
`

type LoanInstRequest struct {
	Acc      string          `json:"acc"`
	Dnum     int64           `json:"dnum"`
	DueDate  time.Time       `json:"dueDate"`
	InstFlag int64           `json:"instFlag"`
	DuePrin  decimal.Decimal `json:"duePrin"`
	DueInt   decimal.Decimal `json:"dueInt"`
	UpInt    decimal.Decimal `json:"upInt"`
}

func (q *QueriesLocal) CreateLoanInst(ctx context.Context, arg LoanInstRequest) error {
	_, err := q.db.ExecContext(ctx, createLoanInst,
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
DELETE FROM LoanInst WHERE Acc = $1 and dNum = $2
`

func (q *QueriesLocal) DeleteLoanInst(ctx context.Context, Acc string, dNum int64) error {
	_, err := q.db.ExecContext(ctx, deleteLoanInst, Acc, dNum)
	return err
}

type LoanInstInfo struct {
	ModCtr    int64           `json:"modCtr"`
	BrCode    string          `json:"brCode"`
	ModAction string          `json:"modAction"`
	Acc       string          `json:"acc"`
	Dnum      int64           `json:"dnum"`
	DueDate   time.Time       `json:"dueDate"`
	InstFlag  int64           `json:"instFlag"`
	DuePrin   decimal.Decimal `json:"duePrin"`
	DueInt    decimal.Decimal `json:"dueInt"`
	UpInt     decimal.Decimal `json:"upInt"`
}

// -- name: GetLoanInst :one
const getLoanInst = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, Acc, Dnum, DueDate, InstFlag, Prin, IntR, UpInt
FROM OrgParms, LoanInst d
INNER JOIN Modified m on m.UniqueKeyString1 = d.Acc and m.UniqueKeyInt1 = dNum
`

func scanRowLoanInst(row *sql.Row) (LoanInstInfo, error) {
	var i LoanInstInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
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

func scanRowsLoanInst(rows *sql.Rows) ([]LoanInstInfo, error) {
	items := []LoanInstInfo{}
	for rows.Next() {
		var i LoanInstInfo
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

func (q *QueriesLocal) GetLoanInst(ctx context.Context, acc string, dNum int64) (LoanInstInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'LoanInst' AND Uploaded = 0 and Acc = $1 and dNum = $2", getLoanInst)
	row := q.db.QueryRowContext(ctx, sql, acc, dNum)
	return scanRowLoanInst(row)
}

type ListLoanInstParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) LoanInstCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, Acc, Dnum, DueDate, InstFlag, Prin, IntR, UpInt
FROM OrgParms, LoanInst d
`, filenamePath)
}

func (q *QueriesLocal) ListLoanInst(ctx context.Context) ([]LoanInstInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'LoanInst' AND Uploaded = 0`,
		getLoanInst)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsLoanInst(rows)
}

// -- name: UpdateLoanInst :one
const updateLoanInst = `
UPDATE LoanInst SET 
  DueDate = $3,
  InstFlag = $4,
  Prin = $5,
  IntR = $6,
  UpInt = $7
WHERE Acc = $1 and dNum = $2`

func (q *QueriesLocal) UpdateLoanInst(ctx context.Context, arg LoanInstRequest) error {
	_, err := q.db.ExecContext(ctx, updateLoanInst,
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
