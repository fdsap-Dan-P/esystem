package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createIdentitySpecsNumber = `-- name: CreateIdentitySpecsNumber: one
INSERT INTO Identity_Specs_Number 
  (IIID, Specs_ID, Value, Value2, Measure_Id) 
VALUES 
  ($1, $2, $3, $4, $5) 
ON CONFLICT( IIID, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, IIID, Specs_Code, Specs_ID, Value, Value2, Measure_Id
`

type IdentitySpecsNumberRequest struct {
	Uuid      uuid.UUID       `json:"uuid"`
	Iiid      int64           `json:"iiid"`
	SpecsId   int64           `json:"specsId"`
	Value     decimal.Decimal `json:"value"`
	Value2    decimal.Decimal `json:"value2"`
	MeasureId decimal.Decimal `json:"measureId"`
}

func (q *QueriesIdentity) CreateIdentitySpecsNumber(ctx context.Context, arg IdentitySpecsNumberRequest) (model.IdentitySpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, createIdentitySpecsNumber,
		arg.Iiid,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.IdentitySpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const updateIdentitySpecsNumber = `-- name: UpdateIdentitySpecsNumber :one
UPDATE Identity_Specs_Number SET 
  IIID = $2,
  Specs_ID = $3,
  Value = $4,
  Value2 = $5
  Measure_Id = $6
WHERE uuid = $1
RETURNING UUID, IIID, Specs_Code, Specs_Id, Value, Value2, Measure_Id
`

func (q *QueriesIdentity) UpdateIdentitySpecsNumber(ctx context.Context, arg IdentitySpecsNumberRequest) (model.IdentitySpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, updateIdentitySpecsNumber,
		arg.Uuid,
		arg.Iiid,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.IdentitySpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const deleteIdentitySpecsNumber = `-- name: DeleteIdentitySpecsNumber :exec
DELETE FROM Identity_Specs_Number
WHERE uuid = $1
`

func (q *QueriesIdentity) DeleteIdentitySpecsNumber(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteIdentitySpecsNumber, uuid)
	return err
}

type IdentitySpecsNumberInfo struct {
	Uuid            uuid.UUID       `json:"uuid"`
	Iiid            int64           `json:"iiid"`
	SpecsCode       int64           `json:"specsCode"`
	SpecsId         int64           `json:"specsId"`
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

func populateIdentitySpecNumber(q *QueriesIdentity, ctx context.Context, sql string) (IdentitySpecsNumberInfo, error) {
	var i IdentitySpecsNumberInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
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

func populateIdentitySpecNumber2(q *QueriesIdentity, ctx context.Context, sql string) ([]IdentitySpecsNumberInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []IdentitySpecsNumberInfo{}
	for rows.Next() {
		var i IdentitySpecsNumberInfo
		err := rows.Scan(
			&i.Uuid,
			&i.Iiid,
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

const identitySpecsNumberSQL = `-- name: identitySpecsNumberSQL
SELECT 
  mr.UUID, d.Iiid, d.Specs_Code, d.Specs_ID, ref.Title Item, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2, d.Measure_Id, mea.Title Measure, mea.Short_Name Measure_Unit, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Identity_Specs_Number d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
LEFT JOIN Reference mea on mea.ID = Measure_Id`

func (q *QueriesIdentity) GetIdentitySpecsNumber(ctx context.Context, iiid int64, specsId int64) (IdentitySpecsNumberInfo, error) {
	return populateIdentitySpecNumber(q, ctx, fmt.Sprintf("%s WHERE d.IIID = %v and d.Specs_ID = %v",
		identitySpecsNumberSQL, iiid, specsId))
}

func (q *QueriesIdentity) GetIdentitySpecsNumberbyUuid(ctx context.Context, uuid uuid.UUID) (IdentitySpecsNumberInfo, error) {
	return populateIdentitySpecNumber(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", identitySpecsNumberSQL, uuid))
}

type ListIdentitySpecsNumberParams struct {
	Iiid   int64 `json:"Iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListIdentitySpecsNumber(ctx context.Context, arg ListIdentitySpecsNumberParams) ([]IdentitySpecsNumberInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Iiid = %v LIMIT %d OFFSET %d",
			identitySpecsNumberSQL, arg.Iiid, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.IIID = %v ", identitySpecsNumberSQL, arg.Iiid)
	}
	return populateIdentitySpecNumber2(q, ctx, sql)
}
