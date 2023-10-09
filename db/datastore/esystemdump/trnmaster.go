package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	model "simplebank/db/datastore/esystemlocal"
)

const createTrnMaster = `-- name: CreateTrnMaster: one
INSERT INTO esystemdump.TrnMaster(
   ModCtr, BrCode, ModAction, Acc, TrnDate, Trn, TrnType, OrNo, Prin, IntR, WaivedInt, RefNo, UserName, Particular )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
ON CONFLICT (brCode, trnDate, trn, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	Acc =  EXCLUDED.Acc,
	TrnType =  EXCLUDED.TrnType,
	OrNo =  EXCLUDED.OrNo,
	Prin =  EXCLUDED.Prin,
	IntR =  EXCLUDED.IntR,
	WaivedInt =  EXCLUDED.WaivedInt,
	RefNo =  EXCLUDED.RefNo,
	UserName =  EXCLUDED.UserName,
	Particular =  EXCLUDED.Particular
`

func (q *QueriesDump) CreateTrnMaster(ctx context.Context, arg model.TrnMaster) error {
	_, err := q.db.ExecContext(ctx, createTrnMaster,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.TrnDate,
		arg.Trn,
		arg.TrnType,
		arg.OrNo,
		arg.Prin,
		arg.IntR,
		arg.WaivedInt,
		arg.RefNo,
		arg.UserName,
		arg.Particular,
	)
	return err
}

const deleteTrnMaster = `-- name: DeleteTrnMaster :exec
DELETE FROM esystemdump.TrnMaster WHERE BrCode = $1 and TrnDate = $2 and Trn = $3
`

func (q *QueriesDump) DeleteTrnMaster(ctx context.Context, brCode string, trnDate time.Time, trn int64) error {
	_, err := q.db.ExecContext(ctx, deleteTrnMaster, brCode, trnDate, trn)
	return err
}

const getTrnMaster = `-- name: GetTrnMaster :one
SELECT
ModCtr, BrCode, ModAction, Acc, TrnDate, Trn, TrnType, OrNo, Prin, IntR, WaivedInt, RefNo, UserName, Particular
FROM esystemdump.TrnMaster
`

func scanRowTrnMaster(row *sql.Row) (model.TrnMaster, error) {
	var i model.TrnMaster
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.Acc,
		&i.TrnDate,
		&i.Trn,
		&i.TrnType,
		&i.OrNo,
		&i.Prin,
		&i.IntR,
		&i.WaivedInt,
		&i.RefNo,
		&i.UserName,
		&i.Particular,
	)
	return i, err
}

func scanRowsTrnMaster(rows *sql.Rows) ([]model.TrnMaster, error) {
	items := []model.TrnMaster{}
	for rows.Next() {
		var i model.TrnMaster
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.TrnDate,
			&i.Trn,
			&i.TrnType,
			&i.OrNo,
			&i.Prin,
			&i.IntR,
			&i.WaivedInt,
			&i.RefNo,
			&i.UserName,
			&i.Particular,
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

func (q *QueriesDump) GetTrnMaster(ctx context.Context, brCode string, trnDate time.Time, trn int64) (model.TrnMaster, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and TrnDate = $2 and Trn = $3", getTrnMaster)
	row := q.db.QueryRowContext(ctx, sql, brCode, trnDate, trn)
	return scanRowTrnMaster(row)
}

type ListTrnMasterParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListTrnMaster(ctx context.Context, lastModCtr int64) ([]model.TrnMaster, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getTrnMaster)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsTrnMaster(rows)
}

const updateTrnMaster = `-- name: UpdateTrnMaster :one
UPDATE esystemdump.TrnMaster SET 
	ModCtr = $1,
	Acc = $4,
	TrnType = $7,
	OrNo = $8,
	Prin = $9,
	IntR = $10,
	WaivedInt = $11,
	RefNo = $12,
	UserName = $13,
	Particular = $14
WHERE BrCode = $2 and TrnDate = $5 and Trn = $6 and ModAction = $3
`

func (q *QueriesDump) UpdateTrnMaster(ctx context.Context, arg model.TrnMaster) error {
	_, err := q.db.ExecContext(ctx, updateTrnMaster,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.TrnDate,
		arg.Trn,
		arg.TrnType,
		arg.OrNo,
		arg.Prin,
		arg.IntR,
		arg.WaivedInt,
		arg.RefNo,
		arg.UserName,
		arg.Particular,
	)
	return err
}
