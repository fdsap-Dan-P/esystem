package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createAccountInventory = `-- name: CreateAccountInventory: one
INSERT INTO Account_Inventory (
  Uuid, Account_ID, Bar_Code, Code, Quantity, Unit_Price, 
  Book_Value, Discount, Tax_Rate, Remarks, Other_Info) 
VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11 ) 
ON CONFLICT(UUID)
DO UPDATE SET
	account_id =  EXCLUDED.account_id,
	repository_id =  EXCLUDED.repository_id,
	bar_code =  EXCLUDED.bar_code,
	code =  EXCLUDED.code,
	quantity =  EXCLUDED.quantity,
	unit_price =  EXCLUDED.unit_price,
	book_value =  EXCLUDED.book_value,
	discount =  EXCLUDED.discount,
	tax_rate =  EXCLUDED.tax_rate,
	remarks =  EXCLUDED.remarks,
	other_info =  EXCLUDED.other_info
RETURNING 
  Id, UUID, Account_ID, Bar_Code, Code, Quantity, Unit_Price, 
   Book_Value, Discount, Tax_Rate, Remarks, Other_Info 
`

type AccountInventoryRequest struct {
	Id        int64           `json:"id"`
	Uuid      uuid.UUID       `json:"uuid"`
	AccountId int64           `json:"accountId"`
	BarCode   sql.NullString  `json:"packageId"`
	Code      string          `json:"itemName"`
	Quantity  decimal.Decimal `json:"quantity"`
	UnitPrice decimal.Decimal `json:"unitPrice"`
	BookValue decimal.Decimal `json:"bookValue"`
	Discount  decimal.Decimal `json:"discount"`
	TaxRate   decimal.Decimal `json:"taxRate"`
	Remarks   string          `json:"remarks"`
	OtherInfo sql.NullString  `json:"otherInfo"`
}

func (q *QueriesAccount) CreateAccountInventory(ctx context.Context, arg AccountInventoryRequest) (model.AccountInventory, error) {
	row := q.db.QueryRowContext(ctx, createAccountInventory,
		arg.Uuid,
		arg.AccountId,
		arg.BarCode,
		arg.Code,
		arg.Quantity,
		arg.UnitPrice,
		arg.BookValue,
		arg.Discount,
		arg.TaxRate,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.AccountInventory
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountId,
		&i.BarCode,
		&i.Code,
		&i.Quantity,
		&i.UnitPrice,
		&i.BookValue,
		&i.Discount,
		&i.TaxRate,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccountInventory = `-- name: DeleteAccountInventory :exec
DELETE FROM Account_Inventory
WHERE id = $1
`

func (q *QueriesAccount) DeleteAccountInventory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccountInventory, id)
	return err
}

type AccountInventoryInfo struct {
	Id        int64           `json:"id"`
	Uuid      uuid.UUID       `json:"uuid"`
	AccountId int64           `json:"accountId"`
	BarCode   sql.NullString  `json:"packageId"`
	Code      string          `json:"itemName"`
	Quantity  decimal.Decimal `json:"quantity"`
	UnitPrice decimal.Decimal `json:"unitPrice"`
	BookValue decimal.Decimal `json:"bookValue"`
	Discount  decimal.Decimal `json:"discount"`
	TaxRate   decimal.Decimal `json:"taxRate"`
	Remarks   string          `json:"remarks"`
	OtherInfo sql.NullString  `json:"otherInfo"`
	ModCtr    int64           `json:"modCtr"`
	Created   sql.NullTime    `json:"created"`
	Updated   sql.NullTime    `json:"updated"`
}

const getAccountInventory = `-- name: GetAccountInventory :one
SELECT 
Id, mr.UUID, Account_ID, Bar_Code, Code, Quantity, Unit_Price, 
Book_Value, Discount, Tax_Rate, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Inventory d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesAccount) GetAccountInventory(ctx context.Context, id int64) (AccountInventoryInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccountInventory, id)
	var i AccountInventoryInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountId,
		&i.BarCode,
		&i.Code,
		&i.Quantity,
		&i.UnitPrice,
		&i.BookValue,
		&i.Discount,
		&i.TaxRate,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAccountInventorybyUuid = `-- name: GetAccountInventorybyUuid :one
SELECT 
Id, mr.UUID, Account_ID, Bar_Code, Code, Quantity, Unit_Price, 
Book_Value, Discount, Tax_Rate, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Inventory d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccount) GetAccountInventorybyUuid(ctx context.Context, uuid uuid.UUID) (AccountInventoryInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccountInventorybyUuid, uuid)
	var i AccountInventoryInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountId,
		&i.BarCode,
		&i.Code,
		&i.Quantity,
		&i.UnitPrice,
		&i.BookValue,
		&i.Discount,
		&i.TaxRate,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAccountInventorybyName = `-- name: GetAccountInventorybyName :one
SELECT 
Id, mr.UUID, Account_ID, Bar_Code, Code, Quantity, Unit_Price, 
Book_Value, Discount, Tax_Rate, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Inventory d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Title = $1 LIMIT 1
`

func (q *QueriesAccount) GetAccountInventorybyName(ctx context.Context, name string) (AccountInventoryInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccountInventorybyName, name)
	var i AccountInventoryInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountId,
		&i.BarCode,
		&i.Code,
		&i.Quantity,
		&i.UnitPrice,
		&i.BookValue,
		&i.Discount,
		&i.TaxRate,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listAccountInventory = `-- name: ListAccountInventory:many
SELECT 
  Id, mr.UUID, Account_ID, Bar_Code, Code, Quantity, Unit_Price, 
  Book_Value, Discount, Tax_Rate, Remarks, Other_Info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Inventory d 
INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Account_Id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListAccountInventoryParams struct {
	AccountId int64 `json:"accountId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountInventory(ctx context.Context, arg ListAccountInventoryParams) ([]AccountInventoryInfo, error) {
	rows, err := q.db.QueryContext(ctx, listAccountInventory, arg.AccountId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountInventoryInfo{}
	for rows.Next() {
		var i AccountInventoryInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.AccountId,
			&i.BarCode,
			&i.Code,
			&i.Quantity,
			&i.UnitPrice,
			&i.BookValue,
			&i.Discount,
			&i.TaxRate,
			&i.Remarks,
			&i.OtherInfo,
			&i.ModCtr,
			&i.Created,
			&i.Updated,
		); err != nil {
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

const updateAccountInventory = `-- name: UpdateAccountInventory :one
UPDATE Account_Inventory SET 
	Account_ID = $2,
	Bar_Code = $3,
	Code = $4,
	Quantity = $5,
	Unit_Price = $6,
	Book_Value = $7,
	Discount = $8,
	Tax_Rate = $9,
	Remarks = $10,
	Other_Info = $11
WHERE id = $1
RETURNING Id, UUID, Account_ID, Bar_Code, Code, Quantity, Unit_Price, 
Book_Value, Discount, Tax_Rate, Remarks, Other_Info 
`

func (q *QueriesAccount) UpdateAccountInventory(ctx context.Context, arg AccountInventoryRequest) (model.AccountInventory, error) {
	row := q.db.QueryRowContext(ctx, updateAccountInventory,
		arg.Id,
		arg.AccountId,
		arg.BarCode,
		arg.Code,
		arg.Quantity,
		arg.UnitPrice,
		arg.BookValue,
		arg.Discount,
		arg.TaxRate,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.AccountInventory
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.AccountId,
		&i.BarCode,
		&i.Code,
		&i.Quantity,
		&i.UnitPrice,
		&i.BookValue,
		&i.Discount,
		&i.TaxRate,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
