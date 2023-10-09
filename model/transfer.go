package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transfer struct {
	Id            int64 `json:"Id"`
	FromAccountId int64 `json:"fromAccountId"`
	ToAccountId   int64 `json:"toAccountId"`
	// must be positive
	Amount    decimal.Decimal `json:"amount"`
	CreatedAt time.Time       `json:"createdAt"`
}
