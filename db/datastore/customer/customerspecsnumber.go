package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createCustomerSpecsNumber = `-- name: CreateCustomerSpecsNumber: one
INSERT INTO Customer_Specs_Number 
  (Customer_Id, Specs_ID, Value, Value2, Measure_Id) 
VALUES 
  ($1, $2, $3, $4, $5) 
ON CONFLICT( Customer_Id, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Customer_Id, Specs_Code, Specs_ID, Value, Value2, Measure_Id
`

type CustomerSpecsNumberRequest struct {
	Uuid       uuid.UUID       `json:"uuid"`
	CustomerId int64           `json:"customerId"`
	SpecsId    int64           `json:"specsId"`
	Value      decimal.Decimal `json:"value"`
	Value2     decimal.Decimal `json:"value2"`
	MeasureId  sql.NullInt64   `json:"measureId"`
}

func (q *QueriesCustomer) CreateCustomerSpecsNumber(ctx context.Context, arg CustomerSpecsNumberRequest) (model.CustomerSpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, createCustomerSpecsNumber,
		arg.CustomerId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.CustomerSpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const updateCustomerSpecsNumber = `-- name: UpdateCustomerSpecsNumber :one
UPDATE Customer_Specs_Number SET 
  Customer_Id = $2,
  Specs_ID = $3,
  Value = $4,
  Value2 = $5,
  Measure_Id = $6
WHERE uuid = $1
RETURNING UUID, Customer_Id, Specs_Code, Specs_ID, Value, Value2
`

func (q *QueriesCustomer) UpdateCustomerSpecsNumber(ctx context.Context, arg CustomerSpecsNumberRequest) (model.CustomerSpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, updateCustomerSpecsNumber,
		arg.Uuid,
		arg.CustomerId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.CustomerSpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const deleteCustomerSpecsNumber = `-- name: DeleteCustomerSpecsNumber :exec
DELETE FROM Customer_Specs_Number
WHERE uuid = $1
`

func (q *QueriesCustomer) DeleteCustomerSpecsNumber(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCustomerSpecsNumber, uuid)
	return err
}

type CustomerSpecsNumberInfo struct {
	Uuid            uuid.UUID       `json:"uuid"`
	CustomerId      int64           `json:"customerId"`
	SpecsCode       string          `json:"specsCode"`
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

func populateIdentitySpecNumber(q *QueriesCustomer, ctx context.Context, sql string) (CustomerSpecsNumberInfo, error) {
	var i CustomerSpecsNumberInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
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

func populateIdentitySpecNumber2(q *QueriesCustomer, ctx context.Context, sql string) ([]CustomerSpecsNumberInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []CustomerSpecsNumberInfo{}
	for rows.Next() {
		var i CustomerSpecsNumberInfo
		err := rows.Scan(
			&i.Uuid,
			&i.CustomerId,
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

const customerSpecsNumberSQL = `-- name: customerSpecsNumberSQL
SELECT 
  mr.UUID, d.Customer_Id, d.Specs_Code, d.Specs_ID, ref.Title Item, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2, d.Measure_Id, mea.Title Measure, mea.Short_Name Measure_Unit, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Specs_Number d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
LEFT JOIN Reference mea on mea.ID = Measure_Id`

func (q *QueriesCustomer) GetCustomerSpecsNumber(ctx context.Context, customerId int64, specsId int64) (CustomerSpecsNumberInfo, error) {
	return populateIdentitySpecNumber(q, ctx, fmt.Sprintf("%s WHERE d.Customer_Id = %v and d.Specs_ID = %v",
		customerSpecsNumberSQL, customerId, specsId))
}

func (q *QueriesCustomer) GetCustomerSpecsNumberbyUuid(ctx context.Context, uuid uuid.UUID) (CustomerSpecsNumberInfo, error) {
	return populateIdentitySpecNumber(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", customerSpecsNumberSQL, uuid))
}

type ListCustomerSpecsNumberParams struct {
	CustomerId int64 `json:"CustomerId"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *QueriesCustomer) ListCustomerSpecsNumber(ctx context.Context, arg ListCustomerSpecsNumberParams) ([]CustomerSpecsNumberInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Customer_Id = %v LIMIT %d OFFSET %d",
			customerSpecsNumberSQL, arg.CustomerId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Customer_Id = %v ", customerSpecsNumberSQL, arg.CustomerId)
	}
	return populateIdentitySpecNumber2(q, ctx, sql)
}
