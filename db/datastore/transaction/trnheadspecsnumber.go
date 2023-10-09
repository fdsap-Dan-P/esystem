package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createTrnHeadSpecsNumber = `-- name: CreateTrnHeadSpecsNumber: one
INSERT INTO Trn_Head_Specs_Number 
  (Trn_Head_Id, Specs_ID, Value, Value2) 
VALUES 
  ($1, $2, $3, $4, $5) 
ON CONFLICT( Trn_Head_Id, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Trn_Head_Id, Specs_Code, Specs_ID, Value, Value2
`

type TrnHeadSpecsNumberRequest struct {
	Uuid      uuid.UUID       `json:"uuid"`
	TrnHeadId int64           `json:"trnHeadId"`
	SpecsId   int64           `json:"specsId"`
	Value     decimal.Decimal `json:"value"`
	Value2    decimal.Decimal `json:"value2"`
	MeasureId sql.NullInt64   `json:"measureId"`
}

func (q *QueriesTransaction) CreateTrnHeadSpecsNumber(ctx context.Context, arg TrnHeadSpecsNumberRequest) (model.TrnHeadSpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, createTrnHeadSpecsNumber,
		arg.TrnHeadId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.TrnHeadSpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const updateTrnHeadSpecsNumber = `-- name: UpdateTrnHeadSpecsNumber :one
UPDATE Trn_Head_Specs_Number SET 
  Trn_Head_Id = $2,
  Specs_ID = $3,
  Value = $4,
  Value2 = $5,
  Measure_Id = $6
WHERE uuid = $1
RETURNING UUID, Trn_Head_Id, Specs_Code, Specs_ID, Value, Value2, Measure_Id
`

func (q *QueriesTransaction) UpdateTrnHeadSpecsNumber(ctx context.Context, arg TrnHeadSpecsNumberRequest) (model.TrnHeadSpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, updateTrnHeadSpecsNumber,
		arg.Uuid,
		arg.TrnHeadId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.TrnHeadSpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const deleteTrnHeadSpecsNumber = `-- name: DeleteTrnHeadSpecsNumber :exec
DELETE FROM Trn_Head_Specs_Number
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTrnHeadSpecsNumber(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTrnHeadSpecsNumber, uuid)
	return err
}

type TrnHeadSpecsNumberInfo struct {
	Uuid            uuid.UUID       `json:"uuid"`
	TrnHeadId       int64           `json:"trnHeadId"`
	SpecsId         int64           `json:"specsId"`
	SpecsCode       string          `json:"specsCode"`
	Item            string          `json:"item"`
	ItemShortName   string          `json:"itemShortName"`
	ItemDescription string          `json:"itemDescription"`
	Value           decimal.Decimal `json:"value"`
	Value2          decimal.Decimal `json:"value2"`
	MeasureId       sql.NullInt64   `json:"measureId"`
	Measure         sql.NullString  `json:"measure"`
	MeasureUnit     sql.NullString  `json:"measureUnit"`
	ModCtr          int64           `json:"mod_ctr"`
	Created         sql.NullTime    `json:"created"`
	Updated         sql.NullTime    `json:"updated"`
}

func populateIdentitySpecNumber(q *QueriesTransaction, ctx context.Context, sql string) (TrnHeadSpecsNumberInfo, error) {
	var i TrnHeadSpecsNumberInfo
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
		&i.MeasureId,
		&i.Measure,
		&i.MeasureUnit,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateIdentitySpecNumber2(q *QueriesTransaction, ctx context.Context, sql string) ([]TrnHeadSpecsNumberInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []TrnHeadSpecsNumberInfo{}
	for rows.Next() {
		var i TrnHeadSpecsNumberInfo
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
			&i.MeasureId,
			&i.Measure,
			&i.MeasureUnit,

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

const trnHeadSpecsNumberSQL = `-- name: trnHeadSpecsNumberSQL
SELECT 
  mr.UUID, d.Trn_Head_Id, d.Specs_Code, d.Specs_ID, ref.Title Item, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2, d.Measure_Id, mea.Title Measure, mea.Short_Name Measure_Unit, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Trn_Head_Specs_Number d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
LEFT JOIN Reference mea on mea.ID = Measure_Id`

func (q *QueriesTransaction) GetTrnHeadSpecsNumber(ctx context.Context, trnHeadId int64, specsId int64) (TrnHeadSpecsNumberInfo, error) {
	return populateIdentitySpecNumber(q, ctx, fmt.Sprintf("%s WHERE d.Trn_Head_Id = %v and d.Specs_ID = %v",
		trnHeadSpecsNumberSQL, trnHeadId, specsId))
}

func (q *QueriesTransaction) GetTrnHeadSpecsNumberbyUuid(ctx context.Context, uuid uuid.UUID) (TrnHeadSpecsNumberInfo, error) {
	return populateIdentitySpecNumber(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", trnHeadSpecsNumberSQL, uuid))
}

type ListTrnHeadSpecsNumberParams struct {
	TrnHeadId int64 `json:"TrnHeadId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTrnHeadSpecsNumber(ctx context.Context, arg ListTrnHeadSpecsNumberParams) ([]TrnHeadSpecsNumberInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Trn_Head_Id = %v LIMIT %d OFFSET %d",
			trnHeadSpecsNumberSQL, arg.TrnHeadId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Trn_Head_Id = %v ", trnHeadSpecsNumberSQL, arg.TrnHeadId)
	}
	return populateIdentitySpecNumber2(q, ctx, sql)
}
