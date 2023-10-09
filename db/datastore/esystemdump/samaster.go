package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createSaMaster = `-- name: CreateSaMaster: one
INSERT INTO esystemdump.SaMaster(
   ModCtr, BrCode, ModAction, Acc, CID, Type, Balance, DoLastTrn, DoStatus, Dopen, DoMaturity, Status )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
ON CONFLICT (brCode, acc, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	CID =  EXCLUDED.CID,
	Type =  EXCLUDED.Type,
	Balance =  EXCLUDED.Balance,
	DoLastTrn =  EXCLUDED.DoLastTrn,
	DoStatus =  EXCLUDED.DoStatus,
	Dopen =  EXCLUDED.Dopen,
	DoMaturity =  EXCLUDED.DoMaturity,
	Status =  EXCLUDED.Status
`

func (q *QueriesDump) CreateSaMaster(ctx context.Context, arg model.SaMaster) error {
	_, err := q.db.ExecContext(ctx, createSaMaster,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.CID,
		arg.Type,
		arg.Balance,
		arg.DoLastTrn,
		arg.DoStatus,
		arg.Dopen,
		arg.DoMaturity,
		arg.Status,
	)
	return err
}

const deleteSaMaster = `-- name: DeleteSaMaster :exec
DELETE FROM esystemdump.SaMaster WHERE BrCode = $1 and Acc = $2
`

func (q *QueriesDump) DeleteSaMaster(ctx context.Context, brCode string, acc string) error {
	_, err := q.db.ExecContext(ctx, deleteSaMaster, brCode, acc)
	return err
}

const getSaMaster = `-- name: GetSaMaster :one
SELECT
ModCtr, BrCode, ModAction, Acc, CID, Type, Balance, DoLastTrn, DoStatus, Dopen, DoMaturity, Status
FROM esystemdump.SaMaster
`

func scanRowSaMaster(row *sql.Row) (model.SaMaster, error) {
	var i model.SaMaster
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.Acc,
		&i.CID,
		&i.Type,
		&i.Balance,
		&i.DoLastTrn,
		&i.DoStatus,
		&i.Dopen,
		&i.DoMaturity,
		&i.Status,
	)
	return i, err
}

func scanRowsSaMaster(rows *sql.Rows) ([]model.SaMaster, error) {
	items := []model.SaMaster{}
	for rows.Next() {
		var i model.SaMaster
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.CID,
			&i.Type,
			&i.Balance,
			&i.DoLastTrn,
			&i.DoStatus,
			&i.Dopen,
			&i.DoMaturity,
			&i.Status,
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

func (q *QueriesDump) GetSaMaster(ctx context.Context, brCode string, acc string) (model.SaMaster, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and Acc = $2", getSaMaster)
	row := q.db.QueryRowContext(ctx, sql, brCode, acc)
	return scanRowSaMaster(row)
}

type ListSaMasterParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListSaMaster(ctx context.Context, lastModCtr int64) ([]model.SaMaster, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getSaMaster)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsSaMaster(rows)
}

const updateSaMaster = `-- name: UpdateSaMaster :one
UPDATE esystemdump.SaMaster SET 
	ModCtr = $1,
	CID = $5,
	Type = $6,
	Balance = $7,
	DoLastTrn = $8,
	DoStatus = $9,
	Dopen = $10,
	DoMaturity = $11,
	Status = $12
WHERE BrCode = $2 and Acc = $4 and ModAction = $3
`

func (q *QueriesDump) UpdateSaMaster(ctx context.Context, arg model.SaMaster) error {
	_, err := q.db.ExecContext(ctx, updateSaMaster,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.CID,
		arg.Type,
		arg.Balance,
		arg.DoLastTrn,
		arg.DoStatus,
		arg.Dopen,
		arg.DoMaturity,
		arg.Status,
	)
	return err
}
