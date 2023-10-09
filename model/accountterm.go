package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccountTerm struct {
	AccountId      int64           `json:"accountId"`
	Uuid           uuid.UUID       `json:"uuid"`
	Frequency      int16           `json:"frequency"`
	N              int16           `json:"n"`
	PaidN          int16           `json:"paidN"`
	FixedDue       decimal.Decimal `json:"fixedDue"`
	CummulativeDue decimal.Decimal `json:"cummulativeDue"`
	DateStart      time.Time       `json:"dateStart"`
	Maturity       time.Time       `json:"maturity"`
}
