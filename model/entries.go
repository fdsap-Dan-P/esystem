package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Entries struct {
	Id        int64           `json:"id"`
	AccountId int64           `json:"accountId"`
	Amount    decimal.Decimal `json:"amount"`
	CreatedAt time.Time       `json:"createdAt"`
}
