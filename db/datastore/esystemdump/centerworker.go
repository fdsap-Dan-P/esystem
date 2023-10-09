package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createCenterWorker = `-- name: CreateCenterWorker: one
INSERT INTO esystemdump.CenterWorker(
   ModCtr, BrCode, ModAction, AOID, Lname, FName, Mname, PhoneNumber)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT (brCode, aOID, ModAction)
DO UPDATE SET
  ModCtr =  EXCLUDED.ModCtr,
  Lname =  EXCLUDED.Lname,
  FName =  EXCLUDED.FName,
  Mname =  EXCLUDED.Mname,
  PhoneNumber =  EXCLUDED.PhoneNumber
`

func (q *QueriesDump) CreateCenterWorker(ctx context.Context, arg model.CenterWorker) error {
	_, err := q.db.ExecContext(ctx, createCenterWorker,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.AOID,
		arg.Lname,
		arg.FName,
		arg.Mname,
		arg.PhoneNumber,
	)
	return err
}

const deleteCenterWorker = `-- name: DeleteCenterWorker :exec
DELETE FROM esystemdump.CenterWorker WHERE BrCode = $1 and AOID = $2
`

func (q *QueriesDump) DeleteCenterWorker(ctx context.Context, brCode string, aOID int64) error {
	_, err := q.db.ExecContext(ctx, deleteCenterWorker, brCode, aOID)
	return err
}

const getCenterWorker = `-- name: GetCenterWorker :one
SELECT
ModCtr, BrCode, ModAction, AOID, Lname, FName, Mname, PhoneNumber
FROM esystemdump.CenterWorker
`

func scanRowCenterWorker(row *sql.Row) (model.CenterWorker, error) {
	var i model.CenterWorker
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.AOID,
		&i.Lname,
		&i.FName,
		&i.Mname,
		&i.PhoneNumber,
	)
	return i, err
}

func scanRowsCenterWorker(rows *sql.Rows) ([]model.CenterWorker, error) {
	items := []model.CenterWorker{}
	for rows.Next() {
		var i model.CenterWorker
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.AOID,
			&i.Lname,
			&i.FName,
			&i.Mname,
			&i.PhoneNumber,
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

func (q *QueriesDump) GetCenterWorker(ctx context.Context, brCode string, aOID int64) (model.CenterWorker, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and AOID = $2", getCenterWorker)
	row := q.db.QueryRowContext(ctx, sql, brCode, aOID)
	return scanRowCenterWorker(row)
}

type ListCenterWorkerParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListCenterWorker(ctx context.Context, lastModCtr int64) ([]model.CenterWorker, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getCenterWorker)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCenterWorker(rows)
}

const updateCenterWorker = `-- name: UpdateCenterWorker :one
UPDATE esystemdump.CenterWorker SET 
  ModCtr = $1,
  Lname = $5,
  FName = $6,
  Mname = $7,
  PhoneNumber = $8
WHERE BrCode = $2 and AOID = $4 and ModAction = $3
`

func (q *QueriesDump) UpdateCenterWorker(ctx context.Context, arg model.CenterWorker) error {
	_, err := q.db.ExecContext(ctx, updateCenterWorker,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.AOID,
		arg.Lname,
		arg.FName,
		arg.Mname,
		arg.PhoneNumber,
	)
	return err
}
