package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createLnChrgData = `-- name: CreateLnChrgData: one
INSERT INTO esystemdump.LnChrgData(
   ModCtr, BrCode, ModAction, Acc, ChrgCode, RefAcc, ChrAmnt )
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (brCode, acc, chrgCode, refAcc, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	ChrAmnt =  EXCLUDED.ChrAmnt
`

func (q *QueriesDump) CreateLnChrgData(ctx context.Context, arg model.LnChrgData) error {
	_, err := q.db.ExecContext(ctx, createLnChrgData,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.ChrgCode,
		arg.RefAcc,
		arg.ChrAmnt,
	)
	return err
}

const deleteLnChrgData = `-- name: DeleteLnChrgData :exec
DELETE FROM esystemdump.LnChrgData WHERE BrCode = $1 and Acc = $2 and ChrgCode = $3 and RefAcc = $4
`

func (q *QueriesDump) DeleteLnChrgData(ctx context.Context, brCode string, acc string, chrgCode int64, refAcc string) error {
	_, err := q.db.ExecContext(ctx, deleteLnChrgData, brCode, acc, chrgCode, refAcc)
	return err
}

const getLnChrgData = `-- name: GetLnChrgData :one
SELECT
ModCtr, BrCode, ModAction, Acc, ChrgCode, RefAcc, ChrAmnt
FROM esystemdump.LnChrgData
`

func scanRowLnChrgData(row *sql.Row) (model.LnChrgData, error) {
	var i model.LnChrgData
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.Acc,
		&i.ChrgCode,
		&i.RefAcc,
		&i.ChrAmnt,
	)
	return i, err
}

func scanRowsLnChrgData(rows *sql.Rows) ([]model.LnChrgData, error) {
	items := []model.LnChrgData{}
	for rows.Next() {
		var i model.LnChrgData
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.ChrgCode,
			&i.RefAcc,
			&i.ChrAmnt,
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

func (q *QueriesDump) GetLnChrgData(ctx context.Context, brCode string, acc string, chrgCode int64, refAcc string) (model.LnChrgData, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and Acc = $2 and ChrgCode = $3 and RefAcc = $4", getLnChrgData)
	row := q.db.QueryRowContext(ctx, sql, brCode, acc, chrgCode, refAcc)
	return scanRowLnChrgData(row)
}

type ListLnChrgDataParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListLnChrgData(ctx context.Context, lastModCtr int64) ([]model.LnChrgData, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getLnChrgData)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsLnChrgData(rows)
}

const updateLnChrgData = `-- name: UpdateLnChrgData :one
UPDATE esystemdump.LnChrgData SET 
ModCtr = $1,
ChrAmnt = $7
WHERE BrCode = $2 and Acc = $4 and ChrgCode = $5 and RefAcc = $6 and ModAction = $3
`

func (q *QueriesDump) UpdateLnChrgData(ctx context.Context, arg model.LnChrgData) error {
	_, err := q.db.ExecContext(ctx, updateLnChrgData,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.ChrgCode,
		arg.RefAcc,
		arg.ChrAmnt,
	)
	return err
}
