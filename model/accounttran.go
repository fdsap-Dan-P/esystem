package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccountTran struct {
	Uuid           uuid.UUID       `json:"uuid"`
	TrnHeadId      int64           `json:"trnHeadId"`
	Series         int64           `json:"series"`
	AlternateKey   sql.NullString  `json:"alternateKey"`
	ValueDate      time.Time       `json:"valueDate"`
	AccountId      int64           `json:"accountId"`
	TrnTypeCode    int64           `json:"trnTypeCode"`
	Currency       string          `json:"currency"`
	ItemId         sql.NullInt64   `json:"itemId"`
	PassbookPosted bool            `json:"passbookPosted"`
	TrnPrin        decimal.Decimal `json:"trnPrin"`
	TrnInt         decimal.Decimal `json:"trnInt"`
	BalPrin        decimal.Decimal `json:"balPrin"`
	BalInt         decimal.Decimal `json:"balInt"`
	Cancelled      bool            `json:"cancelled"`
	OtherInfo      sql.NullString  `json:"otherInfo"`
}
