package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const createAccounts = `-- name: CreateAccounts: one
IF EXISTS (SELECT Acnt_CD FROM Accounts WHERE Acnt_CD = $1) 
BEGIN
  UPDATE Accounts SET 
    Acnt_Title = $2,
    Acnt_Acat_Cd = $3,
    Acnt_Typ = $4,
    MainCD = $5,
    Acnt_Parent_Cd = $6
  WHERE Acnt_Cd = $1;
END ELSE BEGIN
  INSERT INTO Accounts (
	Acnt_CD, Acnt_Title, Acnt_Acat_Cd, Acnt_Typ, MainCD, Acnt_Parent_CD, Acnt_Level
  ) 
  VALUES ($1, $2, $3, $4, $5, $6, 1);
END
`

type AccountsRequest struct {
	Acc      string         `json:"acc"`
	Title    string         `json:"title"`
	Category int64          `json:"category"`
	Type     string         `json:"type"`
	MainCD   sql.NullString `json:"mainCD"`
	Parent   sql.NullString `json:"parent"`
}

func (q *QueriesLocal) CreateAccounts(ctx context.Context, arg AccountsRequest) error {
	_, err := q.db.ExecContext(ctx, createAccounts,
		arg.Acc,
		arg.Title,
		arg.Category,
		arg.Type,
		arg.MainCD,
		arg.Parent,
	)
	return err
}

const deleteAccounts = `-- name: DeleteAccounts :exec
DELETE FROM Accounts WHERE Acnt_CD = $1
`

func (q *QueriesLocal) DeleteAccounts(ctx context.Context, acc string) error {
	_, err := q.db.ExecContext(ctx, deleteAccounts, acc)
	return err
}

type AccountsInfo struct {
	ModCtr    int64          `json:"modCtr"`
	BrCode    string         `json:"brCode"`
	ModAction string         `json:"modAction"`
	Acc       string         `json:"acc"`
	Title     string         `json:"title"`
	Category  int64          `json:"category"`
	Type      string         `json:"type"`
	MainCD    sql.NullString `json:"mainCD"`
	Parent    sql.NullString `json:"parent"`
}

// -- name: GetAccounts :one

const getAccounts = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, Acnt_CD, Acnt_Title, Acnt_Acat_Cd, Acnt_Typ, MainCD, Acnt_Parent_CD
FROM OrgParms, Accounts d
INNER JOIN Modified m on m.UniqueKeyString1 = d.Acnt_Cd
`

func scanRowAccounts(row *sql.Row) (AccountsInfo, error) {
	var i AccountsInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.Acc,
		&i.Title,
		&i.Category,
		&i.Type,
		&i.MainCD,
		&i.Parent,
	)
	return i, err
}

func scanRowsAccounts(rows *sql.Rows) ([]AccountsInfo, error) {
	items := []AccountsInfo{}
	for rows.Next() {
		var i AccountsInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.Acc,
			&i.Title,
			&i.Category,
			&i.Type,
			&i.MainCD,
			&i.Parent,
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

func (q *QueriesLocal) GetAccounts(ctx context.Context, acc string) (AccountsInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'Accounts' AND Uploaded = 0 and Acnt_Cd = $1", getAccounts)
	row := q.db.QueryRowContext(ctx, sql, acc)
	return scanRowAccounts(row)
}

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) AccountsCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx,
		`SELECT 
  		   0 ModCtr, OrgParms.DefBranch_Code BrCode, Acnt_CD, Acnt_Title, Acnt_Acat_Cd, Acnt_Typ, MainCD, Acnt_Parent_CD
		 FROM OrgParms, Accounts d
	`, filenamePath)
}

func (q *QueriesLocal) ListAccounts(ctx context.Context) ([]AccountsInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'Accounts' AND Uploaded = 0`,
		getAccounts)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsAccounts(rows)
}

// -- name: UpdateAccounts :one
const updateAccounts = `
UPDATE Accounts SET 
  Acnt_Title = $2,
  Acnt_Acat_Cd = $3,
  Acnt_Typ = $4,
  MainCD = $5,
  Acnt_Parent_Cd = $6
WHERE Acnt_Cd = $1`

func (q *QueriesLocal) UpdateAccounts(ctx context.Context, arg AccountsRequest) error {
	_, err := q.db.ExecContext(ctx, updateAccounts,
		arg.Acc,
		arg.Title,
		arg.Category,
		arg.Type,
		arg.MainCD,
		arg.Parent,
	)
	return err
}
