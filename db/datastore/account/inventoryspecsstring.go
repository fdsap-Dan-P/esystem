package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createInventorySpecsString = `-- name: CreateInventorySpecsString: one
INSERT INTO Inventory_Specs_String 
  (Inventory_Item_ID, Specs_ID, Value) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Inventory_Item_ID, Specs_ID ) DO UPDATE SET
  Value = excluded.Value
RETURNING 
  UUID, Inventory_Item_ID, Specs_Code, Specs_ID, Value
`

type InventorySpecsStringRequest struct {
	Uuid            uuid.UUID `json:"uuid"`
	InventoryItemId int64     `json:"inventoryItemId"`
	SpecsId         int64     `json:"specsId"`
	Value           string    `json:"value"`
}

func (q *QueriesAccount) CreateInventorySpecsString(ctx context.Context, arg InventorySpecsStringRequest) (model.InventorySpecsString, error) {
	row := q.db.QueryRowContext(ctx, createInventorySpecsString,
		arg.InventoryItemId,
		arg.SpecsId,
		arg.Value,
	)
	var i model.InventorySpecsString
	err := row.Scan(
		&i.Uuid,
		&i.InventoryItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const updateInventorySpecsString = `-- name: UpdateInventorySpecsString :one
UPDATE Inventory_Specs_String SET 
Inventory_Item_ID = $2,
  Specs_ID = $3,
  Value = $4
WHERE uuid = $1
RETURNING UUID, Inventory_Item_ID, Specs_Code, Specs_ID, Value
`

func (q *QueriesAccount) UpdateInventorySpecsString(ctx context.Context, arg InventorySpecsStringRequest) (model.InventorySpecsString, error) {
	row := q.db.QueryRowContext(ctx, updateInventorySpecsString,
		arg.Uuid,
		arg.InventoryItemId,
		arg.SpecsId,
		arg.Value,
	)
	var i model.InventorySpecsString
	err := row.Scan(
		&i.Uuid,
		&i.InventoryItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const deleteInventorySpecsString = `-- name: DeleteInventorySpecsString :exec
DELETE FROM Inventory_Specs_String
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteInventorySpecsString(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteInventorySpecsString, uuid)
	return err
}

type InventorySpecsStringInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	InventoryItemId int64        `json:"inventoryItemId"`
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

func populateInventorySpecString(q *QueriesAccount, ctx context.Context, sql string) (InventorySpecsStringInfo, error) {
	var i InventorySpecsStringInfo
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

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateInventorySpecString2(q *QueriesAccount, ctx context.Context, sql string) ([]InventorySpecsStringInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []InventorySpecsStringInfo{}
	for rows.Next() {
		var i InventorySpecsStringInfo
		err := rows.Scan(
			&i.Uuid,
			&i.InventoryItemId,
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

const inventorySpecsStringSQL = `-- name: inventorySpecsStringSQL
SELECT 
  mr.UUID, d.Inventory_Item_ID, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name Item_Short_Name, 
  ref.Remark Item_Description, d.Value,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Inventory_Specs_String d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesAccount) GetInventorySpecsString(ctx context.Context, inventoryItemId int64, specsId int64) (InventorySpecsStringInfo, error) {
	return populateInventorySpecString(q, ctx, fmt.Sprintf("%s WHERE d.Inventory_Item_ID = %v and d.Specs_ID = %v",
		inventorySpecsStringSQL, inventoryItemId, specsId))
}

func (q *QueriesAccount) GetInventorySpecsStringbyUuid(ctx context.Context, uuid uuid.UUID) (InventorySpecsStringInfo, error) {
	return populateInventorySpecString(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", inventorySpecsStringSQL, uuid))
}

type ListInventorySpecsStringParams struct {
	InventoryItemId int64 `json:"InventoryItemId"`
	Limit           int32 `json:"limit"`
	Offset          int32 `json:"offset"`
}

func (q *QueriesAccount) ListInventorySpecsString(ctx context.Context, arg ListInventorySpecsStringParams) ([]InventorySpecsStringInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Inventory_Item_ID = %v LIMIT %d OFFSET %d",
			inventorySpecsStringSQL, arg.InventoryItemId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Inventory_Item_ID = %v ", inventorySpecsStringSQL, arg.InventoryItemId)
	}
	return populateInventorySpecString2(q, ctx, sql)
}
