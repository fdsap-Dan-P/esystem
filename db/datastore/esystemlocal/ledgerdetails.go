package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

const createLedgerDetails = `-- name: CreateLedgerDetails: one
INSERT INTO Ledger_Details (
	Ledgdet_Trndate, Ledgdet_Acct_Code, Ledgdet_Amount
) 
VALUES ($1, $2, $3) 
`

type LedgerDetailsRequest struct {
	TrnDate time.Time       `json:"trnDate"`
	Acc     string          `json:"acc"`
	Balance decimal.Decimal `json:"balance"`
}

func (q *QueriesLocal) CreateLedgerDetails(ctx context.Context, arg LedgerDetailsRequest) error {
	_, err := q.db.ExecContext(ctx, createLedgerDetails,
		arg.TrnDate,
		arg.Acc,
		arg.Balance,
	)
	return err
}

const deleteLedgerDetails = `-- name: DeleteLedgerDetails :exec
DELETE FROM Ledger_Details WHERE Ledgdet_Trndate = $1 and Ledgdet_Acct_Code = $2
`

func (q *QueriesLocal) DeleteLedgerDetails(ctx context.Context, trnDate time.Time, acc string) error {
	_, err := q.db.ExecContext(ctx, deleteLedgerDetails, trnDate, acc)
	return err
}

type LedgerDetailsInfo struct {
	ModCtr    int64           `json:"modCtr"`
	BrCode    string          `json:"brCode"`
	ModAction string          `json:"modAction"`
	TrnDate   time.Time       `json:"trnDate"`
	Acc       string          `json:"acc"`
	Balance   decimal.Decimal `json:"balance"`
}

// -- name: GetLedgerDetails :one
const getLedgerDetails = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, Ledgdet_Trndate, Ledgdet_Acct_Code, Ledgdet_Amount
FROM OrgParms, Ledger_Details d
INNER JOIN Modified m on m.UniqueKeyDate = Ledgdet_Trndate and m.UniqueKeyString1 = d.Ledgdet_Acct_Code 
`

func scanRowLedgerDetails(row *sql.Row) (LedgerDetailsInfo, error) {
	var i LedgerDetailsInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.TrnDate,
		&i.Acc,
		&i.Balance,
	)
	return i, err
}

func scanRowsLedgerDetails(rows *sql.Rows) ([]LedgerDetailsInfo, error) {
	items := []LedgerDetailsInfo{}
	for rows.Next() {
		var i LedgerDetailsInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.TrnDate,
			&i.Acc,
			&i.Balance,
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

func (q *QueriesLocal) GetLedgerDetails(ctx context.Context, trnDate time.Time, acc string) (LedgerDetailsInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'Ledger_Details' AND Uploaded = 0 and Ledgdet_Trndate = $1 and Ledgdet_Acct_Code = $2", getLedgerDetails)
	row := q.db.QueryRowContext(ctx, sql, trnDate, acc)
	return scanRowLedgerDetails(row)
}

type ListLedgerDetailsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) LedgerDetailsCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, Ledgdet_Trndate, 
  Ledgdet_Acct_Code, Ledgdet_Amount
FROM OrgParms, Ledger_Details d
INNER JOIN Modified m on m.UniqueKeyDate = Ledgdet_Trndate and m.UniqueKeyString1 = d.Ledgdet_Acct_Code 
`, filenamePath)
}

func (q *QueriesLocal) ListLedgerDetails(ctx context.Context) ([]LedgerDetailsInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'Ledger_Details' AND Uploaded = 0`,
		getLedgerDetails)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsLedgerDetails(rows)
}

// -- name: UpdateLedgerDetails :one
const updateLedgerDetails = `
UPDATE Ledger_Details SET 
	Ledgdet_Trndate = $1,
	Ledgdet_Acct_Code = $2,
	Ledgdet_Amount = $3
WHERE Ledgdet_Trndate = $1 and Ledgdet_Acct_Code = $2`

func (q *QueriesLocal) UpdateLedgerDetails(ctx context.Context, arg LedgerDetailsRequest) error {
	_, err := q.db.ExecContext(ctx, updateLedgerDetails,
		arg.TrnDate,
		arg.Acc,
		arg.Balance,
	)
	return err
}
