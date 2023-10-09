package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	model "simplebank/db/datastore/esystemlocal"
)

const createSaTrnMaster = `-- name: CreateSaTrnMaster: one
INSERT INTO esystemdump.SaTrnMaster(
   ModCtr, BrCode, ModAction, Acc, TrnDate, Trn, TrnType, OrNo, TrnAmt, RefNo, Particular, TermId, UserName, PendApprove )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
ON CONFLICT (brCode, trnDate, trn, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	Acc =  EXCLUDED.Acc,
	TrnType =  EXCLUDED.TrnType,
	OrNo =  EXCLUDED.OrNo,
	TrnAmt =  EXCLUDED.TrnAmt,
	RefNo =  EXCLUDED.RefNo,
	Particular =  EXCLUDED.Particular,
	TermId =  EXCLUDED.TermId,
	UserName =  EXCLUDED.UserName,
	PendApprove = EXCLUDED.PendApprove
`

func (q *QueriesDump) CreateSaTrnMaster(ctx context.Context, arg model.SaTrnMaster) error {
	_, err := q.db.ExecContext(ctx, createSaTrnMaster,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.TrnDate,
		arg.Trn,
		arg.TrnType,
		arg.OrNo,
		arg.TrnAmt,
		arg.RefNo,
		arg.Particular,
		arg.TermId,
		arg.UserName,
		arg.PendApprove,
	)
	return err
}

const deleteSaTrnMaster = `-- name: DeleteSaTrnMaster :exec
DELETE FROM esystemdump.SaTrnMaster WHERE BrCode = $1 and TrnDate = $2 and Trn = $3
`

func (q *QueriesDump) DeleteSaTrnMaster(ctx context.Context, brCode string, trnDate time.Time, trn int64) error {
	_, err := q.db.ExecContext(ctx, deleteSaTrnMaster, brCode, trnDate, trn)
	return err
}

const getSaTrnMaster = `-- name: GetSaTrnMaster :one
SELECT
ModCtr, BrCode, ModAction, Acc, TrnDate, Trn, TrnType, OrNo, TrnAmt, RefNo, Particular, TermId, UserName, PendApprove
FROM esystemdump.SaTrnMaster
`

func scanRowSaTrnMaster(row *sql.Row) (model.SaTrnMaster, error) {
	var i model.SaTrnMaster
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.Acc,
		&i.TrnDate,
		&i.Trn,
		&i.TrnType,
		&i.OrNo,
		&i.TrnAmt,
		&i.RefNo,
		&i.Particular,
		&i.TermId,
		&i.UserName,
		&i.PendApprove,
	)
	return i, err
}

func scanRowsSaTrnMaster(rows *sql.Rows) ([]model.SaTrnMaster, error) {
	items := []model.SaTrnMaster{}
	for rows.Next() {
		var i model.SaTrnMaster
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.TrnDate,
			&i.Trn,
			&i.TrnType,
			&i.OrNo,
			&i.TrnAmt,
			&i.RefNo,
			&i.Particular,
			&i.TermId,
			&i.UserName,
			&i.PendApprove,
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

func (q *QueriesDump) GetSaTrnMaster(ctx context.Context, brCode string, trnDate time.Time, trn int64) (model.SaTrnMaster, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and TrnDate = $2 and Trn = $3", getSaTrnMaster)
	row := q.db.QueryRowContext(ctx, sql, brCode, trnDate, trn)
	return scanRowSaTrnMaster(row)
}

type ListSaTrnMasterParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListSaTrnMaster(ctx context.Context, lastModCtr int64) ([]model.SaTrnMaster, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getSaTrnMaster)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsSaTrnMaster(rows)
}

const updateSaTrnMaster = `-- name: UpdateSaTrnMaster :one
UPDATE esystemdump.SaTrnMaster SET 
	ModCtr = $1,
	Acc = $4,
	TrnType = $7,
	OrNo = $8,
	TrnAmt = $9,
	RefNo = $10,
	Particular = $11,
	TermId = $12,
	UserName = $13,
	PendApprove = $14
WHERE BrCode = $2 and TrnDate = $5 and Trn = $6 and ModAction = $3
`

func (q *QueriesDump) UpdateSaTrnMaster(ctx context.Context, arg model.SaTrnMaster) error {
	_, err := q.db.ExecContext(ctx, updateSaTrnMaster,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.TrnDate,
		arg.Trn,
		arg.TrnType,
		arg.OrNo,
		arg.TrnAmt,
		arg.RefNo,
		arg.Particular,
		arg.TermId,
		arg.UserName,
		arg.PendApprove,
	)
	return err
}
