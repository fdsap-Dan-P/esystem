package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createUnit = `-- name: CreateUnit: one
INSERT INTO esystemdump.Unit(
   ModCtr, BrCode, ModAction, UnitCode, Unit, AreaCode, FName, LName, MName, VatReg, UnitAddress )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT (brCode, unitCode, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	Unit =  EXCLUDED.Unit,
	AreaCode =  EXCLUDED.AreaCode,
	FName =  EXCLUDED.FName,
	LName =  EXCLUDED.LName,
	MName =  EXCLUDED.MName,
	VatReg =  EXCLUDED.VatReg,
	UnitAddress =  EXCLUDED.UnitAddress
`

func (q *QueriesDump) CreateUnit(ctx context.Context, arg model.Unit) error {
	_, err := q.db.ExecContext(ctx, createUnit,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.UnitCode,
		arg.Unit,
		arg.AreaCode,
		arg.FName,
		arg.LName,
		arg.MName,
		arg.VatReg,
		arg.UnitAddress,
	)
	return err
}

const deleteUnit = `-- name: DeleteUnit :exec
DELETE FROM esystemdump.Unit WHERE BrCode = $1 and UnitCode = $2
`

func (q *QueriesDump) DeleteUnit(ctx context.Context, brCode string, unitCode int64) error {
	_, err := q.db.ExecContext(ctx, deleteUnit, brCode, unitCode)
	return err
}

const getUnit = `-- name: GetUnit :one
SELECT
ModCtr, BrCode, ModAction, UnitCode, Unit, AreaCode, FName, LName, MName, VatReg, UnitAddress
FROM esystemdump.Unit
`

func scanRowUnit(row *sql.Row) (model.Unit, error) {
	var i model.Unit
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.UnitCode,
		&i.Unit,
		&i.AreaCode,
		&i.FName,
		&i.LName,
		&i.MName,
		&i.VatReg,
		&i.UnitAddress,
	)
	return i, err
}

func scanRowsUnit(rows *sql.Rows) ([]model.Unit, error) {
	items := []model.Unit{}
	for rows.Next() {
		var i model.Unit
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.UnitCode,
			&i.Unit,
			&i.AreaCode,
			&i.FName,
			&i.LName,
			&i.MName,
			&i.VatReg,
			&i.UnitAddress,
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

func (q *QueriesDump) GetUnit(ctx context.Context, brCode string, unitCode int64) (model.Unit, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and UnitCode = $2", getUnit)
	row := q.db.QueryRowContext(ctx, sql, brCode, unitCode)
	return scanRowUnit(row)
}

type ListUnitParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListUnit(ctx context.Context, lastModCtr int64) ([]model.Unit, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getUnit)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsUnit(rows)
}

const updateUnit = `-- name: UpdateUnit :one
UPDATE esystemdump.Unit SET 
	ModCtr = $1,
	Unit = $5,
	AreaCode = $6,
	FName = $7,
	LName = $8,
	MName = $9,
	VatReg = $10,
	UnitAddress = $11
WHERE BrCode = $2 and UnitCode = $4 and ModAction = $3
`

func (q *QueriesDump) UpdateUnit(ctx context.Context, arg model.Unit) error {
	_, err := q.db.ExecContext(ctx, updateUnit,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.UnitCode,
		arg.Unit,
		arg.AreaCode,
		arg.FName,
		arg.LName,
		arg.MName,
		arg.VatReg,
		arg.UnitAddress,
	)
	return err
}
