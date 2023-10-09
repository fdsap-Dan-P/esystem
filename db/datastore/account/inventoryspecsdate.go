package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
)

const createInventorySpecsDate = `-- name: CreateInventorySpecsDate: one
INSERT INTO Inventory_Specs_Date (
Inventory_Item_ID, Specs_ID, Value, Value2
) VALUES (
$1, $2, $3, $4
) 
ON CONFLICT(Inventory_Item_ID, Specs_ID) DO UPDATE SET
Value = excluded.Value,
Value2 = excluded.Value2
RETURNING UUID, Inventory_Item_ID, Specs_Code, Specs_ID, Value, Value2
`

type InventorySpecsDateRequest struct {
	Uuid            uuid.UUID `json:"uuid"`
	InventoryItemId int64     `json:"inventory_Item_id"`
	SpecsId         int64     `json:"specsId"`
	Value           time.Time `json:"value"`
	Value2          time.Time `json:"value2"`
}

func (q *QueriesAccount) CreateInventorySpecsDate(ctx context.Context, arg InventorySpecsDateRequest) (model.InventorySpecsDate, error) {
	row := q.db.QueryRowContext(ctx, createInventorySpecsDate,
		arg.InventoryItemId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.InventorySpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.InventoryItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const deleteInventorySpecsDate = `-- name: DeleteInventorySpecsDate :exec
DELETE FROM Inventory_Specs_Date
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteInventorySpecsDate(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteInventorySpecsDate, uuid)
	return err
}

type InventorySpecsDateInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	InventoryItemId int64        `json:"inventory_Item_id"`
	SpecsCode       string       `json:"specsCode"`
	SpecsId         int64        `json:"specsId"`
	Item            string       `json:"item"`
	ItemDescription string       `json:"itemDescription"`
	Value           time.Time    `json:"value"`
	Value2          time.Time    `json:"value2"`
	ModCtr          int64        `json:"mod_ctr"`
	Created         sql.NullTime `json:"created"`
	Updated         sql.NullTime `json:"updated"`
}

func populateInventorySpecs(q *QueriesAccount, ctx context.Context, sql string) (InventorySpecsDateInfo, error) {
	var i InventorySpecsDateInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.InventoryItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Item,
		&i.ItemDescription,
		&i.Value,
		&i.Value2,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateInventorySpecs2(q *QueriesAccount, ctx context.Context, sql string) ([]InventorySpecsDateInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []InventorySpecsDateInfo{}
	for rows.Next() {
		var i InventorySpecsDateInfo
		err := rows.Scan(
			&i.Uuid,
			&i.InventoryItemId,
			&i.SpecsCode,
			&i.SpecsId,
			&i.Item,
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

const inventorySpecsSQLDate = `-- name: inventorySpecsSQLDate
SELECT 
  mr.UUID, d.Inventory_Item_ID, d.Specs_Code, d.Specs_ID, ref.Title, ref.Remark ItemDescription, 
  d.Value, d.Value2, mr.Mod_Ctr, mr.Created, mr.Updated
FROM Inventory_Specs_Date d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesAccount) GetInventorySpecsDate(ctx context.Context, inventoryItemId int64, specsId int64) (InventorySpecsDateInfo, error) {
	return populateInventorySpecs(q, ctx, fmt.Sprintf("%s WHERE d.Inventory_Item_ID = %v and d.Specs_ID = %v",
		inventorySpecsSQLDate, inventoryItemId, specsId))
}

func (q *QueriesAccount) GetInventorySpecsDatebyUuid(ctx context.Context, uuid uuid.UUID) (InventorySpecsDateInfo, error) {
	return populateInventorySpecs(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'",
		inventorySpecsSQLDate, uuid))
}

type ListInventorySpecsDateParams struct {
	InventoryItemId int64 `json:"inventoryItemId"`
	Limit           int32 `json:"limit"`
	Offset          int32 `json:"offset"`
}

func (q *QueriesAccount) ListInventorySpecsDate(ctx context.Context, arg ListInventorySpecsDateParams) ([]InventorySpecsDateInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Inventory_Item_ID = %v LIMIT %d OFFSET %d",
			inventorySpecsSQLDate, arg.InventoryItemId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Inventory_Item_ID = %v ", inventorySpecsSQLDate, arg.InventoryItemId)
	}
	return populateInventorySpecs2(q, ctx, sql)
}

const updateInventorySpecsDate = `-- name: UpdateInventorySpecsDate :one
UPDATE Inventory_Specs_Date SET 
Inventory_Item_ID = $2,
Specs_ID = $3,
Value = $4,
Value2 = $5
WHERE uuid = $1
RETURNING UUID, Inventory_Item_ID, Specs_Code, Specs_ID, Value, Value2
`

func (q *QueriesAccount) UpdateInventorySpecsDate(ctx context.Context, arg InventorySpecsDateRequest) (model.InventorySpecsDate, error) {
	row := q.db.QueryRowContext(ctx, updateInventorySpecsDate,

		arg.Uuid,
		arg.InventoryItemId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.InventorySpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.InventoryItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}
