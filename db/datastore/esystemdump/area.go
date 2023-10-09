package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createArea = `-- name: CreateArea: one
INSERT INTO esystemdump.Area(
   ModCtr, BrCode, ModAction, AreaCode, Area )
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (brCode, areaCode, ModAction)
DO UPDATE SET
  ModCtr =  EXCLUDED.ModCtr,
  Area =  EXCLUDED.Area
`

func (q *QueriesDump) CreateArea(ctx context.Context, arg model.Area) error {
	_, err := q.db.ExecContext(ctx, createArea,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.AreaCode,
		arg.Area,
	)
	return err
}

const deleteArea = `-- name: DeleteArea :exec
DELETE FROM esystemdump.Area WHERE BrCode = $1 and AreaCode = $2
`

func (q *QueriesDump) DeleteArea(ctx context.Context, brCode string, areaCode int64) error {
	_, err := q.db.ExecContext(ctx, deleteArea, brCode, areaCode)
	return err
}

const getArea = `-- name: GetArea :one
SELECT
ModCtr, BrCode, ModAction, AreaCode, Area
FROM esystemdump.Area
`

func scanRowArea(row *sql.Row) (model.Area, error) {
	var i model.Area
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.AreaCode,
		&i.Area,
	)
	return i, err
}

func scanRowsArea(rows *sql.Rows) ([]model.Area, error) {
	items := []model.Area{}
	for rows.Next() {
		var i model.Area
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.AreaCode,
			&i.Area,
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

func (q *QueriesDump) GetArea(ctx context.Context, brCode string, areaCode int64) (model.Area, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and AreaCode = $2", getArea)
	row := q.db.QueryRowContext(ctx, sql, brCode, areaCode)
	return scanRowArea(row)
}

type ListAreaParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListArea(ctx context.Context, lastModCtr int64) ([]model.Area, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getArea)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsArea(rows)
}

const updateArea = `-- name: UpdateArea :one
UPDATE esystemdump.Area SET 
  ModCtr = $1,
  Area = $5
WHERE BrCode = $2 and AreaCode = $4 and ModAction = $3
`

func (q *QueriesDump) UpdateArea(ctx context.Context, arg model.Area) error {
	_, err := q.db.ExecContext(ctx, updateArea,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.AreaCode,
		arg.Area,
	)
	return err
}
