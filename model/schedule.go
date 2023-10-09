package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Schedule struct {
	Id            int64           `json:"id"`
	Uuid          uuid.UUID       `json:"uuid"`
	AccountId     int64           `json:"accountId"`
	Series        int16           `json:"series"`
	DueDate       time.Time       `json:"dueDate"`
	DuePrin       decimal.Decimal `json:"duePrin"`
	DueInt        decimal.Decimal `json:"dueInt"`
	EndPrin       decimal.Decimal `json:"endPrin"`
	EndInt        decimal.Decimal `json:"endInt"`
	CarryingValue decimal.Decimal `json:"carryingValue"`
	Realizable    decimal.Decimal `json:"realizable"`
	OtherInfo     sql.NullString  `json:"otherInfo"`
}
