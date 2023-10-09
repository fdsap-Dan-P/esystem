package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	model "simplebank/db/datastore/esystemlocal"
)

const createInActiveCID = `-- name: CreateInActiveCID: one
INSERT INTO esystemdump.InActiveCID(
   ModCtr, BrCode, ModAction, CID, InActive, DateStart, DateEnd, UserId)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (BrCode, CID, DateStart, ModAction)
DO UPDATE SET
ModCtr = EXCLUDED.ModCtr,
BrCode = EXCLUDED.BrCode,
InActive = EXCLUDED.InActive,
DateEnd = EXCLUDED.DateEnd,
UserId = EXCLUDED.UserId
`

func (q *QueriesDump) CreateInActiveCID(ctx context.Context, arg model.InActiveCID) error {
	_, err := q.db.ExecContext(ctx, createInActiveCID,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CID,
		arg.InActive,
		arg.DateStart,
		arg.DateEnd,
		arg.UserId,
	)
	return err
}

const deleteInActiveCID = `-- name: DeleteInActiveCID :exec
DELETE FROM esystemdump.InActiveCID WHERE BrCode = $1 and CID = $2 and DateStart = $3
`

func (q *QueriesDump) DeleteInActiveCID(ctx context.Context, brCode string, cid int64, dateStart time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteInActiveCID, brCode, cid, dateStart)
	return err
}

const getInActiveCID = `-- name: GetInActiveCID :one
SELECT
  ModCtr, BrCode, ModAction, CID, InActive, DateStart, DateEnd, UserId
FROM esystemdump.InActiveCID
`

func scanRowInActiveCID(row *sql.Row) (model.InActiveCID, error) {
	var i model.InActiveCID
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.CID,
		&i.InActive,
		&i.DateStart,
		&i.DateEnd,
		&i.UserId,
	)
	return i, err
}

func scanRowsInActiveCID(rows *sql.Rows) ([]model.InActiveCID, error) {
	items := []model.InActiveCID{}
	for rows.Next() {
		var i model.InActiveCID
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.CID,
			&i.InActive,
			&i.DateStart,
			&i.DateEnd,
			&i.UserId,
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

func (q *QueriesDump) GetInActiveCID(ctx context.Context, brCode string, cid int64, dateStart time.Time) (model.InActiveCID, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and CID = $2 and DateStart = $3", getInActiveCID)
	row := q.db.QueryRowContext(ctx, sql, brCode, cid, dateStart)
	return scanRowInActiveCID(row)
}

type ListInActiveCIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListInActiveCID(ctx context.Context, lastModCtr int64) ([]model.InActiveCID, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getInActiveCID)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsInActiveCID(rows)
}

const updateInActiveCID = `-- name: UpdateInActiveCID :one
UPDATE esystemdump.InActiveCID SET 
  ModCtr = $1,
  InActive = $5,
  DateEnd = $7,
  UserId = $8
WHERE BrCode = $2 and CID = $4 and DateStart = $6 and ModAction = $3
`

func (q *QueriesDump) UpdateInActiveCID(ctx context.Context, arg model.InActiveCID) error {
	_, err := q.db.ExecContext(ctx, updateInActiveCID,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.CID,
		arg.InActive,
		arg.DateStart,
		arg.DateEnd,
		arg.UserId,
	)
	return err
}
