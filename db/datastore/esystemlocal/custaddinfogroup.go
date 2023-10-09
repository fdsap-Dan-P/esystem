package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const createCustAddInfoGroup = `-- name: CreateCustAddInfoGroup: one
INSERT INTO CustAddInfoGroup (
	InfoGroup, GroupTitle, Remarks, ReqOnEntry, ReqOnExit, Link2Loan, Link2Save
) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
`

type CustAddInfoGroupRequest struct {
	InfoGroup  int64          `json:"infoGroup"`
	GroupTitle sql.NullString `json:"groupTitle"`
	Remarks    sql.NullString `json:"remarks"`
	ReqOnEntry sql.NullBool   `json:"reqOnEntry"`
	ReqOnExit  sql.NullBool   `json:"reqOnExit"`
	Link2Loan  sql.NullInt64  `json:"link2Loan"`
	Link2Save  sql.NullInt64  `json:"link2Save"`
}

func (q *QueriesLocal) CreateCustAddInfoGroup(ctx context.Context, arg CustAddInfoGroupRequest) error {
	_, err := q.db.ExecContext(ctx, createCustAddInfoGroup,
		arg.InfoGroup,
		arg.GroupTitle,
		arg.Remarks,
		arg.ReqOnEntry,
		arg.ReqOnExit,
		arg.Link2Loan,
		arg.Link2Save,
	)
	return err
}

const deleteCustAddInfoGroup = `-- name: DeleteCustAddInfoGroup :exec
DELETE FROM CustAddInfoGroup WHERE InfoGroup = $1
`

func (q *QueriesLocal) DeleteCustAddInfoGroup(ctx context.Context, infoGroup int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustAddInfoGroup, infoGroup)
	return err
}

type CustAddInfoGroupInfo struct {
	ModCtr     int64          `json:"modCtr"`
	BrCode     string         `json:"brCode"`
	ModAction  string         `json:"modAction"`
	InfoGroup  int64          `json:"infoGroup"`
	GroupTitle sql.NullString `json:"groupTitle"`
	Remarks    sql.NullString `json:"remarks"`
	ReqOnEntry sql.NullBool   `json:"reqOnEntry"`
	ReqOnExit  sql.NullBool   `json:"reqOnExit"`
	Link2Loan  sql.NullInt64  `json:"link2Loan"`
	Link2Save  sql.NullInt64  `json:"link2Save"`
}

// -- name: GetCustAddInfoGroup :one
const getCustAddInfoGroup = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, InfoGroup, GroupTitle, Remarks, ReqOnEntry, ReqOnExit, Link2Loan, Link2Save
FROM OrgParms, CustAddInfoGroup d
INNER JOIN Modified m on m.UniqueKeyInt1 = d.InfoGroup 
`

func scanRowCustAddInfoGroup(row *sql.Row) (CustAddInfoGroupInfo, error) {
	var i CustAddInfoGroupInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.InfoGroup,
		&i.GroupTitle,
		&i.Remarks,
		&i.ReqOnEntry,
		&i.ReqOnExit,
		&i.Link2Loan,
		&i.Link2Save,
	)
	return i, err
}

func scanRowsCustAddInfoGroup(rows *sql.Rows) ([]CustAddInfoGroupInfo, error) {
	items := []CustAddInfoGroupInfo{}
	for rows.Next() {
		var i CustAddInfoGroupInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.InfoGroup,
			&i.GroupTitle,
			&i.Remarks,
			&i.ReqOnEntry,
			&i.ReqOnExit,
			&i.Link2Loan,
			&i.Link2Save,
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

func (q *QueriesLocal) GetCustAddInfoGroup(ctx context.Context, infoGroup int64) (CustAddInfoGroupInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'CustAddInfoGroup' AND Uploaded = 0 and InfoGroup = $1", getCustAddInfoGroup)
	row := q.db.QueryRowContext(ctx, sql, infoGroup)
	return scanRowCustAddInfoGroup(row)
}

type ListCustAddInfoGroupParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) CustAddInfoGroupCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, InfoGroup, 
  GroupTitle, Remarks, ReqOnEntry, ReqOnExit, Link2Loan, Link2Save
FROM OrgParms, CustAddInfoGroup d
`, filenamePath)
}

func (q *QueriesLocal) ListCustAddInfoGroup(ctx context.Context) ([]CustAddInfoGroupInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE  m.TableName = 'CustAddInfoGroup' AND Uploaded = 0`,
		getCustAddInfoGroup)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCustAddInfoGroup(rows)
}

// -- name: UpdateCustAddInfoGroup :one
const updateCustAddInfoGroup = `
UPDATE CustAddInfoGroup SET 
	GroupTitle = $2,
	Remarks = $3,
	ReqOnEntry = $4,
	ReqOnExit = $5,
	Link2Loan = $6,
	Link2Save = $7
WHERE InfoGroup = $1`

func (q *QueriesLocal) UpdateCustAddInfoGroup(ctx context.Context, arg CustAddInfoGroupRequest) error {
	_, err := q.db.ExecContext(ctx, updateCustAddInfoGroup,
		arg.InfoGroup,
		arg.GroupTitle,
		arg.Remarks,
		arg.ReqOnEntry,
		arg.ReqOnExit,
		arg.Link2Loan,
		arg.Link2Save,
	)
	return err
}
