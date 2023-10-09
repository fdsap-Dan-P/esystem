package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	model "simplebank/db/datastore/esystemlocal"
)

const createLedgerDetails = `-- name: CreateLedgerDetails: one
INSERT INTO esystemdump.LedgerDetails(
   ModCtr, BrCode, ModAction, TrnDate, Acc, Balance )
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (brCode, trnDate, acc, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	Balance =  EXCLUDED.Balance
`

func (q *QueriesDump) CreateLedgerDetails(ctx context.Context, arg model.LedgerDetails) error {
	_, err := q.db.ExecContext(ctx, createLedgerDetails,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.TrnDate,
		arg.Acc,
		arg.Balance,
	)
	return err
}

const deleteLedgerDetails = `-- name: DeleteLedgerDetails :exec
DELETE FROM esystemdump.LedgerDetails WHERE BrCode = $1 and TrnDate = $2 and Acc = $3
`

func (q *QueriesDump) DeleteLedgerDetails(ctx context.Context, brCode string, trnDate time.Time, acc string) error {
	_, err := q.db.ExecContext(ctx, deleteLedgerDetails, brCode, trnDate, acc)
	return err
}

const getLedgerDetails = `-- name: GetLedgerDetails :one
SELECT
ModCtr, BrCode, ModAction, TrnDate, Acc, Balance
FROM esystemdump.LedgerDetails
`

func scanRowLedgerDetails(row *sql.Row) (model.LedgerDetails, error) {
	var i model.LedgerDetails
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.TrnDate,
		&i.Acc,
		&i.Balance,
	)
	return i, err
}

func scanRowsLedgerDetails(rows *sql.Rows) ([]model.LedgerDetails, error) {
	items := []model.LedgerDetails{}
	for rows.Next() {
		var i model.LedgerDetails
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

func (q *QueriesDump) GetLedgerDetails(ctx context.Context, brCode string, trnDate time.Time, acc string) (model.LedgerDetails, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and TrnDate = $2 and Acc = $3", getLedgerDetails)
	row := q.db.QueryRowContext(ctx, sql, brCode, trnDate, acc)
	return scanRowLedgerDetails(row)
}

type ListLedgerDetailsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListLedgerDetails(ctx context.Context, lastModCtr int64) ([]model.LedgerDetails, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getLedgerDetails)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsLedgerDetails(rows)
}

const updateLedgerDetails = `-- name: UpdateLedgerDetails :one
UPDATE esystemdump.LedgerDetails SET 
	ModCtr = $1,
	Balance = $6
WHERE BrCode = $2 and TrnDate = $4 and Acc = $5 and ModAction = $3
`

func (q *QueriesDump) UpdateLedgerDetails(ctx context.Context, arg model.LedgerDetails) error {
	_, err := q.db.ExecContext(ctx, updateLedgerDetails,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.TrnDate,
		arg.Acc,
		arg.Balance,
	)
	return err
}
