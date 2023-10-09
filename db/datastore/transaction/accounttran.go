package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
	"simplebank/util"

	"github.com/shopspring/decimal"
)

const createAccountTran = `-- name: CreateAccountTran: one
INSERT INTO Account_Tran (
  UUID, Trn_Head_Id, Series, Alternate_Key, Value_Date, Account_Id, Trn_Type_Code, Currency, 
  Item_Id, Passbook_Posted, Trn_Prin, Trn_Int, Bal_Prin, Bal_Int, Cancelled, Other_Info
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) 
ON CONFLICT (Trn_Head_ID, Series) DO UPDATE SET
  Alternate_Key = EXCLUDED.Alternate_Key, 
  Value_Date = EXCLUDED.Value_Date, 
  Account_Id = EXCLUDED.Account_Id, 
  Trn_Type_Code = EXCLUDED.Trn_Type_Code, 
  Currency = EXCLUDED.Currency, 
  Item_Id = EXCLUDED.Item_Id, 
  Passbook_Posted = EXCLUDED.Passbook_Posted, 
  Trn_Prin = EXCLUDED.Trn_Prin, 
  Trn_Int = EXCLUDED.Trn_Int, 
  Bal_Prin = EXCLUDED.Bal_Prin, 
  Bal_Int = EXCLUDED.Bal_Int, 
  Cancelled = EXCLUDED.Cancelled, 
  Other_Info = EXCLUDED.Other_Info
  RETURNING 
    UUID, Trn_Head_Id, Series, Alternate_Key, Value_Date, Account_Id, Trn_Type_Code, Currency, 
    Item_Id, Passbook_Posted, Trn_Prin, Trn_Int, Bal_Prin, Bal_Int, Cancelled, Other_Info
`

type AccountTranRequest struct {
	Uuid           uuid.UUID       `json:"uuid"`
	TrnHeadId      int64           `json:"trnHeadId"`
	Series         int64           `json:"series"`
	AlternateKey   sql.NullString  `json:"alternateKey"`
	ValueDate      time.Time       `json:"valueDate"`
	AccountId      int64           `json:"accountId"`
	TrnTypeCode    int64           `json:"trnTypeCode"`
	Currency       string          `json:"currency"`
	ItemId         sql.NullInt64   `json:"itemid"`
	ItemCode       int64           `json:"ItemCode"`
	PassbookPosted bool            `json:"passbookPosted"`
	TrnPrin        decimal.Decimal `json:"trnPrin"`
	TrnInt         decimal.Decimal `json:"trnInt"`
	BalPrin        decimal.Decimal `json:"balPrin"`
	BalInt         decimal.Decimal `json:"balInt"`
	Cancelled      bool            `json:"cancelled"`
	OtherInfo      sql.NullString  `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateAccountTran(ctx context.Context, arg AccountTranRequest) (model.AccountTran, error) {
	if arg.Uuid == uuid.Nil {
		arg.Uuid = util.UUID()
	}

	row := q.db.QueryRowContext(ctx, createAccountTran,
		arg.Uuid,
		arg.TrnHeadId,
		arg.Series,
		arg.AlternateKey,
		arg.ValueDate,
		arg.AccountId,
		arg.TrnTypeCode,
		arg.Currency,
		arg.ItemId,
		arg.PassbookPosted,
		arg.TrnPrin,
		arg.TrnInt,
		arg.BalPrin,
		arg.BalInt,
		arg.Cancelled,
		arg.OtherInfo,
	)
	var i model.AccountTran
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.AlternateKey,
		&i.ValueDate,
		&i.AccountId,
		&i.TrnTypeCode,
		&i.Currency,
		&i.ItemId,
		&i.PassbookPosted,
		&i.TrnPrin,
		&i.TrnInt,
		&i.BalPrin,
		&i.BalInt,
		&i.Cancelled,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccountTran = `-- name: DeleteAccountTran :exec
DELETE FROM Account_Tran
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteAccountTran(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountTran, uuid)
	return err
}

type AccountTranInfo struct {
	Uuid           uuid.UUID       `json:"uuid"`
	TrnHeadId      int64           `json:"trnHeadId"`
	Series         int64           `json:"series"`
	AlternateKey   sql.NullString  `json:"alternateKey"`
	ValueDate      time.Time       `json:"valueDate"`
	AccountId      int64           `json:"accountId"`
	TrnTypeCode    int64           `json:"trnTypeCode"`
	Currency       string          `json:"currency"`
	ItemId         sql.NullInt64   `json:"itemid"`
	ItemCode       int64           `json:"ItemCode"`
	PassbookPosted bool            `json:"passbookPosted"`
	TrnPrin        decimal.Decimal `json:"trnPrin"`
	TrnInt         decimal.Decimal `json:"trnInt"`
	BalPrin        decimal.Decimal `json:"balPrin"`
	BalInt         decimal.Decimal `json:"balInt"`
	Cancelled      bool            `json:"cancelled"`
	OtherInfo      sql.NullString  `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getAccountTran = `-- name: GetAccountTran :one
SELECT 
  mr.UUID, Trn_Head_Id, Series, Alternate_Key, Value_Date, Account_Id, Trn_Type_Code, Currency, 
  Item_Id, Passbook_Posted, Trn_Prin, Trn_Int, Bal_Prin, Bal_Int, Cancelled, Other_Info,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Tran d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateAccountTran(q *QueriesTransaction, ctx context.Context, sql string, param ...interface{}) ([]AccountTranInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql, param...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountTranInfo{}
	for rows.Next() {
		var i AccountTranInfo
		err := rows.Scan(
			&i.Uuid,
			&i.TrnHeadId,
			&i.Series,
			&i.AlternateKey,
			&i.ValueDate,
			&i.AccountId,
			&i.TrnTypeCode,
			&i.Currency,
			&i.ItemId,
			&i.PassbookPosted,
			&i.TrnPrin,
			&i.TrnInt,
			&i.BalPrin,
			&i.BalInt,
			&i.Cancelled,
			&i.OtherInfo,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
		)
		if err != nil {
			return items, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return items, err
	}
	if err := rows.Err(); err != nil {
		return items, err
	}
	return items, nil
}

func (q *QueriesTransaction) GetAccountTranbyAcc(ctx context.Context, acc string) ([]AccountTranInfo, error) {
	sql := fmt.Sprintf(`%v WHERE acc.Acc = %s `, getAccountTran, acc)
	return populateAccountTran(q, ctx, sql)
}

func (q *QueriesTransaction) GetAccountTranbyUuid(ctx context.Context, uuid uuid.UUID) (AccountTranInfo, error) {
	sql := fmt.Sprintf(`%v WHERE mr.UUID = $1 `, getAccountTran)
	items, err := populateAccountTran(q, ctx, sql, uuid)

	log.Printf("----------- items: %v err:%v", items, err)
	for _, val := range items {
		return val, err
	}
	return AccountTranInfo{}, fmt.Errorf("schedule UUID:%v not found", uuid)
}

func (q *QueriesTransaction) GetAccountTran(ctx context.Context, trnHeadId int64, series int64) (AccountTranInfo, error) {
	sql := fmt.Sprintf(`%v WHERE d.Trn_Head_Id = $1 and d.Series = $2`, getAccountTran)
	items, err := populateAccountTran(q, ctx, sql, trnHeadId, series)

	log.Printf("----------- items: %v err:%v", items, err)
	for _, val := range items {
		return val, err
	}
	return AccountTranInfo{}, fmt.Errorf("schedule trnHeadId:%v series:%v not found", trnHeadId, series)
}

type ListAccountTranParams struct {
	TrnHeadId int64 `json:"trnHeadId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesTransaction) ListAccountTran(ctx context.Context, arg ListAccountTranParams) ([]AccountTranInfo, error) {
	sql := fmt.Sprintf(`%v WHERE trn_head_id = $1
	ORDER BY uuid LIMIT $2 OFFSET $3`, getAccountTran)
	return populateAccountTran(q, ctx, sql, arg.TrnHeadId, arg.Limit, arg.Offset)
}

const updateAccountTran = `-- name: UpdateAccountTran :one
UPDATE Account_Tran SET 
  Trn_Head_Id = $2,
  Series = $3,
  Alternate_Key = $4,
  Value_Date = $5,
  Account_Id = $6,
  Trn_Type_Code = $7,
  Currency = $8,
  Item_Id = $9,
  Passbook_Posted = $10,
  Trn_Prin = $11,
  Trn_Int = $12,
  Bal_Prin = $13,
  Bal_Int = $14,
  Cancelled = $15,
  Other_Info = $16
WHERE uuid = $1
RETURNING 
  UUID, Trn_Head_Id, Series, Alternate_Key, Value_Date, Account_Id, Trn_Type_Code, Currency, 
  Item_Id, Passbook_Posted, Trn_Prin, Trn_Int, Bal_Prin, Bal_Int, Cancelled, Other_Info
`

func (q *QueriesTransaction) UpdateAccountTran(ctx context.Context, arg AccountTranRequest) (model.AccountTran, error) {
	row := q.db.QueryRowContext(ctx, updateAccountTran,

		arg.Uuid,
		arg.TrnHeadId,
		arg.Series,
		arg.AlternateKey,
		arg.ValueDate,
		arg.AccountId,
		arg.TrnTypeCode,
		arg.Currency,
		arg.ItemId,
		arg.PassbookPosted,
		arg.TrnPrin,
		arg.TrnInt,
		arg.BalPrin,
		arg.BalInt,
		arg.Cancelled,
		arg.OtherInfo,
	)
	var i model.AccountTran
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.AlternateKey,
		&i.ValueDate,
		&i.AccountId,
		&i.TrnTypeCode,
		&i.Currency,
		&i.ItemId,
		&i.PassbookPosted,
		&i.TrnPrin,
		&i.TrnInt,
		&i.BalPrin,
		&i.BalInt,
		&i.Cancelled,
		&i.OtherInfo,
	)
	return i, err
}
