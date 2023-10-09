package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	model "simplebank/db/datastore/esystemlocal"
)

const createLnBeneficiary = `-- name: CreateLnBeneficiary: one
INSERT INTO esystemdump.LnBeneficiary(
	ModCtr, BrCode, ModAction, Acc, BDay, Educ_Lvl, Gender, Last_Name, First_Name, Middle_Name, Remarks)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT(BrCode, Acc, ModAction) DO UPDATE SET
ModCtr = excluded.ModCtr,
ModAction = excluded.ModAction,
BDay = excluded.BDay,
Educ_Lvl = excluded.Educ_Lvl,
Gender = excluded.Gender,
Last_Name = excluded.Last_Name,
First_Name = excluded.First_Name,
Middle_Name = excluded.Middle_Name,
Remarks = excluded.Remarks
`

func (q *QueriesDump) CreateLnBeneficiary(ctx context.Context, arg model.LnBeneficiary) error {
	_, err := q.db.ExecContext(ctx, createLnBeneficiary,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.BDay,
		arg.EducLvl,
		arg.Gender,
		arg.LastName,
		arg.FirstName,
		arg.MiddleName,
		arg.Remarks,
	)
	return err
}

const deleteLnBeneficiary = `-- name: DeleteLnBeneficiary :exec
  DELETE FROM esystemdump.LnBeneficiary WHERE BrCode = $1 and Acc = $2
`

func (q *QueriesDump) DeleteLnBeneficiary(ctx context.Context, brCode string, acc string) error {
	_, err := q.db.ExecContext(ctx, deleteLnBeneficiary, brCode, acc)
	return err
}

const getLnBeneficiary = `-- name: GetLnBeneficiary :one
SELECT 
  ModCtr, ModAction, BrCode, Acc, BDay, Educ_Lvl, Gender, Last_Name, First_Name, Middle_Name, Remarks
FROM esystemdump.LnBeneficiary
`

func scanRowLnBeneficiary(row *sql.Row) (model.LnBeneficiary, error) {
	var i model.LnBeneficiary
	err := row.Scan(
		&i.ModCtr,
		&i.ModAction,
		&i.BrCode,
		&i.Acc,
		&i.BDay,
		&i.EducLvl,
		&i.Gender,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.Remarks,
	)
	return i, err
}

func scanRowsLnBeneficiary(rows *sql.Rows) ([]model.LnBeneficiary, error) {
	items := []model.LnBeneficiary{}
	for rows.Next() {
		var i model.LnBeneficiary
		if err := rows.Scan(
			&i.ModCtr,
			&i.ModAction,
			&i.BrCode,
			&i.Acc,
			&i.BDay,
			&i.EducLvl,
			&i.Gender,
			&i.LastName,
			&i.FirstName,
			&i.MiddleName,
			&i.Remarks,
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

func (q *QueriesDump) GetLnBeneficiary(ctx context.Context, brCode string, acc string) (model.LnBeneficiary, error) {
	sql := fmt.Sprintf("%s WHERE BrCode = $1 and Acc = $2", getLnBeneficiary)
	row := q.db.QueryRowContext(ctx, sql, brCode, acc)
	return scanRowLnBeneficiary(row)
}

type ListLnBeneficiaryParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesDump) ListLnBeneficiary(ctx context.Context, lastModCtr int64) ([]model.LnBeneficiary, error) {
	sql := fmt.Sprintf(`%v WHERE ModCtr > $1`, getLnBeneficiary)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql, lastModCtr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsLnBeneficiary(rows)
}

const updateLnBeneficiary = `-- name: UpdateLnBeneficiary :one
UPDATE esystemdump.LnBeneficiary SET 
  ModCtr = $1,
  BDay = $5,
  Educ_Lvl = $6,
  Gender = $7,
  Last_Name = $8,
  First_Name = $9,
  Middle_Name = $10,
  Remarks = $11
WHERE BrCode = $2 and Acc = $4 and ModAction = $3
`

func (q *QueriesDump) UpdateLnBeneficiary(ctx context.Context, arg model.LnBeneficiary) error {
	_, err := q.db.ExecContext(ctx, updateLnBeneficiary,
		arg.ModCtr,
		arg.BrCode,
		arg.ModAction,
		arg.Acc,
		arg.BDay,
		arg.EducLvl,
		arg.Gender,
		arg.LastName,
		arg.FirstName,
		arg.MiddleName,
		arg.Remarks,
	)
	return err
}
