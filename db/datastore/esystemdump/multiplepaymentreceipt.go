package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	model "simplebank/db/datastore/esystemlocal"
)

const createMultiplePaymentReceipt = `-- name: CreateMultiplePaymentReceipt: one
INSERT INTO esystemdump.MultiplePaymentReceipt(
	ModCtr, BrCode, ModAction, TrnDate, OrNo, CID, PrNo, UserName, TermId, AmtPaid
	)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT (BrCode, OrNo, ModAction) DO UPDATE SET
ModCtr = excluded.ModCtr, 
TrnDate = excluded.TrnDate, 
CID = excluded.CID, 
PrNo = excluded.PrNo, 
UserName = excluded.UserName, 
TermId = excluded.TermId, 
AmtPaid = excluded.AmtPaid
;
`

func (q *QueriesDump) CreateMultiplePaymentReceipt(ctx context.Context, arg model.MultiplePaymentReceipt) error {
	_, err := q.db.ExecContext(ctx, createMultiplePaymentReceipt,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.TrnDate,
		arg.OrNo,
		arg.CID,
		arg.PrNo,
		arg.UserName,
		arg.TermId,
		arg.AmtPaid,
	)
	return err
}

const deleteMultiplePaymentReceipt = `-- name: DeleteMultiplePaymentReceipt :exec
DELETE FROM esystemdump.MultiplePaymentReceipt WHERE BrCode = $1 and OrNo = $2
`

func (q *QueriesDump) DeleteMultiplePaymentReceipt(ctx context.Context, brCode string, orNo int64) error {
	_, err := q.db.ExecContext(ctx, deleteMultiplePaymentReceipt, brCode, orNo)
	return err
}

const getMultiplePaymentReceipt = `-- name: GetMultiplePaymentReceipt :one
SELECT 
  ModCtr, BrCode, ModAction, TrnDate, OrNo, CID, PrNo, UserName, TermId, AmtPaid
FROM esystemdump.MultiplePaymentReceipt
`

func scanRowMultiplePaymentReceipt(row *sql.Row) (model.MultiplePaymentReceipt, error) {
	var i model.MultiplePaymentReceipt
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.TrnDate,
		&i.OrNo,
		&i.CID,
		&i.PrNo,
		&i.UserName,
		&i.TermId,
		&i.AmtPaid,
	)
	return i, err
}

func scanRowsMultiplePaymentReceipt(rows *sql.Rows) ([]model.MultiplePaymentReceipt, error) {
	items := []model.MultiplePaymentReceipt{}
	for rows.Next() {
		var i model.MultiplePaymentReceipt
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.TrnDate,
			&i.OrNo,
			&i.CID,
			&i.PrNo,
			&i.UserName,
			&i.TermId,
			&i.AmtPaid,
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

func (q *QueriesDump) GetMultiplePaymentReceipt(ctx context.Context, brCode string, orNo int64) (model.MultiplePaymentReceipt, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and OrNo = $2", getMultiplePaymentReceipt)
	row := q.db.QueryRowContext(ctx, sql, brCode, orNo)
	return scanRowMultiplePaymentReceipt(row)
}

type ListMultiplePaymentReceiptParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListMultiplePaymentReceipt(ctx context.Context, modCtr int64) ([]model.MultiplePaymentReceipt, error) {
	sql := fmt.Sprintf(`%v WHERE  ModCtr > $1`, getMultiplePaymentReceipt)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, modCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsMultiplePaymentReceipt(rows)
}

const updateMultiplePaymentReceipt = `-- name: UpdateMultiplePaymentReceipt :one
UPDATE esystemdump.MultiplePaymentReceipt SET 
  ModCtr = $1,
  TrnDate = $4,
  CID = $6,
  PrNo = $7,
  UserName = $8,
  TermId = $9,
  AmtPaid = $10
WHERE BrCode = $2 and OrNo = $5 and ModAction = $3
`

func (q *QueriesDump) UpdateMultiplePaymentReceipt(ctx context.Context, arg model.MultiplePaymentReceipt) error {
	_, err := q.db.ExecContext(ctx, updateMultiplePaymentReceipt,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.TrnDate,
		arg.OrNo,
		arg.CID,
		arg.PrNo,
		arg.UserName,
		arg.TermId,
		arg.AmtPaid,
	)
	return err
}
