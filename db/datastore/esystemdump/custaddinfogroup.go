package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createCustAddInfoGroup = `-- name: CreateCustAddInfoGroup: one
INSERT INTO esystemdump.CustAddInfoGroup(
   ModCtr, BrCode, ModAction, InfoGroup, GroupTitle, Remarks, ReqOnEntry, ReqOnExit, Link2Loan, Link2Save )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT (brCode, infoGroup, ModAction)
DO UPDATE SET
  ModCtr =  EXCLUDED.ModCtr,
  GroupTitle =  EXCLUDED.GroupTitle,
  Remarks =  EXCLUDED.Remarks,
  ReqOnEntry =  EXCLUDED.ReqOnEntry,
  ReqOnExit =  EXCLUDED.ReqOnExit,
  Link2Loan =  EXCLUDED.Link2Loan,
  Link2Save =  EXCLUDED.Link2Save
`

func (q *QueriesDump) CreateCustAddInfoGroup(ctx context.Context, arg model.CustAddInfoGroup) error {
	_, err := q.db.ExecContext(ctx, createCustAddInfoGroup,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
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
DELETE FROM esystemdump.CustAddInfoGroup WHERE BrCode = $1 and InfoGroup = $2
`

func (q *QueriesDump) DeleteCustAddInfoGroup(ctx context.Context, brCode string, infoGroup int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustAddInfoGroup, brCode, infoGroup)
	return err
}

const getCustAddInfoGroup = `-- name: GetCustAddInfoGroup :one
SELECT
ModCtr, BrCode, ModAction, InfoGroup, GroupTitle, Remarks, ReqOnEntry, ReqOnExit, Link2Loan, Link2Save
FROM esystemdump.CustAddInfoGroup
`

func scanRowCustAddInfoGroup(row *sql.Row) (model.CustAddInfoGroup, error) {
	var i model.CustAddInfoGroup
	err := row.Scan(
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
	)
	return i, err
}

func scanRowsCustAddInfoGroup(rows *sql.Rows) ([]model.CustAddInfoGroup, error) {
	items := []model.CustAddInfoGroup{}
	for rows.Next() {
		var i model.CustAddInfoGroup
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

func (q *QueriesDump) GetCustAddInfoGroup(ctx context.Context, brCode string, infoGroup int64) (model.CustAddInfoGroup, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and InfoGroup = $2", getCustAddInfoGroup)
	row := q.db.QueryRowContext(ctx, sql, brCode, infoGroup)
	return scanRowCustAddInfoGroup(row)
}

type ListCustAddInfoGroupParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListCustAddInfoGroup(ctx context.Context, lastModCtr int64) ([]model.CustAddInfoGroup, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getCustAddInfoGroup)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCustAddInfoGroup(rows)
}

const updateCustAddInfoGroup = `-- name: UpdateCustAddInfoGroup :one
UPDATE esystemdump.CustAddInfoGroup SET 
  ModCtr = $1,
  GroupTitle = $5,
  Remarks = $6,
  ReqOnEntry = $7,
  ReqOnExit = $8,
  Link2Loan = $9,
  Link2Save = $10
WHERE BrCode = $2 and InfoGroup = $4 and  ModAction = $3
`

func (q *QueriesDump) UpdateCustAddInfoGroup(ctx context.Context, arg model.CustAddInfoGroup) error {
	_, err := q.db.ExecContext(ctx, updateCustAddInfoGroup,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
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
