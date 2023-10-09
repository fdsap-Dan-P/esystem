package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createTrnHeadSpecsRef = `-- name: CreateTrnHeadSpecsRef: one
INSERT INTO Trn_Head_Specs_Ref 
  (Trn_Head_Id, Specs_ID, Ref_Id) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Trn_Head_Id, Specs_ID ) DO UPDATE SET
Ref_Id = excluded.Ref_Id
RETURNING 
  UUID, Trn_Head_Id, Specs_Code, Specs_ID, Ref_Id
`

type TrnHeadSpecsRefRequest struct {
	Uuid      uuid.UUID `json:"uuid"`
	TrnHeadId int64     `json:"trnHeadId"`
	SpecsId   int64     `json:"specsId"`
	RefId     int64     `json:"refId"`
}

func (q *QueriesTransaction) CreateTrnHeadSpecsRef(ctx context.Context, arg TrnHeadSpecsRefRequest) (model.TrnHeadSpecsRef, error) {
	row := q.db.QueryRowContext(ctx, createTrnHeadSpecsRef,
		arg.TrnHeadId,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.TrnHeadSpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const updateTrnHeadSpecsRef = `-- name: UpdateTrnHeadSpecsRef :one
UPDATE Trn_Head_Specs_Ref SET 
  Trn_Head_Id = $2,
  Specs_ID = $3,
  Ref_ID = $4
WHERE uuid = $1
RETURNING UUID, Trn_Head_Id, Specs_Code, Specs_ID, Ref_ID
`

func (q *QueriesTransaction) UpdateTrnHeadSpecsRef(ctx context.Context, arg TrnHeadSpecsRefRequest) (model.TrnHeadSpecsRef, error) {
	row := q.db.QueryRowContext(ctx, updateTrnHeadSpecsRef,
		arg.Uuid,
		arg.TrnHeadId,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.TrnHeadSpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const deleteTrnHeadSpecsRef = `-- name: DeleteTrnHeadSpecsRef :exec
DELETE FROM Trn_Head_Specs_Ref
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTrnHeadSpecsRef(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTrnHeadSpecsRef, uuid)
	return err
}

type TrnHeadSpecsRefInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	TrnHeadId       int64          `json:"trnHeadId"`
	SpecsCode       string         `json:"specsCode"`
	SpecsId         int64          `json:"specsId"`
	Item            string         `json:"item"`
	ItemShortName   string         `json:"itemShortName"`
	ItemDescription string         `json:"itemDescription"`
	RefId           int64          `json:"refId"`
	MeasureId       sql.NullInt64  `json:"measureId"`
	Measure         sql.NullString `json:"measure"`
	MeasureUnit     sql.NullString `json:"measureUnit"`
	ModCtr          int64          `json:"mod_ctr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

func populateTrnHeadSpecRef(q *QueriesTransaction, ctx context.Context, sql string) (TrnHeadSpecsRefInfo, error) {
	var i TrnHeadSpecsRefInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Item,
		&i.ItemShortName,
		&i.ItemDescription,
		&i.RefId,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTrnHeadSpecRef2(q *QueriesTransaction, ctx context.Context, sql string) ([]TrnHeadSpecsRefInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []TrnHeadSpecsRefInfo{}
	for rows.Next() {
		var i TrnHeadSpecsRefInfo
		err := rows.Scan(
			&i.Uuid,
			&i.TrnHeadId,
			&i.SpecsCode,
			&i.SpecsId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
			&i.RefId,

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

const trnHeadSpecsRefSQL = `-- name: trnHeadSpecsRefSQL
SELECT 
  mr.UUID, d.Trn_Head_Id, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, d.Ref_Id,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Trn_Head_Specs_Ref d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesTransaction) GetTrnHeadSpecsRef(ctx context.Context, trnHeadId int64, specsId int64) (TrnHeadSpecsRefInfo, error) {
	return populateTrnHeadSpecRef(q, ctx, fmt.Sprintf("%s WHERE d.Trn_Head_Id = %v and d.Specs_ID = %v",
		trnHeadSpecsRefSQL, trnHeadId, specsId))
}

func (q *QueriesTransaction) GetTrnHeadSpecsRefbyUuid(ctx context.Context, uuid uuid.UUID) (TrnHeadSpecsRefInfo, error) {
	return populateTrnHeadSpecRef(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", trnHeadSpecsRefSQL, uuid))
}

type ListTrnHeadSpecsRefParams struct {
	TrnHeadId int64 `json:"TrnHeadId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTrnHeadSpecsRef(ctx context.Context, arg ListTrnHeadSpecsRefParams) ([]TrnHeadSpecsRefInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Trn_Head_Id = %v LIMIT %d OFFSET %d",
			trnHeadSpecsRefSQL, arg.TrnHeadId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Trn_Head_Id = %v ", trnHeadSpecsRefSQL, arg.TrnHeadId)
	}
	return populateTrnHeadSpecRef2(q, ctx, sql)
}
