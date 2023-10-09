package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const createCustAddInfoGroupNeed = `-- name: CreateCustAddInfoGroupNeed: one
INSERT INTO CustAddInfoGroupNeed (
	InfoGroup, InfoCode, InfoProcess
) 
VALUES ($1, $2, $3) 
`

type CustAddInfoGroupNeedRequest struct {
	InfoGroup   int64          `json:"infoGroup"`
	InfoCode    int64          `json:"infoCode"`
	InfoProcess sql.NullString `json:"infoProcess"`
}

func (q *QueriesLocal) CreateCustAddInfoGroupNeed(ctx context.Context, arg CustAddInfoGroupNeedRequest) error {
	_, err := q.db.ExecContext(ctx, createCustAddInfoGroupNeed,
		arg.InfoGroup,
		arg.InfoCode,
		arg.InfoProcess,
	)
	return err
}

const deleteCustAddInfoGroupNeed = `-- name: DeleteCustAddInfoGroupNeed :exec
DELETE FROM CustAddInfoGroupNeed WHERE InfoGroup = $1 and InfoCode = $2
`

func (q *QueriesLocal) DeleteCustAddInfoGroupNeed(ctx context.Context, infoGroup int64, infoCode int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustAddInfoGroupNeed, infoGroup, infoCode)
	return err
}

type CustAddInfoGroupNeedInfo struct {
	ModCtr    int64  `json:"modCtr"`
	BrCode    string `json:"brCode"`
	ModAction string `json:"modAction"`

	InfoGroup   int64          `json:"infoGroup"`
	InfoCode    int64          `json:"infoCode"`
	InfoProcess sql.NullString `json:"infoProcess"`
}

// -- name: GetCustAddInfoGroupNeed :one
const getCustAddInfoGroupNeed = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, InfoGroup, InfoCode, InfoProcess
FROM OrgParms, CustAddInfoGroupNeed d
INNER JOIN Modified m on m.UniqueKeyInt1 = d.InfoGroup and m.UniqueKeyInt2 = d.InfoCode 
`

func scanRowCustAddInfoGroupNeed(row *sql.Row) (CustAddInfoGroupNeedInfo, error) {
	var i CustAddInfoGroupNeedInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.InfoGroup,
		&i.InfoCode,
		&i.InfoProcess,
	)
	return i, err
}

func scanRowsCustAddInfoGroupNeed(rows *sql.Rows) ([]CustAddInfoGroupNeedInfo, error) {
	items := []CustAddInfoGroupNeedInfo{}
	for rows.Next() {
		var i CustAddInfoGroupNeedInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.InfoGroup,
			&i.InfoCode,
			&i.InfoProcess,
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

func (q *QueriesLocal) GetCustAddInfoGroupNeed(ctx context.Context, infoGroup int64, infoCode int64) (CustAddInfoGroupNeedInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'CustAddInfoGroupNeed' AND Uploaded = 0 and InfoGroup = $1 and InfoCode = $2", getCustAddInfoGroupNeed)
	row := q.db.QueryRowContext(ctx, sql, infoGroup, infoCode)
	return scanRowCustAddInfoGroupNeed(row)
}

type ListCustAddInfoGroupNeedParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) CustAddInfoGroupNeedCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
	0 ModCtr, OrgParms.DefBranch_Code BrCode, InfoGroup, InfoCode, InfoProcess
FROM OrgParms, CustAddInfoGroupNeed d
`, filenamePath)
}

func (q *QueriesLocal) ListCustAddInfoGroupNeed(ctx context.Context) ([]CustAddInfoGroupNeedInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'CustAddInfoGroupNeed' AND Uploaded = 0`,
		getCustAddInfoGroupNeed)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCustAddInfoGroupNeed(rows)
}

// -- name: UpdateCustAddInfoGroupNeed :one
const updateCustAddInfoGroupNeed = `
UPDATE CustAddInfoGroupNeed SET 
	InfoProcess = $3
WHERE InfoGroup = $1 and InfoCode = $2`

func (q *QueriesLocal) UpdateCustAddInfoGroupNeed(ctx context.Context, arg CustAddInfoGroupNeedRequest) error {
	_, err := q.db.ExecContext(ctx, updateCustAddInfoGroupNeed,
		arg.InfoGroup,
		arg.InfoCode,
		arg.InfoProcess,
	)
	return err
}
