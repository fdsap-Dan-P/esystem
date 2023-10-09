package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const createReactivateWriteoff = `-- name: CreateReactivateWriteoff: one
INSERT INTO ReactivateWriteoff(
   CID, DeactivateBy, ReactivateBy, Stat, StatusDate)
VALUES ($1, $2, $3, $4, $5)

`

type ReactivateWriteoffRequest struct {
	ID           int64          `json:"id"`
	CID          int64          `json:"cid"`
	DeactivateBy sql.NullString `json:"deactivateBy"`
	ReactivateBy sql.NullString `json:"reactivateBy"`
	Status       int64          `json:"status"`
	StatusDate   time.Time      `json:"statusDate"`
}

func (q *QueriesLocal) CreateReactivateWriteoff(ctx context.Context, arg ReactivateWriteoffRequest) error {
	_, err := q.db.ExecContext(ctx, createReactivateWriteoff,
		arg.CID,
		arg.DeactivateBy,
		arg.ReactivateBy,
		arg.Status,
		arg.StatusDate,
	)
	return err
}

const deleteReactivateWriteoff = `-- name: DeleteReactivateWriteoff :exec
  DELETE FROM ReactivateWriteoff WHERE ID = $1
`

func (q *QueriesLocal) DeleteReactivateWriteoff(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteReactivateWriteoff, id)
	return err
}

const getReactivateWriteoff = `-- name: GetReactivateWriteoff :one
SELECT 
  m.ModCtr, m.ModAction, OrgParms.DefBranch_Code, ID, CID, DeactivateBy, ReactivateBy, Stat, StatusDate
FROM OrgParms,
 ReactivateWriteoff d
INNER JOIN
    (SELECT Max(ModCtr) ModCtr, ModAction, UniqueKeyInt1, min(CASE WHEN Uploaded = 1 THEN 1 ELSE 0 END) Uploaded
     FROM Modified
     WHERE TableName = 'ReactivateWriteoff'
     GROUP BY ModAction, UniqueKeyInt1 
     ) m on m.UniqueKeyInt1 = ID  
`

func scanRowReactivateWriteoff(row *sql.Row) (ReactivateWriteoff, error) {
	var i ReactivateWriteoff
	err := row.Scan(
		&i.ModCtr,
		&i.ModAction,
		&i.BrCode,
		&i.ID,
		&i.CID,
		&i.DeactivateBy,
		&i.ReactivateBy,
		&i.Status,
		&i.StatusDate,
	)
	return i, err
}

func scanRowsReactivateWriteoff(rows *sql.Rows) ([]ReactivateWriteoff, error) {
	items := []ReactivateWriteoff{}
	for rows.Next() {
		var i ReactivateWriteoff
		if err := rows.Scan(
			&i.ModCtr,
			&i.ModAction,
			&i.BrCode,
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

func (q *QueriesLocal) GetReactivateWriteoff(ctx context.Context, id int64) (ReactivateWriteoff, error) {
	sql := fmt.Sprintf("%s WHERE ID = $1", getReactivateWriteoff)
	row := q.db.QueryRowContext(ctx, sql, id)
	return scanRowReactivateWriteoff(row)
}

func (q *QueriesLocal) GetReactivateWriteoffbyCID(ctx context.Context, cid int64) (ReactivateWriteoff, error) {
	sql := fmt.Sprintf("%s WHERE ID in (SELECT Max(ID) ID FROM ReactivateWriteoff WHERE CID = $1)", getReactivateWriteoff)
	row := q.db.QueryRowContext(ctx, sql, cid)
	return scanRowReactivateWriteoff(row)
}

type ListReactivateWriteoffParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) ListReactivateWriteoff(ctx context.Context) ([]ReactivateWriteoff, error) {
	sql := fmt.Sprintf(`%v WHERE Uploaded = 0`, getReactivateWriteoff)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsReactivateWriteoff(rows)
}

const updateReactivateWriteoff = `-- name: UpdateReactivateWriteoff :one
UPDATE ReactivateWriteoff SET 
  CID = $2,
  DeactivateBy = $3,
  ReactivateBy = $4,
  Stat = $5,
  StatusDate = $6
WHERE ID = $1 
`

func (q *QueriesLocal) UpdateReactivateWriteoff(ctx context.Context, arg ReactivateWriteoffRequest) error {
	_, err := q.db.ExecContext(ctx, updateReactivateWriteoff,
		arg.ID,
		arg.CID,
		arg.DeactivateBy,
		arg.ReactivateBy,
		arg.Status,
		arg.StatusDate,
	)
	return err
}

func (q *QueriesLocal) ReactivateWriteoffCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
  SELECT 
	0 ModCtr, OrgParms.DefBranch_Code, CID, DeactivateBy, ReactivateBy, Stat, StatusDate
  FROM OrgParms,
   ReactivateWriteoff d
`, filenamePath)
}
