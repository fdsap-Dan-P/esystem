package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createInventorySpecsNumber = `-- name: CreateInventorySpecsNumber: one
INSERT INTO Inventory_Specs_Number 
  (Inventory_Item_ID, Specs_ID, Value, Value2, Measure_Id) 
VALUES 
  ($1, $2, $3, $4, $5) 
ON CONFLICT( Inventory_Item_ID, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Inventory_Item_ID, Specs_Code, Specs_ID, Value, Value2, Measure_Id
`

type InventorySpecsNumberRequest struct {
	Uuid            uuid.UUID       `json:"uuid"`
	InventoryItemId int64           `json:"inventoryItemId"`
	SpecsId         int64           `json:"specsId"`
	Value           decimal.Decimal `json:"value"`
	Value2          decimal.Decimal `json:"value2"`
	MeasureId       sql.NullInt64   `json:"measureId"`
}

func (q *QueriesAccount) CreateInventorySpecsNumber(ctx context.Context, arg InventorySpecsNumberRequest) (model.InventorySpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, createInventorySpecsNumber,
		arg.InventoryItemId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.InventorySpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.InventoryItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const updateInventorySpecsNumber = `-- name: UpdateInventorySpecsNumber :one
UPDATE Inventory_Specs_Number SET 
  Inventory_Item_ID = $2,
  Specs_ID = $3,
  Value = $4,
  Value2 = $5,
  Measure_Id = $6
WHERE uuid = $1
RETURNING UUID, Inventory_Item_ID, Specs_Code, Specs_ID, Value, Value2, Measure_Id
`

func (q *QueriesAccount) UpdateInventorySpecsNumber(ctx context.Context, arg InventorySpecsNumberRequest) (model.InventorySpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, updateInventorySpecsNumber,
		arg.Uuid,
		arg.InventoryItemId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.InventorySpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.InventoryItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const deleteInventorySpecsNumber = `-- name: DeleteInventorySpecsNumber :exec
DELETE FROM Inventory_Specs_Number
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteInventorySpecsNumber(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteInventorySpecsNumber, uuid)
	return err
}

type InventorySpecsNumberInfo struct {
	Uuid            uuid.UUID       `json:"uuid"`
	InventoryItemId int64           `json:"inventoryItemId"`
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

func populateInventorySpecNumber(q *QueriesAccount, ctx context.Context, sql string) (InventorySpecsNumberInfo, error) {
	var i InventorySpecsNumberInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.InventoryItemId,
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

func populateInventorySpecNumber2(q *QueriesAccount, ctx context.Context, sql string) ([]InventorySpecsNumberInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []InventorySpecsNumberInfo{}
	for rows.Next() {
		var i InventorySpecsNumberInfo
		err := rows.Scan(
			&i.Uuid,
			&i.InventoryItemId,
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

const inventorySpecsNumberSQL = `-- name: inventorySpecsNumberSQL
SELECT 
  mr.UUID, d.Inventory_Item_ID, d.Specs_Code, ref.id Specs_Id, ref.Title Item, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2, d.Measure_Id, mea.Title Measure, mea.Short_Name Measure_Unit, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Inventory_Specs_Number d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
LEFT JOIN Reference mea on mea.ID = Measure_Id`

func (q *QueriesAccount) GetInventorySpecsNumber(ctx context.Context, inventoryItemId int64, specsId int64) (InventorySpecsNumberInfo, error) {
	return populateInventorySpecNumber(q, ctx, fmt.Sprintf("%s WHERE d.Inventory_Item_ID = %v and d.Specs_ID = %v",
		inventorySpecsNumberSQL, inventoryItemId, specsId))
}

func (q *QueriesAccount) GetInventorySpecsNumberbyUuid(ctx context.Context, uuid uuid.UUID) (InventorySpecsNumberInfo, error) {
	return populateInventorySpecNumber(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", inventorySpecsNumberSQL, uuid))
}

type ListInventorySpecsNumberParams struct {
	InventoryItemId int64 `json:"InventoryItemId"`
	Limit           int32 `json:"limit"`
	Offset          int32 `json:"offset"`
}

func (q *QueriesAccount) ListInventorySpecsNumber(ctx context.Context, arg ListInventorySpecsNumberParams) ([]InventorySpecsNumberInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Inventory_Item_ID = %v LIMIT %d OFFSET %d",
			inventorySpecsNumberSQL, arg.InventoryItemId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Inventory_Item_ID = %v ", inventorySpecsNumberSQL, arg.InventoryItemId)
	}
	return populateInventorySpecNumber2(q, ctx, sql)
}
