package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const createCenterWorker = `-- name: CreateCenterWorker: one
IF EXISTS (SELECT AreaCode FROM Area WHERE AreaCode = $1)
BEGIN
  UPDATE Center_Worker SET 
	CenterW_LName = $2,
	CenterW_FName = $3,
	CenterW_MName = $4,
	PhoneNumber = $5
  WHERE CenterW_ID = $1
END ELSE BEGIN
  INSERT INTO Center_Worker (
	CenterW_ID, CenterW_LName, CenterW_FName, CenterW_MName, PhoneNumber
  ) 
  VALUES ($1, $2, $3, $4, $5) 
END
`

type CenterWorkerRequest struct {
	AOID        int64          `json:"aoID"`
	Lname       sql.NullString `json:"lname"`
	FName       sql.NullString `json:"fName"`
	Mname       sql.NullString `json:"mname"`
	PhoneNumber sql.NullString `json:"phoneNumber"`
}

func (q *QueriesLocal) CreateCenterWorker(ctx context.Context, arg CenterWorkerRequest) error {
	_, err := q.db.ExecContext(ctx, createCenterWorker,
		arg.AOID,
		arg.Lname,
		arg.FName,
		arg.Mname,
		arg.PhoneNumber,
	)
	return err
}

const deleteCenterWorker = `-- name: DeleteCenterWorker :exec
DELETE FROM Center_Worker WHERE CenterW_ID = $1
`

func (q *QueriesLocal) DeleteCenterWorker(ctx context.Context, aoid int64) error {
	_, err := q.db.ExecContext(ctx, deleteCenterWorker, aoid)
	return err
}

type CenterWorkerInfo struct {
	ModCtr      int64          `json:"modCtr"`
	BrCode      string         `json:"brCode"`
	ModAction   string         `json:"modAction"`
	AOID        int64          `json:"aoID"`
	Lname       sql.NullString `json:"lname"`
	FName       sql.NullString `json:"fName"`
	Mname       sql.NullString `json:"mname"`
	PhoneNumber sql.NullString `json:"phoneNumber"`
}

// -- name: GetCenterWorker :one
const getCenterWorker = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, CenterW_ID, CenterW_LName, CenterW_FName, CenterW_MName, PhoneNumber
FROM OrgParms, Center_Worker d
INNER JOIN Modified m on m.UniqueKeyInt1 = d.CenterW_ID 
`

func scanRowCenterWorker(row *sql.Row) (CenterWorkerInfo, error) {
	var i CenterWorkerInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.AOID,
		&i.Lname,
		&i.FName,
		&i.Mname,
		&i.PhoneNumber,
	)
	return i, err
}

func scanRowsCenterWorker(rows *sql.Rows) ([]CenterWorkerInfo, error) {
	items := []CenterWorkerInfo{}
	for rows.Next() {
		var i CenterWorkerInfo
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

func (q *QueriesLocal) GetCenterWorker(ctx context.Context, aoid int64) (CenterWorkerInfo, error) {
	sql := fmt.Sprintf("%s WHERE  m.TableName = 'Center_Worker' AND Uploaded = 0 and CenterW_ID = $1", getCenterWorker)
	row := q.db.QueryRowContext(ctx, sql, aoid)
	return scanRowCenterWorker(row)
}

type ListCenterWorkerParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) CenterWorkerCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, CenterW_ID, CenterW_LName, 
  CenterW_FName, CenterW_MName, PhoneNumber
FROM OrgParms, Center_Worker d
`, filenamePath)
}

func (q *QueriesLocal) ListCenterWorker(ctx context.Context) ([]CenterWorkerInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'Center_Worker' AND Uploaded = 0`,
		getCenterWorker)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCenterWorker(rows)
}

// -- name: UpdateCenterWorker :one
const updateCenterWorker = `
UPDATE Center_Worker SET 
	CenterW_LName = $2,
	CenterW_FName = $3,
	CenterW_MName = $4,
	PhoneNumber = $5
WHERE CenterW_ID = $1`

func (q *QueriesLocal) UpdateCenterWorker(ctx context.Context, arg CenterWorkerRequest) error {
	_, err := q.db.ExecContext(ctx, updateCenterWorker,
		arg.AOID,
		arg.Lname,
		arg.FName,
		arg.Mname,
		arg.PhoneNumber,
	)
	return err
}
