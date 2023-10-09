package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const createJnlHeaders = `-- name: CreateJnlHeaders: one
INSERT INTO JnlHeaders (
  Jnlh_Tran, Jnlh_Tran_Dt, Jnlh_Explanation, Jnlh_Post_By, Jnlh_Code, JNLH_PDTS_AY, JNLH_PDTS_AP
) 
VALUES ($1, $2, $3, $4, $5, 1999, 1) 
`

type JnlHeadersRequest struct {
	Trn         string         `json:"trn"`
	TrnDate     time.Time      `json:"trnDate"`
	Particulars string         `json:"particulars"`
	UserName    sql.NullString `json:"userName"`
	Code        int64          `json:"code"`
}

func (q *QueriesLocal) CreateJnlHeaders(ctx context.Context, arg JnlHeadersRequest) error {
	_, err := q.db.ExecContext(ctx, createJnlHeaders,
		arg.Trn,
		arg.TrnDate,
		arg.Particulars,
		arg.UserName,
		arg.Code,
	)
	return err
}

const deleteJnlHeaders = `-- name: DeleteJnlHeaders :exec
DELETE FROM JnlHeaders WHERE Jnlh_Tran = $1
`

func (q *QueriesLocal) DeleteJnlHeaders(ctx context.Context, Jnlh_Tran string) error {
	_, err := q.db.ExecContext(ctx, deleteJnlHeaders, Jnlh_Tran)
	return err
}

type JnlHeadersInfo struct {
	ModCtr      int64          `json:"modCtr"`
	BrCode      string         `json:"brCode"`
	ModAction   string         `json:"modAction"`
	Trn         string         `json:"trn"`
	TrnDate     time.Time      `json:"trnDate"`
	Particulars string         `json:"particulars"`
	UserName    sql.NullString `json:"userName"`
	Code        int64          `json:"code"`
}

// -- name: GetJnlHeaders :one
const getJnlHeaders = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, Jnlh_Tran, Jnlh_Tran_Dt, Jnlh_Explanation, Jnlh_Post_By, Jnlh_Code
FROM OrgParms, JnlHeaders d
INNER JOIN Modified m on m.UniqueKeyString1 = d.Jnlh_Tran
`

func scanRowJnlHeaders(row *sql.Row) (JnlHeadersInfo, error) {
	var i JnlHeadersInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.Trn,
		&i.TrnDate,
		&i.Particulars,
		&i.UserName,
		&i.Code,
	)
	return i, err
}

func scanRowsJnlHeaders(rows *sql.Rows) ([]JnlHeadersInfo, error) {
	items := []JnlHeadersInfo{}
	for rows.Next() {
		var i JnlHeadersInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Trn,
			&i.TrnDate,
			&i.Particulars,
			&i.UserName,
			&i.Code,
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

func (q *QueriesLocal) GetJnlHeaders(ctx context.Context, Jnlh_Tran string) (JnlHeadersInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'JnlHeaders' AND Uploaded = 0 and Jnlh_Tran = $1", getJnlHeaders)
	row := q.db.QueryRowContext(ctx, sql, Jnlh_Tran)
	return scanRowJnlHeaders(row)
}

type ListJnlHeadersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) JnlHeadersCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, Jnlh_Tran, Jnlh_Tran_Dt, 
  Jnlh_Explanation, Jnlh_Post_By, Jnlh_Code
FROM OrgParms, JnlHeaders d
`, filenamePath)
}

func (q *QueriesLocal) ListJnlHeaders(ctx context.Context) ([]JnlHeadersInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'JnlHeaders' AND Uploaded = 0`,
		getJnlHeaders)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsJnlHeaders(rows)
}

// -- name: UpdateJnlHeaders :one
const updateJnlHeaders = `
UPDATE JnlHeaders SET 
	Jnlh_Tran_Dt = $2,
	Jnlh_Explanation = $3,
	Jnlh_Post_By = $4,
	Jnlh_Code = $5
WHERE Jnlh_Tran = $1`

func (q *QueriesLocal) UpdateJnlHeaders(ctx context.Context, arg JnlHeadersRequest) error {
	_, err := q.db.ExecContext(ctx, updateJnlHeaders,
		arg.Trn,
		arg.TrnDate,
		arg.Particulars,
		arg.UserName,
		arg.Code,
	)
	return err
}
