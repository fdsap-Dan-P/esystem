package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createTrnHeadSpecsString = `-- name: CreateTrnHeadSpecsString: one
INSERT INTO Trn_Head_Specs_String 
  (Trn_Head_ID, Specs_ID, Value) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Trn_Head_ID, Specs_ID ) DO UPDATE SET
  Value = excluded.Value
RETURNING 
  UUID, Trn_Head_ID, Specs_Code, Specs_ID, Value
`

type TrnHeadSpecsStringRequest struct {
	Uuid      uuid.UUID `json:"uuid"`
	TrnHeadId int64     `json:"trnHeadId"`
	SpecsId   int64     `json:"specsId"`
	Value     string    `json:"value"`
}

func (q *QueriesTransaction) CreateTrnHeadSpecsString(ctx context.Context, arg TrnHeadSpecsStringRequest) (model.TrnHeadSpecsString, error) {
	row := q.db.QueryRowContext(ctx, createTrnHeadSpecsString,
		arg.TrnHeadId,
		arg.SpecsId,
		arg.Value,
	)
	var i model.TrnHeadSpecsString
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const updateTrnHeadSpecsString = `-- name: UpdateTrnHeadSpecsString :one
UPDATE Trn_Head_Specs_String SET 
Trn_Head_ID = $2,
  Specs_ID = $3,
  Value = $4
WHERE uuid = $1
RETURNING UUID, Trn_Head_ID, Specs_Code, Specs_ID, Value
`

func (q *QueriesTransaction) UpdateTrnHeadSpecsString(ctx context.Context, arg TrnHeadSpecsStringRequest) (model.TrnHeadSpecsString, error) {
	row := q.db.QueryRowContext(ctx, updateTrnHeadSpecsString,
		arg.Uuid,
		arg.TrnHeadId,
		arg.SpecsId,
		arg.Value,
	)
	var i model.TrnHeadSpecsString
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const deleteTrnHeadSpecsString = `-- name: DeleteTrnHeadSpecsString :exec
DELETE FROM Trn_Head_Specs_String
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTrnHeadSpecsString(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTrnHeadSpecsString, uuid)
	return err
}

type TrnHeadSpecsStringInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	TrnHeadId       int64        `json:"trnHeadId"`
	SpecsId         int64        `json:"specsId"`
	SpecsCode       string       `json:"specsCode"`
	Item            string       `json:"item"`
	ItemShortName   string       `json:"itemShortName"`
	ItemDescription string       `json:"itemDescription"`
	Value           string       `json:"value"`
	ModCtr          int64        `json:"mod_ctr"`
	Created         sql.NullTime `json:"created"`
	Updated         sql.NullTime `json:"updated"`
}

func populateTrnHeadSpecString(q *QueriesTransaction, ctx context.Context, sql string) (TrnHeadSpecsStringInfo, error) {
	var i TrnHeadSpecsStringInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Item,
		&i.ItemShortName,
		&i.ItemDescription,
		&i.Value,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTrnHeadSpecString2(q *QueriesTransaction, ctx context.Context, sql string) ([]TrnHeadSpecsStringInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []TrnHeadSpecsStringInfo{}
	for rows.Next() {
		var i TrnHeadSpecsStringInfo
		err := rows.Scan(
			&i.Uuid,
			&i.TrnHeadId,
			&i.SpecsCode,
			&i.SpecsId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
			&i.Value,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
		)
		if err != nil {
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

const trnHeadSpecsStringSQL = `-- name: trnHeadSpecsStringSQL
SELECT 
  mr.UUID, d.Trn_Head_ID, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name Item_Short_Name, 
  ref.Remark Item_Description, d.Value,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Trn_Head_Specs_String d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesTransaction) GetTrnHeadSpecsString(ctx context.Context, trnHeadId int64, specsId int64) (TrnHeadSpecsStringInfo, error) {
	return populateTrnHeadSpecString(q, ctx, fmt.Sprintf("%s WHERE d.Trn_Head_ID = %v and d.Specs_ID = %v",
		trnHeadSpecsStringSQL, trnHeadId, specsId))
}

func (q *QueriesTransaction) GetTrnHeadSpecsStringbyUuid(ctx context.Context, uuid uuid.UUID) (TrnHeadSpecsStringInfo, error) {
	return populateTrnHeadSpecString(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", trnHeadSpecsStringSQL, uuid))
}

type ListTrnHeadSpecsStringParams struct {
	TrnHeadId int64 `json:"TrnHeadId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTrnHeadSpecsString(ctx context.Context, arg ListTrnHeadSpecsStringParams) ([]TrnHeadSpecsStringInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Trn_Head_ID = %v LIMIT %d OFFSET %d",
			trnHeadSpecsStringSQL, arg.TrnHeadId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Trn_Head_ID = %v ", trnHeadSpecsStringSQL, arg.TrnHeadId)
	}
	return populateTrnHeadSpecString2(q, ctx, sql)
}
