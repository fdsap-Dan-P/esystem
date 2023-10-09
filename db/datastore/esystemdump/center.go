package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createCenter = `-- name: CreateCenter: one
INSERT INTO esystemdump.Center(
   ModCtr, BrCode, ModAction, CenterCode, CenterName, CenterAddress, MeetingDay, Unit, DateEstablished, AOID)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT (brCode, centerCode, ModAction)
DO UPDATE SET
  ModCtr =  EXCLUDED.ModCtr,
  CenterName =  EXCLUDED.CenterName,
  CenterAddress =  EXCLUDED.CenterAddress,
  MeetingDay =  EXCLUDED.MeetingDay,
  Unit =  EXCLUDED.Unit,
  DateEstablished =  EXCLUDED.DateEstablished,
  AOID =  EXCLUDED.AOID
`

func (q *QueriesDump) CreateCenter(ctx context.Context, arg model.Center) error {
	_, err := q.db.ExecContext(ctx, createCenter,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CenterCode,
		arg.CenterName,
		arg.CenterAddress,
		arg.MeetingDay,
		arg.Unit,
		arg.DateEstablished,
		arg.AOID,
	)
	return err
}

const deleteCenter = `-- name: DeleteCenter :exec
DELETE FROM esystemdump.Center WHERE BrCode = $1 and CenterCode = $2
`

func (q *QueriesDump) DeleteCenter(ctx context.Context, brCode string, centerCode string) error {
	_, err := q.db.ExecContext(ctx, deleteCenter, brCode, centerCode)
	return err
}

const getCenter = `-- name: GetCenter :one
SELECT
  ModCtr, BrCode, ModAction, CenterCode, CenterName, CenterAddress, 
  MeetingDay, Unit, DateEstablished, AOID
FROM esystemdump.Center
`

func scanRowCenter(row *sql.Row) (model.Center, error) {
	var i model.Center
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.CenterCode,
		&i.CenterName,
		&i.CenterAddress,
		&i.MeetingDay,
		&i.Unit,
		&i.DateEstablished,
		&i.AOID,
	)
	return i, err
}

func scanRowsCenter(rows *sql.Rows) ([]model.Center, error) {
	items := []model.Center{}
	for rows.Next() {
		var i model.Center
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.CenterCode,
			&i.CenterName,
			&i.CenterAddress,
			&i.MeetingDay,
			&i.Unit,
			&i.DateEstablished,
			&i.AOID,
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

func (q *QueriesDump) GetCenter(ctx context.Context, brCode string, centerCode string) (model.Center, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and CenterCode = $2", getCenter)
	row := q.db.QueryRowContext(ctx, sql, brCode, centerCode)
	return scanRowCenter(row)
}

type ListCenterParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListCenter(ctx context.Context, lastModCtr int64) ([]model.Center, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getCenter)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCenter(rows)
}

const updateCenter = `-- name: UpdateCenter :one
UPDATE esystemdump.Center SET 
	ModCtr = $1,
	CenterName = $5,
	CenterAddress = $6,
	MeetingDay = $7,
	Unit = $8,
	DateEstablished = $9,
	AOID = $10
WHERE BrCode = $2 and CenterCode = $4 and ModAction = $3
`

func (q *QueriesDump) UpdateCenter(ctx context.Context, arg model.Center) error {
	_, err := q.db.ExecContext(ctx, updateCenter,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CenterCode,
		arg.CenterName,
		arg.CenterAddress,
		arg.MeetingDay,
		arg.Unit,
		arg.DateEstablished,
		arg.AOID,
	)
	return err
}
