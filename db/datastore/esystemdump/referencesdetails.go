package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	model "simplebank/db/datastore/esystemlocal"
)

const createReferencesDetails = `-- name: CreateReferencesDetails: one
INSERT INTO esystemdump.ReferencesDetails(
   ModCtr, BrCode, ModAction, ID, RefID, PurposeDescription, ParentID, CodeID, Stat )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT (brCode, iD, ModAction)
DO UPDATE SET
	ModCtr =  EXCLUDED.ModCtr,
	RefID =  EXCLUDED.RefID,
	PurposeDescription =  EXCLUDED.PurposeDescription,
	ParentID =  EXCLUDED.ParentID,
	CodeID =  EXCLUDED.CodeID,
	Stat =  EXCLUDED.Stat
`

func (q *QueriesDump) CreateReferencesDetails(ctx context.Context, arg model.ReferencesDetails) error {
	_, err := q.db.ExecContext(ctx, createReferencesDetails,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
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
DELETE FROM esystemdump.ReferencesDetails WHERE BrCode = $1 and ID = $2
`

func (q *QueriesDump) DeleteReferencesDetails(ctx context.Context, brCode string, iD int64) error {
	_, err := q.db.ExecContext(ctx, deleteReferencesDetails, brCode, iD)
	return err
}

const getReferencesDetails = `-- name: GetReferencesDetails :one
SELECT
ModCtr, BrCode, ModAction, ID, RefID, PurposeDescription, ParentID, CodeID, Stat
FROM esystemdump.ReferencesDetails
`

func scanRowReferencesDetails(row *sql.Row) (model.ReferencesDetails, error) {
	var i model.ReferencesDetails
	err := row.Scan(
		&i.ModCtr,
		&i.BrCode,
		&i.ModAction,
		&i.ID,
		&i.RefID,
		&i.PurposeDescription,
		&i.ParentID,
		&i.CodeID,
		&i.Stat,
	)
	return i, err
}

func scanRowsReferencesDetails(rows *sql.Rows) ([]model.ReferencesDetails, error) {
	items := []model.ReferencesDetails{}
	for rows.Next() {
		var i model.ReferencesDetails
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

func (q *QueriesDump) GetReferencesDetails(ctx context.Context, brCode string, iD int64) (model.ReferencesDetails, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and ID = $2", getReferencesDetails)
	row := q.db.QueryRowContext(ctx, sql, brCode, iD)
	return scanRowReferencesDetails(row)
}

type ListReferencesDetailsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListReferencesDetails(ctx context.Context, lastModCtr int64) ([]model.ReferencesDetails, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getReferencesDetails)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsReferencesDetails(rows)
}

const updateReferencesDetails = `-- name: UpdateReferencesDetails :one
UPDATE esystemdump.ReferencesDetails SET 
	ModCtr = $1,
	RefID = $5,
	PurposeDescription = $6,
	ParentID = $7,
	CodeID = $8,
	Stat = $9
WHERE BrCode = $2 and ID = $4 and ModAction = $3
`

func (q *QueriesDump) UpdateReferencesDetails(ctx context.Context, arg model.ReferencesDetails) error {
	_, err := q.db.ExecContext(ctx, updateReferencesDetails,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.ID,
		arg.RefID,
		arg.PurposeDescription,
		arg.ParentID,
		arg.CodeID,
		arg.Stat,
	)
	return err
}
