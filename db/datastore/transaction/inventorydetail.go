package db

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"simplebank/model"

// 	"github.com/google/uuid"
// 	"github.com/shopspring/decimal"
// )

// const createInventoryDetail = `-- name: CreateInventoryDetail: one
// INSERT INTO Inventory_Detail (
//    uuid, account_inventory_id, inventory_item_id, repository_id, supplier_id, unit_price,
//    Book_Value, Unit, Measure_ID, Batch_Number, Date_Manufactured, Date_Expired,
//    remarks, other_info )
// VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
// ON CONFLICT(UUID)
// DO UPDATE SET
// 	account_inventory_id =  EXCLUDED.account_inventory_id,
// 	inventory_item_id =  EXCLUDED.inventory_item_id,
// 	repository_id =  EXCLUDED.repository_id,
// 	supplier_id =  EXCLUDED.supplier_id,
// 	unit_price =  EXCLUDED.unit_price,
// 	book_value =  EXCLUDED.book_value,
// 	unit =  EXCLUDED.unit,
// 	measure_id =  EXCLUDED.measure_id,
// 	batch_number =  EXCLUDED.batch_number,
// 	date_manufactured =  EXCLUDED.date_manufactured,
// 	date_expired =  EXCLUDED.date_expired,
// 	remarks =  EXCLUDED.remarks,
// 	other_info =  EXCLUDED.other_info

// RETURNING
//   id, uuid, account_inventory_id, inventory_item_id, repository_id, supplier_id, unit_price,
//   book_value, unit, measure_id, batch_number, date_manufactured, date_expired, remarks, other_info`

// type InventoryDetailRequest struct {
// 	Id                 int64           `json:"id"`
// 	Uuid               uuid.UUID       `json:"uuid"`
// 	AccountInventoryId int64           `json:"accountInventoryId"`
// 	InventoryItemId    int64           `json:"inventoryItemId"`
// 	RepositoryId       sql.NullInt64   `json:"repositoryId"`
// 	SupplierId         sql.NullInt64   `json:"supplierId"`
// 	UnitPrice          decimal.Decimal `json:"unitPrice"`
// 	BookValue          decimal.Decimal `json:"bookValue"`
// 	Unit               decimal.Decimal `json:"unit"`
// 	MeasureId          int64           `json:"measureId"`
// 	BatchNumber        sql.NullString  `json:"batchNumber"`
// 	DateManufactured   sql.NullTime    `json:"dateManufactured"`
// 	DateExpired        sql.NullTime    `json:"dateExpired"`
// 	Remarks            string          `json:"remarks"`
// 	OtherInfo          sql.NullString  `json:"otherInfo"`
// }

// type InventoryDetailFullRequest struct {
// 	Id                 int64                `json:"id"`
// 	Uuid               uuid.UUID            `json:"uuid"`
// 	AccountInventoryId int64                `json:"accountInventoryId"`
// 	InventoryItem      InventoryItemRequest `json:"inventoryItem"`
// 	SupplierId         sql.NullInt64        `json:"supplierId"`
// 	UnitPrice          decimal.Decimal      `json:"unitPrice"`
// 	BookValue          decimal.Decimal      `json:"bookValue"`
// 	Unit               decimal.Decimal      `json:"unit"`
// 	MeasureId          int64                `json:"measureId"`
// 	BatchNumber        sql.NullString       `json:"batchNumber"`
// 	DateManufactured   sql.NullTime         `json:"dateManufactured"`
// 	DateExpired        sql.NullTime         `json:"dateExpired"`
// 	Remarks            string               `json:"remarks"`
// 	OtherInfo          sql.NullString       `json:"otherInfo"`
// }

// func (q *QueriesTransaction) CreateInventoryDetail(ctx context.Context, arg InventoryDetailRequest) (model.InventoryDetail, error) {
// 	row := q.db.QueryRowContext(ctx, createInventoryDetail,
// 		arg.Uuid,
// 		arg.AccountInventoryId,
// 		arg.InventoryItemId,
// 		arg.RepositoryId,
// 		arg.SupplierId,
// 		arg.UnitPrice,
// 		arg.BookValue,
// 		arg.Unit,
// 		arg.MeasureId,
// 		arg.BatchNumber,
// 		arg.DateManufactured,
// 		arg.DateExpired,
// 		arg.Remarks,
// 		arg.OtherInfo,
// 	)
// 	var i model.InventoryDetail
// 	err := row.Scan(
// 		&i.Id,
// 		&i.Uuid,
// 		&i.AccountInventoryId,
// 		&i.InventoryItemId,
// 		&i.RepositoryId,
// 		&i.SupplierId,
// 		&i.UnitPrice,
// 		&i.BookValue,
// 		&i.Unit,
// 		&i.MeasureId,
// 		&i.BatchNumber,
// 		&i.DateManufactured,
// 		&i.DateExpired,
// 		&i.Remarks,
// 		&i.OtherInfo,
// 	)
// 	return i, err
// }

// const deleteInventoryDetail = `-- name: DeleteInventoryDetail :exec
// DELETE FROM Inventory_Detail
// WHERE uuid = $1
// `

// func (q *QueriesTransaction) DeleteInventoryDetail(ctx context.Context, uuid uuid.UUID) error {
// 	_, err := q.db.ExecContext(ctx, deleteInventoryDetail, uuid)
// 	return err
// }

// type InventoryDetailInfo struct {
// 	Id                 int64           `json:"id"`
// 	Uuid               uuid.UUID       `json:"uuid"`
// 	AccountInventoryId int64           `json:"accountInventoryId"`
// 	InventoryItemId    int64           `json:"inventoryItemId"`
// 	RepositoryId       sql.NullInt64   `json:"repositoryId"`
// 	SupplierId         sql.NullInt64   `json:"supplierId"`
// 	UnitPrice          decimal.Decimal `json:"unitPrice"`
// 	BookValue          decimal.Decimal `json:"bookValue"`
// 	Unit               decimal.Decimal `json:"unit"`
// 	MeasureId          int64           `json:"measureId"`
// 	BatchNumber        sql.NullString  `json:"batchNumber"`
// 	DateManufactured   sql.NullTime    `json:"dateManufactured"`
// 	DateExpired        sql.NullTime    `json:"dateExpired"`
// 	Remarks            string          `json:"remarks"`
// 	OtherInfo          sql.NullString  `json:"otherInfo"`
// 	ModCtr             int64           `json:"modCtr"`
// 	Created            sql.NullTime    `json:"created"`
// 	Updated            sql.NullTime    `json:"updated"`
// }

// const inventoryDetailSQL = `-- name: InventoryDetailSQL :one
// SELECT
//   Id, mr.UUID, Account_Inventory_Id, Inventory_Item_Id, Repository_Id, Supplier_Id,
//   Unit_Price, Book_Value, Unit, Measure_ID, Batch_Number, Date_Manufactured, Date_Expired,
//   Remarks, Other_Info
// ,mr.Mod_Ctr, mr.Created, mr.Updated
// FROM Inventory_Detail d INNER JOIN Main_Record mr on mr.UUID = d.UUID
// `

// func populateInventoryDetail(q *QueriesTransaction, ctx context.Context, sql string) (InventoryDetailInfo, error) {
// 	var i InventoryDetailInfo
// 	row := q.db.QueryRowContext(ctx, sql)
// 	err := row.Scan(
// 		&i.Id,
// 		&i.Uuid,
// 		&i.AccountInventoryId,
// 		&i.InventoryItemId,
// 		&i.RepositoryId,
// 		&i.SupplierId,
// 		&i.UnitPrice,
// 		&i.BookValue,
// 		&i.Unit,
// 		&i.MeasureId,
// 		&i.BatchNumber,
// 		&i.DateManufactured,
// 		&i.DateExpired,
// 		&i.Remarks,
// 		&i.OtherInfo,
// 		&i.ModCtr,
// 		&i.Created,
// 		&i.Updated,
// 	)
// 	return i, err
// }

// func populateInventoryDetails(q *QueriesTransaction, ctx context.Context, sql string) ([]InventoryDetailInfo, error) {
// 	rows, err := q.db.QueryContext(ctx, sql)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	items := []InventoryDetailInfo{}
// 	for rows.Next() {
// 		var i InventoryDetailInfo
// 		err := rows.Scan(
// 			&i.Id,
// 			&i.Uuid,
// 			&i.AccountInventoryId,
// 			&i.InventoryItemId,
// 			&i.RepositoryId,
// 			&i.SupplierId,
// 			&i.UnitPrice,
// 			&i.BookValue,
// 			&i.Unit,
// 			&i.MeasureId,
// 			&i.BatchNumber,
// 			&i.DateManufactured,
// 			&i.DateExpired,
// 			&i.Remarks,
// 			&i.OtherInfo,
// 			&i.ModCtr,
// 			&i.Created,
// 			&i.Updated,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// func (q *QueriesTransaction) GetInventoryDetail(ctx context.Context, id int64) (InventoryDetailInfo, error) {
// 	return populateInventoryDetail(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", inventoryDetailSQL, id))
// }

// func (q *QueriesTransaction) GetInventoryDetailbyUuid(ctx context.Context, uuid uuid.UUID) (InventoryDetailInfo, error) {
// 	return populateInventoryDetail(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", inventoryDetailSQL, uuid))
// }

// type ListInventoryDetailParams struct {
// 	Limit  int32 `json:"limit"`
// 	Offset int32 `json:"offset"`
// }

// func (q *QueriesTransaction) ListInventoryDetail(ctx context.Context, arg ListInventoryDetailParams) ([]InventoryDetailInfo, error) {
// 	var sql string
// 	if arg.Limit != 0 {
// 		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
// 			inventoryDetailSQL, arg.Limit, arg.Offset)
// 	} else {
// 		sql = inventoryDetailSQL
// 	}
// 	return populateInventoryDetails(q, ctx, sql)
// }

// const updateInventoryDetail = `-- name: UpdateInventoryDetail :one
// UPDATE Inventory_Detail SET
// 	uuid = $2,
// 	account_inventory_id = $3,
// 	inventory_item_id = $4,
// 	repository_id = $5,
// 	supplier_id = $6,
// 	unit_price = $7,
// 	book_value = $8,
// 	unit = $9,
// 	measure_id = $10,
// 	batch_number = $11,
// 	date_manufactured = $12,
// 	date_expired = $13,
// 	remarks = $14,
// 	other_info = $15
// WHERE id = $1
// RETURNING
//   id, uuid, account_inventory_id, inventory_item_id, repository_id,
//   supplier_id, unit_price, book_value, unit, measure_id, batch_number, date_manufactured, date_expired, remarks, other_info
// `

// func (q *QueriesTransaction) UpdateInventoryDetail(ctx context.Context, arg InventoryDetailRequest) (model.InventoryDetail, error) {
// 	row := q.db.QueryRowContext(ctx, updateInventoryDetail,
// 		arg.Id,
// 		arg.Uuid,
// 		arg.AccountInventoryId,
// 		arg.InventoryItemId,
// 		arg.RepositoryId,
// 		arg.SupplierId,
// 		arg.UnitPrice,
// 		arg.BookValue,
// 		arg.Unit,
// 		arg.MeasureId,
// 		arg.BatchNumber,
// 		arg.DateManufactured,
// 		arg.DateExpired,
// 		arg.Remarks,
// 		arg.OtherInfo,
// 	)
// 	var i model.InventoryDetail
// 	err := row.Scan(
// 		&i.Id,
// 		&i.Uuid,
// 		&i.AccountInventoryId,
// 		&i.InventoryItemId,
// 		&i.RepositoryId,
// 		&i.SupplierId,
// 		&i.UnitPrice,
// 		&i.BookValue,
// 		&i.Unit,
// 		&i.MeasureId,
// 		&i.BatchNumber,
// 		&i.DateManufactured,
// 		&i.DateExpired,
// 		&i.Remarks,
// 		&i.OtherInfo,
// 	)
// 	return i, err
// }

// func (store *QueriesTransaction) CreateInventoryDetailFull(ctx context.Context, arg InventoryDetailFullRequest) (model.InventoryDetail, error) {
// 	var result model.InventoryDetail

// 	err := store.ExecTx(ctx, func(q *QueriesTransaction) error {
// 		var err error

// 		InvItem, r := store.CreateInventoryItemFull(context.Background(), arg.InventoryItem)
// 		if r != nil {
// 			fmt.Printf("Error CreateInventoryItemFull: %+v\n", r)
// 			return r
// 		}

// 		row := q.db.QueryRowContext(ctx, createInventoryDetail,
// 			arg.AccountInventoryId,
// 			InvItem.Id,
// 			arg.SupplierId,
// 			arg.UnitPrice,
// 			arg.BookValue,
// 			arg.Unit,
// 			arg.MeasureId,
// 			arg.BatchNumber,
// 			arg.DateManufactured,
// 			arg.DateExpired,
// 			arg.Remarks,
// 			arg.OtherInfo,
// 		)

// 		err = row.Scan(
// 			&result.Id,
// 			&result.Uuid,
// 			&result.AccountInventoryId,
// 			&result.InventoryItemId,
// 			&result.SupplierId,
// 			&result.UnitPrice,
// 			&result.BookValue,
// 			&result.Unit,
// 			&result.MeasureId,
// 			&result.BatchNumber,
// 			&result.DateManufactured,
// 			&result.DateExpired,
// 			&result.Remarks,
// 			&result.OtherInfo,
// 		)

// 		fmt.Printf("InventoryItem: %+v\n", arg.InventoryItem)
// 		fmt.Printf("InvItem:err: %+v\n", err)

// 		return err
// 	})

// 	return result, err
// }
