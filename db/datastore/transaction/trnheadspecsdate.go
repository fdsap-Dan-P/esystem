package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
)

const createTrnHeadSpecsDate = `-- name: CreateTrnHeadSpecsDate: one
INSERT INTO Trn_Head_Specs_Date 
  (Trn_Head_Id, Specs_ID, Value, Value2) 
VALUES 
  ($1, $2, $3, $4) 
ON CONFLICT( Trn_Head_Id, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Trn_Head_Id, Specs_Code, Specs_ID, Value, Value2
`

type TrnHeadSpecsDateRequest struct {
	Uuid      uuid.UUID `json:"uuid"`
	TrnHeadId int64     `json:"trnHeadId"`
	SpecsId   int64     `json:"specsId"`
	Value     time.Time `json:"value"`
	Value2    time.Time `json:"value2"`
}

func (q *QueriesTransaction) CreateTrnHeadSpecsDate(ctx context.Context, arg TrnHeadSpecsDateRequest) (model.TrnHeadSpecsDate, error) {
	row := q.db.QueryRowContext(ctx, createTrnHeadSpecsDate,
		arg.TrnHeadId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.TrnHeadSpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const updateTrnHeadSpecsDate = `-- name: UpdateTrnHeadSpecsDate :one
UPDATE Trn_Head_Specs_Date SET 
Trn_Head_Id = $2,
	Specs_ID = $3,
	Value = $4,
	Value2 = $5
WHERE uuid = $1
RETURNING UUID, Trn_Head_Id, Specs_Code, Specs_ID, Value, Value2
`

func (q *QueriesTransaction) UpdateTrnHeadSpecsDate(ctx context.Context, arg TrnHeadSpecsDateRequest) (model.TrnHeadSpecsDate, error) {
	row := q.db.QueryRowContext(ctx, updateTrnHeadSpecsDate,
		arg.Uuid,
		arg.TrnHeadId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.TrnHeadSpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const deleteTrnHeadSpecsDate = `-- name: DeleteTrnHeadSpecsDate :exec
DELETE FROM Trn_Head_Specs_Date
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTrnHeadSpecsDate(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTrnHeadSpecsDate, uuid)
	return err
}

type TrnHeadSpecsDateInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	TrnHeadId       int64        `json:"trnHeadId"`
	SpecsCode       string       `json:"specsCode"`
	SpecsId         int64        `json:"specsId"`
	Item            string       `json:"item"`
	ItemShortName   string       `json:"itemShortName"`
	ItemDescription string       `json:"itemDescription"`
	Value           time.Time    `json:"value"`
	Value2          time.Time    `json:"value2"`
	ModCtr          int64        `json:"modCtr"`
	Created         sql.NullTime `json:"created"`
	Updated         sql.NullTime `json:"updated"`
}

func populateIdentitySpecDate(q *QueriesTransaction, ctx context.Context, sql string) (TrnHeadSpecsDateInfo, error) {
	var i TrnHeadSpecsDateInfo
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
		&i.Value2,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateIdentitySpecDate2(q *QueriesTransaction, ctx context.Context, sql string) ([]TrnHeadSpecsDateInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []TrnHeadSpecsDateInfo{}
	for rows.Next() {
		var i TrnHeadSpecsDateInfo
		err := rows.Scan(
			&i.Uuid,
			&i.TrnHeadId,
			&i.SpecsCode,
			&i.SpecsId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
			&i.Value,
			&i.Value2,

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

const TrnHeadSpecsDateSQL = `-- name: TrnHeadSpecsDateSQL
SELECT 
  mr.UUID, d.Trn_Head_ID, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Trn_Head_Specs_Date d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesTransaction) GetTrnHeadSpecsDate(ctx context.Context, trnHeadId int64, specsId int64) (TrnHeadSpecsDateInfo, error) {
	return populateIdentitySpecDate(q, ctx, fmt.Sprintf("%s WHERE d.Trn_Head_ID = %v and d.Specs_ID = %v",
		TrnHeadSpecsDateSQL, trnHeadId, specsId))
}

func (q *QueriesTransaction) GetTrnHeadSpecsDatebyUuid(ctx context.Context, uuid uuid.UUID) (TrnHeadSpecsDateInfo, error) {
	return populateIdentitySpecDate(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", TrnHeadSpecsDateSQL, uuid))
}

type ListTrnHeadSpecsDateParams struct {
	TrnHeadId int64 `json:"TrnHeadId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTrnHeadSpecsDate(ctx context.Context, arg ListTrnHeadSpecsDateParams) ([]TrnHeadSpecsDateInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Trn_Head_ID = %v LIMIT %d OFFSET %d",
			TrnHeadSpecsDateSQL, arg.TrnHeadId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Trn_Head_ID = %v ", TrnHeadSpecsDateSQL, arg.TrnHeadId)
	}
	return populateIdentitySpecDate2(q, ctx, sql)
}
