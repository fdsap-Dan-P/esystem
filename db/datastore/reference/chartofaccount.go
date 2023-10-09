package db

import (
	"context"
	"database/sql"

	"simplebank/model"

	"github.com/google/uuid"
)

const createChartofAccount = `-- name: CreateChartofAccount: one
INSERT INTO Chartof_Account (
Acc, Active, Contra_Account, Normal_Balance, Title, 
Parent_Id, Short_Name, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8
) RETURNING Id, UUId, Acc, Active, Contra_Account, Normal_Balance, Title, 
Parent_Id, Short_Name, Other_Info
`

type ChartofAccountRequest struct {
	Id            int64          `json:"id"`
	Uuid          uuid.UUID      `json:"uuid"`
	Acc           string         `json:"acc"`
	Active        bool           `json:"active"`
	ContraAccount bool           `json:"contraAccount"`
	NormalBalance bool           `json:"normalBalance"`
	Title         string         `json:"title"`
	ParentId      int64          `json:"parentId"`
	ShortName     string         `json:"shortName"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}

func (q *QueriesReference) CreateChartofAccount(ctx context.Context, arg ChartofAccountRequest) (model.ChartofAccount, error) {
	row := q.db.QueryRowContext(ctx, createChartofAccount,
		arg.Acc,
		arg.Active,
		arg.ContraAccount,
		arg.NormalBalance,
		arg.Title,
		arg.ParentId,
		arg.ShortName,
		arg.OtherInfo,
	)
	var i model.ChartofAccount
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Acc,
		&i.Active,
		&i.ContraAccount,
		&i.NormalBalance,
		&i.Title,
		&i.ParentId,
		&i.ShortName,
		&i.OtherInfo,
	)
	return i, err
}

const deleteChartofAccount = `-- name: DeleteChartofAccount :exec
DELETE FROM Chartof_Account
WHERE id = $1
`

func (q *QueriesReference) DeleteChartofAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteChartofAccount, id)
	return err
}

type ChartofAccountInfo struct {
	Id            int64          `json:"id"`
	Uuid          uuid.UUID      `json:"uuid"`
	Acc           string         `json:"acc"`
	Active        bool           `json:"active"`
	ContraAccount bool           `json:"contraAccount"`
	NormalBalance bool           `json:"normalBalance"`
	Title         string         `json:"title"`
	ParentId      int64          `json:"parentId"`
	ShortName     string         `json:"shortName"`
	OtherInfo     sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getChartofAccount = `-- name: GetChartofAccount :one
SELECT 
Id, mr.UUId, Acc, Active, Contra_Account, 
Normal_Balance, Title, Parent_Id, Short_Name, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Chartof_Account d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesReference) GetChartofAccount(ctx context.Context, id int64) (ChartofAccountInfo, error) {
	row := q.db.QueryRowContext(ctx, getChartofAccount, id)
	var i ChartofAccountInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Acc,
		&i.Active,
		&i.ContraAccount,
		&i.NormalBalance,
		&i.Title,
		&i.ParentId,
		&i.ShortName,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getChartofAccountbyAcc = `-- name: GetChartofAccountbyAcc :one
SELECT 
Id, mr.UUId, Acc, Active, Contra_Account, 
Normal_Balance, Title, Parent_Id, Short_Name, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Chartof_Account d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE acc = $1 LIMIT 1
`

func (q *QueriesReference) GetChartofAccountbyAcc(ctx context.Context, acc string) (ChartofAccountInfo, error) {
	row := q.db.QueryRowContext(ctx, getChartofAccountbyAcc, acc)
	var i ChartofAccountInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Acc,
		&i.Active,
		&i.ContraAccount,
		&i.NormalBalance,
		&i.Title,
		&i.ParentId,
		&i.ShortName,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getChartofAccountbyTitle = `-- name: GetChartofAccountbyTitle :one
SELECT 
Id, mr.UUId, Acc, Active, Contra_Account, 
Normal_Balance, Title, Parent_Id, Short_Name, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Chartof_Account d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE d.Title = $1 LIMIT 1
`

func (q *QueriesReference) GetChartofAccountbyTitle(ctx context.Context, title string) (ChartofAccountInfo, error) {
	row := q.db.QueryRowContext(ctx, getChartofAccountbyTitle, title)
	var i ChartofAccountInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Acc,
		&i.Active,
		&i.ContraAccount,
		&i.NormalBalance,
		&i.Title,
		&i.ParentId,
		&i.ShortName,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getChartofAccountbyUuId = `-- name: GetChartofAccountbyUuId :one
SELECT 
Id, mr.UUId, Acc, Active, Contra_Account, 
Normal_Balance, Title, Parent_Id, Short_Name, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Chartof_Account d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesReference) GetChartofAccountbyUuId(ctx context.Context, uuid uuid.UUID) (ChartofAccountInfo, error) {
	row := q.db.QueryRowContext(ctx, getChartofAccountbyUuId, uuid)
	var i ChartofAccountInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Acc,
		&i.Active,
		&i.ContraAccount,
		&i.NormalBalance,
		&i.Title,
		&i.ParentId,
		&i.ShortName,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listChartofAccount = `-- name: ListChartofAccount:many
SELECT 
Id, mr.UUId, Acc, Active, Contra_Account, 
Normal_Balance, Title, Parent_Id, Short_Name, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Chartof_Account d INNER JOIN Main_Record mr on mr.UUId = d.UUId
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListChartofAccountParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesReference) ListChartofAccount(ctx context.Context, arg ListChartofAccountParams) ([]ChartofAccountInfo, error) {
	rows, err := q.db.QueryContext(ctx, listChartofAccount, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ChartofAccountInfo{}
	for rows.Next() {
		var i ChartofAccountInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.Acc,
			&i.Active,
			&i.ContraAccount,
			&i.NormalBalance,
			&i.Title,
			&i.ParentId,
			&i.ShortName,
			&i.OtherInfo,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
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

const updateChartofAccount = `-- name: UpdateChartofAccount :one
UPDATE Chartof_Account SET 
Acc = $2,
Active = $3,
Contra_Account = $4,
Normal_Balance = $5,
Title = $6,
Parent_Id = $7,
Short_Name = $8,
Other_Info = $9
WHERE id = $1
RETURNING Id, UUId, Acc, Active, Contra_Account, Normal_Balance, Title, 
Parent_Id, Short_Name, Other_Info
`

func (q *QueriesReference) UpdateChartofAccount(ctx context.Context, arg ChartofAccountRequest) (model.ChartofAccount, error) {
	row := q.db.QueryRowContext(ctx, updateChartofAccount,
		arg.Id,
		arg.Acc,
		arg.Active,
		arg.ContraAccount,
		arg.NormalBalance,
		arg.Title,
		arg.ParentId,
		arg.ShortName,
		arg.OtherInfo,
	)
	var i model.ChartofAccount
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.Acc,
		&i.Active,
		&i.ContraAccount,
		&i.NormalBalance,
		&i.Title,
		&i.ParentId,
		&i.ShortName,
		&i.OtherInfo,
	)
	return i, err
}
