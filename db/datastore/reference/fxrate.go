package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createFxrate = `-- name: CreateFxrate: one
INSERT INTO fxRate (
Buy_Rate, Cutof_Date, Sell_Rate, Base_Currency, Currency, 
Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6
)
ON CONFLICT(Base_Currency,   Currency,   Cutof_Date)
  DO UPDATE SET 
	   Buy_Rate = EXCLUDED.Buy_Rate, 
		 Sell_Rate = EXCLUDED.Sell_Rate
RETURNING UUId, Buy_Rate, Cutof_Date, Sell_Rate, Base_Currency, Currency, 
Other_Info
`

type FxrateRequest struct {
	Uuid         uuid.UUID       `json:"uuid"`
	BuyRate      decimal.Decimal `json:"buyRate"`
	CutofDate    time.Time       `json:"cutofDate"`
	SellRate     decimal.Decimal `json:"sellRate"`
	BaseCurrency string          `json:"baseCurrency"`
	Currency     string          `json:"currency"`
	OtherInfo    sql.NullString  `json:"otherInfo"`
}

func (q *QueriesReference) CreateFxrate(ctx context.Context, arg FxrateRequest) (model.Fxrate, error) {
	row := q.db.QueryRowContext(ctx, createFxrate,
		arg.BuyRate,
		arg.CutofDate,
		arg.SellRate,
		arg.BaseCurrency,
		arg.Currency,
		arg.OtherInfo,
	)
	var i model.Fxrate
	err := row.Scan(
		&i.Uuid,
		&i.BuyRate,
		&i.CutofDate,
		&i.SellRate,
		&i.BaseCurrency,
		&i.Currency,
		&i.OtherInfo,
	)
	return i, err
}

const deleteFxrate = `-- name: DeleteFxrate :exec
DELETE FROM fxRate
WHERE uuid = $1
`

func (q *QueriesReference) DeleteFxrate(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteFxrate, uuid)
	return err
}

type FxrateInfo struct {
	Uuid         uuid.UUID       `json:"uuid"`
	BuyRate      decimal.Decimal `json:"buyRate"`
	CutofDate    time.Time       `json:"cutofDate"`
	SellRate     decimal.Decimal `json:"sellRate"`
	BaseCurrency string          `json:"baseCurrency"`
	Currency     string          `json:"currency"`
	OtherInfo    sql.NullString  `json:"otherInfo"`
	ModCtr       int64           `json:"modCtr"`
	Created      sql.NullTime    `json:"created"`
	Updated      sql.NullTime    `json:"updated"`
}

const getFxrate = `-- name: GetFxrate :one
SELECT 
mr.UUId, Buy_Rate, 
Cutof_Date, Sell_Rate, Base_Currency, Currency, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM fxRate d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Base_Currency = $1 and Currency = $2 and Cutof_Date = $3 LIMIT 1
`

func (q *QueriesReference) GetFxrate(ctx context.Context, baseCurrency string, currency string, cutDate time.Time) (FxrateInfo, error) {
	row := q.db.QueryRowContext(ctx, getFxrate, baseCurrency, currency, cutDate)
	var i FxrateInfo
	err := row.Scan(
		&i.Uuid,
		&i.BuyRate,
		&i.CutofDate,
		&i.SellRate,
		&i.BaseCurrency,
		&i.Currency,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getFxratebyUuId = `-- name: GetFxratebyUuId :one
SELECT 
mr.UUId, Buy_Rate, 
Cutof_Date, Sell_Rate, Base_Currency, Currency, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM fxRate d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesReference) GetFxratebyUuId(ctx context.Context, uuid uuid.UUID) (FxrateInfo, error) {
	row := q.db.QueryRowContext(ctx, getFxratebyUuId, uuid)
	var i FxrateInfo
	err := row.Scan(
		&i.Uuid,
		&i.BuyRate,
		&i.CutofDate,
		&i.SellRate,
		&i.BaseCurrency,
		&i.Currency,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listFxrate = `-- name: ListFxrate:many
SELECT 
mr.UUId, Buy_Rate, 
Cutof_Date, Sell_Rate, Base_Currency, Currency, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM fxRate d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Cutof_Date = $1
ORDER BY Base_Currency
LIMIT $2
OFFSET $3
`

type ListFxrateParams struct {
	CutDate time.Time `json:"cutDate"`
	Limit   int32     `json:"limit"`
	Offset  int32     `json:"offset"`
}

func (q *QueriesReference) ListFxrate(ctx context.Context, arg ListFxrateParams) ([]FxrateInfo, error) {
	rows, err := q.db.QueryContext(ctx, listFxrate, arg.CutDate, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FxrateInfo{}
	for rows.Next() {
		var i FxrateInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.BuyRate,
			&i.CutofDate,
			&i.SellRate,
			&i.BaseCurrency,
			&i.Currency,
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

const updateFxrate = `-- name: UpdateFxrate :one
UPDATE fxRate SET 
Buy_Rate = $2,
Cutof_Date = $3,
Sell_Rate = $4,
Base_Currency = $5,
Currency = $6,
Other_Info = $7
WHERE uuid = $1
RETURNING UUId, Buy_Rate, Cutof_Date, Sell_Rate, Base_Currency, Currency, 
Other_Info
`

func (q *QueriesReference) UpdateFxrate(ctx context.Context, arg FxrateRequest) (model.Fxrate, error) {
	row := q.db.QueryRowContext(ctx, updateFxrate,
		arg.Uuid,
		arg.BuyRate,
		arg.CutofDate,
		arg.SellRate,
		arg.BaseCurrency,
		arg.Currency,
		arg.OtherInfo,
	)
	var i model.Fxrate
	err := row.Scan(
		&i.Uuid,
		&i.BuyRate,
		&i.CutofDate,
		&i.SellRate,
		&i.BaseCurrency,
		&i.Currency,
		&i.OtherInfo,
	)
	return i, err
}
