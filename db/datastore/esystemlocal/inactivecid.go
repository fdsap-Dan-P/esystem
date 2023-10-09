package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const createInActiveCID = `-- name: CreateInActiveCID: one
INSERT INTO InActiveCID(
   CID, InActive, Date_Start, Date_End, Posted_By, DeActivated_By)
VALUES ($1, $2, $3, $4, $5, $6)

`

type InActiveCIDRequest struct {
	CID           int64          `json:"cid"`
	InActive      bool           `json:"inActive"`
	DateStart     time.Time      `json:"dateStart"`
	DateEnd       sql.NullTime   `json:"dateEnd"`
	UserId        string         `json:"userId"`
	DeactivatedBy sql.NullString `json:"deactivatedBy"`
}

func (q *QueriesLocal) CreateInActiveCID(ctx context.Context, arg InActiveCIDRequest) error {
	_, err := q.db.ExecContext(ctx, createInActiveCID,
		arg.CID,
		arg.InActive,
		arg.DateStart,
		arg.DateEnd,
		arg.UserId,
		arg.DeactivatedBy,
	)
	return err
}

const deleteInActiveCID = `-- name: DeleteInActiveCID :exec
  DELETE FROM InActiveCID WHERE CID = $1 and Date_Start = $2
`

func (q *QueriesLocal) DeleteInActiveCID(ctx context.Context, cid int64, dateStart time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteInActiveCID, cid, dateStart)
	return err
}

const getInActiveCID = `-- name: GetInActiveCID :one
SELECT 
  m.ModCtr, m.ModAction, OrgParms.DefBranch_Code, CID, InActive, DateStart, DateEnd, UserId, DeactivatedBy
FROM OrgParms,
 (SELECT
    CID, max(CASE WHEN InActive = 1 THEN 1 ELSE 0 END) InActive, Date_Start DateStart, max(Date_End) DateEnd, max(Posted_By) UserId, max(Deactivated_By) DeactivatedBy
  FROM InActiveCID
  GROUP BY CID, Date_Start
  ) d
INNER JOIN
    (SELECT Max(ModCtr) ModCtr, ModAction, UniqueKeyInt1, UniqueKeyDate, min(CASE WHEN Uploaded = 1 THEN 1 ELSE 0 END) Uploaded
     FROM Modified
     WHERE TableName = 'InActiveCID'
     GROUP BY ModAction, UniqueKeyInt1, UniqueKeyDate 
     ) m on m.UniqueKeyInt1 = CID and m.UniqueKeyDate = d.DateStart  
`

func scanRowInActiveCID(row *sql.Row) (InActiveCID, error) {
	var i InActiveCID
	err := row.Scan(
		&i.ModCtr,
		&i.ModAction,
		&i.BrCode,
		&i.CID,
		&i.InActive,
		&i.DateStart,
		&i.DateEnd,
		&i.UserId,
		&i.DeactivatedBy,
	)
	return i, err
}

func scanRowsInActiveCID(rows *sql.Rows) ([]InActiveCID, error) {
	items := []InActiveCID{}
	for rows.Next() {
		var i InActiveCID
		if err := rows.Scan(
			&i.ModCtr,
			&i.ModAction,
			&i.BrCode,
			&i.CID,
			&i.InActive,
			&i.DateStart,
			&i.DateEnd,
			&i.UserId,
			&i.DeactivatedBy,
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

func (q *QueriesLocal) GetInActiveCID(ctx context.Context, cid int64, dateStart time.Time) (InActiveCID, error) {
	sql := fmt.Sprintf("%s WHERE CID = $1 and DateStart = $2", getInActiveCID)
	row := q.db.QueryRowContext(ctx, sql, cid, dateStart)
	return scanRowInActiveCID(row)
}

type ListInActiveCIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) ListInActiveCID(ctx context.Context) ([]InActiveCID, error) {
	sql := fmt.Sprintf(`%v WHERE Uploaded = 0`, getInActiveCID)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsInActiveCID(rows)
}

const updateInActiveCID = `-- name: UpdateInActiveCID :one
UPDATE InActiveCID SET 
  InActive = $2,
  Date_End = $4,
  Posted_By = $5,
  DeActivated_By = $6
WHERE CID = $1 and Date_Start = $3
`

func (q *QueriesLocal) UpdateInActiveCID(ctx context.Context, arg InActiveCIDRequest) error {
	_, err := q.db.ExecContext(ctx, updateInActiveCID,
		arg.CID,
		arg.InActive,
		arg.DateStart,
		arg.DateEnd,
		arg.UserId,
		arg.DeactivatedBy,
	)
	return err
}

func (q *QueriesLocal) InActiveCIDCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code, CID, InActive, DateStart, DateEnd, UserId, DeactivatedBy
FROM OrgParms,
 (SELECT
    CID, max(CASE WHEN InActive = 1 THEN 1 ELSE 0 END) InActive, Date_Start DateStart, 
	max(Date_End) DateEnd, max(Posted_By) UserId, max(Deactivated_By) DeactivatedBy
  FROM InActiveCID
  GROUP BY CID, Date_Start
) d  
`, filenamePath)
}
