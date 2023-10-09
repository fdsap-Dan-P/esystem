package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const createInventoryTran = `-- name: CreateInventoryTran: one
INSERT INTO Inventory_Tran(
  uuid, Trn_Head_Id, Series, inventory_detail_id, repository_id, quantity, unit_price, discount, tax_amt, net_trn_amt, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT(UUID)
DO UPDATE SET 
	trn_head_id =  EXCLUDED.trn_head_id,
	series =  EXCLUDED.series,
	inventory_detail_id =  EXCLUDED.inventory_detail_id,
	repository_id =  EXCLUDED.repository_id,
	quantity =  EXCLUDED.quantity,
	unit_price =  EXCLUDED.unit_price,
	discount =  EXCLUDED.discount,
	tax_amt =  EXCLUDED.tax_amt,
	net_trn_amt =  EXCLUDED.net_trn_amt,
	other_info =  EXCLUDED.other_info
RETURNING 
  UUID, Trn_Head_Id, Series, inventory_detail_id, repository_id, quantity, unit_price, discount, tax_amt, net_trn_amt, other_info`

type InventoryTranRequest struct {
	Uuid              uuid.UUID       `json:"uuid"`
	TrnHeadId         int64           `json:"trnHeadId"`
	Series            int16           `json:"series"`
	InventoryDetailId int64           `json:"inventoryDetailId"`
	RepositoryId      int64           `json:"repositoryId"`
	Quantity          decimal.Decimal `json:"quantity"`
	UnitPrice         decimal.Decimal `json:"unitPrice"`
	Discount          decimal.Decimal `json:"discount"`
	TaxAmt            decimal.Decimal `json:"taxAmt"`
	NetTrnAmt         decimal.Decimal `json:"netTrnAmt"`
	OtherInfo         sql.NullString  `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateInventoryTran(ctx context.Context, arg InventoryTranRequest) (model.InventoryTran, error) {
	row := q.db.QueryRowContext(ctx, createInventoryTran,
		arg.Uuid,
		arg.TrnHeadId,
		arg.Series,
		arg.InventoryDetailId,
		arg.RepositoryId,
		arg.Quantity,
		arg.UnitPrice,
		arg.Discount,
		arg.TaxAmt,
		arg.NetTrnAmt,
		arg.OtherInfo,
	)
	var i model.InventoryTran
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.InventoryDetailId,
		&i.RepositoryId,
		&i.Quantity,
		&i.UnitPrice,
		&i.Discount,
		&i.TaxAmt,
		&i.NetTrnAmt,
		&i.OtherInfo,
	)
	return i, err
}

const deleteInventoryTran = `-- name: DeleteInventoryTran :exec
DELETE FROM Inventory_Tran
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteInventoryTran(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteInventoryTran, uuid)
	return err
}

type InventoryTranInfo struct {
	Uuid              uuid.UUID       `json:"uuid"`
	TrnHeadId         int64           `json:"trnHeadId"`
	Series            int16           `json:"series"`
	InventoryDetailId int64           `json:"inventoryDetailId"`
	RepositoryId      int64           `json:"repositoryId"`
	Quantity          decimal.Decimal `json:"quantity"`
	UnitPrice         decimal.Decimal `json:"unitPrice"`
	Discount          decimal.Decimal `json:"discount"`
	TaxAmt            decimal.Decimal `json:"taxAmt"`
	NetTrnAmt         decimal.Decimal `json:"netTrnAmt"`
	OtherInfo         sql.NullString  `json:"otherInfo"`
	ModCtr            int64           `json:"modCtr"`
	Created           sql.NullTime    `json:"created"`
	Updated           sql.NullTime    `json:"updated"`
}

const inventoryTranSQL = `-- name: InventoryTranSQL :one
SELECT
  mr.UUID, Trn_Head_Id, Series, inventory_detail_id, repository_id, quantity, unit_price, discount, tax_amt, net_trn_amt, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Inventory_Tran d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateInventoryTran(q *QueriesTransaction, ctx context.Context, sql string) (InventoryTranInfo, error) {
	var i InventoryTranInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.InventoryDetailId,
		&i.RepositoryId,
		&i.Quantity,
		&i.UnitPrice,
		&i.Discount,
		&i.TaxAmt,
		&i.NetTrnAmt,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateInventoryTrans(q *QueriesTransaction, ctx context.Context, sql string) ([]InventoryTranInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []InventoryTranInfo{}
	for rows.Next() {
		var i InventoryTranInfo
		err := rows.Scan(
			&i.Uuid,
			&i.TrnHeadId,
			&i.Series,
			&i.InventoryDetailId,
			&i.RepositoryId,
			&i.Quantity,
			&i.UnitPrice,
			&i.Discount,
			&i.TaxAmt,
			&i.NetTrnAmt,
			&i.OtherInfo,
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

// func (q *QueriesTransaction) GetInventoryTran(ctx context.Context, id int64) (InventoryTranInfo, error) {
// 	return populateInventoryTran(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", inventoryTranSQL, id))
// }

func (q *QueriesTransaction) GetInventoryTranbyUuid(ctx context.Context, uuid uuid.UUID) (InventoryTranInfo, error) {
	return populateInventoryTran(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", inventoryTranSQL, uuid))
}

type ListInventoryTranParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesTransaction) ListInventoryTran(ctx context.Context, arg ListInventoryTranParams) ([]InventoryTranInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			inventoryTranSQL, arg.Limit, arg.Offset)
	} else {
		sql = inventoryTranSQL
	}
	return populateInventoryTrans(q, ctx, sql)
}

const updateInventoryTran = `-- name: UpdateInventoryTran :one
UPDATE Inventory_Tran SET 
    trn_head_id = $2,
	series = $3,
	inventory_detail_id = $4,
	repository_id = $5,
	quantity = $6,
	unit_price = $7,
	discount = $8,
	tax_amt = $9,
	net_trn_amt = $10,
	other_info = $11
WHERE uuid = $1
RETURNING Uuid, trn_head_id, series, inventory_detail_id, repository_id, quantity, unit_price, discount, tax_amt, net_trn_amt, other_info
`

func (q *QueriesTransaction) UpdateInventoryTran(ctx context.Context, arg InventoryTranRequest) (model.InventoryTran, error) {
	row := q.db.QueryRowContext(ctx, updateInventoryTran,
		arg.Uuid,
		arg.TrnHeadId,
		arg.Series,
		arg.InventoryDetailId,
		arg.RepositoryId,
		arg.Quantity,
		arg.UnitPrice,
		arg.Discount,
		arg.TaxAmt,
		arg.NetTrnAmt,
		arg.OtherInfo,
	)
	var i model.InventoryTran
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.InventoryDetailId,
		&i.RepositoryId,
		&i.Quantity,
		&i.UnitPrice,
		&i.Discount,
		&i.TaxAmt,
		&i.NetTrnAmt,
		&i.OtherInfo,
	)
	return i, err
}
