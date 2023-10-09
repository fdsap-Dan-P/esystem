package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccountInterest struct {
	AccountId       int64           `json:"accountId"`
	Uuid            uuid.UUID       `json:"uuid"`
	Interest        decimal.Decimal `json:"interest"`
	EffectiveRate   decimal.Decimal `json:"effectiveRate"`
	InterestRate    decimal.Decimal `json:"interestRate"`
	Credit          decimal.Decimal `json:"credit"`
	Debit           decimal.Decimal `json:"debit"`
	Accruals        decimal.Decimal `json:"accruals"`
	WaivedInt       decimal.Decimal `json:"waivedInt"`
	LastAccruedDate sql.NullTime    `json:"lastAccruedDate"`
}
