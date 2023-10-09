package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const createLnBeneficiary = `-- name: CreateLnBeneficiary: one
INSERT INTO LnBeneficiary(
	Acc, Beneficiary, BDay, Age, Educ_Lvl, Gender, BFLName, BffName, BFMName, Remarks)
SELECT $1, dbo.fullname($5, $6, $7), $2, floor(round(datediff(dd, '1979-04-14', '2023-04-14')/365.25,3)), 
  $3, $4, $5, $6, $7, $8
`

type LnBeneficiaryRequest struct {
	Acc        string         `json:"acc"`
	BDay       time.Time      `json:"bDay"`
	EducLvl    string         `json:"educLvl"`
	Gender     bool           `json:"gender"`
	LastName   sql.NullString `json:"lastName"`
	FirstName  sql.NullString `json:"firstName"`
	MiddleName sql.NullString `json:"middleName"`
	Remarks    sql.NullString `json:"remarks"`
}

func (q *QueriesLocal) CreateLnBeneficiary(ctx context.Context, arg LnBeneficiaryRequest) error {
	_, err := q.db.ExecContext(ctx, createLnBeneficiary,
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
  DELETE FROM LnBeneficiary WHERE Acc = $1
`

func (q *QueriesLocal) DeleteLnBeneficiary(ctx context.Context, acc string) error {
	_, err := q.db.ExecContext(ctx, deleteLnBeneficiary, acc)
	return err
}

const getLnBeneficiary = `-- name: GetLnBeneficiary :one
SELECT 
  m.ModCtr, m.ModAction, OrgParms.DefBranch_Code, Acc, BDay, Educ_Lvl, Gender, BFLName, BffName, BFMName, Remarks
FROM OrgParms,
 (SELECT
    Acc, Max(BDay) BDay, Max(Educ_Lvl) Educ_Lvl, Max(CASE WHEN Gender = 1 THEN 1 ELSE 0 END) Gender, Max(BffName) BffName, Max(BFLName) BFLName, Max(BFMName) BFMName, Max(Remarks) Remarks
  FROM LnBeneficiary
  GROUP BY Acc
  ) d
INNER JOIN
    (SELECT Max(ModCtr) ModCtr, ModAction, UniqueKeyString1, min(CASE WHEN Uploaded = 1 THEN 1 ELSE 0 END) Uploaded
     FROM Modified
     WHERE TableName = 'LnBeneficiary'
     GROUP BY ModAction, UniqueKeyString1 
     ) m on m.UniqueKeyString1 = Acc
`

func scanRowLnBeneficiary(row *sql.Row) (LnBeneficiary, error) {
	var i LnBeneficiary
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

func scanRowsLnBeneficiary(rows *sql.Rows) ([]LnBeneficiary, error) {
	items := []LnBeneficiary{}
	for rows.Next() {
		var i LnBeneficiary
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

func (q *QueriesLocal) GetLnBeneficiary(ctx context.Context, acc string) (LnBeneficiary, error) {
	sql := fmt.Sprintf("%s WHERE Acc = $1", getLnBeneficiary)
	row := q.db.QueryRowContext(ctx, sql, acc)
	return scanRowLnBeneficiary(row)
}

type ListLnBeneficiaryParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) ListLnBeneficiary(ctx context.Context) ([]LnBeneficiary, error) {
	sql := fmt.Sprintf(`%v WHERE Uploaded = 0`, getLnBeneficiary)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsLnBeneficiary(rows)
}

const updateLnBeneficiary = `-- name: UpdateLnBeneficiary :one
UPDATE LnBeneficiary SET 
  BDay = $2,
  Educ_Lvl = $3,
  Gender = $4,
  BfLName = $5,
  BffName = $6,
  BfMName = $7,
  Remarks = $8
WHERE Acc = $1 
`

func (q *QueriesLocal) UpdateLnBeneficiary(ctx context.Context, arg LnBeneficiaryRequest) error {
	_, err := q.db.ExecContext(ctx, updateLnBeneficiary,
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

func (q *QueriesLocal) LnBeneficiaryCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code, 
  Acc, BDay, Educ_Lvl, Gender, BffName, BFLName, BFMName, Remarks
FROM OrgParms,
   (SELECT
      Acc, Max(BDay) BDay, Max(Educ_Lvl) Educ_Lvl, Max(CASE WHEN Gender = 1 THEN 1 ELSE 0 END) Gender, 
	  Max(BffName) BffName, Max(BFLName) BFLName, Max(BFMName) BFMName, Max(Remarks) Remarks
	FROM LnBeneficiary
    GROUP BY Acc
  ) d 
`, filenamePath)
}
