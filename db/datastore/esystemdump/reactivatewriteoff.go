package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	model "simplebank/db/datastore/esystemlocal"
)

const createReactivateWriteoff = `-- name: CreateReactivateWriteoff: one
INSERT INTO esystemdump.ReactivateWriteoff(
	ModCtr, BrCode, ModAction, ID, CID, DeactivateBy, ReactivateBy, Status, StatusDate)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT(BrCode, ID, ModAction) DO UPDATE SET
  ModCtr = EXCLUDED.ModCtr, 
  CID = EXCLUDED.CID, 
  DeactivateBy = EXCLUDED.DeactivateBy, 
  ReactivateBy = EXCLUDED.ReactivateBy, 
  Status = EXCLUDED.Status, 
  StatusDate = EXCLUDED.StatusDate
`

func (q *QueriesDump) CreateReactivateWriteoff(ctx context.Context, arg model.ReactivateWriteoff) error {
	_, err := q.db.ExecContext(ctx, createReactivateWriteoff,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.ID,
		arg.CID,
		arg.DeactivateBy,
		arg.ReactivateBy,
		arg.Status,
		arg.StatusDate,
	)
	return err
}

const deleteReactivateWriteoff = `-- name: DeleteReactivateWriteoff :exec
  DELETE FROM esystemdump.ReactivateWriteoff WHERE BrCode = $1 and ID = $2
`

func (q *QueriesDump) DeleteReactivateWriteoff(ctx context.Context, brCode string, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteReactivateWriteoff, brCode, id)
	return err
}

const getReactivateWriteoff = `-- name: GetReactivateWriteoff :one
SELECT 
  ModCtr, BrCode, ModAction, ID, CID, DeactivateBy, ReactivateBy, Status, StatusDate
FROM esystemdump.ReactivateWriteoff d
`

func scanRowReactivateWriteoff(row *sql.Row) (model.ReactivateWriteoff, error) {
	var i model.ReactivateWriteoff
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.ID,
		&i.CID,
		&i.DeactivateBy,
		&i.ReactivateBy,
		&i.Status,
		&i.StatusDate,
	)
	return i, err
}

func scanRowsReactivateWriteoff(rows *sql.Rows) ([]model.ReactivateWriteoff, error) {
	items := []model.ReactivateWriteoff{}
	for rows.Next() {
		var i model.ReactivateWriteoff
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.ID,
			&i.CID,
			&i.DeactivateBy,
			&i.ReactivateBy,
			&i.Status,
			&i.StatusDate,
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

func (q *QueriesDump) GetReactivateWriteoff(ctx context.Context, brCode string, id int64) (model.ReactivateWriteoff, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and ID = $2", getReactivateWriteoff)
	row := q.db.QueryRowContext(ctx, sql, id)
	return scanRowReactivateWriteoff(row)
}

func (q *QueriesDump) GetReactivateWriteoffbyCID(ctx context.Context, brCode string, cid int64) (model.ReactivateWriteoff, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and  CID = $2 ORDER BY ID DESC LIMIT 1", getReactivateWriteoff)
	row := q.db.QueryRowContext(ctx, sql, brCode, cid)
	return scanRowReactivateWriteoff(row)
}

type ListReactivateWriteoffParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListReactivateWriteoff(ctx context.Context, lastModCtr int64) ([]model.ReactivateWriteoff, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getReactivateWriteoff)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsReactivateWriteoff(rows)
}

const updateReactivateWriteoff = `-- name: UpdateReactivateWriteoff :one
UPDATE esystemdump.ReactivateWriteoff SET 
  ModCtr = $1,
  CID = $5,
  DeactivateBy = $6,
  ReactivateBy = $7,
  Status = $8,
  StatusDate = $9
WHERE BrCode = $2 and ID = $4 and ModAction = $3
`

func (q *QueriesDump) UpdateReactivateWriteoff(ctx context.Context, arg model.ReactivateWriteoff) error {
	_, err := q.db.ExecContext(ctx, updateReactivateWriteoff,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.ID,
		arg.CID,
		arg.DeactivateBy,
		arg.ReactivateBy,
		arg.Status,
		arg.StatusDate,
	)
	return err
}
