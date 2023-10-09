package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createInventorySpecsRef = `-- name: CreateInventorySpecsRef: one
INSERT INTO Inventory_Specs_Ref 
  (Inventory_Item_Id, Specs_ID, Ref_Id) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Inventory_Item_Id, Specs_ID ) DO UPDATE SET
Ref_Id = excluded.Ref_Id
RETURNING 
  UUID, Inventory_Item_Id, Specs_Code, Specs_ID, Ref_Id
`

type InventorySpecsRefRequest struct {
	Uuid            uuid.UUID `json:"uuid"`
	InventoryItemId int64     `json:"inventoryItemId"`
	SpecsId         int64     `json:"specsId"`
	RefId           int64     `json:"refId"`
}

func (q *QueriesAccount) CreateInventorySpecsRef(ctx context.Context, arg InventorySpecsRefRequest) (model.InventorySpecsRef, error) {
	row := q.db.QueryRowContext(ctx, createInventorySpecsRef,
		arg.InventoryItemId,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.InventorySpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.InventoryItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const updateInventorySpecsRef = `-- name: UpdateInventorySpecsRef :one
UPDATE Inventory_Specs_Ref SET 
  Inventory_Item_Id = $2,
  Specs_ID = $3,
  Ref_ID = $4
WHERE uuid = $1
RETURNING UUID, Inventory_Item_Id, Specs_Code, Specs_ID, Ref_ID
`

func (q *QueriesAccount) UpdateInventorySpecsRef(ctx context.Context, arg InventorySpecsRefRequest) (model.InventorySpecsRef, error) {
	row := q.db.QueryRowContext(ctx, updateInventorySpecsRef,
		arg.Uuid,
		arg.InventoryItemId,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.InventorySpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.InventoryItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const deleteInventorySpecsRef = `-- name: DeleteInventorySpecsRef :exec
DELETE FROM Inventory_Specs_Ref
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteInventorySpecsRef(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteInventorySpecsRef, uuid)
	return err
}

type InventorySpecsRefInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	InventoryItemId int64          `json:"inventoryItemId"`
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

func populateInventorySpecRef(q *QueriesAccount, ctx context.Context, sql string) (InventorySpecsRefInfo, error) {
	var i InventorySpecsRefInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.InventoryItemId,
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

func populateInventorySpecRef2(q *QueriesAccount, ctx context.Context, sql string) ([]InventorySpecsRefInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []InventorySpecsRefInfo{}
	for rows.Next() {
		var i InventorySpecsRefInfo
		err := rows.Scan(
			&i.Uuid,
			&i.InventoryItemId,
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

const inventorySpecsRefSQL = `-- name: inventorySpecsRefSQL
SELECT 
  mr.UUID, d.Inventory_Item_Id, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, d.Ref_Id,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Inventory_Specs_Ref d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesAccount) GetInventorySpecsRef(ctx context.Context, inventoryItemId int64, specsId int64) (InventorySpecsRefInfo, error) {
	return populateInventorySpecRef(q, ctx, fmt.Sprintf("%s WHERE d.Inventory_Item_Id = %v and d.Specs_ID = %v",
		inventorySpecsRefSQL, inventoryItemId, specsId))
}

func (q *QueriesAccount) GetInventorySpecsRefbyUuid(ctx context.Context, uuid uuid.UUID) (InventorySpecsRefInfo, error) {
	return populateInventorySpecRef(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", inventorySpecsRefSQL, uuid))
}

type ListInventorySpecsRefParams struct {
	InventoryItemId int64 `json:"InventoryItemId"`
	Limit           int32 `json:"limit"`
	Offset          int32 `json:"offset"`
}

func (q *QueriesAccount) ListInventorySpecsRef(ctx context.Context, arg ListInventorySpecsRefParams) ([]InventorySpecsRefInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Inventory_Item_Id = %v LIMIT %d OFFSET %d",
			inventorySpecsRefSQL, arg.InventoryItemId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Inventory_Item_Id = %v ", inventorySpecsRefSQL, arg.InventoryItemId)
	}
	return populateInventorySpecRef2(q, ctx, sql)
}
