package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const createReferencesDetails = `-- name: CreateReferencesDetails: one
IF EXISTS (SELECT ID FROM ReferencesDetails WHERE ID = $1)
BEGIN
  UPDATE ReferencesDetails SET
	RefID = $2, 
	PurposeDescription = $3, 
	ParentID = $4, 
	CodeID = $5, 
	Stat = $6
  WHERE ID = $1

END ELSE
BEGIN
  INSERT INTO ReferencesDetails (
	ID, RefID, PurposeDescription, ParentID, CodeID, Stat
) 
VALUES ($1, $2, $3, $4, $5, $6) 
END
`

type ReferencesDetailsRequest struct {
	ID                 int64          `json:"id"`
	RefID              int64          `json:"refID"`
	PurposeDescription sql.NullString `json:"purposeDescription"`
	ParentID           sql.NullInt64  `json:"parentID"`
	CodeID             sql.NullInt64  `json:"codeID"`
	Stat               sql.NullInt64  `json:"stat"`
}

func (q *QueriesLocal) CreateReferencesDetails(ctx context.Context, arg ReferencesDetailsRequest) error {
	_, err := q.db.ExecContext(ctx, createReferencesDetails,
		arg.ID,
		arg.RefID,
		arg.PurposeDescription,
		arg.ParentID,
		arg.CodeID,
		arg.Stat,
	)
	return err
}

const deleteReferencesDetails = `-- name: DeleteReferencesDetails :exec
DELETE FROM ReferencesDetails WHERE id = $1
`

func (q *QueriesLocal) DeleteReferencesDetails(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteReferencesDetails, id)
	return err
}

type ReferencesDetailsInfo struct {
	ModCtr             int64          `json:"modCtr"`
	BrCode             string         `json:"brCode"`
	ModAction          string         `json:"modAction"`
	ID                 int64          `json:"id"`
	RefID              int64          `json:"refID"`
	PurposeDescription sql.NullString `json:"purposeDescription"`
	ParentID           sql.NullInt64  `json:"parentID"`
	CodeID             sql.NullInt64  `json:"codeID"`
	Stat               sql.NullInt64  `json:"stat"`
}

// -- name: GetReferencesDetails :one
const getReferencesDetails = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, ID, RefID, PurposeDescription, ParentID, CodeID, Stat
FROM OrgParms, ReferencesDetails d
INNER JOIN Modified m on m.UniqueKeyInt1 = d.id 
`

func scanRowReferencesDetails(row *sql.Row) (ReferencesDetailsInfo, error) {
	var i ReferencesDetailsInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.ID,
		&i.RefID,
		&i.PurposeDescription,
		&i.ParentID,
		&i.CodeID,
		&i.Stat,
	)
	return i, err
}

func scanRowsReferencesDetails(rows *sql.Rows) ([]ReferencesDetailsInfo, error) {
	items := []ReferencesDetailsInfo{}
	for rows.Next() {
		var i ReferencesDetailsInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.ID,
			&i.RefID,
			&i.PurposeDescription,
			&i.ParentID,
			&i.CodeID,
			&i.Stat,
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

func (q *QueriesLocal) GetReferencesDetails(ctx context.Context, id int64) (ReferencesDetailsInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'ReferencesDetails' AND Uploaded = 0 and id = $1", getReferencesDetails)
	row := q.db.QueryRowContext(ctx, sql, id)
	return scanRowReferencesDetails(row)
}

type ListReferencesDetailsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) ReferencesDetailsCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, ID, RefID, PurposeDescription, ParentID, CodeID, Stat
FROM OrgParms, ReferencesDetails d
`, filenamePath)
}

func (q *QueriesLocal) ListReferencesDetails(ctx context.Context) ([]ReferencesDetailsInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'ReferencesDetails' AND Uploaded = 0`,
		getReferencesDetails)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsReferencesDetails(rows)
}

// -- name: UpdateReferencesDetails :one
const updateReferencesDetails = `
UPDATE ReferencesDetails SET 
	RefID = $2,
	PurposeDescription = $3,
	ParentID = $4,
	CodeID = $5,
	Stat = $6
WHERE id = $1`

func (q *QueriesLocal) UpdateReferencesDetails(ctx context.Context, arg ReferencesDetailsRequest) error {
	_, err := q.db.ExecContext(ctx, updateReferencesDetails,
		arg.ID,
		arg.RefID,
		arg.PurposeDescription,
		arg.ParentID,
		arg.CodeID,
		arg.Stat,
	)
	return err
}
